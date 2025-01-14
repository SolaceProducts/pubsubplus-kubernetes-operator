/*
Copyright 2023 Solace Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"fmt"
	"strconv"
	"strings"

	eventbrokerv1beta1 "github.com/SolaceProducts/pubsubplus-operator/api/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *PubSubPlusEventBrokerReconciler) newDeploymentForPrometheusExporter(monitoringDeploymentName string, monitoringSecret *corev1.Secret, m *eventbrokerv1beta1.PubSubPlusEventBroker) *appsv1.Deployment {
	monitoringDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      monitoringDeploymentName,
			Namespace: m.Namespace,
			Labels:    getObjectLabels(m.Name),
		},
	}
	r.updateDeploymentForPrometheusExporter(monitoringDeployment, monitoringSecret, m)
	// Set PubSubPlusEventBroker instance as the owner and controller
	ctrl.SetControllerReference(m, monitoringDeployment, r.Scheme)
	return monitoringDeployment
}

func (r *PubSubPlusEventBrokerReconciler) updateDeploymentForPrometheusExporter(monitoringDeployment *appsv1.Deployment, monitoringSecret *corev1.Secret, eventBroker *eventbrokerv1beta1.PubSubPlusEventBroker) *appsv1.Deployment {
	if monitoringDeployment.Annotations == nil {
		monitoringDeployment.Annotations = map[string]string{}
	}
	monitoringDeployment.Annotations[monitoringSpecSignatureAnnotationName] = monitoringSpecHash(eventBroker.Spec)
	monitoringDeployment.Spec = appsv1.DeploymentSpec{
		Selector: &metav1.LabelSelector{
			MatchLabels: getMonitoringDeploymentSelector(eventBroker.Name),
		},
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: getMonitoringDeploymentSelector(eventBroker.Name),
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:            "exporter",
						Image:           r.getExporterImageDetails(eventBroker.Spec.Monitoring.MonitoringImage),
						ImagePullPolicy: getExporterImagePullPolicy(eventBroker.Spec.Monitoring.MonitoringImage),
						Ports: []corev1.ContainerPort{{
							Name:          getExporterHttpProtocolType(&eventBroker.Spec.Monitoring),
							ContainerPort: getExporterContainerPort(&eventBroker.Spec.Monitoring),
						}},

						Env: []corev1.EnvVar{
							{
								Name:  monitoringExporterListenAddress,
								Value: fmt.Sprintf("%s://%s.%s.svc.cluster.local:%d", getExporterHttpProtocolType(&eventBroker.Spec.Monitoring), getObjectName("PrometheusExporterService", eventBroker.Name), eventBroker.Namespace, getExporterContainerPort(&eventBroker.Spec.Monitoring)),
							},
							{
								Name:  monitoringExporterScrapeURI,
								Value: fmt.Sprintf("%s://%s.%s.svc.cluster.local:%d", getPubSubPlusEventBrokerProtocol(&eventBroker.Spec), getObjectName("BrokerService", eventBroker.Name), eventBroker.Namespace, getPubSubPlusEventBrokerPort(&eventBroker.Spec.Service, &eventBroker.Spec.BrokerTLS)),
							},
							{
								Name:  monitoringExporterListenTLS,
								Value: getExporterTLSConfiguration(&eventBroker.Spec.Monitoring),
							},
							{
								Name:  monitoringExporterBrokerUsername,
								Value: "monitor",
							},
							{
								Name: monitoringExporterBrokerPassword,
								ValueFrom: &corev1.EnvVarSource{
									SecretKeyRef: &corev1.SecretKeySelector{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: monitoringSecret.Name,
										},
										Key: monitorSecretKeyName,
									},
								},
							},
							{
								Name:  monitoringExporterScrapeTimeout,
								Value: fmt.Sprint(eventBroker.Spec.Monitoring.TimeOut) + "s",
							},
							{
								Name:  monitoringExporterSSLVerify,
								Value: strconv.FormatBool(eventBroker.Spec.Monitoring.SSLVerify),
							},
							{
								Name:  monitoringExporterIncludeRates,
								Value: strconv.FormatBool(eventBroker.Spec.Monitoring.IncludeRates),
							},
						},
						SecurityContext: &corev1.SecurityContext{
							Privileged: &[]bool{false}[0], // Set to false
							Capabilities: &corev1.Capabilities{
								Drop: []corev1.Capability{
									corev1.Capability("ALL"),
								},
							},
							RunAsNonRoot:             &[]bool{true}[0],  // Set to true
							AllowPrivilegeEscalation: &[]bool{false}[0], // Set to false
							SeccompProfile: &corev1.SeccompProfile{
								Type: corev1.SeccompProfileTypeRuntimeDefault,
							},
						},
						Resources: eventBroker.Spec.Monitoring.Resources,
					},
				},
				ImagePullSecrets: getExporterImagePullSecrets(eventBroker.Spec.Monitoring.MonitoringImage),
			},
		},
		Strategy: appsv1.DeploymentStrategy{
			Type: appsv1.RollingUpdateDeploymentStrategyType,
		},
	}

	// OpenShift and namespace must be considered for RunAsUser
	// Only set it if not on OpenShift or using the "default" namespace
	// otherwise leave it undefined
	if !r.IsOpenShift || eventBroker.Namespace == corev1.NamespaceDefault {
		monitoringDeployment.Spec.Template.Spec.SecurityContext = &corev1.PodSecurityContext{
			RunAsUser:  &[]int64{10001}[0],
			RunAsGroup: &[]int64{10001}[0],
			FSGroup:    &[]int64{10002}[0],
		}
	}

	//Set TLS configuration
	if eventBroker.Spec.Monitoring.MonitoringMetricsEndpoint != nil && len(strings.TrimSpace(eventBroker.Spec.Monitoring.MonitoringMetricsEndpoint.EndpointTLSConfigSecret)) > 0 {
		allVolumes := monitoringDeployment.Spec.Template.Spec.Volumes
		allVolumes = append(allVolumes, corev1.Volume{
			Name: "server-certs",
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName:  eventBroker.Spec.Monitoring.MonitoringMetricsEndpoint.EndpointTLSConfigSecret,
					DefaultMode: &[]int32{0400}[0],
				},
			},
		})
		allContainerVolumeMounts := monitoringDeployment.Spec.Template.Spec.Containers[0].VolumeMounts
		allContainerVolumeMounts = append(allContainerVolumeMounts, corev1.VolumeMount{
			Name:      "server-certs",
			MountPath: "/mnt/disks/solace",
			ReadOnly:  true,
		})
		allEnv := monitoringDeployment.Spec.Template.Spec.Containers[0].Env

		allEnv = append(allEnv, corev1.EnvVar{
			Name:  monitoringExporterServerCert,
			Value: "/mnt/disks/solace/" + eventBroker.Spec.Monitoring.MonitoringMetricsEndpoint.EndpointTlsConfigServerCertName,
		})
		allEnv = append(allEnv, corev1.EnvVar{
			Name:  monitoringExporterPrivateKey,
			Value: "/mnt/disks/solace/" + eventBroker.Spec.Monitoring.MonitoringMetricsEndpoint.EndpointTlsConfigPrivateKeyName,
		})

		monitoringDeployment.Spec.Template.Spec.Containers[0].Env = allEnv
		monitoringDeployment.Spec.Template.Spec.Volumes = allVolumes
		monitoringDeployment.Spec.Template.Spec.Containers[0].VolumeMounts = allContainerVolumeMounts
	} else {
		allEnv := monitoringDeployment.Spec.Template.Spec.Containers[0].Env
		allEnv = append(allEnv, corev1.EnvVar{
			Name:  monitoringExporterServerCert,
			Value: ".", //This is a mandatory parameter for older versions of the exporter.
		})
		allEnv = append(allEnv, corev1.EnvVar{
			Name:  monitoringExporterPrivateKey,
			Value: ".", //This is a mandatory parameter for older versions of the exporter.
		})
		monitoringDeployment.Spec.Template.Spec.Containers[0].Env = allEnv
	}

	//Set Extra environment variables
	if len(eventBroker.Spec.Monitoring.ExtraEnvVars) > 0 {
		allEnv := monitoringDeployment.Spec.Template.Spec.Containers[0].Env
		for _, envV := range eventBroker.Spec.Monitoring.ExtraEnvVars {
			if strings.ToUpper(envV.Name) != monitoringExporterIncludeRates &&
				strings.ToUpper(envV.Name) != monitoringExporterPrivateKey &&
				strings.ToUpper(envV.Name) != monitoringExporterServerCert &&
				strings.ToUpper(envV.Name) != monitoringExporterBrokerUsername &&
				strings.ToUpper(envV.Name) != monitoringExporterBrokerPassword &&
				strings.ToUpper(envV.Name) != monitoringExporterListenAddress &&
				strings.ToUpper(envV.Name) != monitoringExporterListenTLS &&
				strings.ToUpper(envV.Name) != monitoringExporterScrapeTimeout &&
				strings.ToUpper(envV.Name) != monitoringExporterScrapeURI &&
				strings.ToUpper(envV.Name) != monitoringExporterSSLVerify {
				allEnv = append(allEnv, corev1.EnvVar{
					Name:  envV.Name,
					Value: envV.Value,
				})
			}
		}
		monitoringDeployment.Spec.Template.Spec.Containers[0].Env = allEnv
	}

	return monitoringDeployment
}

func (r *PubSubPlusEventBrokerReconciler) newServiceForPrometheusExporter(exporter *eventbrokerv1beta1.Monitoring, svcName string, m *eventbrokerv1beta1.PubSubPlusEventBroker) *corev1.Service {
	dep := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      svcName,
			Namespace: m.Namespace,
			Labels:    getMonitoringDeploymentSelector(m.Name),
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       getExporterHttpProtocolType(&m.Spec.Monitoring),
					Protocol:   getExporterServiceProtocol(&m.Spec.Monitoring),
					Port:       getExporterServicePort(exporter),
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: getExporterContainerPort(exporter)},
				},
			},
			Selector: getMonitoringDeploymentSelector(m.Name),
		},
	}
	if exporter.MonitoringMetricsEndpoint != nil && exporter.MonitoringMetricsEndpoint.ServiceType != "" {
		dep.Spec.Type = exporter.MonitoringMetricsEndpoint.ServiceType
	} else {
		dep.Spec.Type = corev1.ServiceTypeClusterIP
	}
	// Set PubSubPlusEventBroker instance as the owner and controller
	ctrl.SetControllerReference(m, dep, r.Scheme)
	return dep
}
func getPubSubPlusEventBrokerPort(m *eventbrokerv1beta1.Service, b *eventbrokerv1beta1.BrokerTLS) int32 {
	if len(m.Ports) == 0 {
		if b.Enabled {
			return 1943
		}
		return 8080
	}
	for i := range m.Ports {
		if b.Enabled {
			if m.Ports[i].Name == tlsSempPortName {
				return m.Ports[i].ContainerPort
			}
		} else {
			if m.Ports[i].Name == tcpSempPortName {
				return m.Ports[i].ContainerPort
			}
		}
	}
	return 0
}

func getPubSubPlusEventBrokerProtocol(m *eventbrokerv1beta1.EventBrokerSpec) string {
	if m.BrokerTLS.Enabled {
		return "https"
	}
	return "http"
}

func getExporterContainerPort(m *eventbrokerv1beta1.Monitoring) int32 {
	if m.MonitoringMetricsEndpoint == nil || m.MonitoringMetricsEndpoint.ContainerPort == 0 {
		return 9628
	}
	return m.MonitoringMetricsEndpoint.ContainerPort
}

func getExporterServicePort(m *eventbrokerv1beta1.Monitoring) int32 {
	if m.MonitoringMetricsEndpoint == nil || m.MonitoringMetricsEndpoint.ServicePort == 0 {
		return 9628
	}
	return m.MonitoringMetricsEndpoint.ServicePort
}

func getExporterHttpProtocolType(m *eventbrokerv1beta1.Monitoring) string {
	if m.MonitoringMetricsEndpoint != nil && len(strings.TrimSpace(m.MonitoringMetricsEndpoint.Name)) > 0 {
		return m.MonitoringMetricsEndpoint.Name
	} else if m.MonitoringMetricsEndpoint != nil && m.MonitoringMetricsEndpoint.ListenTLS {
		return "tls-metrics"
	}
	return "tcp-metrics"
}

func getExporterServiceProtocol(m *eventbrokerv1beta1.Monitoring) corev1.Protocol {
	if m.MonitoringMetricsEndpoint == nil || m.MonitoringMetricsEndpoint.Protocol == "" {
		return corev1.ProtocolTCP
	}
	return m.MonitoringMetricsEndpoint.Protocol
}

func getExporterTLSConfiguration(m *eventbrokerv1beta1.Monitoring) string {
	if m.MonitoringMetricsEndpoint == nil || !m.MonitoringMetricsEndpoint.ListenTLS {
		return "false"
	}
	return strconv.FormatBool(m.MonitoringMetricsEndpoint.ListenTLS)
}

func (r *PubSubPlusEventBrokerReconciler) getExporterImageDetails(bm *eventbrokerv1beta1.MonitoringImage) string {
	imageRepo := (map[bool]string{true: DefaultExporterImageRepoOpenShift, false: DefaultExporterImageRepoK8s})[r.IsOpenShift]
	imageTag := (map[bool]string{true: DefaultExporterImageTagOpenShift, false: DefaultExporterImageTagK8s})[r.IsOpenShift]

	if bm != nil && len(strings.TrimSpace(bm.Repository)) > 0 {
		imageRepo = bm.Repository
	}
	if bm != nil && len(strings.TrimSpace(bm.Tag)) > 0 {
		imageTag = bm.Tag
	}
	return imageRepo + ":" + imageTag
}

func getExporterImagePullPolicy(bm *eventbrokerv1beta1.MonitoringImage) corev1.PullPolicy {
	imagePullPolicy := corev1.PullIfNotPresent
	if bm != nil && len(bm.ImagePullPolicy) > 0 {
		imagePullPolicy = bm.ImagePullPolicy
	}
	return imagePullPolicy
}

func getExporterImagePullSecrets(bm *eventbrokerv1beta1.MonitoringImage) []corev1.LocalObjectReference {
	var imagePullSecret []corev1.LocalObjectReference
	if bm != nil && len(bm.ImagePullSecrets) > 0 {
		imagePullSecret = bm.ImagePullSecrets
	}
	return imagePullSecret
}

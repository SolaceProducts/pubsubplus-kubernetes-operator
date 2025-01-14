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
	"time"

	pubsubplus "github.com/SolaceProducts/pubsubplus-operator/api/v1beta1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("Monitoring Exporter Test", func() {

	const (
		broker_nonha      = "m-nonha"
		broker_ha         = "m-ha"
		tls_secret        = "monitoring-tls"
		monitoring_secret = "monitoring-user-secret"
		namespace         = "default"
		memLimits         = "1Gi"
		cpuLimits         = "1"
		memRequests       = "500Mi"
		cpuRequests       = "500m"
	)

	Context("When cluster is created, Prometheus Monitoring Exporter is set up", func() {

		It("allows monitoring exporter to be created when enabled and not created when not enabled", func() {

			By("Sets up in Non HA mode when Monitoring is enabled", func() {
				tlsSecret := &corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      tls_secret,
						Namespace: namespace,
					},
					Data: map[string][]byte{
						"tls.crt": []byte("-----BEGIN CERTIFICATE-----\nMIIClDCCAXwCCQDR8jzOOfj9PjANBgkqhkiG9w0BAQsFADAMMQowCAYDVQQDDAEq\nMB4XDTIzMDMwODExNTc1OFoXDTI0MDMwNzExNTc1OFowDDEKMAgGA1UEAwwBKjCC\nASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANU8Gaoh1S426Q7q7rTUg1mM\ndKqFiQXCW/NJ1s4EaL9SjaTaCRrgcoN2eUr2L1lvgBNB0dN7E02OkAeYumKqL20M\nPdOZN8aU/WlYvt9o81Adyy2C03SMugE7t5djIqwk6p6x49uBRK9eLVVjEdWFiyBa\n7wnJPCUdUiSqlJl4PPf+N7GDyOCqERie002gLw+KQHejcoT6z4cfMSCyjAcM++yV\n/LhCa8wW2oB9Q/RMTpEpez6xD41vJ8YRR07CjB7SCFV2fb2EQBoBMTYvTRIwypr7\nqbS4v9sbU3W9I0mYZPtR/ukklonxSmr268HahjB3Dh+1DgzqR8DlZCp8nIuV3W8C\nAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAhBG8kyXasTe9Owxhx2YbPVk0QIQqJa2H\nSC3Ygl792Jt+AUPPSJDKoclGnKWeyKZ2usVU3Katj8V/SIOiosDr0e3XsyhpKVRJ\nDwar43Vkou+R7XUU3is+Oax16Q2Dh9xTESjVB0fzm+QzapO8oSiuk9OERq2W70jt\nksh8J6lHCJcjPTSDZiD84puRhKAcNt5gVul2mA9DLuwKfVKUlthX1uJdN6HaQfy0\n5sxuhjzJhDhYI/vfpu6mI5rFnQXqgj+eoiglHJ9j0qnvDRxffBqR+Zh4Mez51Y+m\nBtmHqY52WLt4RacRByOtnHMcTdiM5bpuxrntqKaGck124AcoBQ6ZLw==\n-----END CERTIFICATE-----\n"),
						"tls.key": []byte("-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDVPBmqIdUuNukO\n6u601INZjHSqhYkFwlvzSdbOBGi/Uo2k2gka4HKDdnlK9i9Zb4ATQdHTexNNjpAH\nmLpiqi9tDD3TmTfGlP1pWL7faPNQHcstgtN0jLoBO7eXYyKsJOqesePbgUSvXi1V\nYxHVhYsgWu8JyTwlHVIkqpSZeDz3/jexg8jgqhEYntNNoC8PikB3o3KE+s+HHzEg\nsowHDPvslfy4QmvMFtqAfUP0TE6RKXs+sQ+NbyfGEUdOwowe0ghVdn29hEAaATE2\nL00SMMqa+6m0uL/bG1N1vSNJmGT7Uf7pJJaJ8Upq9uvB2oYwdw4ftQ4M6kfA5WQq\nfJyLld1vAgMBAAECggEBALKS1ltoYgOF8L+Rd77wid+ghMOZeRrdneus1rtJbf9r\nvztjbWSYus3llcZ1TUn02qlF4dbdp1i4H159RPoD1BvauJxQICmp9F8Y9yBZ4Aok\nKVc/zJ46jDskK6gYWZ0YfXPRPiVBqKfEkuqDQRgz8kNyY+UqJbhfSb9zK2crDsQP\nFnFRDzxsNd+PieZv9XUmKddabdsPiAzbnfHmR0spORrER3Xu6IV8x/voBMJNYxf/\nob2QS2RpqpU2pjwVtXoQer1h5ulOJSMjk6qHnx+aC58rkG4kKOzzyDkFnV6sBo/I\nr476XCwsgglRedDwvoqu32pzWy3DjeKRWc9VhHnQNcECgYEA7GUN4hZnZk6Zc4Ip\nQGgwZejNOC4pX8/BVHxpDVG/l3i5rCyw0Lzf/GLcns7OnR+w4O4CQFFXASaZpaWG\ngawiUwamO+JuRf6EWDicmWwIdeLtnlNs8Rzd7/3ZuB+i1I+LyKoZrraEeIFc+Hx8\nPkW+E4LtxxC9GsKLhvwSc49itWkCgYEA5utU7luuB9ERI+neIKaWeg3bJdiheNaG\nujrtTXITqi0C8alMPoQnVnQfsyW9/8eebqHe9cvrmfNKM/Xvl/doK+e1gEAS+vpx\nHhtaknrF+Nf2ia8Agpi3WOxNg0loyY0eaLwwOX8qG+ZWrN1lhcAofDhZxMTYGVse\nrjeUI/Nj6RcCgYBJ0lz9h6WOq2j8S196f47tpD/CFZhSFVz4d0mPIUJFmSvSerpU\n1UbVWEIxTb/0DVt9QpZtY3laIKXGtuRERm8Jon/zH4j0TsEhk7xDpRsXRWCTGtZg\njXU5ZvrApxCAdLtgVM5kYxcHUs6nwqhCAiGTkkWS7sU/QBW2d62DbPmUUQKBgARv\nousNUdOOnaCt/nlsGdnwaDRa7AcxP9dWCHcDaQNM6BCSaweMbGEJzA4Z/INsZ0vC\nylC4gScs+FD1OYwW0aZ+RgtXr8WoiAHHDr9fomv8Yh0VApJ/so3/xCFwiJXOozXp\n35dLLRjqHOInQqsGHQD96COSkIA0Muuv36WtKE8zAoGASEc60XkoXcpZnyKUJLIt\nmZpJdvXD7aD+V5ACBChsLwyRjWJvmV5geJxSH3uwx+/twtHF3RUpCdvBVgGoFBFV\nE7QOfmctQZ2+RF31tL9TJUV2etfCaI5lTQ7tL2zMj1dRYPLR6D8X9RWjFyyKMV22\nbDdmIg7zk1T9GQNyMHZJXqU=\n-----END PRIVATE KEY-----\n"),
					},
					Type: corev1.SecretTypeTLS,
				}
				Expect(k8sClient.Create(ctx, tlsSecret)).Should(Succeed())

				monitoringSecret := &corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      monitoring_secret,
						Namespace: namespace,
					},
					Data: map[string][]byte{
						monitorSecretKeyName: []byte("dummymonitoringsecretbroker"),
					},
					Type: corev1.SecretTypeOpaque,
				}
				Expect(k8sClient.Create(ctx, monitoringSecret)).Should(Succeed())

				brokerMNonHA := &pubsubplus.PubSubPlusEventBroker{
					ObjectMeta: metav1.ObjectMeta{
						Name:      broker_nonha,
						Namespace: namespace,
					},
					Spec: pubsubplus.EventBrokerSpec{
						Developer:      true,
						Redundancy:     false,
						UpdateStrategy: pubsubplus.AutomatedRollingUpdateStrategy,
						Storage: pubsubplus.Storage{
							Slow: false,
						},
						BrokerTLS: pubsubplus.BrokerTLS{
							Enabled: false,
						},
						Monitoring: pubsubplus.Monitoring{
							Enabled: true,
							MonitoringImage: &pubsubplus.MonitoringImage{
								Repository: "ghcr.io/solacedev/pubsubplus-prometheus-exporter",
								Tag:        "latest",
								ImagePullSecrets: []corev1.LocalObjectReference{
									{
										Name: "regcred",
									},
								},
							},
							MonitoringMetricsEndpoint: &pubsubplus.MonitoringMetricsEndpoint{
								Name:          "monitoring-tcp",
								Protocol:      corev1.ProtocolTCP,
								ServiceType:   corev1.ServiceTypeClusterIP,
								ServicePort:   7629,
								ContainerPort: 7629,
							},
							Resources: corev1.ResourceRequirements{
								Limits: map[corev1.ResourceName]resource.Quantity{
									corev1.ResourceCPU:    resource.MustParse(cpuLimits),
									corev1.ResourceMemory: resource.MustParse(memLimits),
								},
								Requests: map[corev1.ResourceName]resource.Quantity{
									corev1.ResourceCPU:    resource.MustParse(cpuRequests),
									corev1.ResourceMemory: resource.MustParse(memRequests),
								},
							},
						},
					},
				}
				Expect(k8sClient.Create(ctx, brokerMNonHA)).Should(Succeed())

				time.Sleep(100 * time.Second)

				statefulset := &appsv1.StatefulSet{}
				statefulsetName := getStatefulsetName(brokerMNonHA.Name, "p")
				_ = k8sClient.Get(ctx, types.NamespacedName{Name: statefulsetName, Namespace: brokerMNonHA.Namespace}, statefulset)

				//confirm Monitoring Exporter Deployment is found
				EventuallyWithOffset(60, func() bool {
					monitoringDeployment := &appsv1.Deployment{}
					monitoringExporter := getObjectName("PrometheusExporterDeployment", brokerMNonHA.Name)
					err := k8sClient.Get(ctx, types.NamespacedName{Name: monitoringExporter, Namespace: brokerMNonHA.Namespace}, monitoringDeployment)
					// confirm the user provided resources are applied
					Expect(monitoringDeployment.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu().String()).To(Equal(cpuLimits))
					Expect(monitoringDeployment.Spec.Template.Spec.Containers[0].Resources.Limits.Memory().String()).To(Equal(memLimits))
					Expect(monitoringDeployment.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu().String()).To(Equal(cpuRequests))
					Expect(monitoringDeployment.Spec.Template.Spec.Containers[0].Resources.Requests.Memory().String()).To(Equal(memRequests))

					return err == nil
				}).WithTimeout(90 * time.Second).Should(BeTrue())

				//delete broker
				Expect(k8sClient.Delete(ctx, brokerMNonHA)).Should(Succeed())

				Expect(k8sClient.Delete(ctx, tlsSecret)).Should(Succeed())
				Expect(k8sClient.Delete(ctx, monitoringSecret)).Should(Succeed())
			})

			By("Does not setup Monitoring Exporter when Enabled flag is False", func() {
				brokerHA := &pubsubplus.PubSubPlusEventBroker{
					ObjectMeta: metav1.ObjectMeta{
						Name:      broker_ha,
						Namespace: namespace,
					},
					Spec: pubsubplus.EventBrokerSpec{
						Developer:      false,
						Redundancy:     true,
						UpdateStrategy: pubsubplus.ManualPodRestartUpdateStrategy,
						BrokerNodeAssignment: []pubsubplus.NodeAssignment{
							{
								Name: "Primary",
								Spec: pubsubplus.NodeAssignmentSpec{
									NodeSelector: map[string]string{
										"kubernetes.io/os": "linux",
									},
								},
							}, {
								Name: "Backup",
								Spec: pubsubplus.NodeAssignmentSpec{
									NodeSelector: map[string]string{
										"kubernetes.io/os": "linux",
									},
								},
							},
						},
						Monitoring: pubsubplus.Monitoring{
							Enabled: false,
						},
					},
				}
				Expect(k8sClient.Create(ctx, brokerHA)).Should(Succeed())

				statefulset := &appsv1.StatefulSet{}
				statefulsetName := getStatefulsetName(brokerHA.Name, "p")
				_ = k8sClient.Get(ctx, types.NamespacedName{Name: statefulsetName, Namespace: brokerHA.Namespace}, statefulset)

				statefulsetB := &appsv1.StatefulSet{}
				statefulsetBName := getStatefulsetName(brokerHA.Name, "b")
				_ = k8sClient.Get(ctx, types.NamespacedName{Name: statefulsetBName, Namespace: brokerHA.Namespace}, statefulsetB)
				statefulsetM := &appsv1.StatefulSet{}
				statefulsetMName := getStatefulsetName(brokerHA.Name, "m")
				_ = k8sClient.Get(ctx, types.NamespacedName{Name: statefulsetMName, Namespace: brokerHA.Namespace}, statefulsetM)

				//confirm Monitoring Deployment is not found
				EventuallyWithOffset(3, func() bool {
					monitoringDeployment := &appsv1.Deployment{}
					monitoringExporter := getObjectName("PrometheusExporterDeployment", brokerHA.Name)
					err := k8sClient.Get(ctx, types.NamespacedName{Name: monitoringExporter, Namespace: brokerHA.Namespace}, monitoringDeployment)
					return err != nil
				}).WithTimeout(3 * time.Second).Should(BeTrue())

				//delete broker
				Expect(k8sClient.Delete(ctx, brokerHA)).To(Succeed())

			})

		})
	})

})

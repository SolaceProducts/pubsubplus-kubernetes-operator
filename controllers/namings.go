/*
Copyright 2022.

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

// One place for consistent naming

package controllers

import (
	"fmt"
	"strconv"
)

const (
	dependenciesSignatureAnnotationName = "lastAppliedConfig"
	appKubernetesIoNameLabel            = "pubsubpluseventbroker"
	appKubernetesIoManagedByLabel       = "solace-pubsubplus-operator"
	secretKeyName                       = "username_admin_password"
)

type BrokerRole int // Notice that this is about the current role, not the broker node designation
const (
	Active = iota
	Standby
	Monitor
	PrometheusExporter
)

// Provides the object names for the current EventBroker deployment
func getObjectName(objectType string, deploymentName string) string {
	nameSuffix := map[string]string{
		"ConfigMap":                    "-pubsubplus",
		"DiscoveryService":             "-pubsubplus-discovery",
		"Role":                         "-pubsubplus-podtagupdater",
		"RoleBinding":                  "-pubsubplus-sa-to-podtagupdater",
		"ServiceAccount":               "-pubsubplus-sa",
		"Secret":                       "-pubsubplus-secrets",
		"Service":                      "-pubsubplus",
		"StatefulSet":                  "-pubsubplus-%s",
		"PrometheusExporterDeployment": "-pubsubplus-prometheus-exporter",
		"PrometheusExporterService":    "-pubsubplus-prometheus-exporter-service",
	}
	return deploymentName + nameSuffix[objectType]
}

// Provides the name of a StatefulSet of a role specified by roleSuffix
// roleSuffix can be `p` (Primary), `b` (Backup) or `m` (Monitor)
func getStatefulsetName(deploymentName string, roleSuffix string) string {
	return fmt.Sprintf(getObjectName("StatefulSet", deploymentName), roleSuffix)
}

// Provides the labels for an object in the current EventBroker deployment
// These labels are used for all the objects except for Pods
func getObjectLabels(deploymentName string) map[string]string {
	return map[string]string{
		"app.kubernetes.io/instance":   deploymentName,
		"app.kubernetes.io/name":       appKubernetesIoNameLabel,
		"app.kubernetes.io/managed-by": appKubernetesIoManagedByLabel,
	}
}

func getBrokerNodeType(statefulSetDeploymentName string) string {
	nodeTypeSymbol := string(statefulSetDeploymentName[len(statefulSetDeploymentName)-1]) // Last char of the StatefulSet deployment name
	return (map[string]string{"p": "message-routing-primary", "b": "message-routing-backup", "m": "monitor"})[nodeTypeSymbol]
}

// Provides the labels for a broker Pod in the current EventBroker deployment
func getPodLabels(deploymentName string, nodeType string) map[string]string {
	return map[string]string{
		"app.kubernetes.io/instance": deploymentName,
		"app.kubernetes.io/name":     appKubernetesIoNameLabel,
		"node-type":                  nodeType,
	}
}

// Provides the selector (from Pods) to be used for broker services
func getServiceSelector(deploymentName string) map[string]string {
	return map[string]string{
		"app.kubernetes.io/instance": deploymentName,
		"app.kubernetes.io/name":     appKubernetesIoNameLabel,
		"active":                     "true",
	}
}

// Provides the selector (from Pods) to be used for broker nodes discovery services
func getDiscoveryServiceSelector(deploymentName string) map[string]string {
	return map[string]string{
		"app.kubernetes.io/instance": deploymentName,
		"app.kubernetes.io/name":     appKubernetesIoNameLabel,
	}
}

// Provides the Pod selector for broker pods
func getBrokerPodSelector(deploymentName string, brokerRole BrokerRole) map[string]string {
	if brokerRole == Monitor {
		return map[string]string{
			"app.kubernetes.io/instance": deploymentName,
			"app.kubernetes.io/name":     appKubernetesIoNameLabel,
			"node-type":                  "monitor",
		}
	} else if brokerRole == PrometheusExporter {
		return map[string]string{
			"app.kubernetes.io/instance": deploymentName,
			"app.kubernetes.io/name":     appKubernetesIoNameLabel,
			"solace-prometheus-exporter": "true",
		}
	} else {
		return map[string]string{
			"app.kubernetes.io/instance": deploymentName,
			"app.kubernetes.io/name":     appKubernetesIoNameLabel,
			"active":                     strconv.FormatBool(brokerRole == Active),
		}
	}
}

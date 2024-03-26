//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022, Solace Corporation

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerImage) DeepCopyInto(out *BrokerImage) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerImage.
func (in *BrokerImage) DeepCopy() *BrokerImage {
	if in == nil {
		return nil
	}
	out := new(BrokerImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerPersistentVolumeClaim) DeepCopyInto(out *BrokerPersistentVolumeClaim) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerPersistentVolumeClaim.
func (in *BrokerPersistentVolumeClaim) DeepCopy() *BrokerPersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(BrokerPersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerPort) DeepCopyInto(out *BrokerPort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerPort.
func (in *BrokerPort) DeepCopy() *BrokerPort {
	if in == nil {
		return nil
	}
	out := new(BrokerPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerServiceAccount) DeepCopyInto(out *BrokerServiceAccount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerServiceAccount.
func (in *BrokerServiceAccount) DeepCopy() *BrokerServiceAccount {
	if in == nil {
		return nil
	}
	out := new(BrokerServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSubStatus) DeepCopyInto(out *BrokerSubStatus) {
	*out = *in
	if in.StatefulSets != nil {
		in, out := &in.StatefulSets, &out.StatefulSets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSubStatus.
func (in *BrokerSubStatus) DeepCopy() *BrokerSubStatus {
	if in == nil {
		return nil
	}
	out := new(BrokerSubStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerTLS) DeepCopyInto(out *BrokerTLS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerTLS.
func (in *BrokerTLS) DeepCopy() *BrokerTLS {
	if in == nil {
		return nil
	}
	out := new(BrokerTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSpec) DeepCopyInto(out *ContainerSpec) {
	*out = *in
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSpec.
func (in *ContainerSpec) DeepCopy() *ContainerSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventBrokerSpec) DeepCopyInto(out *EventBrokerSpec) {
	*out = *in
	if in.SystemScaling != nil {
		in, out := &in.SystemScaling, &out.SystemScaling
		*out = new(SystemScaling)
		**out = **in
	}
	if in.ExtraEnvVars != nil {
		in, out := &in.ExtraEnvVars, &out.ExtraEnvVars
		*out = make([]*ExtraEnvVar, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ExtraEnvVar)
				**out = **in
			}
		}
	}
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.BrokerImage.DeepCopyInto(&out.BrokerImage)
	if in.BrokerNodeAssignment != nil {
		in, out := &in.BrokerNodeAssignment, &out.BrokerNodeAssignment
		*out = make([]NodeAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.SecurityContext = in.SecurityContext
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
	out.ServiceAccount = in.ServiceAccount
	out.BrokerTLS = in.BrokerTLS
	in.Service.DeepCopyInto(&out.Service)
	in.Storage.DeepCopyInto(&out.Storage)
	in.Monitoring.DeepCopyInto(&out.Monitoring)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventBrokerSpec.
func (in *EventBrokerSpec) DeepCopy() *EventBrokerSpec {
	if in == nil {
		return nil
	}
	out := new(EventBrokerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventBrokerStatus) DeepCopyInto(out *EventBrokerStatus) {
	*out = *in
	if in.PodsList != nil {
		in, out := &in.PodsList, &out.PodsList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Broker.DeepCopyInto(&out.Broker)
	out.Monitoring = in.Monitoring
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventBrokerStatus.
func (in *EventBrokerStatus) DeepCopy() *EventBrokerStatus {
	if in == nil {
		return nil
	}
	out := new(EventBrokerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraEnvVar) DeepCopyInto(out *ExtraEnvVar) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraEnvVar.
func (in *ExtraEnvVar) DeepCopy() *ExtraEnvVar {
	if in == nil {
		return nil
	}
	out := new(ExtraEnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
	if in.MonitoringImage != nil {
		in, out := &in.MonitoringImage, &out.MonitoringImage
		*out = new(MonitoringImage)
		(*in).DeepCopyInto(*out)
	}
	if in.MonitoringMetricsEndpoint != nil {
		in, out := &in.MonitoringMetricsEndpoint, &out.MonitoringMetricsEndpoint
		*out = new(MonitoringMetricsEndpoint)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringImage) DeepCopyInto(out *MonitoringImage) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringImage.
func (in *MonitoringImage) DeepCopy() *MonitoringImage {
	if in == nil {
		return nil
	}
	out := new(MonitoringImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringMetricsEndpoint) DeepCopyInto(out *MonitoringMetricsEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringMetricsEndpoint.
func (in *MonitoringMetricsEndpoint) DeepCopy() *MonitoringMetricsEndpoint {
	if in == nil {
		return nil
	}
	out := new(MonitoringMetricsEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringSubStatus) DeepCopyInto(out *MonitoringSubStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringSubStatus.
func (in *MonitoringSubStatus) DeepCopy() *MonitoringSubStatus {
	if in == nil {
		return nil
	}
	out := new(MonitoringSubStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAssignment) DeepCopyInto(out *NodeAssignment) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAssignment.
func (in *NodeAssignment) DeepCopy() *NodeAssignment {
	if in == nil {
		return nil
	}
	out := new(NodeAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAssignmentSpec) DeepCopyInto(out *NodeAssignmentSpec) {
	*out = *in
	in.Affinity.DeepCopyInto(&out.Affinity)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAssignmentSpec.
func (in *NodeAssignmentSpec) DeepCopy() *NodeAssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(NodeAssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubSubPlusEventBroker) DeepCopyInto(out *PubSubPlusEventBroker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubSubPlusEventBroker.
func (in *PubSubPlusEventBroker) DeepCopy() *PubSubPlusEventBroker {
	if in == nil {
		return nil
	}
	out := new(PubSubPlusEventBroker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PubSubPlusEventBroker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubSubPlusEventBrokerList) DeepCopyInto(out *PubSubPlusEventBrokerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PubSubPlusEventBroker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubSubPlusEventBrokerList.
func (in *PubSubPlusEventBrokerList) DeepCopy() *PubSubPlusEventBrokerList {
	if in == nil {
		return nil
	}
	out := new(PubSubPlusEventBrokerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PubSubPlusEventBrokerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityContext) DeepCopyInto(out *SecurityContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityContext.
func (in *SecurityContext) DeepCopy() *SecurityContext {
	if in == nil {
		return nil
	}
	out := new(SecurityContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]*BrokerPort, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BrokerPort)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.CustomVolumeMount != nil {
		in, out := &in.CustomVolumeMount, &out.CustomVolumeMount
		*out = make([]StorageCustomVolumeMount, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageCustomVolumeMount) DeepCopyInto(out *StorageCustomVolumeMount) {
	*out = *in
	out.PersistentVolumeClaim = in.PersistentVolumeClaim
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageCustomVolumeMount.
func (in *StorageCustomVolumeMount) DeepCopy() *StorageCustomVolumeMount {
	if in == nil {
		return nil
	}
	out := new(StorageCustomVolumeMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemScaling) DeepCopyInto(out *SystemScaling) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemScaling.
func (in *SystemScaling) DeepCopy() *SystemScaling {
	if in == nil {
		return nil
	}
	out := new(SystemScaling)
	in.DeepCopyInto(out)
	return out
}

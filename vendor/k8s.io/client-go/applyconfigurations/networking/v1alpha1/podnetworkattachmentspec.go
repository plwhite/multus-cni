/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// PodNetworkAttachmentSpecApplyConfiguration represents an declarative configuration of the PodNetworkAttachmentSpec type for use
// with apply.
type PodNetworkAttachmentSpecApplyConfiguration struct {
	PodNetworkName *string                           `json:"podNetworkName,omitempty"`
	ParametersRefs []ParametersRefApplyConfiguration `json:"parametersRefs,omitempty"`
}

// PodNetworkAttachmentSpecApplyConfiguration constructs an declarative configuration of the PodNetworkAttachmentSpec type for use with
// apply.
func PodNetworkAttachmentSpec() *PodNetworkAttachmentSpecApplyConfiguration {
	return &PodNetworkAttachmentSpecApplyConfiguration{}
}

// WithPodNetworkName sets the PodNetworkName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodNetworkName field is set to the value of the last call.
func (b *PodNetworkAttachmentSpecApplyConfiguration) WithPodNetworkName(value string) *PodNetworkAttachmentSpecApplyConfiguration {
	b.PodNetworkName = &value
	return b
}

// WithParametersRefs adds the given value to the ParametersRefs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ParametersRefs field.
func (b *PodNetworkAttachmentSpecApplyConfiguration) WithParametersRefs(values ...*ParametersRefApplyConfiguration) *PodNetworkAttachmentSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithParametersRefs")
		}
		b.ParametersRefs = append(b.ParametersRefs, *values[i])
	}
	return b
}
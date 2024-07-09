package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GatewayClassSpec defines the desired state of GatewayClass
type GatewayClassSpec struct {
	ControllerName string         `json:"controllerName,omitempty"`
	ParametersRef  *ParametersRef `json:"parametersRef,omitempty"`
}

// ParametersRef defines the reference to the configuration parameters
type ParametersRef struct {
	Group string `json:"group,omitempty"`
	Kind  string `json:"kind,omitempty"`
	Name  string `json:"name,omitempty"`
}

// GatewayClassStatus defines the observed state of GatewayClass
type GatewayClassStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GatewayClass is the Schema for the gatewayclasses API
type GatewayClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GatewayClassSpec   `json:"spec,omitempty"`
	Status GatewayClassStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayClassList contains a list of GatewayClass
type GatewayClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GatewayClass{}, &GatewayClassList{})
}

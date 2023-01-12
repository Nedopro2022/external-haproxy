package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExternalHAProxySpec defines the desired state of ExternalHAProxy
type ExternalHAProxySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ExternalHAProxy. Edit externalhaproxy_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ExternalHAProxyStatus defines the observed state of ExternalHAProxy
type ExternalHAProxyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ExternalHAProxy is the Schema for the externalhaproxies API
type ExternalHAProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalHAProxySpec   `json:"spec,omitempty"`
	Status ExternalHAProxyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ExternalHAProxyList contains a list of ExternalHAProxy
type ExternalHAProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalHAProxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExternalHAProxy{}, &ExternalHAProxyList{})
}

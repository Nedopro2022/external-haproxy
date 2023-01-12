package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TCPFrontendSpec defines the desired state of TCPFrontend
type TCPFrontendSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of TCPFrontend. Edit tcpfrontend_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// TCPFrontendStatus defines the observed state of TCPFrontend
type TCPFrontendStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TCPFrontend is the Schema for the tcpfrontends API
type TCPFrontend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TCPFrontendSpec   `json:"spec,omitempty"`
	Status TCPFrontendStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TCPFrontendList contains a list of TCPFrontend
type TCPFrontendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TCPFrontend `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TCPFrontend{}, &TCPFrontendList{})
}

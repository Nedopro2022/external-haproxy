package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TCPFrontendSpec defines the desired state of TCPFrontend
type TCPFrontendSpec struct {
	// HAProxy specifies the associated ExternalHAProxy resource.
	HAProxy corev1.LocalObjectReference `json:"haproxy"`

	// Binds specifies bind addresses.
	// E.g., [ ":8080" ]
	Binds []string `json:"binds"`

	// Backend specifies the associated TCPBackend resource.
	Backend corev1.LocalObjectReference `json:"backend"`
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

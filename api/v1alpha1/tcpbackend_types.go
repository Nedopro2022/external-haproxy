package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackendServer defines a backend server.
type BackendServer struct {
	// Name specifies the server name.
	// E.g., "sv1"
	Name string `json:"name"`

	// Address specifies the address.
	// E.g., "10.0.0.101:8080"
	Address string `json:"address"`

	// Weight specifies the weight. (range: 0-255)
	Weight uint8 `json:"weight"`
}

// TCPBackendSpec defines the desired state of TCPBackend
type TCPBackendSpec struct {
	// HAProxy specifies the associated ExternalHAProxy resource.
	HAProxy corev1.LocalObjectReference `json:"haproxy"`

	// Servers specifies backend servers.
	Servers []BackendServer `json:"servers"`
}

// TCPBackendStatus defines the observed state of TCPBackend
type TCPBackendStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TCPBackend is the Schema for the tcpbackends API
type TCPBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TCPBackendSpec   `json:"spec,omitempty"`
	Status TCPBackendStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TCPBackendList contains a list of TCPBackend
type TCPBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TCPBackend `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TCPBackend{}, &TCPBackendList{})
}

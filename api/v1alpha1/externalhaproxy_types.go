package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataPlaneAPI defines the settings of a HAProxy Data Plane API endpoint.
type DataPlaneAPI struct {
	// HTTPS specifies whether to use HTTPS.
	// Using HTTPS is strongly recommended as basic auth sends passwords in plain text.
	HTTPS bool `json:"https"`

	// Host specifies the host name.
	// E.g., "10.0.0.10"
	Host string `json:"host"`

	// Port specifies the port number.
	// E.g., "5555"
	Port string `json:"port"`

	// BasicAuthSecret specifies the Secret resource used in basic authentication.
	// The Secret resource should be `type: kubernetes.io/basic-auth`.
	BasicAuthSecret corev1.LocalObjectReference `json:"basicAuthSecret"`
}

// ExternalHAProxySpec defines the desired state of ExternalHAProxy
type ExternalHAProxySpec struct {
	// DataPlaneAPI specifies the HAProxy Data Plane API endpoint.
	DataPlaneAPI DataPlaneAPI `json:"dataplaneAPI"`

	// BaseConfig specifies the base HAProxy config.
	BaseConfig string `json:"baseConfig"`
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

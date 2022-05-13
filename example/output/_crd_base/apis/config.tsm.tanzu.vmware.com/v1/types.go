// Code generated by nexus. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	gnstsmtanzuvmwarecomv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/_generated/apis/gns.tsm.tanzu.vmware.com/v1"

	"gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/_generated/common"
)

// +k8s:openapi-gen=true
type Child struct {
	Group string `json:"group" yaml:"group"`
	Kind  string `json:"kind" yaml:"kind"`
	Name  string `json:"name" yaml:"name"`
}

// +k8s:openapi-gen=true
type Link struct {
	Group string `json:"group" yaml:"group"`
	Kind  string `json:"kind" yaml:"kind"`
	Name  string `json:"name" yaml:"name"`
}

/* ------------------- CRDs definitions ------------------- */

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type Config struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              ConfigSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

func (c *Config) CRDName() string {
	return "configs.config.tsm.tanzu.vmware.com"
}

func (c *Config) DisplayName() string {
	if c.GetLabels() != nil {
		return c.GetLabels()[common.DISPLAY_NAME_LABEL]
	}
	return ""
}

// +k8s:openapi-gen=true
type ConfigSpec struct {
	MyStr  gnstsmtanzuvmwarecomv1.MyStr `json:"myStr" yaml:"myStr"`
	GNSGvk *Child                       `json:"gNSGvk,omitempty" yaml:"gNSGvk,omitempty" nexus:"child"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ConfigList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Config `json:"items" yaml:"items"`
}

// Code generated by nexus. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"nexustempmodule/common"
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

// +k8s:openapi-gen=true
type NexusStatus struct {
	SourceGeneration int64 `json:"sourceGeneration" yaml:"sourceGeneration"`
	RemoteGeneration int64 `json:"remoteGeneration" yaml:"remoteGeneration"`
}

/* ------------------- CRDs definitions ------------------- */

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type Gns struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              GnsSpec        `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status            GnsNexusStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// +k8s:openapi-gen=true
type GnsNexusStatus struct {
	Nexus NexusStatus `json:"nexus,omitempty" yaml:"nexus,omitempty"`
}

func (c *Gns) CRDName() string {
	return "gnses.gns.tsm.tanzu.vmware.com"
}

func (c *Gns) DisplayName() string {
	if c.GetLabels() != nil {
		return c.GetLabels()[common.DISPLAY_NAME_LABEL]
	}
	return ""
}

// +k8s:openapi-gen=true
type GnsSpec struct {
	//nexus-validation: MaxLength=8, MinLength=2
	//nexus-validation: Pattern=abc
	Domain           string           `json:"domain" yaml:"domain"`
	UseSharedGateway bool             `json:"useSharedGateway" yaml:"useSharedGateway"`
	Mydesc           Description      `json:"mydesc" yaml:"mydesc"`
	HostPort         HostPort         `json:"hostPort" yaml:"hostPort"`
	TestArray        EmptyData        `json:"testArray" yaml:"testArray"`
	Instance         Instance         `json:"instance" yaml:"instance"`
	Array1           []uint32         `json:"array1" yaml:"array1"`
	Array2           []Description    `json:"array2" yaml:"array2"`
	Array3           []Bar            `json:"array3" yaml:"array3"`
	Array4           []Instance       `json:"array4" yaml:"array4"`
	Array5           AliasArr         `json:"array5" yaml:"array5"`
	FooChildGvk      *Child           `json:"fooChildGvk,omitempty" yaml:"fooChildGvk,omitempty" nexus:"child"`
	FooChildrenGvk   map[string]Child `json:"fooChildrenGvk,omitempty" yaml:"fooChildrenGvk,omitempty" nexus:"children"`
	FooLinkGvk       *Link            `json:"fooLinkGvk,omitempty" yaml:"fooLinkGvk,omitempty" nexus:"link"`
	FooLinksGvk      map[string]Link  `json:"fooLinksGvk,omitempty" yaml:"fooLinksGvk,omitempty" nexus:"links"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type GnsList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Gns `json:"items" yaml:"items"`
}

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type Bar struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              BarSpec        `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status            BarNexusStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// +k8s:openapi-gen=true
type BarNexusStatus struct {
	Nexus NexusStatus `json:"nexus,omitempty" yaml:"nexus,omitempty"`
}

func (c *Bar) CRDName() string {
	return "bars.gns.tsm.tanzu.vmware.com"
}

func (c *Bar) DisplayName() string {
	if c.GetLabels() != nil {
		return c.GetLabels()[common.DISPLAY_NAME_LABEL]
	}
	return ""
}

// +k8s:openapi-gen=true
type BarSpec struct {
	Name []string `json:"name" yaml:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BarList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Bar `json:"items" yaml:"items"`
}

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type EmptyData struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`

	Status EmptyDataNexusStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// +k8s:openapi-gen=true
type EmptyDataNexusStatus struct {
	Nexus NexusStatus `json:"nexus,omitempty" yaml:"nexus,omitempty"`
}

func (c *EmptyData) CRDName() string {
	return "emptydatas.gns.tsm.tanzu.vmware.com"
}

func (c *EmptyData) DisplayName() string {
	if c.GetLabels() != nil {
		return c.GetLabels()[common.DISPLAY_NAME_LABEL]
	}
	return ""
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EmptyDataList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []EmptyData `json:"items" yaml:"items"`
}

// +k8s:openapi-gen=true
type HostPort struct {
	Host Host
	Port Port
}

// +k8s:openapi-gen=true
type Description struct {
	Color     string
	Version   string
	ProjectID []string
	TestAns   []Answer
	Instance  Instance
	HostPort  HostPort
}

// +k8s:openapi-gen=true
type Answer struct {
	Name string
}

type Port uint16

// Host the IP of the endpoint
type Host string
type Instance float32
type AliasArr []int

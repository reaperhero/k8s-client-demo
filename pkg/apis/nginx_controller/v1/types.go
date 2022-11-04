package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Nginx struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               Spec   `json:"spec"`
	Status             Status `json:"status, omitempty"`
}

type NginxList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []Nginx `json:"items"`
}

type Spec struct {
	Replicas int `json:"replicas"`
}

type Status struct {
	Message string `json:"message,omitempty"`
}

package internal

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Resource struct {
	metav1.GroupVersionKind `json:"groupVersionKind"`
	SubResource             string `json:"subResource"`
}

type Metadata struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace"`
	Annotations map[string]string `json:"annotations"`
	Labels      map[string]string `json:"labels"`
}

func (m Metadata) isEmpty() bool {
	return m.Name == ""
}

type ObjectWatched struct {
	Metadata   `json:"metadata"`
	Spec       map[string]interface{} `json:"spec"`
	APIVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Status     map[string]interface{} `json:"status"`
}

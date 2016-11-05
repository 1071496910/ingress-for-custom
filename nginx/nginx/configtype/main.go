/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package configtype

import (
	"errors"
	"strings"

	"k8s.io/kubernetes/pkg/apis/extensions"
)

const (
	secureUpstream = "ingress.kubernetes.io/configtype"
)

type ingAnnotations map[string]string

func (a ingAnnotations) configType() []string {
	servers, ok := a[secureUpstream]
	if ok {
		return strings.Split(servers, ",")
	}
	return nil
}

// Get the servers which has custom config
func ParseAnnotations(ing *extensions.Ingress) ([]string, error) {
	if ing.GetAnnotations() == nil {
		return nil, errors.New("no annotations present")
	}

	return ingAnnotations(ing.GetAnnotations()).configType(), nil
}

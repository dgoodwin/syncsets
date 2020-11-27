package api

import (
	"fmt"
)

type Registry struct {
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) GetResource(resourceName string) (APIResource, error) {
	switch resourceName {
	case "clusters":
		return &Cluster{}, nil
	case "syncsets":
		return &SelectorSyncSet{}, nil
	}
	return nil, fmt.Errorf("no registered resource: %s", resourceName)
}

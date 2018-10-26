/*
Copyright 2018 The Knative Authors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/knative/eventing/pkg/apis/flows/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FlowLister helps list Flows.
type FlowLister interface {
	// List lists all Flows in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Flow, err error)
	// Flows returns an object that can list and get Flows.
	Flows(namespace string) FlowNamespaceLister
	FlowListerExpansion
}

// flowLister implements the FlowLister interface.
type flowLister struct {
	indexer cache.Indexer
}

// NewFlowLister returns a new FlowLister.
func NewFlowLister(indexer cache.Indexer) FlowLister {
	return &flowLister{indexer: indexer}
}

// List lists all Flows in the indexer.
func (s *flowLister) List(selector labels.Selector) (ret []*v1alpha1.Flow, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Flow))
	})
	return ret, err
}

// Flows returns an object that can list and get Flows.
func (s *flowLister) Flows(namespace string) FlowNamespaceLister {
	return flowNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FlowNamespaceLister helps list and get Flows.
type FlowNamespaceLister interface {
	// List lists all Flows in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Flow, err error)
	// Get retrieves the Flow from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Flow, error)
	FlowNamespaceListerExpansion
}

// flowNamespaceLister implements the FlowNamespaceLister
// interface.
type flowNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Flows in the indexer for a given namespace.
func (s flowNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Flow, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Flow))
	})
	return ret, err
}

// Get retrieves the Flow from the indexer for a given namespace and name.
func (s flowNamespaceLister) Get(name string) (*v1alpha1.Flow, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("flow"), name)
	}
	return obj.(*v1alpha1.Flow), nil
}

/*
Copyright The Kubernetes Authors.

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

package v1

import (
	v1 "github.com/logston/k8s-catzzz-logger/pkg/apis/catzzzlogger/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CatzzzLoggerLister helps list CatzzzLoggers.
type CatzzzLoggerLister interface {
	// List lists all CatzzzLoggers in the indexer.
	List(selector labels.Selector) (ret []*v1.CatzzzLogger, err error)
	// CatzzzLoggers returns an object that can list and get CatzzzLoggers.
	CatzzzLoggers(namespace string) CatzzzLoggerNamespaceLister
	CatzzzLoggerListerExpansion
}

// catzzzLoggerLister implements the CatzzzLoggerLister interface.
type catzzzLoggerLister struct {
	indexer cache.Indexer
}

// NewCatzzzLoggerLister returns a new CatzzzLoggerLister.
func NewCatzzzLoggerLister(indexer cache.Indexer) CatzzzLoggerLister {
	return &catzzzLoggerLister{indexer: indexer}
}

// List lists all CatzzzLoggers in the indexer.
func (s *catzzzLoggerLister) List(selector labels.Selector) (ret []*v1.CatzzzLogger, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CatzzzLogger))
	})
	return ret, err
}

// CatzzzLoggers returns an object that can list and get CatzzzLoggers.
func (s *catzzzLoggerLister) CatzzzLoggers(namespace string) CatzzzLoggerNamespaceLister {
	return catzzzLoggerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CatzzzLoggerNamespaceLister helps list and get CatzzzLoggers.
type CatzzzLoggerNamespaceLister interface {
	// List lists all CatzzzLoggers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.CatzzzLogger, err error)
	// Get retrieves the CatzzzLogger from the indexer for a given namespace and name.
	Get(name string) (*v1.CatzzzLogger, error)
	CatzzzLoggerNamespaceListerExpansion
}

// catzzzLoggerNamespaceLister implements the CatzzzLoggerNamespaceLister
// interface.
type catzzzLoggerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CatzzzLoggers in the indexer for a given namespace.
func (s catzzzLoggerNamespaceLister) List(selector labels.Selector) (ret []*v1.CatzzzLogger, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CatzzzLogger))
	})
	return ret, err
}

// Get retrieves the CatzzzLogger from the indexer for a given namespace and name.
func (s catzzzLoggerNamespaceLister) Get(name string) (*v1.CatzzzLogger, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("catzzzlogger"), name)
	}
	return obj.(*v1.CatzzzLogger), nil
}

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	catzzzloggerv1 "github.com/logston/k8s-catzzz-logger/pkg/apis/catzzzlogger/v1"
	versioned "github.com/logston/k8s-catzzz-logger/pkg/client/clientset/versioned"
	internalinterfaces "github.com/logston/k8s-catzzz-logger/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/logston/k8s-catzzz-logger/pkg/client/listers/catzzzlogger/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CatzzzLoggerInformer provides access to a shared informer and lister for
// CatzzzLoggers.
type CatzzzLoggerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CatzzzLoggerLister
}

type catzzzLoggerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCatzzzLoggerInformer constructs a new informer for CatzzzLogger type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCatzzzLoggerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCatzzzLoggerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCatzzzLoggerInformer constructs a new informer for CatzzzLogger type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCatzzzLoggerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PaullogstonV1().CatzzzLoggers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PaullogstonV1().CatzzzLoggers(namespace).Watch(context.TODO(), options)
			},
		},
		&catzzzloggerv1.CatzzzLogger{},
		resyncPeriod,
		indexers,
	)
}

func (f *catzzzLoggerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCatzzzLoggerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *catzzzLoggerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&catzzzloggerv1.CatzzzLogger{}, f.defaultInformer)
}

func (f *catzzzLoggerInformer) Lister() v1.CatzzzLoggerLister {
	return v1.NewCatzzzLoggerLister(f.Informer().GetIndexer())
}

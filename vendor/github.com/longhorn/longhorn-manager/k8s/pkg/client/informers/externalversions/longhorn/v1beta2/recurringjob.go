/*
Copyright The Longhorn Authors.

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

package v1beta2

import (
	context "context"
	time "time"

	apislonghornv1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
	versioned "github.com/longhorn/longhorn-manager/k8s/pkg/client/clientset/versioned"
	internalinterfaces "github.com/longhorn/longhorn-manager/k8s/pkg/client/informers/externalversions/internalinterfaces"
	longhornv1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/client/listers/longhorn/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RecurringJobInformer provides access to a shared informer and lister for
// RecurringJobs.
type RecurringJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() longhornv1beta2.RecurringJobLister
}

type recurringJobInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRecurringJobInformer constructs a new informer for RecurringJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRecurringJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRecurringJobInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRecurringJobInformer constructs a new informer for RecurringJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRecurringJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LonghornV1beta2().RecurringJobs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LonghornV1beta2().RecurringJobs(namespace).Watch(context.TODO(), options)
			},
		},
		&apislonghornv1beta2.RecurringJob{},
		resyncPeriod,
		indexers,
	)
}

func (f *recurringJobInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRecurringJobInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *recurringJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apislonghornv1beta2.RecurringJob{}, f.defaultInformer)
}

func (f *recurringJobInformer) Lister() longhornv1beta2.RecurringJobLister {
	return longhornv1beta2.NewRecurringJobLister(f.Informer().GetIndexer())
}

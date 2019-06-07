/*
 * Copyright (c) 2019 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	meshv1alpha1 "github.com/cellery-io/mesh-controller/pkg/apis/mesh/v1alpha1"
	versioned "github.com/cellery-io/mesh-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/cellery-io/mesh-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/cellery-io/mesh-controller/pkg/client/listers/mesh/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CellInformer provides access to a shared informer and lister for
// Cells.
type CellInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CellLister
}

type cellInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCellInformer constructs a new informer for Cell type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCellInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCellInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCellInformer constructs a new informer for Cell type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCellInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeshV1alpha1().Cells(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeshV1alpha1().Cells(namespace).Watch(options)
			},
		},
		&meshv1alpha1.Cell{},
		resyncPeriod,
		indexers,
	)
}

func (f *cellInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCellInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cellInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&meshv1alpha1.Cell{}, f.defaultInformer)
}

func (f *cellInformer) Lister() v1alpha1.CellLister {
	return v1alpha1.NewCellLister(f.Informer().GetIndexer())
}

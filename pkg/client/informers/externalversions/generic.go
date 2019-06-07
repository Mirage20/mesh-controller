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

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/cellery-io/mesh-controller/pkg/apis/istio/authentication/v1alpha1"
	v1alpha3 "github.com/cellery-io/mesh-controller/pkg/apis/istio/networking/v1alpha3"
	servingv1alpha1 "github.com/cellery-io/mesh-controller/pkg/apis/knative/serving/v1alpha1"
	v1beta1 "github.com/cellery-io/mesh-controller/pkg/apis/knative/serving/v1beta1"
	meshv1alpha1 "github.com/cellery-io/mesh-controller/pkg/apis/mesh/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=authentication, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("policies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Authentication().V1alpha1().Policies().Informer()}, nil

		// Group=mesh, Version=v1alpha1
	case meshv1alpha1.SchemeGroupVersion.WithResource("autoscalepolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Mesh().V1alpha1().AutoscalePolicies().Informer()}, nil
	case meshv1alpha1.SchemeGroupVersion.WithResource("cells"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Mesh().V1alpha1().Cells().Informer()}, nil
	case meshv1alpha1.SchemeGroupVersion.WithResource("gateways"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Mesh().V1alpha1().Gateways().Informer()}, nil
	case meshv1alpha1.SchemeGroupVersion.WithResource("services"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Mesh().V1alpha1().Services().Informer()}, nil
	case meshv1alpha1.SchemeGroupVersion.WithResource("tokenservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Mesh().V1alpha1().TokenServices().Informer()}, nil

		// Group=networking, Version=v1alpha3
	case v1alpha3.SchemeGroupVersion.WithResource("destinationrules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().DestinationRules().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("envoyfilters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().EnvoyFilters().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("gateways"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().Gateways().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("virtualservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().VirtualServices().Informer()}, nil

		// Group=serving.knative.dev, Version=v1alpha1
	case servingv1alpha1.SchemeGroupVersion.WithResource("configurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1alpha1().Configurations().Informer()}, nil
	case servingv1alpha1.SchemeGroupVersion.WithResource("revisions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1alpha1().Revisions().Informer()}, nil
	case servingv1alpha1.SchemeGroupVersion.WithResource("routes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1alpha1().Routes().Informer()}, nil
	case servingv1alpha1.SchemeGroupVersion.WithResource("services"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1alpha1().Services().Informer()}, nil

		// Group=serving.knative.dev, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("configurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1beta1().Configurations().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("revisions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1beta1().Revisions().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("routes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1beta1().Routes().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("services"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Serving().V1beta1().Services().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

/*
 * Copyright (c) 2018 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/cellery-io/mesh-controller/pkg/apis/mesh/v1alpha1"
	scheme "github.com/cellery-io/mesh-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WebCellsGetter has a method to return a WebCellInterface.
// A group's client should implement this interface.
type WebCellsGetter interface {
	WebCells(namespace string) WebCellInterface
}

// WebCellInterface has methods to work with WebCell resources.
type WebCellInterface interface {
	Create(*v1alpha1.WebCell) (*v1alpha1.WebCell, error)
	Update(*v1alpha1.WebCell) (*v1alpha1.WebCell, error)
	UpdateStatus(*v1alpha1.WebCell) (*v1alpha1.WebCell, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.WebCell, error)
	List(opts v1.ListOptions) (*v1alpha1.WebCellList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WebCell, err error)
	WebCellExpansion
}

// webCells implements WebCellInterface
type webCells struct {
	client rest.Interface
	ns     string
}

// newWebCells returns a WebCells
func newWebCells(c *MeshV1alpha1Client, namespace string) *webCells {
	return &webCells{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the webCell, and returns the corresponding webCell object, and an error if there is any.
func (c *webCells) Get(name string, options v1.GetOptions) (result *v1alpha1.WebCell, err error) {
	result = &v1alpha1.WebCell{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("webcells").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WebCells that match those selectors.
func (c *webCells) List(opts v1.ListOptions) (result *v1alpha1.WebCellList, err error) {
	result = &v1alpha1.WebCellList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("webcells").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested webCells.
func (c *webCells) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("webcells").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a webCell and creates it.  Returns the server's representation of the webCell, and an error, if there is any.
func (c *webCells) Create(webCell *v1alpha1.WebCell) (result *v1alpha1.WebCell, err error) {
	result = &v1alpha1.WebCell{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("webcells").
		Body(webCell).
		Do().
		Into(result)
	return
}

// Update takes the representation of a webCell and updates it. Returns the server's representation of the webCell, and an error, if there is any.
func (c *webCells) Update(webCell *v1alpha1.WebCell) (result *v1alpha1.WebCell, err error) {
	result = &v1alpha1.WebCell{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("webcells").
		Name(webCell.Name).
		Body(webCell).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *webCells) UpdateStatus(webCell *v1alpha1.WebCell) (result *v1alpha1.WebCell, err error) {
	result = &v1alpha1.WebCell{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("webcells").
		Name(webCell.Name).
		SubResource("status").
		Body(webCell).
		Do().
		Into(result)
	return
}

// Delete takes name of the webCell and deletes it. Returns an error if one occurs.
func (c *webCells) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("webcells").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *webCells) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("webcells").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched webCell.
func (c *webCells) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WebCell, err error) {
	result = &v1alpha1.WebCell{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("webcells").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

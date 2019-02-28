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

package fake

import (
	v1alpha1 "github.com/cellery-io/mesh-controller/pkg/apis/mesh/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStandardCells implements StandardCellInterface
type FakeStandardCells struct {
	Fake *FakeMeshV1alpha1
	ns   string
}

var standardcellsResource = schema.GroupVersionResource{Group: "mesh", Version: "v1alpha1", Resource: "standardcells"}

var standardcellsKind = schema.GroupVersionKind{Group: "mesh", Version: "v1alpha1", Kind: "StandardCell"}

// Get takes name of the standardCell, and returns the corresponding standardCell object, and an error if there is any.
func (c *FakeStandardCells) Get(name string, options v1.GetOptions) (result *v1alpha1.StandardCell, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(standardcellsResource, c.ns, name), &v1alpha1.StandardCell{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StandardCell), err
}

// List takes label and field selectors, and returns the list of StandardCells that match those selectors.
func (c *FakeStandardCells) List(opts v1.ListOptions) (result *v1alpha1.StandardCellList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(standardcellsResource, standardcellsKind, c.ns, opts), &v1alpha1.StandardCellList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StandardCellList{ListMeta: obj.(*v1alpha1.StandardCellList).ListMeta}
	for _, item := range obj.(*v1alpha1.StandardCellList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested standardCells.
func (c *FakeStandardCells) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(standardcellsResource, c.ns, opts))

}

// Create takes the representation of a standardCell and creates it.  Returns the server's representation of the standardCell, and an error, if there is any.
func (c *FakeStandardCells) Create(standardCell *v1alpha1.StandardCell) (result *v1alpha1.StandardCell, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(standardcellsResource, c.ns, standardCell), &v1alpha1.StandardCell{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StandardCell), err
}

// Update takes the representation of a standardCell and updates it. Returns the server's representation of the standardCell, and an error, if there is any.
func (c *FakeStandardCells) Update(standardCell *v1alpha1.StandardCell) (result *v1alpha1.StandardCell, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(standardcellsResource, c.ns, standardCell), &v1alpha1.StandardCell{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StandardCell), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStandardCells) UpdateStatus(standardCell *v1alpha1.StandardCell) (*v1alpha1.StandardCell, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(standardcellsResource, "status", c.ns, standardCell), &v1alpha1.StandardCell{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StandardCell), err
}

// Delete takes name of the standardCell and deletes it. Returns an error if one occurs.
func (c *FakeStandardCells) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(standardcellsResource, c.ns, name), &v1alpha1.StandardCell{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStandardCells) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(standardcellsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StandardCellList{})
	return err
}

// Patch applies the patch and returns the patched standardCell.
func (c *FakeStandardCells) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StandardCell, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(standardcellsResource, c.ns, name, data, subresources...), &v1alpha1.StandardCell{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StandardCell), err
}

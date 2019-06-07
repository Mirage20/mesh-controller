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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/cellery-io/mesh-controller/pkg/apis/knative/serving/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRevisions implements RevisionInterface
type FakeRevisions struct {
	Fake *FakeServingV1beta1
	ns   string
}

var revisionsResource = schema.GroupVersionResource{Group: "serving.knative.dev", Version: "v1beta1", Resource: "revisions"}

var revisionsKind = schema.GroupVersionKind{Group: "serving.knative.dev", Version: "v1beta1", Kind: "Revision"}

// Get takes name of the revision, and returns the corresponding revision object, and an error if there is any.
func (c *FakeRevisions) Get(name string, options v1.GetOptions) (result *v1beta1.Revision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(revisionsResource, c.ns, name), &v1beta1.Revision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Revision), err
}

// List takes label and field selectors, and returns the list of Revisions that match those selectors.
func (c *FakeRevisions) List(opts v1.ListOptions) (result *v1beta1.RevisionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(revisionsResource, revisionsKind, c.ns, opts), &v1beta1.RevisionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.RevisionList{ListMeta: obj.(*v1beta1.RevisionList).ListMeta}
	for _, item := range obj.(*v1beta1.RevisionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested revisions.
func (c *FakeRevisions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(revisionsResource, c.ns, opts))

}

// Create takes the representation of a revision and creates it.  Returns the server's representation of the revision, and an error, if there is any.
func (c *FakeRevisions) Create(revision *v1beta1.Revision) (result *v1beta1.Revision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(revisionsResource, c.ns, revision), &v1beta1.Revision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Revision), err
}

// Update takes the representation of a revision and updates it. Returns the server's representation of the revision, and an error, if there is any.
func (c *FakeRevisions) Update(revision *v1beta1.Revision) (result *v1beta1.Revision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(revisionsResource, c.ns, revision), &v1beta1.Revision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Revision), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRevisions) UpdateStatus(revision *v1beta1.Revision) (*v1beta1.Revision, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(revisionsResource, "status", c.ns, revision), &v1beta1.Revision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Revision), err
}

// Delete takes name of the revision and deletes it. Returns an error if one occurs.
func (c *FakeRevisions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(revisionsResource, c.ns, name), &v1beta1.Revision{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRevisions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(revisionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.RevisionList{})
	return err
}

// Patch applies the patch and returns the patched revision.
func (c *FakeRevisions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Revision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(revisionsResource, c.ns, name, data, subresources...), &v1beta1.Revision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Revision), err
}

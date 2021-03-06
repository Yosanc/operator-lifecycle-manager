/*
Copyright 2020 Red Hat, Inc.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	operatorsv1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOperatorGroups implements OperatorGroupInterface
type FakeOperatorGroups struct {
	Fake *FakeOperatorsV1
	ns   string
}

var operatorgroupsResource = schema.GroupVersionResource{Group: "operators.coreos.com", Version: "v1", Resource: "operatorgroups"}

var operatorgroupsKind = schema.GroupVersionKind{Group: "operators.coreos.com", Version: "v1", Kind: "OperatorGroup"}

// Get takes name of the operatorGroup, and returns the corresponding operatorGroup object, and an error if there is any.
func (c *FakeOperatorGroups) Get(name string, options v1.GetOptions) (result *operatorsv1.OperatorGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(operatorgroupsResource, c.ns, name), &operatorsv1.OperatorGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OperatorGroup), err
}

// List takes label and field selectors, and returns the list of OperatorGroups that match those selectors.
func (c *FakeOperatorGroups) List(opts v1.ListOptions) (result *operatorsv1.OperatorGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(operatorgroupsResource, operatorgroupsKind, c.ns, opts), &operatorsv1.OperatorGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorsv1.OperatorGroupList{ListMeta: obj.(*operatorsv1.OperatorGroupList).ListMeta}
	for _, item := range obj.(*operatorsv1.OperatorGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested operatorGroups.
func (c *FakeOperatorGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(operatorgroupsResource, c.ns, opts))

}

// Create takes the representation of a operatorGroup and creates it.  Returns the server's representation of the operatorGroup, and an error, if there is any.
func (c *FakeOperatorGroups) Create(operatorGroup *operatorsv1.OperatorGroup) (result *operatorsv1.OperatorGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(operatorgroupsResource, c.ns, operatorGroup), &operatorsv1.OperatorGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OperatorGroup), err
}

// Update takes the representation of a operatorGroup and updates it. Returns the server's representation of the operatorGroup, and an error, if there is any.
func (c *FakeOperatorGroups) Update(operatorGroup *operatorsv1.OperatorGroup) (result *operatorsv1.OperatorGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(operatorgroupsResource, c.ns, operatorGroup), &operatorsv1.OperatorGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OperatorGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOperatorGroups) UpdateStatus(operatorGroup *operatorsv1.OperatorGroup) (*operatorsv1.OperatorGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(operatorgroupsResource, "status", c.ns, operatorGroup), &operatorsv1.OperatorGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OperatorGroup), err
}

// Delete takes name of the operatorGroup and deletes it. Returns an error if one occurs.
func (c *FakeOperatorGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(operatorgroupsResource, c.ns, name), &operatorsv1.OperatorGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOperatorGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(operatorgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &operatorsv1.OperatorGroupList{})
	return err
}

// Patch applies the patch and returns the patched operatorGroup.
func (c *FakeOperatorGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *operatorsv1.OperatorGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(operatorgroupsResource, c.ns, name, pt, data, subresources...), &operatorsv1.OperatorGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OperatorGroup), err
}

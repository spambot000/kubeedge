/*
Copyright The KubeEdge Authors.

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
	"context"

	v1alpha1 "github.com/kubeedge/kubeedge/pkg/apis/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceAccountAccesses implements ServiceAccountAccessInterface
type FakeServiceAccountAccesses struct {
	Fake *FakePolicyV1alpha1
	ns   string
}

var serviceaccountaccessesResource = v1alpha1.SchemeGroupVersion.WithResource("serviceaccountaccesses")

var serviceaccountaccessesKind = v1alpha1.SchemeGroupVersion.WithKind("ServiceAccountAccess")

// Get takes name of the serviceAccountAccess, and returns the corresponding serviceAccountAccess object, and an error if there is any.
func (c *FakeServiceAccountAccesses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceAccountAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serviceaccountaccessesResource, c.ns, name), &v1alpha1.ServiceAccountAccess{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAccountAccess), err
}

// List takes label and field selectors, and returns the list of ServiceAccountAccesses that match those selectors.
func (c *FakeServiceAccountAccesses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceAccountAccessList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serviceaccountaccessesResource, serviceaccountaccessesKind, c.ns, opts), &v1alpha1.ServiceAccountAccessList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceAccountAccessList{ListMeta: obj.(*v1alpha1.ServiceAccountAccessList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceAccountAccessList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceAccountAccesses.
func (c *FakeServiceAccountAccesses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serviceaccountaccessesResource, c.ns, opts))

}

// Create takes the representation of a serviceAccountAccess and creates it.  Returns the server's representation of the serviceAccountAccess, and an error, if there is any.
func (c *FakeServiceAccountAccesses) Create(ctx context.Context, serviceAccountAccess *v1alpha1.ServiceAccountAccess, opts v1.CreateOptions) (result *v1alpha1.ServiceAccountAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serviceaccountaccessesResource, c.ns, serviceAccountAccess), &v1alpha1.ServiceAccountAccess{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAccountAccess), err
}

// Update takes the representation of a serviceAccountAccess and updates it. Returns the server's representation of the serviceAccountAccess, and an error, if there is any.
func (c *FakeServiceAccountAccesses) Update(ctx context.Context, serviceAccountAccess *v1alpha1.ServiceAccountAccess, opts v1.UpdateOptions) (result *v1alpha1.ServiceAccountAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serviceaccountaccessesResource, c.ns, serviceAccountAccess), &v1alpha1.ServiceAccountAccess{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAccountAccess), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceAccountAccesses) UpdateStatus(ctx context.Context, serviceAccountAccess *v1alpha1.ServiceAccountAccess, opts v1.UpdateOptions) (*v1alpha1.ServiceAccountAccess, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serviceaccountaccessesResource, "status", c.ns, serviceAccountAccess), &v1alpha1.ServiceAccountAccess{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAccountAccess), err
}

// Delete takes name of the serviceAccountAccess and deletes it. Returns an error if one occurs.
func (c *FakeServiceAccountAccesses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(serviceaccountaccessesResource, c.ns, name, opts), &v1alpha1.ServiceAccountAccess{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceAccountAccesses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serviceaccountaccessesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceAccountAccessList{})
	return err
}

// Patch applies the patch and returns the patched serviceAccountAccess.
func (c *FakeServiceAccountAccesses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceAccountAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serviceaccountaccessesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServiceAccountAccess{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAccountAccess), err
}

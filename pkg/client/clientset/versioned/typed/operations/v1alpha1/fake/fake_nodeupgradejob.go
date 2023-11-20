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

	v1alpha1 "github.com/kubeedge/kubeedge/pkg/apis/operations/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeUpgradeJobs implements NodeUpgradeJobInterface
type FakeNodeUpgradeJobs struct {
	Fake *FakeOperationsV1alpha1
}

var nodeupgradejobsResource = v1alpha1.SchemeGroupVersion.WithResource("nodeupgradejobs")

var nodeupgradejobsKind = v1alpha1.SchemeGroupVersion.WithKind("NodeUpgradeJob")

// Get takes name of the nodeUpgradeJob, and returns the corresponding nodeUpgradeJob object, and an error if there is any.
func (c *FakeNodeUpgradeJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodeUpgradeJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodeupgradejobsResource, name), &v1alpha1.NodeUpgradeJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeUpgradeJob), err
}

// List takes label and field selectors, and returns the list of NodeUpgradeJobs that match those selectors.
func (c *FakeNodeUpgradeJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeUpgradeJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodeupgradejobsResource, nodeupgradejobsKind, opts), &v1alpha1.NodeUpgradeJobList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeUpgradeJobList{ListMeta: obj.(*v1alpha1.NodeUpgradeJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodeUpgradeJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeUpgradeJobs.
func (c *FakeNodeUpgradeJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodeupgradejobsResource, opts))
}

// Create takes the representation of a nodeUpgradeJob and creates it.  Returns the server's representation of the nodeUpgradeJob, and an error, if there is any.
func (c *FakeNodeUpgradeJobs) Create(ctx context.Context, nodeUpgradeJob *v1alpha1.NodeUpgradeJob, opts v1.CreateOptions) (result *v1alpha1.NodeUpgradeJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodeupgradejobsResource, nodeUpgradeJob), &v1alpha1.NodeUpgradeJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeUpgradeJob), err
}

// Update takes the representation of a nodeUpgradeJob and updates it. Returns the server's representation of the nodeUpgradeJob, and an error, if there is any.
func (c *FakeNodeUpgradeJobs) Update(ctx context.Context, nodeUpgradeJob *v1alpha1.NodeUpgradeJob, opts v1.UpdateOptions) (result *v1alpha1.NodeUpgradeJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodeupgradejobsResource, nodeUpgradeJob), &v1alpha1.NodeUpgradeJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeUpgradeJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeUpgradeJobs) UpdateStatus(ctx context.Context, nodeUpgradeJob *v1alpha1.NodeUpgradeJob, opts v1.UpdateOptions) (*v1alpha1.NodeUpgradeJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nodeupgradejobsResource, "status", nodeUpgradeJob), &v1alpha1.NodeUpgradeJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeUpgradeJob), err
}

// Delete takes name of the nodeUpgradeJob and deletes it. Returns an error if one occurs.
func (c *FakeNodeUpgradeJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(nodeupgradejobsResource, name, opts), &v1alpha1.NodeUpgradeJob{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeUpgradeJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodeupgradejobsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeUpgradeJobList{})
	return err
}

// Patch applies the patch and returns the patched nodeUpgradeJob.
func (c *FakeNodeUpgradeJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeUpgradeJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodeupgradejobsResource, name, pt, data, subresources...), &v1alpha1.NodeUpgradeJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeUpgradeJob), err
}

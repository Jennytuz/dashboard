/*
Copyright 2016 The Kubernetes Authors.

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

package fake

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	core "k8s.io/kubernetes/pkg/client/testing/core"
	labels "k8s.io/kubernetes/pkg/labels"
	watch "k8s.io/kubernetes/pkg/watch"
)

// FakeNodes implements NodeInterface
type FakeNodes struct {
	Fake *FakeCore
}

var nodesResource = unversioned.GroupVersionResource{Group: "", Version: "", Resource: "nodes"}

func (c *FakeNodes) Create(node *api.Node) (result *api.Node, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootCreateAction(nodesResource, node), &api.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*api.Node), err
}

func (c *FakeNodes) Update(node *api.Node) (result *api.Node, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootUpdateAction(nodesResource, node), &api.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*api.Node), err
}

func (c *FakeNodes) UpdateStatus(node *api.Node) (*api.Node, error) {
	obj, err := c.Fake.
		Invokes(core.NewRootUpdateSubresourceAction(nodesResource, "status", node), &api.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*api.Node), err
}

func (c *FakeNodes) Delete(name string, options *api.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewRootDeleteAction(nodesResource, name), &api.Node{})
	return err
}

func (c *FakeNodes) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	action := core.NewRootDeleteCollectionAction(nodesResource, listOptions)

	_, err := c.Fake.Invokes(action, &api.NodeList{})
	return err
}

func (c *FakeNodes) Get(name string) (result *api.Node, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootGetAction(nodesResource, name), &api.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*api.Node), err
}

func (c *FakeNodes) List(opts api.ListOptions) (result *api.NodeList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootListAction(nodesResource, opts), &api.NodeList{})
	if obj == nil {
		return nil, err
	}

	label := opts.LabelSelector
	if label == nil {
		label = labels.Everything()
	}
	list := &api.NodeList{}
	for _, item := range obj.(*api.NodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodes.
func (c *FakeNodes) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewRootWatchAction(nodesResource, opts))
}

// Patch applies the patch and returns the patched node.
func (c *FakeNodes) Patch(name string, pt api.PatchType, data []byte) (result *api.Node, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootPatchAction(nodesResource, name, data), &api.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*api.Node), err
}

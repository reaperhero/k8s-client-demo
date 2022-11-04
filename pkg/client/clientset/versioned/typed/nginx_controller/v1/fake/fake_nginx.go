/*
Copyright The Kubernetes Authors.

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
	nginxcontrollerv1 "github.com/reaperhero/k8s-client-demo/pkg/apis/nginx_controller/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNginxes implements NginxInterface
type FakeNginxes struct {
	Fake *FakeMycompanyV1
	ns   string
}

var nginxesResource = schema.GroupVersionResource{Group: "mycompany.com", Version: "v1", Resource: "nginxes"}

var nginxesKind = schema.GroupVersionKind{Group: "mycompany.com", Version: "v1", Kind: "Nginx"}

// Get takes name of the nginx, and returns the corresponding nginx object, and an error if there is any.
func (c *FakeNginxes) Get(name string, options v1.GetOptions) (result *nginxcontrollerv1.Nginx, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nginxesResource, c.ns, name), &nginxcontrollerv1.Nginx{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nginxcontrollerv1.Nginx), err
}

// List takes label and field selectors, and returns the list of Nginxes that match those selectors.
func (c *FakeNginxes) List(opts v1.ListOptions) (result *nginxcontrollerv1.NginxList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nginxesResource, nginxesKind, c.ns, opts), &nginxcontrollerv1.NginxList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &nginxcontrollerv1.NginxList{ListMeta: obj.(*nginxcontrollerv1.NginxList).ListMeta}
	for _, item := range obj.(*nginxcontrollerv1.NginxList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nginxes.
func (c *FakeNginxes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nginxesResource, c.ns, opts))

}

// Create takes the representation of a nginx and creates it.  Returns the server's representation of the nginx, and an error, if there is any.
func (c *FakeNginxes) Create(nginx *nginxcontrollerv1.Nginx) (result *nginxcontrollerv1.Nginx, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nginxesResource, c.ns, nginx), &nginxcontrollerv1.Nginx{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nginxcontrollerv1.Nginx), err
}

// Update takes the representation of a nginx and updates it. Returns the server's representation of the nginx, and an error, if there is any.
func (c *FakeNginxes) Update(nginx *nginxcontrollerv1.Nginx) (result *nginxcontrollerv1.Nginx, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nginxesResource, c.ns, nginx), &nginxcontrollerv1.Nginx{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nginxcontrollerv1.Nginx), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNginxes) UpdateStatus(nginx *nginxcontrollerv1.Nginx) (*nginxcontrollerv1.Nginx, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nginxesResource, "status", c.ns, nginx), &nginxcontrollerv1.Nginx{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nginxcontrollerv1.Nginx), err
}

// Delete takes name of the nginx and deletes it. Returns an error if one occurs.
func (c *FakeNginxes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nginxesResource, c.ns, name), &nginxcontrollerv1.Nginx{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNginxes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nginxesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &nginxcontrollerv1.NginxList{})
	return err
}

// Patch applies the patch and returns the patched nginx.
func (c *FakeNginxes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *nginxcontrollerv1.Nginx, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nginxesResource, c.ns, name, pt, data, subresources...), &nginxcontrollerv1.Nginx{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nginxcontrollerv1.Nginx), err
}

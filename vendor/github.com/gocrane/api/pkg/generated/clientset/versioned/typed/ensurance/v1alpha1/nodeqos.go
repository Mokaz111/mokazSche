// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/gocrane/api/ensurance/v1alpha1"
	scheme "github.com/gocrane/api/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodeQOSsGetter has a method to return a NodeQOSInterface.
// A group's client should implement this interface.
type NodeQOSsGetter interface {
	NodeQOSs() NodeQOSInterface
}

// NodeQOSInterface has methods to work with NodeQOS resources.
type NodeQOSInterface interface {
	Create(ctx context.Context, nodeQOS *v1alpha1.NodeQOS, opts v1.CreateOptions) (*v1alpha1.NodeQOS, error)
	Update(ctx context.Context, nodeQOS *v1alpha1.NodeQOS, opts v1.UpdateOptions) (*v1alpha1.NodeQOS, error)
	UpdateStatus(ctx context.Context, nodeQOS *v1alpha1.NodeQOS, opts v1.UpdateOptions) (*v1alpha1.NodeQOS, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NodeQOS, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NodeQOSList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeQOS, err error)
	NodeQOSExpansion
}

// nodeQOSs implements NodeQOSInterface
type nodeQOSs struct {
	client rest.Interface
}

// newNodeQOSs returns a NodeQOSs
func newNodeQOSs(c *EnsuranceV1alpha1Client) *nodeQOSs {
	return &nodeQOSs{
		client: c.RESTClient(),
	}
}

// Get takes name of the nodeQOS, and returns the corresponding nodeQOS object, and an error if there is any.
func (c *nodeQOSs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodeQOS, err error) {
	result = &v1alpha1.NodeQOS{}
	err = c.client.Get().
		Resource("nodeqoss").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeQOSs that match those selectors.
func (c *nodeQOSs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeQOSList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NodeQOSList{}
	err = c.client.Get().
		Resource("nodeqoss").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeQOSs.
func (c *nodeQOSs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("nodeqoss").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a nodeQOS and creates it.  Returns the server's representation of the nodeQOS, and an error, if there is any.
func (c *nodeQOSs) Create(ctx context.Context, nodeQOS *v1alpha1.NodeQOS, opts v1.CreateOptions) (result *v1alpha1.NodeQOS, err error) {
	result = &v1alpha1.NodeQOS{}
	err = c.client.Post().
		Resource("nodeqoss").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeQOS).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a nodeQOS and updates it. Returns the server's representation of the nodeQOS, and an error, if there is any.
func (c *nodeQOSs) Update(ctx context.Context, nodeQOS *v1alpha1.NodeQOS, opts v1.UpdateOptions) (result *v1alpha1.NodeQOS, err error) {
	result = &v1alpha1.NodeQOS{}
	err = c.client.Put().
		Resource("nodeqoss").
		Name(nodeQOS.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeQOS).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *nodeQOSs) UpdateStatus(ctx context.Context, nodeQOS *v1alpha1.NodeQOS, opts v1.UpdateOptions) (result *v1alpha1.NodeQOS, err error) {
	result = &v1alpha1.NodeQOS{}
	err = c.client.Put().
		Resource("nodeqoss").
		Name(nodeQOS.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeQOS).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the nodeQOS and deletes it. Returns an error if one occurs.
func (c *nodeQOSs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nodeqoss").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeQOSs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("nodeqoss").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched nodeQOS.
func (c *nodeQOSs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeQOS, err error) {
	result = &v1alpha1.NodeQOS{}
	err = c.client.Patch(pt).
		Resource("nodeqoss").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
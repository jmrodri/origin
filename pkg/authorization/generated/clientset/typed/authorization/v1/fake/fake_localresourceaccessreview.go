package fake

import (
	v1 "github.com/openshift/origin/pkg/authorization/apis/authorization/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeLocalResourceAccessReviews implements LocalResourceAccessReviewInterface
type FakeLocalResourceAccessReviews struct {
	Fake *FakeAuthorizationV1
	ns   string
}

var localresourceaccessreviewsResource = schema.GroupVersionResource{Group: "authorization.openshift.io", Version: "v1", Resource: "localresourceaccessreviews"}

var localresourceaccessreviewsKind = schema.GroupVersionKind{Group: "authorization.openshift.io", Version: "v1", Kind: "LocalResourceAccessReview"}

// Create takes the representation of a localResourceAccessReview and creates it.  Returns the server's representation of the localResourceAccessReview, and an error, if there is any.
func (c *FakeLocalResourceAccessReviews) Create(localResourceAccessReview *v1.LocalResourceAccessReview) (result *v1.LocalResourceAccessReview, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(localresourceaccessreviewsResource, c.ns, localResourceAccessReview), &v1.LocalResourceAccessReview{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.LocalResourceAccessReview), err
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// EnterpriseDirCommentService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnterpriseDirCommentService] method instead.
type EnterpriseDirCommentService struct {
	Options []option.RequestOption
}

// NewEnterpriseDirCommentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnterpriseDirCommentService(opts ...option.RequestOption) (r EnterpriseDirCommentService) {
	r = EnterpriseDirCommentService{}
	r.Options = opts
	return
}

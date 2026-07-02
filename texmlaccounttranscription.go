// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// TexmlAccountTranscriptionService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountTranscriptionService] method instead.
type TexmlAccountTranscriptionService struct {
	Options []option.RequestOption
	// TeXML REST Commands
	Json TexmlAccountTranscriptionJsonService
}

// NewTexmlAccountTranscriptionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTexmlAccountTranscriptionService(opts ...option.RequestOption) (r TexmlAccountTranscriptionService) {
	r = TexmlAccountTranscriptionService{}
	r.Options = opts
	r.Json = NewTexmlAccountTranscriptionJsonService(opts...)
	return
}

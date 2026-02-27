// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// AIOpenAIService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOpenAIService] method instead.
type AIOpenAIService struct {
	Options []option.RequestOption
	// OpenAI-compatible embeddings endpoints for generating vector representations of
	// text
	Embeddings AIOpenAIEmbeddingService
}

// NewAIOpenAIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIOpenAIService(opts ...option.RequestOption) (r AIOpenAIService) {
	r = AIOpenAIService{}
	r.Options = opts
	r.Embeddings = NewAIOpenAIEmbeddingService(opts...)
	return
}

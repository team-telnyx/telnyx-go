// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// AIKnowledgeService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIKnowledgeService] method instead.
type AIKnowledgeService struct {
	Options []option.RequestOption
	// Create and manage logical collections of your Telnyx data, tune retrieval
	// settings, manage sources, and run collection-scoped semantic search.
	Collections AIKnowledgeCollectionService
}

// NewAIKnowledgeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIKnowledgeService(opts ...option.RequestOption) (r AIKnowledgeService) {
	r = AIKnowledgeService{}
	r.Options = opts
	r.Collections = NewAIKnowledgeCollectionService(opts...)
	return
}

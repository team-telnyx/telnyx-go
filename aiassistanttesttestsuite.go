// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// AIAssistantTestTestSuiteService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantTestTestSuiteService] method instead.
type AIAssistantTestTestSuiteService struct {
	Options []option.RequestOption
	Runs    AIAssistantTestTestSuiteRunService
}

// NewAIAssistantTestTestSuiteService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIAssistantTestTestSuiteService(opts ...option.RequestOption) (r AIAssistantTestTestSuiteService) {
	r = AIAssistantTestTestSuiteService{}
	r.Options = opts
	r.Runs = NewAIAssistantTestTestSuiteRunService(opts...)
	return
}

// Retrieves a list of all distinct test suite names available to the current user
func (r *AIAssistantTestTestSuiteService) List(ctx context.Context, opts ...option.RequestOption) (res *AIAssistantTestTestSuiteListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/assistants/tests/test-suites"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response containing all available test suite names.
//
// Returns a list of distinct test suite names that can be used for filtering and
// organizing tests.
type AIAssistantTestTestSuiteListResponse struct {
	// Array of unique test suite names available to the user.
	Data []string `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAssistantTestTestSuiteListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantTestTestSuiteListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

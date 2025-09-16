// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// MessageRcService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageRcService] method instead.
type MessageRcService struct {
	Options []option.RequestOption
}

// NewMessageRcService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessageRcService(opts ...option.RequestOption) (r MessageRcService) {
	r = MessageRcService{}
	r.Options = opts
	return
}

// Generate a deeplink URL that can be used to start an RCS conversation with a
// specific agent.
func (r *MessageRcService) GenerateDeeplink(ctx context.Context, agentID string, query MessageRcGenerateDeeplinkParams, opts ...option.RequestOption) (res *MessageRcGenerateDeeplinkResponse, err error) {
	opts = append(r.Options[:], opts...)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("messages/rcs/deeplinks/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MessageRcGenerateDeeplinkResponse struct {
	Data MessageRcGenerateDeeplinkResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageRcGenerateDeeplinkResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageRcGenerateDeeplinkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRcGenerateDeeplinkResponseData struct {
	// The generated deeplink URL
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageRcGenerateDeeplinkResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessageRcGenerateDeeplinkResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRcGenerateDeeplinkParams struct {
	// Pre-filled message body (URL encoded)
	Body param.Opt[string] `query:"body,omitzero" json:"-"`
	// Phone number in E164 format (URL encoded)
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageRcGenerateDeeplinkParams]'s query parameters as
// `url.Values`.
func (r MessageRcGenerateDeeplinkParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

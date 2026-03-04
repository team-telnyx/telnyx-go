// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// MessagingProfileMetricService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingProfileMetricService] method instead.
type MessagingProfileMetricService struct {
	Options []option.RequestOption
}

// NewMessagingProfileMetricService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingProfileMetricService(opts ...option.RequestOption) (r MessagingProfileMetricService) {
	r = MessagingProfileMetricService{}
	r.Options = opts
	return
}

// List high-level metrics for all messaging profiles belonging to the
// authenticated user.
func (r *MessagingProfileMetricService) List(ctx context.Context, query MessagingProfileMetricListParams, opts ...option.RequestOption) (res *MessagingProfileMetricListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messaging_profile_metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MessagingProfileMetricListResponse struct {
	Data []map[string]any               `json:"data"`
	Meta shared.MessagingPaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingProfileMetricListResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingProfileMetricListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingProfileMetricListParams struct {
	// The time frame for metrics.
	//
	// Any of "1h", "3h", "24h", "3d", "7d", "30d".
	TimeFrame MessagingProfileMetricListParamsTimeFrame `query:"time_frame,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingProfileMetricListParams]'s query parameters as
// `url.Values`.
func (r MessagingProfileMetricListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The time frame for metrics.
type MessagingProfileMetricListParamsTimeFrame string

const (
	MessagingProfileMetricListParamsTimeFrame1h  MessagingProfileMetricListParamsTimeFrame = "1h"
	MessagingProfileMetricListParamsTimeFrame3h  MessagingProfileMetricListParamsTimeFrame = "3h"
	MessagingProfileMetricListParamsTimeFrame24h MessagingProfileMetricListParamsTimeFrame = "24h"
	MessagingProfileMetricListParamsTimeFrame3d  MessagingProfileMetricListParamsTimeFrame = "3d"
	MessagingProfileMetricListParamsTimeFrame7d  MessagingProfileMetricListParamsTimeFrame = "7d"
	MessagingProfileMetricListParamsTimeFrame30d MessagingProfileMetricListParamsTimeFrame = "30d"
)

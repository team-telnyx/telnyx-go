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

// ExternalConnectionLogMessageService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalConnectionLogMessageService] method instead.
type ExternalConnectionLogMessageService struct {
	Options []option.RequestOption
}

// NewExternalConnectionLogMessageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExternalConnectionLogMessageService(opts ...option.RequestOption) (r ExternalConnectionLogMessageService) {
	r = ExternalConnectionLogMessageService{}
	r.Options = opts
	return
}

// Retrieve a log message for an external connection associated with your account.
func (r *ExternalConnectionLogMessageService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ExternalConnectionLogMessageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/log_messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of log messages for all external connections associated with
// your account.
func (r *ExternalConnectionLogMessageService) List(ctx context.Context, query ExternalConnectionLogMessageListParams, opts ...option.RequestOption) (res *ExternalConnectionLogMessageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "external_connections/log_messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Dismiss a log message for an external connection associated with your account.
func (r *ExternalConnectionLogMessageService) Dismiss(ctx context.Context, id string, opts ...option.RequestOption) (res *ExternalConnectionLogMessageDismissResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/log_messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ExternalConnectionLogMessageGetResponse struct {
	LogMessages []ExternalConnectionLogMessageGetResponseLogMessage `json:"log_messages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LogMessages respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionLogMessageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionLogMessageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionLogMessageGetResponseLogMessage struct {
	Code   string                                                  `json:"code,required" format:"int64"`
	Title  string                                                  `json:"title,required"`
	Detail string                                                  `json:"detail"`
	Meta   ExternalConnectionLogMessageGetResponseLogMessageMeta   `json:"meta"`
	Source ExternalConnectionLogMessageGetResponseLogMessageSource `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Title       respjson.Field
		Detail      respjson.Field
		Meta        respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionLogMessageGetResponseLogMessage) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionLogMessageGetResponseLogMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionLogMessageGetResponseLogMessageMeta struct {
	// The external connection the log message is associated with, if any.
	ExternalConnectionID string `json:"external_connection_id" format:"int64"`
	// The telephone number the log message is associated with, if any.
	TelephoneNumber string `json:"telephone_number"`
	// The ticket ID for an operation that generated the log message, if any.
	TicketID string `json:"ticket_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExternalConnectionID respjson.Field
		TelephoneNumber      respjson.Field
		TicketID             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionLogMessageGetResponseLogMessageMeta) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionLogMessageGetResponseLogMessageMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionLogMessageGetResponseLogMessageSource struct {
	// JSON pointer (RFC6901) to the offending entity.
	Pointer string `json:"pointer" format:"json-pointer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pointer     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionLogMessageGetResponseLogMessageSource) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionLogMessageGetResponseLogMessageSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionLogMessageListResponse struct {
	LogMessages []ExternalConnectionLogMessageListResponseLogMessage `json:"log_messages"`
	Meta        ExternalVoiceIntegrationsPaginationMeta              `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LogMessages respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionLogMessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionLogMessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionLogMessageListResponseLogMessage struct {
	Code   string                                                   `json:"code,required" format:"int64"`
	Title  string                                                   `json:"title,required"`
	Detail string                                                   `json:"detail"`
	Meta   ExternalConnectionLogMessageListResponseLogMessageMeta   `json:"meta"`
	Source ExternalConnectionLogMessageListResponseLogMessageSource `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Title       respjson.Field
		Detail      respjson.Field
		Meta        respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionLogMessageListResponseLogMessage) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionLogMessageListResponseLogMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionLogMessageListResponseLogMessageMeta struct {
	// The external connection the log message is associated with, if any.
	ExternalConnectionID string `json:"external_connection_id" format:"int64"`
	// The telephone number the log message is associated with, if any.
	TelephoneNumber string `json:"telephone_number"`
	// The ticket ID for an operation that generated the log message, if any.
	TicketID string `json:"ticket_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExternalConnectionID respjson.Field
		TelephoneNumber      respjson.Field
		TicketID             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionLogMessageListResponseLogMessageMeta) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionLogMessageListResponseLogMessageMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionLogMessageListResponseLogMessageSource struct {
	// JSON pointer (RFC6901) to the offending entity.
	Pointer string `json:"pointer" format:"json-pointer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pointer     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionLogMessageListResponseLogMessageSource) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionLogMessageListResponseLogMessageSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionLogMessageDismissResponse struct {
	// Describes wether or not the operation was successful
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionLogMessageDismissResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionLogMessageDismissResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionLogMessageListParams struct {
	// Filter parameter for log messages (deepObject style). Supports filtering by
	// external_connection_id and telephone_number with eq/contains operations.
	Filter ExternalConnectionLogMessageListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page ExternalConnectionLogMessageListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionLogMessageListParams]'s query parameters
// as `url.Values`.
func (r ExternalConnectionLogMessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter parameter for log messages (deepObject style). Supports filtering by
// external_connection_id and telephone_number with eq/contains operations.
type ExternalConnectionLogMessageListParamsFilter struct {
	// The external connection ID to filter by or "null" to filter for logs without an
	// external connection ID
	ExternalConnectionID param.Opt[string] `query:"external_connection_id,omitzero" json:"-"`
	// Telephone number filter operations for log messages. Use 'eq' for exact matches
	// or 'contains' for partial matches.
	TelephoneNumber ExternalConnectionLogMessageListParamsFilterTelephoneNumber `query:"telephone_number,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionLogMessageListParamsFilter]'s query
// parameters as `url.Values`.
func (r ExternalConnectionLogMessageListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Telephone number filter operations for log messages. Use 'eq' for exact matches
// or 'contains' for partial matches.
type ExternalConnectionLogMessageListParamsFilterTelephoneNumber struct {
	// The partial phone number to filter log messages for. Requires 3-15 digits.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// The phone number to filter log messages for or "null" to filter for logs without
	// a phone number
	Eq param.Opt[string] `query:"eq,omitzero" format:"E164" json:"-"`
	paramObj
}

// URLQuery serializes
// [ExternalConnectionLogMessageListParamsFilterTelephoneNumber]'s query parameters
// as `url.Values`.
func (r ExternalConnectionLogMessageListParamsFilterTelephoneNumber) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type ExternalConnectionLogMessageListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionLogMessageListParamsPage]'s query
// parameters as `url.Values`.
func (r ExternalConnectionLogMessageListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

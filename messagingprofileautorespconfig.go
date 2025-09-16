// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/telnyx-go/internal/encoding/json"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// MessagingProfileAutorespConfigService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingProfileAutorespConfigService] method instead.
type MessagingProfileAutorespConfigService struct {
	Options []option.RequestOption
}

// NewMessagingProfileAutorespConfigService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessagingProfileAutorespConfigService(opts ...option.RequestOption) (r MessagingProfileAutorespConfigService) {
	r = MessagingProfileAutorespConfigService{}
	r.Options = opts
	return
}

// Create Auto-Reponse Setting
func (r *MessagingProfileAutorespConfigService) New(ctx context.Context, profileID string, body MessagingProfileAutorespConfigNewParams, opts ...option.RequestOption) (res *AutoRespConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("messaging_profiles/%s/autoresp_configs", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Auto-Response Setting
func (r *MessagingProfileAutorespConfigService) Get(ctx context.Context, autorespCfgID string, query MessagingProfileAutorespConfigGetParams, opts ...option.RequestOption) (res *AutoRespConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.ProfileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	if autorespCfgID == "" {
		err = errors.New("missing required autoresp_cfg_id parameter")
		return
	}
	path := fmt.Sprintf("messaging_profiles/%s/autoresp_configs/%s", query.ProfileID, autorespCfgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Auto-Response Setting
func (r *MessagingProfileAutorespConfigService) Update(ctx context.Context, autorespCfgID string, params MessagingProfileAutorespConfigUpdateParams, opts ...option.RequestOption) (res *AutoRespConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.ProfileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	if autorespCfgID == "" {
		err = errors.New("missing required autoresp_cfg_id parameter")
		return
	}
	path := fmt.Sprintf("messaging_profiles/%s/autoresp_configs/%s", params.ProfileID, autorespCfgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// List Auto-Response Settings
func (r *MessagingProfileAutorespConfigService) List(ctx context.Context, profileID string, query MessagingProfileAutorespConfigListParams, opts ...option.RequestOption) (res *MessagingProfileAutorespConfigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("messaging_profiles/%s/autoresp_configs", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete Auto-Response Setting
func (r *MessagingProfileAutorespConfigService) Delete(ctx context.Context, autorespCfgID string, body MessagingProfileAutorespConfigDeleteParams, opts ...option.RequestOption) (res *MessagingProfileAutorespConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ProfileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	if autorespCfgID == "" {
		err = errors.New("missing required autoresp_cfg_id parameter")
		return
	}
	path := fmt.Sprintf("messaging_profiles/%s/autoresp_configs/%s", body.ProfileID, autorespCfgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AutoRespConfig struct {
	ID          string    `json:"id,required"`
	CountryCode string    `json:"country_code,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	Keywords    []string  `json:"keywords,required"`
	// Any of "start", "stop", "info".
	Op        AutoRespConfigOp `json:"op,required"`
	UpdatedAt time.Time        `json:"updated_at,required" format:"date-time"`
	RespText  string           `json:"resp_text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CountryCode respjson.Field
		CreatedAt   respjson.Field
		Keywords    respjson.Field
		Op          respjson.Field
		UpdatedAt   respjson.Field
		RespText    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoRespConfig) RawJSON() string { return r.JSON.raw }
func (r *AutoRespConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutoRespConfigOp string

const (
	AutoRespConfigOpStart AutoRespConfigOp = "start"
	AutoRespConfigOpStop  AutoRespConfigOp = "stop"
	AutoRespConfigOpInfo  AutoRespConfigOp = "info"
)

// The properties CountryCode, Keywords, Op are required.
type AutoRespConfigCreateParam struct {
	CountryCode string   `json:"country_code,required"`
	Keywords    []string `json:"keywords,omitzero,required"`
	// Any of "start", "stop", "info".
	Op       AutoRespConfigCreateOp `json:"op,omitzero,required"`
	RespText param.Opt[string]      `json:"resp_text,omitzero"`
	paramObj
}

func (r AutoRespConfigCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow AutoRespConfigCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutoRespConfigCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutoRespConfigCreateOp string

const (
	AutoRespConfigCreateOpStart AutoRespConfigCreateOp = "start"
	AutoRespConfigCreateOpStop  AutoRespConfigCreateOp = "stop"
	AutoRespConfigCreateOpInfo  AutoRespConfigCreateOp = "info"
)

type AutoRespConfigResponse struct {
	Data AutoRespConfig `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoRespConfigResponse) RawJSON() string { return r.JSON.raw }
func (r *AutoRespConfigResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of Auto-Response Settings
type MessagingProfileAutorespConfigListResponse struct {
	Data []AutoRespConfig `json:"data,required"`
	Meta PaginationMeta   `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingProfileAutorespConfigListResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingProfileAutorespConfigListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingProfileAutorespConfigDeleteResponse = any

type MessagingProfileAutorespConfigNewParams struct {
	AutoRespConfigCreate AutoRespConfigCreateParam
	paramObj
}

func (r MessagingProfileAutorespConfigNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AutoRespConfigCreate)
}
func (r *MessagingProfileAutorespConfigNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AutoRespConfigCreate)
}

type MessagingProfileAutorespConfigGetParams struct {
	ProfileID string `path:"profile_id,required" format:"uuid" json:"-"`
	paramObj
}

type MessagingProfileAutorespConfigUpdateParams struct {
	ProfileID            string `path:"profile_id,required" format:"uuid" json:"-"`
	AutoRespConfigCreate AutoRespConfigCreateParam
	paramObj
}

func (r MessagingProfileAutorespConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AutoRespConfigCreate)
}
func (r *MessagingProfileAutorespConfigUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AutoRespConfigCreate)
}

type MessagingProfileAutorespConfigListParams struct {
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Consolidated created_at parameter (deepObject style). Originally:
	// created_at[gte], created_at[lte]
	CreatedAt MessagingProfileAutorespConfigListParamsCreatedAt `query:"created_at,omitzero" json:"-"`
	// Consolidated updated_at parameter (deepObject style). Originally:
	// updated_at[gte], updated_at[lte]
	UpdatedAt MessagingProfileAutorespConfigListParamsUpdatedAt `query:"updated_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingProfileAutorespConfigListParams]'s query
// parameters as `url.Values`.
func (r MessagingProfileAutorespConfigListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated created_at parameter (deepObject style). Originally:
// created_at[gte], created_at[lte]
type MessagingProfileAutorespConfigListParamsCreatedAt struct {
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingProfileAutorespConfigListParamsCreatedAt]'s query
// parameters as `url.Values`.
func (r MessagingProfileAutorespConfigListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated updated_at parameter (deepObject style). Originally:
// updated_at[gte], updated_at[lte]
type MessagingProfileAutorespConfigListParamsUpdatedAt struct {
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingProfileAutorespConfigListParamsUpdatedAt]'s query
// parameters as `url.Values`.
func (r MessagingProfileAutorespConfigListParamsUpdatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingProfileAutorespConfigDeleteParams struct {
	ProfileID string `path:"profile_id,required" format:"uuid" json:"-"`
	paramObj
}

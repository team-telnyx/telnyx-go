// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Traffic Policy Profiles operations
//
// TrafficPolicyProfileService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTrafficPolicyProfileService] method instead.
type TrafficPolicyProfileService struct {
	Options []option.RequestOption
}

// NewTrafficPolicyProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTrafficPolicyProfileService(opts ...option.RequestOption) (r TrafficPolicyProfileService) {
	r = TrafficPolicyProfileService{}
	r.Options = opts
	return
}

// Create a new traffic policy profile. At least one of `services`, `ip_ranges`, or
// `domains` must be provided.
func (r *TrafficPolicyProfileService) New(ctx context.Context, body TrafficPolicyProfileNewParams, opts ...option.RequestOption) (res *TrafficPolicyProfileNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "traffic_policy_profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the details regarding a specific traffic policy profile.
func (r *TrafficPolicyProfileService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TrafficPolicyProfileGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("traffic_policy_profiles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a traffic policy profile.
func (r *TrafficPolicyProfileService) Update(ctx context.Context, id string, body TrafficPolicyProfileUpdateParams, opts ...option.RequestOption) (res *TrafficPolicyProfileUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("traffic_policy_profiles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get all traffic policy profiles belonging to the user that match the given
// filters.
func (r *TrafficPolicyProfileService) List(ctx context.Context, query TrafficPolicyProfileListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[TrafficPolicyProfile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "traffic_policy_profiles"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get all traffic policy profiles belonging to the user that match the given
// filters.
func (r *TrafficPolicyProfileService) ListAutoPaging(ctx context.Context, query TrafficPolicyProfileListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[TrafficPolicyProfile] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes the traffic policy profile.
func (r *TrafficPolicyProfileService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TrafficPolicyProfileDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("traffic_policy_profiles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get all available PCEF services that can be used in traffic policy profiles.
func (r *TrafficPolicyProfileService) ListServices(ctx context.Context, query TrafficPolicyProfileListServicesParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[TrafficPolicyProfileListServicesResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "traffic_policy_profiles/services"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get all available PCEF services that can be used in traffic policy profiles.
func (r *TrafficPolicyProfileService) ListServicesAutoPaging(ctx context.Context, query TrafficPolicyProfileListServicesParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[TrafficPolicyProfileListServicesResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.ListServices(ctx, query, opts...))
}

type TrafficPolicyProfile struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Array of domain names.
	Domains []string `json:"domains"`
	// Array of IP ranges in CIDR notation.
	IPRanges []string `json:"ip_ranges"`
	// Bandwidth limit in kbps.
	LimitBwKbps int64  `json:"limit_bw_kbps" api:"nullable"`
	RecordType  string `json:"record_type"`
	// Array of PCEF service IDs included in the profile.
	Services []string `json:"services"`
	// The type of the traffic policy profile.
	//
	// Any of "whitelist", "blacklist", "throttling".
	Type TrafficPolicyProfileType `json:"type"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Domains     respjson.Field
		IPRanges    respjson.Field
		LimitBwKbps respjson.Field
		RecordType  respjson.Field
		Services    respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrafficPolicyProfile) RawJSON() string { return r.JSON.raw }
func (r *TrafficPolicyProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the traffic policy profile.
type TrafficPolicyProfileType string

const (
	TrafficPolicyProfileTypeWhitelist  TrafficPolicyProfileType = "whitelist"
	TrafficPolicyProfileTypeBlacklist  TrafficPolicyProfileType = "blacklist"
	TrafficPolicyProfileTypeThrottling TrafficPolicyProfileType = "throttling"
)

type TrafficPolicyProfileNewResponse struct {
	Data TrafficPolicyProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrafficPolicyProfileNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TrafficPolicyProfileNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrafficPolicyProfileGetResponse struct {
	Data TrafficPolicyProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrafficPolicyProfileGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TrafficPolicyProfileGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrafficPolicyProfileUpdateResponse struct {
	Data TrafficPolicyProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrafficPolicyProfileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TrafficPolicyProfileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrafficPolicyProfileDeleteResponse struct {
	Data TrafficPolicyProfileDeleteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrafficPolicyProfileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TrafficPolicyProfileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrafficPolicyProfileDeleteResponseData struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrafficPolicyProfileDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *TrafficPolicyProfileDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrafficPolicyProfileListServicesResponse struct {
	// The service identifier.
	ID string `json:"id"`
	// The group the service belongs to.
	Group string `json:"group"`
	// The name of the service.
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Group        respjson.Field
		Name         respjson.Field
		ResourceType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrafficPolicyProfileListServicesResponse) RawJSON() string { return r.JSON.raw }
func (r *TrafficPolicyProfileListServicesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrafficPolicyProfileNewParams struct {
	// The type of the traffic policy profile.
	//
	// Any of "whitelist", "blacklist".
	Type TrafficPolicyProfileNewParamsType `json:"type,omitzero" api:"required"`
	// Array of domain names.
	Domains []string `json:"domains,omitzero"`
	// Array of IP ranges in CIDR notation.
	IPRanges []string `json:"ip_ranges,omitzero"`
	// Bandwidth limit in kbps. Must be 512 or 1024.
	//
	// Any of 512, 1024.
	LimitBwKbps int64 `json:"limit_bw_kbps,omitzero"`
	// Array of PCEF service IDs to include in the profile.
	Services []string `json:"services,omitzero"`
	paramObj
}

func (r TrafficPolicyProfileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TrafficPolicyProfileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrafficPolicyProfileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the traffic policy profile.
type TrafficPolicyProfileNewParamsType string

const (
	TrafficPolicyProfileNewParamsTypeWhitelist TrafficPolicyProfileNewParamsType = "whitelist"
	TrafficPolicyProfileNewParamsTypeBlacklist TrafficPolicyProfileNewParamsType = "blacklist"
)

type TrafficPolicyProfileUpdateParams struct {
	// Bandwidth limit in kbps. Must be 512 or 1024, or null to remove.
	//
	// Any of 512, 1024.
	LimitBwKbps int64 `json:"limit_bw_kbps,omitzero"`
	// Array of domain names.
	Domains []string `json:"domains,omitzero"`
	// Array of IP ranges in CIDR notation.
	IPRanges []string `json:"ip_ranges,omitzero"`
	// Array of PCEF service IDs to include in the profile.
	Services []string `json:"services,omitzero"`
	// The type of the traffic policy profile.
	//
	// Any of "whitelist", "blacklist", "throttling".
	Type TrafficPolicyProfileUpdateParamsType `json:"type,omitzero"`
	paramObj
}

func (r TrafficPolicyProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TrafficPolicyProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrafficPolicyProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the traffic policy profile.
type TrafficPolicyProfileUpdateParamsType string

const (
	TrafficPolicyProfileUpdateParamsTypeWhitelist  TrafficPolicyProfileUpdateParamsType = "whitelist"
	TrafficPolicyProfileUpdateParamsTypeBlacklist  TrafficPolicyProfileUpdateParamsType = "blacklist"
	TrafficPolicyProfileUpdateParamsTypeThrottling TrafficPolicyProfileUpdateParamsType = "throttling"
)

type TrafficPolicyProfileListParams struct {
	// Filter by service ID.
	FilterService param.Opt[string] `query:"filter[service],omitzero" json:"-"`
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by traffic policy profile type.
	//
	// Any of "whitelist", "blacklist", "throttling".
	FilterType TrafficPolicyProfileListParamsFilterType `query:"filter[type],omitzero" json:"-"`
	// Sorts traffic policy profiles by the given field. Defaults to ascending order
	// unless field is prefixed with a minus sign.
	//
	// Any of "service", "-service", "type", "-type".
	Sort TrafficPolicyProfileListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TrafficPolicyProfileListParams]'s query parameters as
// `url.Values`.
func (r TrafficPolicyProfileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by traffic policy profile type.
type TrafficPolicyProfileListParamsFilterType string

const (
	TrafficPolicyProfileListParamsFilterTypeWhitelist  TrafficPolicyProfileListParamsFilterType = "whitelist"
	TrafficPolicyProfileListParamsFilterTypeBlacklist  TrafficPolicyProfileListParamsFilterType = "blacklist"
	TrafficPolicyProfileListParamsFilterTypeThrottling TrafficPolicyProfileListParamsFilterType = "throttling"
)

// Sorts traffic policy profiles by the given field. Defaults to ascending order
// unless field is prefixed with a minus sign.
type TrafficPolicyProfileListParamsSort string

const (
	TrafficPolicyProfileListParamsSortService     TrafficPolicyProfileListParamsSort = "service"
	TrafficPolicyProfileListParamsSortDescService TrafficPolicyProfileListParamsSort = "-service"
	TrafficPolicyProfileListParamsSortType        TrafficPolicyProfileListParamsSort = "type"
	TrafficPolicyProfileListParamsSortDescType    TrafficPolicyProfileListParamsSort = "-type"
)

type TrafficPolicyProfileListServicesParams struct {
	// Filter services by group.
	FilterGroup param.Opt[string] `query:"filter[group],omitzero" json:"-"`
	// Filter services by name.
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TrafficPolicyProfileListServicesParams]'s query parameters
// as `url.Values`.
func (r TrafficPolicyProfileListServicesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

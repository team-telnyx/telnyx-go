// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// PortingLoaConfigurationService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingLoaConfigurationService] method instead.
type PortingLoaConfigurationService struct {
	Options []option.RequestOption
}

// NewPortingLoaConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPortingLoaConfigurationService(opts ...option.RequestOption) (r PortingLoaConfigurationService) {
	r = PortingLoaConfigurationService{}
	r.Options = opts
	return
}

// Create a LOA configuration.
func (r *PortingLoaConfigurationService) New(ctx context.Context, body PortingLoaConfigurationNewParams, opts ...option.RequestOption) (res *PortingLoaConfigurationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "porting/loa_configurations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific LOA configuration.
func (r *PortingLoaConfigurationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PortingLoaConfigurationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting/loa_configurations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific LOA configuration.
func (r *PortingLoaConfigurationService) Update(ctx context.Context, id string, body PortingLoaConfigurationUpdateParams, opts ...option.RequestOption) (res *PortingLoaConfigurationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting/loa_configurations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List the LOA configurations.
func (r *PortingLoaConfigurationService) List(ctx context.Context, query PortingLoaConfigurationListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[PortingLoaConfiguration], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "porting/loa_configurations"
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

// List the LOA configurations.
func (r *PortingLoaConfigurationService) ListAutoPaging(ctx context.Context, query PortingLoaConfigurationListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[PortingLoaConfiguration] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a specific LOA configuration.
func (r *PortingLoaConfigurationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting/loa_configurations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Preview the LOA template that would be generated without need to create LOA
// configuration.
func (r *PortingLoaConfigurationService) Preview0(ctx context.Context, body PortingLoaConfigurationPreview0Params, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	path := "porting/loa_configuration/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Preview a specific LOA configuration.
func (r *PortingLoaConfigurationService) Preview1(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting/loa_configurations/%s/preview", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PortingLoaConfiguration struct {
	// Uniquely identifies the LOA configuration.
	ID string `json:"id" format:"uuid"`
	// The address of the company.
	Address PortingLoaConfigurationAddress `json:"address"`
	// The name of the company
	CompanyName string `json:"company_name"`
	// The contact information of the company.
	Contact PortingLoaConfigurationContact `json:"contact"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The logo to be used in the LOA.
	Logo PortingLoaConfigurationLogo `json:"logo"`
	// The name of the LOA configuration
	Name string `json:"name"`
	// The organization that owns the LOA configuration
	OrganizationID string `json:"organization_id"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Address        respjson.Field
		CompanyName    respjson.Field
		Contact        respjson.Field
		CreatedAt      respjson.Field
		Logo           respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		RecordType     respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingLoaConfiguration) RawJSON() string { return r.JSON.raw }
func (r *PortingLoaConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The address of the company.
type PortingLoaConfigurationAddress struct {
	// The locality of the company
	City string `json:"city"`
	// The country code of the company
	CountryCode string `json:"country_code"`
	// The extended address of the company
	ExtendedAddress string `json:"extended_address"`
	// The administrative area of the company
	State string `json:"state"`
	// The street address of the company
	StreetAddress string `json:"street_address"`
	// The postal code of the company
	ZipCode string `json:"zip_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City            respjson.Field
		CountryCode     respjson.Field
		ExtendedAddress respjson.Field
		State           respjson.Field
		StreetAddress   respjson.Field
		ZipCode         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingLoaConfigurationAddress) RawJSON() string { return r.JSON.raw }
func (r *PortingLoaConfigurationAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The contact information of the company.
type PortingLoaConfigurationContact struct {
	// The email address of the contact
	Email string `json:"email"`
	// The phone number of the contact
	PhoneNumber string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingLoaConfigurationContact) RawJSON() string { return r.JSON.raw }
func (r *PortingLoaConfigurationContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The logo to be used in the LOA.
type PortingLoaConfigurationLogo struct {
	// The content type of the logo.
	//
	// Any of "image/png".
	ContentType string `json:"content_type"`
	// Identifies the document that contains the logo.
	DocumentID string `json:"document_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		DocumentID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingLoaConfigurationLogo) RawJSON() string { return r.JSON.raw }
func (r *PortingLoaConfigurationLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingLoaConfigurationNewResponse struct {
	Data PortingLoaConfiguration `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingLoaConfigurationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingLoaConfigurationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingLoaConfigurationGetResponse struct {
	Data PortingLoaConfiguration `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingLoaConfigurationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingLoaConfigurationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingLoaConfigurationUpdateResponse struct {
	Data PortingLoaConfiguration `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingLoaConfigurationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingLoaConfigurationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingLoaConfigurationNewParams struct {
	// The address of the company.
	Address PortingLoaConfigurationNewParamsAddress `json:"address,omitzero,required"`
	// The name of the company
	CompanyName string `json:"company_name,required"`
	// The contact information of the company.
	Contact PortingLoaConfigurationNewParamsContact `json:"contact,omitzero,required"`
	// The logo of the LOA configuration
	Logo PortingLoaConfigurationNewParamsLogo `json:"logo,omitzero,required"`
	// The name of the LOA configuration
	Name string `json:"name,required"`
	paramObj
}

func (r PortingLoaConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingLoaConfigurationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingLoaConfigurationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The address of the company.
//
// The properties City, CountryCode, State, StreetAddress, ZipCode are required.
type PortingLoaConfigurationNewParamsAddress struct {
	// The locality of the company
	City string `json:"city,required"`
	// The country code of the company
	CountryCode string `json:"country_code,required"`
	// The administrative area of the company
	State string `json:"state,required"`
	// The street address of the company
	StreetAddress string `json:"street_address,required"`
	// The postal code of the company
	ZipCode string `json:"zip_code,required"`
	// The extended address of the company
	ExtendedAddress param.Opt[string] `json:"extended_address,omitzero"`
	paramObj
}

func (r PortingLoaConfigurationNewParamsAddress) MarshalJSON() (data []byte, err error) {
	type shadow PortingLoaConfigurationNewParamsAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingLoaConfigurationNewParamsAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The contact information of the company.
//
// The properties Email, PhoneNumber are required.
type PortingLoaConfigurationNewParamsContact struct {
	// The email address of the contact
	Email string `json:"email,required"`
	// The phone number of the contact
	PhoneNumber string `json:"phone_number,required"`
	paramObj
}

func (r PortingLoaConfigurationNewParamsContact) MarshalJSON() (data []byte, err error) {
	type shadow PortingLoaConfigurationNewParamsContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingLoaConfigurationNewParamsContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The logo of the LOA configuration
//
// The property DocumentID is required.
type PortingLoaConfigurationNewParamsLogo struct {
	// The document identification
	DocumentID string `json:"document_id,required" format:"uuid"`
	paramObj
}

func (r PortingLoaConfigurationNewParamsLogo) MarshalJSON() (data []byte, err error) {
	type shadow PortingLoaConfigurationNewParamsLogo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingLoaConfigurationNewParamsLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingLoaConfigurationUpdateParams struct {
	// The address of the company.
	Address PortingLoaConfigurationUpdateParamsAddress `json:"address,omitzero,required"`
	// The name of the company
	CompanyName string `json:"company_name,required"`
	// The contact information of the company.
	Contact PortingLoaConfigurationUpdateParamsContact `json:"contact,omitzero,required"`
	// The logo of the LOA configuration
	Logo PortingLoaConfigurationUpdateParamsLogo `json:"logo,omitzero,required"`
	// The name of the LOA configuration
	Name string `json:"name,required"`
	paramObj
}

func (r PortingLoaConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingLoaConfigurationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingLoaConfigurationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The address of the company.
//
// The properties City, CountryCode, State, StreetAddress, ZipCode are required.
type PortingLoaConfigurationUpdateParamsAddress struct {
	// The locality of the company
	City string `json:"city,required"`
	// The country code of the company
	CountryCode string `json:"country_code,required"`
	// The administrative area of the company
	State string `json:"state,required"`
	// The street address of the company
	StreetAddress string `json:"street_address,required"`
	// The postal code of the company
	ZipCode string `json:"zip_code,required"`
	// The extended address of the company
	ExtendedAddress param.Opt[string] `json:"extended_address,omitzero"`
	paramObj
}

func (r PortingLoaConfigurationUpdateParamsAddress) MarshalJSON() (data []byte, err error) {
	type shadow PortingLoaConfigurationUpdateParamsAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingLoaConfigurationUpdateParamsAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The contact information of the company.
//
// The properties Email, PhoneNumber are required.
type PortingLoaConfigurationUpdateParamsContact struct {
	// The email address of the contact
	Email string `json:"email,required"`
	// The phone number of the contact
	PhoneNumber string `json:"phone_number,required"`
	paramObj
}

func (r PortingLoaConfigurationUpdateParamsContact) MarshalJSON() (data []byte, err error) {
	type shadow PortingLoaConfigurationUpdateParamsContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingLoaConfigurationUpdateParamsContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The logo of the LOA configuration
//
// The property DocumentID is required.
type PortingLoaConfigurationUpdateParamsLogo struct {
	// The document identification
	DocumentID string `json:"document_id,required" format:"uuid"`
	paramObj
}

func (r PortingLoaConfigurationUpdateParamsLogo) MarshalJSON() (data []byte, err error) {
	type shadow PortingLoaConfigurationUpdateParamsLogo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingLoaConfigurationUpdateParamsLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingLoaConfigurationListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PortingLoaConfigurationListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingLoaConfigurationListParams]'s query parameters as
// `url.Values`.
func (r PortingLoaConfigurationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PortingLoaConfigurationListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingLoaConfigurationListParamsPage]'s query parameters
// as `url.Values`.
func (r PortingLoaConfigurationListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingLoaConfigurationPreview0Params struct {
	// The address of the company.
	Address PortingLoaConfigurationPreview0ParamsAddress `json:"address,omitzero,required"`
	// The name of the company
	CompanyName string `json:"company_name,required"`
	// The contact information of the company.
	Contact PortingLoaConfigurationPreview0ParamsContact `json:"contact,omitzero,required"`
	// The logo of the LOA configuration
	Logo PortingLoaConfigurationPreview0ParamsLogo `json:"logo,omitzero,required"`
	// The name of the LOA configuration
	Name string `json:"name,required"`
	paramObj
}

func (r PortingLoaConfigurationPreview0Params) MarshalJSON() (data []byte, err error) {
	type shadow PortingLoaConfigurationPreview0Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingLoaConfigurationPreview0Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The address of the company.
//
// The properties City, CountryCode, State, StreetAddress, ZipCode are required.
type PortingLoaConfigurationPreview0ParamsAddress struct {
	// The locality of the company
	City string `json:"city,required"`
	// The country code of the company
	CountryCode string `json:"country_code,required"`
	// The administrative area of the company
	State string `json:"state,required"`
	// The street address of the company
	StreetAddress string `json:"street_address,required"`
	// The postal code of the company
	ZipCode string `json:"zip_code,required"`
	// The extended address of the company
	ExtendedAddress param.Opt[string] `json:"extended_address,omitzero"`
	paramObj
}

func (r PortingLoaConfigurationPreview0ParamsAddress) MarshalJSON() (data []byte, err error) {
	type shadow PortingLoaConfigurationPreview0ParamsAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingLoaConfigurationPreview0ParamsAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The contact information of the company.
//
// The properties Email, PhoneNumber are required.
type PortingLoaConfigurationPreview0ParamsContact struct {
	// The email address of the contact
	Email string `json:"email,required"`
	// The phone number of the contact
	PhoneNumber string `json:"phone_number,required"`
	paramObj
}

func (r PortingLoaConfigurationPreview0ParamsContact) MarshalJSON() (data []byte, err error) {
	type shadow PortingLoaConfigurationPreview0ParamsContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingLoaConfigurationPreview0ParamsContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The logo of the LOA configuration
//
// The property DocumentID is required.
type PortingLoaConfigurationPreview0ParamsLogo struct {
	// The document identification
	DocumentID string `json:"document_id,required" format:"uuid"`
	paramObj
}

func (r PortingLoaConfigurationPreview0ParamsLogo) MarshalJSON() (data []byte, err error) {
	type shadow PortingLoaConfigurationPreview0ParamsLogo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingLoaConfigurationPreview0ParamsLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// CustomerServiceRecordService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerServiceRecordService] method instead.
type CustomerServiceRecordService struct {
	Options []option.RequestOption
}

// NewCustomerServiceRecordService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomerServiceRecordService(opts ...option.RequestOption) (r CustomerServiceRecordService) {
	r = CustomerServiceRecordService{}
	r.Options = opts
	return
}

// Create a new customer service record for the provided phone number.
func (r *CustomerServiceRecordService) New(ctx context.Context, body CustomerServiceRecordNewParams, opts ...option.RequestOption) (res *CustomerServiceRecordNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer_service_records"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific customer service record.
func (r *CustomerServiceRecordService) Get(ctx context.Context, customerServiceRecordID string, opts ...option.RequestOption) (res *CustomerServiceRecordGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if customerServiceRecordID == "" {
		err = errors.New("missing required customer_service_record_id parameter")
		return
	}
	path := fmt.Sprintf("customer_service_records/%s", customerServiceRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List customer service records.
func (r *CustomerServiceRecordService) List(ctx context.Context, query CustomerServiceRecordListParams, opts ...option.RequestOption) (res *CustomerServiceRecordListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer_service_records"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Verify the coverage for a list of phone numbers.
func (r *CustomerServiceRecordService) VerifyPhoneNumberCoverage(ctx context.Context, body CustomerServiceRecordVerifyPhoneNumberCoverageParams, opts ...option.RequestOption) (res *CustomerServiceRecordVerifyPhoneNumberCoverageResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer_service_records/phone_number_coverages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CustomerServiceRecord struct {
	// Uniquely identifies this customer service record
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The error message in case status is `failed`. This field would be null in case
	// of `pending` or `completed` status.
	ErrorMessage string `json:"error_message"`
	// The phone number of the customer service record.
	PhoneNumber string `json:"phone_number"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The result of the CSR request. This field would be null in case of `pending` or
	// `failed` status.
	Result CustomerServiceRecordResult `json:"result"`
	// The status of the customer service record
	//
	// Any of "pending", "completed", "failed".
	Status CustomerServiceRecordStatus `json:"status"`
	// ISO 8601 formatted date indicating when the resource was created.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		ErrorMessage respjson.Field
		PhoneNumber  respjson.Field
		RecordType   respjson.Field
		Result       respjson.Field
		Status       respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerServiceRecord) RawJSON() string { return r.JSON.raw }
func (r *CustomerServiceRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result of the CSR request. This field would be null in case of `pending` or
// `failed` status.
type CustomerServiceRecordResult struct {
	// The address of the customer service record
	Address CustomerServiceRecordResultAddress `json:"address"`
	// The admin of the customer service record.
	Admin CustomerServiceRecordResultAdmin `json:"admin"`
	// The associated phone numbers of the customer service record.
	AssociatedPhoneNumbers []string `json:"associated_phone_numbers"`
	// The name of the carrier that the customer service record is for.
	CarrierName string `json:"carrier_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address                respjson.Field
		Admin                  respjson.Field
		AssociatedPhoneNumbers respjson.Field
		CarrierName            respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerServiceRecordResult) RawJSON() string { return r.JSON.raw }
func (r *CustomerServiceRecordResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The address of the customer service record
type CustomerServiceRecordResultAddress struct {
	// The state of the address
	AdministrativeArea string `json:"administrative_area"`
	// The full address
	FullAddress string `json:"full_address"`
	// The city of the address
	Locality string `json:"locality"`
	// The zip code of the address
	PostalCode string `json:"postal_code"`
	// The street address
	StreetAddress string `json:"street_address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea respjson.Field
		FullAddress        respjson.Field
		Locality           respjson.Field
		PostalCode         respjson.Field
		StreetAddress      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerServiceRecordResultAddress) RawJSON() string { return r.JSON.raw }
func (r *CustomerServiceRecordResultAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The admin of the customer service record.
type CustomerServiceRecordResultAdmin struct {
	// The account number of the customer service record.
	AccountNumber string `json:"account_number"`
	// The authorized person name of the customer service record.
	AuthorizedPersonName string `json:"authorized_person_name"`
	// The billing phone number of the customer service record.
	BillingPhoneNumber string `json:"billing_phone_number"`
	// The name of the customer service record.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountNumber        respjson.Field
		AuthorizedPersonName respjson.Field
		BillingPhoneNumber   respjson.Field
		Name                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerServiceRecordResultAdmin) RawJSON() string { return r.JSON.raw }
func (r *CustomerServiceRecordResultAdmin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the customer service record
type CustomerServiceRecordStatus string

const (
	CustomerServiceRecordStatusPending   CustomerServiceRecordStatus = "pending"
	CustomerServiceRecordStatusCompleted CustomerServiceRecordStatus = "completed"
	CustomerServiceRecordStatusFailed    CustomerServiceRecordStatus = "failed"
)

type CustomerServiceRecordNewResponse struct {
	Data CustomerServiceRecord `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerServiceRecordNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerServiceRecordNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerServiceRecordGetResponse struct {
	Data CustomerServiceRecord `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerServiceRecordGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerServiceRecordGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerServiceRecordListResponse struct {
	Data []CustomerServiceRecord `json:"data"`
	Meta PaginationMeta          `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerServiceRecordListResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerServiceRecordListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerServiceRecordVerifyPhoneNumberCoverageResponse struct {
	Data []CustomerServiceRecordVerifyPhoneNumberCoverageResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerServiceRecordVerifyPhoneNumberCoverageResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerServiceRecordVerifyPhoneNumberCoverageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerServiceRecordVerifyPhoneNumberCoverageResponseData struct {
	// Additional data required to perform CSR for the phone number. Only returned if
	// `has_csr_coverage` is true.
	//
	// Any of "name", "authorized_person_name", "account_number", "customer_code",
	// "pin", "address_line_1", "city", "state", "zip_code", "billing_phone_number".
	AdditionalDataRequired []string `json:"additional_data_required"`
	// Indicates whether the phone number is covered or not.
	HasCsrCoverage bool `json:"has_csr_coverage"`
	// The phone number that is being verified.
	PhoneNumber string `json:"phone_number"`
	// The reason why the phone number is not covered. Only returned if
	// `has_csr_coverage` is false.
	Reason string `json:"reason"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalDataRequired respjson.Field
		HasCsrCoverage         respjson.Field
		PhoneNumber            respjson.Field
		Reason                 respjson.Field
		RecordType             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerServiceRecordVerifyPhoneNumberCoverageResponseData) RawJSON() string {
	return r.JSON.raw
}
func (r *CustomerServiceRecordVerifyPhoneNumberCoverageResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerServiceRecordNewParams struct {
	// A valid US phone number in E164 format.
	PhoneNumber string `json:"phone_number,required"`
	// Callback URL to receive webhook notifications.
	WebhookURL     param.Opt[string]                            `json:"webhook_url,omitzero"`
	AdditionalData CustomerServiceRecordNewParamsAdditionalData `json:"additional_data,omitzero"`
	paramObj
}

func (r CustomerServiceRecordNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerServiceRecordNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerServiceRecordNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerServiceRecordNewParamsAdditionalData struct {
	// The account number of the customer service record.
	AccountNumber param.Opt[string] `json:"account_number,omitzero"`
	// The first line of the address of the customer service record.
	AddressLine1 param.Opt[string] `json:"address_line_1,omitzero"`
	// The name of the authorized person.
	AuthorizedPersonName param.Opt[string] `json:"authorized_person_name,omitzero"`
	// The billing phone number of the customer service record.
	BillingPhoneNumber param.Opt[string] `json:"billing_phone_number,omitzero"`
	// The city of the customer service record.
	City param.Opt[string] `json:"city,omitzero"`
	// The customer code of the customer service record.
	CustomerCode param.Opt[string] `json:"customer_code,omitzero"`
	// The name of the administrator of CSR.
	Name param.Opt[string] `json:"name,omitzero"`
	// The PIN of the customer service record.
	Pin param.Opt[string] `json:"pin,omitzero"`
	// The state of the customer service record.
	State param.Opt[string] `json:"state,omitzero"`
	// The zip code of the customer service record.
	ZipCode param.Opt[string] `json:"zip_code,omitzero"`
	paramObj
}

func (r CustomerServiceRecordNewParamsAdditionalData) MarshalJSON() (data []byte, err error) {
	type shadow CustomerServiceRecordNewParamsAdditionalData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerServiceRecordNewParamsAdditionalData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerServiceRecordListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[phone_number][eq], filter[phone_number][in][], filter[status][eq],
	// filter[status][in][], filter[created_at][lt], filter[created_at][gt]
	Filter CustomerServiceRecordListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page CustomerServiceRecordListParamsPage `query:"page,omitzero" json:"-"`
	// Consolidated sort parameter (deepObject style). Originally: sort[value]
	Sort CustomerServiceRecordListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomerServiceRecordListParams]'s query parameters as
// `url.Values`.
func (r CustomerServiceRecordListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[phone_number][eq], filter[phone_number][in][], filter[status][eq],
// filter[status][in][], filter[created_at][lt], filter[created_at][gt]
type CustomerServiceRecordListParamsFilter struct {
	CreatedAt   CustomerServiceRecordListParamsFilterCreatedAt   `query:"created_at,omitzero" json:"-"`
	PhoneNumber CustomerServiceRecordListParamsFilterPhoneNumber `query:"phone_number,omitzero" json:"-"`
	Status      CustomerServiceRecordListParamsFilterStatus      `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomerServiceRecordListParamsFilter]'s query parameters
// as `url.Values`.
func (r CustomerServiceRecordListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerServiceRecordListParamsFilterCreatedAt struct {
	// Filters records to those created after a specific date.
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// Filters records to those created before a specific date.
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [CustomerServiceRecordListParamsFilterCreatedAt]'s query
// parameters as `url.Values`.
func (r CustomerServiceRecordListParamsFilterCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerServiceRecordListParamsFilterPhoneNumber struct {
	// Filters records to those with a specified number.
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	// Filters records to those with at least one number in the list.
	In []string `query:"in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomerServiceRecordListParamsFilterPhoneNumber]'s query
// parameters as `url.Values`.
func (r CustomerServiceRecordListParamsFilterPhoneNumber) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerServiceRecordListParamsFilterStatus struct {
	// Filters records to those with a specific status.
	//
	// Any of "pending", "completed", "failed".
	Eq string `query:"eq,omitzero" json:"-"`
	// Filters records to those with a least one status in the list.
	//
	// Any of "pending", "completed", "failed".
	In []string `query:"in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomerServiceRecordListParamsFilterStatus]'s query
// parameters as `url.Values`.
func (r CustomerServiceRecordListParamsFilterStatus) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[CustomerServiceRecordListParamsFilterStatus](
		"eq", "pending", "completed", "failed",
	)
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type CustomerServiceRecordListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomerServiceRecordListParamsPage]'s query parameters as
// `url.Values`.
func (r CustomerServiceRecordListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated sort parameter (deepObject style). Originally: sort[value]
type CustomerServiceRecordListParamsSort struct {
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order.
	//
	// Any of "created_at", "-created_at".
	Value string `query:"value,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomerServiceRecordListParamsSort]'s query parameters as
// `url.Values`.
func (r CustomerServiceRecordListParamsSort) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[CustomerServiceRecordListParamsSort](
		"value", "created_at", "-created_at",
	)
}

type CustomerServiceRecordVerifyPhoneNumberCoverageParams struct {
	// The phone numbers list to be verified.
	PhoneNumbers []string `json:"phone_numbers,omitzero,required"`
	paramObj
}

func (r CustomerServiceRecordVerifyPhoneNumberCoverageParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerServiceRecordVerifyPhoneNumberCoverageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerServiceRecordVerifyPhoneNumberCoverageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"encoding/json"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// LegacyReportingBatchDetailRecordService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLegacyReportingBatchDetailRecordService] method instead.
type LegacyReportingBatchDetailRecordService struct {
	Options      []option.RequestOption
	Messaging    LegacyReportingBatchDetailRecordMessagingService
	SpeechToText LegacyReportingBatchDetailRecordSpeechToTextService
	Voice        LegacyReportingBatchDetailRecordVoiceService
}

// NewLegacyReportingBatchDetailRecordService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLegacyReportingBatchDetailRecordService(opts ...option.RequestOption) (r LegacyReportingBatchDetailRecordService) {
	r = LegacyReportingBatchDetailRecordService{}
	r.Options = opts
	r.Messaging = NewLegacyReportingBatchDetailRecordMessagingService(opts...)
	r.SpeechToText = NewLegacyReportingBatchDetailRecordSpeechToTextService(opts...)
	r.Voice = NewLegacyReportingBatchDetailRecordVoiceService(opts...)
	return
}

// Query filter criteria. Note: The first filter object must specify filter_type as
// 'and'. You cannot follow an 'or' with another 'and'.
type Filter struct {
	// Billing group UUID to filter by
	BillingGroup string `json:"billing_group"`
	// Called line identification (destination number)
	Cld string `json:"cld"`
	// Filter type for CLD matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CldFilter FilterCldFilter `json:"cld_filter"`
	// Calling line identification (caller ID)
	Cli string `json:"cli"`
	// Filter type for CLI matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CliFilter FilterCliFilter `json:"cli_filter"`
	// Logical operator for combining filters
	//
	// Any of "and", "or".
	FilterType FilterFilterType `json:"filter_type"`
	// Tag name to filter by
	TagsList string `json:"tags_list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingGroup respjson.Field
		Cld          respjson.Field
		CldFilter    respjson.Field
		Cli          respjson.Field
		CliFilter    respjson.Field
		FilterType   respjson.Field
		TagsList     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Filter) RawJSON() string { return r.JSON.raw }
func (r *Filter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Filter to a FilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FilterParam.Overrides()
func (r Filter) ToParam() FilterParam {
	return param.Override[FilterParam](json.RawMessage(r.RawJSON()))
}

// Filter type for CLD matching
type FilterCldFilter string

const (
	FilterCldFilterContains   FilterCldFilter = "contains"
	FilterCldFilterStartsWith FilterCldFilter = "starts_with"
	FilterCldFilterEndsWith   FilterCldFilter = "ends_with"
)

// Filter type for CLI matching
type FilterCliFilter string

const (
	FilterCliFilterContains   FilterCliFilter = "contains"
	FilterCliFilterStartsWith FilterCliFilter = "starts_with"
	FilterCliFilterEndsWith   FilterCliFilter = "ends_with"
)

// Logical operator for combining filters
type FilterFilterType string

const (
	FilterFilterTypeAnd FilterFilterType = "and"
	FilterFilterTypeOr  FilterFilterType = "or"
)

// Query filter criteria. Note: The first filter object must specify filter_type as
// 'and'. You cannot follow an 'or' with another 'and'.
type FilterParam struct {
	// Billing group UUID to filter by
	BillingGroup param.Opt[string] `json:"billing_group,omitzero"`
	// Called line identification (destination number)
	Cld param.Opt[string] `json:"cld,omitzero"`
	// Calling line identification (caller ID)
	Cli param.Opt[string] `json:"cli,omitzero"`
	// Tag name to filter by
	TagsList param.Opt[string] `json:"tags_list,omitzero"`
	// Filter type for CLD matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CldFilter FilterCldFilter `json:"cld_filter,omitzero"`
	// Filter type for CLI matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CliFilter FilterCliFilter `json:"cli_filter,omitzero"`
	// Logical operator for combining filters
	//
	// Any of "and", "or".
	FilterType FilterFilterType `json:"filter_type,omitzero"`
	paramObj
}

func (r FilterParam) MarshalJSON() (data []byte, err error) {
	type shadow FilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

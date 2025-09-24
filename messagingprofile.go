// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
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
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared"
)

// MessagingProfileService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingProfileService] method instead.
type MessagingProfileService struct {
	Options         []option.RequestOption
	AutorespConfigs MessagingProfileAutorespConfigService
}

// NewMessagingProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingProfileService(opts ...option.RequestOption) (r MessagingProfileService) {
	r = MessagingProfileService{}
	r.Options = opts
	r.AutorespConfigs = NewMessagingProfileAutorespConfigService(opts...)
	return
}

// Create a messaging profile
func (r *MessagingProfileService) New(ctx context.Context, body MessagingProfileNewParams, opts ...option.RequestOption) (res *MessagingProfileNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messaging_profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a messaging profile
func (r *MessagingProfileService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingProfileGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_profiles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a messaging profile
func (r *MessagingProfileService) Update(ctx context.Context, id string, body MessagingProfileUpdateParams, opts ...option.RequestOption) (res *MessagingProfileUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_profiles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List messaging profiles
func (r *MessagingProfileService) List(ctx context.Context, query MessagingProfileListParams, opts ...option.RequestOption) (res *MessagingProfileListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messaging_profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a messaging profile
func (r *MessagingProfileService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingProfileDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_profiles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List phone numbers associated with a messaging profile
func (r *MessagingProfileService) ListPhoneNumbers(ctx context.Context, id string, query MessagingProfileListPhoneNumbersParams, opts ...option.RequestOption) (res *MessagingProfileListPhoneNumbersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_profiles/%s/phone_numbers", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List short codes associated with a messaging profile
func (r *MessagingProfileService) ListShortCodes(ctx context.Context, id string, query MessagingProfileListShortCodesParams, opts ...option.RequestOption) (res *MessagingProfileListShortCodesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_profiles/%s/short_codes", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MessagingProfile struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The alphanumeric sender ID to use when sending to destinations that require an
	// alphanumeric sender ID.
	AlphaSender string `json:"alpha_sender,nullable"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The maximum amount of money (in USD) that can be spent by this profile before
	// midnight UTC.
	DailySpendLimit string `json:"daily_spend_limit"`
	// Whether to enforce the value configured by `daily_spend_limit`.
	DailySpendLimitEnabled bool `json:"daily_spend_limit_enabled"`
	// Specifies whether the messaging profile is enabled or not.
	Enabled bool `json:"enabled"`
	// enables SMS fallback for MMS messages.
	MmsFallBackToSMS bool `json:"mms_fall_back_to_sms"`
	// enables automated resizing of MMS media.
	MmsTranscoding bool `json:"mms_transcoding"`
	// A user friendly name for the messaging profile.
	Name string `json:"name"`
	// Number Pool allows you to send messages from a pool of numbers of different
	// types, assigning weights to each type. The pool consists of all the long code
	// and toll free numbers assigned to the messaging profile.
	//
	// To disable this feature, set the object field to `null`.
	NumberPoolSettings NumberPoolSettings `json:"number_pool_settings,nullable"`
	// Identifies the type of the resource.
	//
	// Any of "messaging_profile".
	RecordType MessagingProfileRecordType `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The URL shortener feature allows automatic replacement of URLs that were
	// generated using a public URL shortener service. Some examples include bit.do,
	// bit.ly, goo.gl, ht.ly, is.gd, ow.ly, rebrand.ly, t.co, tiny.cc, and tinyurl.com.
	// Such URLs are replaced with with links generated by Telnyx. The use of custom
	// links can improve branding and message deliverability.
	//
	// To disable this feature, set the object field to `null`.
	URLShortenerSettings URLShortenerSettings `json:"url_shortener_settings,nullable"`
	// Secret used to authenticate with v1 endpoints.
	V1Secret string `json:"v1_secret"`
	// Determines which webhook format will be used, Telnyx API v1, v2, or a legacy
	// 2010-04-01 format.
	//
	// Any of "1", "2", "2010-04-01".
	WebhookAPIVersion MessagingProfileWebhookAPIVersion `json:"webhook_api_version"`
	// The failover URL where webhooks related to this messaging profile will be sent
	// if sending to the primary URL fails.
	WebhookFailoverURL string `json:"webhook_failover_url,nullable" format:"url"`
	// The URL where webhooks related to this messaging profile will be sent.
	WebhookURL string `json:"webhook_url,nullable" format:"url"`
	// Destinations to which the messaging profile is allowed to send. The elements in
	// the list must be valid ISO 3166-1 alpha-2 country codes. If set to `["*"]`, all
	// destinations will be allowed.
	WhitelistedDestinations []string `json:"whitelisted_destinations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		AlphaSender             respjson.Field
		CreatedAt               respjson.Field
		DailySpendLimit         respjson.Field
		DailySpendLimitEnabled  respjson.Field
		Enabled                 respjson.Field
		MmsFallBackToSMS        respjson.Field
		MmsTranscoding          respjson.Field
		Name                    respjson.Field
		NumberPoolSettings      respjson.Field
		RecordType              respjson.Field
		UpdatedAt               respjson.Field
		URLShortenerSettings    respjson.Field
		V1Secret                respjson.Field
		WebhookAPIVersion       respjson.Field
		WebhookFailoverURL      respjson.Field
		WebhookURL              respjson.Field
		WhitelistedDestinations respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingProfile) RawJSON() string { return r.JSON.raw }
func (r *MessagingProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type MessagingProfileRecordType string

const (
	MessagingProfileRecordTypeMessagingProfile MessagingProfileRecordType = "messaging_profile"
)

// Determines which webhook format will be used, Telnyx API v1, v2, or a legacy
// 2010-04-01 format.
type MessagingProfileWebhookAPIVersion string

const (
	MessagingProfileWebhookAPIVersion1          MessagingProfileWebhookAPIVersion = "1"
	MessagingProfileWebhookAPIVersion2          MessagingProfileWebhookAPIVersion = "2"
	MessagingProfileWebhookAPIVersion2010_04_01 MessagingProfileWebhookAPIVersion = "2010-04-01"
)

// Number Pool allows you to send messages from a pool of numbers of different
// types, assigning weights to each type. The pool consists of all the long code
// and toll free numbers assigned to the messaging profile.
//
// To disable this feature, set the object field to `null`.
type NumberPoolSettings struct {
	// Defines the probability weight for a Long Code number to be selected when
	// sending a message. The higher the weight the higher the probability. The sum of
	// the weights for all number types does not necessarily need to add to 100. Weight
	// must be a non-negative number, and when equal to zero it will remove the number
	// type from the pool.
	LongCodeWeight float64 `json:"long_code_weight,required"`
	// If set to true all unhealthy numbers will be automatically excluded from the
	// pool. Health metrics per number are calculated on a regular basis, taking into
	// account the deliverability rate and the amount of messages marked as spam by
	// upstream carriers. Numbers with a deliverability rate below 25% or spam ratio
	// over 75% will be considered unhealthy.
	SkipUnhealthy bool `json:"skip_unhealthy,required"`
	// Defines the probability weight for a Toll Free number to be selected when
	// sending a message. The higher the weight the higher the probability. The sum of
	// the weights for all number types does not necessarily need to add to 100. Weight
	// must be a non-negative number, and when equal to zero it will remove the number
	// type from the pool.
	TollFreeWeight float64 `json:"toll_free_weight,required"`
	// If set to true, Number Pool will try to choose a sending number with the same
	// area code as the destination number. If there are no such numbers available, a
	// nunber with a different area code will be chosen. Currently only NANP numbers
	// are supported.
	Geomatch bool `json:"geomatch"`
	// If set to true, Number Pool will try to choose the same sending number for all
	// messages to a particular recipient. If the sending number becomes unhealthy and
	// `skip_unhealthy` is set to true, a new number will be chosen.
	StickySender bool `json:"sticky_sender"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LongCodeWeight respjson.Field
		SkipUnhealthy  respjson.Field
		TollFreeWeight respjson.Field
		Geomatch       respjson.Field
		StickySender   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberPoolSettings) RawJSON() string { return r.JSON.raw }
func (r *NumberPoolSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NumberPoolSettings to a NumberPoolSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NumberPoolSettingsParam.Overrides()
func (r NumberPoolSettings) ToParam() NumberPoolSettingsParam {
	return param.Override[NumberPoolSettingsParam](json.RawMessage(r.RawJSON()))
}

// Number Pool allows you to send messages from a pool of numbers of different
// types, assigning weights to each type. The pool consists of all the long code
// and toll free numbers assigned to the messaging profile.
//
// To disable this feature, set the object field to `null`.
//
// The properties LongCodeWeight, SkipUnhealthy, TollFreeWeight are required.
type NumberPoolSettingsParam struct {
	// Defines the probability weight for a Long Code number to be selected when
	// sending a message. The higher the weight the higher the probability. The sum of
	// the weights for all number types does not necessarily need to add to 100. Weight
	// must be a non-negative number, and when equal to zero it will remove the number
	// type from the pool.
	LongCodeWeight float64 `json:"long_code_weight,required"`
	// If set to true all unhealthy numbers will be automatically excluded from the
	// pool. Health metrics per number are calculated on a regular basis, taking into
	// account the deliverability rate and the amount of messages marked as spam by
	// upstream carriers. Numbers with a deliverability rate below 25% or spam ratio
	// over 75% will be considered unhealthy.
	SkipUnhealthy bool `json:"skip_unhealthy,required"`
	// Defines the probability weight for a Toll Free number to be selected when
	// sending a message. The higher the weight the higher the probability. The sum of
	// the weights for all number types does not necessarily need to add to 100. Weight
	// must be a non-negative number, and when equal to zero it will remove the number
	// type from the pool.
	TollFreeWeight float64 `json:"toll_free_weight,required"`
	// If set to true, Number Pool will try to choose a sending number with the same
	// area code as the destination number. If there are no such numbers available, a
	// nunber with a different area code will be chosen. Currently only NANP numbers
	// are supported.
	Geomatch param.Opt[bool] `json:"geomatch,omitzero"`
	// If set to true, Number Pool will try to choose the same sending number for all
	// messages to a particular recipient. If the sending number becomes unhealthy and
	// `skip_unhealthy` is set to true, a new number will be chosen.
	StickySender param.Opt[bool] `json:"sticky_sender,omitzero"`
	paramObj
}

func (r NumberPoolSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow NumberPoolSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberPoolSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The URL shortener feature allows automatic replacement of URLs that were
// generated using a public URL shortener service. Some examples include bit.do,
// bit.ly, goo.gl, ht.ly, is.gd, ow.ly, rebrand.ly, t.co, tiny.cc, and tinyurl.com.
// Such URLs are replaced with with links generated by Telnyx. The use of custom
// links can improve branding and message deliverability.
//
// To disable this feature, set the object field to `null`.
type URLShortenerSettings struct {
	// One of the domains provided by the Telnyx URL shortener service.
	Domain string `json:"domain,required"`
	// Optional prefix that can be used to identify your brand, and will appear in the
	// Telnyx generated URLs after the domain name.
	Prefix string `json:"prefix"`
	// Use the link replacement tool only for links that are specifically blacklisted
	// by Telnyx.
	ReplaceBlacklistOnly bool `json:"replace_blacklist_only"`
	// Receive webhooks for when your replaced links are clicked. Webhooks are sent to
	// the webhooks on the messaging profile.
	SendWebhooks bool `json:"send_webhooks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain               respjson.Field
		Prefix               respjson.Field
		ReplaceBlacklistOnly respjson.Field
		SendWebhooks         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URLShortenerSettings) RawJSON() string { return r.JSON.raw }
func (r *URLShortenerSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this URLShortenerSettings to a URLShortenerSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// URLShortenerSettingsParam.Overrides()
func (r URLShortenerSettings) ToParam() URLShortenerSettingsParam {
	return param.Override[URLShortenerSettingsParam](json.RawMessage(r.RawJSON()))
}

// The URL shortener feature allows automatic replacement of URLs that were
// generated using a public URL shortener service. Some examples include bit.do,
// bit.ly, goo.gl, ht.ly, is.gd, ow.ly, rebrand.ly, t.co, tiny.cc, and tinyurl.com.
// Such URLs are replaced with with links generated by Telnyx. The use of custom
// links can improve branding and message deliverability.
//
// To disable this feature, set the object field to `null`.
//
// The property Domain is required.
type URLShortenerSettingsParam struct {
	// One of the domains provided by the Telnyx URL shortener service.
	Domain string `json:"domain,required"`
	// Optional prefix that can be used to identify your brand, and will appear in the
	// Telnyx generated URLs after the domain name.
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	// Use the link replacement tool only for links that are specifically blacklisted
	// by Telnyx.
	ReplaceBlacklistOnly param.Opt[bool] `json:"replace_blacklist_only,omitzero"`
	// Receive webhooks for when your replaced links are clicked. Webhooks are sent to
	// the webhooks on the messaging profile.
	SendWebhooks param.Opt[bool] `json:"send_webhooks,omitzero"`
	paramObj
}

func (r URLShortenerSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow URLShortenerSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLShortenerSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingProfileNewResponse struct {
	Data MessagingProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingProfileNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingProfileNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingProfileGetResponse struct {
	Data MessagingProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingProfileGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingProfileGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingProfileUpdateResponse struct {
	Data MessagingProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingProfileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingProfileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingProfileListResponse struct {
	Data []MessagingProfile `json:"data"`
	Meta PaginationMeta     `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingProfileListResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingProfileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingProfileDeleteResponse struct {
	Data MessagingProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingProfileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingProfileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingProfileListPhoneNumbersResponse struct {
	Data []shared.PhoneNumberWithMessagingSettings `json:"data"`
	Meta PaginationMeta                            `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingProfileListPhoneNumbersResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingProfileListPhoneNumbersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingProfileListShortCodesResponse struct {
	Data []shared.ShortCode `json:"data"`
	Meta PaginationMeta     `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingProfileListShortCodesResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingProfileListShortCodesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingProfileNewParams struct {
	// A user friendly name for the messaging profile.
	Name string `json:"name,required"`
	// Destinations to which the messaging profile is allowed to send. The elements in
	// the list must be valid ISO 3166-1 alpha-2 country codes. If set to `["*"]` all
	// destinations will be allowed.
	WhitelistedDestinations []string `json:"whitelisted_destinations,omitzero,required"`
	// The alphanumeric sender ID to use when sending to destinations that require an
	// alphanumeric sender ID.
	AlphaSender param.Opt[string] `json:"alpha_sender,omitzero"`
	// The failover URL where webhooks related to this messaging profile will be sent
	// if sending to the primary URL fails.
	WebhookFailoverURL param.Opt[string] `json:"webhook_failover_url,omitzero" format:"url"`
	// The URL where webhooks related to this messaging profile will be sent.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	// The maximum amount of money (in USD) that can be spent by this profile before
	// midnight UTC.
	DailySpendLimit param.Opt[string] `json:"daily_spend_limit,omitzero"`
	// Whether to enforce the value configured by `daily_spend_limit`.
	DailySpendLimitEnabled param.Opt[bool] `json:"daily_spend_limit_enabled,omitzero"`
	// Specifies whether the messaging profile is enabled or not.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// enables SMS fallback for MMS messages.
	MmsFallBackToSMS param.Opt[bool] `json:"mms_fall_back_to_sms,omitzero"`
	// enables automated resizing of MMS media.
	MmsTranscoding param.Opt[bool] `json:"mms_transcoding,omitzero"`
	// Number Pool allows you to send messages from a pool of numbers of different
	// types, assigning weights to each type. The pool consists of all the long code
	// and toll free numbers assigned to the messaging profile.
	//
	// To disable this feature, set the object field to `null`.
	NumberPoolSettings NumberPoolSettingsParam `json:"number_pool_settings,omitzero"`
	// The URL shortener feature allows automatic replacement of URLs that were
	// generated using a public URL shortener service. Some examples include bit.do,
	// bit.ly, goo.gl, ht.ly, is.gd, ow.ly, rebrand.ly, t.co, tiny.cc, and tinyurl.com.
	// Such URLs are replaced with with links generated by Telnyx. The use of custom
	// links can improve branding and message deliverability.
	//
	// To disable this feature, set the object field to `null`.
	URLShortenerSettings URLShortenerSettingsParam `json:"url_shortener_settings,omitzero"`
	// Determines which webhook format will be used, Telnyx API v1, v2, or a legacy
	// 2010-04-01 format.
	//
	// Any of "1", "2", "2010-04-01".
	WebhookAPIVersion MessagingProfileNewParamsWebhookAPIVersion `json:"webhook_api_version,omitzero"`
	paramObj
}

func (r MessagingProfileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MessagingProfileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagingProfileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determines which webhook format will be used, Telnyx API v1, v2, or a legacy
// 2010-04-01 format.
type MessagingProfileNewParamsWebhookAPIVersion string

const (
	MessagingProfileNewParamsWebhookAPIVersion1          MessagingProfileNewParamsWebhookAPIVersion = "1"
	MessagingProfileNewParamsWebhookAPIVersion2          MessagingProfileNewParamsWebhookAPIVersion = "2"
	MessagingProfileNewParamsWebhookAPIVersion2010_04_01 MessagingProfileNewParamsWebhookAPIVersion = "2010-04-01"
)

type MessagingProfileUpdateParams struct {
	// The alphanumeric sender ID to use when sending to destinations that require an
	// alphanumeric sender ID.
	AlphaSender param.Opt[string] `json:"alpha_sender,omitzero"`
	// The failover URL where webhooks related to this messaging profile will be sent
	// if sending to the primary URL fails.
	WebhookFailoverURL param.Opt[string] `json:"webhook_failover_url,omitzero" format:"url"`
	// The URL where webhooks related to this messaging profile will be sent.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	// The maximum amount of money (in USD) that can be spent by this profile before
	// midnight UTC.
	DailySpendLimit param.Opt[string] `json:"daily_spend_limit,omitzero"`
	// Whether to enforce the value configured by `daily_spend_limit`.
	DailySpendLimitEnabled param.Opt[bool] `json:"daily_spend_limit_enabled,omitzero"`
	// Specifies whether the messaging profile is enabled or not.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// enables SMS fallback for MMS messages.
	MmsFallBackToSMS param.Opt[bool] `json:"mms_fall_back_to_sms,omitzero"`
	// enables automated resizing of MMS media.
	MmsTranscoding param.Opt[bool] `json:"mms_transcoding,omitzero"`
	// A user friendly name for the messaging profile.
	Name param.Opt[string] `json:"name,omitzero"`
	// Secret used to authenticate with v1 endpoints.
	V1Secret param.Opt[string] `json:"v1_secret,omitzero"`
	// Number Pool allows you to send messages from a pool of numbers of different
	// types, assigning weights to each type. The pool consists of all the long code
	// and toll free numbers assigned to the messaging profile.
	//
	// To disable this feature, set the object field to `null`.
	NumberPoolSettings NumberPoolSettingsParam `json:"number_pool_settings,omitzero"`
	// The URL shortener feature allows automatic replacement of URLs that were
	// generated using a public URL shortener service. Some examples include bit.do,
	// bit.ly, goo.gl, ht.ly, is.gd, ow.ly, rebrand.ly, t.co, tiny.cc, and tinyurl.com.
	// Such URLs are replaced with with links generated by Telnyx. The use of custom
	// links can improve branding and message deliverability.
	//
	// To disable this feature, set the object field to `null`.
	URLShortenerSettings URLShortenerSettingsParam `json:"url_shortener_settings,omitzero"`
	// Determines which webhook format will be used, Telnyx API v1, v2, or a legacy
	// 2010-04-01 format.
	//
	// Any of "1", "2", "2010-04-01".
	WebhookAPIVersion MessagingProfileUpdateParamsWebhookAPIVersion `json:"webhook_api_version,omitzero"`
	// Destinations to which the messaging profile is allowed to send. The elements in
	// the list must be valid ISO 3166-1 alpha-2 country codes. If set to `["*"]`, all
	// destinations will be allowed.
	//
	// This field is required if the messaging profile doesn't have it defined yet.
	WhitelistedDestinations []string `json:"whitelisted_destinations,omitzero"`
	paramObj
}

func (r MessagingProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MessagingProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagingProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determines which webhook format will be used, Telnyx API v1, v2, or a legacy
// 2010-04-01 format.
type MessagingProfileUpdateParamsWebhookAPIVersion string

const (
	MessagingProfileUpdateParamsWebhookAPIVersion1          MessagingProfileUpdateParamsWebhookAPIVersion = "1"
	MessagingProfileUpdateParamsWebhookAPIVersion2          MessagingProfileUpdateParamsWebhookAPIVersion = "2"
	MessagingProfileUpdateParamsWebhookAPIVersion2010_04_01 MessagingProfileUpdateParamsWebhookAPIVersion = "2010-04-01"
)

type MessagingProfileListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[name]
	Filter MessagingProfileListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page MessagingProfileListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingProfileListParams]'s query parameters as
// `url.Values`.
func (r MessagingProfileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[name]
type MessagingProfileListParamsFilter struct {
	// Filter by name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingProfileListParamsFilter]'s query parameters as
// `url.Values`.
func (r MessagingProfileListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type MessagingProfileListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingProfileListParamsPage]'s query parameters as
// `url.Values`.
func (r MessagingProfileListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingProfileListPhoneNumbersParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page MessagingProfileListPhoneNumbersParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingProfileListPhoneNumbersParams]'s query parameters
// as `url.Values`.
func (r MessagingProfileListPhoneNumbersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type MessagingProfileListPhoneNumbersParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingProfileListPhoneNumbersParamsPage]'s query
// parameters as `url.Values`.
func (r MessagingProfileListPhoneNumbersParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingProfileListShortCodesParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page MessagingProfileListShortCodesParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingProfileListShortCodesParams]'s query parameters as
// `url.Values`.
func (r MessagingProfileListShortCodesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type MessagingProfileListShortCodesParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingProfileListShortCodesParamsPage]'s query
// parameters as `url.Values`.
func (r MessagingProfileListShortCodesParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

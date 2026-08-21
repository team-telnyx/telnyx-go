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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage RCS agent registration, testing, verification, and launch.
//
// RcAgentService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRcAgentService] method instead.
type RcAgentService struct {
	Options []option.RequestOption
	// Manage RCS agent registration, testing, verification, and launch.
	TestDevices RcAgentTestDeviceService
}

// NewRcAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRcAgentService(opts ...option.RequestOption) (r RcAgentService) {
	r = RcAgentService{}
	r.Options = opts
	r.TestDevices = NewRcAgentTestDeviceService(opts...)
	return
}

// Creates an editable RCS agent draft under a brand. The `Idempotency-Key` is
// scoped to the authenticated organization. Reusing the key with the same request
// returns the original agent, while reusing it with a different request returns a
// conflict.
func (r *RcAgentService) New(ctx context.Context, params RcAgentNewParams, opts ...option.RequestOption) (res *AgentResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "rcs/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves an RCS agent, section statuses, test devices, carrier approvals, and
// provider capabilities.
func (r *RcAgentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rcs/agents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates one or more fields on an agent while its status is `CREATED`. Submitted
// agents cannot be changed through this endpoint.
func (r *RcAgentService) Update(ctx context.Context, id string, body RcAgentUpdateParams, opts ...option.RequestOption) (res *AgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rcs/agents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists RCS agents owned by the authenticated organization, optionally filtered by
// brand.
func (r *RcAgentService) List(ctx context.Context, query RcAgentListParams, opts ...option.RequestOption) (res *[]AgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rcs/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Adds the campaign and testing configuration, then starts asynchronous carrier
// launch. Agent basics must already be submitted. Repeating a launch that is
// already in progress returns the current agent without creating new work.
func (r *RcAgentService) Launch(ctx context.Context, id string, body RcAgentLaunchParams, opts ...option.RequestOption) (res *AgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rcs/agents/%s/launch", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists carrier approval records for an RCS agent. The provider may expose
// per-carrier, hub-level, or bot-level approval status.
func (r *RcAgentService) GetCarrierApprovals(ctx context.Context, id string, opts ...option.RequestOption) (res *[]CarrierApprovalResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rcs/agents/%s/carrier_approvals", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Starts asynchronous provider provisioning and submits the agent's basic
// configuration. The brand must be `VERIFIED`. Repeating this request for an
// in-progress agent returns its current state without creating new work.
func (r *RcAgentService) Submit(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rcs/agents/%s/submit", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type AgentCampaignConfiguration struct {
	CompanyOverview       string                    `json:"company_overview" api:"required"`
	AdditionalInformation string                    `json:"additional_information" api:"nullable"`
	AgentOverview         string                    `json:"agent_overview" api:"nullable"`
	ConsentSettings       AgentConsentConfiguration `json:"consent_settings" api:"nullable"`
	Interactions          []AgentInteraction        `json:"interactions"`
	MessageExamples       []string                  `json:"message_examples"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyOverview       respjson.Field
		AdditionalInformation respjson.Field
		AgentOverview         respjson.Field
		ConsentSettings       respjson.Field
		Interactions          respjson.Field
		MessageExamples       respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentCampaignConfiguration) RawJSON() string { return r.JSON.raw }
func (r *AgentCampaignConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentCampaignConfiguration to a
// AgentCampaignConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentCampaignConfigurationParam.Overrides()
func (r AgentCampaignConfiguration) ToParam() AgentCampaignConfigurationParam {
	return param.Override[AgentCampaignConfigurationParam](json.RawMessage(r.RawJSON()))
}

// The property CompanyOverview is required.
type AgentCampaignConfigurationParam struct {
	CompanyOverview       string                         `json:"company_overview" api:"required"`
	AdditionalInformation param.Opt[string]              `json:"additional_information,omitzero"`
	AgentOverview         param.Opt[string]              `json:"agent_overview,omitzero"`
	ConsentSettings       AgentConsentConfigurationParam `json:"consent_settings,omitzero"`
	Interactions          []AgentInteractionParam        `json:"interactions,omitzero"`
	MessageExamples       []string                       `json:"message_examples,omitzero"`
	paramObj
}

func (r AgentCampaignConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentCampaignConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentCampaignConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentConfiguration struct {
	// Basic agent identity and contact information. At least one complete phone,
	// website, or email contact is required.
	Basics   AgentConfigurationBasicsUnion `json:"basics" api:"required"`
	Campaign AgentCampaignConfiguration    `json:"campaign" api:"nullable"`
	Testing  AgentTestingConfiguration     `json:"testing" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Basics      respjson.Field
		Campaign    respjson.Field
		Testing     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentConfiguration) RawJSON() string { return r.JSON.raw }
func (r *AgentConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentConfiguration to a AgentConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentConfigurationParam.Overrides()
func (r AgentConfiguration) ToParam() AgentConfigurationParam {
	return param.Override[AgentConfigurationParam](json.RawMessage(r.RawJSON()))
}

// AgentConfigurationBasicsUnion contains all possible properties and values from
// [AgentConfigurationBasicsAgentPhoneContactRequirement],
// [AgentConfigurationBasicsAgentWebhookContactRequirement],
// [AgentConfigurationBasicsAgentProfileContactRequirement].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AgentConfigurationBasicsUnion struct {
	// This field is from variant
	// [AgentConfigurationBasicsAgentPhoneContactRequirement].
	PhoneNumber AgentPhoneContact `json:"phone_number"`
	BrandColor  string            `json:"brand_color"`
	Description string            `json:"description"`
	// This field is from variant
	// [AgentConfigurationBasicsAgentPhoneContactRequirement].
	Email                 AgentEmailContact `json:"email"`
	HeroURL               string            `json:"hero_url"`
	LogoURL               string            `json:"logo_url"`
	PrivacyPolicyURL      string            `json:"privacy_policy_url"`
	TermsAndConditionsURL string            `json:"terms_and_conditions_url"`
	// This field is from variant
	// [AgentConfigurationBasicsAgentPhoneContactRequirement].
	Website AgentWebsiteContact `json:"website"`
	JSON    struct {
		PhoneNumber           respjson.Field
		BrandColor            respjson.Field
		Description           respjson.Field
		Email                 respjson.Field
		HeroURL               respjson.Field
		LogoURL               respjson.Field
		PrivacyPolicyURL      respjson.Field
		TermsAndConditionsURL respjson.Field
		Website               respjson.Field
		raw                   string
	} `json:"-"`
}

func (u AgentConfigurationBasicsUnion) AsAgentPhoneContactRequirement() (v AgentConfigurationBasicsAgentPhoneContactRequirement) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentConfigurationBasicsUnion) AsAgentWebhookContactRequirement() (v AgentConfigurationBasicsAgentWebhookContactRequirement) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentConfigurationBasicsUnion) AsAgentProfileContactRequirement() (v AgentConfigurationBasicsAgentProfileContactRequirement) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentConfigurationBasicsUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentConfigurationBasicsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentConfigurationBasicsAgentPhoneContactRequirement struct {
	PhoneNumber           AgentPhoneContact   `json:"phone_number" api:"required"`
	BrandColor            string              `json:"brand_color"`
	Description           string              `json:"description"`
	Email                 AgentEmailContact   `json:"email" api:"nullable"`
	HeroURL               string              `json:"hero_url" format:"uri"`
	LogoURL               string              `json:"logo_url" format:"uri"`
	PrivacyPolicyURL      string              `json:"privacy_policy_url" format:"uri"`
	TermsAndConditionsURL string              `json:"terms_and_conditions_url" format:"uri"`
	Website               AgentWebsiteContact `json:"website" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber           respjson.Field
		BrandColor            respjson.Field
		Description           respjson.Field
		Email                 respjson.Field
		HeroURL               respjson.Field
		LogoURL               respjson.Field
		PrivacyPolicyURL      respjson.Field
		TermsAndConditionsURL respjson.Field
		Website               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentConfigurationBasicsAgentPhoneContactRequirement) RawJSON() string { return r.JSON.raw }
func (r *AgentConfigurationBasicsAgentPhoneContactRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentConfigurationBasicsAgentWebhookContactRequirement struct {
	Website               AgentWebsiteContact `json:"website" api:"required"`
	BrandColor            string              `json:"brand_color"`
	Description           string              `json:"description"`
	Email                 AgentEmailContact   `json:"email" api:"nullable"`
	HeroURL               string              `json:"hero_url" format:"uri"`
	LogoURL               string              `json:"logo_url" format:"uri"`
	PhoneNumber           AgentPhoneContact   `json:"phone_number" api:"nullable"`
	PrivacyPolicyURL      string              `json:"privacy_policy_url" format:"uri"`
	TermsAndConditionsURL string              `json:"terms_and_conditions_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Website               respjson.Field
		BrandColor            respjson.Field
		Description           respjson.Field
		Email                 respjson.Field
		HeroURL               respjson.Field
		LogoURL               respjson.Field
		PhoneNumber           respjson.Field
		PrivacyPolicyURL      respjson.Field
		TermsAndConditionsURL respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentConfigurationBasicsAgentWebhookContactRequirement) RawJSON() string { return r.JSON.raw }
func (r *AgentConfigurationBasicsAgentWebhookContactRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentConfigurationBasicsAgentProfileContactRequirement struct {
	Email                 AgentEmailContact   `json:"email" api:"required"`
	BrandColor            string              `json:"brand_color"`
	Description           string              `json:"description"`
	HeroURL               string              `json:"hero_url" format:"uri"`
	LogoURL               string              `json:"logo_url" format:"uri"`
	PhoneNumber           AgentPhoneContact   `json:"phone_number" api:"nullable"`
	PrivacyPolicyURL      string              `json:"privacy_policy_url" format:"uri"`
	TermsAndConditionsURL string              `json:"terms_and_conditions_url" format:"uri"`
	Website               AgentWebsiteContact `json:"website" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email                 respjson.Field
		BrandColor            respjson.Field
		Description           respjson.Field
		HeroURL               respjson.Field
		LogoURL               respjson.Field
		PhoneNumber           respjson.Field
		PrivacyPolicyURL      respjson.Field
		TermsAndConditionsURL respjson.Field
		Website               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentConfigurationBasicsAgentProfileContactRequirement) RawJSON() string { return r.JSON.raw }
func (r *AgentConfigurationBasicsAgentProfileContactRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Basics is required.
type AgentConfigurationParam struct {
	// Basic agent identity and contact information. At least one complete phone,
	// website, or email contact is required.
	Basics   AgentConfigurationBasicsUnionParam `json:"basics,omitzero" api:"required"`
	Campaign AgentCampaignConfigurationParam    `json:"campaign,omitzero"`
	Testing  AgentTestingConfigurationParam     `json:"testing,omitzero"`
	paramObj
}

func (r AgentConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentConfigurationBasicsUnionParam struct {
	OfAgentPhoneContactRequirement   *AgentConfigurationBasicsAgentPhoneContactRequirementParam   `json:",omitzero,inline"`
	OfAgentWebhookContactRequirement *AgentConfigurationBasicsAgentWebhookContactRequirementParam `json:",omitzero,inline"`
	OfAgentProfileContactRequirement *AgentConfigurationBasicsAgentProfileContactRequirementParam `json:",omitzero,inline"`
	paramUnion
}

func (u AgentConfigurationBasicsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAgentPhoneContactRequirement, u.OfAgentWebhookContactRequirement, u.OfAgentProfileContactRequirement)
}
func (u *AgentConfigurationBasicsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentConfigurationBasicsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAgentPhoneContactRequirement) {
		return u.OfAgentPhoneContactRequirement
	} else if !param.IsOmitted(u.OfAgentWebhookContactRequirement) {
		return u.OfAgentWebhookContactRequirement
	} else if !param.IsOmitted(u.OfAgentProfileContactRequirement) {
		return u.OfAgentProfileContactRequirement
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentConfigurationBasicsUnionParam) GetBrandColor() *string {
	if vt := u.OfAgentPhoneContactRequirement; vt != nil && vt.BrandColor.Valid() {
		return &vt.BrandColor.Value
	} else if vt := u.OfAgentWebhookContactRequirement; vt != nil && vt.BrandColor.Valid() {
		return &vt.BrandColor.Value
	} else if vt := u.OfAgentProfileContactRequirement; vt != nil && vt.BrandColor.Valid() {
		return &vt.BrandColor.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentConfigurationBasicsUnionParam) GetDescription() *string {
	if vt := u.OfAgentPhoneContactRequirement; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfAgentWebhookContactRequirement; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfAgentProfileContactRequirement; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentConfigurationBasicsUnionParam) GetHeroURL() *string {
	if vt := u.OfAgentPhoneContactRequirement; vt != nil && vt.HeroURL.Valid() {
		return &vt.HeroURL.Value
	} else if vt := u.OfAgentWebhookContactRequirement; vt != nil && vt.HeroURL.Valid() {
		return &vt.HeroURL.Value
	} else if vt := u.OfAgentProfileContactRequirement; vt != nil && vt.HeroURL.Valid() {
		return &vt.HeroURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentConfigurationBasicsUnionParam) GetLogoURL() *string {
	if vt := u.OfAgentPhoneContactRequirement; vt != nil && vt.LogoURL.Valid() {
		return &vt.LogoURL.Value
	} else if vt := u.OfAgentWebhookContactRequirement; vt != nil && vt.LogoURL.Valid() {
		return &vt.LogoURL.Value
	} else if vt := u.OfAgentProfileContactRequirement; vt != nil && vt.LogoURL.Valid() {
		return &vt.LogoURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentConfigurationBasicsUnionParam) GetPrivacyPolicyURL() *string {
	if vt := u.OfAgentPhoneContactRequirement; vt != nil && vt.PrivacyPolicyURL.Valid() {
		return &vt.PrivacyPolicyURL.Value
	} else if vt := u.OfAgentWebhookContactRequirement; vt != nil && vt.PrivacyPolicyURL.Valid() {
		return &vt.PrivacyPolicyURL.Value
	} else if vt := u.OfAgentProfileContactRequirement; vt != nil && vt.PrivacyPolicyURL.Valid() {
		return &vt.PrivacyPolicyURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentConfigurationBasicsUnionParam) GetTermsAndConditionsURL() *string {
	if vt := u.OfAgentPhoneContactRequirement; vt != nil && vt.TermsAndConditionsURL.Valid() {
		return &vt.TermsAndConditionsURL.Value
	} else if vt := u.OfAgentWebhookContactRequirement; vt != nil && vt.TermsAndConditionsURL.Valid() {
		return &vt.TermsAndConditionsURL.Value
	} else if vt := u.OfAgentProfileContactRequirement; vt != nil && vt.TermsAndConditionsURL.Valid() {
		return &vt.TermsAndConditionsURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's PhoneNumber property, if present.
func (u AgentConfigurationBasicsUnionParam) GetPhoneNumber() *AgentPhoneContactParam {
	if vt := u.OfAgentPhoneContactRequirement; vt != nil {
		return &vt.PhoneNumber
	} else if vt := u.OfAgentWebhookContactRequirement; vt != nil {
		return &vt.PhoneNumber
	} else if vt := u.OfAgentProfileContactRequirement; vt != nil {
		return &vt.PhoneNumber
	}
	return nil
}

// Returns a pointer to the underlying variant's Email property, if present.
func (u AgentConfigurationBasicsUnionParam) GetEmail() *AgentEmailContactParam {
	if vt := u.OfAgentPhoneContactRequirement; vt != nil {
		return &vt.Email
	} else if vt := u.OfAgentWebhookContactRequirement; vt != nil {
		return &vt.Email
	} else if vt := u.OfAgentProfileContactRequirement; vt != nil {
		return &vt.Email
	}
	return nil
}

// Returns a pointer to the underlying variant's Website property, if present.
func (u AgentConfigurationBasicsUnionParam) GetWebsite() *AgentWebsiteContactParam {
	if vt := u.OfAgentPhoneContactRequirement; vt != nil {
		return &vt.Website
	} else if vt := u.OfAgentWebhookContactRequirement; vt != nil {
		return &vt.Website
	} else if vt := u.OfAgentProfileContactRequirement; vt != nil {
		return &vt.Website
	}
	return nil
}

// The property PhoneNumber is required.
type AgentConfigurationBasicsAgentPhoneContactRequirementParam struct {
	PhoneNumber           AgentPhoneContactParam   `json:"phone_number,omitzero" api:"required"`
	BrandColor            param.Opt[string]        `json:"brand_color,omitzero"`
	Description           param.Opt[string]        `json:"description,omitzero"`
	HeroURL               param.Opt[string]        `json:"hero_url,omitzero" format:"uri"`
	LogoURL               param.Opt[string]        `json:"logo_url,omitzero" format:"uri"`
	PrivacyPolicyURL      param.Opt[string]        `json:"privacy_policy_url,omitzero" format:"uri"`
	TermsAndConditionsURL param.Opt[string]        `json:"terms_and_conditions_url,omitzero" format:"uri"`
	Email                 AgentEmailContactParam   `json:"email,omitzero"`
	Website               AgentWebsiteContactParam `json:"website,omitzero"`
	paramObj
}

func (r AgentConfigurationBasicsAgentPhoneContactRequirementParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentConfigurationBasicsAgentPhoneContactRequirementParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentConfigurationBasicsAgentPhoneContactRequirementParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Website is required.
type AgentConfigurationBasicsAgentWebhookContactRequirementParam struct {
	Website               AgentWebsiteContactParam `json:"website,omitzero" api:"required"`
	BrandColor            param.Opt[string]        `json:"brand_color,omitzero"`
	Description           param.Opt[string]        `json:"description,omitzero"`
	HeroURL               param.Opt[string]        `json:"hero_url,omitzero" format:"uri"`
	LogoURL               param.Opt[string]        `json:"logo_url,omitzero" format:"uri"`
	PrivacyPolicyURL      param.Opt[string]        `json:"privacy_policy_url,omitzero" format:"uri"`
	TermsAndConditionsURL param.Opt[string]        `json:"terms_and_conditions_url,omitzero" format:"uri"`
	Email                 AgentEmailContactParam   `json:"email,omitzero"`
	PhoneNumber           AgentPhoneContactParam   `json:"phone_number,omitzero"`
	paramObj
}

func (r AgentConfigurationBasicsAgentWebhookContactRequirementParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentConfigurationBasicsAgentWebhookContactRequirementParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentConfigurationBasicsAgentWebhookContactRequirementParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Email is required.
type AgentConfigurationBasicsAgentProfileContactRequirementParam struct {
	Email                 AgentEmailContactParam   `json:"email,omitzero" api:"required"`
	BrandColor            param.Opt[string]        `json:"brand_color,omitzero"`
	Description           param.Opt[string]        `json:"description,omitzero"`
	HeroURL               param.Opt[string]        `json:"hero_url,omitzero" format:"uri"`
	LogoURL               param.Opt[string]        `json:"logo_url,omitzero" format:"uri"`
	PrivacyPolicyURL      param.Opt[string]        `json:"privacy_policy_url,omitzero" format:"uri"`
	TermsAndConditionsURL param.Opt[string]        `json:"terms_and_conditions_url,omitzero" format:"uri"`
	PhoneNumber           AgentPhoneContactParam   `json:"phone_number,omitzero"`
	Website               AgentWebsiteContactParam `json:"website,omitzero"`
	paramObj
}

func (r AgentConfigurationBasicsAgentProfileContactRequirementParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentConfigurationBasicsAgentProfileContactRequirementParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentConfigurationBasicsAgentProfileContactRequirementParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentConsentConfiguration struct {
	CallToAction   string                                 `json:"call_to_action" api:"required"`
	DoubleOptIn    bool                                   `json:"double_opt_in" api:"required"`
	HelpResponse   string                                 `json:"help_response" api:"required"`
	OptInMessage   string                                 `json:"opt_in_message" api:"required"`
	OptInMethods   []AgentConsentConfigurationOptInMethod `json:"opt_in_methods" api:"required"`
	OptOutResponse string                                 `json:"opt_out_response" api:"required"`
	// Required when an opt-in method is `WEBSITE` or `MOBILE_APP`.
	CallToActionMediaURL string `json:"call_to_action_media_url" api:"nullable" format:"uri"`
	// Required when an opt-in method is `WEBSITE`.
	CallToActionURL string `json:"call_to_action_url" api:"nullable" format:"uri"`
	// Required when double_opt_in is true.
	DoubleOptInMessage string `json:"double_opt_in_message" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallToAction         respjson.Field
		DoubleOptIn          respjson.Field
		HelpResponse         respjson.Field
		OptInMessage         respjson.Field
		OptInMethods         respjson.Field
		OptOutResponse       respjson.Field
		CallToActionMediaURL respjson.Field
		CallToActionURL      respjson.Field
		DoubleOptInMessage   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentConsentConfiguration) RawJSON() string { return r.JSON.raw }
func (r *AgentConsentConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentConsentConfiguration to a
// AgentConsentConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentConsentConfigurationParam.Overrides()
func (r AgentConsentConfiguration) ToParam() AgentConsentConfigurationParam {
	return param.Override[AgentConsentConfigurationParam](json.RawMessage(r.RawJSON()))
}

type AgentConsentConfigurationOptInMethod struct {
	// Any of "SMS", "WEBSITE", "MOBILE_APP", "QR_CODE", "SALE_POINT", "OTHER".
	MethodType string `json:"method_type" api:"required"`
	// Required when method_type is `OTHER`.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MethodType  respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentConsentConfigurationOptInMethod) RawJSON() string { return r.JSON.raw }
func (r *AgentConsentConfigurationOptInMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CallToAction, DoubleOptIn, HelpResponse, OptInMessage,
// OptInMethods, OptOutResponse are required.
type AgentConsentConfigurationParam struct {
	CallToAction   string                                      `json:"call_to_action" api:"required"`
	DoubleOptIn    bool                                        `json:"double_opt_in" api:"required"`
	HelpResponse   string                                      `json:"help_response" api:"required"`
	OptInMessage   string                                      `json:"opt_in_message" api:"required"`
	OptInMethods   []AgentConsentConfigurationOptInMethodParam `json:"opt_in_methods,omitzero" api:"required"`
	OptOutResponse string                                      `json:"opt_out_response" api:"required"`
	// Required when an opt-in method is `WEBSITE` or `MOBILE_APP`.
	CallToActionMediaURL param.Opt[string] `json:"call_to_action_media_url,omitzero" format:"uri"`
	// Required when an opt-in method is `WEBSITE`.
	CallToActionURL param.Opt[string] `json:"call_to_action_url,omitzero" format:"uri"`
	// Required when double_opt_in is true.
	DoubleOptInMessage param.Opt[string] `json:"double_opt_in_message,omitzero"`
	paramObj
}

func (r AgentConsentConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentConsentConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentConsentConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property MethodType is required.
type AgentConsentConfigurationOptInMethodParam struct {
	// Any of "SMS", "WEBSITE", "MOBILE_APP", "QR_CODE", "SALE_POINT", "OTHER".
	MethodType string `json:"method_type,omitzero" api:"required"`
	// Required when method_type is `OTHER`.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r AgentConsentConfigurationOptInMethodParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentConsentConfigurationOptInMethodParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentConsentConfigurationOptInMethodParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AgentConsentConfigurationOptInMethodParam](
		"method_type", "SMS", "WEBSITE", "MOBILE_APP", "QR_CODE", "SALE_POINT", "OTHER",
	)
}

type AgentEmailContact struct {
	Address string `json:"address" api:"required" format:"email"`
	Label   string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentEmailContact) RawJSON() string { return r.JSON.raw }
func (r *AgentEmailContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentEmailContact to a AgentEmailContactParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentEmailContactParam.Overrides()
func (r AgentEmailContact) ToParam() AgentEmailContactParam {
	return param.Override[AgentEmailContactParam](json.RawMessage(r.RawJSON()))
}

// The properties Address, Label are required.
type AgentEmailContactParam struct {
	Address string `json:"address" api:"required" format:"email"`
	Label   string `json:"label" api:"required"`
	paramObj
}

func (r AgentEmailContactParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentEmailContactParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentEmailContactParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentInteraction struct {
	// Any of "TRANSACTIONAL_UPDATES", "CUSTOMER_SUPPORT", "LOYALTY_OR_REWARD",
	// "MARKETING_OR_PROMOTIONAL", "ACCOUNT_ALERTS", "TWO_WAY_CONVERSATION", "OTHER".
	InteractionType AgentInteractionInteractionType `json:"interaction_type" api:"required"`
	// Required when interaction_type is `OTHER`.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InteractionType respjson.Field
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentInteraction) RawJSON() string { return r.JSON.raw }
func (r *AgentInteraction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentInteraction to a AgentInteractionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentInteractionParam.Overrides()
func (r AgentInteraction) ToParam() AgentInteractionParam {
	return param.Override[AgentInteractionParam](json.RawMessage(r.RawJSON()))
}

type AgentInteractionInteractionType string

const (
	AgentInteractionInteractionTypeTransactionalUpdates   AgentInteractionInteractionType = "TRANSACTIONAL_UPDATES"
	AgentInteractionInteractionTypeCustomerSupport        AgentInteractionInteractionType = "CUSTOMER_SUPPORT"
	AgentInteractionInteractionTypeLoyaltyOrReward        AgentInteractionInteractionType = "LOYALTY_OR_REWARD"
	AgentInteractionInteractionTypeMarketingOrPromotional AgentInteractionInteractionType = "MARKETING_OR_PROMOTIONAL"
	AgentInteractionInteractionTypeAccountAlerts          AgentInteractionInteractionType = "ACCOUNT_ALERTS"
	AgentInteractionInteractionTypeTwoWayConversation     AgentInteractionInteractionType = "TWO_WAY_CONVERSATION"
	AgentInteractionInteractionTypeOther                  AgentInteractionInteractionType = "OTHER"
)

// The property InteractionType is required.
type AgentInteractionParam struct {
	// Any of "TRANSACTIONAL_UPDATES", "CUSTOMER_SUPPORT", "LOYALTY_OR_REWARD",
	// "MARKETING_OR_PROMOTIONAL", "ACCOUNT_ALERTS", "TWO_WAY_CONVERSATION", "OTHER".
	InteractionType AgentInteractionInteractionType `json:"interaction_type,omitzero" api:"required"`
	// Required when interaction_type is `OTHER`.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r AgentInteractionParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentInteractionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentInteractionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentPhoneContact struct {
	Label  string `json:"label" api:"required"`
	Number string `json:"number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Number      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentPhoneContact) RawJSON() string { return r.JSON.raw }
func (r *AgentPhoneContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentPhoneContact to a AgentPhoneContactParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentPhoneContactParam.Overrides()
func (r AgentPhoneContact) ToParam() AgentPhoneContactParam {
	return param.Override[AgentPhoneContactParam](json.RawMessage(r.RawJSON()))
}

// The properties Label, Number are required.
type AgentPhoneContactParam struct {
	Label  string `json:"label" api:"required"`
	Number string `json:"number" api:"required"`
	paramObj
}

func (r AgentPhoneContactParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentPhoneContactParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentPhoneContactParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponse struct {
	AgentID string `json:"agent_id" api:"required" format:"uuid"`
	// Any of "SUBMITTED", "APPROVED", "REJECTED".
	BasicsStatus AgentSubmissionStatus `json:"basics_status" api:"required"`
	// Any of "NON_CONVERSATIONAL", "CONVERSATIONAL".
	BillingCategory AgentResponseBillingCategory `json:"billing_category" api:"required"`
	BrandID         string                       `json:"brand_id" api:"required" format:"uuid"`
	// Any of "SUBMITTED", "APPROVED", "REJECTED".
	CampaignStatus   AgentSubmissionStatus     `json:"campaign_status" api:"required"`
	Capabilities     CapabilitiesResponse      `json:"capabilities" api:"required"`
	CarrierApprovals []CarrierApprovalResponse `json:"carrier_approvals" api:"required"`
	Configuration    AgentConfiguration        `json:"configuration" api:"required"`
	DisplayName      string                    `json:"display_name" api:"required"`
	HostingRegion    string                    `json:"hosting_region" api:"required"`
	ProfileID        string                    `json:"profile_id" api:"required"`
	// Any of "CREATED", "SUBMITTED", "VERIFYING", "VERIFIED", "LAUNCHING", "LAUNCHED",
	// "LIVE", "REJECTED", "FAILED".
	Status      AgentResponseStatus  `json:"status" api:"required"`
	TestDevices []TestDeviceResponse `json:"test_devices" api:"required"`
	// Any of "SUBMITTED", "APPROVED", "REJECTED".
	TestingStatus AgentSubmissionStatus `json:"testing_status" api:"required"`
	// Any of "MULTI_USE", "PROMOTIONAL", "TRANSACTIONAL", "OTP".
	UseCase AgentUseCase `json:"use_case" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID          respjson.Field
		BasicsStatus     respjson.Field
		BillingCategory  respjson.Field
		BrandID          respjson.Field
		CampaignStatus   respjson.Field
		Capabilities     respjson.Field
		CarrierApprovals respjson.Field
		Configuration    respjson.Field
		DisplayName      respjson.Field
		HostingRegion    respjson.Field
		ProfileID        respjson.Field
		Status           respjson.Field
		TestDevices      respjson.Field
		TestingStatus    respjson.Field
		UseCase          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseBillingCategory string

const (
	AgentResponseBillingCategoryNonConversational AgentResponseBillingCategory = "NON_CONVERSATIONAL"
	AgentResponseBillingCategoryConversational    AgentResponseBillingCategory = "CONVERSATIONAL"
)

type AgentResponseStatus string

const (
	AgentResponseStatusCreated   AgentResponseStatus = "CREATED"
	AgentResponseStatusSubmitted AgentResponseStatus = "SUBMITTED"
	AgentResponseStatusVerifying AgentResponseStatus = "VERIFYING"
	AgentResponseStatusVerified  AgentResponseStatus = "VERIFIED"
	AgentResponseStatusLaunching AgentResponseStatus = "LAUNCHING"
	AgentResponseStatusLaunched  AgentResponseStatus = "LAUNCHED"
	AgentResponseStatusLive      AgentResponseStatus = "LIVE"
	AgentResponseStatusRejected  AgentResponseStatus = "REJECTED"
	AgentResponseStatusFailed    AgentResponseStatus = "FAILED"
)

type AgentSubmissionStatus string

const (
	AgentSubmissionStatusSubmitted AgentSubmissionStatus = "SUBMITTED"
	AgentSubmissionStatusApproved  AgentSubmissionStatus = "APPROVED"
	AgentSubmissionStatusRejected  AgentSubmissionStatus = "REJECTED"
)

type AgentTestingConfiguration struct {
	// A publicly accessible test video or evidence URL.
	TestURL               string `json:"test_url" api:"required" format:"uri"`
	AdditionalInformation string `json:"additional_information" api:"nullable"`
	MessageID             string `json:"message_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TestURL               respjson.Field
		AdditionalInformation respjson.Field
		MessageID             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTestingConfiguration) RawJSON() string { return r.JSON.raw }
func (r *AgentTestingConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentTestingConfiguration to a
// AgentTestingConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentTestingConfigurationParam.Overrides()
func (r AgentTestingConfiguration) ToParam() AgentTestingConfigurationParam {
	return param.Override[AgentTestingConfigurationParam](json.RawMessage(r.RawJSON()))
}

// The property TestURL is required.
type AgentTestingConfigurationParam struct {
	// A publicly accessible test video or evidence URL.
	TestURL               string            `json:"test_url" api:"required" format:"uri"`
	AdditionalInformation param.Opt[string] `json:"additional_information,omitzero"`
	MessageID             param.Opt[string] `json:"message_id,omitzero"`
	paramObj
}

func (r AgentTestingConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentTestingConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTestingConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentUseCase string

const (
	AgentUseCaseMultiUse      AgentUseCase = "MULTI_USE"
	AgentUseCasePromotional   AgentUseCase = "PROMOTIONAL"
	AgentUseCaseTransactional AgentUseCase = "TRANSACTIONAL"
	AgentUseCaseOtp           AgentUseCase = "OTP"
)

type AgentWebsiteContact struct {
	Label string `json:"label" api:"required"`
	URL   string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentWebsiteContact) RawJSON() string { return r.JSON.raw }
func (r *AgentWebsiteContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentWebsiteContact to a AgentWebsiteContactParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentWebsiteContactParam.Overrides()
func (r AgentWebsiteContact) ToParam() AgentWebsiteContactParam {
	return param.Override[AgentWebsiteContactParam](json.RawMessage(r.RawJSON()))
}

// The properties Label, URL are required.
type AgentWebsiteContactParam struct {
	Label string `json:"label" api:"required"`
	URL   string `json:"url" api:"required" format:"uri"`
	paramObj
}

func (r AgentWebsiteContactParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentWebsiteContactParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentWebsiteContactParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilitiesResponse struct {
	BrandEntity         bool `json:"brand_entity" api:"required"`
	BrandVerification   bool `json:"brand_verification" api:"required"`
	Campaigns           bool `json:"campaigns" api:"required"`
	DistinctLaunchPhase bool `json:"distinct_launch_phase" api:"required"`
	InviteTestDevices   bool `json:"invite_test_devices" api:"required"`
	PerCarrierApproval  bool `json:"per_carrier_approval" api:"required"`
	SubmissionSections  bool `json:"submission_sections" api:"required"`
	Templates           bool `json:"templates" api:"required"`
	VendorWebhooks      bool `json:"vendor_webhooks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandEntity         respjson.Field
		BrandVerification   respjson.Field
		Campaigns           respjson.Field
		DistinctLaunchPhase respjson.Field
		InviteTestDevices   respjson.Field
		PerCarrierApproval  respjson.Field
		SubmissionSections  respjson.Field
		Templates           respjson.Field
		VendorWebhooks      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *CapabilitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CarrierApprovalResponse struct {
	ApprovalID     string    `json:"approval_id" api:"required" format:"uuid"`
	ApprovedAt     time.Time `json:"approved_at" api:"required" format:"date-time"`
	Carrier        string    `json:"carrier" api:"required"`
	RejectedReason string    `json:"rejected_reason" api:"required"`
	// Any of "carrier", "hub", "bot".
	ScopeType CarrierApprovalResponseScopeType `json:"scope_type" api:"required"`
	// Any of "PENDING", "SUBMITTED", "APPROVED", "REJECTED".
	Status      CarrierApprovalResponseStatus `json:"status" api:"required"`
	SubmittedAt time.Time                     `json:"submitted_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalID     respjson.Field
		ApprovedAt     respjson.Field
		Carrier        respjson.Field
		RejectedReason respjson.Field
		ScopeType      respjson.Field
		Status         respjson.Field
		SubmittedAt    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CarrierApprovalResponse) RawJSON() string { return r.JSON.raw }
func (r *CarrierApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CarrierApprovalResponseScopeType string

const (
	CarrierApprovalResponseScopeTypeCarrier CarrierApprovalResponseScopeType = "carrier"
	CarrierApprovalResponseScopeTypeHub     CarrierApprovalResponseScopeType = "hub"
	CarrierApprovalResponseScopeTypeBot     CarrierApprovalResponseScopeType = "bot"
)

type CarrierApprovalResponseStatus string

const (
	CarrierApprovalResponseStatusPending   CarrierApprovalResponseStatus = "PENDING"
	CarrierApprovalResponseStatusSubmitted CarrierApprovalResponseStatus = "SUBMITTED"
	CarrierApprovalResponseStatusApproved  CarrierApprovalResponseStatus = "APPROVED"
	CarrierApprovalResponseStatusRejected  CarrierApprovalResponseStatus = "REJECTED"
)

type RcsAgent struct {
	// RCS Agent ID
	AgentID string `json:"agent_id"`
	// Human readable agent name
	AgentName string `json:"agent_name"`
	// Date and time when the resource was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies whether the agent is enabled
	Enabled bool `json:"enabled"`
	// Messaging profile ID associated with the RCS Agent
	ProfileID string `json:"profile_id" api:"nullable" format:"uuid"`
	// Date and time when the resource was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// User ID associated with the RCS Agent
	UserID string `json:"user_id"`
	// Failover URL to receive RCS events
	WebhookFailoverURL string `json:"webhook_failover_url" api:"nullable" format:"url"`
	// URL to receive RCS events
	WebhookURL string `json:"webhook_url" api:"nullable" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID            respjson.Field
		AgentName          respjson.Field
		CreatedAt          respjson.Field
		Enabled            respjson.Field
		ProfileID          respjson.Field
		UpdatedAt          respjson.Field
		UserID             respjson.Field
		WebhookFailoverURL respjson.Field
		WebhookURL         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgent) RawJSON() string { return r.JSON.raw }
func (r *RcsAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsAgentResponse struct {
	Data RcsAgent `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentResponse) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcAgentNewParams struct {
	BrandID       string                  `json:"brand_id" api:"required" format:"uuid"`
	Configuration AgentConfigurationParam `json:"configuration,omitzero" api:"required"`
	DisplayName   string                  `json:"display_name" api:"required"`
	// Any of "MULTI_USE", "PROMOTIONAL", "TRANSACTIONAL", "OTP".
	UseCase        AgentUseCase      `json:"use_case,omitzero" api:"required"`
	IdempotencyKey string            `header:"Idempotency-Key" api:"required" json:"-"`
	HostingRegion  param.Opt[string] `json:"hosting_region,omitzero"`
	// A Messaging Profile owned by the authenticated organization. When omitted, the
	// agent inherits the brand profile.
	ProfileID param.Opt[string] `json:"profile_id,omitzero"`
	paramObj
}

func (r RcAgentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RcAgentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcAgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcAgentUpdateParams struct {
	DisplayName   param.Opt[string]       `json:"display_name,omitzero"`
	HostingRegion param.Opt[string]       `json:"hosting_region,omitzero"`
	ProfileID     param.Opt[string]       `json:"profile_id,omitzero"`
	Configuration AgentConfigurationParam `json:"configuration,omitzero"`
	// Any of "MULTI_USE", "PROMOTIONAL", "TRANSACTIONAL", "OTP".
	UseCase AgentUseCase `json:"use_case,omitzero"`
	paramObj
}

func (r RcAgentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RcAgentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcAgentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcAgentListParams struct {
	// Only return agents belonging to this brand.
	BrandID param.Opt[string] `query:"brand_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [RcAgentListParams]'s query parameters as `url.Values`.
func (r RcAgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RcAgentLaunchParams struct {
	Campaign RcAgentLaunchParamsCampaign    `json:"campaign,omitzero" api:"required"`
	Testing  AgentTestingConfigurationParam `json:"testing,omitzero" api:"required"`
	paramObj
}

func (r RcAgentLaunchParams) MarshalJSON() (data []byte, err error) {
	type shadow RcAgentLaunchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcAgentLaunchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcAgentLaunchParamsCampaign struct {
	AgentOverview   string                         `json:"agent_overview" api:"required"`
	ConsentSettings AgentConsentConfigurationParam `json:"consent_settings,omitzero" api:"required"`
	Interactions    []AgentInteractionParam        `json:"interactions,omitzero" api:"required"`
	MessageExamples []string                       `json:"message_examples,omitzero" api:"required"`
	AgentCampaignConfigurationParam
}

func (r RcAgentLaunchParamsCampaign) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*RcAgentLaunchParamsCampaign
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

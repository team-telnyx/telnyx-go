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
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// Manage Whatsapp message templates
//
// WhatsappTemplateService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappTemplateService] method instead.
type WhatsappTemplateService struct {
	Options []option.RequestOption
}

// NewWhatsappTemplateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhatsappTemplateService(opts ...option.RequestOption) (r WhatsappTemplateService) {
	r = WhatsappTemplateService{}
	r.Options = opts
	return
}

// Create a Whatsapp message template
func (r *WhatsappTemplateService) New(ctx context.Context, body WhatsappTemplateNewParams, opts ...option.RequestOption) (res *WhatsappTemplateNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/whatsapp/message_templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List Whatsapp message templates
func (r *WhatsappTemplateService) List(ctx context.Context, query WhatsappTemplateListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[shared.WhatsappTemplateData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/whatsapp/message_templates"
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

// List Whatsapp message templates
func (r *WhatsappTemplateService) ListAutoPaging(ctx context.Context, query WhatsappTemplateListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[shared.WhatsappTemplateData] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type WhatsappTemplateNewResponse struct {
	Data shared.WhatsappTemplateData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappTemplateNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappTemplateNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappTemplateNewParams struct {
	// Template category: AUTHENTICATION, UTILITY, or MARKETING.
	//
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	Category WhatsappTemplateNewParamsCategory `json:"category,omitzero" api:"required"`
	// Template components defining message structure. Passed through to Meta Graph
	// API. Templates with variables must include example values. Supports HEADER,
	// BODY, FOOTER, BUTTONS, CAROUSEL and any future Meta component types.
	Components []WhatsappTemplateNewParamsComponentUnion `json:"components,omitzero" api:"required"`
	// Template language code (e.g. en_US, es, pt_BR).
	Language string `json:"language" api:"required"`
	// Template name. Lowercase letters, numbers, and underscores only.
	Name string `json:"name" api:"required"`
	// The WhatsApp Business Account ID.
	WabaID string `json:"waba_id" api:"required"`
	paramObj
}

func (r WhatsappTemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template category: AUTHENTICATION, UTILITY, or MARKETING.
type WhatsappTemplateNewParamsCategory string

const (
	WhatsappTemplateNewParamsCategoryMarketing      WhatsappTemplateNewParamsCategory = "MARKETING"
	WhatsappTemplateNewParamsCategoryUtility        WhatsappTemplateNewParamsCategory = "UTILITY"
	WhatsappTemplateNewParamsCategoryAuthentication WhatsappTemplateNewParamsCategory = "AUTHENTICATION"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WhatsappTemplateNewParamsComponentUnion struct {
	OfHeader   *WhatsappTemplateNewParamsComponentHeader   `json:",omitzero,inline"`
	OfBody     *WhatsappTemplateNewParamsComponentBody     `json:",omitzero,inline"`
	OfFooter   *WhatsappTemplateNewParamsComponentFooter   `json:",omitzero,inline"`
	OfButtons  *WhatsappTemplateNewParamsComponentButtons  `json:",omitzero,inline"`
	OfCarousel *WhatsappTemplateNewParamsComponentCarousel `json:",omitzero,inline"`
	paramUnion
}

func (u WhatsappTemplateNewParamsComponentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHeader,
		u.OfBody,
		u.OfFooter,
		u.OfButtons,
		u.OfCarousel)
}
func (u *WhatsappTemplateNewParamsComponentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WhatsappTemplateNewParamsComponentUnion) asAny() any {
	if !param.IsOmitted(u.OfHeader) {
		return u.OfHeader
	} else if !param.IsOmitted(u.OfBody) {
		return u.OfBody
	} else if !param.IsOmitted(u.OfFooter) {
		return u.OfFooter
	} else if !param.IsOmitted(u.OfButtons) {
		return u.OfButtons
	} else if !param.IsOmitted(u.OfCarousel) {
		return u.OfCarousel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappTemplateNewParamsComponentUnion) GetFormat() *string {
	if vt := u.OfHeader; vt != nil {
		return &vt.Format
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappTemplateNewParamsComponentUnion) GetCodeExpirationMinutes() *int64 {
	if vt := u.OfFooter; vt != nil && vt.CodeExpirationMinutes.Valid() {
		return &vt.CodeExpirationMinutes.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappTemplateNewParamsComponentUnion) GetButtons() []WhatsappTemplateNewParamsComponentButtonsButton {
	if vt := u.OfButtons; vt != nil {
		return vt.Buttons
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappTemplateNewParamsComponentUnion) GetCards() []WhatsappTemplateNewParamsComponentCarouselCard {
	if vt := u.OfCarousel; vt != nil {
		return vt.Cards
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappTemplateNewParamsComponentUnion) GetType() *string {
	if vt := u.OfHeader; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBody; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFooter; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfButtons; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCarousel; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappTemplateNewParamsComponentUnion) GetText() *string {
	if vt := u.OfHeader; vt != nil && vt.Text.Valid() {
		return &vt.Text.Value
	} else if vt := u.OfBody; vt != nil && vt.Text.Valid() {
		return &vt.Text.Value
	} else if vt := u.OfFooter; vt != nil && vt.Text.Valid() {
		return &vt.Text.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WhatsappTemplateNewParamsComponentUnion) GetExample() (res whatsappTemplateNewParamsComponentUnionExample) {
	if vt := u.OfHeader; vt != nil {
		res.any = &vt.Example
	} else if vt := u.OfBody; vt != nil {
		res.any = &vt.Example
	}
	return
}

// Can have the runtime types [*WhatsappTemplateNewParamsComponentHeaderExample],
// [*WhatsappTemplateNewParamsComponentBodyExample]
type whatsappTemplateNewParamsComponentUnionExample struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *telnyx.WhatsappTemplateNewParamsComponentHeaderExample:
//	case *telnyx.WhatsappTemplateNewParamsComponentBodyExample:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u whatsappTemplateNewParamsComponentUnionExample) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[WhatsappTemplateNewParamsComponentUnion](
		"type",
		apijson.Discriminator[WhatsappTemplateNewParamsComponentHeader]("HEADER"),
		apijson.Discriminator[WhatsappTemplateNewParamsComponentBody]("BODY"),
		apijson.Discriminator[WhatsappTemplateNewParamsComponentFooter]("FOOTER"),
		apijson.Discriminator[WhatsappTemplateNewParamsComponentButtons]("BUTTONS"),
		apijson.Discriminator[WhatsappTemplateNewParamsComponentCarousel]("CAROUSEL"),
	)
}

// Optional header displayed at the top of the message.
//
// The properties Format, Type are required.
type WhatsappTemplateNewParamsComponentHeader struct {
	// Header format type: TEXT (supports one variable), IMAGE, VIDEO, DOCUMENT, or
	// LOCATION.
	//
	// Any of "TEXT", "IMAGE", "VIDEO", "DOCUMENT", "LOCATION".
	Format string `json:"format,omitzero" api:"required"`
	// Header text. Required when format is TEXT. Supports one variable ({{1}}).
	// Variables cannot be at the start or end.
	Text param.Opt[string] `json:"text,omitzero"`
	// Sample values for header variables.
	Example WhatsappTemplateNewParamsComponentHeaderExample `json:"example,omitzero"`
	// This field can be elided, and will marshal its zero value as "HEADER".
	Type constant.Header `json:"type" api:"required"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentHeader) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentHeader](
		"format", "TEXT", "IMAGE", "VIDEO", "DOCUMENT", "LOCATION",
	)
}

// Sample values for header variables.
type WhatsappTemplateNewParamsComponentHeaderExample struct {
	// Media handle for IMAGE, VIDEO, or DOCUMENT headers.
	HeaderHandle []string `json:"header_handle,omitzero"`
	// Sample values for text header variables.
	HeaderText []string `json:"header_text,omitzero"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentHeaderExample) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentHeaderExample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentHeaderExample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The main text content of the message. Supports multiple variable parameters
// ({{1}}, {{2}}, etc.). Variables cannot be at the start or end. Maximum 1024
// characters.
//
// The property Type is required.
type WhatsappTemplateNewParamsComponentBody struct {
	// Body text content. Use {{1}}, {{2}}, etc. for variable placeholders. Required
	// for MARKETING and UTILITY templates. Optional for AUTHENTICATION templates where
	// Meta provides the built-in OTP body.
	Text param.Opt[string] `json:"text,omitzero"`
	// Sample values for body variables. Required when body text contains parameters.
	Example WhatsappTemplateNewParamsComponentBodyExample `json:"example,omitzero"`
	// This field can be elided, and will marshal its zero value as "BODY".
	Type constant.Body `json:"type" api:"required"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentBody) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sample values for body variables. Required when body text contains parameters.
type WhatsappTemplateNewParamsComponentBodyExample struct {
	// Array containing one array of sample values, one per variable in order.
	BodyText [][]string `json:"body_text,omitzero"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentBodyExample) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentBodyExample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentBodyExample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional footer displayed at the bottom of the message. Does not support
// variables.
//
// The property Type is required.
type WhatsappTemplateNewParamsComponentFooter struct {
	// OTP code expiration time in minutes. Used in AUTHENTICATION template footers
	// instead of free-form text.
	CodeExpirationMinutes param.Opt[int64] `json:"code_expiration_minutes,omitzero"`
	// Footer text. Maximum 60 characters. For non-authentication templates.
	Text param.Opt[string] `json:"text,omitzero"`
	// This field can be elided, and will marshal its zero value as "FOOTER".
	Type constant.Footer `json:"type" api:"required"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentFooter) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentFooter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentFooter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional interactive buttons. Maximum 3 buttons per template.
//
// The properties Buttons, Type are required.
type WhatsappTemplateNewParamsComponentButtons struct {
	// Array of button objects. Meta supports various combinations of button types.
	Buttons []WhatsappTemplateNewParamsComponentButtonsButton `json:"buttons,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "BUTTONS".
	Type constant.Buttons `json:"type" api:"required"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentButtons) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentButtons
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentButtons) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type WhatsappTemplateNewParamsComponentButtonsButton struct {
	// Any of "URL", "PHONE_NUMBER", "QUICK_REPLY", "OTP", "COPY_CODE", "FLOW".
	Type string `json:"type,omitzero" api:"required"`
	// Custom autofill button text for ONE_TAP OTP buttons.
	AutofillText param.Opt[string] `json:"autofill_text,omitzero"`
	// Flow ID for FLOW-type buttons.
	FlowID param.Opt[string] `json:"flow_id,omitzero"`
	// Target screen name for FLOW buttons with navigate action.
	NavigateScreen param.Opt[string] `json:"navigate_screen,omitzero"`
	// Android package name. Required for ONE_TAP OTP buttons.
	PackageName param.Opt[string] `json:"package_name,omitzero"`
	// Phone number in E.164 format.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// Android app signing key hash. Required for ONE_TAP OTP buttons.
	SignatureHash param.Opt[string] `json:"signature_hash,omitzero"`
	// Button label text. Maximum 25 characters. Required for URL, PHONE_NUMBER, and
	// QUICK_REPLY buttons. Not required for OTP buttons (Meta supplies the label).
	Text param.Opt[string] `json:"text,omitzero"`
	// URL for URL-type buttons. Supports one variable ({{1}}).
	URL param.Opt[string] `json:"url,omitzero"`
	// Whether zero-tap terms have been accepted.
	ZeroTapTermsAccepted param.Opt[bool] `json:"zero_tap_terms_accepted,omitzero"`
	// Sample values for URL variable.
	Example []string `json:"example,omitzero"`
	// Flow action type for FLOW-type buttons.
	//
	// Any of "navigate", "data_exchange".
	FlowAction string `json:"flow_action,omitzero"`
	// Any of "COPY_CODE", "ONE_TAP".
	OtpType string `json:"otp_type,omitzero"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentButtonsButton) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentButtonsButton
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentButtonsButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentButtonsButton](
		"type", "URL", "PHONE_NUMBER", "QUICK_REPLY", "OTP", "COPY_CODE", "FLOW",
	)
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentButtonsButton](
		"flow_action", "navigate", "data_exchange",
	)
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentButtonsButton](
		"otp_type", "COPY_CODE", "ONE_TAP",
	)
}

// Carousel component for multi-card templates. Each card can contain its own
// header, body, and buttons.
//
// The properties Cards, Type are required.
type WhatsappTemplateNewParamsComponentCarousel struct {
	// Array of card objects, each with its own components.
	Cards []WhatsappTemplateNewParamsComponentCarouselCard `json:"cards,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "CAROUSEL".
	Type constant.Carousel `json:"type" api:"required"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentCarousel) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentCarousel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentCarousel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappTemplateNewParamsComponentCarouselCard struct {
	Components []map[string]any `json:"components,omitzero"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentCarouselCard) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentCarouselCard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentCarouselCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappTemplateListParams struct {
	// Search templates by name
	FilterSearch param.Opt[string] `query:"filter[search],omitzero" json:"-"`
	// Filter by template status
	FilterStatus param.Opt[string] `query:"filter[status],omitzero" json:"-"`
	// Filter by WABA ID
	FilterWabaID param.Opt[string] `query:"filter[waba_id],omitzero" json:"-"`
	PageNumber   param.Opt[int64]  `query:"page[number],omitzero" json:"-"`
	PageSize     param.Opt[int64]  `query:"page[size],omitzero" json:"-"`
	// Filter by category
	//
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	FilterCategory WhatsappTemplateListParamsFilterCategory `query:"filter[category],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WhatsappTemplateListParams]'s query parameters as
// `url.Values`.
func (r WhatsappTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by category
type WhatsappTemplateListParamsFilterCategory string

const (
	WhatsappTemplateListParamsFilterCategoryMarketing      WhatsappTemplateListParamsFilterCategory = "MARKETING"
	WhatsappTemplateListParamsFilterCategoryUtility        WhatsappTemplateListParamsFilterCategory = "UTILITY"
	WhatsappTemplateListParamsFilterCategoryAuthentication WhatsappTemplateListParamsFilterCategory = "AUTHENTICATION"
)

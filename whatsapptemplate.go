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
	OfWhatsappTemplateNewsComponentWhatsappTemplateHeaderComponent   *WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponent   `json:",omitzero,inline"`
	OfWhatsappTemplateNewsComponentWhatsappTemplateBodyComponent     *WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponent     `json:",omitzero,inline"`
	OfWhatsappTemplateNewsComponentWhatsappTemplateFooterComponent   *WhatsappTemplateNewParamsComponentWhatsappTemplateFooterComponent   `json:",omitzero,inline"`
	OfWhatsappTemplateNewsComponentWhatsappTemplateButtonsComponent  *WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponent  `json:",omitzero,inline"`
	OfWhatsappTemplateNewsComponentWhatsappTemplateCarouselComponent *WhatsappTemplateNewParamsComponentWhatsappTemplateCarouselComponent `json:",omitzero,inline"`
	paramUnion
}

func (u WhatsappTemplateNewParamsComponentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWhatsappTemplateNewsComponentWhatsappTemplateHeaderComponent,
		u.OfWhatsappTemplateNewsComponentWhatsappTemplateBodyComponent,
		u.OfWhatsappTemplateNewsComponentWhatsappTemplateFooterComponent,
		u.OfWhatsappTemplateNewsComponentWhatsappTemplateButtonsComponent,
		u.OfWhatsappTemplateNewsComponentWhatsappTemplateCarouselComponent)
}
func (u *WhatsappTemplateNewParamsComponentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WhatsappTemplateNewParamsComponentUnion) asAny() any {
	if !param.IsOmitted(u.OfWhatsappTemplateNewsComponentWhatsappTemplateHeaderComponent) {
		return u.OfWhatsappTemplateNewsComponentWhatsappTemplateHeaderComponent
	} else if !param.IsOmitted(u.OfWhatsappTemplateNewsComponentWhatsappTemplateBodyComponent) {
		return u.OfWhatsappTemplateNewsComponentWhatsappTemplateBodyComponent
	} else if !param.IsOmitted(u.OfWhatsappTemplateNewsComponentWhatsappTemplateFooterComponent) {
		return u.OfWhatsappTemplateNewsComponentWhatsappTemplateFooterComponent
	} else if !param.IsOmitted(u.OfWhatsappTemplateNewsComponentWhatsappTemplateButtonsComponent) {
		return u.OfWhatsappTemplateNewsComponentWhatsappTemplateButtonsComponent
	} else if !param.IsOmitted(u.OfWhatsappTemplateNewsComponentWhatsappTemplateCarouselComponent) {
		return u.OfWhatsappTemplateNewsComponentWhatsappTemplateCarouselComponent
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappTemplateNewParamsComponentUnion) GetFormat() *string {
	if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateHeaderComponent; vt != nil {
		return &vt.Format
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappTemplateNewParamsComponentUnion) GetCodeExpirationMinutes() *int64 {
	if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateFooterComponent; vt != nil && vt.CodeExpirationMinutes.Valid() {
		return &vt.CodeExpirationMinutes.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappTemplateNewParamsComponentUnion) GetButtons() []WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponentButton {
	if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateButtonsComponent; vt != nil {
		return vt.Buttons
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappTemplateNewParamsComponentUnion) GetCards() []WhatsappTemplateNewParamsComponentWhatsappTemplateCarouselComponentCard {
	if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateCarouselComponent; vt != nil {
		return vt.Cards
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappTemplateNewParamsComponentUnion) GetType() *string {
	if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateHeaderComponent; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateBodyComponent; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateFooterComponent; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateButtonsComponent; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateCarouselComponent; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappTemplateNewParamsComponentUnion) GetText() *string {
	if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateHeaderComponent; vt != nil && vt.Text.Valid() {
		return &vt.Text.Value
	} else if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateBodyComponent; vt != nil && vt.Text.Valid() {
		return &vt.Text.Value
	} else if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateFooterComponent; vt != nil && vt.Text.Valid() {
		return &vt.Text.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WhatsappTemplateNewParamsComponentUnion) GetExample() (res whatsappTemplateNewParamsComponentUnionExample) {
	if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateHeaderComponent; vt != nil {
		res.any = &vt.Example
	} else if vt := u.OfWhatsappTemplateNewsComponentWhatsappTemplateBodyComponent; vt != nil {
		res.any = &vt.Example
	}
	return
}

// Can have the runtime types
// [*WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponentExample],
// [*WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponentExample]
type whatsappTemplateNewParamsComponentUnionExample struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *telnyx.WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponentExample:
//	case *telnyx.WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponentExample:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u whatsappTemplateNewParamsComponentUnionExample) AsAny() any { return u.any }

// Optional header displayed at the top of the message.
//
// The properties Format, Type are required.
type WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponent struct {
	// Header format type: TEXT (supports one variable), IMAGE, VIDEO, DOCUMENT, or
	// LOCATION.
	//
	// Any of "TEXT", "IMAGE", "VIDEO", "DOCUMENT", "LOCATION".
	Format string `json:"format,omitzero" api:"required"`
	// Any of "HEADER".
	Type string `json:"type,omitzero" api:"required"`
	// Header text. Required when format is TEXT. Supports one variable ({{1}}).
	// Variables cannot be at the start or end.
	Text param.Opt[string] `json:"text,omitzero"`
	// Sample values for header variables.
	Example WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponentExample `json:"example,omitzero"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponent) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponent](
		"format", "TEXT", "IMAGE", "VIDEO", "DOCUMENT", "LOCATION",
	)
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponent](
		"type", "HEADER",
	)
}

// Sample values for header variables.
type WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponentExample struct {
	// Media handle for IMAGE, VIDEO, or DOCUMENT headers.
	HeaderHandle []string `json:"header_handle,omitzero"`
	// Sample values for text header variables.
	HeaderText []string `json:"header_text,omitzero"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponentExample) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponentExample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentWhatsappTemplateHeaderComponentExample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The main text content of the message. Supports multiple variable parameters
// ({{1}}, {{2}}, etc.). Variables cannot be at the start or end. Maximum 1024
// characters.
//
// The property Type is required.
type WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponent struct {
	// Any of "BODY".
	Type string `json:"type,omitzero" api:"required"`
	// Body text content. Use {{1}}, {{2}}, etc. for variable placeholders. Required
	// for MARKETING and UTILITY templates. Optional for AUTHENTICATION templates where
	// Meta provides the built-in OTP body.
	Text param.Opt[string] `json:"text,omitzero"`
	// Sample values for body variables. Required when body text contains parameters.
	Example WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponentExample `json:"example,omitzero"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponent) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponent](
		"type", "BODY",
	)
}

// Sample values for body variables. Required when body text contains parameters.
type WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponentExample struct {
	// Array containing one array of sample values, one per variable in order.
	BodyText [][]string `json:"body_text,omitzero"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponentExample) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponentExample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentWhatsappTemplateBodyComponentExample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional footer displayed at the bottom of the message. Does not support
// variables.
//
// The property Type is required.
type WhatsappTemplateNewParamsComponentWhatsappTemplateFooterComponent struct {
	// Any of "FOOTER".
	Type string `json:"type,omitzero" api:"required"`
	// OTP code expiration time in minutes. Used in AUTHENTICATION template footers
	// instead of free-form text.
	CodeExpirationMinutes param.Opt[int64] `json:"code_expiration_minutes,omitzero"`
	// Footer text. Maximum 60 characters. For non-authentication templates.
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentWhatsappTemplateFooterComponent) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentWhatsappTemplateFooterComponent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentWhatsappTemplateFooterComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentWhatsappTemplateFooterComponent](
		"type", "FOOTER",
	)
}

// Optional interactive buttons. Maximum 3 buttons per template.
//
// The properties Buttons, Type are required.
type WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponent struct {
	// Array of button objects. Meta supports various combinations of button types.
	Buttons []WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponentButton `json:"buttons,omitzero" api:"required"`
	// Any of "BUTTONS".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponent) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponent](
		"type", "BUTTONS",
	)
}

// The property Type is required.
type WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponentButton struct {
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

func (r WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponentButton) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponentButton
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponentButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponentButton](
		"type", "URL", "PHONE_NUMBER", "QUICK_REPLY", "OTP", "COPY_CODE", "FLOW",
	)
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponentButton](
		"flow_action", "navigate", "data_exchange",
	)
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentWhatsappTemplateButtonsComponentButton](
		"otp_type", "COPY_CODE", "ONE_TAP",
	)
}

// Carousel component for multi-card templates. Each card can contain its own
// header, body, and buttons.
//
// The properties Cards, Type are required.
type WhatsappTemplateNewParamsComponentWhatsappTemplateCarouselComponent struct {
	// Array of card objects, each with its own components.
	Cards []WhatsappTemplateNewParamsComponentWhatsappTemplateCarouselComponentCard `json:"cards,omitzero" api:"required"`
	// Any of "CAROUSEL".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentWhatsappTemplateCarouselComponent) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentWhatsappTemplateCarouselComponent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentWhatsappTemplateCarouselComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappTemplateNewParamsComponentWhatsappTemplateCarouselComponent](
		"type", "CAROUSEL",
	)
}

type WhatsappTemplateNewParamsComponentWhatsappTemplateCarouselComponentCard struct {
	Components []map[string]any `json:"components,omitzero"`
	paramObj
}

func (r WhatsappTemplateNewParamsComponentWhatsappTemplateCarouselComponentCard) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParamsComponentWhatsappTemplateCarouselComponentCard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParamsComponentWhatsappTemplateCarouselComponentCard) UnmarshalJSON(data []byte) error {
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// Manage Whatsapp message templates
//
// WhatsappMessageTemplateService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappMessageTemplateService] method instead.
type WhatsappMessageTemplateService struct {
	Options []option.RequestOption
}

// NewWhatsappMessageTemplateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhatsappMessageTemplateService(opts ...option.RequestOption) (r WhatsappMessageTemplateService) {
	r = WhatsappMessageTemplateService{}
	r.Options = opts
	return
}

// Get a Whatsapp message template by ID
func (r *WhatsappMessageTemplateService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WhatsappMessageTemplateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp_message_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a Whatsapp message template
func (r *WhatsappMessageTemplateService) Update(ctx context.Context, id string, body WhatsappMessageTemplateUpdateParams, opts ...option.RequestOption) (res *WhatsappMessageTemplateUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp_message_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a Whatsapp message template
func (r *WhatsappMessageTemplateService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v2/whatsapp_message_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type WhatsappMessageTemplateGetResponse struct {
	Data shared.WhatsappTemplateData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappMessageTemplateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateUpdateResponse struct {
	Data shared.WhatsappTemplateData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappMessageTemplateUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateUpdateParams struct {
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	Category WhatsappMessageTemplateUpdateParamsCategory `json:"category,omitzero"`
	// Updated template components. Same structure as the create request.
	Components []WhatsappMessageTemplateUpdateParamsComponentUnion `json:"components,omitzero"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateUpdateParamsCategory string

const (
	WhatsappMessageTemplateUpdateParamsCategoryMarketing      WhatsappMessageTemplateUpdateParamsCategory = "MARKETING"
	WhatsappMessageTemplateUpdateParamsCategoryUtility        WhatsappMessageTemplateUpdateParamsCategory = "UTILITY"
	WhatsappMessageTemplateUpdateParamsCategoryAuthentication WhatsappMessageTemplateUpdateParamsCategory = "AUTHENTICATION"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WhatsappMessageTemplateUpdateParamsComponentUnion struct {
	OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateHeaderComponent   *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponent   `json:",omitzero,inline"`
	OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateBodyComponent     *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponent     `json:",omitzero,inline"`
	OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateFooterComponent   *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateFooterComponent   `json:",omitzero,inline"`
	OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateButtonsComponent  *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponent  `json:",omitzero,inline"`
	OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateCarouselComponent *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateCarouselComponent `json:",omitzero,inline"`
	paramUnion
}

func (u WhatsappMessageTemplateUpdateParamsComponentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateHeaderComponent,
		u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateBodyComponent,
		u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateFooterComponent,
		u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateButtonsComponent,
		u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateCarouselComponent)
}
func (u *WhatsappMessageTemplateUpdateParamsComponentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WhatsappMessageTemplateUpdateParamsComponentUnion) asAny() any {
	if !param.IsOmitted(u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateHeaderComponent) {
		return u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateHeaderComponent
	} else if !param.IsOmitted(u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateBodyComponent) {
		return u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateBodyComponent
	} else if !param.IsOmitted(u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateFooterComponent) {
		return u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateFooterComponent
	} else if !param.IsOmitted(u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateButtonsComponent) {
		return u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateButtonsComponent
	} else if !param.IsOmitted(u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateCarouselComponent) {
		return u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateCarouselComponent
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetFormat() *string {
	if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateHeaderComponent; vt != nil {
		return &vt.Format
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetCodeExpirationMinutes() *int64 {
	if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateFooterComponent; vt != nil && vt.CodeExpirationMinutes.Valid() {
		return &vt.CodeExpirationMinutes.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetButtons() []WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponentButton {
	if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateButtonsComponent; vt != nil {
		return vt.Buttons
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetCards() []WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateCarouselComponentCard {
	if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateCarouselComponent; vt != nil {
		return vt.Cards
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetType() *string {
	if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateHeaderComponent; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateBodyComponent; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateFooterComponent; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateButtonsComponent; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateCarouselComponent; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetText() *string {
	if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateHeaderComponent; vt != nil && vt.Text.Valid() {
		return &vt.Text.Value
	} else if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateBodyComponent; vt != nil && vt.Text.Valid() {
		return &vt.Text.Value
	} else if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateFooterComponent; vt != nil && vt.Text.Valid() {
		return &vt.Text.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetExample() (res whatsappMessageTemplateUpdateParamsComponentUnionExample) {
	if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateHeaderComponent; vt != nil {
		res.any = &vt.Example
	} else if vt := u.OfWhatsappMessageTemplateUpdatesComponentWhatsappTemplateBodyComponent; vt != nil {
		res.any = &vt.Example
	}
	return
}

// Can have the runtime types
// [*WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponentExample],
// [*WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponentExample]
type whatsappMessageTemplateUpdateParamsComponentUnionExample struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *telnyx.WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponentExample:
//	case *telnyx.WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponentExample:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u whatsappMessageTemplateUpdateParamsComponentUnionExample) AsAny() any { return u.any }

// Optional header displayed at the top of the message.
//
// The properties Format, Type are required.
type WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponent struct {
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
	Example WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponentExample `json:"example,omitzero"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponent) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponent](
		"format", "TEXT", "IMAGE", "VIDEO", "DOCUMENT", "LOCATION",
	)
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponent](
		"type", "HEADER",
	)
}

// Sample values for header variables.
type WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponentExample struct {
	// Media handle for IMAGE, VIDEO, or DOCUMENT headers.
	HeaderHandle []string `json:"header_handle,omitzero"`
	// Sample values for text header variables.
	HeaderText []string `json:"header_text,omitzero"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponentExample) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponentExample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateHeaderComponentExample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The main text content of the message. Supports multiple variable parameters
// ({{1}}, {{2}}, etc.). Variables cannot be at the start or end. Maximum 1024
// characters.
//
// The property Type is required.
type WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponent struct {
	// Any of "BODY".
	Type string `json:"type,omitzero" api:"required"`
	// Body text content. Use {{1}}, {{2}}, etc. for variable placeholders. Required
	// for MARKETING and UTILITY templates. Optional for AUTHENTICATION templates where
	// Meta provides the built-in OTP body.
	Text param.Opt[string] `json:"text,omitzero"`
	// Sample values for body variables. Required when body text contains parameters.
	Example WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponentExample `json:"example,omitzero"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponent) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponent](
		"type", "BODY",
	)
}

// Sample values for body variables. Required when body text contains parameters.
type WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponentExample struct {
	// Array containing one array of sample values, one per variable in order.
	BodyText [][]string `json:"body_text,omitzero"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponentExample) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponentExample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateBodyComponentExample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional footer displayed at the bottom of the message. Does not support
// variables.
//
// The property Type is required.
type WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateFooterComponent struct {
	// Any of "FOOTER".
	Type string `json:"type,omitzero" api:"required"`
	// OTP code expiration time in minutes. Used in AUTHENTICATION template footers
	// instead of free-form text.
	CodeExpirationMinutes param.Opt[int64] `json:"code_expiration_minutes,omitzero"`
	// Footer text. Maximum 60 characters. For non-authentication templates.
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateFooterComponent) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateFooterComponent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateFooterComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateFooterComponent](
		"type", "FOOTER",
	)
}

// Optional interactive buttons. Maximum 3 buttons per template.
//
// The properties Buttons, Type are required.
type WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponent struct {
	// Array of button objects. Meta supports various combinations of button types.
	Buttons []WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponentButton `json:"buttons,omitzero" api:"required"`
	// Any of "BUTTONS".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponent) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponent](
		"type", "BUTTONS",
	)
}

// The property Type is required.
type WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponentButton struct {
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

func (r WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponentButton) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponentButton
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponentButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponentButton](
		"type", "URL", "PHONE_NUMBER", "QUICK_REPLY", "OTP", "COPY_CODE", "FLOW",
	)
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponentButton](
		"flow_action", "navigate", "data_exchange",
	)
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateButtonsComponentButton](
		"otp_type", "COPY_CODE", "ONE_TAP",
	)
}

// Carousel component for multi-card templates. Each card can contain its own
// header, body, and buttons.
//
// The properties Cards, Type are required.
type WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateCarouselComponent struct {
	// Array of card objects, each with its own components.
	Cards []WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateCarouselComponentCard `json:"cards,omitzero" api:"required"`
	// Any of "CAROUSEL".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateCarouselComponent) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateCarouselComponent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateCarouselComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateCarouselComponent](
		"type", "CAROUSEL",
	)
}

type WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateCarouselComponentCard struct {
	Components []any `json:"components,omitzero"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateCarouselComponentCard) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateCarouselComponentCard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentWhatsappTemplateCarouselComponentCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

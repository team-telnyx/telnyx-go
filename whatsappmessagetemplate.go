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
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
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
	OfHeader   *WhatsappMessageTemplateUpdateParamsComponentHeader   `json:",omitzero,inline"`
	OfBody     *WhatsappMessageTemplateUpdateParamsComponentBody     `json:",omitzero,inline"`
	OfFooter   *WhatsappMessageTemplateUpdateParamsComponentFooter   `json:",omitzero,inline"`
	OfButtons  *WhatsappMessageTemplateUpdateParamsComponentButtons  `json:",omitzero,inline"`
	OfCarousel *WhatsappMessageTemplateUpdateParamsComponentCarousel `json:",omitzero,inline"`
	paramUnion
}

func (u WhatsappMessageTemplateUpdateParamsComponentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHeader,
		u.OfBody,
		u.OfFooter,
		u.OfButtons,
		u.OfCarousel)
}
func (u *WhatsappMessageTemplateUpdateParamsComponentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WhatsappMessageTemplateUpdateParamsComponentUnion) asAny() any {
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
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetFormat() *string {
	if vt := u.OfHeader; vt != nil {
		return &vt.Format
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetCodeExpirationMinutes() *int64 {
	if vt := u.OfFooter; vt != nil && vt.CodeExpirationMinutes.Valid() {
		return &vt.CodeExpirationMinutes.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetButtons() []WhatsappMessageTemplateUpdateParamsComponentButtonsButton {
	if vt := u.OfButtons; vt != nil {
		return vt.Buttons
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetCards() []WhatsappMessageTemplateUpdateParamsComponentCarouselCard {
	if vt := u.OfCarousel; vt != nil {
		return vt.Cards
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetType() *string {
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
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetText() *string {
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
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetExample() (res whatsappMessageTemplateUpdateParamsComponentUnionExample) {
	if vt := u.OfHeader; vt != nil {
		res.any = &vt.Example
	} else if vt := u.OfBody; vt != nil {
		res.any = &vt.Example
	}
	return
}

// Can have the runtime types
// [*WhatsappMessageTemplateUpdateParamsComponentHeaderExample],
// [*WhatsappMessageTemplateUpdateParamsComponentBodyExample]
type whatsappMessageTemplateUpdateParamsComponentUnionExample struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *telnyx.WhatsappMessageTemplateUpdateParamsComponentHeaderExample:
//	case *telnyx.WhatsappMessageTemplateUpdateParamsComponentBodyExample:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u whatsappMessageTemplateUpdateParamsComponentUnionExample) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[WhatsappMessageTemplateUpdateParamsComponentUnion](
		"type",
		apijson.Discriminator[WhatsappMessageTemplateUpdateParamsComponentHeader]("HEADER"),
		apijson.Discriminator[WhatsappMessageTemplateUpdateParamsComponentBody]("BODY"),
		apijson.Discriminator[WhatsappMessageTemplateUpdateParamsComponentFooter]("FOOTER"),
		apijson.Discriminator[WhatsappMessageTemplateUpdateParamsComponentButtons]("BUTTONS"),
		apijson.Discriminator[WhatsappMessageTemplateUpdateParamsComponentCarousel]("CAROUSEL"),
	)
}

// Optional header displayed at the top of the message.
//
// The properties Format, Type are required.
type WhatsappMessageTemplateUpdateParamsComponentHeader struct {
	// Header format type: TEXT (supports one variable), IMAGE, VIDEO, DOCUMENT, or
	// LOCATION.
	//
	// Any of "TEXT", "IMAGE", "VIDEO", "DOCUMENT", "LOCATION".
	Format string `json:"format,omitzero" api:"required"`
	// Header text. Required when format is TEXT. Supports one variable ({{1}}).
	// Variables cannot be at the start or end.
	Text param.Opt[string] `json:"text,omitzero"`
	// Sample values for header variables.
	Example WhatsappMessageTemplateUpdateParamsComponentHeaderExample `json:"example,omitzero"`
	// This field can be elided, and will marshal its zero value as "HEADER".
	Type constant.Header `json:"type" api:"required"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentHeader) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentHeader](
		"format", "TEXT", "IMAGE", "VIDEO", "DOCUMENT", "LOCATION",
	)
}

// Sample values for header variables.
type WhatsappMessageTemplateUpdateParamsComponentHeaderExample struct {
	// Media handle for IMAGE, VIDEO, or DOCUMENT headers.
	HeaderHandle []string `json:"header_handle,omitzero"`
	// Sample values for text header variables.
	HeaderText []string `json:"header_text,omitzero"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentHeaderExample) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentHeaderExample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentHeaderExample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The main text content of the message. Supports multiple variable parameters
// ({{1}}, {{2}}, etc.). Variables cannot be at the start or end. Maximum 1024
// characters.
//
// The property Type is required.
type WhatsappMessageTemplateUpdateParamsComponentBody struct {
	// Body text content. Use {{1}}, {{2}}, etc. for variable placeholders. Required
	// for MARKETING and UTILITY templates. Optional for AUTHENTICATION templates where
	// Meta provides the built-in OTP body.
	Text param.Opt[string] `json:"text,omitzero"`
	// Sample values for body variables. Required when body text contains parameters.
	Example WhatsappMessageTemplateUpdateParamsComponentBodyExample `json:"example,omitzero"`
	// This field can be elided, and will marshal its zero value as "BODY".
	Type constant.Body `json:"type" api:"required"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentBody) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sample values for body variables. Required when body text contains parameters.
type WhatsappMessageTemplateUpdateParamsComponentBodyExample struct {
	// Array containing one array of sample values, one per variable in order.
	BodyText [][]string `json:"body_text,omitzero"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentBodyExample) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentBodyExample
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentBodyExample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional footer displayed at the bottom of the message. Does not support
// variables.
//
// The property Type is required.
type WhatsappMessageTemplateUpdateParamsComponentFooter struct {
	// OTP code expiration time in minutes. Used in AUTHENTICATION template footers
	// instead of free-form text.
	CodeExpirationMinutes param.Opt[int64] `json:"code_expiration_minutes,omitzero"`
	// Footer text. Maximum 60 characters. For non-authentication templates.
	Text param.Opt[string] `json:"text,omitzero"`
	// This field can be elided, and will marshal its zero value as "FOOTER".
	Type constant.Footer `json:"type" api:"required"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentFooter) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentFooter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentFooter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional interactive buttons. Maximum 3 buttons per template.
//
// The properties Buttons, Type are required.
type WhatsappMessageTemplateUpdateParamsComponentButtons struct {
	// Array of button objects. Meta supports various combinations of button types.
	Buttons []WhatsappMessageTemplateUpdateParamsComponentButtonsButton `json:"buttons,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "BUTTONS".
	Type constant.Buttons `json:"type" api:"required"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentButtons) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentButtons
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentButtons) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type WhatsappMessageTemplateUpdateParamsComponentButtonsButton struct {
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

func (r WhatsappMessageTemplateUpdateParamsComponentButtonsButton) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentButtonsButton
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentButtonsButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentButtonsButton](
		"type", "URL", "PHONE_NUMBER", "QUICK_REPLY", "OTP", "COPY_CODE", "FLOW",
	)
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentButtonsButton](
		"flow_action", "navigate", "data_exchange",
	)
	apijson.RegisterFieldValidator[WhatsappMessageTemplateUpdateParamsComponentButtonsButton](
		"otp_type", "COPY_CODE", "ONE_TAP",
	)
}

// Carousel component for multi-card templates. Each card can contain its own
// header, body, and buttons.
//
// The properties Cards, Type are required.
type WhatsappMessageTemplateUpdateParamsComponentCarousel struct {
	// Array of card objects, each with its own components.
	Cards []WhatsappMessageTemplateUpdateParamsComponentCarouselCard `json:"cards,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "CAROUSEL".
	Type constant.Carousel `json:"type" api:"required"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentCarousel) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentCarousel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentCarousel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateUpdateParamsComponentCarouselCard struct {
	Components []map[string]any `json:"components,omitzero"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParamsComponentCarouselCard) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParamsComponentCarouselCard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParamsComponentCarouselCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

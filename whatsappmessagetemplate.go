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
	OfHeader   *WhatsappTemplateHeaderComponentParam   `json:",omitzero,inline"`
	OfBody     *WhatsappTemplateBodyComponentParam     `json:",omitzero,inline"`
	OfFooter   *WhatsappTemplateFooterComponentParam   `json:",omitzero,inline"`
	OfButtons  *WhatsappTemplateButtonsComponentParam  `json:",omitzero,inline"`
	OfCarousel *WhatsappTemplateCarouselComponentParam `json:",omitzero,inline"`
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
		return (*string)(&vt.Format)
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
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetButtons() []WhatsappTemplateButtonsComponentButtonParam {
	if vt := u.OfButtons; vt != nil {
		return vt.Buttons
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WhatsappMessageTemplateUpdateParamsComponentUnion) GetCards() []WhatsappTemplateCarouselComponentCardParam {
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

// Can have the runtime types [*WhatsappTemplateHeaderComponentExampleParam],
// [*WhatsappTemplateBodyComponentExampleParam]
type whatsappMessageTemplateUpdateParamsComponentUnionExample struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *telnyx.WhatsappTemplateHeaderComponentExampleParam:
//	case *telnyx.WhatsappTemplateBodyComponentExampleParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u whatsappMessageTemplateUpdateParamsComponentUnionExample) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[WhatsappMessageTemplateUpdateParamsComponentUnion](
		"type",
		apijson.Discriminator[WhatsappTemplateHeaderComponentParam]("HEADER"),
		apijson.Discriminator[WhatsappTemplateBodyComponentParam]("BODY"),
		apijson.Discriminator[WhatsappTemplateFooterComponentParam]("FOOTER"),
		apijson.Discriminator[WhatsappTemplateButtonsComponentParam]("BUTTONS"),
		apijson.Discriminator[WhatsappTemplateCarouselComponentParam]("CAROUSEL"),
	)
}

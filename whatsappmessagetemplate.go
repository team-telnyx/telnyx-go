// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
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
	Data WhatsappMessageTemplateGetResponseData `json:"data"`
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

type WhatsappMessageTemplateGetResponseData struct {
	ID string `json:"id"`
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	Category string `json:"category"`
	// Whatsapp template components (header, body, footer, buttons)
	Components              []any                                                         `json:"components"`
	CreatedAt               time.Time                                                     `json:"created_at" format:"date-time"`
	Language                string                                                        `json:"language"`
	Name                    string                                                        `json:"name"`
	RecordType              string                                                        `json:"record_type"`
	RejectionReason         string                                                        `json:"rejection_reason"`
	Status                  string                                                        `json:"status"`
	TemplateID              string                                                        `json:"template_id"`
	UpdatedAt               time.Time                                                     `json:"updated_at" format:"date-time"`
	WhatsappBusinessAccount WhatsappMessageTemplateGetResponseDataWhatsappBusinessAccount `json:"whatsapp_business_account"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Category                respjson.Field
		Components              respjson.Field
		CreatedAt               respjson.Field
		Language                respjson.Field
		Name                    respjson.Field
		RecordType              respjson.Field
		RejectionReason         respjson.Field
		Status                  respjson.Field
		TemplateID              respjson.Field
		UpdatedAt               respjson.Field
		WhatsappBusinessAccount respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappMessageTemplateGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateGetResponseDataWhatsappBusinessAccount struct {
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateGetResponseDataWhatsappBusinessAccount) RawJSON() string {
	return r.JSON.raw
}
func (r *WhatsappMessageTemplateGetResponseDataWhatsappBusinessAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateUpdateResponse struct {
	Data WhatsappMessageTemplateUpdateResponseData `json:"data"`
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

type WhatsappMessageTemplateUpdateResponseData struct {
	ID string `json:"id"`
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	Category string `json:"category"`
	// Whatsapp template components (header, body, footer, buttons)
	Components              []any                                                            `json:"components"`
	CreatedAt               time.Time                                                        `json:"created_at" format:"date-time"`
	Language                string                                                           `json:"language"`
	Name                    string                                                           `json:"name"`
	RecordType              string                                                           `json:"record_type"`
	RejectionReason         string                                                           `json:"rejection_reason"`
	Status                  string                                                           `json:"status"`
	TemplateID              string                                                           `json:"template_id"`
	UpdatedAt               time.Time                                                        `json:"updated_at" format:"date-time"`
	WhatsappBusinessAccount WhatsappMessageTemplateUpdateResponseDataWhatsappBusinessAccount `json:"whatsapp_business_account"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Category                respjson.Field
		Components              respjson.Field
		CreatedAt               respjson.Field
		Language                respjson.Field
		Name                    respjson.Field
		RecordType              respjson.Field
		RejectionReason         respjson.Field
		Status                  respjson.Field
		TemplateID              respjson.Field
		UpdatedAt               respjson.Field
		WhatsappBusinessAccount respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappMessageTemplateUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateUpdateResponseDataWhatsappBusinessAccount struct {
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateUpdateResponseDataWhatsappBusinessAccount) RawJSON() string {
	return r.JSON.raw
}
func (r *WhatsappMessageTemplateUpdateResponseDataWhatsappBusinessAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateUpdateParams struct {
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	Category   WhatsappMessageTemplateUpdateParamsCategory `json:"category,omitzero"`
	Components []any                                       `json:"components,omitzero"`
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

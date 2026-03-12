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
	Category   WhatsappMessageTemplateUpdateParamsCategory `json:"category,omitzero"`
	Components []map[string]any                            `json:"components,omitzero"`
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

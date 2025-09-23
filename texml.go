// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// TexmlService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlService] method instead.
type TexmlService struct {
	Options  []option.RequestOption
	Accounts TexmlAccountService
	Calls    TexmlCallService
}

// NewTexmlService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTexmlService(opts ...option.RequestOption) (r TexmlService) {
	r = TexmlService{}
	r.Options = opts
	r.Accounts = NewTexmlAccountService(opts...)
	r.Calls = NewTexmlCallService(opts...)
	return
}

// Create a TeXML secret which can be later used as a Dynamic Parameter for TeXML
// when using Mustache Templates in your TeXML. In your TeXML you will be able to
// use your secret name, and this name will be replaced by the actual secret value
// when processing the TeXML on Telnyx side. The secrets are not visible in any
// logs.
func (r *TexmlService) Secrets(ctx context.Context, body TexmlSecretsParams, opts ...option.RequestOption) (res *TexmlSecretsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "texml/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TexmlSecretsResponse struct {
	Data TexmlSecretsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlSecretsResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlSecretsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlSecretsResponseData struct {
	Name string `json:"name"`
	// Any of "REDACTED".
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlSecretsResponseData) RawJSON() string { return r.JSON.raw }
func (r *TexmlSecretsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlSecretsParams struct {
	// Name used as a reference for the secret, if the name already exists within the
	// account its value will be replaced
	Name string `json:"name,required"`
	// Secret value which will be used when rendering the TeXML template
	Value string `json:"value,required"`
	paramObj
}

func (r TexmlSecretsParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlSecretsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlSecretsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

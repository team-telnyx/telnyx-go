// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apiform"
	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared"
)

// MessagingHostedNumberOrderActionService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingHostedNumberOrderActionService] method instead.
type MessagingHostedNumberOrderActionService struct {
	Options []option.RequestOption
}

// NewMessagingHostedNumberOrderActionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessagingHostedNumberOrderActionService(opts ...option.RequestOption) (r MessagingHostedNumberOrderActionService) {
	r = MessagingHostedNumberOrderActionService{}
	r.Options = opts
	return
}

// Upload file required for a messaging hosted number order
func (r *MessagingHostedNumberOrderActionService) UploadFile(ctx context.Context, id string, body MessagingHostedNumberOrderActionUploadFileParams, opts ...option.RequestOption) (res *MessagingHostedNumberOrderActionUploadFileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_hosted_number_orders/%s/actions/file_upload", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MessagingHostedNumberOrderActionUploadFileResponse struct {
	Data shared.MessagingHostedNumberOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrderActionUploadFileResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberOrderActionUploadFileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderActionUploadFileParams struct {
	// Must be the last month's bill with proof of ownership of all of the numbers in
	// the order in PDF format.
	Bill io.Reader `json:"bill,omitzero" format:"binary"`
	// Must be a signed LOA for the numbers in the order in PDF format.
	Loa io.Reader `json:"loa,omitzero" format:"binary"`
	paramObj
}

func (r MessagingHostedNumberOrderActionUploadFileParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

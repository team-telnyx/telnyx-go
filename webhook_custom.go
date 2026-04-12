// Custom webhook service methods — persists across codegen.
//
// This file contains hand-written methods on WebhookService that provide
// webhook parsing and signature verification. These are separated from
// webhook.go to avoid merge conflicts when Stainless regenerates event types.
package telnyx

import (
	"errors"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/lib"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

// UnsafeUnwrap parses a webhook payload without verifying the signature.
// Use this only if you have already verified the signature separately.
func (r *WebhookService) UnsafeUnwrap(payload []byte, opts ...option.RequestOption) (*UnsafeUnwrapWebhookEventUnion, error) {
	res := &UnsafeUnwrapWebhookEventUnion{}
	err := res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Unwrap verifies the webhook signature and parses the payload.
// Returns an error if signature verification fails or the payload is invalid.
func (r *WebhookService) Unwrap(payload []byte, headers http.Header, opts ...option.RequestOption) (*UnwrapWebhookEventUnion, error) {
	opts = slices.Concat(r.Options, opts)
	cfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	key := cfg.PublicKey
	if key == "" {
		return nil, errors.New("The PublicKey option must be set in order to verify webhook headers")
	}

	// Verify the webhook signature using ED25519 (via lib package)
	if err := lib.VerifyWebhookSignature(payload, headers, key); err != nil {
		return nil, err
	}

	res := &UnwrapWebhookEventUnion{}
	err = res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Verify checks the webhook signature without parsing the payload.
// Returns nil if the signature is valid, or an error describing why verification failed.
func (r *WebhookService) Verify(payload []byte, headers http.Header, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	cfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	key := cfg.PublicKey
	if key == "" {
		return errors.New("The PublicKey option must be set in order to verify webhook headers")
	}

	return lib.VerifyWebhookSignature(payload, headers, key)
}

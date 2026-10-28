// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// Purchase and manage dedicated sending IP addresses for your Telnyx email
// traffic.
//
// EmailDedicatedIPActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailDedicatedIPActionService] method instead.
type EmailDedicatedIPActionService struct {
	Options []option.RequestOption
}

// NewEmailDedicatedIPActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailDedicatedIPActionService(opts ...option.RequestOption) (r EmailDedicatedIPActionService) {
	r = EmailDedicatedIPActionService{}
	r.Options = opts
	return
}

// Manually complete the warm-up ramp (`warming_up` to `active`), bypassing the
// schedule. Use only for an IP already warmed elsewhere. Completion sets
// `warmup.progress` to `100`, clears `warmup.daily_cap`, and applies the
// ramp-complete allowance to the current UTC day: `warmup.today.cap` becomes
// `100000000`, `warmup.today.granted` keeps the sends already released today, and
// `warmup.today.remaining` becomes `cap` minus `granted`. Once completion applies
// this allowance, an exhausted stage allowance does not block the next eligible
// first-attempt send. The response reports these values. On an `active` IP,
// `warmup.progress` and `warmup.daily_cap` are informational: they can keep
// earlier values after a suspension is lifted and do not limit sending.
func (r *EmailDedicatedIPActionService) CompleteWarmup(ctx context.Context, id string, opts ...option.RequestOption) (res *DedicatedIPShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email/dedicated_ips/%s/actions/complete_warmup", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Pause the warm-up ramp at its current stage. While paused, the IP releases no
// first-attempt sends from hold and the ramp does not advance. Only a `warming_up`
// IP can be paused; pausing an already paused IP refreshes `paused_at`.
func (r *EmailDedicatedIPActionService) PauseWarmup(ctx context.Context, id string, opts ...option.RequestOption) (res *DedicatedIPShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email/dedicated_ips/%s/actions/pause_warmup", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Resume a paused warm-up ramp from its held stage. Only a `warming_up` IP can be
// resumed; resuming a ramp that is not paused returns the IP unchanged.
func (r *EmailDedicatedIPActionService) ResumeWarmup(ctx context.Context, id string, opts ...option.RequestOption) (res *DedicatedIPShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email/dedicated_ips/%s/actions/resume_warmup", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Start the warm-up ramp (`provisioning` to `warming_up`). The ramp never starts
// automatically. Only a `provisioning` IP can be started; a `suspended` IP returns
// `409`. Starting does not check DNS: the IP carries mail only while
// `dns_confirmed_at` is set and the IP is in a pool that routes mail.
func (r *EmailDedicatedIPActionService) StartWarmup(ctx context.Context, id string, opts ...option.RequestOption) (res *DedicatedIPShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email/dedicated_ips/%s/actions/start_warmup", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

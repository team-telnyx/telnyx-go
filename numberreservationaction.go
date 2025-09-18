// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// NumberReservationActionService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumberReservationActionService] method instead.
type NumberReservationActionService struct {
	Options []option.RequestOption
}

// NewNumberReservationActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNumberReservationActionService(opts ...option.RequestOption) (r NumberReservationActionService) {
	r = NumberReservationActionService{}
	r.Options = opts
	return
}

// Extends reservation expiry time on all phone numbers.
func (r *NumberReservationActionService) Extend(ctx context.Context, numberReservationID string, opts ...option.RequestOption) (res *NumberReservationActionExtendResponse, err error) {
	opts = append(r.Options[:], opts...)
	if numberReservationID == "" {
		err = errors.New("missing required number_reservation_id parameter")
		return
	}
	path := fmt.Sprintf("number_reservations/%s/actions/extend", numberReservationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type NumberReservationActionExtendResponse struct {
	Data NumberReservation `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberReservationActionExtendResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberReservationActionExtendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

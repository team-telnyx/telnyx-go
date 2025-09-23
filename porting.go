// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// PortingService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingService] method instead.
type PortingService struct {
	Options           []option.RequestOption
	Events            PortingEventService
	Reports           PortingReportService
	LoaConfigurations PortingLoaConfigurationService
}

// NewPortingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPortingService(opts ...option.RequestOption) (r PortingService) {
	r = PortingService{}
	r.Options = opts
	r.Events = NewPortingEventService(opts...)
	r.Reports = NewPortingReportService(opts...)
	r.LoaConfigurations = NewPortingLoaConfigurationService(opts...)
	return
}

// List available carriers in the UK.
func (r *PortingService) ListUkCarriers(ctx context.Context, opts ...option.RequestOption) (res *PortingListUkCarriersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "porting/uk_carriers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PortingListUkCarriersResponse struct {
	Data []PortingListUkCarriersResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingListUkCarriersResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingListUkCarriersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingListUkCarriersResponseData struct {
	// Identifies the UK carrier.
	ID string `json:"id" format:"uuid"`
	// Alternative CUPIDs of the carrier.
	AlternativeCupids []string `json:"alternative_cupids"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The CUPID of the carrier. This is a 3 digit number code that identifies the
	// carrier in the UK.
	Cupid string `json:"cupid"`
	// The name of the carrier.
	Name string `json:"name"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AlternativeCupids respjson.Field
		CreatedAt         respjson.Field
		Cupid             respjson.Field
		Name              respjson.Field
		RecordType        respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingListUkCarriersResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingListUkCarriersResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

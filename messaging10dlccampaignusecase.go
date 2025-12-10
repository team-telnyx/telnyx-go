// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// Messaging10dlcCampaignUsecaseService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessaging10dlcCampaignUsecaseService] method instead.
type Messaging10dlcCampaignUsecaseService struct {
	Options []option.RequestOption
}

// NewMessaging10dlcCampaignUsecaseService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessaging10dlcCampaignUsecaseService(opts ...option.RequestOption) (r Messaging10dlcCampaignUsecaseService) {
	r = Messaging10dlcCampaignUsecaseService{}
	r.Options = opts
	return
}

// Get Campaign Cost
func (r *Messaging10dlcCampaignUsecaseService) GetCost(ctx context.Context, query Messaging10dlcCampaignUsecaseGetCostParams, opts ...option.RequestOption) (res *Messaging10dlcCampaignUsecaseGetCostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "10dlc/campaign/usecase/cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Messaging10dlcCampaignUsecaseGetCostResponse struct {
	CampaignUsecase string `json:"campaignUsecase,required"`
	Description     string `json:"description,required"`
	MonthlyCost     string `json:"monthlyCost,required"`
	UpFrontCost     string `json:"upFrontCost,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CampaignUsecase respjson.Field
		Description     respjson.Field
		MonthlyCost     respjson.Field
		UpFrontCost     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Messaging10dlcCampaignUsecaseGetCostResponse) RawJSON() string { return r.JSON.raw }
func (r *Messaging10dlcCampaignUsecaseGetCostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Messaging10dlcCampaignUsecaseGetCostParams struct {
	Usecase string `query:"usecase,required" json:"-"`
	paramObj
}

// URLQuery serializes [Messaging10dlcCampaignUsecaseGetCostParams]'s query
// parameters as `url.Values`.
func (r Messaging10dlcCampaignUsecaseGetCostParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

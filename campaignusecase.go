// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// CampaignUsecaseService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignUsecaseService] method instead.
type CampaignUsecaseService struct {
	Options []option.RequestOption
}

// NewCampaignUsecaseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignUsecaseService(opts ...option.RequestOption) (r CampaignUsecaseService) {
	r = CampaignUsecaseService{}
	r.Options = opts
	return
}

// Get Campaign Cost
func (r *CampaignUsecaseService) GetCost(ctx context.Context, query CampaignUsecaseGetCostParams, opts ...option.RequestOption) (res *CampaignUsecaseGetCostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "campaign/usecase/cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CampaignUsecaseGetCostResponse struct {
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
func (r CampaignUsecaseGetCostResponse) RawJSON() string { return r.JSON.raw }
func (r *CampaignUsecaseGetCostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignUsecaseGetCostParams struct {
	Usecase string `query:"usecase,required" json:"-"`
	paramObj
}

// URLQuery serializes [CampaignUsecaseGetCostParams]'s query parameters as
// `url.Values`.
func (r CampaignUsecaseGetCostParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

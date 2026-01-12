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
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Messaging10dlcCampaignBuilderBrandService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessaging10dlcCampaignBuilderBrandService] method instead.
type Messaging10dlcCampaignBuilderBrandService struct {
	Options []option.RequestOption
}

// NewMessaging10dlcCampaignBuilderBrandService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewMessaging10dlcCampaignBuilderBrandService(opts ...option.RequestOption) (r Messaging10dlcCampaignBuilderBrandService) {
	r = Messaging10dlcCampaignBuilderBrandService{}
	r.Options = opts
	return
}

// This endpoint allows you to see whether or not the supplied brand is suitable
// for your desired campaign use case.
func (r *Messaging10dlcCampaignBuilderBrandService) QualifyByUsecase(ctx context.Context, usecase string, query Messaging10dlcCampaignBuilderBrandQualifyByUsecaseParams, opts ...option.RequestOption) (res *Messaging10dlcCampaignBuilderBrandQualifyByUsecaseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.BrandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	if usecase == "" {
		err = errors.New("missing required usecase parameter")
		return
	}
	path := fmt.Sprintf("10dlc/campaignBuilder/brand/%s/usecase/%s", query.BrandID, usecase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Messaging10dlcCampaignBuilderBrandQualifyByUsecaseResponse struct {
	// Campaign annual subscription fee
	AnnualFee float64 `json:"annualFee"`
	// Maximum number of sub-usecases declaration required.
	MaxSubUsecases int64 `json:"maxSubUsecases"`
	// Minimum number of sub-usecases declaration required.
	MinSubUsecases int64 `json:"minSubUsecases"`
	// Map of usecase metadata for each MNO. Key is the network ID of the MNO (e.g.
	// 10017), Value is the mno metadata for the usecase.
	MnoMetadata map[string]any `json:"mnoMetadata"`
	// Campaign monthly subscription fee
	MonthlyFee float64 `json:"monthlyFee"`
	// Campaign quarterly subscription fee
	QuarterlyFee float64 `json:"quarterlyFee"`
	// Campaign usecase
	Usecase string `json:"usecase"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnnualFee      respjson.Field
		MaxSubUsecases respjson.Field
		MinSubUsecases respjson.Field
		MnoMetadata    respjson.Field
		MonthlyFee     respjson.Field
		QuarterlyFee   respjson.Field
		Usecase        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Messaging10dlcCampaignBuilderBrandQualifyByUsecaseResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *Messaging10dlcCampaignBuilderBrandQualifyByUsecaseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Messaging10dlcCampaignBuilderBrandQualifyByUsecaseParams struct {
	BrandID string `path:"brandId,required" json:"-"`
	paramObj
}

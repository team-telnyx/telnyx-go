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

// CampaignBuilderBrandService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignBuilderBrandService] method instead.
type CampaignBuilderBrandService struct {
	Options []option.RequestOption
}

// NewCampaignBuilderBrandService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCampaignBuilderBrandService(opts ...option.RequestOption) (r CampaignBuilderBrandService) {
	r = CampaignBuilderBrandService{}
	r.Options = opts
	return
}

// This endpoint allows you to see whether or not the supplied brand is suitable
// for your desired campaign use case.
func (r *CampaignBuilderBrandService) QualifyByUsecase(ctx context.Context, usecase string, query CampaignBuilderBrandQualifyByUsecaseParams, opts ...option.RequestOption) (res *CampaignBuilderBrandQualifyByUsecaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.BrandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	if usecase == "" {
		err = errors.New("missing required usecase parameter")
		return
	}
	path := fmt.Sprintf("campaignBuilder/brand/%s/usecase/%s", query.BrandID, usecase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CampaignBuilderBrandQualifyByUsecaseResponse struct {
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
func (r CampaignBuilderBrandQualifyByUsecaseResponse) RawJSON() string { return r.JSON.raw }
func (r *CampaignBuilderBrandQualifyByUsecaseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignBuilderBrandQualifyByUsecaseParams struct {
	BrandID string `path:"brandId,required" json:"-"`
	paramObj
}

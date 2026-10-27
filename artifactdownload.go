// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// Download screenshots and other artifacts.
//
// ArtifactDownloadService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArtifactDownloadService] method instead.
type ArtifactDownloadService struct {
	Options []option.RequestOption
}

// NewArtifactDownloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewArtifactDownloadService(opts ...option.RequestOption) (r ArtifactDownloadService) {
	r = ArtifactDownloadService{}
	r.Options = opts
	return
}

// Downloads the artifact content. The signed URL (with expires and sig query
// params) is obtained from the artifact metadata endpoint.
func (r *ArtifactDownloadService) Get(ctx context.Context, artifactID string, query ArtifactDownloadGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if artifactID == "" {
		err = errors.New("missing required artifactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("artifact_downloads/%s", url.PathEscape(artifactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ArtifactDownloadGetParams struct {
	Expires string `query:"expires" api:"required" json:"-"`
	Sig     string `query:"sig" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [ArtifactDownloadGetParams]'s query parameters as
// `url.Values`.
func (r ArtifactDownloadGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

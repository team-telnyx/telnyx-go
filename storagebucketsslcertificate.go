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
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apiform"
	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// StorageBucketSslCertificateService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageBucketSslCertificateService] method instead.
type StorageBucketSslCertificateService struct {
	Options []option.RequestOption
}

// NewStorageBucketSslCertificateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewStorageBucketSslCertificateService(opts ...option.RequestOption) (r StorageBucketSslCertificateService) {
	r = StorageBucketSslCertificateService{}
	r.Options = opts
	return
}

// Uploads an SSL certificate and its matching secret so that you can use Telnyx's
// storage as your CDN.
func (r *StorageBucketSslCertificateService) New(ctx context.Context, bucketName string, body StorageBucketSslCertificateNewParams, opts ...option.RequestOption) (res *StorageBucketSslCertificateNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bucketName == "" {
		err = errors.New("missing required bucketName parameter")
		return
	}
	path := fmt.Sprintf("storage/buckets/%s/ssl_certificate", bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns the stored certificate detail of a bucket, if applicable.
func (r *StorageBucketSslCertificateService) Get(ctx context.Context, bucketName string, opts ...option.RequestOption) (res *StorageBucketSslCertificateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bucketName == "" {
		err = errors.New("missing required bucketName parameter")
		return
	}
	path := fmt.Sprintf("storage/buckets/%s/ssl_certificate", bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an SSL certificate and its matching secret.
func (r *StorageBucketSslCertificateService) Delete(ctx context.Context, bucketName string, opts ...option.RequestOption) (res *StorageBucketSslCertificateDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bucketName == "" {
		err = errors.New("missing required bucketName parameter")
		return
	}
	path := fmt.Sprintf("storage/buckets/%s/ssl_certificate", bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type SslCertificate struct {
	// Unique identifier for the SSL certificate
	ID string `json:"id"`
	// Time when SSL certificate was uploaded
	CreatedAt time.Time              `json:"created_at" format:"date-time"`
	IssuedBy  SslCertificateIssuedBy `json:"issued_by"`
	IssuedTo  SslCertificateIssuedTo `json:"issued_to"`
	// The time the certificate is valid from
	ValidFrom time.Time `json:"valid_from" format:"date-time"`
	// The time the certificate is valid to
	ValidTo time.Time `json:"valid_to" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IssuedBy    respjson.Field
		IssuedTo    respjson.Field
		ValidFrom   respjson.Field
		ValidTo     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SslCertificate) RawJSON() string { return r.JSON.raw }
func (r *SslCertificate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SslCertificateIssuedBy struct {
	// The common name of the entity the certificate was issued by
	CommonName string `json:"common_name"`
	// The organization the certificate was issued by
	Organization string `json:"organization"`
	// The organizational unit the certificate was issued by
	OrganizationUnit string `json:"organization_unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommonName       respjson.Field
		Organization     respjson.Field
		OrganizationUnit respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SslCertificateIssuedBy) RawJSON() string { return r.JSON.raw }
func (r *SslCertificateIssuedBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SslCertificateIssuedTo struct {
	// The common name of the entity the certificate was issued to
	CommonName string `json:"common_name"`
	// The organization the certificate was issued to
	Organization string `json:"organization"`
	// The organizational unit the certificate was issued to
	OrganizationUnit string `json:"organization_unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommonName       respjson.Field
		Organization     respjson.Field
		OrganizationUnit respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SslCertificateIssuedTo) RawJSON() string { return r.JSON.raw }
func (r *SslCertificateIssuedTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketSslCertificateNewResponse struct {
	Data SslCertificate `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageBucketSslCertificateNewResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageBucketSslCertificateNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketSslCertificateGetResponse struct {
	Data SslCertificate `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageBucketSslCertificateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageBucketSslCertificateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketSslCertificateDeleteResponse struct {
	Data SslCertificate `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageBucketSslCertificateDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageBucketSslCertificateDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageBucketSslCertificateNewParams struct {
	// The SSL certificate file
	Certificate io.Reader `json:"certificate,omitzero" format:"binary"`
	// The private key file
	PrivateKey io.Reader `json:"private_key,omitzero" format:"binary"`
	paramObj
}

func (r StorageBucketSslCertificateNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

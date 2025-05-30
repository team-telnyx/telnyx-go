/*
Telnyx API

SIP trunking, SMS, MMS, Call Control and Telephony Data Services.

API version: 2.0.0
Contact: support@telnyx.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package telnyx

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// MDRDetailReportsAPIService MDRDetailReportsAPI service
type MDRDetailReportsAPIService service

type ApiGetPaginatedMdrsRequest struct {
	ctx context.Context
	ApiService *MDRDetailReportsAPIService
	startDate *string
	endDate *string
	id *string
	direction *string
	profile *string
	cld *string
	cli *string
	status *string
	messageType *string
}

// Pagination start date
func (r ApiGetPaginatedMdrsRequest) StartDate(startDate string) ApiGetPaginatedMdrsRequest {
	r.startDate = &startDate
	return r
}

// Pagination end date
func (r ApiGetPaginatedMdrsRequest) EndDate(endDate string) ApiGetPaginatedMdrsRequest {
	r.endDate = &endDate
	return r
}

func (r ApiGetPaginatedMdrsRequest) Id(id string) ApiGetPaginatedMdrsRequest {
	r.id = &id
	return r
}

func (r ApiGetPaginatedMdrsRequest) Direction(direction string) ApiGetPaginatedMdrsRequest {
	r.direction = &direction
	return r
}

func (r ApiGetPaginatedMdrsRequest) Profile(profile string) ApiGetPaginatedMdrsRequest {
	r.profile = &profile
	return r
}

func (r ApiGetPaginatedMdrsRequest) Cld(cld string) ApiGetPaginatedMdrsRequest {
	r.cld = &cld
	return r
}

func (r ApiGetPaginatedMdrsRequest) Cli(cli string) ApiGetPaginatedMdrsRequest {
	r.cli = &cli
	return r
}

func (r ApiGetPaginatedMdrsRequest) Status(status string) ApiGetPaginatedMdrsRequest {
	r.status = &status
	return r
}

func (r ApiGetPaginatedMdrsRequest) MessageType(messageType string) ApiGetPaginatedMdrsRequest {
	r.messageType = &messageType
	return r
}

func (r ApiGetPaginatedMdrsRequest) Execute() (*MdrGetDetailResponse, *http.Response, error) {
	return r.ApiService.GetPaginatedMdrsExecute(r)
}

/*
GetPaginatedMdrs Fetch all Mdr records

Fetch all Mdr records 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPaginatedMdrsRequest
*/
func (a *MDRDetailReportsAPIService) GetPaginatedMdrs(ctx context.Context) ApiGetPaginatedMdrsRequest {
	return ApiGetPaginatedMdrsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MdrGetDetailResponse
func (a *MDRDetailReportsAPIService) GetPaginatedMdrsExecute(r ApiGetPaginatedMdrsRequest) (*MdrGetDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MdrGetDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MDRDetailReportsAPIService.GetPaginatedMdrs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/mdrs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	}
	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "form", "")
	}
	if r.profile != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "profile", r.profile, "form", "")
	}
	if r.cld != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cld", r.cld, "form", "")
	}
	if r.cli != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cli", r.cli, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.messageType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "message_type", r.messageType, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

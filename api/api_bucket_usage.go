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
	"strings"
	"time"
)


// BucketUsageAPIService BucketUsageAPI service
type BucketUsageAPIService service

type ApiGetBucketUsageRequest struct {
	ctx context.Context
	ApiService *BucketUsageAPIService
	bucketName string
}

func (r ApiGetBucketUsageRequest) Execute() (*GetBucketUsage200Response, *http.Response, error) {
	return r.ApiService.GetBucketUsageExecute(r)
}

/*
GetBucketUsage Get Bucket Usage

Returns the amount of storage space and number of files a bucket takes up.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bucketName The name of the bucket
 @return ApiGetBucketUsageRequest
*/
func (a *BucketUsageAPIService) GetBucketUsage(ctx context.Context, bucketName string) ApiGetBucketUsageRequest {
	return ApiGetBucketUsageRequest{
		ApiService: a,
		ctx: ctx,
		bucketName: bucketName,
	}
}

// Execute executes the request
//  @return GetBucketUsage200Response
func (a *BucketUsageAPIService) GetBucketUsageExecute(r ApiGetBucketUsageRequest) (*GetBucketUsage200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetBucketUsage200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketUsageAPIService.GetBucketUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/storage/buckets/{bucketName}/usage/storage"
	localVarPath = strings.Replace(localVarPath, "{"+"bucketName"+"}", url.PathEscape(parameterValueToString(r.bucketName, "bucketName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetStorageAPIUsageRequest struct {
	ctx context.Context
	ApiService *BucketUsageAPIService
	bucketName string
	filterStartTime *time.Time
	filterEndTime *time.Time
}

// The start time of the period to filter the usage (ISO microsecond format)
func (r ApiGetStorageAPIUsageRequest) FilterStartTime(filterStartTime time.Time) ApiGetStorageAPIUsageRequest {
	r.filterStartTime = &filterStartTime
	return r
}

// The end time of the period to filter the usage (ISO microsecond format)
func (r ApiGetStorageAPIUsageRequest) FilterEndTime(filterEndTime time.Time) ApiGetStorageAPIUsageRequest {
	r.filterEndTime = &filterEndTime
	return r
}

func (r ApiGetStorageAPIUsageRequest) Execute() (*GetStorageAPIUsage200Response, *http.Response, error) {
	return r.ApiService.GetStorageAPIUsageExecute(r)
}

/*
GetStorageAPIUsage Get API Usage

Returns the detail on API usage on a bucket of a particular time period, group by method category.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bucketName The name of the bucket
 @return ApiGetStorageAPIUsageRequest
*/
func (a *BucketUsageAPIService) GetStorageAPIUsage(ctx context.Context, bucketName string) ApiGetStorageAPIUsageRequest {
	return ApiGetStorageAPIUsageRequest{
		ApiService: a,
		ctx: ctx,
		bucketName: bucketName,
	}
}

// Execute executes the request
//  @return GetStorageAPIUsage200Response
func (a *BucketUsageAPIService) GetStorageAPIUsageExecute(r ApiGetStorageAPIUsageRequest) (*GetStorageAPIUsage200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetStorageAPIUsage200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketUsageAPIService.GetStorageAPIUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/storage/buckets/{bucketName}/usage/api"
	localVarPath = strings.Replace(localVarPath, "{"+"bucketName"+"}", url.PathEscape(parameterValueToString(r.bucketName, "bucketName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.filterStartTime == nil {
		return localVarReturnValue, nil, reportError("filterStartTime is required and must be specified")
	}
	if r.filterEndTime == nil {
		return localVarReturnValue, nil, reportError("filterEndTime is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "filter[start_time]", r.filterStartTime, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "filter[end_time]", r.filterEndTime, "form", "")
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

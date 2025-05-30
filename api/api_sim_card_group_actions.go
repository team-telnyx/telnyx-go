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
)


// SIMCardGroupActionsAPIService SIMCardGroupActionsAPI service
type SIMCardGroupActionsAPIService service

type ApiGetSimCardGroupActionRequest struct {
	ctx context.Context
	ApiService *SIMCardGroupActionsAPIService
	id string
}

func (r ApiGetSimCardGroupActionRequest) Execute() (*GetSimCardGroupAction200Response, *http.Response, error) {
	return r.ApiService.GetSimCardGroupActionExecute(r)
}

/*
GetSimCardGroupAction Get SIM card group action details

This API allows fetching detailed information about a SIM card group action resource to make follow-ups in an existing asynchronous operation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifies the resource.
 @return ApiGetSimCardGroupActionRequest
*/
func (a *SIMCardGroupActionsAPIService) GetSimCardGroupAction(ctx context.Context, id string) ApiGetSimCardGroupActionRequest {
	return ApiGetSimCardGroupActionRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetSimCardGroupAction200Response
func (a *SIMCardGroupActionsAPIService) GetSimCardGroupActionExecute(r ApiGetSimCardGroupActionRequest) (*GetSimCardGroupAction200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSimCardGroupAction200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SIMCardGroupActionsAPIService.GetSimCardGroupAction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sim_card_group_actions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetSimCardGroupActionsRequest struct {
	ctx context.Context
	ApiService *SIMCardGroupActionsAPIService
	pageNumber *int32
	pageSize *int32
	filterSimCardGroupId *string
	filterStatus *string
	filterType *string
}

// The page number to load.
func (r ApiGetSimCardGroupActionsRequest) PageNumber(pageNumber int32) ApiGetSimCardGroupActionsRequest {
	r.pageNumber = &pageNumber
	return r
}

// The size of the page.
func (r ApiGetSimCardGroupActionsRequest) PageSize(pageSize int32) ApiGetSimCardGroupActionsRequest {
	r.pageSize = &pageSize
	return r
}

// A valid SIM card group ID.
func (r ApiGetSimCardGroupActionsRequest) FilterSimCardGroupId(filterSimCardGroupId string) ApiGetSimCardGroupActionsRequest {
	r.filterSimCardGroupId = &filterSimCardGroupId
	return r
}

// Filter by a specific status of the resource&#39;s lifecycle.
func (r ApiGetSimCardGroupActionsRequest) FilterStatus(filterStatus string) ApiGetSimCardGroupActionsRequest {
	r.filterStatus = &filterStatus
	return r
}

// Filter by action type.
func (r ApiGetSimCardGroupActionsRequest) FilterType(filterType string) ApiGetSimCardGroupActionsRequest {
	r.filterType = &filterType
	return r
}

func (r ApiGetSimCardGroupActionsRequest) Execute() (*GetSimCardGroupActions200Response, *http.Response, error) {
	return r.ApiService.GetSimCardGroupActionsExecute(r)
}

/*
GetSimCardGroupActions List SIM card group actions

This API allows listing a paginated collection a SIM card group actions. It allows to explore a collection of existing asynchronous operation using specific filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSimCardGroupActionsRequest
*/
func (a *SIMCardGroupActionsAPIService) GetSimCardGroupActions(ctx context.Context) ApiGetSimCardGroupActionsRequest {
	return ApiGetSimCardGroupActionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetSimCardGroupActions200Response
func (a *SIMCardGroupActionsAPIService) GetSimCardGroupActionsExecute(r ApiGetSimCardGroupActionsRequest) (*GetSimCardGroupActions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSimCardGroupActions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SIMCardGroupActionsAPIService.GetSimCardGroupActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sim_card_group_actions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[number]", r.pageNumber, "form", "")
	} else {
		var defaultValue int32 = 1
		r.pageNumber = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[size]", r.pageSize, "form", "")
	} else {
		var defaultValue int32 = 20
		r.pageSize = &defaultValue
	}
	if r.filterSimCardGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[sim_card_group_id]", r.filterSimCardGroupId, "form", "")
	}
	if r.filterStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[status]", r.filterStatus, "form", "")
	}
	if r.filterType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[type]", r.filterType, "form", "")
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
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

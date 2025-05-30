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


// SIMCardActionsAPIService SIMCardActionsAPI service
type SIMCardActionsAPIService service

type ApiGetBulkSimCardActionRequest struct {
	ctx context.Context
	ApiService *SIMCardActionsAPIService
	id string
}

func (r ApiGetBulkSimCardActionRequest) Execute() (*GetBulkSimCardAction200Response, *http.Response, error) {
	return r.ApiService.GetBulkSimCardActionExecute(r)
}

/*
GetBulkSimCardAction Get bulk SIM card action details

This API fetches information about a bulk SIM card action. A bulk SIM card action contains details about a collection of individual SIM card actions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifies the resource.
 @return ApiGetBulkSimCardActionRequest
*/
func (a *SIMCardActionsAPIService) GetBulkSimCardAction(ctx context.Context, id string) ApiGetBulkSimCardActionRequest {
	return ApiGetBulkSimCardActionRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetBulkSimCardAction200Response
func (a *SIMCardActionsAPIService) GetBulkSimCardActionExecute(r ApiGetBulkSimCardActionRequest) (*GetBulkSimCardAction200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetBulkSimCardAction200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SIMCardActionsAPIService.GetBulkSimCardAction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bulk_sim_card_actions/{id}"
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

type ApiGetSimCardActionRequest struct {
	ctx context.Context
	ApiService *SIMCardActionsAPIService
	id string
}

func (r ApiGetSimCardActionRequest) Execute() (*GetSimCardAction200Response, *http.Response, error) {
	return r.ApiService.GetSimCardActionExecute(r)
}

/*
GetSimCardAction Get SIM card action details

This API fetches detailed information about a SIM card action to follow-up on an existing asynchronous operation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Identifies the resource.
 @return ApiGetSimCardActionRequest
*/
func (a *SIMCardActionsAPIService) GetSimCardAction(ctx context.Context, id string) ApiGetSimCardActionRequest {
	return ApiGetSimCardActionRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetSimCardAction200Response
func (a *SIMCardActionsAPIService) GetSimCardActionExecute(r ApiGetSimCardActionRequest) (*GetSimCardAction200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSimCardAction200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SIMCardActionsAPIService.GetSimCardAction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sim_card_actions/{id}"
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

type ApiListBulkSimCardActionsRequest struct {
	ctx context.Context
	ApiService *SIMCardActionsAPIService
	pageNumber *int32
	pageSize *int32
	filterActionType *string
}

// The page number to load.
func (r ApiListBulkSimCardActionsRequest) PageNumber(pageNumber int32) ApiListBulkSimCardActionsRequest {
	r.pageNumber = &pageNumber
	return r
}

// The size of the page.
func (r ApiListBulkSimCardActionsRequest) PageSize(pageSize int32) ApiListBulkSimCardActionsRequest {
	r.pageSize = &pageSize
	return r
}

// Filter by action type.
func (r ApiListBulkSimCardActionsRequest) FilterActionType(filterActionType string) ApiListBulkSimCardActionsRequest {
	r.filterActionType = &filterActionType
	return r
}

func (r ApiListBulkSimCardActionsRequest) Execute() (*ListBulkSimCardActions200Response, *http.Response, error) {
	return r.ApiService.ListBulkSimCardActionsExecute(r)
}

/*
ListBulkSimCardActions List bulk SIM card actions

This API lists a paginated collection of bulk SIM card actions. A bulk SIM card action contains details about a collection of individual SIM card actions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListBulkSimCardActionsRequest
*/
func (a *SIMCardActionsAPIService) ListBulkSimCardActions(ctx context.Context) ApiListBulkSimCardActionsRequest {
	return ApiListBulkSimCardActionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListBulkSimCardActions200Response
func (a *SIMCardActionsAPIService) ListBulkSimCardActionsExecute(r ApiListBulkSimCardActionsRequest) (*ListBulkSimCardActions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListBulkSimCardActions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SIMCardActionsAPIService.ListBulkSimCardActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bulk_sim_card_actions"

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
	if r.filterActionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[action_type]", r.filterActionType, "form", "")
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

type ApiListSimCardActionsRequest struct {
	ctx context.Context
	ApiService *SIMCardActionsAPIService
	pageNumber *int32
	pageSize *int32
	filterSimCardId *string
	filterStatus *string
	filterBulkSimCardActionId *string
	filterActionType *string
}

// The page number to load.
func (r ApiListSimCardActionsRequest) PageNumber(pageNumber int32) ApiListSimCardActionsRequest {
	r.pageNumber = &pageNumber
	return r
}

// The size of the page.
func (r ApiListSimCardActionsRequest) PageSize(pageSize int32) ApiListSimCardActionsRequest {
	r.pageSize = &pageSize
	return r
}

// A valid SIM card ID.
func (r ApiListSimCardActionsRequest) FilterSimCardId(filterSimCardId string) ApiListSimCardActionsRequest {
	r.filterSimCardId = &filterSimCardId
	return r
}

// Filter by a specific status of the resource&#39;s lifecycle.
func (r ApiListSimCardActionsRequest) FilterStatus(filterStatus string) ApiListSimCardActionsRequest {
	r.filterStatus = &filterStatus
	return r
}

// Filter by a bulk SIM card action ID.
func (r ApiListSimCardActionsRequest) FilterBulkSimCardActionId(filterBulkSimCardActionId string) ApiListSimCardActionsRequest {
	r.filterBulkSimCardActionId = &filterBulkSimCardActionId
	return r
}

// Filter by action type.
func (r ApiListSimCardActionsRequest) FilterActionType(filterActionType string) ApiListSimCardActionsRequest {
	r.filterActionType = &filterActionType
	return r
}

func (r ApiListSimCardActionsRequest) Execute() (*ListSimCardActions200Response, *http.Response, error) {
	return r.ApiService.ListSimCardActionsExecute(r)
}

/*
ListSimCardActions List SIM card actions

This API lists a paginated collection of SIM card actions. It enables exploring a collection of existing asynchronous operations using specific filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListSimCardActionsRequest
*/
func (a *SIMCardActionsAPIService) ListSimCardActions(ctx context.Context) ApiListSimCardActionsRequest {
	return ApiListSimCardActionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListSimCardActions200Response
func (a *SIMCardActionsAPIService) ListSimCardActionsExecute(r ApiListSimCardActionsRequest) (*ListSimCardActions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListSimCardActions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SIMCardActionsAPIService.ListSimCardActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sim_card_actions"

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
	if r.filterSimCardId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[sim_card_id]", r.filterSimCardId, "form", "")
	}
	if r.filterStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[status]", r.filterStatus, "form", "")
	}
	if r.filterBulkSimCardActionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[bulk_sim_card_action_id]", r.filterBulkSimCardActionId, "form", "")
	}
	if r.filterActionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[action_type]", r.filterActionType, "form", "")
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

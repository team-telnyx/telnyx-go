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
	"reflect"
)


// UserBundlesAPIService UserBundlesAPI service
type UserBundlesAPIService service

type ApiCreateUserBundlesBulkRequest struct {
	ctx context.Context
	ApiService *UserBundlesAPIService
	createUserBundlesBulkRequest *CreateUserBundlesBulkRequest
	authorizationBearer *string
}

func (r ApiCreateUserBundlesBulkRequest) CreateUserBundlesBulkRequest(createUserBundlesBulkRequest CreateUserBundlesBulkRequest) ApiCreateUserBundlesBulkRequest {
	r.createUserBundlesBulkRequest = &createUserBundlesBulkRequest
	return r
}

// Format: Bearer &lt;TOKEN&gt;
func (r ApiCreateUserBundlesBulkRequest) AuthorizationBearer(authorizationBearer string) ApiCreateUserBundlesBulkRequest {
	r.authorizationBearer = &authorizationBearer
	return r
}

func (r ApiCreateUserBundlesBulkRequest) Execute() (*CreatedUserBundlesResponse, *http.Response, error) {
	return r.ApiService.CreateUserBundlesBulkExecute(r)
}

/*
CreateUserBundlesBulk Create User Bundles

Creates multiple user bundles for the user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateUserBundlesBulkRequest
*/
func (a *UserBundlesAPIService) CreateUserBundlesBulk(ctx context.Context) ApiCreateUserBundlesBulkRequest {
	return ApiCreateUserBundlesBulkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreatedUserBundlesResponse
func (a *UserBundlesAPIService) CreateUserBundlesBulkExecute(r ApiCreateUserBundlesBulkRequest) (*CreatedUserBundlesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreatedUserBundlesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserBundlesAPIService.CreateUserBundlesBulk")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundle_pricing/user_bundles/bulk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createUserBundlesBulkRequest == nil {
		return localVarReturnValue, nil, reportError("createUserBundlesBulkRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	if r.authorizationBearer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "authorization_bearer", r.authorizationBearer, "simple", "")
	}
	// body params
	localVarPostBody = r.createUserBundlesBulkRequest
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

type ApiDeactivateUserBundleRequest struct {
	ctx context.Context
	ApiService *UserBundlesAPIService
	userBundleId string
	authorizationBearer *string
}

// Format: Bearer &lt;TOKEN&gt;
func (r ApiDeactivateUserBundleRequest) AuthorizationBearer(authorizationBearer string) ApiDeactivateUserBundleRequest {
	r.authorizationBearer = &authorizationBearer
	return r
}

func (r ApiDeactivateUserBundleRequest) Execute() (*UserBundleCreateResponse, *http.Response, error) {
	return r.ApiService.DeactivateUserBundleExecute(r)
}

/*
DeactivateUserBundle Deactivate User Bundle

Deactivates a user bundle by its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userBundleId
 @return ApiDeactivateUserBundleRequest
*/
func (a *UserBundlesAPIService) DeactivateUserBundle(ctx context.Context, userBundleId string) ApiDeactivateUserBundleRequest {
	return ApiDeactivateUserBundleRequest{
		ApiService: a,
		ctx: ctx,
		userBundleId: userBundleId,
	}
}

// Execute executes the request
//  @return UserBundleCreateResponse
func (a *UserBundlesAPIService) DeactivateUserBundleExecute(r ApiDeactivateUserBundleRequest) (*UserBundleCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserBundleCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserBundlesAPIService.DeactivateUserBundle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundle_pricing/user_bundles/{user_bundle_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_bundle_id"+"}", url.PathEscape(parameterValueToString(r.userBundleId, "userBundleId")), -1)

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
	if r.authorizationBearer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "authorization_bearer", r.authorizationBearer, "simple", "")
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

type ApiGetUnusedUserBundlesRequest struct {
	ctx context.Context
	ApiService *UserBundlesAPIService
	filterCountryIso *[]string
	authorizationBearer *string
}

// Filter by country code.
func (r ApiGetUnusedUserBundlesRequest) FilterCountryIso(filterCountryIso []string) ApiGetUnusedUserBundlesRequest {
	r.filterCountryIso = &filterCountryIso
	return r
}

// Format: Bearer &lt;TOKEN&gt;
func (r ApiGetUnusedUserBundlesRequest) AuthorizationBearer(authorizationBearer string) ApiGetUnusedUserBundlesRequest {
	r.authorizationBearer = &authorizationBearer
	return r
}

func (r ApiGetUnusedUserBundlesRequest) Execute() (*UnusedUserBundlesResponse, *http.Response, error) {
	return r.ApiService.GetUnusedUserBundlesExecute(r)
}

/*
GetUnusedUserBundles Get Unused User Bundles

Returns all user bundles that aren't in use.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUnusedUserBundlesRequest
*/
func (a *UserBundlesAPIService) GetUnusedUserBundles(ctx context.Context) ApiGetUnusedUserBundlesRequest {
	return ApiGetUnusedUserBundlesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UnusedUserBundlesResponse
func (a *UserBundlesAPIService) GetUnusedUserBundlesExecute(r ApiGetUnusedUserBundlesRequest) (*UnusedUserBundlesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UnusedUserBundlesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserBundlesAPIService.GetUnusedUserBundles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundle_pricing/user_bundles/unused"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterCountryIso != nil {
		t := *r.filterCountryIso
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter[country_iso]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter[country_iso]", t, "form", "multi")
		}
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
	if r.authorizationBearer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "authorization_bearer", r.authorizationBearer, "simple", "")
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

type ApiGetUserBundleByIdRequest struct {
	ctx context.Context
	ApiService *UserBundlesAPIService
	userBundleId string
	authorizationBearer *string
}

// Format: Bearer &lt;TOKEN&gt;
func (r ApiGetUserBundleByIdRequest) AuthorizationBearer(authorizationBearer string) ApiGetUserBundleByIdRequest {
	r.authorizationBearer = &authorizationBearer
	return r
}

func (r ApiGetUserBundleByIdRequest) Execute() (*UserBundleResponse, *http.Response, error) {
	return r.ApiService.GetUserBundleByIdExecute(r)
}

/*
GetUserBundleById Get User Bundle by Id

Retrieves a user bundle by its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userBundleId
 @return ApiGetUserBundleByIdRequest
*/
func (a *UserBundlesAPIService) GetUserBundleById(ctx context.Context, userBundleId string) ApiGetUserBundleByIdRequest {
	return ApiGetUserBundleByIdRequest{
		ApiService: a,
		ctx: ctx,
		userBundleId: userBundleId,
	}
}

// Execute executes the request
//  @return UserBundleResponse
func (a *UserBundlesAPIService) GetUserBundleByIdExecute(r ApiGetUserBundleByIdRequest) (*UserBundleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserBundleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserBundlesAPIService.GetUserBundleById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundle_pricing/user_bundles/{user_bundle_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_bundle_id"+"}", url.PathEscape(parameterValueToString(r.userBundleId, "userBundleId")), -1)

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
	if r.authorizationBearer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "authorization_bearer", r.authorizationBearer, "simple", "")
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

type ApiGetUserBundleResourcesRequest struct {
	ctx context.Context
	ApiService *UserBundlesAPIService
	userBundleId string
	authorizationBearer *string
}

// Format: Bearer &lt;TOKEN&gt;
func (r ApiGetUserBundleResourcesRequest) AuthorizationBearer(authorizationBearer string) ApiGetUserBundleResourcesRequest {
	r.authorizationBearer = &authorizationBearer
	return r
}

func (r ApiGetUserBundleResourcesRequest) Execute() (*UserBundleResourcesResponse, *http.Response, error) {
	return r.ApiService.GetUserBundleResourcesExecute(r)
}

/*
GetUserBundleResources Get User Bundle Resources

Retrieves the resources of a user bundle by its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userBundleId
 @return ApiGetUserBundleResourcesRequest
*/
func (a *UserBundlesAPIService) GetUserBundleResources(ctx context.Context, userBundleId string) ApiGetUserBundleResourcesRequest {
	return ApiGetUserBundleResourcesRequest{
		ApiService: a,
		ctx: ctx,
		userBundleId: userBundleId,
	}
}

// Execute executes the request
//  @return UserBundleResourcesResponse
func (a *UserBundlesAPIService) GetUserBundleResourcesExecute(r ApiGetUserBundleResourcesRequest) (*UserBundleResourcesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserBundleResourcesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserBundlesAPIService.GetUserBundleResources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundle_pricing/user_bundles/{user_bundle_id}/resources"
	localVarPath = strings.Replace(localVarPath, "{"+"user_bundle_id"+"}", url.PathEscape(parameterValueToString(r.userBundleId, "userBundleId")), -1)

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
	if r.authorizationBearer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "authorization_bearer", r.authorizationBearer, "simple", "")
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

type ApiGetUserBundlesRequest struct {
	ctx context.Context
	ApiService *UserBundlesAPIService
	filterCountryIso *[]string
	filterResource *[]string
	pageNumber *int32
	pageSize *int32
	authorizationBearer *string
}

// Filter by country code.
func (r ApiGetUserBundlesRequest) FilterCountryIso(filterCountryIso []string) ApiGetUserBundlesRequest {
	r.filterCountryIso = &filterCountryIso
	return r
}

// Filter by resource.
func (r ApiGetUserBundlesRequest) FilterResource(filterResource []string) ApiGetUserBundlesRequest {
	r.filterResource = &filterResource
	return r
}

// The page number to load.
func (r ApiGetUserBundlesRequest) PageNumber(pageNumber int32) ApiGetUserBundlesRequest {
	r.pageNumber = &pageNumber
	return r
}

// The size of the page.
func (r ApiGetUserBundlesRequest) PageSize(pageSize int32) ApiGetUserBundlesRequest {
	r.pageSize = &pageSize
	return r
}

// Format: Bearer &lt;TOKEN&gt;
func (r ApiGetUserBundlesRequest) AuthorizationBearer(authorizationBearer string) ApiGetUserBundlesRequest {
	r.authorizationBearer = &authorizationBearer
	return r
}

func (r ApiGetUserBundlesRequest) Execute() (*PaginatedUserBundlesResponse, *http.Response, error) {
	return r.ApiService.GetUserBundlesExecute(r)
}

/*
GetUserBundles Get User Bundles

Get a paginated list of user bundles.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserBundlesRequest
*/
func (a *UserBundlesAPIService) GetUserBundles(ctx context.Context) ApiGetUserBundlesRequest {
	return ApiGetUserBundlesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedUserBundlesResponse
func (a *UserBundlesAPIService) GetUserBundlesExecute(r ApiGetUserBundlesRequest) (*PaginatedUserBundlesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedUserBundlesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserBundlesAPIService.GetUserBundles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundle_pricing/user_bundles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterCountryIso != nil {
		t := *r.filterCountryIso
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter[country_iso]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter[country_iso]", t, "form", "multi")
		}
	}
	if r.filterResource != nil {
		t := *r.filterResource
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filter[resource]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filter[resource]", t, "form", "multi")
		}
	}
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
	if r.authorizationBearer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "authorization_bearer", r.authorizationBearer, "simple", "")
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

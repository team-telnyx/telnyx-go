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


// NumbersFeaturesAPIService NumbersFeaturesAPI service
type NumbersFeaturesAPIService service

type ApiPostNumbersFeaturesRequest struct {
	ctx context.Context
	ApiService *NumbersFeaturesAPIService
	postNumbersFeaturesRequest *PostNumbersFeaturesRequest
}

func (r ApiPostNumbersFeaturesRequest) PostNumbersFeaturesRequest(postNumbersFeaturesRequest PostNumbersFeaturesRequest) ApiPostNumbersFeaturesRequest {
	r.postNumbersFeaturesRequest = &postNumbersFeaturesRequest
	return r
}

func (r ApiPostNumbersFeaturesRequest) Execute() (*PostNumbersFeatures200Response, *http.Response, error) {
	return r.ApiService.PostNumbersFeaturesExecute(r)
}

/*
PostNumbersFeatures Retrieve the features for a list of numbers

Retrieve the features for a list of numbers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostNumbersFeaturesRequest
*/
func (a *NumbersFeaturesAPIService) PostNumbersFeatures(ctx context.Context) ApiPostNumbersFeaturesRequest {
	return ApiPostNumbersFeaturesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostNumbersFeatures200Response
func (a *NumbersFeaturesAPIService) PostNumbersFeaturesExecute(r ApiPostNumbersFeaturesRequest) (*PostNumbersFeatures200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostNumbersFeatures200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NumbersFeaturesAPIService.PostNumbersFeatures")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/numbers_features"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postNumbersFeaturesRequest == nil {
		return localVarReturnValue, nil, reportError("postNumbersFeaturesRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.postNumbersFeaturesRequest
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

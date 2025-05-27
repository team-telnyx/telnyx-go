# WebhookToolParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the tool. | 
**Description** | **string** | The description of the tool. | 
**Url** | **string** | The URL of the external tool to be called. This URL is going to be used by the assistant. The URL can be templated like: &#x60;https://example.com/api/v1/{id}&#x60;, where &#x60;{id}&#x60; is a placeholder for a value that will be provided by the assistant if &#x60;path_parameters&#x60; are provided with the &#x60;id&#x60; attribute. | 
**Method** | Pointer to **string** | The HTTP method to be used when calling the external tool. | [optional] [default to "POST"]
**Headers** | Pointer to [**[]WebhookToolParamsHeadersInner**](WebhookToolParamsHeadersInner.md) | The headers to be sent to the external tool. | [optional] 
**BodyParameters** | Pointer to [**WebhookToolParamsBodyParameters**](WebhookToolParamsBodyParameters.md) |  | [optional] 
**PathParameters** | Pointer to [**WebhookToolParamsPathParameters**](WebhookToolParamsPathParameters.md) |  | [optional] 
**QueryParameters** | Pointer to [**WebhookToolParamsQueryParameters**](WebhookToolParamsQueryParameters.md) |  | [optional] 

## Methods

### NewWebhookToolParams

`func NewWebhookToolParams(name string, description string, url string, ) *WebhookToolParams`

NewWebhookToolParams instantiates a new WebhookToolParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookToolParamsWithDefaults

`func NewWebhookToolParamsWithDefaults() *WebhookToolParams`

NewWebhookToolParamsWithDefaults instantiates a new WebhookToolParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookToolParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookToolParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookToolParams) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WebhookToolParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookToolParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookToolParams) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUrl

`func (o *WebhookToolParams) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookToolParams) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookToolParams) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *WebhookToolParams) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WebhookToolParams) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WebhookToolParams) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *WebhookToolParams) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetHeaders

`func (o *WebhookToolParams) GetHeaders() []WebhookToolParamsHeadersInner`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookToolParams) GetHeadersOk() (*[]WebhookToolParamsHeadersInner, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookToolParams) SetHeaders(v []WebhookToolParamsHeadersInner)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebhookToolParams) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBodyParameters

`func (o *WebhookToolParams) GetBodyParameters() WebhookToolParamsBodyParameters`

GetBodyParameters returns the BodyParameters field if non-nil, zero value otherwise.

### GetBodyParametersOk

`func (o *WebhookToolParams) GetBodyParametersOk() (*WebhookToolParamsBodyParameters, bool)`

GetBodyParametersOk returns a tuple with the BodyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyParameters

`func (o *WebhookToolParams) SetBodyParameters(v WebhookToolParamsBodyParameters)`

SetBodyParameters sets BodyParameters field to given value.

### HasBodyParameters

`func (o *WebhookToolParams) HasBodyParameters() bool`

HasBodyParameters returns a boolean if a field has been set.

### GetPathParameters

`func (o *WebhookToolParams) GetPathParameters() WebhookToolParamsPathParameters`

GetPathParameters returns the PathParameters field if non-nil, zero value otherwise.

### GetPathParametersOk

`func (o *WebhookToolParams) GetPathParametersOk() (*WebhookToolParamsPathParameters, bool)`

GetPathParametersOk returns a tuple with the PathParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathParameters

`func (o *WebhookToolParams) SetPathParameters(v WebhookToolParamsPathParameters)`

SetPathParameters sets PathParameters field to given value.

### HasPathParameters

`func (o *WebhookToolParams) HasPathParameters() bool`

HasPathParameters returns a boolean if a field has been set.

### GetQueryParameters

`func (o *WebhookToolParams) GetQueryParameters() WebhookToolParamsQueryParameters`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *WebhookToolParams) GetQueryParametersOk() (*WebhookToolParamsQueryParameters, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *WebhookToolParams) SetQueryParameters(v WebhookToolParamsQueryParameters)`

SetQueryParameters sets QueryParameters field to given value.

### HasQueryParameters

`func (o *WebhookToolParams) HasQueryParameters() bool`

HasQueryParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



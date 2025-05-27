# UpdateCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The value to set the call status to. Setting the status to completed ends the call. | [optional] 
**Url** | Pointer to **string** | The URL where TeXML will make a request to retrieve a new set of TeXML instructions to continue the call flow. | [optional] 
**Method** | Pointer to **string** | HTTP request type used for &#x60;Url&#x60;. | [optional] 
**FallbackUrl** | Pointer to **string** | A failover URL for which Telnyx will retrieve the TeXML call instructions if the Url is not responding. | [optional] 
**FallbackMethod** | Pointer to **string** | HTTP request type used for &#x60;FallbackUrl&#x60;. | [optional] 
**StatusCallback** | Pointer to **string** | URL destination for Telnyx to send status callback events to for the call. | [optional] 
**StatusCallbackMethod** | Pointer to **string** | HTTP request type used for &#x60;StatusCallback&#x60;. | [optional] 
**Texml** | Pointer to **string** | TeXML to replace the current one with. | [optional] 

## Methods

### NewUpdateCallRequest

`func NewUpdateCallRequest() *UpdateCallRequest`

NewUpdateCallRequest instantiates a new UpdateCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCallRequestWithDefaults

`func NewUpdateCallRequestWithDefaults() *UpdateCallRequest`

NewUpdateCallRequestWithDefaults instantiates a new UpdateCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateCallRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateCallRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateCallRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateCallRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateCallRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateCallRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateCallRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateCallRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMethod

`func (o *UpdateCallRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *UpdateCallRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *UpdateCallRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *UpdateCallRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetFallbackUrl

`func (o *UpdateCallRequest) GetFallbackUrl() string`

GetFallbackUrl returns the FallbackUrl field if non-nil, zero value otherwise.

### GetFallbackUrlOk

`func (o *UpdateCallRequest) GetFallbackUrlOk() (*string, bool)`

GetFallbackUrlOk returns a tuple with the FallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUrl

`func (o *UpdateCallRequest) SetFallbackUrl(v string)`

SetFallbackUrl sets FallbackUrl field to given value.

### HasFallbackUrl

`func (o *UpdateCallRequest) HasFallbackUrl() bool`

HasFallbackUrl returns a boolean if a field has been set.

### GetFallbackMethod

`func (o *UpdateCallRequest) GetFallbackMethod() string`

GetFallbackMethod returns the FallbackMethod field if non-nil, zero value otherwise.

### GetFallbackMethodOk

`func (o *UpdateCallRequest) GetFallbackMethodOk() (*string, bool)`

GetFallbackMethodOk returns a tuple with the FallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackMethod

`func (o *UpdateCallRequest) SetFallbackMethod(v string)`

SetFallbackMethod sets FallbackMethod field to given value.

### HasFallbackMethod

`func (o *UpdateCallRequest) HasFallbackMethod() bool`

HasFallbackMethod returns a boolean if a field has been set.

### GetStatusCallback

`func (o *UpdateCallRequest) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *UpdateCallRequest) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *UpdateCallRequest) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *UpdateCallRequest) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### GetStatusCallbackMethod

`func (o *UpdateCallRequest) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *UpdateCallRequest) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *UpdateCallRequest) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *UpdateCallRequest) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### GetTexml

`func (o *UpdateCallRequest) GetTexml() string`

GetTexml returns the Texml field if non-nil, zero value otherwise.

### GetTexmlOk

`func (o *UpdateCallRequest) GetTexmlOk() (*string, bool)`

GetTexmlOk returns a tuple with the Texml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTexml

`func (o *UpdateCallRequest) SetTexml(v string)`

SetTexml sets Texml field to given value.

### HasTexml

`func (o *UpdateCallRequest) HasTexml() bool`

HasTexml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



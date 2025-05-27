# HttpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Headers** | Pointer to **[][]string** | List of headers, limited to 10kB. | [optional] 
**Body** | Pointer to **string** | Raw response body, limited to 10kB. | [optional] 

## Methods

### NewHttpResponse

`func NewHttpResponse() *HttpResponse`

NewHttpResponse instantiates a new HttpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpResponseWithDefaults

`func NewHttpResponseWithDefaults() *HttpResponse`

NewHttpResponseWithDefaults instantiates a new HttpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HttpResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HttpResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HttpResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HttpResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHeaders

`func (o *HttpResponse) GetHeaders() [][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpResponse) GetHeadersOk() (*[][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpResponse) SetHeaders(v [][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *HttpResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *HttpResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *HttpResponse) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *HttpResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



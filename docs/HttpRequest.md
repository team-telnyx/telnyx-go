# HttpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **[][]string** | List of headers, limited to 10kB. | [optional] 

## Methods

### NewHttpRequest

`func NewHttpRequest() *HttpRequest`

NewHttpRequest instantiates a new HttpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpRequestWithDefaults

`func NewHttpRequestWithDefaults() *HttpRequest`

NewHttpRequestWithDefaults instantiates a new HttpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *HttpRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HttpRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHeaders

`func (o *HttpRequest) GetHeaders() [][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpRequest) GetHeadersOk() (*[][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpRequest) SetHeaders(v [][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Http

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**HttpRequest**](HttpRequest.md) |  | [optional] 
**Response** | Pointer to [**NullableHttpResponse**](HttpResponse.md) |  | [optional] 

## Methods

### NewHttp

`func NewHttp() *Http`

NewHttp instantiates a new Http object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpWithDefaults

`func NewHttpWithDefaults() *Http`

NewHttpWithDefaults instantiates a new Http object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *Http) GetRequest() HttpRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *Http) GetRequestOk() (*HttpRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *Http) SetRequest(v HttpRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *Http) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *Http) GetResponse() HttpResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Http) GetResponseOk() (*HttpResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Http) SetResponse(v HttpResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Http) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *Http) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *Http) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



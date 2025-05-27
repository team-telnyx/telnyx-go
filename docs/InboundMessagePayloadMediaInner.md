# InboundMessagePayloadMediaInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The url of the media requested to be sent. | [optional] 
**ContentType** | Pointer to **string** | The MIME type of the requested media. | [optional] 
**Size** | Pointer to **int32** | The size of the requested media. | [optional] 
**HashSha256** | Pointer to **string** | The SHA256 hash of the requested media. | [optional] 

## Methods

### NewInboundMessagePayloadMediaInner

`func NewInboundMessagePayloadMediaInner() *InboundMessagePayloadMediaInner`

NewInboundMessagePayloadMediaInner instantiates a new InboundMessagePayloadMediaInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundMessagePayloadMediaInnerWithDefaults

`func NewInboundMessagePayloadMediaInnerWithDefaults() *InboundMessagePayloadMediaInner`

NewInboundMessagePayloadMediaInnerWithDefaults instantiates a new InboundMessagePayloadMediaInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *InboundMessagePayloadMediaInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InboundMessagePayloadMediaInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InboundMessagePayloadMediaInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InboundMessagePayloadMediaInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetContentType

`func (o *InboundMessagePayloadMediaInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *InboundMessagePayloadMediaInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *InboundMessagePayloadMediaInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *InboundMessagePayloadMediaInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSize

`func (o *InboundMessagePayloadMediaInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *InboundMessagePayloadMediaInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *InboundMessagePayloadMediaInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *InboundMessagePayloadMediaInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetHashSha256

`func (o *InboundMessagePayloadMediaInner) GetHashSha256() string`

GetHashSha256 returns the HashSha256 field if non-nil, zero value otherwise.

### GetHashSha256Ok

`func (o *InboundMessagePayloadMediaInner) GetHashSha256Ok() (*string, bool)`

GetHashSha256Ok returns a tuple with the HashSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashSha256

`func (o *InboundMessagePayloadMediaInner) SetHashSha256(v string)`

SetHashSha256 sets HashSha256 field to given value.

### HasHashSha256

`func (o *InboundMessagePayloadMediaInner) HasHashSha256() bool`

HasHashSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



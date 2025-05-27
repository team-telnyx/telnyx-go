# OutboundMessagePayloadMediaInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The url of the media requested to be sent. | [optional] 
**ContentType** | Pointer to **NullableString** | The MIME type of the requested media. | [optional] 
**Sha256** | Pointer to **NullableString** | The SHA256 hash of the requested media. | [optional] 
**Size** | Pointer to **NullableInt32** | The size of the requested media. | [optional] 

## Methods

### NewOutboundMessagePayloadMediaInner

`func NewOutboundMessagePayloadMediaInner() *OutboundMessagePayloadMediaInner`

NewOutboundMessagePayloadMediaInner instantiates a new OutboundMessagePayloadMediaInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundMessagePayloadMediaInnerWithDefaults

`func NewOutboundMessagePayloadMediaInnerWithDefaults() *OutboundMessagePayloadMediaInner`

NewOutboundMessagePayloadMediaInnerWithDefaults instantiates a new OutboundMessagePayloadMediaInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OutboundMessagePayloadMediaInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OutboundMessagePayloadMediaInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OutboundMessagePayloadMediaInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OutboundMessagePayloadMediaInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetContentType

`func (o *OutboundMessagePayloadMediaInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *OutboundMessagePayloadMediaInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *OutboundMessagePayloadMediaInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *OutboundMessagePayloadMediaInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *OutboundMessagePayloadMediaInner) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *OutboundMessagePayloadMediaInner) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetSha256

`func (o *OutboundMessagePayloadMediaInner) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *OutboundMessagePayloadMediaInner) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *OutboundMessagePayloadMediaInner) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *OutboundMessagePayloadMediaInner) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### SetSha256Nil

`func (o *OutboundMessagePayloadMediaInner) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *OutboundMessagePayloadMediaInner) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
### GetSize

`func (o *OutboundMessagePayloadMediaInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OutboundMessagePayloadMediaInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OutboundMessagePayloadMediaInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *OutboundMessagePayloadMediaInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *OutboundMessagePayloadMediaInner) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *OutboundMessagePayloadMediaInner) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



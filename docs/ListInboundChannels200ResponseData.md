# ListInboundChannels200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to **int32** | The current number of concurrent channels set for the account | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the response | [optional] 

## Methods

### NewListInboundChannels200ResponseData

`func NewListInboundChannels200ResponseData() *ListInboundChannels200ResponseData`

NewListInboundChannels200ResponseData instantiates a new ListInboundChannels200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInboundChannels200ResponseDataWithDefaults

`func NewListInboundChannels200ResponseDataWithDefaults() *ListInboundChannels200ResponseData`

NewListInboundChannels200ResponseDataWithDefaults instantiates a new ListInboundChannels200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ListInboundChannels200ResponseData) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ListInboundChannels200ResponseData) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ListInboundChannels200ResponseData) SetChannels(v int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ListInboundChannels200ResponseData) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetRecordType

`func (o *ListInboundChannels200ResponseData) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ListInboundChannels200ResponseData) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ListInboundChannels200ResponseData) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ListInboundChannels200ResponseData) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GcbChannelZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** |  | 
**Countries** | **[]string** | List of countries (in ISO 3166-2, capitalized) members of the billing channel zone | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Channels** | **int64** |  | 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date of when the channel zone was created | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date of when the channel zone was updated | [optional] 

## Methods

### NewGcbChannelZone

`func NewGcbChannelZone(recordType string, countries []string, id string, name string, channels int64, ) *GcbChannelZone`

NewGcbChannelZone instantiates a new GcbChannelZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcbChannelZoneWithDefaults

`func NewGcbChannelZoneWithDefaults() *GcbChannelZone`

NewGcbChannelZoneWithDefaults instantiates a new GcbChannelZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *GcbChannelZone) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *GcbChannelZone) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *GcbChannelZone) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetCountries

`func (o *GcbChannelZone) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *GcbChannelZone) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *GcbChannelZone) SetCountries(v []string)`

SetCountries sets Countries field to given value.


### GetId

`func (o *GcbChannelZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GcbChannelZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GcbChannelZone) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GcbChannelZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcbChannelZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcbChannelZone) SetName(v string)`

SetName sets Name field to given value.


### GetChannels

`func (o *GcbChannelZone) GetChannels() int64`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *GcbChannelZone) GetChannelsOk() (*int64, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *GcbChannelZone) SetChannels(v int64)`

SetChannels sets Channels field to given value.


### GetCreatedAt

`func (o *GcbChannelZone) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GcbChannelZone) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GcbChannelZone) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GcbChannelZone) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GcbChannelZone) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GcbChannelZone) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GcbChannelZone) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GcbChannelZone) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



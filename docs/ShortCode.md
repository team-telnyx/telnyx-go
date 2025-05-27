# ShortCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] [readonly] 
**ShortCode** | Pointer to **string** | Short digit sequence used to address messages. | [optional] [readonly] 
**CountryCode** | Pointer to **string** | ISO 3166-1 alpha-2 country code. | [optional] [readonly] 
**MessagingProfileId** | **NullableString** | Unique identifier for a messaging profile. | 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] [readonly] 

## Methods

### NewShortCode

`func NewShortCode(messagingProfileId NullableString, ) *ShortCode`

NewShortCode instantiates a new ShortCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortCodeWithDefaults

`func NewShortCodeWithDefaults() *ShortCode`

NewShortCodeWithDefaults instantiates a new ShortCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *ShortCode) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ShortCode) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ShortCode) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ShortCode) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *ShortCode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShortCode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShortCode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShortCode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShortCode

`func (o *ShortCode) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *ShortCode) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *ShortCode) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.

### HasShortCode

`func (o *ShortCode) HasShortCode() bool`

HasShortCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *ShortCode) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ShortCode) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ShortCode) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ShortCode) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *ShortCode) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *ShortCode) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *ShortCode) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.


### SetMessagingProfileIdNil

`func (o *ShortCode) SetMessagingProfileIdNil(b bool)`

 SetMessagingProfileIdNil sets the value for MessagingProfileId to be an explicit nil

### UnsetMessagingProfileId
`func (o *ShortCode) UnsetMessagingProfileId()`

UnsetMessagingProfileId ensures that no value is present for MessagingProfileId, not even an explicit nil
### GetCreatedAt

`func (o *ShortCode) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShortCode) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShortCode) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ShortCode) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ShortCode) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ShortCode) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ShortCode) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ShortCode) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



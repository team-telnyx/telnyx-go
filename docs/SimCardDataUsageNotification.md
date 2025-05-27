# SimCardDataUsageNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**SimCardId** | Pointer to **string** | The identification UUID of the related SIM card resource. | [optional] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**Threshold** | Pointer to [**PostSimCardDataUsageNotificationRequestThreshold**](PostSimCardDataUsageNotificationRequestThreshold.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 

## Methods

### NewSimCardDataUsageNotification

`func NewSimCardDataUsageNotification() *SimCardDataUsageNotification`

NewSimCardDataUsageNotification instantiates a new SimCardDataUsageNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimCardDataUsageNotificationWithDefaults

`func NewSimCardDataUsageNotificationWithDefaults() *SimCardDataUsageNotification`

NewSimCardDataUsageNotificationWithDefaults instantiates a new SimCardDataUsageNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimCardDataUsageNotification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimCardDataUsageNotification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimCardDataUsageNotification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimCardDataUsageNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSimCardId

`func (o *SimCardDataUsageNotification) GetSimCardId() string`

GetSimCardId returns the SimCardId field if non-nil, zero value otherwise.

### GetSimCardIdOk

`func (o *SimCardDataUsageNotification) GetSimCardIdOk() (*string, bool)`

GetSimCardIdOk returns a tuple with the SimCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardId

`func (o *SimCardDataUsageNotification) SetSimCardId(v string)`

SetSimCardId sets SimCardId field to given value.

### HasSimCardId

`func (o *SimCardDataUsageNotification) HasSimCardId() bool`

HasSimCardId returns a boolean if a field has been set.

### GetRecordType

`func (o *SimCardDataUsageNotification) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SimCardDataUsageNotification) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SimCardDataUsageNotification) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SimCardDataUsageNotification) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetThreshold

`func (o *SimCardDataUsageNotification) GetThreshold() PostSimCardDataUsageNotificationRequestThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *SimCardDataUsageNotification) GetThresholdOk() (*PostSimCardDataUsageNotificationRequestThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *SimCardDataUsageNotification) SetThreshold(v PostSimCardDataUsageNotificationRequestThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *SimCardDataUsageNotification) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SimCardDataUsageNotification) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimCardDataUsageNotification) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimCardDataUsageNotification) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SimCardDataUsageNotification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SimCardDataUsageNotification) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SimCardDataUsageNotification) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SimCardDataUsageNotification) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SimCardDataUsageNotification) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



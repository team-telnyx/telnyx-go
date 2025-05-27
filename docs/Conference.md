# Conference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** |  | 
**Id** | **string** | Uniquely identifies the conference | 
**Name** | **string** | Name of the conference | 
**CreatedAt** | **string** | ISO 8601 formatted date of when the conference was created | 
**ExpiresAt** | **string** | ISO 8601 formatted date of when the conference will expire | 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date of when the conference was last updated | [optional] 
**Region** | Pointer to **string** | Region where the conference is hosted | [optional] 
**Status** | Pointer to **string** | Status of the conference | [optional] 
**EndReason** | Pointer to **string** | Reason why the conference ended | [optional] 
**EndedBy** | Pointer to [**ConferenceEndedBy**](ConferenceEndedBy.md) |  | [optional] 
**ConnectionId** | Pointer to **string** | Identifies the connection associated with the conference | [optional] 

## Methods

### NewConference

`func NewConference(recordType string, id string, name string, createdAt string, expiresAt string, ) *Conference`

NewConference instantiates a new Conference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceWithDefaults

`func NewConferenceWithDefaults() *Conference`

NewConferenceWithDefaults instantiates a new Conference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *Conference) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *Conference) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *Conference) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetId

`func (o *Conference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Conference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Conference) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Conference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Conference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Conference) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *Conference) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Conference) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Conference) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *Conference) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Conference) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Conference) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetUpdatedAt

`func (o *Conference) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Conference) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Conference) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Conference) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRegion

`func (o *Conference) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Conference) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Conference) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Conference) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *Conference) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Conference) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Conference) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Conference) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEndReason

`func (o *Conference) GetEndReason() string`

GetEndReason returns the EndReason field if non-nil, zero value otherwise.

### GetEndReasonOk

`func (o *Conference) GetEndReasonOk() (*string, bool)`

GetEndReasonOk returns a tuple with the EndReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndReason

`func (o *Conference) SetEndReason(v string)`

SetEndReason sets EndReason field to given value.

### HasEndReason

`func (o *Conference) HasEndReason() bool`

HasEndReason returns a boolean if a field has been set.

### GetEndedBy

`func (o *Conference) GetEndedBy() ConferenceEndedBy`

GetEndedBy returns the EndedBy field if non-nil, zero value otherwise.

### GetEndedByOk

`func (o *Conference) GetEndedByOk() (*ConferenceEndedBy, bool)`

GetEndedByOk returns a tuple with the EndedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedBy

`func (o *Conference) SetEndedBy(v ConferenceEndedBy)`

SetEndedBy sets EndedBy field to given value.

### HasEndedBy

`func (o *Conference) HasEndedBy() bool`

HasEndedBy returns a boolean if a field has been set.

### GetConnectionId

`func (o *Conference) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Conference) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Conference) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Conference) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



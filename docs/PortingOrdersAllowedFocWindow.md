# PortingOrdersAllowedFocWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating the start of the range of foc window. | [optional] [readonly] 
**EndedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating the end of the range of foc window | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 

## Methods

### NewPortingOrdersAllowedFocWindow

`func NewPortingOrdersAllowedFocWindow() *PortingOrdersAllowedFocWindow`

NewPortingOrdersAllowedFocWindow instantiates a new PortingOrdersAllowedFocWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingOrdersAllowedFocWindowWithDefaults

`func NewPortingOrdersAllowedFocWindowWithDefaults() *PortingOrdersAllowedFocWindow`

NewPortingOrdersAllowedFocWindowWithDefaults instantiates a new PortingOrdersAllowedFocWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartedAt

`func (o *PortingOrdersAllowedFocWindow) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *PortingOrdersAllowedFocWindow) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *PortingOrdersAllowedFocWindow) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *PortingOrdersAllowedFocWindow) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *PortingOrdersAllowedFocWindow) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *PortingOrdersAllowedFocWindow) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *PortingOrdersAllowedFocWindow) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *PortingOrdersAllowedFocWindow) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingOrdersAllowedFocWindow) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingOrdersAllowedFocWindow) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingOrdersAllowedFocWindow) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingOrdersAllowedFocWindow) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



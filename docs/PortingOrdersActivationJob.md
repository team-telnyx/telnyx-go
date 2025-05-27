# PortingOrdersActivationJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies this activation job | [optional] [readonly] 
**Status** | Pointer to **string** | Specifies the status of this activation job | [optional] 
**ActivationType** | Pointer to **string** | Specifies the type of this activation job | [optional] 
**ActivateAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the activation job should be executed. This time should be between some activation window. | [optional] 
**ActivationWindows** | Pointer to [**[]PortingOrdersActivationJobActivationWindowsInner**](PortingOrdersActivationJobActivationWindowsInner.md) | List of allowed activation windows for this activation job | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 

## Methods

### NewPortingOrdersActivationJob

`func NewPortingOrdersActivationJob() *PortingOrdersActivationJob`

NewPortingOrdersActivationJob instantiates a new PortingOrdersActivationJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingOrdersActivationJobWithDefaults

`func NewPortingOrdersActivationJobWithDefaults() *PortingOrdersActivationJob`

NewPortingOrdersActivationJobWithDefaults instantiates a new PortingOrdersActivationJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortingOrdersActivationJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortingOrdersActivationJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortingOrdersActivationJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortingOrdersActivationJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PortingOrdersActivationJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortingOrdersActivationJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortingOrdersActivationJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PortingOrdersActivationJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActivationType

`func (o *PortingOrdersActivationJob) GetActivationType() string`

GetActivationType returns the ActivationType field if non-nil, zero value otherwise.

### GetActivationTypeOk

`func (o *PortingOrdersActivationJob) GetActivationTypeOk() (*string, bool)`

GetActivationTypeOk returns a tuple with the ActivationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationType

`func (o *PortingOrdersActivationJob) SetActivationType(v string)`

SetActivationType sets ActivationType field to given value.

### HasActivationType

`func (o *PortingOrdersActivationJob) HasActivationType() bool`

HasActivationType returns a boolean if a field has been set.

### GetActivateAt

`func (o *PortingOrdersActivationJob) GetActivateAt() time.Time`

GetActivateAt returns the ActivateAt field if non-nil, zero value otherwise.

### GetActivateAtOk

`func (o *PortingOrdersActivationJob) GetActivateAtOk() (*time.Time, bool)`

GetActivateAtOk returns a tuple with the ActivateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateAt

`func (o *PortingOrdersActivationJob) SetActivateAt(v time.Time)`

SetActivateAt sets ActivateAt field to given value.

### HasActivateAt

`func (o *PortingOrdersActivationJob) HasActivateAt() bool`

HasActivateAt returns a boolean if a field has been set.

### GetActivationWindows

`func (o *PortingOrdersActivationJob) GetActivationWindows() []PortingOrdersActivationJobActivationWindowsInner`

GetActivationWindows returns the ActivationWindows field if non-nil, zero value otherwise.

### GetActivationWindowsOk

`func (o *PortingOrdersActivationJob) GetActivationWindowsOk() (*[]PortingOrdersActivationJobActivationWindowsInner, bool)`

GetActivationWindowsOk returns a tuple with the ActivationWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationWindows

`func (o *PortingOrdersActivationJob) SetActivationWindows(v []PortingOrdersActivationJobActivationWindowsInner)`

SetActivationWindows sets ActivationWindows field to given value.

### HasActivationWindows

`func (o *PortingOrdersActivationJob) HasActivationWindows() bool`

HasActivationWindows returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingOrdersActivationJob) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingOrdersActivationJob) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingOrdersActivationJob) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingOrdersActivationJob) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortingOrdersActivationJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortingOrdersActivationJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortingOrdersActivationJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortingOrdersActivationJob) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortingOrdersActivationJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortingOrdersActivationJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortingOrdersActivationJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortingOrdersActivationJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



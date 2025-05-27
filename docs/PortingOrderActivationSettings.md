# PortingOrderActivationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FocDatetimeRequested** | Pointer to **time.Time** | ISO 8601 formatted Date/Time requested for the FOC date | [optional] 
**FocDatetimeActual** | Pointer to **time.Time** | ISO 8601 formatted Date/Time of the FOC date | [optional] 
**FastPortEligible** | Pointer to **bool** | Indicates whether this porting order is eligible for FastPort | [optional] [readonly] 
**ActivationStatus** | Pointer to [**PortingOrderActivationStatus**](PortingOrderActivationStatus.md) |  | [optional] 

## Methods

### NewPortingOrderActivationSettings

`func NewPortingOrderActivationSettings() *PortingOrderActivationSettings`

NewPortingOrderActivationSettings instantiates a new PortingOrderActivationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingOrderActivationSettingsWithDefaults

`func NewPortingOrderActivationSettingsWithDefaults() *PortingOrderActivationSettings`

NewPortingOrderActivationSettingsWithDefaults instantiates a new PortingOrderActivationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFocDatetimeRequested

`func (o *PortingOrderActivationSettings) GetFocDatetimeRequested() time.Time`

GetFocDatetimeRequested returns the FocDatetimeRequested field if non-nil, zero value otherwise.

### GetFocDatetimeRequestedOk

`func (o *PortingOrderActivationSettings) GetFocDatetimeRequestedOk() (*time.Time, bool)`

GetFocDatetimeRequestedOk returns a tuple with the FocDatetimeRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocDatetimeRequested

`func (o *PortingOrderActivationSettings) SetFocDatetimeRequested(v time.Time)`

SetFocDatetimeRequested sets FocDatetimeRequested field to given value.

### HasFocDatetimeRequested

`func (o *PortingOrderActivationSettings) HasFocDatetimeRequested() bool`

HasFocDatetimeRequested returns a boolean if a field has been set.

### GetFocDatetimeActual

`func (o *PortingOrderActivationSettings) GetFocDatetimeActual() time.Time`

GetFocDatetimeActual returns the FocDatetimeActual field if non-nil, zero value otherwise.

### GetFocDatetimeActualOk

`func (o *PortingOrderActivationSettings) GetFocDatetimeActualOk() (*time.Time, bool)`

GetFocDatetimeActualOk returns a tuple with the FocDatetimeActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocDatetimeActual

`func (o *PortingOrderActivationSettings) SetFocDatetimeActual(v time.Time)`

SetFocDatetimeActual sets FocDatetimeActual field to given value.

### HasFocDatetimeActual

`func (o *PortingOrderActivationSettings) HasFocDatetimeActual() bool`

HasFocDatetimeActual returns a boolean if a field has been set.

### GetFastPortEligible

`func (o *PortingOrderActivationSettings) GetFastPortEligible() bool`

GetFastPortEligible returns the FastPortEligible field if non-nil, zero value otherwise.

### GetFastPortEligibleOk

`func (o *PortingOrderActivationSettings) GetFastPortEligibleOk() (*bool, bool)`

GetFastPortEligibleOk returns a tuple with the FastPortEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastPortEligible

`func (o *PortingOrderActivationSettings) SetFastPortEligible(v bool)`

SetFastPortEligible sets FastPortEligible field to given value.

### HasFastPortEligible

`func (o *PortingOrderActivationSettings) HasFastPortEligible() bool`

HasFastPortEligible returns a boolean if a field has been set.

### GetActivationStatus

`func (o *PortingOrderActivationSettings) GetActivationStatus() PortingOrderActivationStatus`

GetActivationStatus returns the ActivationStatus field if non-nil, zero value otherwise.

### GetActivationStatusOk

`func (o *PortingOrderActivationSettings) GetActivationStatusOk() (*PortingOrderActivationStatus, bool)`

GetActivationStatusOk returns a tuple with the ActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationStatus

`func (o *PortingOrderActivationSettings) SetActivationStatus(v PortingOrderActivationStatus)`

SetActivationStatus sets ActivationStatus field to given value.

### HasActivationStatus

`func (o *PortingOrderActivationSettings) HasActivationStatus() bool`

HasActivationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



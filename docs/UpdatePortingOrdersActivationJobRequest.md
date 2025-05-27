# UpdatePortingOrdersActivationJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivateAt** | Pointer to **time.Time** | The desired activation time. The activation time should be between any of the activation windows. | [optional] 

## Methods

### NewUpdatePortingOrdersActivationJobRequest

`func NewUpdatePortingOrdersActivationJobRequest() *UpdatePortingOrdersActivationJobRequest`

NewUpdatePortingOrdersActivationJobRequest instantiates a new UpdatePortingOrdersActivationJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePortingOrdersActivationJobRequestWithDefaults

`func NewUpdatePortingOrdersActivationJobRequestWithDefaults() *UpdatePortingOrdersActivationJobRequest`

NewUpdatePortingOrdersActivationJobRequestWithDefaults instantiates a new UpdatePortingOrdersActivationJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivateAt

`func (o *UpdatePortingOrdersActivationJobRequest) GetActivateAt() time.Time`

GetActivateAt returns the ActivateAt field if non-nil, zero value otherwise.

### GetActivateAtOk

`func (o *UpdatePortingOrdersActivationJobRequest) GetActivateAtOk() (*time.Time, bool)`

GetActivateAtOk returns a tuple with the ActivateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateAt

`func (o *UpdatePortingOrdersActivationJobRequest) SetActivateAt(v time.Time)`

SetActivateAt sets ActivateAt field to given value.

### HasActivateAt

`func (o *UpdatePortingOrdersActivationJobRequest) HasActivateAt() bool`

HasActivateAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



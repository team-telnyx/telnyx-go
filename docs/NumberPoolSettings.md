# NumberPoolSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TollFreeWeight** | **float32** | Defines the probability weight for a Toll Free number to be selected when sending a message. The higher the weight the higher the probability. The sum of the weights for all number types does not necessarily need to add to 100. Weight must be a non-negative number, and when equal to zero it will remove the number type from the pool.  | 
**LongCodeWeight** | **float32** | Defines the probability weight for a Long Code number to be selected when sending a message. The higher the weight the higher the probability. The sum of the weights for all number types does not necessarily need to add to 100.  Weight must be a non-negative number, and when equal to zero it will remove the number type from the pool.  | 
**SkipUnhealthy** | **bool** | If set to true all unhealthy numbers will be automatically excluded from the pool. Health metrics per number are calculated on a regular basis, taking into account the deliverability rate and the amount of messages marked as spam by upstream carriers. Numbers with a deliverability rate below 25% or spam ratio over 75% will be considered unhealthy.  | 
**StickySender** | Pointer to **bool** | If set to true, Number Pool will try to choose the same sending number for all messages to a particular recipient. If the sending number becomes unhealthy and &#x60;skip_unhealthy&#x60; is set to true, a new number will be chosen.  | [optional] [default to false]
**Geomatch** | Pointer to **bool** | If set to true, Number Pool will try to choose a sending number with the same area code as the destination number. If there are no such numbers available, a nunber with a different area code will be chosen. Currently only NANP numbers are supported.  | [optional] [default to false]

## Methods

### NewNumberPoolSettings

`func NewNumberPoolSettings(tollFreeWeight float32, longCodeWeight float32, skipUnhealthy bool, ) *NumberPoolSettings`

NewNumberPoolSettings instantiates a new NumberPoolSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberPoolSettingsWithDefaults

`func NewNumberPoolSettingsWithDefaults() *NumberPoolSettings`

NewNumberPoolSettingsWithDefaults instantiates a new NumberPoolSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTollFreeWeight

`func (o *NumberPoolSettings) GetTollFreeWeight() float32`

GetTollFreeWeight returns the TollFreeWeight field if non-nil, zero value otherwise.

### GetTollFreeWeightOk

`func (o *NumberPoolSettings) GetTollFreeWeightOk() (*float32, bool)`

GetTollFreeWeightOk returns a tuple with the TollFreeWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTollFreeWeight

`func (o *NumberPoolSettings) SetTollFreeWeight(v float32)`

SetTollFreeWeight sets TollFreeWeight field to given value.


### GetLongCodeWeight

`func (o *NumberPoolSettings) GetLongCodeWeight() float32`

GetLongCodeWeight returns the LongCodeWeight field if non-nil, zero value otherwise.

### GetLongCodeWeightOk

`func (o *NumberPoolSettings) GetLongCodeWeightOk() (*float32, bool)`

GetLongCodeWeightOk returns a tuple with the LongCodeWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongCodeWeight

`func (o *NumberPoolSettings) SetLongCodeWeight(v float32)`

SetLongCodeWeight sets LongCodeWeight field to given value.


### GetSkipUnhealthy

`func (o *NumberPoolSettings) GetSkipUnhealthy() bool`

GetSkipUnhealthy returns the SkipUnhealthy field if non-nil, zero value otherwise.

### GetSkipUnhealthyOk

`func (o *NumberPoolSettings) GetSkipUnhealthyOk() (*bool, bool)`

GetSkipUnhealthyOk returns a tuple with the SkipUnhealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipUnhealthy

`func (o *NumberPoolSettings) SetSkipUnhealthy(v bool)`

SetSkipUnhealthy sets SkipUnhealthy field to given value.


### GetStickySender

`func (o *NumberPoolSettings) GetStickySender() bool`

GetStickySender returns the StickySender field if non-nil, zero value otherwise.

### GetStickySenderOk

`func (o *NumberPoolSettings) GetStickySenderOk() (*bool, bool)`

GetStickySenderOk returns a tuple with the StickySender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickySender

`func (o *NumberPoolSettings) SetStickySender(v bool)`

SetStickySender sets StickySender field to given value.

### HasStickySender

`func (o *NumberPoolSettings) HasStickySender() bool`

HasStickySender returns a boolean if a field has been set.

### GetGeomatch

`func (o *NumberPoolSettings) GetGeomatch() bool`

GetGeomatch returns the Geomatch field if non-nil, zero value otherwise.

### GetGeomatchOk

`func (o *NumberPoolSettings) GetGeomatchOk() (*bool, bool)`

GetGeomatchOk returns a tuple with the Geomatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeomatch

`func (o *NumberPoolSettings) SetGeomatch(v bool)`

SetGeomatch sets Geomatch field to given value.

### HasGeomatch

`func (o *NumberPoolSettings) HasGeomatch() bool`

HasGeomatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OrderExternalVetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvpId** | **string** | External vetting provider ID for the brand. | 
**VettingClass** | **string** | Identifies the vetting classification. | 

## Methods

### NewOrderExternalVetting

`func NewOrderExternalVetting(evpId string, vettingClass string, ) *OrderExternalVetting`

NewOrderExternalVetting instantiates a new OrderExternalVetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderExternalVettingWithDefaults

`func NewOrderExternalVettingWithDefaults() *OrderExternalVetting`

NewOrderExternalVettingWithDefaults instantiates a new OrderExternalVetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvpId

`func (o *OrderExternalVetting) GetEvpId() string`

GetEvpId returns the EvpId field if non-nil, zero value otherwise.

### GetEvpIdOk

`func (o *OrderExternalVetting) GetEvpIdOk() (*string, bool)`

GetEvpIdOk returns a tuple with the EvpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpId

`func (o *OrderExternalVetting) SetEvpId(v string)`

SetEvpId sets EvpId field to given value.


### GetVettingClass

`func (o *OrderExternalVetting) GetVettingClass() string`

GetVettingClass returns the VettingClass field if non-nil, zero value otherwise.

### GetVettingClassOk

`func (o *OrderExternalVetting) GetVettingClassOk() (*string, bool)`

GetVettingClassOk returns a tuple with the VettingClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVettingClass

`func (o *OrderExternalVetting) SetVettingClass(v string)`

SetVettingClass sets VettingClass field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



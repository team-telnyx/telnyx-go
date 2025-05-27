# ManagedAccountsGlobalOutboundChannels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagedAccountAllowCustomPricing** | Pointer to **bool** | Boolean value that indicates if the managed account is able to have custom pricing set for it or not. If false, uses the pricing of the manager account. Defaults to false. This value may be changed, but there may be time lag between when the value is changed and pricing changes take effect. | [optional] 
**AllocatableGlobalOutboundChannels** | Pointer to **int32** | The total amount of allocatable global outbound channels available to the authenticated manager. Will be 0 if the feature is not enabled for their account. | [optional] 
**RecordType** | Pointer to **string** | The type of the data contained in this record. | [optional] 
**TotalGlobalChannelsAllocated** | Pointer to **int32** | The total number of allocatable global outbound channels currently allocated across all managed accounts for the authenticated user. This includes any amount of channels allocated by default at managed account creation time. Will be 0 if the feature is not enabled for their account. | [optional] 

## Methods

### NewManagedAccountsGlobalOutboundChannels

`func NewManagedAccountsGlobalOutboundChannels() *ManagedAccountsGlobalOutboundChannels`

NewManagedAccountsGlobalOutboundChannels instantiates a new ManagedAccountsGlobalOutboundChannels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAccountsGlobalOutboundChannelsWithDefaults

`func NewManagedAccountsGlobalOutboundChannelsWithDefaults() *ManagedAccountsGlobalOutboundChannels`

NewManagedAccountsGlobalOutboundChannelsWithDefaults instantiates a new ManagedAccountsGlobalOutboundChannels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagedAccountAllowCustomPricing

`func (o *ManagedAccountsGlobalOutboundChannels) GetManagedAccountAllowCustomPricing() bool`

GetManagedAccountAllowCustomPricing returns the ManagedAccountAllowCustomPricing field if non-nil, zero value otherwise.

### GetManagedAccountAllowCustomPricingOk

`func (o *ManagedAccountsGlobalOutboundChannels) GetManagedAccountAllowCustomPricingOk() (*bool, bool)`

GetManagedAccountAllowCustomPricingOk returns a tuple with the ManagedAccountAllowCustomPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAccountAllowCustomPricing

`func (o *ManagedAccountsGlobalOutboundChannels) SetManagedAccountAllowCustomPricing(v bool)`

SetManagedAccountAllowCustomPricing sets ManagedAccountAllowCustomPricing field to given value.

### HasManagedAccountAllowCustomPricing

`func (o *ManagedAccountsGlobalOutboundChannels) HasManagedAccountAllowCustomPricing() bool`

HasManagedAccountAllowCustomPricing returns a boolean if a field has been set.

### GetAllocatableGlobalOutboundChannels

`func (o *ManagedAccountsGlobalOutboundChannels) GetAllocatableGlobalOutboundChannels() int32`

GetAllocatableGlobalOutboundChannels returns the AllocatableGlobalOutboundChannels field if non-nil, zero value otherwise.

### GetAllocatableGlobalOutboundChannelsOk

`func (o *ManagedAccountsGlobalOutboundChannels) GetAllocatableGlobalOutboundChannelsOk() (*int32, bool)`

GetAllocatableGlobalOutboundChannelsOk returns a tuple with the AllocatableGlobalOutboundChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatableGlobalOutboundChannels

`func (o *ManagedAccountsGlobalOutboundChannels) SetAllocatableGlobalOutboundChannels(v int32)`

SetAllocatableGlobalOutboundChannels sets AllocatableGlobalOutboundChannels field to given value.

### HasAllocatableGlobalOutboundChannels

`func (o *ManagedAccountsGlobalOutboundChannels) HasAllocatableGlobalOutboundChannels() bool`

HasAllocatableGlobalOutboundChannels returns a boolean if a field has been set.

### GetRecordType

`func (o *ManagedAccountsGlobalOutboundChannels) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ManagedAccountsGlobalOutboundChannels) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ManagedAccountsGlobalOutboundChannels) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ManagedAccountsGlobalOutboundChannels) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetTotalGlobalChannelsAllocated

`func (o *ManagedAccountsGlobalOutboundChannels) GetTotalGlobalChannelsAllocated() int32`

GetTotalGlobalChannelsAllocated returns the TotalGlobalChannelsAllocated field if non-nil, zero value otherwise.

### GetTotalGlobalChannelsAllocatedOk

`func (o *ManagedAccountsGlobalOutboundChannels) GetTotalGlobalChannelsAllocatedOk() (*int32, bool)`

GetTotalGlobalChannelsAllocatedOk returns a tuple with the TotalGlobalChannelsAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalChannelsAllocated

`func (o *ManagedAccountsGlobalOutboundChannels) SetTotalGlobalChannelsAllocated(v int32)`

SetTotalGlobalChannelsAllocated sets TotalGlobalChannelsAllocated field to given value.

### HasTotalGlobalChannelsAllocated

`func (o *ManagedAccountsGlobalOutboundChannels) HasTotalGlobalChannelsAllocated() bool`

HasTotalGlobalChannelsAllocated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



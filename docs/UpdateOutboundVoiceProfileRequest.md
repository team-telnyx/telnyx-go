# UpdateOutboundVoiceProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user-supplied name to help with organization. | 
**TrafficType** | Pointer to [**TrafficType**](TrafficType.md) |  | [optional] [default to CONVERSATIONAL]
**ServicePlan** | Pointer to [**ServicePlan**](ServicePlan.md) |  | [optional] [default to GLOBAL]
**ConcurrentCallLimit** | Pointer to **NullableInt32** | Must be no more than your global concurrent call limit. Null means no limit. | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether the outbound voice profile can be used. Disabled profiles will result in outbound calls being blocked for the associated Connections. | [optional] [default to true]
**Tags** | Pointer to **[]string** |  | [optional] 
**UsagePaymentMethod** | Pointer to [**UsagePaymentMethod**](UsagePaymentMethod.md) |  | [optional] [default to RATE_DECK]
**WhitelistedDestinations** | Pointer to **[]string** | The list of destinations you want to be able to call using this outbound voice profile formatted in alpha2. | [optional] [default to ["US","CA"]]
**MaxDestinationRate** | Pointer to **float32** | Maximum rate (price per minute) for a Destination to be allowed when making outbound calls. | [optional] 
**DailySpendLimit** | Pointer to **string** | The maximum amount of usage charges, in USD, you want Telnyx to allow on this outbound voice profile in a day before disallowing new calls. | [optional] 
**DailySpendLimitEnabled** | Pointer to **bool** | Specifies whether to enforce the daily_spend_limit on this outbound voice profile. | [optional] [default to false]
**CallRecording** | Pointer to [**OutboundCallRecording**](OutboundCallRecording.md) |  | [optional] 
**BillingGroupId** | Pointer to **NullableString** | The ID of the billing group associated with the outbound proflile. Defaults to null (for no group assigned). | [optional] 

## Methods

### NewUpdateOutboundVoiceProfileRequest

`func NewUpdateOutboundVoiceProfileRequest(name string, ) *UpdateOutboundVoiceProfileRequest`

NewUpdateOutboundVoiceProfileRequest instantiates a new UpdateOutboundVoiceProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOutboundVoiceProfileRequestWithDefaults

`func NewUpdateOutboundVoiceProfileRequestWithDefaults() *UpdateOutboundVoiceProfileRequest`

NewUpdateOutboundVoiceProfileRequestWithDefaults instantiates a new UpdateOutboundVoiceProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOutboundVoiceProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOutboundVoiceProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOutboundVoiceProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTrafficType

`func (o *UpdateOutboundVoiceProfileRequest) GetTrafficType() TrafficType`

GetTrafficType returns the TrafficType field if non-nil, zero value otherwise.

### GetTrafficTypeOk

`func (o *UpdateOutboundVoiceProfileRequest) GetTrafficTypeOk() (*TrafficType, bool)`

GetTrafficTypeOk returns a tuple with the TrafficType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficType

`func (o *UpdateOutboundVoiceProfileRequest) SetTrafficType(v TrafficType)`

SetTrafficType sets TrafficType field to given value.

### HasTrafficType

`func (o *UpdateOutboundVoiceProfileRequest) HasTrafficType() bool`

HasTrafficType returns a boolean if a field has been set.

### GetServicePlan

`func (o *UpdateOutboundVoiceProfileRequest) GetServicePlan() ServicePlan`

GetServicePlan returns the ServicePlan field if non-nil, zero value otherwise.

### GetServicePlanOk

`func (o *UpdateOutboundVoiceProfileRequest) GetServicePlanOk() (*ServicePlan, bool)`

GetServicePlanOk returns a tuple with the ServicePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlan

`func (o *UpdateOutboundVoiceProfileRequest) SetServicePlan(v ServicePlan)`

SetServicePlan sets ServicePlan field to given value.

### HasServicePlan

`func (o *UpdateOutboundVoiceProfileRequest) HasServicePlan() bool`

HasServicePlan returns a boolean if a field has been set.

### GetConcurrentCallLimit

`func (o *UpdateOutboundVoiceProfileRequest) GetConcurrentCallLimit() int32`

GetConcurrentCallLimit returns the ConcurrentCallLimit field if non-nil, zero value otherwise.

### GetConcurrentCallLimitOk

`func (o *UpdateOutboundVoiceProfileRequest) GetConcurrentCallLimitOk() (*int32, bool)`

GetConcurrentCallLimitOk returns a tuple with the ConcurrentCallLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentCallLimit

`func (o *UpdateOutboundVoiceProfileRequest) SetConcurrentCallLimit(v int32)`

SetConcurrentCallLimit sets ConcurrentCallLimit field to given value.

### HasConcurrentCallLimit

`func (o *UpdateOutboundVoiceProfileRequest) HasConcurrentCallLimit() bool`

HasConcurrentCallLimit returns a boolean if a field has been set.

### SetConcurrentCallLimitNil

`func (o *UpdateOutboundVoiceProfileRequest) SetConcurrentCallLimitNil(b bool)`

 SetConcurrentCallLimitNil sets the value for ConcurrentCallLimit to be an explicit nil

### UnsetConcurrentCallLimit
`func (o *UpdateOutboundVoiceProfileRequest) UnsetConcurrentCallLimit()`

UnsetConcurrentCallLimit ensures that no value is present for ConcurrentCallLimit, not even an explicit nil
### GetEnabled

`func (o *UpdateOutboundVoiceProfileRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateOutboundVoiceProfileRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateOutboundVoiceProfileRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateOutboundVoiceProfileRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTags

`func (o *UpdateOutboundVoiceProfileRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateOutboundVoiceProfileRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateOutboundVoiceProfileRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateOutboundVoiceProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUsagePaymentMethod

`func (o *UpdateOutboundVoiceProfileRequest) GetUsagePaymentMethod() UsagePaymentMethod`

GetUsagePaymentMethod returns the UsagePaymentMethod field if non-nil, zero value otherwise.

### GetUsagePaymentMethodOk

`func (o *UpdateOutboundVoiceProfileRequest) GetUsagePaymentMethodOk() (*UsagePaymentMethod, bool)`

GetUsagePaymentMethodOk returns a tuple with the UsagePaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePaymentMethod

`func (o *UpdateOutboundVoiceProfileRequest) SetUsagePaymentMethod(v UsagePaymentMethod)`

SetUsagePaymentMethod sets UsagePaymentMethod field to given value.

### HasUsagePaymentMethod

`func (o *UpdateOutboundVoiceProfileRequest) HasUsagePaymentMethod() bool`

HasUsagePaymentMethod returns a boolean if a field has been set.

### GetWhitelistedDestinations

`func (o *UpdateOutboundVoiceProfileRequest) GetWhitelistedDestinations() []string`

GetWhitelistedDestinations returns the WhitelistedDestinations field if non-nil, zero value otherwise.

### GetWhitelistedDestinationsOk

`func (o *UpdateOutboundVoiceProfileRequest) GetWhitelistedDestinationsOk() (*[]string, bool)`

GetWhitelistedDestinationsOk returns a tuple with the WhitelistedDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedDestinations

`func (o *UpdateOutboundVoiceProfileRequest) SetWhitelistedDestinations(v []string)`

SetWhitelistedDestinations sets WhitelistedDestinations field to given value.

### HasWhitelistedDestinations

`func (o *UpdateOutboundVoiceProfileRequest) HasWhitelistedDestinations() bool`

HasWhitelistedDestinations returns a boolean if a field has been set.

### GetMaxDestinationRate

`func (o *UpdateOutboundVoiceProfileRequest) GetMaxDestinationRate() float32`

GetMaxDestinationRate returns the MaxDestinationRate field if non-nil, zero value otherwise.

### GetMaxDestinationRateOk

`func (o *UpdateOutboundVoiceProfileRequest) GetMaxDestinationRateOk() (*float32, bool)`

GetMaxDestinationRateOk returns a tuple with the MaxDestinationRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDestinationRate

`func (o *UpdateOutboundVoiceProfileRequest) SetMaxDestinationRate(v float32)`

SetMaxDestinationRate sets MaxDestinationRate field to given value.

### HasMaxDestinationRate

`func (o *UpdateOutboundVoiceProfileRequest) HasMaxDestinationRate() bool`

HasMaxDestinationRate returns a boolean if a field has been set.

### GetDailySpendLimit

`func (o *UpdateOutboundVoiceProfileRequest) GetDailySpendLimit() string`

GetDailySpendLimit returns the DailySpendLimit field if non-nil, zero value otherwise.

### GetDailySpendLimitOk

`func (o *UpdateOutboundVoiceProfileRequest) GetDailySpendLimitOk() (*string, bool)`

GetDailySpendLimitOk returns a tuple with the DailySpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySpendLimit

`func (o *UpdateOutboundVoiceProfileRequest) SetDailySpendLimit(v string)`

SetDailySpendLimit sets DailySpendLimit field to given value.

### HasDailySpendLimit

`func (o *UpdateOutboundVoiceProfileRequest) HasDailySpendLimit() bool`

HasDailySpendLimit returns a boolean if a field has been set.

### GetDailySpendLimitEnabled

`func (o *UpdateOutboundVoiceProfileRequest) GetDailySpendLimitEnabled() bool`

GetDailySpendLimitEnabled returns the DailySpendLimitEnabled field if non-nil, zero value otherwise.

### GetDailySpendLimitEnabledOk

`func (o *UpdateOutboundVoiceProfileRequest) GetDailySpendLimitEnabledOk() (*bool, bool)`

GetDailySpendLimitEnabledOk returns a tuple with the DailySpendLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySpendLimitEnabled

`func (o *UpdateOutboundVoiceProfileRequest) SetDailySpendLimitEnabled(v bool)`

SetDailySpendLimitEnabled sets DailySpendLimitEnabled field to given value.

### HasDailySpendLimitEnabled

`func (o *UpdateOutboundVoiceProfileRequest) HasDailySpendLimitEnabled() bool`

HasDailySpendLimitEnabled returns a boolean if a field has been set.

### GetCallRecording

`func (o *UpdateOutboundVoiceProfileRequest) GetCallRecording() OutboundCallRecording`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *UpdateOutboundVoiceProfileRequest) GetCallRecordingOk() (*OutboundCallRecording, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *UpdateOutboundVoiceProfileRequest) SetCallRecording(v OutboundCallRecording)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *UpdateOutboundVoiceProfileRequest) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *UpdateOutboundVoiceProfileRequest) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *UpdateOutboundVoiceProfileRequest) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *UpdateOutboundVoiceProfileRequest) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *UpdateOutboundVoiceProfileRequest) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### SetBillingGroupIdNil

`func (o *UpdateOutboundVoiceProfileRequest) SetBillingGroupIdNil(b bool)`

 SetBillingGroupIdNil sets the value for BillingGroupId to be an explicit nil

### UnsetBillingGroupId
`func (o *UpdateOutboundVoiceProfileRequest) UnsetBillingGroupId()`

UnsetBillingGroupId ensures that no value is present for BillingGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



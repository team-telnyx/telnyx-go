# OutboundVoiceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Name** | **string** | A user-supplied name to help with organization. | 
**ConnectionsCount** | Pointer to **int32** | Amount of connections associated with this outbound voice profile. | [optional] 
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
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] 

## Methods

### NewOutboundVoiceProfile

`func NewOutboundVoiceProfile(name string, ) *OutboundVoiceProfile`

NewOutboundVoiceProfile instantiates a new OutboundVoiceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundVoiceProfileWithDefaults

`func NewOutboundVoiceProfileWithDefaults() *OutboundVoiceProfile`

NewOutboundVoiceProfileWithDefaults instantiates a new OutboundVoiceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutboundVoiceProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutboundVoiceProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutboundVoiceProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OutboundVoiceProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *OutboundVoiceProfile) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *OutboundVoiceProfile) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *OutboundVoiceProfile) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *OutboundVoiceProfile) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetName

`func (o *OutboundVoiceProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutboundVoiceProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutboundVoiceProfile) SetName(v string)`

SetName sets Name field to given value.


### GetConnectionsCount

`func (o *OutboundVoiceProfile) GetConnectionsCount() int32`

GetConnectionsCount returns the ConnectionsCount field if non-nil, zero value otherwise.

### GetConnectionsCountOk

`func (o *OutboundVoiceProfile) GetConnectionsCountOk() (*int32, bool)`

GetConnectionsCountOk returns a tuple with the ConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsCount

`func (o *OutboundVoiceProfile) SetConnectionsCount(v int32)`

SetConnectionsCount sets ConnectionsCount field to given value.

### HasConnectionsCount

`func (o *OutboundVoiceProfile) HasConnectionsCount() bool`

HasConnectionsCount returns a boolean if a field has been set.

### GetTrafficType

`func (o *OutboundVoiceProfile) GetTrafficType() TrafficType`

GetTrafficType returns the TrafficType field if non-nil, zero value otherwise.

### GetTrafficTypeOk

`func (o *OutboundVoiceProfile) GetTrafficTypeOk() (*TrafficType, bool)`

GetTrafficTypeOk returns a tuple with the TrafficType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficType

`func (o *OutboundVoiceProfile) SetTrafficType(v TrafficType)`

SetTrafficType sets TrafficType field to given value.

### HasTrafficType

`func (o *OutboundVoiceProfile) HasTrafficType() bool`

HasTrafficType returns a boolean if a field has been set.

### GetServicePlan

`func (o *OutboundVoiceProfile) GetServicePlan() ServicePlan`

GetServicePlan returns the ServicePlan field if non-nil, zero value otherwise.

### GetServicePlanOk

`func (o *OutboundVoiceProfile) GetServicePlanOk() (*ServicePlan, bool)`

GetServicePlanOk returns a tuple with the ServicePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlan

`func (o *OutboundVoiceProfile) SetServicePlan(v ServicePlan)`

SetServicePlan sets ServicePlan field to given value.

### HasServicePlan

`func (o *OutboundVoiceProfile) HasServicePlan() bool`

HasServicePlan returns a boolean if a field has been set.

### GetConcurrentCallLimit

`func (o *OutboundVoiceProfile) GetConcurrentCallLimit() int32`

GetConcurrentCallLimit returns the ConcurrentCallLimit field if non-nil, zero value otherwise.

### GetConcurrentCallLimitOk

`func (o *OutboundVoiceProfile) GetConcurrentCallLimitOk() (*int32, bool)`

GetConcurrentCallLimitOk returns a tuple with the ConcurrentCallLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentCallLimit

`func (o *OutboundVoiceProfile) SetConcurrentCallLimit(v int32)`

SetConcurrentCallLimit sets ConcurrentCallLimit field to given value.

### HasConcurrentCallLimit

`func (o *OutboundVoiceProfile) HasConcurrentCallLimit() bool`

HasConcurrentCallLimit returns a boolean if a field has been set.

### SetConcurrentCallLimitNil

`func (o *OutboundVoiceProfile) SetConcurrentCallLimitNil(b bool)`

 SetConcurrentCallLimitNil sets the value for ConcurrentCallLimit to be an explicit nil

### UnsetConcurrentCallLimit
`func (o *OutboundVoiceProfile) UnsetConcurrentCallLimit()`

UnsetConcurrentCallLimit ensures that no value is present for ConcurrentCallLimit, not even an explicit nil
### GetEnabled

`func (o *OutboundVoiceProfile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OutboundVoiceProfile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OutboundVoiceProfile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OutboundVoiceProfile) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTags

`func (o *OutboundVoiceProfile) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OutboundVoiceProfile) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OutboundVoiceProfile) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OutboundVoiceProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUsagePaymentMethod

`func (o *OutboundVoiceProfile) GetUsagePaymentMethod() UsagePaymentMethod`

GetUsagePaymentMethod returns the UsagePaymentMethod field if non-nil, zero value otherwise.

### GetUsagePaymentMethodOk

`func (o *OutboundVoiceProfile) GetUsagePaymentMethodOk() (*UsagePaymentMethod, bool)`

GetUsagePaymentMethodOk returns a tuple with the UsagePaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePaymentMethod

`func (o *OutboundVoiceProfile) SetUsagePaymentMethod(v UsagePaymentMethod)`

SetUsagePaymentMethod sets UsagePaymentMethod field to given value.

### HasUsagePaymentMethod

`func (o *OutboundVoiceProfile) HasUsagePaymentMethod() bool`

HasUsagePaymentMethod returns a boolean if a field has been set.

### GetWhitelistedDestinations

`func (o *OutboundVoiceProfile) GetWhitelistedDestinations() []string`

GetWhitelistedDestinations returns the WhitelistedDestinations field if non-nil, zero value otherwise.

### GetWhitelistedDestinationsOk

`func (o *OutboundVoiceProfile) GetWhitelistedDestinationsOk() (*[]string, bool)`

GetWhitelistedDestinationsOk returns a tuple with the WhitelistedDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedDestinations

`func (o *OutboundVoiceProfile) SetWhitelistedDestinations(v []string)`

SetWhitelistedDestinations sets WhitelistedDestinations field to given value.

### HasWhitelistedDestinations

`func (o *OutboundVoiceProfile) HasWhitelistedDestinations() bool`

HasWhitelistedDestinations returns a boolean if a field has been set.

### GetMaxDestinationRate

`func (o *OutboundVoiceProfile) GetMaxDestinationRate() float32`

GetMaxDestinationRate returns the MaxDestinationRate field if non-nil, zero value otherwise.

### GetMaxDestinationRateOk

`func (o *OutboundVoiceProfile) GetMaxDestinationRateOk() (*float32, bool)`

GetMaxDestinationRateOk returns a tuple with the MaxDestinationRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDestinationRate

`func (o *OutboundVoiceProfile) SetMaxDestinationRate(v float32)`

SetMaxDestinationRate sets MaxDestinationRate field to given value.

### HasMaxDestinationRate

`func (o *OutboundVoiceProfile) HasMaxDestinationRate() bool`

HasMaxDestinationRate returns a boolean if a field has been set.

### GetDailySpendLimit

`func (o *OutboundVoiceProfile) GetDailySpendLimit() string`

GetDailySpendLimit returns the DailySpendLimit field if non-nil, zero value otherwise.

### GetDailySpendLimitOk

`func (o *OutboundVoiceProfile) GetDailySpendLimitOk() (*string, bool)`

GetDailySpendLimitOk returns a tuple with the DailySpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySpendLimit

`func (o *OutboundVoiceProfile) SetDailySpendLimit(v string)`

SetDailySpendLimit sets DailySpendLimit field to given value.

### HasDailySpendLimit

`func (o *OutboundVoiceProfile) HasDailySpendLimit() bool`

HasDailySpendLimit returns a boolean if a field has been set.

### GetDailySpendLimitEnabled

`func (o *OutboundVoiceProfile) GetDailySpendLimitEnabled() bool`

GetDailySpendLimitEnabled returns the DailySpendLimitEnabled field if non-nil, zero value otherwise.

### GetDailySpendLimitEnabledOk

`func (o *OutboundVoiceProfile) GetDailySpendLimitEnabledOk() (*bool, bool)`

GetDailySpendLimitEnabledOk returns a tuple with the DailySpendLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySpendLimitEnabled

`func (o *OutboundVoiceProfile) SetDailySpendLimitEnabled(v bool)`

SetDailySpendLimitEnabled sets DailySpendLimitEnabled field to given value.

### HasDailySpendLimitEnabled

`func (o *OutboundVoiceProfile) HasDailySpendLimitEnabled() bool`

HasDailySpendLimitEnabled returns a boolean if a field has been set.

### GetCallRecording

`func (o *OutboundVoiceProfile) GetCallRecording() OutboundCallRecording`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *OutboundVoiceProfile) GetCallRecordingOk() (*OutboundCallRecording, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *OutboundVoiceProfile) SetCallRecording(v OutboundCallRecording)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *OutboundVoiceProfile) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *OutboundVoiceProfile) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *OutboundVoiceProfile) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *OutboundVoiceProfile) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *OutboundVoiceProfile) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### SetBillingGroupIdNil

`func (o *OutboundVoiceProfile) SetBillingGroupIdNil(b bool)`

 SetBillingGroupIdNil sets the value for BillingGroupId to be an explicit nil

### UnsetBillingGroupId
`func (o *OutboundVoiceProfile) UnsetBillingGroupId()`

UnsetBillingGroupId ensures that no value is present for BillingGroupId, not even an explicit nil
### GetCreatedAt

`func (o *OutboundVoiceProfile) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OutboundVoiceProfile) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OutboundVoiceProfile) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OutboundVoiceProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OutboundVoiceProfile) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OutboundVoiceProfile) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OutboundVoiceProfile) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OutboundVoiceProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



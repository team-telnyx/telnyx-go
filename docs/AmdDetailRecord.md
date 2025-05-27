# AmdDetailRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Feature invocation id | [optional] 
**InvokedAt** | Pointer to **time.Time** | Feature invocation time | [optional] 
**Feature** | Pointer to **string** | Feature name | [optional] 
**Tags** | Pointer to **string** | User-provided tags | [optional] 
**BillingGroupId** | Pointer to **string** | Billing Group id | [optional] 
**BillingGroupName** | Pointer to **string** | Name of the Billing Group specified in billing_group_id | [optional] 
**ConnectionId** | Pointer to **string** | Connection id | [optional] 
**ConnectionName** | Pointer to **string** | Connection name | [optional] 
**CallLegId** | Pointer to **string** | Telnyx UUID that identifies the related call leg | [optional] 
**CallSessionId** | Pointer to **string** | Telnyx UUID that identifies the related call session | [optional] 
**IsTelnyxBillable** | Pointer to **bool** | Indicates whether Telnyx billing charges might be applicable | [optional] 
**Rate** | Pointer to **string** | Currency amount per billing unit used to calculate the Telnyx billing cost | [optional] 
**RateMeasuredIn** | Pointer to **string** | Billing unit used to calculate the Telnyx billing cost | [optional] 
**Cost** | Pointer to **string** | Currency amount for Telnyx billing cost | [optional] 
**Currency** | Pointer to **string** | Telnyx account currency used to describe monetary values, including billing cost | [optional] 
**RecordType** | **string** |  | [default to "amd_detail_record"]

## Methods

### NewAmdDetailRecord

`func NewAmdDetailRecord(recordType string, ) *AmdDetailRecord`

NewAmdDetailRecord instantiates a new AmdDetailRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmdDetailRecordWithDefaults

`func NewAmdDetailRecordWithDefaults() *AmdDetailRecord`

NewAmdDetailRecordWithDefaults instantiates a new AmdDetailRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AmdDetailRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AmdDetailRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AmdDetailRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AmdDetailRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvokedAt

`func (o *AmdDetailRecord) GetInvokedAt() time.Time`

GetInvokedAt returns the InvokedAt field if non-nil, zero value otherwise.

### GetInvokedAtOk

`func (o *AmdDetailRecord) GetInvokedAtOk() (*time.Time, bool)`

GetInvokedAtOk returns a tuple with the InvokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokedAt

`func (o *AmdDetailRecord) SetInvokedAt(v time.Time)`

SetInvokedAt sets InvokedAt field to given value.

### HasInvokedAt

`func (o *AmdDetailRecord) HasInvokedAt() bool`

HasInvokedAt returns a boolean if a field has been set.

### GetFeature

`func (o *AmdDetailRecord) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *AmdDetailRecord) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *AmdDetailRecord) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *AmdDetailRecord) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetTags

`func (o *AmdDetailRecord) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AmdDetailRecord) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AmdDetailRecord) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AmdDetailRecord) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *AmdDetailRecord) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *AmdDetailRecord) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *AmdDetailRecord) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *AmdDetailRecord) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### GetBillingGroupName

`func (o *AmdDetailRecord) GetBillingGroupName() string`

GetBillingGroupName returns the BillingGroupName field if non-nil, zero value otherwise.

### GetBillingGroupNameOk

`func (o *AmdDetailRecord) GetBillingGroupNameOk() (*string, bool)`

GetBillingGroupNameOk returns a tuple with the BillingGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupName

`func (o *AmdDetailRecord) SetBillingGroupName(v string)`

SetBillingGroupName sets BillingGroupName field to given value.

### HasBillingGroupName

`func (o *AmdDetailRecord) HasBillingGroupName() bool`

HasBillingGroupName returns a boolean if a field has been set.

### GetConnectionId

`func (o *AmdDetailRecord) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *AmdDetailRecord) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *AmdDetailRecord) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *AmdDetailRecord) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnectionName

`func (o *AmdDetailRecord) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *AmdDetailRecord) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *AmdDetailRecord) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *AmdDetailRecord) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetCallLegId

`func (o *AmdDetailRecord) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *AmdDetailRecord) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *AmdDetailRecord) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *AmdDetailRecord) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *AmdDetailRecord) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *AmdDetailRecord) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *AmdDetailRecord) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *AmdDetailRecord) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetIsTelnyxBillable

`func (o *AmdDetailRecord) GetIsTelnyxBillable() bool`

GetIsTelnyxBillable returns the IsTelnyxBillable field if non-nil, zero value otherwise.

### GetIsTelnyxBillableOk

`func (o *AmdDetailRecord) GetIsTelnyxBillableOk() (*bool, bool)`

GetIsTelnyxBillableOk returns a tuple with the IsTelnyxBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTelnyxBillable

`func (o *AmdDetailRecord) SetIsTelnyxBillable(v bool)`

SetIsTelnyxBillable sets IsTelnyxBillable field to given value.

### HasIsTelnyxBillable

`func (o *AmdDetailRecord) HasIsTelnyxBillable() bool`

HasIsTelnyxBillable returns a boolean if a field has been set.

### GetRate

`func (o *AmdDetailRecord) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *AmdDetailRecord) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *AmdDetailRecord) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *AmdDetailRecord) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRateMeasuredIn

`func (o *AmdDetailRecord) GetRateMeasuredIn() string`

GetRateMeasuredIn returns the RateMeasuredIn field if non-nil, zero value otherwise.

### GetRateMeasuredInOk

`func (o *AmdDetailRecord) GetRateMeasuredInOk() (*string, bool)`

GetRateMeasuredInOk returns a tuple with the RateMeasuredIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateMeasuredIn

`func (o *AmdDetailRecord) SetRateMeasuredIn(v string)`

SetRateMeasuredIn sets RateMeasuredIn field to given value.

### HasRateMeasuredIn

`func (o *AmdDetailRecord) HasRateMeasuredIn() bool`

HasRateMeasuredIn returns a boolean if a field has been set.

### GetCost

`func (o *AmdDetailRecord) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AmdDetailRecord) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AmdDetailRecord) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AmdDetailRecord) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *AmdDetailRecord) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AmdDetailRecord) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AmdDetailRecord) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AmdDetailRecord) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRecordType

`func (o *AmdDetailRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *AmdDetailRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *AmdDetailRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



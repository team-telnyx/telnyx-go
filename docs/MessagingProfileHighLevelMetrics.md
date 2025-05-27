# MessagingProfileHighLevelMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**MessagingProfileId** | Pointer to **string** | Identifies the type of resource. | [optional] [readonly] 
**PhoneNumbers** | Pointer to **float32** | The number of phone numbers associated with the messaging profile. | [optional] [readonly] 
**Outbound** | Pointer to [**MessagingProfileHighLevelMetricsOutbound**](MessagingProfileHighLevelMetricsOutbound.md) |  | [optional] 
**Inbound** | Pointer to [**MessagingProfileHighLevelMetricsInbound**](MessagingProfileHighLevelMetricsInbound.md) |  | [optional] 

## Methods

### NewMessagingProfileHighLevelMetrics

`func NewMessagingProfileHighLevelMetrics() *MessagingProfileHighLevelMetrics`

NewMessagingProfileHighLevelMetrics instantiates a new MessagingProfileHighLevelMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingProfileHighLevelMetricsWithDefaults

`func NewMessagingProfileHighLevelMetricsWithDefaults() *MessagingProfileHighLevelMetrics`

NewMessagingProfileHighLevelMetricsWithDefaults instantiates a new MessagingProfileHighLevelMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *MessagingProfileHighLevelMetrics) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *MessagingProfileHighLevelMetrics) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *MessagingProfileHighLevelMetrics) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *MessagingProfileHighLevelMetrics) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *MessagingProfileHighLevelMetrics) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *MessagingProfileHighLevelMetrics) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *MessagingProfileHighLevelMetrics) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *MessagingProfileHighLevelMetrics) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *MessagingProfileHighLevelMetrics) GetPhoneNumbers() float32`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *MessagingProfileHighLevelMetrics) GetPhoneNumbersOk() (*float32, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *MessagingProfileHighLevelMetrics) SetPhoneNumbers(v float32)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *MessagingProfileHighLevelMetrics) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetOutbound

`func (o *MessagingProfileHighLevelMetrics) GetOutbound() MessagingProfileHighLevelMetricsOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *MessagingProfileHighLevelMetrics) GetOutboundOk() (*MessagingProfileHighLevelMetricsOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *MessagingProfileHighLevelMetrics) SetOutbound(v MessagingProfileHighLevelMetricsOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *MessagingProfileHighLevelMetrics) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

### GetInbound

`func (o *MessagingProfileHighLevelMetrics) GetInbound() MessagingProfileHighLevelMetricsInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *MessagingProfileHighLevelMetrics) GetInboundOk() (*MessagingProfileHighLevelMetricsInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *MessagingProfileHighLevelMetrics) SetInbound(v MessagingProfileHighLevelMetricsInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *MessagingProfileHighLevelMetrics) HasInbound() bool`

HasInbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



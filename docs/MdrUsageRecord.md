# MdrUsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Connection** | Pointer to **string** |  | [optional] 
**Received** | Pointer to **string** |  | [optional] 
**Delivered** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Parts** | Pointer to **string** |  | [optional] 
**Sent** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**MessageType** | Pointer to **string** |  | [optional] 
**TnType** | Pointer to **string** |  | [optional] 
**CarrierPassthroughFee** | Pointer to **string** |  | [optional] 

## Methods

### NewMdrUsageRecord

`func NewMdrUsageRecord() *MdrUsageRecord`

NewMdrUsageRecord instantiates a new MdrUsageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdrUsageRecordWithDefaults

`func NewMdrUsageRecordWithDefaults() *MdrUsageRecord`

NewMdrUsageRecordWithDefaults instantiates a new MdrUsageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *MdrUsageRecord) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *MdrUsageRecord) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *MdrUsageRecord) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *MdrUsageRecord) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDirection

`func (o *MdrUsageRecord) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MdrUsageRecord) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MdrUsageRecord) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MdrUsageRecord) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProduct

`func (o *MdrUsageRecord) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *MdrUsageRecord) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *MdrUsageRecord) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *MdrUsageRecord) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetConnection

`func (o *MdrUsageRecord) GetConnection() string`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *MdrUsageRecord) GetConnectionOk() (*string, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *MdrUsageRecord) SetConnection(v string)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *MdrUsageRecord) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetReceived

`func (o *MdrUsageRecord) GetReceived() string`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *MdrUsageRecord) GetReceivedOk() (*string, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *MdrUsageRecord) SetReceived(v string)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *MdrUsageRecord) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetDelivered

`func (o *MdrUsageRecord) GetDelivered() string`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *MdrUsageRecord) GetDeliveredOk() (*string, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *MdrUsageRecord) SetDelivered(v string)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *MdrUsageRecord) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetCurrency

`func (o *MdrUsageRecord) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MdrUsageRecord) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MdrUsageRecord) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MdrUsageRecord) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetParts

`func (o *MdrUsageRecord) GetParts() string`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *MdrUsageRecord) GetPartsOk() (*string, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *MdrUsageRecord) SetParts(v string)`

SetParts sets Parts field to given value.

### HasParts

`func (o *MdrUsageRecord) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetSent

`func (o *MdrUsageRecord) GetSent() string`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *MdrUsageRecord) GetSentOk() (*string, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *MdrUsageRecord) SetSent(v string)`

SetSent sets Sent field to given value.

### HasSent

`func (o *MdrUsageRecord) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetProfileId

`func (o *MdrUsageRecord) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *MdrUsageRecord) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *MdrUsageRecord) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *MdrUsageRecord) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetTags

`func (o *MdrUsageRecord) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MdrUsageRecord) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MdrUsageRecord) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MdrUsageRecord) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMessageType

`func (o *MdrUsageRecord) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *MdrUsageRecord) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *MdrUsageRecord) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *MdrUsageRecord) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetTnType

`func (o *MdrUsageRecord) GetTnType() string`

GetTnType returns the TnType field if non-nil, zero value otherwise.

### GetTnTypeOk

`func (o *MdrUsageRecord) GetTnTypeOk() (*string, bool)`

GetTnTypeOk returns a tuple with the TnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTnType

`func (o *MdrUsageRecord) SetTnType(v string)`

SetTnType sets TnType field to given value.

### HasTnType

`func (o *MdrUsageRecord) HasTnType() bool`

HasTnType returns a boolean if a field has been set.

### GetCarrierPassthroughFee

`func (o *MdrUsageRecord) GetCarrierPassthroughFee() string`

GetCarrierPassthroughFee returns the CarrierPassthroughFee field if non-nil, zero value otherwise.

### GetCarrierPassthroughFeeOk

`func (o *MdrUsageRecord) GetCarrierPassthroughFeeOk() (*string, bool)`

GetCarrierPassthroughFeeOk returns a tuple with the CarrierPassthroughFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierPassthroughFee

`func (o *MdrUsageRecord) SetCarrierPassthroughFee(v string)`

SetCarrierPassthroughFee sets CarrierPassthroughFee field to given value.

### HasCarrierPassthroughFee

`func (o *MdrUsageRecord) HasCarrierPassthroughFee() bool`

HasCarrierPassthroughFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



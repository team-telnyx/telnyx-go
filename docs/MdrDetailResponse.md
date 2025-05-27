# MdrDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Message sent time | [optional] 
**ProfileName** | Pointer to **string** | Configured profile name. New profiles can be created and configured on Telnyx portal | [optional] 
**Direction** | Pointer to **string** | Direction of message - inbound or outbound. | [optional] 
**Parts** | Pointer to **float32** | Number of parts this message has. Max number of character is 160. If message contains more characters then that it will be broken down in multiple parts | [optional] 
**Status** | Pointer to **string** | Message status | [optional] 
**Cld** | Pointer to **string** | The destination number for a call, or the callee | [optional] 
**Cli** | Pointer to **string** | The number associated with the person initiating the call, or the caller | [optional] 
**Rate** | Pointer to **string** | Rate applied to the message | [optional] 
**Cost** | Pointer to **string** | Final cost. Cost is calculated as rate * parts | [optional] 
**Currency** | Pointer to **string** | Currency of the rate and cost | [optional] 
**Id** | Pointer to **string** | Id of message detail record | [optional] 
**MessageType** | Pointer to **string** | Type of message | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 

## Methods

### NewMdrDetailResponse

`func NewMdrDetailResponse() *MdrDetailResponse`

NewMdrDetailResponse instantiates a new MdrDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdrDetailResponseWithDefaults

`func NewMdrDetailResponseWithDefaults() *MdrDetailResponse`

NewMdrDetailResponseWithDefaults instantiates a new MdrDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MdrDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MdrDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MdrDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MdrDetailResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetProfileName

`func (o *MdrDetailResponse) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *MdrDetailResponse) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *MdrDetailResponse) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *MdrDetailResponse) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetDirection

`func (o *MdrDetailResponse) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MdrDetailResponse) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MdrDetailResponse) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MdrDetailResponse) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetParts

`func (o *MdrDetailResponse) GetParts() float32`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *MdrDetailResponse) GetPartsOk() (*float32, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *MdrDetailResponse) SetParts(v float32)`

SetParts sets Parts field to given value.

### HasParts

`func (o *MdrDetailResponse) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetStatus

`func (o *MdrDetailResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MdrDetailResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MdrDetailResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MdrDetailResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCld

`func (o *MdrDetailResponse) GetCld() string`

GetCld returns the Cld field if non-nil, zero value otherwise.

### GetCldOk

`func (o *MdrDetailResponse) GetCldOk() (*string, bool)`

GetCldOk returns a tuple with the Cld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCld

`func (o *MdrDetailResponse) SetCld(v string)`

SetCld sets Cld field to given value.

### HasCld

`func (o *MdrDetailResponse) HasCld() bool`

HasCld returns a boolean if a field has been set.

### GetCli

`func (o *MdrDetailResponse) GetCli() string`

GetCli returns the Cli field if non-nil, zero value otherwise.

### GetCliOk

`func (o *MdrDetailResponse) GetCliOk() (*string, bool)`

GetCliOk returns a tuple with the Cli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCli

`func (o *MdrDetailResponse) SetCli(v string)`

SetCli sets Cli field to given value.

### HasCli

`func (o *MdrDetailResponse) HasCli() bool`

HasCli returns a boolean if a field has been set.

### GetRate

`func (o *MdrDetailResponse) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *MdrDetailResponse) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *MdrDetailResponse) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *MdrDetailResponse) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetCost

`func (o *MdrDetailResponse) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *MdrDetailResponse) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *MdrDetailResponse) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *MdrDetailResponse) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *MdrDetailResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MdrDetailResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MdrDetailResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MdrDetailResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *MdrDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MdrDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MdrDetailResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MdrDetailResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessageType

`func (o *MdrDetailResponse) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *MdrDetailResponse) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *MdrDetailResponse) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *MdrDetailResponse) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetRecordType

`func (o *MdrDetailResponse) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *MdrDetailResponse) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *MdrDetailResponse) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *MdrDetailResponse) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



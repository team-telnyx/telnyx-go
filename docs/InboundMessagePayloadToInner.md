# InboundMessagePayloadToInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | Receiving address (+E.164 formatted phone number or short code). | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Carrier** | Pointer to **string** | The carrier of the receiver. | [optional] 
**LineType** | Pointer to **string** | The line-type of the receiver. | [optional] 

## Methods

### NewInboundMessagePayloadToInner

`func NewInboundMessagePayloadToInner() *InboundMessagePayloadToInner`

NewInboundMessagePayloadToInner instantiates a new InboundMessagePayloadToInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundMessagePayloadToInnerWithDefaults

`func NewInboundMessagePayloadToInnerWithDefaults() *InboundMessagePayloadToInner`

NewInboundMessagePayloadToInnerWithDefaults instantiates a new InboundMessagePayloadToInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *InboundMessagePayloadToInner) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *InboundMessagePayloadToInner) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *InboundMessagePayloadToInner) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *InboundMessagePayloadToInner) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *InboundMessagePayloadToInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InboundMessagePayloadToInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InboundMessagePayloadToInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InboundMessagePayloadToInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCarrier

`func (o *InboundMessagePayloadToInner) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *InboundMessagePayloadToInner) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *InboundMessagePayloadToInner) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *InboundMessagePayloadToInner) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetLineType

`func (o *InboundMessagePayloadToInner) GetLineType() string`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *InboundMessagePayloadToInner) GetLineTypeOk() (*string, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *InboundMessagePayloadToInner) SetLineType(v string)`

SetLineType sets LineType field to given value.

### HasLineType

`func (o *InboundMessagePayloadToInner) HasLineType() bool`

HasLineType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



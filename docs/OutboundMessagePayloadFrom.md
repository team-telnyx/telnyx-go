# OutboundMessagePayloadFrom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code). | [optional] 
**Carrier** | Pointer to **string** | The carrier of the receiver. | [optional] 
**LineType** | Pointer to **string** | The line-type of the receiver. | [optional] 

## Methods

### NewOutboundMessagePayloadFrom

`func NewOutboundMessagePayloadFrom() *OutboundMessagePayloadFrom`

NewOutboundMessagePayloadFrom instantiates a new OutboundMessagePayloadFrom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundMessagePayloadFromWithDefaults

`func NewOutboundMessagePayloadFromWithDefaults() *OutboundMessagePayloadFrom`

NewOutboundMessagePayloadFromWithDefaults instantiates a new OutboundMessagePayloadFrom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *OutboundMessagePayloadFrom) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OutboundMessagePayloadFrom) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OutboundMessagePayloadFrom) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OutboundMessagePayloadFrom) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCarrier

`func (o *OutboundMessagePayloadFrom) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *OutboundMessagePayloadFrom) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *OutboundMessagePayloadFrom) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *OutboundMessagePayloadFrom) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetLineType

`func (o *OutboundMessagePayloadFrom) GetLineType() string`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *OutboundMessagePayloadFrom) GetLineTypeOk() (*string, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *OutboundMessagePayloadFrom) SetLineType(v string)`

SetLineType sets LineType field to given value.

### HasLineType

`func (o *OutboundMessagePayloadFrom) HasLineType() bool`

HasLineType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



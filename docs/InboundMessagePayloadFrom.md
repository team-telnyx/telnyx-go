# InboundMessagePayloadFrom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code). | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Carrier** | Pointer to **string** | The carrier of the sender. | [optional] 
**LineType** | Pointer to **string** | The line-type of the sender. | [optional] 

## Methods

### NewInboundMessagePayloadFrom

`func NewInboundMessagePayloadFrom() *InboundMessagePayloadFrom`

NewInboundMessagePayloadFrom instantiates a new InboundMessagePayloadFrom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundMessagePayloadFromWithDefaults

`func NewInboundMessagePayloadFromWithDefaults() *InboundMessagePayloadFrom`

NewInboundMessagePayloadFromWithDefaults instantiates a new InboundMessagePayloadFrom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *InboundMessagePayloadFrom) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *InboundMessagePayloadFrom) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *InboundMessagePayloadFrom) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *InboundMessagePayloadFrom) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *InboundMessagePayloadFrom) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InboundMessagePayloadFrom) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InboundMessagePayloadFrom) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InboundMessagePayloadFrom) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCarrier

`func (o *InboundMessagePayloadFrom) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *InboundMessagePayloadFrom) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *InboundMessagePayloadFrom) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *InboundMessagePayloadFrom) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetLineType

`func (o *InboundMessagePayloadFrom) GetLineType() string`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *InboundMessagePayloadFrom) GetLineTypeOk() (*string, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *InboundMessagePayloadFrom) SetLineType(v string)`

SetLineType sets LineType field to given value.

### HasLineType

`func (o *InboundMessagePayloadFrom) HasLineType() bool`

HasLineType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



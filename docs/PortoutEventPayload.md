# PortoutEventPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the port-out order that have the FOC date changed. | [optional] 
**Status** | Pointer to **string** | The new status of the port out. | [optional] 
**PhoneNumbers** | Pointer to **[]string** | Phone numbers associated with this port-out order | [optional] 
**CarrierName** | Pointer to **string** | Carrier the number will be ported out to | [optional] 
**Spid** | Pointer to **string** | The new carrier SPID. | [optional] 
**UserId** | Pointer to **string** | Identifies the organization that port-out order belongs to. | [optional] 
**SubscriberName** | Pointer to **string** | The name of the port-out&#39;s end user. | [optional] 
**RejectionReason** | Pointer to **string** | The reason why the order is being rejected by the user. If the order is authorized, this field can be left null | [optional] 
**AttemptedPin** | Pointer to **string** | The PIN that was attempted to be used to authorize the port out. | [optional] 
**PortoutId** | Pointer to **string** | Identifies the port-out order that the comment was added to. | [optional] 
**Comment** | Pointer to **string** | The body of the comment. | [optional] 
**FocDate** | Pointer to **time.Time** | ISO 8601 formatted date indicating the new FOC date. | [optional] 

## Methods

### NewPortoutEventPayload

`func NewPortoutEventPayload() *PortoutEventPayload`

NewPortoutEventPayload instantiates a new PortoutEventPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortoutEventPayloadWithDefaults

`func NewPortoutEventPayloadWithDefaults() *PortoutEventPayload`

NewPortoutEventPayloadWithDefaults instantiates a new PortoutEventPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortoutEventPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortoutEventPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortoutEventPayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortoutEventPayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PortoutEventPayload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortoutEventPayload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortoutEventPayload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PortoutEventPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *PortoutEventPayload) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *PortoutEventPayload) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *PortoutEventPayload) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *PortoutEventPayload) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetCarrierName

`func (o *PortoutEventPayload) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *PortoutEventPayload) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *PortoutEventPayload) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *PortoutEventPayload) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetSpid

`func (o *PortoutEventPayload) GetSpid() string`

GetSpid returns the Spid field if non-nil, zero value otherwise.

### GetSpidOk

`func (o *PortoutEventPayload) GetSpidOk() (*string, bool)`

GetSpidOk returns a tuple with the Spid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpid

`func (o *PortoutEventPayload) SetSpid(v string)`

SetSpid sets Spid field to given value.

### HasSpid

`func (o *PortoutEventPayload) HasSpid() bool`

HasSpid returns a boolean if a field has been set.

### GetUserId

`func (o *PortoutEventPayload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PortoutEventPayload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PortoutEventPayload) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PortoutEventPayload) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSubscriberName

`func (o *PortoutEventPayload) GetSubscriberName() string`

GetSubscriberName returns the SubscriberName field if non-nil, zero value otherwise.

### GetSubscriberNameOk

`func (o *PortoutEventPayload) GetSubscriberNameOk() (*string, bool)`

GetSubscriberNameOk returns a tuple with the SubscriberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberName

`func (o *PortoutEventPayload) SetSubscriberName(v string)`

SetSubscriberName sets SubscriberName field to given value.

### HasSubscriberName

`func (o *PortoutEventPayload) HasSubscriberName() bool`

HasSubscriberName returns a boolean if a field has been set.

### GetRejectionReason

`func (o *PortoutEventPayload) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *PortoutEventPayload) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *PortoutEventPayload) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *PortoutEventPayload) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetAttemptedPin

`func (o *PortoutEventPayload) GetAttemptedPin() string`

GetAttemptedPin returns the AttemptedPin field if non-nil, zero value otherwise.

### GetAttemptedPinOk

`func (o *PortoutEventPayload) GetAttemptedPinOk() (*string, bool)`

GetAttemptedPinOk returns a tuple with the AttemptedPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedPin

`func (o *PortoutEventPayload) SetAttemptedPin(v string)`

SetAttemptedPin sets AttemptedPin field to given value.

### HasAttemptedPin

`func (o *PortoutEventPayload) HasAttemptedPin() bool`

HasAttemptedPin returns a boolean if a field has been set.

### GetPortoutId

`func (o *PortoutEventPayload) GetPortoutId() string`

GetPortoutId returns the PortoutId field if non-nil, zero value otherwise.

### GetPortoutIdOk

`func (o *PortoutEventPayload) GetPortoutIdOk() (*string, bool)`

GetPortoutIdOk returns a tuple with the PortoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortoutId

`func (o *PortoutEventPayload) SetPortoutId(v string)`

SetPortoutId sets PortoutId field to given value.

### HasPortoutId

`func (o *PortoutEventPayload) HasPortoutId() bool`

HasPortoutId returns a boolean if a field has been set.

### GetComment

`func (o *PortoutEventPayload) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PortoutEventPayload) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PortoutEventPayload) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PortoutEventPayload) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetFocDate

`func (o *PortoutEventPayload) GetFocDate() time.Time`

GetFocDate returns the FocDate field if non-nil, zero value otherwise.

### GetFocDateOk

`func (o *PortoutEventPayload) GetFocDateOk() (*time.Time, bool)`

GetFocDateOk returns a tuple with the FocDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocDate

`func (o *PortoutEventPayload) SetFocDate(v time.Time)`

SetFocDate sets FocDate field to given value.

### HasFocDate

`func (o *PortoutEventPayload) HasFocDate() bool`

HasFocDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



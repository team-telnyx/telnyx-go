# CreatePortingPhoneNumberExtensionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortingPhoneNumberId** | **string** | Identifies the porting phone number associated with this porting phone number extension. | 
**ExtensionRange** | [**CreatePortingPhoneNumberExtensionRequestExtensionRange**](CreatePortingPhoneNumberExtensionRequestExtensionRange.md) |  | 
**ActivationRanges** | [**[]CreatePortingPhoneNumberExtensionRequestActivationRangesInner**](CreatePortingPhoneNumberExtensionRequestActivationRangesInner.md) | Specifies the activation ranges for this porting phone number extension. The activation range must be within the extension range and should not overlap with other activation ranges. | 

## Methods

### NewCreatePortingPhoneNumberExtensionRequest

`func NewCreatePortingPhoneNumberExtensionRequest(portingPhoneNumberId string, extensionRange CreatePortingPhoneNumberExtensionRequestExtensionRange, activationRanges []CreatePortingPhoneNumberExtensionRequestActivationRangesInner, ) *CreatePortingPhoneNumberExtensionRequest`

NewCreatePortingPhoneNumberExtensionRequest instantiates a new CreatePortingPhoneNumberExtensionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePortingPhoneNumberExtensionRequestWithDefaults

`func NewCreatePortingPhoneNumberExtensionRequestWithDefaults() *CreatePortingPhoneNumberExtensionRequest`

NewCreatePortingPhoneNumberExtensionRequestWithDefaults instantiates a new CreatePortingPhoneNumberExtensionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortingPhoneNumberId

`func (o *CreatePortingPhoneNumberExtensionRequest) GetPortingPhoneNumberId() string`

GetPortingPhoneNumberId returns the PortingPhoneNumberId field if non-nil, zero value otherwise.

### GetPortingPhoneNumberIdOk

`func (o *CreatePortingPhoneNumberExtensionRequest) GetPortingPhoneNumberIdOk() (*string, bool)`

GetPortingPhoneNumberIdOk returns a tuple with the PortingPhoneNumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingPhoneNumberId

`func (o *CreatePortingPhoneNumberExtensionRequest) SetPortingPhoneNumberId(v string)`

SetPortingPhoneNumberId sets PortingPhoneNumberId field to given value.


### GetExtensionRange

`func (o *CreatePortingPhoneNumberExtensionRequest) GetExtensionRange() CreatePortingPhoneNumberExtensionRequestExtensionRange`

GetExtensionRange returns the ExtensionRange field if non-nil, zero value otherwise.

### GetExtensionRangeOk

`func (o *CreatePortingPhoneNumberExtensionRequest) GetExtensionRangeOk() (*CreatePortingPhoneNumberExtensionRequestExtensionRange, bool)`

GetExtensionRangeOk returns a tuple with the ExtensionRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionRange

`func (o *CreatePortingPhoneNumberExtensionRequest) SetExtensionRange(v CreatePortingPhoneNumberExtensionRequestExtensionRange)`

SetExtensionRange sets ExtensionRange field to given value.


### GetActivationRanges

`func (o *CreatePortingPhoneNumberExtensionRequest) GetActivationRanges() []CreatePortingPhoneNumberExtensionRequestActivationRangesInner`

GetActivationRanges returns the ActivationRanges field if non-nil, zero value otherwise.

### GetActivationRangesOk

`func (o *CreatePortingPhoneNumberExtensionRequest) GetActivationRangesOk() (*[]CreatePortingPhoneNumberExtensionRequestActivationRangesInner, bool)`

GetActivationRangesOk returns a tuple with the ActivationRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationRanges

`func (o *CreatePortingPhoneNumberExtensionRequest) SetActivationRanges(v []CreatePortingPhoneNumberExtensionRequestActivationRangesInner)`

SetActivationRanges sets ActivationRanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



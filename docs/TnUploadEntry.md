# TnUploadEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberId** | Pointer to **string** | Uniquely identifies the resource. | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number in E164 format. | [optional] 
**Status** | Pointer to **string** | Represents the status of the phone number entry upload on Microsoft Teams. | [optional] [default to "pending_upload"]
**ErrorCode** | Pointer to **string** | A code returned by Microsoft Teams if there is an error with the phone number entry upload. | [optional] 
**ErrorMessage** | Pointer to **string** | A message returned by Microsoft Teams if there is an error with the upload process. | [optional] 
**CivicAddressId** | Pointer to **string** | Identifies the civic address assigned to the phone number entry. | [optional] 
**LocationId** | Pointer to **string** | Identifies the location assigned to the phone number entry. | [optional] 
**InternalStatus** | Pointer to **string** | Represents the status of the phone number entry upload on Telnyx. | [optional] [default to "pending_assignment"]

## Methods

### NewTnUploadEntry

`func NewTnUploadEntry() *TnUploadEntry`

NewTnUploadEntry instantiates a new TnUploadEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTnUploadEntryWithDefaults

`func NewTnUploadEntryWithDefaults() *TnUploadEntry`

NewTnUploadEntryWithDefaults instantiates a new TnUploadEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberId

`func (o *TnUploadEntry) GetNumberId() string`

GetNumberId returns the NumberId field if non-nil, zero value otherwise.

### GetNumberIdOk

`func (o *TnUploadEntry) GetNumberIdOk() (*string, bool)`

GetNumberIdOk returns a tuple with the NumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberId

`func (o *TnUploadEntry) SetNumberId(v string)`

SetNumberId sets NumberId field to given value.

### HasNumberId

`func (o *TnUploadEntry) HasNumberId() bool`

HasNumberId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *TnUploadEntry) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *TnUploadEntry) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *TnUploadEntry) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *TnUploadEntry) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *TnUploadEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TnUploadEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TnUploadEntry) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TnUploadEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *TnUploadEntry) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TnUploadEntry) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TnUploadEntry) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *TnUploadEntry) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TnUploadEntry) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TnUploadEntry) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TnUploadEntry) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TnUploadEntry) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCivicAddressId

`func (o *TnUploadEntry) GetCivicAddressId() string`

GetCivicAddressId returns the CivicAddressId field if non-nil, zero value otherwise.

### GetCivicAddressIdOk

`func (o *TnUploadEntry) GetCivicAddressIdOk() (*string, bool)`

GetCivicAddressIdOk returns a tuple with the CivicAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddressId

`func (o *TnUploadEntry) SetCivicAddressId(v string)`

SetCivicAddressId sets CivicAddressId field to given value.

### HasCivicAddressId

`func (o *TnUploadEntry) HasCivicAddressId() bool`

HasCivicAddressId returns a boolean if a field has been set.

### GetLocationId

`func (o *TnUploadEntry) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *TnUploadEntry) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *TnUploadEntry) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *TnUploadEntry) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetInternalStatus

`func (o *TnUploadEntry) GetInternalStatus() string`

GetInternalStatus returns the InternalStatus field if non-nil, zero value otherwise.

### GetInternalStatusOk

`func (o *TnUploadEntry) GetInternalStatusOk() (*string, bool)`

GetInternalStatusOk returns a tuple with the InternalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalStatus

`func (o *TnUploadEntry) SetInternalStatus(v string)`

SetInternalStatus sets InternalStatus field to given value.

### HasInternalStatus

`func (o *TnUploadEntry) HasInternalStatus() bool`

HasInternalStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TnReleaseEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | Phone number in E164 format. | [optional] 
**NumberId** | Pointer to **string** | Phone number ID from the Telnyx API. | [optional] 

## Methods

### NewTnReleaseEntry

`func NewTnReleaseEntry() *TnReleaseEntry`

NewTnReleaseEntry instantiates a new TnReleaseEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTnReleaseEntryWithDefaults

`func NewTnReleaseEntryWithDefaults() *TnReleaseEntry`

NewTnReleaseEntryWithDefaults instantiates a new TnReleaseEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *TnReleaseEntry) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *TnReleaseEntry) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *TnReleaseEntry) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *TnReleaseEntry) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetNumberId

`func (o *TnReleaseEntry) GetNumberId() string`

GetNumberId returns the NumberId field if non-nil, zero value otherwise.

### GetNumberIdOk

`func (o *TnReleaseEntry) GetNumberIdOk() (*string, bool)`

GetNumberIdOk returns a tuple with the NumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberId

`func (o *TnReleaseEntry) SetNumberId(v string)`

SetNumberId sets NumberId field to given value.

### HasNumberId

`func (o *TnReleaseEntry) HasNumberId() bool`

HasNumberId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



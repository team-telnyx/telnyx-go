# Release

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TicketId** | Pointer to **string** | Uniquely identifies the resource. | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Represents the status of the release on Microsoft Teams. | [optional] [default to "pending_upload"]
**ErrorMessage** | Pointer to **string** | A message set if there is an error with the upload process. | [optional] 
**TelephoneNumbers** | Pointer to [**[]TnReleaseEntry**](TnReleaseEntry.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 

## Methods

### NewRelease

`func NewRelease() *Release`

NewRelease instantiates a new Release object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseWithDefaults

`func NewReleaseWithDefaults() *Release`

NewReleaseWithDefaults instantiates a new Release object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicketId

`func (o *Release) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *Release) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *Release) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *Release) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### GetTenantId

`func (o *Release) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Release) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Release) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Release) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetStatus

`func (o *Release) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Release) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Release) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Release) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Release) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Release) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Release) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Release) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetTelephoneNumbers

`func (o *Release) GetTelephoneNumbers() []TnReleaseEntry`

GetTelephoneNumbers returns the TelephoneNumbers field if non-nil, zero value otherwise.

### GetTelephoneNumbersOk

`func (o *Release) GetTelephoneNumbersOk() (*[]TnReleaseEntry, bool)`

GetTelephoneNumbersOk returns a tuple with the TelephoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumbers

`func (o *Release) SetTelephoneNumbers(v []TnReleaseEntry)`

SetTelephoneNumbers sets TelephoneNumbers field to given value.

### HasTelephoneNumbers

`func (o *Release) HasTelephoneNumbers() bool`

HasTelephoneNumbers returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Release) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Release) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Release) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Release) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Upload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TicketId** | Pointer to **string** | Uniquely identifies the resource. | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Represents the status of the upload on Microsoft Teams. | [optional] [default to "pending_upload"]
**AvailableUsages** | Pointer to **[]string** |  | [optional] 
**ErrorCode** | Pointer to **string** | A code returned by Microsoft Teams if there is an error with the upload process. | [optional] 
**ErrorMessage** | Pointer to **string** | A message set if there is an error with the upload process. | [optional] 
**TnUploadEntries** | Pointer to [**[]TnUploadEntry**](TnUploadEntry.md) |  | [optional] 

## Methods

### NewUpload

`func NewUpload() *Upload`

NewUpload instantiates a new Upload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadWithDefaults

`func NewUploadWithDefaults() *Upload`

NewUploadWithDefaults instantiates a new Upload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicketId

`func (o *Upload) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *Upload) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *Upload) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *Upload) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### GetTenantId

`func (o *Upload) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Upload) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Upload) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Upload) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetLocationId

`func (o *Upload) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Upload) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Upload) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *Upload) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetStatus

`func (o *Upload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Upload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Upload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Upload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAvailableUsages

`func (o *Upload) GetAvailableUsages() []string`

GetAvailableUsages returns the AvailableUsages field if non-nil, zero value otherwise.

### GetAvailableUsagesOk

`func (o *Upload) GetAvailableUsagesOk() (*[]string, bool)`

GetAvailableUsagesOk returns a tuple with the AvailableUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUsages

`func (o *Upload) SetAvailableUsages(v []string)`

SetAvailableUsages sets AvailableUsages field to given value.

### HasAvailableUsages

`func (o *Upload) HasAvailableUsages() bool`

HasAvailableUsages returns a boolean if a field has been set.

### GetErrorCode

`func (o *Upload) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Upload) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Upload) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Upload) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Upload) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Upload) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Upload) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Upload) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetTnUploadEntries

`func (o *Upload) GetTnUploadEntries() []TnUploadEntry`

GetTnUploadEntries returns the TnUploadEntries field if non-nil, zero value otherwise.

### GetTnUploadEntriesOk

`func (o *Upload) GetTnUploadEntriesOk() (*[]TnUploadEntry, bool)`

GetTnUploadEntriesOk returns a tuple with the TnUploadEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTnUploadEntries

`func (o *Upload) SetTnUploadEntries(v []TnUploadEntry)`

SetTnUploadEntries sets TnUploadEntries field to given value.

### HasTnUploadEntries

`func (o *Upload) HasTnUploadEntries() bool`

HasTnUploadEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



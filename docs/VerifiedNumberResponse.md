# VerifiedNumberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to [**VerifiedNumberRecordType**](VerifiedNumberRecordType.md) |  | [optional] 
**VerifiedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewVerifiedNumberResponse

`func NewVerifiedNumberResponse() *VerifiedNumberResponse`

NewVerifiedNumberResponse instantiates a new VerifiedNumberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifiedNumberResponseWithDefaults

`func NewVerifiedNumberResponseWithDefaults() *VerifiedNumberResponse`

NewVerifiedNumberResponseWithDefaults instantiates a new VerifiedNumberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *VerifiedNumberResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *VerifiedNumberResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *VerifiedNumberResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *VerifiedNumberResponse) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetRecordType

`func (o *VerifiedNumberResponse) GetRecordType() VerifiedNumberRecordType`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *VerifiedNumberResponse) GetRecordTypeOk() (*VerifiedNumberRecordType, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *VerifiedNumberResponse) SetRecordType(v VerifiedNumberRecordType)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *VerifiedNumberResponse) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *VerifiedNumberResponse) GetVerifiedAt() string`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *VerifiedNumberResponse) GetVerifiedAtOk() (*string, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *VerifiedNumberResponse) SetVerifiedAt(v string)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *VerifiedNumberResponse) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



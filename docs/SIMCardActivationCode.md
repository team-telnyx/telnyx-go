# SIMCardActivationCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**ActivationCode** | Pointer to **string** | Contents of the eSIM activation QR code. | [optional] [readonly] 

## Methods

### NewSIMCardActivationCode

`func NewSIMCardActivationCode() *SIMCardActivationCode`

NewSIMCardActivationCode instantiates a new SIMCardActivationCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardActivationCodeWithDefaults

`func NewSIMCardActivationCodeWithDefaults() *SIMCardActivationCode`

NewSIMCardActivationCodeWithDefaults instantiates a new SIMCardActivationCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *SIMCardActivationCode) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SIMCardActivationCode) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SIMCardActivationCode) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SIMCardActivationCode) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetActivationCode

`func (o *SIMCardActivationCode) GetActivationCode() string`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *SIMCardActivationCode) GetActivationCodeOk() (*string, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *SIMCardActivationCode) SetActivationCode(v string)`

SetActivationCode sets ActivationCode field to given value.

### HasActivationCode

`func (o *SIMCardActivationCode) HasActivationCode() bool`

HasActivationCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



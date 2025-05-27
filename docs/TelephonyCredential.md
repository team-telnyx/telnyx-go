# TelephonyCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **string** | Identifies the resource this credential is associated with. | [optional] 
**Expired** | Pointer to **bool** | Defaults to false | [optional] 
**SipUsername** | Pointer to **string** | The randomly generated SIP username for the credential. | [optional] 
**SipPassword** | Pointer to **string** | The randomly generated SIP password for the credential. | [optional] 
**CreatedAt** | Pointer to **string** | ISO-8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO-8601 formatted date indicating when the resource was updated. | [optional] 
**ExpiresAt** | Pointer to **string** | ISO-8601 formatted date indicating when the resource will expire. | [optional] 

## Methods

### NewTelephonyCredential

`func NewTelephonyCredential() *TelephonyCredential`

NewTelephonyCredential instantiates a new TelephonyCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephonyCredentialWithDefaults

`func NewTelephonyCredentialWithDefaults() *TelephonyCredential`

NewTelephonyCredentialWithDefaults instantiates a new TelephonyCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TelephonyCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TelephonyCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TelephonyCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TelephonyCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *TelephonyCredential) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *TelephonyCredential) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *TelephonyCredential) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *TelephonyCredential) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetName

`func (o *TelephonyCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelephonyCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelephonyCredential) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelephonyCredential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceId

`func (o *TelephonyCredential) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *TelephonyCredential) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *TelephonyCredential) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *TelephonyCredential) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetExpired

`func (o *TelephonyCredential) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *TelephonyCredential) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *TelephonyCredential) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *TelephonyCredential) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetSipUsername

`func (o *TelephonyCredential) GetSipUsername() string`

GetSipUsername returns the SipUsername field if non-nil, zero value otherwise.

### GetSipUsernameOk

`func (o *TelephonyCredential) GetSipUsernameOk() (*string, bool)`

GetSipUsernameOk returns a tuple with the SipUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipUsername

`func (o *TelephonyCredential) SetSipUsername(v string)`

SetSipUsername sets SipUsername field to given value.

### HasSipUsername

`func (o *TelephonyCredential) HasSipUsername() bool`

HasSipUsername returns a boolean if a field has been set.

### GetSipPassword

`func (o *TelephonyCredential) GetSipPassword() string`

GetSipPassword returns the SipPassword field if non-nil, zero value otherwise.

### GetSipPasswordOk

`func (o *TelephonyCredential) GetSipPasswordOk() (*string, bool)`

GetSipPasswordOk returns a tuple with the SipPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipPassword

`func (o *TelephonyCredential) SetSipPassword(v string)`

SetSipPassword sets SipPassword field to given value.

### HasSipPassword

`func (o *TelephonyCredential) HasSipPassword() bool`

HasSipPassword returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TelephonyCredential) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TelephonyCredential) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TelephonyCredential) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TelephonyCredential) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TelephonyCredential) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TelephonyCredential) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TelephonyCredential) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TelephonyCredential) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TelephonyCredential) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TelephonyCredential) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TelephonyCredential) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TelephonyCredential) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



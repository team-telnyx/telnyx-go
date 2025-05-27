# PushCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of a push credential | 
**Certificate** | **string** | Apple certificate for sending push notifications. For iOS only | 
**PrivateKey** | **string** | Apple private key for a given certificate for sending push notifications. For iOS only | 
**ProjectAccountJsonFile** | **map[string]interface{}** | Google server key for sending push notifications. For Android only | 
**Alias** | **string** | Alias to uniquely identify a credential | 
**Type** | **string** | Type of mobile push credential. Either &lt;code&gt;ios&lt;/code&gt; or &lt;code&gt;android&lt;/code&gt; | 
**RecordType** | **string** |  | [readonly] 
**CreatedAt** | **string** | ISO 8601 timestamp when the room was created | 
**UpdatedAt** | **string** | ISO 8601 timestamp when the room was updated. | 

## Methods

### NewPushCredential

`func NewPushCredential(id string, certificate string, privateKey string, projectAccountJsonFile map[string]interface{}, alias string, type_ string, recordType string, createdAt string, updatedAt string, ) *PushCredential`

NewPushCredential instantiates a new PushCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushCredentialWithDefaults

`func NewPushCredentialWithDefaults() *PushCredential`

NewPushCredentialWithDefaults instantiates a new PushCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PushCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PushCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PushCredential) SetId(v string)`

SetId sets Id field to given value.


### GetCertificate

`func (o *PushCredential) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PushCredential) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PushCredential) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetPrivateKey

`func (o *PushCredential) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PushCredential) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PushCredential) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetProjectAccountJsonFile

`func (o *PushCredential) GetProjectAccountJsonFile() map[string]interface{}`

GetProjectAccountJsonFile returns the ProjectAccountJsonFile field if non-nil, zero value otherwise.

### GetProjectAccountJsonFileOk

`func (o *PushCredential) GetProjectAccountJsonFileOk() (*map[string]interface{}, bool)`

GetProjectAccountJsonFileOk returns a tuple with the ProjectAccountJsonFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAccountJsonFile

`func (o *PushCredential) SetProjectAccountJsonFile(v map[string]interface{})`

SetProjectAccountJsonFile sets ProjectAccountJsonFile field to given value.


### GetAlias

`func (o *PushCredential) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PushCredential) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PushCredential) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetType

`func (o *PushCredential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PushCredential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PushCredential) SetType(v string)`

SetType sets Type field to given value.


### GetRecordType

`func (o *PushCredential) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PushCredential) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PushCredential) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetCreatedAt

`func (o *PushCredential) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PushCredential) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PushCredential) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PushCredential) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PushCredential) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PushCredential) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



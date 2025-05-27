# AuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the audit log entry. | [optional] 
**UserId** | Pointer to **string** | Unique identifier for the user who made the change. | [optional] 
**RecordType** | Pointer to **string** | The type of the resource being audited. | [optional] 
**ResourceId** | Pointer to **string** | Unique identifier for the resource that was changed. | [optional] 
**AlternateResourceId** | Pointer to **NullableString** | An alternate identifier for a resource which may be considered unique enough to identify the resource but is not the primary identifier for the resource. For example, this field could be used to store the phone number value for a phone number when the primary database identifier is a separate distinct value. | [optional] 
**ChangeMadeBy** | Pointer to **string** | Indicates if the change was made by Telnyx on your behalf, the organization owner, a member of your organization, or in the case of managed accounts, the account manager. | [optional] 
**Changes** | Pointer to [**[]AuditEventChanges**](AuditEventChanges.md) | Details of the changes made to the resource. | [optional] 
**OrganizationId** | Pointer to **string** | Unique identifier for the organization that owns the resource. | [optional] 
**ChangeType** | Pointer to **string** | The type of change that occurred. | [optional] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the change occurred. | [optional] 

## Methods

### NewAuditLog

`func NewAuditLog() *AuditLog`

NewAuditLog instantiates a new AuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogWithDefaults

`func NewAuditLogWithDefaults() *AuditLog`

NewAuditLogWithDefaults instantiates a new AuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *AuditLog) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuditLog) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuditLog) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuditLog) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRecordType

`func (o *AuditLog) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *AuditLog) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *AuditLog) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *AuditLog) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetResourceId

`func (o *AuditLog) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AuditLog) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AuditLog) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AuditLog) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetAlternateResourceId

`func (o *AuditLog) GetAlternateResourceId() string`

GetAlternateResourceId returns the AlternateResourceId field if non-nil, zero value otherwise.

### GetAlternateResourceIdOk

`func (o *AuditLog) GetAlternateResourceIdOk() (*string, bool)`

GetAlternateResourceIdOk returns a tuple with the AlternateResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateResourceId

`func (o *AuditLog) SetAlternateResourceId(v string)`

SetAlternateResourceId sets AlternateResourceId field to given value.

### HasAlternateResourceId

`func (o *AuditLog) HasAlternateResourceId() bool`

HasAlternateResourceId returns a boolean if a field has been set.

### SetAlternateResourceIdNil

`func (o *AuditLog) SetAlternateResourceIdNil(b bool)`

 SetAlternateResourceIdNil sets the value for AlternateResourceId to be an explicit nil

### UnsetAlternateResourceId
`func (o *AuditLog) UnsetAlternateResourceId()`

UnsetAlternateResourceId ensures that no value is present for AlternateResourceId, not even an explicit nil
### GetChangeMadeBy

`func (o *AuditLog) GetChangeMadeBy() string`

GetChangeMadeBy returns the ChangeMadeBy field if non-nil, zero value otherwise.

### GetChangeMadeByOk

`func (o *AuditLog) GetChangeMadeByOk() (*string, bool)`

GetChangeMadeByOk returns a tuple with the ChangeMadeBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeMadeBy

`func (o *AuditLog) SetChangeMadeBy(v string)`

SetChangeMadeBy sets ChangeMadeBy field to given value.

### HasChangeMadeBy

`func (o *AuditLog) HasChangeMadeBy() bool`

HasChangeMadeBy returns a boolean if a field has been set.

### GetChanges

`func (o *AuditLog) GetChanges() []AuditEventChanges`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *AuditLog) GetChangesOk() (*[]AuditEventChanges, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *AuditLog) SetChanges(v []AuditEventChanges)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *AuditLog) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### SetChangesNil

`func (o *AuditLog) SetChangesNil(b bool)`

 SetChangesNil sets the value for Changes to be an explicit nil

### UnsetChanges
`func (o *AuditLog) UnsetChanges()`

UnsetChanges ensures that no value is present for Changes, not even an explicit nil
### GetOrganizationId

`func (o *AuditLog) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AuditLog) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AuditLog) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AuditLog) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetChangeType

`func (o *AuditLog) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *AuditLog) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *AuditLog) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *AuditLog) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuditLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuditLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



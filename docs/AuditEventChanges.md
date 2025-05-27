# AuditEventChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | The name of the field that was changed. May use the dot notation to indicate nested fields. | [optional] 
**To** | Pointer to [**AuditEventChangesTo**](AuditEventChangesTo.md) |  | [optional] 
**From** | Pointer to [**AuditEventChangesFrom**](AuditEventChangesFrom.md) |  | [optional] 

## Methods

### NewAuditEventChanges

`func NewAuditEventChanges() *AuditEventChanges`

NewAuditEventChanges instantiates a new AuditEventChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventChangesWithDefaults

`func NewAuditEventChangesWithDefaults() *AuditEventChanges`

NewAuditEventChangesWithDefaults instantiates a new AuditEventChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AuditEventChanges) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AuditEventChanges) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AuditEventChanges) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *AuditEventChanges) HasField() bool`

HasField returns a boolean if a field has been set.

### GetTo

`func (o *AuditEventChanges) GetTo() AuditEventChangesTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *AuditEventChanges) GetToOk() (*AuditEventChangesTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *AuditEventChanges) SetTo(v AuditEventChangesTo)`

SetTo sets To field to given value.

### HasTo

`func (o *AuditEventChanges) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *AuditEventChanges) GetFrom() AuditEventChangesFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *AuditEventChanges) GetFromOk() (*AuditEventChangesFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *AuditEventChanges) SetFrom(v AuditEventChangesFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *AuditEventChanges) HasFrom() bool`

HasFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BulkSIMCardAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**ActionType** | Pointer to **string** | The operation type. It can be one of the following: &lt;br/&gt; &lt;ul&gt; &lt;li&gt;&lt;code&gt;bulk_set_public_ips&lt;/code&gt; - set a public IP for each specified SIM card.&lt;/li&gt; &lt;/ul&gt; | [optional] [readonly] 
**Settings** | Pointer to **map[string]interface{}** | A JSON object representation of the bulk action payload. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewBulkSIMCardAction

`func NewBulkSIMCardAction() *BulkSIMCardAction`

NewBulkSIMCardAction instantiates a new BulkSIMCardAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkSIMCardActionWithDefaults

`func NewBulkSIMCardActionWithDefaults() *BulkSIMCardAction`

NewBulkSIMCardActionWithDefaults instantiates a new BulkSIMCardAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkSIMCardAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkSIMCardAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkSIMCardAction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkSIMCardAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *BulkSIMCardAction) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *BulkSIMCardAction) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *BulkSIMCardAction) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *BulkSIMCardAction) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetActionType

`func (o *BulkSIMCardAction) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *BulkSIMCardAction) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *BulkSIMCardAction) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *BulkSIMCardAction) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetSettings

`func (o *BulkSIMCardAction) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *BulkSIMCardAction) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *BulkSIMCardAction) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *BulkSIMCardAction) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BulkSIMCardAction) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BulkSIMCardAction) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BulkSIMCardAction) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BulkSIMCardAction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BulkSIMCardAction) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BulkSIMCardAction) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BulkSIMCardAction) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BulkSIMCardAction) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



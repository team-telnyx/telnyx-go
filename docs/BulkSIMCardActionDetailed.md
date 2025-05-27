# BulkSIMCardActionDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**ActionType** | Pointer to **string** | The operation type. It can be one of the following: &lt;br/&gt; &lt;ul&gt; &lt;li&gt;&lt;code&gt;bulk_set_public_ips&lt;/code&gt; - set a public IP for each specified SIM card.&lt;/li&gt; &lt;/ul&gt; | [optional] [readonly] 
**Settings** | Pointer to **map[string]interface{}** | A JSON object representation of the bulk action payload. | [optional] [readonly] 
**SimCardActionsSummary** | Pointer to [**[]SIMCardActionsSummary**](SIMCardActionsSummary.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewBulkSIMCardActionDetailed

`func NewBulkSIMCardActionDetailed() *BulkSIMCardActionDetailed`

NewBulkSIMCardActionDetailed instantiates a new BulkSIMCardActionDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkSIMCardActionDetailedWithDefaults

`func NewBulkSIMCardActionDetailedWithDefaults() *BulkSIMCardActionDetailed`

NewBulkSIMCardActionDetailedWithDefaults instantiates a new BulkSIMCardActionDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkSIMCardActionDetailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkSIMCardActionDetailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkSIMCardActionDetailed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkSIMCardActionDetailed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *BulkSIMCardActionDetailed) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *BulkSIMCardActionDetailed) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *BulkSIMCardActionDetailed) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *BulkSIMCardActionDetailed) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetActionType

`func (o *BulkSIMCardActionDetailed) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *BulkSIMCardActionDetailed) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *BulkSIMCardActionDetailed) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *BulkSIMCardActionDetailed) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetSettings

`func (o *BulkSIMCardActionDetailed) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *BulkSIMCardActionDetailed) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *BulkSIMCardActionDetailed) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *BulkSIMCardActionDetailed) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSimCardActionsSummary

`func (o *BulkSIMCardActionDetailed) GetSimCardActionsSummary() []SIMCardActionsSummary`

GetSimCardActionsSummary returns the SimCardActionsSummary field if non-nil, zero value otherwise.

### GetSimCardActionsSummaryOk

`func (o *BulkSIMCardActionDetailed) GetSimCardActionsSummaryOk() (*[]SIMCardActionsSummary, bool)`

GetSimCardActionsSummaryOk returns a tuple with the SimCardActionsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardActionsSummary

`func (o *BulkSIMCardActionDetailed) SetSimCardActionsSummary(v []SIMCardActionsSummary)`

SetSimCardActionsSummary sets SimCardActionsSummary field to given value.

### HasSimCardActionsSummary

`func (o *BulkSIMCardActionDetailed) HasSimCardActionsSummary() bool`

HasSimCardActionsSummary returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BulkSIMCardActionDetailed) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BulkSIMCardActionDetailed) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BulkSIMCardActionDetailed) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BulkSIMCardActionDetailed) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BulkSIMCardActionDetailed) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BulkSIMCardActionDetailed) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BulkSIMCardActionDetailed) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BulkSIMCardActionDetailed) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



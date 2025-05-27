# SIMCardAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**SimCardId** | Pointer to **string** | The related SIM card identifier. | [optional] [readonly] 
**ActionType** | Pointer to **string** | The operation type. It can be one of the following: &lt;br/&gt; &lt;ul&gt;  &lt;li&gt;&lt;code&gt;enable&lt;/code&gt; - move the SIM card to the &lt;code&gt;enabled&lt;/code&gt; status&lt;/li&gt;  &lt;li&gt;&lt;code&gt;enable_standby_sim_card&lt;/code&gt; - move a SIM card previously on the &lt;code&gt;standby&lt;/code&gt; status to the &lt;code&gt;enabled&lt;/code&gt; status after it consumes data.&lt;/li&gt;  &lt;li&gt;&lt;code&gt;disable&lt;/code&gt; - move the SIM card to the &lt;code&gt;disabled&lt;/code&gt; status&lt;/li&gt;  &lt;li&gt;&lt;code&gt;set_standby&lt;/code&gt; - move the SIM card to the &lt;code&gt;standby&lt;/code&gt; status&lt;/li&gt;  &lt;/ul&gt; | [optional] [readonly] 
**Status** | Pointer to [**SIMCardActionStatus**](SIMCardActionStatus.md) |  | [optional] 
**Settings** | Pointer to **map[string]interface{}** | A JSON object representation of the action params. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewSIMCardAction

`func NewSIMCardAction() *SIMCardAction`

NewSIMCardAction instantiates a new SIMCardAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardActionWithDefaults

`func NewSIMCardActionWithDefaults() *SIMCardAction`

NewSIMCardActionWithDefaults instantiates a new SIMCardAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SIMCardAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SIMCardAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SIMCardAction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SIMCardAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *SIMCardAction) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SIMCardAction) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SIMCardAction) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SIMCardAction) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetSimCardId

`func (o *SIMCardAction) GetSimCardId() string`

GetSimCardId returns the SimCardId field if non-nil, zero value otherwise.

### GetSimCardIdOk

`func (o *SIMCardAction) GetSimCardIdOk() (*string, bool)`

GetSimCardIdOk returns a tuple with the SimCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardId

`func (o *SIMCardAction) SetSimCardId(v string)`

SetSimCardId sets SimCardId field to given value.

### HasSimCardId

`func (o *SIMCardAction) HasSimCardId() bool`

HasSimCardId returns a boolean if a field has been set.

### GetActionType

`func (o *SIMCardAction) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *SIMCardAction) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *SIMCardAction) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *SIMCardAction) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetStatus

`func (o *SIMCardAction) GetStatus() SIMCardActionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SIMCardAction) GetStatusOk() (*SIMCardActionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SIMCardAction) SetStatus(v SIMCardActionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SIMCardAction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSettings

`func (o *SIMCardAction) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SIMCardAction) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SIMCardAction) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SIMCardAction) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *SIMCardAction) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *SIMCardAction) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetCreatedAt

`func (o *SIMCardAction) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SIMCardAction) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SIMCardAction) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SIMCardAction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SIMCardAction) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SIMCardAction) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SIMCardAction) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SIMCardAction) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



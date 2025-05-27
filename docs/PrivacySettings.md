# PrivacySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataRetention** | Pointer to **bool** | If true, conversation history and insights will be stored. If false, they will not be stored. This in‑tool toggle governs solely the retention of conversation history and insights via the AI assistant. It has no effect on any separate recording, transcription, or storage configuration that you have set at the account, number, or application level. All such external settings remain in force regardless of your selection here. | [optional] 

## Methods

### NewPrivacySettings

`func NewPrivacySettings() *PrivacySettings`

NewPrivacySettings instantiates a new PrivacySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivacySettingsWithDefaults

`func NewPrivacySettingsWithDefaults() *PrivacySettings`

NewPrivacySettingsWithDefaults instantiates a new PrivacySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataRetention

`func (o *PrivacySettings) GetDataRetention() bool`

GetDataRetention returns the DataRetention field if non-nil, zero value otherwise.

### GetDataRetentionOk

`func (o *PrivacySettings) GetDataRetentionOk() (*bool, bool)`

GetDataRetentionOk returns a tuple with the DataRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetention

`func (o *PrivacySettings) SetDataRetention(v bool)`

SetDataRetention sets DataRetention field to given value.

### HasDataRetention

`func (o *PrivacySettings) HasDataRetention() bool`

HasDataRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



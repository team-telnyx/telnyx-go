# MediaStorageDetailRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the Media Storage Event | [optional] 
**CreatedAt** | Pointer to **time.Time** | Event creation time | [optional] 
**AssetId** | Pointer to **string** | Asset id | [optional] 
**UserId** | Pointer to **string** | User id | [optional] 
**OrgId** | Pointer to **string** | Organization owner id | [optional] 
**ActionType** | Pointer to **string** | Type of action performed against the Media Storage API | [optional] 
**LinkChannelType** | Pointer to **string** | Link channel type | [optional] 
**LinkChannelId** | Pointer to **string** | Link channel id | [optional] 
**Status** | Pointer to **string** | Request status | [optional] 
**WebhookId** | Pointer to **string** | Webhook id | [optional] 
**Rate** | Pointer to **string** | Currency amount per billing unit used to calculate the Telnyx billing cost | [optional] 
**RateMeasuredIn** | Pointer to **string** | Billing unit used to calculate the Telnyx billing cost | [optional] 
**Cost** | Pointer to **string** | Currency amount for Telnyx billing cost | [optional] 
**Currency** | Pointer to **string** | Telnyx account currency used to describe monetary values, including billing cost | [optional] 
**RecordType** | **string** |  | [default to "media_storage"]

## Methods

### NewMediaStorageDetailRecord

`func NewMediaStorageDetailRecord(recordType string, ) *MediaStorageDetailRecord`

NewMediaStorageDetailRecord instantiates a new MediaStorageDetailRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaStorageDetailRecordWithDefaults

`func NewMediaStorageDetailRecordWithDefaults() *MediaStorageDetailRecord`

NewMediaStorageDetailRecordWithDefaults instantiates a new MediaStorageDetailRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MediaStorageDetailRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaStorageDetailRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaStorageDetailRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MediaStorageDetailRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MediaStorageDetailRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MediaStorageDetailRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MediaStorageDetailRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MediaStorageDetailRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAssetId

`func (o *MediaStorageDetailRecord) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *MediaStorageDetailRecord) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *MediaStorageDetailRecord) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *MediaStorageDetailRecord) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetUserId

`func (o *MediaStorageDetailRecord) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MediaStorageDetailRecord) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MediaStorageDetailRecord) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MediaStorageDetailRecord) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetOrgId

`func (o *MediaStorageDetailRecord) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MediaStorageDetailRecord) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MediaStorageDetailRecord) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *MediaStorageDetailRecord) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetActionType

`func (o *MediaStorageDetailRecord) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *MediaStorageDetailRecord) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *MediaStorageDetailRecord) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *MediaStorageDetailRecord) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetLinkChannelType

`func (o *MediaStorageDetailRecord) GetLinkChannelType() string`

GetLinkChannelType returns the LinkChannelType field if non-nil, zero value otherwise.

### GetLinkChannelTypeOk

`func (o *MediaStorageDetailRecord) GetLinkChannelTypeOk() (*string, bool)`

GetLinkChannelTypeOk returns a tuple with the LinkChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkChannelType

`func (o *MediaStorageDetailRecord) SetLinkChannelType(v string)`

SetLinkChannelType sets LinkChannelType field to given value.

### HasLinkChannelType

`func (o *MediaStorageDetailRecord) HasLinkChannelType() bool`

HasLinkChannelType returns a boolean if a field has been set.

### GetLinkChannelId

`func (o *MediaStorageDetailRecord) GetLinkChannelId() string`

GetLinkChannelId returns the LinkChannelId field if non-nil, zero value otherwise.

### GetLinkChannelIdOk

`func (o *MediaStorageDetailRecord) GetLinkChannelIdOk() (*string, bool)`

GetLinkChannelIdOk returns a tuple with the LinkChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkChannelId

`func (o *MediaStorageDetailRecord) SetLinkChannelId(v string)`

SetLinkChannelId sets LinkChannelId field to given value.

### HasLinkChannelId

`func (o *MediaStorageDetailRecord) HasLinkChannelId() bool`

HasLinkChannelId returns a boolean if a field has been set.

### GetStatus

`func (o *MediaStorageDetailRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MediaStorageDetailRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MediaStorageDetailRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MediaStorageDetailRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWebhookId

`func (o *MediaStorageDetailRecord) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *MediaStorageDetailRecord) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *MediaStorageDetailRecord) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *MediaStorageDetailRecord) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetRate

`func (o *MediaStorageDetailRecord) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *MediaStorageDetailRecord) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *MediaStorageDetailRecord) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *MediaStorageDetailRecord) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRateMeasuredIn

`func (o *MediaStorageDetailRecord) GetRateMeasuredIn() string`

GetRateMeasuredIn returns the RateMeasuredIn field if non-nil, zero value otherwise.

### GetRateMeasuredInOk

`func (o *MediaStorageDetailRecord) GetRateMeasuredInOk() (*string, bool)`

GetRateMeasuredInOk returns a tuple with the RateMeasuredIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateMeasuredIn

`func (o *MediaStorageDetailRecord) SetRateMeasuredIn(v string)`

SetRateMeasuredIn sets RateMeasuredIn field to given value.

### HasRateMeasuredIn

`func (o *MediaStorageDetailRecord) HasRateMeasuredIn() bool`

HasRateMeasuredIn returns a boolean if a field has been set.

### GetCost

`func (o *MediaStorageDetailRecord) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *MediaStorageDetailRecord) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *MediaStorageDetailRecord) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *MediaStorageDetailRecord) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *MediaStorageDetailRecord) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MediaStorageDetailRecord) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MediaStorageDetailRecord) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MediaStorageDetailRecord) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRecordType

`func (o *MediaStorageDetailRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *MediaStorageDetailRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *MediaStorageDetailRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



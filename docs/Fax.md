# Fax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**ConnectionId** | Pointer to **string** | The ID of the connection used to send the fax. | [optional] 
**Direction** | Pointer to [**Direction**](Direction.md) |  | [optional] 
**MediaUrl** | Pointer to **string** | The URL (or list of URLs) to the PDF used for the fax&#39;s media. media_url and media_name/contents can&#39;t be submitted together. | [optional] 
**MediaName** | Pointer to **string** | The media_name used for the fax&#39;s media. Must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. media_name and media_url/contents can&#39;t be submitted together. | [optional] 
**To** | Pointer to **string** | The phone number, in E.164 format, the fax will be sent to or SIP URI | [optional] 
**From** | Pointer to **string** | The phone number, in E.164 format, the fax will be sent from. | [optional] 
**FromDisplayName** | Pointer to **string** | The string used as the caller id name (SIP From Display Name) presented to the destination (&#x60;to&#x60; number). | [optional] 
**Quality** | Pointer to [**Quality**](Quality.md) |  | [optional] [default to HIGH]
**Status** | Pointer to **string** | Status of the fax | [optional] 
**WebhookUrl** | Pointer to **string** | URL that will receive fax webhooks | [optional] 
**WebhookFailoverUrl** | Pointer to **string** | Optional failover URL that will receive fax webhooks if webhook_url doesn&#39;t return a 2XX response | [optional] 
**StoreMedia** | Pointer to **bool** | Should fax media be stored on temporary URL. It does not support media_name. | [optional] 
**StoredMediaUrl** | Pointer to **string** | If store_media was set to true, this is a link to temporary location. Link expires after 10 minutes. | [optional] 
**PreviewUrl** | Pointer to **string** | If &#x60;store_preview&#x60; was set to &#x60;true&#x60;, this is a link to temporary location. Link expires after 10 minutes. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 timestamp when resource was created | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 timestamp when resource was updated | [optional] 

## Methods

### NewFax

`func NewFax() *Fax`

NewFax instantiates a new Fax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaxWithDefaults

`func NewFaxWithDefaults() *Fax`

NewFaxWithDefaults instantiates a new Fax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *Fax) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *Fax) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *Fax) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *Fax) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *Fax) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Fax) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Fax) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Fax) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConnectionId

`func (o *Fax) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Fax) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Fax) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Fax) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetDirection

`func (o *Fax) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Fax) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Fax) SetDirection(v Direction)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Fax) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetMediaUrl

`func (o *Fax) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *Fax) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *Fax) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *Fax) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### GetMediaName

`func (o *Fax) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *Fax) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *Fax) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *Fax) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.

### GetTo

`func (o *Fax) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Fax) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Fax) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *Fax) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *Fax) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Fax) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Fax) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Fax) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetFromDisplayName

`func (o *Fax) GetFromDisplayName() string`

GetFromDisplayName returns the FromDisplayName field if non-nil, zero value otherwise.

### GetFromDisplayNameOk

`func (o *Fax) GetFromDisplayNameOk() (*string, bool)`

GetFromDisplayNameOk returns a tuple with the FromDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDisplayName

`func (o *Fax) SetFromDisplayName(v string)`

SetFromDisplayName sets FromDisplayName field to given value.

### HasFromDisplayName

`func (o *Fax) HasFromDisplayName() bool`

HasFromDisplayName returns a boolean if a field has been set.

### GetQuality

`func (o *Fax) GetQuality() Quality`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *Fax) GetQualityOk() (*Quality, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *Fax) SetQuality(v Quality)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *Fax) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetStatus

`func (o *Fax) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Fax) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Fax) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Fax) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *Fax) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *Fax) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *Fax) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *Fax) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookFailoverUrl

`func (o *Fax) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *Fax) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *Fax) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *Fax) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### GetStoreMedia

`func (o *Fax) GetStoreMedia() bool`

GetStoreMedia returns the StoreMedia field if non-nil, zero value otherwise.

### GetStoreMediaOk

`func (o *Fax) GetStoreMediaOk() (*bool, bool)`

GetStoreMediaOk returns a tuple with the StoreMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreMedia

`func (o *Fax) SetStoreMedia(v bool)`

SetStoreMedia sets StoreMedia field to given value.

### HasStoreMedia

`func (o *Fax) HasStoreMedia() bool`

HasStoreMedia returns a boolean if a field has been set.

### GetStoredMediaUrl

`func (o *Fax) GetStoredMediaUrl() string`

GetStoredMediaUrl returns the StoredMediaUrl field if non-nil, zero value otherwise.

### GetStoredMediaUrlOk

`func (o *Fax) GetStoredMediaUrlOk() (*string, bool)`

GetStoredMediaUrlOk returns a tuple with the StoredMediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredMediaUrl

`func (o *Fax) SetStoredMediaUrl(v string)`

SetStoredMediaUrl sets StoredMediaUrl field to given value.

### HasStoredMediaUrl

`func (o *Fax) HasStoredMediaUrl() bool`

HasStoredMediaUrl returns a boolean if a field has been set.

### GetPreviewUrl

`func (o *Fax) GetPreviewUrl() string`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *Fax) GetPreviewUrlOk() (*string, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *Fax) SetPreviewUrl(v string)`

SetPreviewUrl sets PreviewUrl field to given value.

### HasPreviewUrl

`func (o *Fax) HasPreviewUrl() bool`

HasPreviewUrl returns a boolean if a field has been set.

### GetClientState

`func (o *Fax) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *Fax) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *Fax) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *Fax) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Fax) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Fax) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Fax) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Fax) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Fax) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Fax) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Fax) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Fax) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



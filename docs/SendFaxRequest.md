# SendFaxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | The connection ID to send the fax with. | 
**MediaUrl** | Pointer to **string** | The URL (or list of URLs) to the PDF used for the fax&#39;s media. media_url and media_name/contents can&#39;t be submitted together. | [optional] 
**MediaName** | Pointer to **string** | The media_name used for the fax&#39;s media. Must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. media_name and media_url/contents can&#39;t be submitted together. | [optional] 
**To** | **string** | The phone number, in E.164 format, the fax will be sent to or SIP URI | 
**From** | **string** | The phone number, in E.164 format, the fax will be sent from. | 
**FromDisplayName** | Pointer to **string** | The &#x60;from_display_name&#x60; string to be used as the caller id name (SIP From Display Name) presented to the destination (&#x60;to&#x60; number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and -_~!.+ special characters. If ommited, the display name will be the same as the number in the &#x60;from&#x60; field. | [optional] 
**Quality** | Pointer to [**Quality**](Quality.md) |  | [optional] [default to HIGH]
**T38Enabled** | Pointer to **bool** | The flag to disable the T.38 protocol. | [optional] [default to true]
**Monochrome** | Pointer to **bool** | The flag to enable monochrome, true black and white fax results. | [optional] [default to false]
**StoreMedia** | Pointer to **bool** | Should fax media be stored on temporary URL. It does not support media_name, they can&#39;t be submitted together. | [optional] [default to false]
**StorePreview** | Pointer to **bool** | Should fax preview be stored on temporary URL. | [optional] [default to false]
**PreviewFormat** | Pointer to **string** | The format for the preview file in case the &#x60;store_preview&#x60; is &#x60;true&#x60;. | [optional] [default to "tiff"]
**WebhookUrl** | Pointer to **string** | Use this field to override the URL to which Telnyx will send subsequent webhooks for this fax. | [optional] 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 

## Methods

### NewSendFaxRequest

`func NewSendFaxRequest(connectionId string, to string, from string, ) *SendFaxRequest`

NewSendFaxRequest instantiates a new SendFaxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendFaxRequestWithDefaults

`func NewSendFaxRequestWithDefaults() *SendFaxRequest`

NewSendFaxRequestWithDefaults instantiates a new SendFaxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *SendFaxRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SendFaxRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SendFaxRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetMediaUrl

`func (o *SendFaxRequest) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *SendFaxRequest) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *SendFaxRequest) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *SendFaxRequest) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### GetMediaName

`func (o *SendFaxRequest) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *SendFaxRequest) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *SendFaxRequest) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *SendFaxRequest) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.

### GetTo

`func (o *SendFaxRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SendFaxRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SendFaxRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *SendFaxRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SendFaxRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SendFaxRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetFromDisplayName

`func (o *SendFaxRequest) GetFromDisplayName() string`

GetFromDisplayName returns the FromDisplayName field if non-nil, zero value otherwise.

### GetFromDisplayNameOk

`func (o *SendFaxRequest) GetFromDisplayNameOk() (*string, bool)`

GetFromDisplayNameOk returns a tuple with the FromDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDisplayName

`func (o *SendFaxRequest) SetFromDisplayName(v string)`

SetFromDisplayName sets FromDisplayName field to given value.

### HasFromDisplayName

`func (o *SendFaxRequest) HasFromDisplayName() bool`

HasFromDisplayName returns a boolean if a field has been set.

### GetQuality

`func (o *SendFaxRequest) GetQuality() Quality`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *SendFaxRequest) GetQualityOk() (*Quality, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *SendFaxRequest) SetQuality(v Quality)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *SendFaxRequest) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetT38Enabled

`func (o *SendFaxRequest) GetT38Enabled() bool`

GetT38Enabled returns the T38Enabled field if non-nil, zero value otherwise.

### GetT38EnabledOk

`func (o *SendFaxRequest) GetT38EnabledOk() (*bool, bool)`

GetT38EnabledOk returns a tuple with the T38Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT38Enabled

`func (o *SendFaxRequest) SetT38Enabled(v bool)`

SetT38Enabled sets T38Enabled field to given value.

### HasT38Enabled

`func (o *SendFaxRequest) HasT38Enabled() bool`

HasT38Enabled returns a boolean if a field has been set.

### GetMonochrome

`func (o *SendFaxRequest) GetMonochrome() bool`

GetMonochrome returns the Monochrome field if non-nil, zero value otherwise.

### GetMonochromeOk

`func (o *SendFaxRequest) GetMonochromeOk() (*bool, bool)`

GetMonochromeOk returns a tuple with the Monochrome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonochrome

`func (o *SendFaxRequest) SetMonochrome(v bool)`

SetMonochrome sets Monochrome field to given value.

### HasMonochrome

`func (o *SendFaxRequest) HasMonochrome() bool`

HasMonochrome returns a boolean if a field has been set.

### GetStoreMedia

`func (o *SendFaxRequest) GetStoreMedia() bool`

GetStoreMedia returns the StoreMedia field if non-nil, zero value otherwise.

### GetStoreMediaOk

`func (o *SendFaxRequest) GetStoreMediaOk() (*bool, bool)`

GetStoreMediaOk returns a tuple with the StoreMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreMedia

`func (o *SendFaxRequest) SetStoreMedia(v bool)`

SetStoreMedia sets StoreMedia field to given value.

### HasStoreMedia

`func (o *SendFaxRequest) HasStoreMedia() bool`

HasStoreMedia returns a boolean if a field has been set.

### GetStorePreview

`func (o *SendFaxRequest) GetStorePreview() bool`

GetStorePreview returns the StorePreview field if non-nil, zero value otherwise.

### GetStorePreviewOk

`func (o *SendFaxRequest) GetStorePreviewOk() (*bool, bool)`

GetStorePreviewOk returns a tuple with the StorePreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePreview

`func (o *SendFaxRequest) SetStorePreview(v bool)`

SetStorePreview sets StorePreview field to given value.

### HasStorePreview

`func (o *SendFaxRequest) HasStorePreview() bool`

HasStorePreview returns a boolean if a field has been set.

### GetPreviewFormat

`func (o *SendFaxRequest) GetPreviewFormat() string`

GetPreviewFormat returns the PreviewFormat field if non-nil, zero value otherwise.

### GetPreviewFormatOk

`func (o *SendFaxRequest) GetPreviewFormatOk() (*string, bool)`

GetPreviewFormatOk returns a tuple with the PreviewFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewFormat

`func (o *SendFaxRequest) SetPreviewFormat(v string)`

SetPreviewFormat sets PreviewFormat field to given value.

### HasPreviewFormat

`func (o *SendFaxRequest) HasPreviewFormat() bool`

HasPreviewFormat returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *SendFaxRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *SendFaxRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *SendFaxRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *SendFaxRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetClientState

`func (o *SendFaxRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *SendFaxRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *SendFaxRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *SendFaxRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



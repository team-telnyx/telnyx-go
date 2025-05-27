# UpdateVerifyProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**WebhookFailoverUrl** | Pointer to **string** |  | [optional] 
**Sms** | Pointer to [**UpdateVerifyProfileSMSRequest**](UpdateVerifyProfileSMSRequest.md) |  | [optional] 
**Call** | Pointer to [**UpdateVerifyProfileCallRequest**](UpdateVerifyProfileCallRequest.md) |  | [optional] 
**Flashcall** | Pointer to [**UpdateVerifyProfileFlashcallRequest**](UpdateVerifyProfileFlashcallRequest.md) |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateVerifyProfileRequest

`func NewUpdateVerifyProfileRequest() *UpdateVerifyProfileRequest`

NewUpdateVerifyProfileRequest instantiates a new UpdateVerifyProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVerifyProfileRequestWithDefaults

`func NewUpdateVerifyProfileRequestWithDefaults() *UpdateVerifyProfileRequest`

NewUpdateVerifyProfileRequestWithDefaults instantiates a new UpdateVerifyProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateVerifyProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVerifyProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVerifyProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVerifyProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *UpdateVerifyProfileRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UpdateVerifyProfileRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UpdateVerifyProfileRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *UpdateVerifyProfileRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookFailoverUrl

`func (o *UpdateVerifyProfileRequest) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *UpdateVerifyProfileRequest) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *UpdateVerifyProfileRequest) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *UpdateVerifyProfileRequest) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### GetSms

`func (o *UpdateVerifyProfileRequest) GetSms() UpdateVerifyProfileSMSRequest`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *UpdateVerifyProfileRequest) GetSmsOk() (*UpdateVerifyProfileSMSRequest, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *UpdateVerifyProfileRequest) SetSms(v UpdateVerifyProfileSMSRequest)`

SetSms sets Sms field to given value.

### HasSms

`func (o *UpdateVerifyProfileRequest) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetCall

`func (o *UpdateVerifyProfileRequest) GetCall() UpdateVerifyProfileCallRequest`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *UpdateVerifyProfileRequest) GetCallOk() (*UpdateVerifyProfileCallRequest, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *UpdateVerifyProfileRequest) SetCall(v UpdateVerifyProfileCallRequest)`

SetCall sets Call field to given value.

### HasCall

`func (o *UpdateVerifyProfileRequest) HasCall() bool`

HasCall returns a boolean if a field has been set.

### GetFlashcall

`func (o *UpdateVerifyProfileRequest) GetFlashcall() UpdateVerifyProfileFlashcallRequest`

GetFlashcall returns the Flashcall field if non-nil, zero value otherwise.

### GetFlashcallOk

`func (o *UpdateVerifyProfileRequest) GetFlashcallOk() (*UpdateVerifyProfileFlashcallRequest, bool)`

GetFlashcallOk returns a tuple with the Flashcall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashcall

`func (o *UpdateVerifyProfileRequest) SetFlashcall(v UpdateVerifyProfileFlashcallRequest)`

SetFlashcall sets Flashcall field to given value.

### HasFlashcall

`func (o *UpdateVerifyProfileRequest) HasFlashcall() bool`

HasFlashcall returns a boolean if a field has been set.

### GetLanguage

`func (o *UpdateVerifyProfileRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UpdateVerifyProfileRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UpdateVerifyProfileRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UpdateVerifyProfileRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



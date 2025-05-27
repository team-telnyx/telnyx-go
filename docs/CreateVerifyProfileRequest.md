# CreateVerifyProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**WebhookFailoverUrl** | Pointer to **string** |  | [optional] 
**Sms** | Pointer to [**CreateVerifyProfileSMSRequest**](CreateVerifyProfileSMSRequest.md) |  | [optional] 
**Call** | Pointer to [**CreateVerifyProfileCallRequest**](CreateVerifyProfileCallRequest.md) |  | [optional] 
**Flashcall** | Pointer to [**CreateVerifyProfileFlashcallRequest**](CreateVerifyProfileFlashcallRequest.md) |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateVerifyProfileRequest

`func NewCreateVerifyProfileRequest(name string, ) *CreateVerifyProfileRequest`

NewCreateVerifyProfileRequest instantiates a new CreateVerifyProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVerifyProfileRequestWithDefaults

`func NewCreateVerifyProfileRequestWithDefaults() *CreateVerifyProfileRequest`

NewCreateVerifyProfileRequestWithDefaults instantiates a new CreateVerifyProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVerifyProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVerifyProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVerifyProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetWebhookUrl

`func (o *CreateVerifyProfileRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateVerifyProfileRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateVerifyProfileRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateVerifyProfileRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookFailoverUrl

`func (o *CreateVerifyProfileRequest) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *CreateVerifyProfileRequest) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *CreateVerifyProfileRequest) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *CreateVerifyProfileRequest) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### GetSms

`func (o *CreateVerifyProfileRequest) GetSms() CreateVerifyProfileSMSRequest`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *CreateVerifyProfileRequest) GetSmsOk() (*CreateVerifyProfileSMSRequest, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *CreateVerifyProfileRequest) SetSms(v CreateVerifyProfileSMSRequest)`

SetSms sets Sms field to given value.

### HasSms

`func (o *CreateVerifyProfileRequest) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetCall

`func (o *CreateVerifyProfileRequest) GetCall() CreateVerifyProfileCallRequest`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *CreateVerifyProfileRequest) GetCallOk() (*CreateVerifyProfileCallRequest, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *CreateVerifyProfileRequest) SetCall(v CreateVerifyProfileCallRequest)`

SetCall sets Call field to given value.

### HasCall

`func (o *CreateVerifyProfileRequest) HasCall() bool`

HasCall returns a boolean if a field has been set.

### GetFlashcall

`func (o *CreateVerifyProfileRequest) GetFlashcall() CreateVerifyProfileFlashcallRequest`

GetFlashcall returns the Flashcall field if non-nil, zero value otherwise.

### GetFlashcallOk

`func (o *CreateVerifyProfileRequest) GetFlashcallOk() (*CreateVerifyProfileFlashcallRequest, bool)`

GetFlashcallOk returns a tuple with the Flashcall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashcall

`func (o *CreateVerifyProfileRequest) SetFlashcall(v CreateVerifyProfileFlashcallRequest)`

SetFlashcall sets Flashcall field to given value.

### HasFlashcall

`func (o *CreateVerifyProfileRequest) HasFlashcall() bool`

HasFlashcall returns a boolean if a field has been set.

### GetLanguage

`func (o *CreateVerifyProfileRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CreateVerifyProfileRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CreateVerifyProfileRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CreateVerifyProfileRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



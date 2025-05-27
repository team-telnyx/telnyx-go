# VerifyProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**WebhookFailoverUrl** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to [**VerificationProfileRecordType**](VerificationProfileRecordType.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Sms** | Pointer to [**VerifyProfileSMSResponse**](VerifyProfileSMSResponse.md) |  | [optional] 
**Call** | Pointer to [**VerifyProfileCallResponse**](VerifyProfileCallResponse.md) |  | [optional] 
**Flashcall** | Pointer to [**VerifyProfileFlashcallResponse**](VerifyProfileFlashcallResponse.md) |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 

## Methods

### NewVerifyProfileResponse

`func NewVerifyProfileResponse() *VerifyProfileResponse`

NewVerifyProfileResponse instantiates a new VerifyProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyProfileResponseWithDefaults

`func NewVerifyProfileResponseWithDefaults() *VerifyProfileResponse`

NewVerifyProfileResponseWithDefaults instantiates a new VerifyProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VerifyProfileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerifyProfileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerifyProfileResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VerifyProfileResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VerifyProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VerifyProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VerifyProfileResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VerifyProfileResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *VerifyProfileResponse) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *VerifyProfileResponse) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *VerifyProfileResponse) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *VerifyProfileResponse) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookFailoverUrl

`func (o *VerifyProfileResponse) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *VerifyProfileResponse) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *VerifyProfileResponse) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *VerifyProfileResponse) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### GetRecordType

`func (o *VerifyProfileResponse) GetRecordType() VerificationProfileRecordType`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *VerifyProfileResponse) GetRecordTypeOk() (*VerificationProfileRecordType, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *VerifyProfileResponse) SetRecordType(v VerificationProfileRecordType)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *VerifyProfileResponse) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VerifyProfileResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VerifyProfileResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VerifyProfileResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VerifyProfileResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VerifyProfileResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VerifyProfileResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VerifyProfileResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VerifyProfileResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSms

`func (o *VerifyProfileResponse) GetSms() VerifyProfileSMSResponse`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *VerifyProfileResponse) GetSmsOk() (*VerifyProfileSMSResponse, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *VerifyProfileResponse) SetSms(v VerifyProfileSMSResponse)`

SetSms sets Sms field to given value.

### HasSms

`func (o *VerifyProfileResponse) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetCall

`func (o *VerifyProfileResponse) GetCall() VerifyProfileCallResponse`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *VerifyProfileResponse) GetCallOk() (*VerifyProfileCallResponse, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *VerifyProfileResponse) SetCall(v VerifyProfileCallResponse)`

SetCall sets Call field to given value.

### HasCall

`func (o *VerifyProfileResponse) HasCall() bool`

HasCall returns a boolean if a field has been set.

### GetFlashcall

`func (o *VerifyProfileResponse) GetFlashcall() VerifyProfileFlashcallResponse`

GetFlashcall returns the Flashcall field if non-nil, zero value otherwise.

### GetFlashcallOk

`func (o *VerifyProfileResponse) GetFlashcallOk() (*VerifyProfileFlashcallResponse, bool)`

GetFlashcallOk returns a tuple with the Flashcall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashcall

`func (o *VerifyProfileResponse) SetFlashcall(v VerifyProfileFlashcallResponse)`

SetFlashcall sets Flashcall field to given value.

### HasFlashcall

`func (o *VerifyProfileResponse) HasFlashcall() bool`

HasFlashcall returns a boolean if a field has been set.

### GetLanguage

`func (o *VerifyProfileResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *VerifyProfileResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *VerifyProfileResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *VerifyProfileResponse) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



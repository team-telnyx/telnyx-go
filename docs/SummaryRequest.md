# SummaryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | The name of the bucket that contains the file to be summarized. | 
**Filename** | **string** | The name of the file to be summarized. | 
**SystemPrompt** | Pointer to **string** | A system prompt to guide the summary generation. | [optional] 

## Methods

### NewSummaryRequest

`func NewSummaryRequest(bucket string, filename string, ) *SummaryRequest`

NewSummaryRequest instantiates a new SummaryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryRequestWithDefaults

`func NewSummaryRequestWithDefaults() *SummaryRequest`

NewSummaryRequestWithDefaults instantiates a new SummaryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *SummaryRequest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *SummaryRequest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *SummaryRequest) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetFilename

`func (o *SummaryRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SummaryRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SummaryRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetSystemPrompt

`func (o *SummaryRequest) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *SummaryRequest) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *SummaryRequest) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *SummaryRequest) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



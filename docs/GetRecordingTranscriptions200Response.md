# GetRecordingTranscriptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RecordingTranscription**](RecordingTranscription.md) |  | [optional] 
**Meta** | Pointer to [**CursorPaginationMeta**](CursorPaginationMeta.md) |  | [optional] 

## Methods

### NewGetRecordingTranscriptions200Response

`func NewGetRecordingTranscriptions200Response() *GetRecordingTranscriptions200Response`

NewGetRecordingTranscriptions200Response instantiates a new GetRecordingTranscriptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordingTranscriptions200ResponseWithDefaults

`func NewGetRecordingTranscriptions200ResponseWithDefaults() *GetRecordingTranscriptions200Response`

NewGetRecordingTranscriptions200ResponseWithDefaults instantiates a new GetRecordingTranscriptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRecordingTranscriptions200Response) GetData() []RecordingTranscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRecordingTranscriptions200Response) GetDataOk() (*[]RecordingTranscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRecordingTranscriptions200Response) SetData(v []RecordingTranscription)`

SetData sets Data field to given value.

### HasData

`func (o *GetRecordingTranscriptions200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetRecordingTranscriptions200Response) GetMeta() CursorPaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetRecordingTranscriptions200Response) GetMetaOk() (*CursorPaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetRecordingTranscriptions200Response) SetMeta(v CursorPaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetRecordingTranscriptions200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



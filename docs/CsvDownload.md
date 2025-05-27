# CsvDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Url** | Pointer to **string** | The URL at which the CSV file can be retrieved. | [optional] 
**Status** | Pointer to **string** | Indicates the completion level of the CSV report. Only complete CSV download requests will be able to be retrieved. | [optional] [default to "pending"]

## Methods

### NewCsvDownload

`func NewCsvDownload() *CsvDownload`

NewCsvDownload instantiates a new CsvDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsvDownloadWithDefaults

`func NewCsvDownloadWithDefaults() *CsvDownload`

NewCsvDownloadWithDefaults instantiates a new CsvDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CsvDownload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CsvDownload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CsvDownload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CsvDownload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *CsvDownload) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CsvDownload) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CsvDownload) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CsvDownload) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetUrl

`func (o *CsvDownload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CsvDownload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CsvDownload) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CsvDownload) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *CsvDownload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CsvDownload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CsvDownload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CsvDownload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



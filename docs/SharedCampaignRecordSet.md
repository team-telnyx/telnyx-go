# SharedCampaignRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** |  | [optional] 
**Records** | Pointer to [**[]SharedCampaign**](SharedCampaign.md) |  | [optional] 
**TotalRecords** | Pointer to **int32** |  | [optional] 

## Methods

### NewSharedCampaignRecordSet

`func NewSharedCampaignRecordSet() *SharedCampaignRecordSet`

NewSharedCampaignRecordSet instantiates a new SharedCampaignRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedCampaignRecordSetWithDefaults

`func NewSharedCampaignRecordSetWithDefaults() *SharedCampaignRecordSet`

NewSharedCampaignRecordSetWithDefaults instantiates a new SharedCampaignRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *SharedCampaignRecordSet) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SharedCampaignRecordSet) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SharedCampaignRecordSet) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SharedCampaignRecordSet) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRecords

`func (o *SharedCampaignRecordSet) GetRecords() []SharedCampaign`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *SharedCampaignRecordSet) GetRecordsOk() (*[]SharedCampaign, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *SharedCampaignRecordSet) SetRecords(v []SharedCampaign)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *SharedCampaignRecordSet) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotalRecords

`func (o *SharedCampaignRecordSet) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *SharedCampaignRecordSet) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *SharedCampaignRecordSet) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *SharedCampaignRecordSet) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



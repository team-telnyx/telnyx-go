# PWGAssignedResourcesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | The type of the resource assigned to the Private Wireless Gateway. | [optional] [readonly] 
**Count** | Pointer to **int32** | The current count of a resource type assigned to the Private Wireless Gateway. | [optional] [readonly] 

## Methods

### NewPWGAssignedResourcesSummary

`func NewPWGAssignedResourcesSummary() *PWGAssignedResourcesSummary`

NewPWGAssignedResourcesSummary instantiates a new PWGAssignedResourcesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPWGAssignedResourcesSummaryWithDefaults

`func NewPWGAssignedResourcesSummaryWithDefaults() *PWGAssignedResourcesSummary`

NewPWGAssignedResourcesSummaryWithDefaults instantiates a new PWGAssignedResourcesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *PWGAssignedResourcesSummary) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PWGAssignedResourcesSummary) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PWGAssignedResourcesSummary) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PWGAssignedResourcesSummary) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCount

`func (o *PWGAssignedResourcesSummary) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PWGAssignedResourcesSummary) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PWGAssignedResourcesSummary) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PWGAssignedResourcesSummary) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



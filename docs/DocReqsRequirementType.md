# DocReqsRequirementType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptanceCriteria** | Pointer to [**DocReqsRequirementTypeAcceptanceCriteria**](DocReqsRequirementTypeAcceptanceCriteria.md) |  | [optional] 
**Description** | Pointer to **string** | Describes the requirement type | [optional] 
**Example** | Pointer to **string** | Provides one or more example of acceptable documents | [optional] 
**Type** | Pointer to **string** | Defines the type of this requirement type | [optional] 
**Name** | Pointer to **string** | A short descriptive name for this requirement_type | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource | [optional] [readonly] 
**Id** | Pointer to **string** | Identifies the associated document | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was last updated. | [optional] [readonly] 

## Methods

### NewDocReqsRequirementType

`func NewDocReqsRequirementType() *DocReqsRequirementType`

NewDocReqsRequirementType instantiates a new DocReqsRequirementType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocReqsRequirementTypeWithDefaults

`func NewDocReqsRequirementTypeWithDefaults() *DocReqsRequirementType`

NewDocReqsRequirementTypeWithDefaults instantiates a new DocReqsRequirementType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptanceCriteria

`func (o *DocReqsRequirementType) GetAcceptanceCriteria() DocReqsRequirementTypeAcceptanceCriteria`

GetAcceptanceCriteria returns the AcceptanceCriteria field if non-nil, zero value otherwise.

### GetAcceptanceCriteriaOk

`func (o *DocReqsRequirementType) GetAcceptanceCriteriaOk() (*DocReqsRequirementTypeAcceptanceCriteria, bool)`

GetAcceptanceCriteriaOk returns a tuple with the AcceptanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceCriteria

`func (o *DocReqsRequirementType) SetAcceptanceCriteria(v DocReqsRequirementTypeAcceptanceCriteria)`

SetAcceptanceCriteria sets AcceptanceCriteria field to given value.

### HasAcceptanceCriteria

`func (o *DocReqsRequirementType) HasAcceptanceCriteria() bool`

HasAcceptanceCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *DocReqsRequirementType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DocReqsRequirementType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DocReqsRequirementType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DocReqsRequirementType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *DocReqsRequirementType) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *DocReqsRequirementType) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *DocReqsRequirementType) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *DocReqsRequirementType) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetType

`func (o *DocReqsRequirementType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocReqsRequirementType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocReqsRequirementType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DocReqsRequirementType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DocReqsRequirementType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocReqsRequirementType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocReqsRequirementType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocReqsRequirementType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecordType

`func (o *DocReqsRequirementType) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *DocReqsRequirementType) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *DocReqsRequirementType) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *DocReqsRequirementType) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *DocReqsRequirementType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocReqsRequirementType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocReqsRequirementType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocReqsRequirementType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocReqsRequirementType) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocReqsRequirementType) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocReqsRequirementType) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocReqsRequirementType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DocReqsRequirementType) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DocReqsRequirementType) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DocReqsRequirementType) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DocReqsRequirementType) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



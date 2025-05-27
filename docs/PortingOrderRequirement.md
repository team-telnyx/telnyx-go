# PortingOrderRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldType** | Pointer to **string** | Type of value expected on field_value field | [optional] 
**FieldValue** | Pointer to **string** | identifies the document that satisfies this requirement | [optional] 
**RequirementTypeId** | Pointer to **string** | Identifies the requirement type that meets this requirement | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 

## Methods

### NewPortingOrderRequirement

`func NewPortingOrderRequirement() *PortingOrderRequirement`

NewPortingOrderRequirement instantiates a new PortingOrderRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingOrderRequirementWithDefaults

`func NewPortingOrderRequirementWithDefaults() *PortingOrderRequirement`

NewPortingOrderRequirementWithDefaults instantiates a new PortingOrderRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldType

`func (o *PortingOrderRequirement) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *PortingOrderRequirement) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *PortingOrderRequirement) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *PortingOrderRequirement) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetFieldValue

`func (o *PortingOrderRequirement) GetFieldValue() string`

GetFieldValue returns the FieldValue field if non-nil, zero value otherwise.

### GetFieldValueOk

`func (o *PortingOrderRequirement) GetFieldValueOk() (*string, bool)`

GetFieldValueOk returns a tuple with the FieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValue

`func (o *PortingOrderRequirement) SetFieldValue(v string)`

SetFieldValue sets FieldValue field to given value.

### HasFieldValue

`func (o *PortingOrderRequirement) HasFieldValue() bool`

HasFieldValue returns a boolean if a field has been set.

### GetRequirementTypeId

`func (o *PortingOrderRequirement) GetRequirementTypeId() string`

GetRequirementTypeId returns the RequirementTypeId field if non-nil, zero value otherwise.

### GetRequirementTypeIdOk

`func (o *PortingOrderRequirement) GetRequirementTypeIdOk() (*string, bool)`

GetRequirementTypeIdOk returns a tuple with the RequirementTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementTypeId

`func (o *PortingOrderRequirement) SetRequirementTypeId(v string)`

SetRequirementTypeId sets RequirementTypeId field to given value.

### HasRequirementTypeId

`func (o *PortingOrderRequirement) HasRequirementTypeId() bool`

HasRequirementTypeId returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingOrderRequirement) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingOrderRequirement) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingOrderRequirement) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingOrderRequirement) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



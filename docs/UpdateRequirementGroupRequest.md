# UpdateRequirementGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerReference** | Pointer to **string** | Reference for the customer | [optional] 
**RegulatoryRequirements** | Pointer to [**[]UpdateRequirementGroupRequestRegulatoryRequirementsInner**](UpdateRequirementGroupRequestRegulatoryRequirementsInner.md) |  | [optional] 

## Methods

### NewUpdateRequirementGroupRequest

`func NewUpdateRequirementGroupRequest() *UpdateRequirementGroupRequest`

NewUpdateRequirementGroupRequest instantiates a new UpdateRequirementGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRequirementGroupRequestWithDefaults

`func NewUpdateRequirementGroupRequestWithDefaults() *UpdateRequirementGroupRequest`

NewUpdateRequirementGroupRequestWithDefaults instantiates a new UpdateRequirementGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerReference

`func (o *UpdateRequirementGroupRequest) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *UpdateRequirementGroupRequest) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *UpdateRequirementGroupRequest) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *UpdateRequirementGroupRequest) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetRegulatoryRequirements

`func (o *UpdateRequirementGroupRequest) GetRegulatoryRequirements() []UpdateRequirementGroupRequestRegulatoryRequirementsInner`

GetRegulatoryRequirements returns the RegulatoryRequirements field if non-nil, zero value otherwise.

### GetRegulatoryRequirementsOk

`func (o *UpdateRequirementGroupRequest) GetRegulatoryRequirementsOk() (*[]UpdateRequirementGroupRequestRegulatoryRequirementsInner, bool)`

GetRegulatoryRequirementsOk returns a tuple with the RegulatoryRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryRequirements

`func (o *UpdateRequirementGroupRequest) SetRegulatoryRequirements(v []UpdateRequirementGroupRequestRegulatoryRequirementsInner)`

SetRegulatoryRequirements sets RegulatoryRequirements field to given value.

### HasRegulatoryRequirements

`func (o *UpdateRequirementGroupRequest) HasRegulatoryRequirements() bool`

HasRegulatoryRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



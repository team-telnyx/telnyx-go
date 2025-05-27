# UpdateNumberOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegulatoryRequirements** | Pointer to [**[]UpdateRegulatoryRequirement**](UpdateRegulatoryRequirement.md) |  | [optional] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 

## Methods

### NewUpdateNumberOrderRequest

`func NewUpdateNumberOrderRequest() *UpdateNumberOrderRequest`

NewUpdateNumberOrderRequest instantiates a new UpdateNumberOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNumberOrderRequestWithDefaults

`func NewUpdateNumberOrderRequestWithDefaults() *UpdateNumberOrderRequest`

NewUpdateNumberOrderRequestWithDefaults instantiates a new UpdateNumberOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegulatoryRequirements

`func (o *UpdateNumberOrderRequest) GetRegulatoryRequirements() []UpdateRegulatoryRequirement`

GetRegulatoryRequirements returns the RegulatoryRequirements field if non-nil, zero value otherwise.

### GetRegulatoryRequirementsOk

`func (o *UpdateNumberOrderRequest) GetRegulatoryRequirementsOk() (*[]UpdateRegulatoryRequirement, bool)`

GetRegulatoryRequirementsOk returns a tuple with the RegulatoryRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryRequirements

`func (o *UpdateNumberOrderRequest) SetRegulatoryRequirements(v []UpdateRegulatoryRequirement)`

SetRegulatoryRequirements sets RegulatoryRequirements field to given value.

### HasRegulatoryRequirements

`func (o *UpdateNumberOrderRequest) HasRegulatoryRequirements() bool`

HasRegulatoryRequirements returns a boolean if a field has been set.

### GetCustomerReference

`func (o *UpdateNumberOrderRequest) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *UpdateNumberOrderRequest) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *UpdateNumberOrderRequest) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *UpdateNumberOrderRequest) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



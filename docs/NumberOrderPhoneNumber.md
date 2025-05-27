# NumberOrderPhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**OrderRequestId** | Pointer to **string** |  | [optional] 
**SubNumberOrderId** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**PhoneNumberType** | Pointer to **string** |  | [optional] 
**RegulatoryRequirements** | Pointer to [**[]SubNumberOrderRegulatoryRequirementWithValue**](SubNumberOrderRegulatoryRequirementWithValue.md) |  | [optional] 
**RequirementsMet** | Pointer to **bool** | True if all requirements are met for a phone number, false otherwise. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the phone number in the order. | [optional] [readonly] 
**BundleId** | Pointer to **string** |  | [optional] [readonly] 
**Locality** | Pointer to **string** |  | [optional] 
**Deadline** | Pointer to **string** |  | [optional] 
**RequirementsStatus** | Pointer to **string** | Status of requirements (if applicable) | [optional] [readonly] 
**IsBlockNumber** | Pointer to **bool** |  | [optional] 

## Methods

### NewNumberOrderPhoneNumber

`func NewNumberOrderPhoneNumber() *NumberOrderPhoneNumber`

NewNumberOrderPhoneNumber instantiates a new NumberOrderPhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberOrderPhoneNumberWithDefaults

`func NewNumberOrderPhoneNumberWithDefaults() *NumberOrderPhoneNumber`

NewNumberOrderPhoneNumberWithDefaults instantiates a new NumberOrderPhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NumberOrderPhoneNumber) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NumberOrderPhoneNumber) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NumberOrderPhoneNumber) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NumberOrderPhoneNumber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *NumberOrderPhoneNumber) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NumberOrderPhoneNumber) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NumberOrderPhoneNumber) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NumberOrderPhoneNumber) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *NumberOrderPhoneNumber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *NumberOrderPhoneNumber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *NumberOrderPhoneNumber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *NumberOrderPhoneNumber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetOrderRequestId

`func (o *NumberOrderPhoneNumber) GetOrderRequestId() string`

GetOrderRequestId returns the OrderRequestId field if non-nil, zero value otherwise.

### GetOrderRequestIdOk

`func (o *NumberOrderPhoneNumber) GetOrderRequestIdOk() (*string, bool)`

GetOrderRequestIdOk returns a tuple with the OrderRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRequestId

`func (o *NumberOrderPhoneNumber) SetOrderRequestId(v string)`

SetOrderRequestId sets OrderRequestId field to given value.

### HasOrderRequestId

`func (o *NumberOrderPhoneNumber) HasOrderRequestId() bool`

HasOrderRequestId returns a boolean if a field has been set.

### GetSubNumberOrderId

`func (o *NumberOrderPhoneNumber) GetSubNumberOrderId() string`

GetSubNumberOrderId returns the SubNumberOrderId field if non-nil, zero value otherwise.

### GetSubNumberOrderIdOk

`func (o *NumberOrderPhoneNumber) GetSubNumberOrderIdOk() (*string, bool)`

GetSubNumberOrderIdOk returns a tuple with the SubNumberOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNumberOrderId

`func (o *NumberOrderPhoneNumber) SetSubNumberOrderId(v string)`

SetSubNumberOrderId sets SubNumberOrderId field to given value.

### HasSubNumberOrderId

`func (o *NumberOrderPhoneNumber) HasSubNumberOrderId() bool`

HasSubNumberOrderId returns a boolean if a field has been set.

### GetCountryCode

`func (o *NumberOrderPhoneNumber) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *NumberOrderPhoneNumber) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *NumberOrderPhoneNumber) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *NumberOrderPhoneNumber) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *NumberOrderPhoneNumber) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *NumberOrderPhoneNumber) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *NumberOrderPhoneNumber) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *NumberOrderPhoneNumber) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetRegulatoryRequirements

`func (o *NumberOrderPhoneNumber) GetRegulatoryRequirements() []SubNumberOrderRegulatoryRequirementWithValue`

GetRegulatoryRequirements returns the RegulatoryRequirements field if non-nil, zero value otherwise.

### GetRegulatoryRequirementsOk

`func (o *NumberOrderPhoneNumber) GetRegulatoryRequirementsOk() (*[]SubNumberOrderRegulatoryRequirementWithValue, bool)`

GetRegulatoryRequirementsOk returns a tuple with the RegulatoryRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryRequirements

`func (o *NumberOrderPhoneNumber) SetRegulatoryRequirements(v []SubNumberOrderRegulatoryRequirementWithValue)`

SetRegulatoryRequirements sets RegulatoryRequirements field to given value.

### HasRegulatoryRequirements

`func (o *NumberOrderPhoneNumber) HasRegulatoryRequirements() bool`

HasRegulatoryRequirements returns a boolean if a field has been set.

### GetRequirementsMet

`func (o *NumberOrderPhoneNumber) GetRequirementsMet() bool`

GetRequirementsMet returns the RequirementsMet field if non-nil, zero value otherwise.

### GetRequirementsMetOk

`func (o *NumberOrderPhoneNumber) GetRequirementsMetOk() (*bool, bool)`

GetRequirementsMetOk returns a tuple with the RequirementsMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsMet

`func (o *NumberOrderPhoneNumber) SetRequirementsMet(v bool)`

SetRequirementsMet sets RequirementsMet field to given value.

### HasRequirementsMet

`func (o *NumberOrderPhoneNumber) HasRequirementsMet() bool`

HasRequirementsMet returns a boolean if a field has been set.

### GetStatus

`func (o *NumberOrderPhoneNumber) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NumberOrderPhoneNumber) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NumberOrderPhoneNumber) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NumberOrderPhoneNumber) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBundleId

`func (o *NumberOrderPhoneNumber) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *NumberOrderPhoneNumber) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *NumberOrderPhoneNumber) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *NumberOrderPhoneNumber) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetLocality

`func (o *NumberOrderPhoneNumber) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *NumberOrderPhoneNumber) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *NumberOrderPhoneNumber) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *NumberOrderPhoneNumber) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetDeadline

`func (o *NumberOrderPhoneNumber) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *NumberOrderPhoneNumber) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *NumberOrderPhoneNumber) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *NumberOrderPhoneNumber) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetRequirementsStatus

`func (o *NumberOrderPhoneNumber) GetRequirementsStatus() string`

GetRequirementsStatus returns the RequirementsStatus field if non-nil, zero value otherwise.

### GetRequirementsStatusOk

`func (o *NumberOrderPhoneNumber) GetRequirementsStatusOk() (*string, bool)`

GetRequirementsStatusOk returns a tuple with the RequirementsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsStatus

`func (o *NumberOrderPhoneNumber) SetRequirementsStatus(v string)`

SetRequirementsStatus sets RequirementsStatus field to given value.

### HasRequirementsStatus

`func (o *NumberOrderPhoneNumber) HasRequirementsStatus() bool`

HasRequirementsStatus returns a boolean if a field has been set.

### GetIsBlockNumber

`func (o *NumberOrderPhoneNumber) GetIsBlockNumber() bool`

GetIsBlockNumber returns the IsBlockNumber field if non-nil, zero value otherwise.

### GetIsBlockNumberOk

`func (o *NumberOrderPhoneNumber) GetIsBlockNumberOk() (*bool, bool)`

GetIsBlockNumberOk returns a tuple with the IsBlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlockNumber

`func (o *NumberOrderPhoneNumber) SetIsBlockNumber(v bool)`

SetIsBlockNumber sets IsBlockNumber field to given value.

### HasIsBlockNumber

`func (o *NumberOrderPhoneNumber) HasIsBlockNumber() bool`

HasIsBlockNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



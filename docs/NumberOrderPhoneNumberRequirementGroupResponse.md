# NumberOrderPhoneNumberRequirementGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**OrderRequestId** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**IsBlockNumber** | Pointer to **bool** |  | [optional] 
**RegulatoryRequirements** | Pointer to [**[]NumberOrderPhoneNumberRequirementGroupResponseRegulatoryRequirementsInner**](NumberOrderPhoneNumberRequirementGroupResponseRegulatoryRequirementsInner.md) |  | [optional] 
**Locality** | Pointer to **string** |  | [optional] 
**PhoneNumberType** | Pointer to **string** |  | [optional] 
**BundleId** | Pointer to **NullableString** |  | [optional] 
**SubNumberOrderId** | Pointer to **string** |  | [optional] 
**Deadline** | Pointer to **time.Time** |  | [optional] 
**RequirementsStatus** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**RequirementsMet** | Pointer to **bool** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 

## Methods

### NewNumberOrderPhoneNumberRequirementGroupResponse

`func NewNumberOrderPhoneNumberRequirementGroupResponse() *NumberOrderPhoneNumberRequirementGroupResponse`

NewNumberOrderPhoneNumberRequirementGroupResponse instantiates a new NumberOrderPhoneNumberRequirementGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberOrderPhoneNumberRequirementGroupResponseWithDefaults

`func NewNumberOrderPhoneNumberRequirementGroupResponseWithDefaults() *NumberOrderPhoneNumberRequirementGroupResponse`

NewNumberOrderPhoneNumberRequirementGroupResponseWithDefaults instantiates a new NumberOrderPhoneNumberRequirementGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrderRequestId

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetOrderRequestId() string`

GetOrderRequestId returns the OrderRequestId field if non-nil, zero value otherwise.

### GetOrderRequestIdOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetOrderRequestIdOk() (*string, bool)`

GetOrderRequestIdOk returns a tuple with the OrderRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRequestId

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetOrderRequestId(v string)`

SetOrderRequestId sets OrderRequestId field to given value.

### HasOrderRequestId

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasOrderRequestId() bool`

HasOrderRequestId returns a boolean if a field has been set.

### GetCountryCode

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetIsBlockNumber

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetIsBlockNumber() bool`

GetIsBlockNumber returns the IsBlockNumber field if non-nil, zero value otherwise.

### GetIsBlockNumberOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetIsBlockNumberOk() (*bool, bool)`

GetIsBlockNumberOk returns a tuple with the IsBlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlockNumber

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetIsBlockNumber(v bool)`

SetIsBlockNumber sets IsBlockNumber field to given value.

### HasIsBlockNumber

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasIsBlockNumber() bool`

HasIsBlockNumber returns a boolean if a field has been set.

### GetRegulatoryRequirements

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetRegulatoryRequirements() []NumberOrderPhoneNumberRequirementGroupResponseRegulatoryRequirementsInner`

GetRegulatoryRequirements returns the RegulatoryRequirements field if non-nil, zero value otherwise.

### GetRegulatoryRequirementsOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetRegulatoryRequirementsOk() (*[]NumberOrderPhoneNumberRequirementGroupResponseRegulatoryRequirementsInner, bool)`

GetRegulatoryRequirementsOk returns a tuple with the RegulatoryRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryRequirements

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetRegulatoryRequirements(v []NumberOrderPhoneNumberRequirementGroupResponseRegulatoryRequirementsInner)`

SetRegulatoryRequirements sets RegulatoryRequirements field to given value.

### HasRegulatoryRequirements

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasRegulatoryRequirements() bool`

HasRegulatoryRequirements returns a boolean if a field has been set.

### GetLocality

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetBundleId

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### SetBundleIdNil

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetBundleIdNil(b bool)`

 SetBundleIdNil sets the value for BundleId to be an explicit nil

### UnsetBundleId
`func (o *NumberOrderPhoneNumberRequirementGroupResponse) UnsetBundleId()`

UnsetBundleId ensures that no value is present for BundleId, not even an explicit nil
### GetSubNumberOrderId

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetSubNumberOrderId() string`

GetSubNumberOrderId returns the SubNumberOrderId field if non-nil, zero value otherwise.

### GetSubNumberOrderIdOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetSubNumberOrderIdOk() (*string, bool)`

GetSubNumberOrderIdOk returns a tuple with the SubNumberOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNumberOrderId

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetSubNumberOrderId(v string)`

SetSubNumberOrderId sets SubNumberOrderId field to given value.

### HasSubNumberOrderId

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasSubNumberOrderId() bool`

HasSubNumberOrderId returns a boolean if a field has been set.

### GetDeadline

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetRequirementsStatus

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetRequirementsStatus() string`

GetRequirementsStatus returns the RequirementsStatus field if non-nil, zero value otherwise.

### GetRequirementsStatusOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetRequirementsStatusOk() (*string, bool)`

GetRequirementsStatusOk returns a tuple with the RequirementsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsStatus

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetRequirementsStatus(v string)`

SetRequirementsStatus sets RequirementsStatus field to given value.

### HasRequirementsStatus

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasRequirementsStatus() bool`

HasRequirementsStatus returns a boolean if a field has been set.

### GetId

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetRequirementsMet

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetRequirementsMet() bool`

GetRequirementsMet returns the RequirementsMet field if non-nil, zero value otherwise.

### GetRequirementsMetOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetRequirementsMetOk() (*bool, bool)`

GetRequirementsMetOk returns a tuple with the RequirementsMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsMet

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetRequirementsMet(v bool)`

SetRequirementsMet sets RequirementsMet field to given value.

### HasRequirementsMet

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasRequirementsMet() bool`

HasRequirementsMet returns a boolean if a field has been set.

### GetRecordType

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NumberOrderPhoneNumberRequirementGroupResponse) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



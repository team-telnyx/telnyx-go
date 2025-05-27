# CustomerServiceRecordResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierName** | Pointer to **string** | The name of the carrier that the customer service record is for. | [optional] 
**AssociatedPhoneNumbers** | Pointer to **[]string** | The associated phone numbers of the customer service record. | [optional] 
**Admin** | Pointer to [**CustomerServiceRecordResultAdmin**](CustomerServiceRecordResultAdmin.md) |  | [optional] 
**Address** | Pointer to [**CustomerServiceRecordResultAddress**](CustomerServiceRecordResultAddress.md) |  | [optional] 

## Methods

### NewCustomerServiceRecordResult

`func NewCustomerServiceRecordResult() *CustomerServiceRecordResult`

NewCustomerServiceRecordResult instantiates a new CustomerServiceRecordResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerServiceRecordResultWithDefaults

`func NewCustomerServiceRecordResultWithDefaults() *CustomerServiceRecordResult`

NewCustomerServiceRecordResultWithDefaults instantiates a new CustomerServiceRecordResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierName

`func (o *CustomerServiceRecordResult) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *CustomerServiceRecordResult) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *CustomerServiceRecordResult) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *CustomerServiceRecordResult) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetAssociatedPhoneNumbers

`func (o *CustomerServiceRecordResult) GetAssociatedPhoneNumbers() []string`

GetAssociatedPhoneNumbers returns the AssociatedPhoneNumbers field if non-nil, zero value otherwise.

### GetAssociatedPhoneNumbersOk

`func (o *CustomerServiceRecordResult) GetAssociatedPhoneNumbersOk() (*[]string, bool)`

GetAssociatedPhoneNumbersOk returns a tuple with the AssociatedPhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedPhoneNumbers

`func (o *CustomerServiceRecordResult) SetAssociatedPhoneNumbers(v []string)`

SetAssociatedPhoneNumbers sets AssociatedPhoneNumbers field to given value.

### HasAssociatedPhoneNumbers

`func (o *CustomerServiceRecordResult) HasAssociatedPhoneNumbers() bool`

HasAssociatedPhoneNumbers returns a boolean if a field has been set.

### GetAdmin

`func (o *CustomerServiceRecordResult) GetAdmin() CustomerServiceRecordResultAdmin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *CustomerServiceRecordResult) GetAdminOk() (*CustomerServiceRecordResultAdmin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *CustomerServiceRecordResult) SetAdmin(v CustomerServiceRecordResultAdmin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *CustomerServiceRecordResult) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetAddress

`func (o *CustomerServiceRecordResult) GetAddress() CustomerServiceRecordResultAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerServiceRecordResult) GetAddressOk() (*CustomerServiceRecordResultAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerServiceRecordResult) SetAddress(v CustomerServiceRecordResultAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CustomerServiceRecordResult) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



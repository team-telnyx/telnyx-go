/*
Telnyx API

SIP trunking, SMS, MMS, Call Control and Telephony Data Services.

API version: 2.0.0
Contact: support@telnyx.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package telnyx

import (
	"encoding/json"
)

// checks if the CustomerServiceRecordAdditionalData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerServiceRecordAdditionalData{}

// CustomerServiceRecordAdditionalData struct for CustomerServiceRecordAdditionalData
type CustomerServiceRecordAdditionalData struct {
	// The name of the administrator of CSR.
	Name *string `json:"name,omitempty"`
	// The name of the authorized person.
	AuthorizedPersonName *string `json:"authorized_person_name,omitempty"`
	// The PIN of the customer service record.
	Pin *string `json:"pin,omitempty"`
	// The account number of the customer service record.
	AccountNumber *string `json:"account_number,omitempty"`
	// The customer code of the customer service record.
	CustomerCode *string `json:"customer_code,omitempty"`
	// The first line of the address of the customer service record.
	AddressLine1 *string `json:"address_line_1,omitempty"`
	// The city of the customer service record.
	City *string `json:"city,omitempty"`
	// The state of the customer service record.
	State *string `json:"state,omitempty" validate:"regexp=^[A-Z]{2}$"`
	// The zip code of the customer service record.
	ZipCode *string `json:"zip_code,omitempty" validate:"regexp=^\\\\d{5}$"`
	// The billing phone number of the customer service record.
	BillingPhoneNumber *string `json:"billing_phone_number,omitempty" validate:"regexp=^\\\\+1\\\\d{10}$"`
}

// NewCustomerServiceRecordAdditionalData instantiates a new CustomerServiceRecordAdditionalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerServiceRecordAdditionalData() *CustomerServiceRecordAdditionalData {
	this := CustomerServiceRecordAdditionalData{}
	return &this
}

// NewCustomerServiceRecordAdditionalDataWithDefaults instantiates a new CustomerServiceRecordAdditionalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerServiceRecordAdditionalDataWithDefaults() *CustomerServiceRecordAdditionalData {
	this := CustomerServiceRecordAdditionalData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerServiceRecordAdditionalData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordAdditionalData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerServiceRecordAdditionalData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerServiceRecordAdditionalData) SetName(v string) {
	o.Name = &v
}

// GetAuthorizedPersonName returns the AuthorizedPersonName field value if set, zero value otherwise.
func (o *CustomerServiceRecordAdditionalData) GetAuthorizedPersonName() string {
	if o == nil || IsNil(o.AuthorizedPersonName) {
		var ret string
		return ret
	}
	return *o.AuthorizedPersonName
}

// GetAuthorizedPersonNameOk returns a tuple with the AuthorizedPersonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordAdditionalData) GetAuthorizedPersonNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizedPersonName) {
		return nil, false
	}
	return o.AuthorizedPersonName, true
}

// HasAuthorizedPersonName returns a boolean if a field has been set.
func (o *CustomerServiceRecordAdditionalData) HasAuthorizedPersonName() bool {
	if o != nil && !IsNil(o.AuthorizedPersonName) {
		return true
	}

	return false
}

// SetAuthorizedPersonName gets a reference to the given string and assigns it to the AuthorizedPersonName field.
func (o *CustomerServiceRecordAdditionalData) SetAuthorizedPersonName(v string) {
	o.AuthorizedPersonName = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *CustomerServiceRecordAdditionalData) GetPin() string {
	if o == nil || IsNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordAdditionalData) GetPinOk() (*string, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *CustomerServiceRecordAdditionalData) HasPin() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *CustomerServiceRecordAdditionalData) SetPin(v string) {
	o.Pin = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *CustomerServiceRecordAdditionalData) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordAdditionalData) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *CustomerServiceRecordAdditionalData) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *CustomerServiceRecordAdditionalData) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetCustomerCode returns the CustomerCode field value if set, zero value otherwise.
func (o *CustomerServiceRecordAdditionalData) GetCustomerCode() string {
	if o == nil || IsNil(o.CustomerCode) {
		var ret string
		return ret
	}
	return *o.CustomerCode
}

// GetCustomerCodeOk returns a tuple with the CustomerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordAdditionalData) GetCustomerCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerCode) {
		return nil, false
	}
	return o.CustomerCode, true
}

// HasCustomerCode returns a boolean if a field has been set.
func (o *CustomerServiceRecordAdditionalData) HasCustomerCode() bool {
	if o != nil && !IsNil(o.CustomerCode) {
		return true
	}

	return false
}

// SetCustomerCode gets a reference to the given string and assigns it to the CustomerCode field.
func (o *CustomerServiceRecordAdditionalData) SetCustomerCode(v string) {
	o.CustomerCode = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *CustomerServiceRecordAdditionalData) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordAdditionalData) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *CustomerServiceRecordAdditionalData) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *CustomerServiceRecordAdditionalData) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CustomerServiceRecordAdditionalData) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordAdditionalData) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CustomerServiceRecordAdditionalData) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CustomerServiceRecordAdditionalData) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CustomerServiceRecordAdditionalData) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordAdditionalData) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CustomerServiceRecordAdditionalData) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CustomerServiceRecordAdditionalData) SetState(v string) {
	o.State = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *CustomerServiceRecordAdditionalData) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode) {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordAdditionalData) GetZipCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *CustomerServiceRecordAdditionalData) HasZipCode() bool {
	if o != nil && !IsNil(o.ZipCode) {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *CustomerServiceRecordAdditionalData) SetZipCode(v string) {
	o.ZipCode = &v
}

// GetBillingPhoneNumber returns the BillingPhoneNumber field value if set, zero value otherwise.
func (o *CustomerServiceRecordAdditionalData) GetBillingPhoneNumber() string {
	if o == nil || IsNil(o.BillingPhoneNumber) {
		var ret string
		return ret
	}
	return *o.BillingPhoneNumber
}

// GetBillingPhoneNumberOk returns a tuple with the BillingPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordAdditionalData) GetBillingPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BillingPhoneNumber) {
		return nil, false
	}
	return o.BillingPhoneNumber, true
}

// HasBillingPhoneNumber returns a boolean if a field has been set.
func (o *CustomerServiceRecordAdditionalData) HasBillingPhoneNumber() bool {
	if o != nil && !IsNil(o.BillingPhoneNumber) {
		return true
	}

	return false
}

// SetBillingPhoneNumber gets a reference to the given string and assigns it to the BillingPhoneNumber field.
func (o *CustomerServiceRecordAdditionalData) SetBillingPhoneNumber(v string) {
	o.BillingPhoneNumber = &v
}

func (o CustomerServiceRecordAdditionalData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerServiceRecordAdditionalData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AuthorizedPersonName) {
		toSerialize["authorized_person_name"] = o.AuthorizedPersonName
	}
	if !IsNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if !IsNil(o.CustomerCode) {
		toSerialize["customer_code"] = o.CustomerCode
	}
	if !IsNil(o.AddressLine1) {
		toSerialize["address_line_1"] = o.AddressLine1
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ZipCode) {
		toSerialize["zip_code"] = o.ZipCode
	}
	if !IsNil(o.BillingPhoneNumber) {
		toSerialize["billing_phone_number"] = o.BillingPhoneNumber
	}
	return toSerialize, nil
}

type NullableCustomerServiceRecordAdditionalData struct {
	value *CustomerServiceRecordAdditionalData
	isSet bool
}

func (v NullableCustomerServiceRecordAdditionalData) Get() *CustomerServiceRecordAdditionalData {
	return v.value
}

func (v *NullableCustomerServiceRecordAdditionalData) Set(val *CustomerServiceRecordAdditionalData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerServiceRecordAdditionalData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerServiceRecordAdditionalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerServiceRecordAdditionalData(val *CustomerServiceRecordAdditionalData) *NullableCustomerServiceRecordAdditionalData {
	return &NullableCustomerServiceRecordAdditionalData{value: val, isSet: true}
}

func (v NullableCustomerServiceRecordAdditionalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerServiceRecordAdditionalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



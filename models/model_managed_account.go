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
	"bytes"
	"fmt"
)

// checks if the ManagedAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedAccount{}

// ManagedAccount struct for ManagedAccount
type ManagedAccount struct {
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Uniquely identifies the managed account.
	Id string `json:"id"`
	// The managed account's email.
	Email string `json:"email"`
	// The managed account's V2 API access key
	ApiKey string `json:"api_key"`
	// The manager account's email, which serves as the V1 API user identifier
	ApiUser string `json:"api_user"`
	// The managed account's V1 API token
	ApiToken string `json:"api_token"`
	// The organization the managed account is associated with.
	OrganizationName *string `json:"organization_name,omitempty"`
	// The ID of the manager account associated with the managed account.
	ManagerAccountId string `json:"manager_account_id"`
	Balance *ManagedAccountBalance `json:"balance,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// Boolean value that indicates if the managed account is able to have custom pricing set for it or not. If false, uses the pricing of the manager account. Defaults to false. There may be time lag between when the value is changed and pricing changes take effect.
	ManagedAccountAllowCustomPricing *bool `json:"managed_account_allow_custom_pricing,omitempty"`
	// Boolean value that indicates if the billing information and charges to the managed account \"roll up\" to the manager account. If true, the managed account will not have its own balance and will use the shared balance with the manager account. This value cannot be changed after account creation without going through Telnyx support as changes require manual updates to the account ledger. Defaults to false.
	RollupBilling *bool `json:"rollup_billing,omitempty"`
}

type _ManagedAccount ManagedAccount

// NewManagedAccount instantiates a new ManagedAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedAccount(recordType string, id string, email string, apiKey string, apiUser string, apiToken string, managerAccountId string, createdAt string, updatedAt string) *ManagedAccount {
	this := ManagedAccount{}
	this.RecordType = recordType
	this.Id = id
	this.Email = email
	this.ApiKey = apiKey
	this.ApiUser = apiUser
	this.ApiToken = apiToken
	this.ManagerAccountId = managerAccountId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewManagedAccountWithDefaults instantiates a new ManagedAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedAccountWithDefaults() *ManagedAccount {
	this := ManagedAccount{}
	return &this
}

// GetRecordType returns the RecordType field value
func (o *ManagedAccount) GetRecordType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetRecordTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordType, true
}

// SetRecordType sets field value
func (o *ManagedAccount) SetRecordType(v string) {
	o.RecordType = v
}

// GetId returns the Id field value
func (o *ManagedAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ManagedAccount) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value
func (o *ManagedAccount) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ManagedAccount) SetEmail(v string) {
	o.Email = v
}

// GetApiKey returns the ApiKey field value
func (o *ManagedAccount) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *ManagedAccount) SetApiKey(v string) {
	o.ApiKey = v
}

// GetApiUser returns the ApiUser field value
func (o *ManagedAccount) GetApiUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiUser
}

// GetApiUserOk returns a tuple with the ApiUser field value
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetApiUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiUser, true
}

// SetApiUser sets field value
func (o *ManagedAccount) SetApiUser(v string) {
	o.ApiUser = v
}

// GetApiToken returns the ApiToken field value
func (o *ManagedAccount) GetApiToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiToken, true
}

// SetApiToken sets field value
func (o *ManagedAccount) SetApiToken(v string) {
	o.ApiToken = v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *ManagedAccount) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *ManagedAccount) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *ManagedAccount) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetManagerAccountId returns the ManagerAccountId field value
func (o *ManagedAccount) GetManagerAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManagerAccountId
}

// GetManagerAccountIdOk returns a tuple with the ManagerAccountId field value
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetManagerAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagerAccountId, true
}

// SetManagerAccountId sets field value
func (o *ManagedAccount) SetManagerAccountId(v string) {
	o.ManagerAccountId = v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *ManagedAccount) GetBalance() ManagedAccountBalance {
	if o == nil || IsNil(o.Balance) {
		var ret ManagedAccountBalance
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetBalanceOk() (*ManagedAccountBalance, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *ManagedAccount) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given ManagedAccountBalance and assigns it to the Balance field.
func (o *ManagedAccount) SetBalance(v ManagedAccountBalance) {
	o.Balance = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ManagedAccount) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ManagedAccount) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ManagedAccount) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ManagedAccount) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetManagedAccountAllowCustomPricing returns the ManagedAccountAllowCustomPricing field value if set, zero value otherwise.
func (o *ManagedAccount) GetManagedAccountAllowCustomPricing() bool {
	if o == nil || IsNil(o.ManagedAccountAllowCustomPricing) {
		var ret bool
		return ret
	}
	return *o.ManagedAccountAllowCustomPricing
}

// GetManagedAccountAllowCustomPricingOk returns a tuple with the ManagedAccountAllowCustomPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetManagedAccountAllowCustomPricingOk() (*bool, bool) {
	if o == nil || IsNil(o.ManagedAccountAllowCustomPricing) {
		return nil, false
	}
	return o.ManagedAccountAllowCustomPricing, true
}

// HasManagedAccountAllowCustomPricing returns a boolean if a field has been set.
func (o *ManagedAccount) HasManagedAccountAllowCustomPricing() bool {
	if o != nil && !IsNil(o.ManagedAccountAllowCustomPricing) {
		return true
	}

	return false
}

// SetManagedAccountAllowCustomPricing gets a reference to the given bool and assigns it to the ManagedAccountAllowCustomPricing field.
func (o *ManagedAccount) SetManagedAccountAllowCustomPricing(v bool) {
	o.ManagedAccountAllowCustomPricing = &v
}

// GetRollupBilling returns the RollupBilling field value if set, zero value otherwise.
func (o *ManagedAccount) GetRollupBilling() bool {
	if o == nil || IsNil(o.RollupBilling) {
		var ret bool
		return ret
	}
	return *o.RollupBilling
}

// GetRollupBillingOk returns a tuple with the RollupBilling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAccount) GetRollupBillingOk() (*bool, bool) {
	if o == nil || IsNil(o.RollupBilling) {
		return nil, false
	}
	return o.RollupBilling, true
}

// HasRollupBilling returns a boolean if a field has been set.
func (o *ManagedAccount) HasRollupBilling() bool {
	if o != nil && !IsNil(o.RollupBilling) {
		return true
	}

	return false
}

// SetRollupBilling gets a reference to the given bool and assigns it to the RollupBilling field.
func (o *ManagedAccount) SetRollupBilling(v bool) {
	o.RollupBilling = &v
}

func (o ManagedAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["record_type"] = o.RecordType
	toSerialize["id"] = o.Id
	toSerialize["email"] = o.Email
	toSerialize["api_key"] = o.ApiKey
	toSerialize["api_user"] = o.ApiUser
	toSerialize["api_token"] = o.ApiToken
	if !IsNil(o.OrganizationName) {
		toSerialize["organization_name"] = o.OrganizationName
	}
	toSerialize["manager_account_id"] = o.ManagerAccountId
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.ManagedAccountAllowCustomPricing) {
		toSerialize["managed_account_allow_custom_pricing"] = o.ManagedAccountAllowCustomPricing
	}
	if !IsNil(o.RollupBilling) {
		toSerialize["rollup_billing"] = o.RollupBilling
	}
	return toSerialize, nil
}

func (o *ManagedAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"record_type",
		"id",
		"email",
		"api_key",
		"api_user",
		"api_token",
		"manager_account_id",
		"created_at",
		"updated_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varManagedAccount := _ManagedAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagedAccount)

	if err != nil {
		return err
	}

	*o = ManagedAccount(varManagedAccount)

	return err
}

type NullableManagedAccount struct {
	value *ManagedAccount
	isSet bool
}

func (v NullableManagedAccount) Get() *ManagedAccount {
	return v.value
}

func (v *NullableManagedAccount) Set(val *ManagedAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedAccount(val *ManagedAccount) *NullableManagedAccount {
	return &NullableManagedAccount{value: val, isSet: true}
}

func (v NullableManagedAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



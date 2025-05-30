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

// checks if the SIMCardOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SIMCardOrder{}

// SIMCardOrder struct for SIMCardOrder
type SIMCardOrder struct {
	// Identifies the resource.
	Id *string `json:"id,omitempty"`
	// Identifies the type of the resource.
	RecordType *string `json:"record_type,omitempty"`
	// The amount of SIM cards requested in the SIM card order.
	Quantity *int32 `json:"quantity,omitempty"`
	Cost *SIMCardOrderCost `json:"cost,omitempty"`
	OrderAddress *SIMCardOrderOrderAddress `json:"order_address,omitempty"`
	// The URL used to get tracking information about the order.
	TrackingUrl *string `json:"tracking_url,omitempty"`
	// The current status of the SIM Card order.<ul> <li><code>pending</code> - the order is waiting to be processed.</li> <li><code>processing</code> - the order is currently being processed.</li> <li><code>ready_to_ship</code> - the order is ready to be shipped to the specified <b>address</b>.</li> <li><code>shipped</code> - the order was shipped and is on its way to be delivered to the specified <b>address</b>.</li> <li><code>delivered</code> - the order was delivered to the specified <b>address</b>.</li> <li><code>canceled</code> - the order was canceled.</li> </ul>
	Status *string `json:"status,omitempty"`
	// ISO 8601 formatted date-time indicating when the resource was last created.
	CreatedAt *string `json:"created_at,omitempty"`
	// ISO 8601 formatted date-time indicating when the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewSIMCardOrder instantiates a new SIMCardOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSIMCardOrder() *SIMCardOrder {
	this := SIMCardOrder{}
	return &this
}

// NewSIMCardOrderWithDefaults instantiates a new SIMCardOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSIMCardOrderWithDefaults() *SIMCardOrder {
	this := SIMCardOrder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SIMCardOrder) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardOrder) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SIMCardOrder) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SIMCardOrder) SetId(v string) {
	o.Id = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *SIMCardOrder) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardOrder) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *SIMCardOrder) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *SIMCardOrder) SetRecordType(v string) {
	o.RecordType = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SIMCardOrder) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardOrder) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SIMCardOrder) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *SIMCardOrder) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *SIMCardOrder) GetCost() SIMCardOrderCost {
	if o == nil || IsNil(o.Cost) {
		var ret SIMCardOrderCost
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardOrder) GetCostOk() (*SIMCardOrderCost, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *SIMCardOrder) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given SIMCardOrderCost and assigns it to the Cost field.
func (o *SIMCardOrder) SetCost(v SIMCardOrderCost) {
	o.Cost = &v
}

// GetOrderAddress returns the OrderAddress field value if set, zero value otherwise.
func (o *SIMCardOrder) GetOrderAddress() SIMCardOrderOrderAddress {
	if o == nil || IsNil(o.OrderAddress) {
		var ret SIMCardOrderOrderAddress
		return ret
	}
	return *o.OrderAddress
}

// GetOrderAddressOk returns a tuple with the OrderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardOrder) GetOrderAddressOk() (*SIMCardOrderOrderAddress, bool) {
	if o == nil || IsNil(o.OrderAddress) {
		return nil, false
	}
	return o.OrderAddress, true
}

// HasOrderAddress returns a boolean if a field has been set.
func (o *SIMCardOrder) HasOrderAddress() bool {
	if o != nil && !IsNil(o.OrderAddress) {
		return true
	}

	return false
}

// SetOrderAddress gets a reference to the given SIMCardOrderOrderAddress and assigns it to the OrderAddress field.
func (o *SIMCardOrder) SetOrderAddress(v SIMCardOrderOrderAddress) {
	o.OrderAddress = &v
}

// GetTrackingUrl returns the TrackingUrl field value if set, zero value otherwise.
func (o *SIMCardOrder) GetTrackingUrl() string {
	if o == nil || IsNil(o.TrackingUrl) {
		var ret string
		return ret
	}
	return *o.TrackingUrl
}

// GetTrackingUrlOk returns a tuple with the TrackingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardOrder) GetTrackingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingUrl) {
		return nil, false
	}
	return o.TrackingUrl, true
}

// HasTrackingUrl returns a boolean if a field has been set.
func (o *SIMCardOrder) HasTrackingUrl() bool {
	if o != nil && !IsNil(o.TrackingUrl) {
		return true
	}

	return false
}

// SetTrackingUrl gets a reference to the given string and assigns it to the TrackingUrl field.
func (o *SIMCardOrder) SetTrackingUrl(v string) {
	o.TrackingUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SIMCardOrder) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardOrder) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SIMCardOrder) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SIMCardOrder) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SIMCardOrder) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardOrder) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SIMCardOrder) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SIMCardOrder) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SIMCardOrder) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardOrder) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SIMCardOrder) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *SIMCardOrder) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o SIMCardOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SIMCardOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.OrderAddress) {
		toSerialize["order_address"] = o.OrderAddress
	}
	if !IsNil(o.TrackingUrl) {
		toSerialize["tracking_url"] = o.TrackingUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableSIMCardOrder struct {
	value *SIMCardOrder
	isSet bool
}

func (v NullableSIMCardOrder) Get() *SIMCardOrder {
	return v.value
}

func (v *NullableSIMCardOrder) Set(val *SIMCardOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableSIMCardOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableSIMCardOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSIMCardOrder(val *SIMCardOrder) *NullableSIMCardOrder {
	return &NullableSIMCardOrder{value: val, isSet: true}
}

func (v NullableSIMCardOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSIMCardOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



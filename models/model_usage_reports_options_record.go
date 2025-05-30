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

// checks if the UsageReportsOptionsRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageReportsOptionsRecord{}

// UsageReportsOptionsRecord An object following one of the schemas published in https://developers.telnyx.com/docs/api/v2/detail-records
type UsageReportsOptionsRecord struct {
	// Telnyx Product
	Product *string `json:"product,omitempty"`
	// Telnyx Product Dimensions
	ProductDimensions []string `json:"product_dimensions,omitempty"`
	// Telnyx Product Metrics
	ProductMetrics []string `json:"product_metrics,omitempty"`
	// Subproducts if applicable
	RecordTypes []RecordType `json:"record_types,omitempty"`
}

// NewUsageReportsOptionsRecord instantiates a new UsageReportsOptionsRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageReportsOptionsRecord() *UsageReportsOptionsRecord {
	this := UsageReportsOptionsRecord{}
	return &this
}

// NewUsageReportsOptionsRecordWithDefaults instantiates a new UsageReportsOptionsRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageReportsOptionsRecordWithDefaults() *UsageReportsOptionsRecord {
	this := UsageReportsOptionsRecord{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *UsageReportsOptionsRecord) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageReportsOptionsRecord) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *UsageReportsOptionsRecord) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *UsageReportsOptionsRecord) SetProduct(v string) {
	o.Product = &v
}

// GetProductDimensions returns the ProductDimensions field value if set, zero value otherwise.
func (o *UsageReportsOptionsRecord) GetProductDimensions() []string {
	if o == nil || IsNil(o.ProductDimensions) {
		var ret []string
		return ret
	}
	return o.ProductDimensions
}

// GetProductDimensionsOk returns a tuple with the ProductDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageReportsOptionsRecord) GetProductDimensionsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductDimensions) {
		return nil, false
	}
	return o.ProductDimensions, true
}

// HasProductDimensions returns a boolean if a field has been set.
func (o *UsageReportsOptionsRecord) HasProductDimensions() bool {
	if o != nil && !IsNil(o.ProductDimensions) {
		return true
	}

	return false
}

// SetProductDimensions gets a reference to the given []string and assigns it to the ProductDimensions field.
func (o *UsageReportsOptionsRecord) SetProductDimensions(v []string) {
	o.ProductDimensions = v
}

// GetProductMetrics returns the ProductMetrics field value if set, zero value otherwise.
func (o *UsageReportsOptionsRecord) GetProductMetrics() []string {
	if o == nil || IsNil(o.ProductMetrics) {
		var ret []string
		return ret
	}
	return o.ProductMetrics
}

// GetProductMetricsOk returns a tuple with the ProductMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageReportsOptionsRecord) GetProductMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductMetrics) {
		return nil, false
	}
	return o.ProductMetrics, true
}

// HasProductMetrics returns a boolean if a field has been set.
func (o *UsageReportsOptionsRecord) HasProductMetrics() bool {
	if o != nil && !IsNil(o.ProductMetrics) {
		return true
	}

	return false
}

// SetProductMetrics gets a reference to the given []string and assigns it to the ProductMetrics field.
func (o *UsageReportsOptionsRecord) SetProductMetrics(v []string) {
	o.ProductMetrics = v
}

// GetRecordTypes returns the RecordTypes field value if set, zero value otherwise.
func (o *UsageReportsOptionsRecord) GetRecordTypes() []RecordType {
	if o == nil || IsNil(o.RecordTypes) {
		var ret []RecordType
		return ret
	}
	return o.RecordTypes
}

// GetRecordTypesOk returns a tuple with the RecordTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageReportsOptionsRecord) GetRecordTypesOk() ([]RecordType, bool) {
	if o == nil || IsNil(o.RecordTypes) {
		return nil, false
	}
	return o.RecordTypes, true
}

// HasRecordTypes returns a boolean if a field has been set.
func (o *UsageReportsOptionsRecord) HasRecordTypes() bool {
	if o != nil && !IsNil(o.RecordTypes) {
		return true
	}

	return false
}

// SetRecordTypes gets a reference to the given []RecordType and assigns it to the RecordTypes field.
func (o *UsageReportsOptionsRecord) SetRecordTypes(v []RecordType) {
	o.RecordTypes = v
}

func (o UsageReportsOptionsRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageReportsOptionsRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.ProductDimensions) {
		toSerialize["product_dimensions"] = o.ProductDimensions
	}
	if !IsNil(o.ProductMetrics) {
		toSerialize["product_metrics"] = o.ProductMetrics
	}
	if !IsNil(o.RecordTypes) {
		toSerialize["record_types"] = o.RecordTypes
	}
	return toSerialize, nil
}

type NullableUsageReportsOptionsRecord struct {
	value *UsageReportsOptionsRecord
	isSet bool
}

func (v NullableUsageReportsOptionsRecord) Get() *UsageReportsOptionsRecord {
	return v.value
}

func (v *NullableUsageReportsOptionsRecord) Set(val *UsageReportsOptionsRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageReportsOptionsRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageReportsOptionsRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageReportsOptionsRecord(val *UsageReportsOptionsRecord) *NullableUsageReportsOptionsRecord {
	return &NullableUsageReportsOptionsRecord{value: val, isSet: true}
}

func (v NullableUsageReportsOptionsRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageReportsOptionsRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



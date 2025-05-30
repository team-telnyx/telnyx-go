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
	"time"
)

// checks if the CdrUsageReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdrUsageReportResponse{}

// CdrUsageReportResponse struct for CdrUsageReportResponse
type CdrUsageReportResponse struct {
	// Identifies the resource
	Id *string `json:"id,omitempty"`
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime *time.Time `json:"end_time,omitempty"`
	Connections []int64 `json:"connections,omitempty"`
	AggregationType *string `json:"aggregation_type,omitempty"`
	Status *string `json:"status,omitempty"`
	ReportUrl *string `json:"report_url,omitempty"`
	Result map[string]interface{} `json:"result,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	RecordType *string `json:"record_type,omitempty"`
	ProductBreakdown *string `json:"product_breakdown,omitempty"`
}

// NewCdrUsageReportResponse instantiates a new CdrUsageReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdrUsageReportResponse() *CdrUsageReportResponse {
	this := CdrUsageReportResponse{}
	return &this
}

// NewCdrUsageReportResponseWithDefaults instantiates a new CdrUsageReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdrUsageReportResponseWithDefaults() *CdrUsageReportResponse {
	this := CdrUsageReportResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CdrUsageReportResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrUsageReportResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CdrUsageReportResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CdrUsageReportResponse) SetId(v string) {
	o.Id = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *CdrUsageReportResponse) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrUsageReportResponse) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *CdrUsageReportResponse) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *CdrUsageReportResponse) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *CdrUsageReportResponse) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrUsageReportResponse) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *CdrUsageReportResponse) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *CdrUsageReportResponse) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *CdrUsageReportResponse) GetConnections() []int64 {
	if o == nil || IsNil(o.Connections) {
		var ret []int64
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrUsageReportResponse) GetConnectionsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Connections) {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *CdrUsageReportResponse) HasConnections() bool {
	if o != nil && !IsNil(o.Connections) {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []int64 and assigns it to the Connections field.
func (o *CdrUsageReportResponse) SetConnections(v []int64) {
	o.Connections = v
}

// GetAggregationType returns the AggregationType field value if set, zero value otherwise.
func (o *CdrUsageReportResponse) GetAggregationType() string {
	if o == nil || IsNil(o.AggregationType) {
		var ret string
		return ret
	}
	return *o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrUsageReportResponse) GetAggregationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationType) {
		return nil, false
	}
	return o.AggregationType, true
}

// HasAggregationType returns a boolean if a field has been set.
func (o *CdrUsageReportResponse) HasAggregationType() bool {
	if o != nil && !IsNil(o.AggregationType) {
		return true
	}

	return false
}

// SetAggregationType gets a reference to the given string and assigns it to the AggregationType field.
func (o *CdrUsageReportResponse) SetAggregationType(v string) {
	o.AggregationType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CdrUsageReportResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrUsageReportResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CdrUsageReportResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CdrUsageReportResponse) SetStatus(v string) {
	o.Status = &v
}

// GetReportUrl returns the ReportUrl field value if set, zero value otherwise.
func (o *CdrUsageReportResponse) GetReportUrl() string {
	if o == nil || IsNil(o.ReportUrl) {
		var ret string
		return ret
	}
	return *o.ReportUrl
}

// GetReportUrlOk returns a tuple with the ReportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrUsageReportResponse) GetReportUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportUrl) {
		return nil, false
	}
	return o.ReportUrl, true
}

// HasReportUrl returns a boolean if a field has been set.
func (o *CdrUsageReportResponse) HasReportUrl() bool {
	if o != nil && !IsNil(o.ReportUrl) {
		return true
	}

	return false
}

// SetReportUrl gets a reference to the given string and assigns it to the ReportUrl field.
func (o *CdrUsageReportResponse) SetReportUrl(v string) {
	o.ReportUrl = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CdrUsageReportResponse) GetResult() map[string]interface{} {
	if o == nil || IsNil(o.Result) {
		var ret map[string]interface{}
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrUsageReportResponse) GetResultOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Result) {
		return map[string]interface{}{}, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CdrUsageReportResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given map[string]interface{} and assigns it to the Result field.
func (o *CdrUsageReportResponse) SetResult(v map[string]interface{}) {
	o.Result = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CdrUsageReportResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrUsageReportResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CdrUsageReportResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CdrUsageReportResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CdrUsageReportResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrUsageReportResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CdrUsageReportResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CdrUsageReportResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *CdrUsageReportResponse) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrUsageReportResponse) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *CdrUsageReportResponse) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *CdrUsageReportResponse) SetRecordType(v string) {
	o.RecordType = &v
}

// GetProductBreakdown returns the ProductBreakdown field value if set, zero value otherwise.
func (o *CdrUsageReportResponse) GetProductBreakdown() string {
	if o == nil || IsNil(o.ProductBreakdown) {
		var ret string
		return ret
	}
	return *o.ProductBreakdown
}

// GetProductBreakdownOk returns a tuple with the ProductBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrUsageReportResponse) GetProductBreakdownOk() (*string, bool) {
	if o == nil || IsNil(o.ProductBreakdown) {
		return nil, false
	}
	return o.ProductBreakdown, true
}

// HasProductBreakdown returns a boolean if a field has been set.
func (o *CdrUsageReportResponse) HasProductBreakdown() bool {
	if o != nil && !IsNil(o.ProductBreakdown) {
		return true
	}

	return false
}

// SetProductBreakdown gets a reference to the given string and assigns it to the ProductBreakdown field.
func (o *CdrUsageReportResponse) SetProductBreakdown(v string) {
	o.ProductBreakdown = &v
}

func (o CdrUsageReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdrUsageReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.Connections) {
		toSerialize["connections"] = o.Connections
	}
	if !IsNil(o.AggregationType) {
		toSerialize["aggregation_type"] = o.AggregationType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ReportUrl) {
		toSerialize["report_url"] = o.ReportUrl
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.ProductBreakdown) {
		toSerialize["product_breakdown"] = o.ProductBreakdown
	}
	return toSerialize, nil
}

type NullableCdrUsageReportResponse struct {
	value *CdrUsageReportResponse
	isSet bool
}

func (v NullableCdrUsageReportResponse) Get() *CdrUsageReportResponse {
	return v.value
}

func (v *NullableCdrUsageReportResponse) Set(val *CdrUsageReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCdrUsageReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCdrUsageReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdrUsageReportResponse(val *CdrUsageReportResponse) *NullableCdrUsageReportResponse {
	return &NullableCdrUsageReportResponse{value: val, isSet: true}
}

func (v NullableCdrUsageReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdrUsageReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the PrivateWirelessGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateWirelessGateway{}

// PrivateWirelessGateway struct for PrivateWirelessGateway
type PrivateWirelessGateway struct {
	// Identifies the resource.
	Id *string `json:"id,omitempty"`
	// The identification of the related network resource.
	NetworkId *string `json:"network_id,omitempty"`
	RecordType *string `json:"record_type,omitempty"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// The private wireless gateway name.
	Name *string `json:"name,omitempty"`
	// The name of the region where the Private Wireless Gateway is deployed.
	RegionCode *string `json:"region_code,omitempty"`
	Status *PrivateWirelessGatewayStatus `json:"status,omitempty"`
	// IP block used to assign IPs to the SIM cards in the Private Wireless Gateway.
	IpRange *string `json:"ip_range,omitempty"`
	// A list of the resources that have been assigned to the Private Wireless Gateway.
	AssignedResources []PWGAssignedResourcesSummary `json:"assigned_resources,omitempty"`
}

// NewPrivateWirelessGateway instantiates a new PrivateWirelessGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateWirelessGateway() *PrivateWirelessGateway {
	this := PrivateWirelessGateway{}
	return &this
}

// NewPrivateWirelessGatewayWithDefaults instantiates a new PrivateWirelessGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateWirelessGatewayWithDefaults() *PrivateWirelessGateway {
	this := PrivateWirelessGateway{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PrivateWirelessGateway) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateWirelessGateway) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PrivateWirelessGateway) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PrivateWirelessGateway) SetId(v string) {
	o.Id = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *PrivateWirelessGateway) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateWirelessGateway) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *PrivateWirelessGateway) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *PrivateWirelessGateway) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *PrivateWirelessGateway) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateWirelessGateway) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *PrivateWirelessGateway) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *PrivateWirelessGateway) SetRecordType(v string) {
	o.RecordType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PrivateWirelessGateway) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateWirelessGateway) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PrivateWirelessGateway) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PrivateWirelessGateway) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PrivateWirelessGateway) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateWirelessGateway) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PrivateWirelessGateway) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *PrivateWirelessGateway) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PrivateWirelessGateway) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateWirelessGateway) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PrivateWirelessGateway) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PrivateWirelessGateway) SetName(v string) {
	o.Name = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *PrivateWirelessGateway) GetRegionCode() string {
	if o == nil || IsNil(o.RegionCode) {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateWirelessGateway) GetRegionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RegionCode) {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *PrivateWirelessGateway) HasRegionCode() bool {
	if o != nil && !IsNil(o.RegionCode) {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *PrivateWirelessGateway) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PrivateWirelessGateway) GetStatus() PrivateWirelessGatewayStatus {
	if o == nil || IsNil(o.Status) {
		var ret PrivateWirelessGatewayStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateWirelessGateway) GetStatusOk() (*PrivateWirelessGatewayStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PrivateWirelessGateway) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PrivateWirelessGatewayStatus and assigns it to the Status field.
func (o *PrivateWirelessGateway) SetStatus(v PrivateWirelessGatewayStatus) {
	o.Status = &v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *PrivateWirelessGateway) GetIpRange() string {
	if o == nil || IsNil(o.IpRange) {
		var ret string
		return ret
	}
	return *o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateWirelessGateway) GetIpRangeOk() (*string, bool) {
	if o == nil || IsNil(o.IpRange) {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *PrivateWirelessGateway) HasIpRange() bool {
	if o != nil && !IsNil(o.IpRange) {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given string and assigns it to the IpRange field.
func (o *PrivateWirelessGateway) SetIpRange(v string) {
	o.IpRange = &v
}

// GetAssignedResources returns the AssignedResources field value if set, zero value otherwise.
func (o *PrivateWirelessGateway) GetAssignedResources() []PWGAssignedResourcesSummary {
	if o == nil || IsNil(o.AssignedResources) {
		var ret []PWGAssignedResourcesSummary
		return ret
	}
	return o.AssignedResources
}

// GetAssignedResourcesOk returns a tuple with the AssignedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateWirelessGateway) GetAssignedResourcesOk() ([]PWGAssignedResourcesSummary, bool) {
	if o == nil || IsNil(o.AssignedResources) {
		return nil, false
	}
	return o.AssignedResources, true
}

// HasAssignedResources returns a boolean if a field has been set.
func (o *PrivateWirelessGateway) HasAssignedResources() bool {
	if o != nil && !IsNil(o.AssignedResources) {
		return true
	}

	return false
}

// SetAssignedResources gets a reference to the given []PWGAssignedResourcesSummary and assigns it to the AssignedResources field.
func (o *PrivateWirelessGateway) SetAssignedResources(v []PWGAssignedResourcesSummary) {
	o.AssignedResources = v
}

func (o PrivateWirelessGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateWirelessGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NetworkId) {
		toSerialize["network_id"] = o.NetworkId
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RegionCode) {
		toSerialize["region_code"] = o.RegionCode
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.IpRange) {
		toSerialize["ip_range"] = o.IpRange
	}
	if !IsNil(o.AssignedResources) {
		toSerialize["assigned_resources"] = o.AssignedResources
	}
	return toSerialize, nil
}

type NullablePrivateWirelessGateway struct {
	value *PrivateWirelessGateway
	isSet bool
}

func (v NullablePrivateWirelessGateway) Get() *PrivateWirelessGateway {
	return v.value
}

func (v *NullablePrivateWirelessGateway) Set(val *PrivateWirelessGateway) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateWirelessGateway) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateWirelessGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateWirelessGateway(val *PrivateWirelessGateway) *NullablePrivateWirelessGateway {
	return &NullablePrivateWirelessGateway{value: val, isSet: true}
}

func (v NullablePrivateWirelessGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateWirelessGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ListTextToSpeechVoices200ResponseVoicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTextToSpeechVoices200ResponseVoicesInner{}

// ListTextToSpeechVoices200ResponseVoicesInner struct for ListTextToSpeechVoices200ResponseVoicesInner
type ListTextToSpeechVoices200ResponseVoicesInner struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Label *string `json:"label,omitempty"`
	Accent *string `json:"accent,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Age *string `json:"age,omitempty"`
	Language *string `json:"language,omitempty"`
}

// NewListTextToSpeechVoices200ResponseVoicesInner instantiates a new ListTextToSpeechVoices200ResponseVoicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTextToSpeechVoices200ResponseVoicesInner() *ListTextToSpeechVoices200ResponseVoicesInner {
	this := ListTextToSpeechVoices200ResponseVoicesInner{}
	return &this
}

// NewListTextToSpeechVoices200ResponseVoicesInnerWithDefaults instantiates a new ListTextToSpeechVoices200ResponseVoicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTextToSpeechVoices200ResponseVoicesInnerWithDefaults() *ListTextToSpeechVoices200ResponseVoicesInner {
	this := ListTextToSpeechVoices200ResponseVoicesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) SetProvider(v string) {
	o.Provider = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) SetLabel(v string) {
	o.Label = &v
}

// GetAccent returns the Accent field value if set, zero value otherwise.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetAccent() string {
	if o == nil || IsNil(o.Accent) {
		var ret string
		return ret
	}
	return *o.Accent
}

// GetAccentOk returns a tuple with the Accent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetAccentOk() (*string, bool) {
	if o == nil || IsNil(o.Accent) {
		return nil, false
	}
	return o.Accent, true
}

// HasAccent returns a boolean if a field has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) HasAccent() bool {
	if o != nil && !IsNil(o.Accent) {
		return true
	}

	return false
}

// SetAccent gets a reference to the given string and assigns it to the Accent field.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) SetAccent(v string) {
	o.Accent = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) SetGender(v string) {
	o.Gender = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetAge() string {
	if o == nil || IsNil(o.Age) {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetAgeOk() (*string, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) SetAge(v string) {
	o.Age = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ListTextToSpeechVoices200ResponseVoicesInner) SetLanguage(v string) {
	o.Language = &v
}

func (o ListTextToSpeechVoices200ResponseVoicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTextToSpeechVoices200ResponseVoicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Accent) {
		toSerialize["accent"] = o.Accent
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.Age) {
		toSerialize["age"] = o.Age
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

type NullableListTextToSpeechVoices200ResponseVoicesInner struct {
	value *ListTextToSpeechVoices200ResponseVoicesInner
	isSet bool
}

func (v NullableListTextToSpeechVoices200ResponseVoicesInner) Get() *ListTextToSpeechVoices200ResponseVoicesInner {
	return v.value
}

func (v *NullableListTextToSpeechVoices200ResponseVoicesInner) Set(val *ListTextToSpeechVoices200ResponseVoicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListTextToSpeechVoices200ResponseVoicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListTextToSpeechVoices200ResponseVoicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTextToSpeechVoices200ResponseVoicesInner(val *ListTextToSpeechVoices200ResponseVoicesInner) *NullableListTextToSpeechVoices200ResponseVoicesInner {
	return &NullableListTextToSpeechVoices200ResponseVoicesInner{value: val, isSet: true}
}

func (v NullableListTextToSpeechVoices200ResponseVoicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTextToSpeechVoices200ResponseVoicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



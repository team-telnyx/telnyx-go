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
	"fmt"
	"gopkg.in/validator.v2"
)

// DetailRecord - An object following one of the schemas published in https://developers.telnyx.com/docs/api/v2/detail-records
type DetailRecord struct {
	AmdDetailRecord *AmdDetailRecord
	ConferenceDetailRecord *ConferenceDetailRecord
	ConferenceParticipantDetailRecord *ConferenceParticipantDetailRecord
	MediaStorageDetailRecord *MediaStorageDetailRecord
	MessageDetailRecord *MessageDetailRecord
	SimCardUsageDetailRecord *SimCardUsageDetailRecord
	VerifyDetailRecord *VerifyDetailRecord
}

// AmdDetailRecordAsDetailRecord is a convenience function that returns AmdDetailRecord wrapped in DetailRecord
func AmdDetailRecordAsDetailRecord(v *AmdDetailRecord) DetailRecord {
	return DetailRecord{
		AmdDetailRecord: v,
	}
}

// ConferenceDetailRecordAsDetailRecord is a convenience function that returns ConferenceDetailRecord wrapped in DetailRecord
func ConferenceDetailRecordAsDetailRecord(v *ConferenceDetailRecord) DetailRecord {
	return DetailRecord{
		ConferenceDetailRecord: v,
	}
}

// ConferenceParticipantDetailRecordAsDetailRecord is a convenience function that returns ConferenceParticipantDetailRecord wrapped in DetailRecord
func ConferenceParticipantDetailRecordAsDetailRecord(v *ConferenceParticipantDetailRecord) DetailRecord {
	return DetailRecord{
		ConferenceParticipantDetailRecord: v,
	}
}

// MediaStorageDetailRecordAsDetailRecord is a convenience function that returns MediaStorageDetailRecord wrapped in DetailRecord
func MediaStorageDetailRecordAsDetailRecord(v *MediaStorageDetailRecord) DetailRecord {
	return DetailRecord{
		MediaStorageDetailRecord: v,
	}
}

// MessageDetailRecordAsDetailRecord is a convenience function that returns MessageDetailRecord wrapped in DetailRecord
func MessageDetailRecordAsDetailRecord(v *MessageDetailRecord) DetailRecord {
	return DetailRecord{
		MessageDetailRecord: v,
	}
}

// SimCardUsageDetailRecordAsDetailRecord is a convenience function that returns SimCardUsageDetailRecord wrapped in DetailRecord
func SimCardUsageDetailRecordAsDetailRecord(v *SimCardUsageDetailRecord) DetailRecord {
	return DetailRecord{
		SimCardUsageDetailRecord: v,
	}
}

// VerifyDetailRecordAsDetailRecord is a convenience function that returns VerifyDetailRecord wrapped in DetailRecord
func VerifyDetailRecordAsDetailRecord(v *VerifyDetailRecord) DetailRecord {
	return DetailRecord{
		VerifyDetailRecord: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DetailRecord) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AmdDetailRecord
	err = newStrictDecoder(data).Decode(&dst.AmdDetailRecord)
	if err == nil {
		jsonAmdDetailRecord, _ := json.Marshal(dst.AmdDetailRecord)
		if string(jsonAmdDetailRecord) == "{}" { // empty struct
			dst.AmdDetailRecord = nil
		} else {
			if err = validator.Validate(dst.AmdDetailRecord); err != nil {
				dst.AmdDetailRecord = nil
			} else {
				match++
			}
		}
	} else {
		dst.AmdDetailRecord = nil
	}

	// try to unmarshal data into ConferenceDetailRecord
	err = newStrictDecoder(data).Decode(&dst.ConferenceDetailRecord)
	if err == nil {
		jsonConferenceDetailRecord, _ := json.Marshal(dst.ConferenceDetailRecord)
		if string(jsonConferenceDetailRecord) == "{}" { // empty struct
			dst.ConferenceDetailRecord = nil
		} else {
			if err = validator.Validate(dst.ConferenceDetailRecord); err != nil {
				dst.ConferenceDetailRecord = nil
			} else {
				match++
			}
		}
	} else {
		dst.ConferenceDetailRecord = nil
	}

	// try to unmarshal data into ConferenceParticipantDetailRecord
	err = newStrictDecoder(data).Decode(&dst.ConferenceParticipantDetailRecord)
	if err == nil {
		jsonConferenceParticipantDetailRecord, _ := json.Marshal(dst.ConferenceParticipantDetailRecord)
		if string(jsonConferenceParticipantDetailRecord) == "{}" { // empty struct
			dst.ConferenceParticipantDetailRecord = nil
		} else {
			if err = validator.Validate(dst.ConferenceParticipantDetailRecord); err != nil {
				dst.ConferenceParticipantDetailRecord = nil
			} else {
				match++
			}
		}
	} else {
		dst.ConferenceParticipantDetailRecord = nil
	}

	// try to unmarshal data into MediaStorageDetailRecord
	err = newStrictDecoder(data).Decode(&dst.MediaStorageDetailRecord)
	if err == nil {
		jsonMediaStorageDetailRecord, _ := json.Marshal(dst.MediaStorageDetailRecord)
		if string(jsonMediaStorageDetailRecord) == "{}" { // empty struct
			dst.MediaStorageDetailRecord = nil
		} else {
			if err = validator.Validate(dst.MediaStorageDetailRecord); err != nil {
				dst.MediaStorageDetailRecord = nil
			} else {
				match++
			}
		}
	} else {
		dst.MediaStorageDetailRecord = nil
	}

	// try to unmarshal data into MessageDetailRecord
	err = newStrictDecoder(data).Decode(&dst.MessageDetailRecord)
	if err == nil {
		jsonMessageDetailRecord, _ := json.Marshal(dst.MessageDetailRecord)
		if string(jsonMessageDetailRecord) == "{}" { // empty struct
			dst.MessageDetailRecord = nil
		} else {
			if err = validator.Validate(dst.MessageDetailRecord); err != nil {
				dst.MessageDetailRecord = nil
			} else {
				match++
			}
		}
	} else {
		dst.MessageDetailRecord = nil
	}

	// try to unmarshal data into SimCardUsageDetailRecord
	err = newStrictDecoder(data).Decode(&dst.SimCardUsageDetailRecord)
	if err == nil {
		jsonSimCardUsageDetailRecord, _ := json.Marshal(dst.SimCardUsageDetailRecord)
		if string(jsonSimCardUsageDetailRecord) == "{}" { // empty struct
			dst.SimCardUsageDetailRecord = nil
		} else {
			if err = validator.Validate(dst.SimCardUsageDetailRecord); err != nil {
				dst.SimCardUsageDetailRecord = nil
			} else {
				match++
			}
		}
	} else {
		dst.SimCardUsageDetailRecord = nil
	}

	// try to unmarshal data into VerifyDetailRecord
	err = newStrictDecoder(data).Decode(&dst.VerifyDetailRecord)
	if err == nil {
		jsonVerifyDetailRecord, _ := json.Marshal(dst.VerifyDetailRecord)
		if string(jsonVerifyDetailRecord) == "{}" { // empty struct
			dst.VerifyDetailRecord = nil
		} else {
			if err = validator.Validate(dst.VerifyDetailRecord); err != nil {
				dst.VerifyDetailRecord = nil
			} else {
				match++
			}
		}
	} else {
		dst.VerifyDetailRecord = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AmdDetailRecord = nil
		dst.ConferenceDetailRecord = nil
		dst.ConferenceParticipantDetailRecord = nil
		dst.MediaStorageDetailRecord = nil
		dst.MessageDetailRecord = nil
		dst.SimCardUsageDetailRecord = nil
		dst.VerifyDetailRecord = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DetailRecord)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DetailRecord)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DetailRecord) MarshalJSON() ([]byte, error) {
	if src.AmdDetailRecord != nil {
		return json.Marshal(&src.AmdDetailRecord)
	}

	if src.ConferenceDetailRecord != nil {
		return json.Marshal(&src.ConferenceDetailRecord)
	}

	if src.ConferenceParticipantDetailRecord != nil {
		return json.Marshal(&src.ConferenceParticipantDetailRecord)
	}

	if src.MediaStorageDetailRecord != nil {
		return json.Marshal(&src.MediaStorageDetailRecord)
	}

	if src.MessageDetailRecord != nil {
		return json.Marshal(&src.MessageDetailRecord)
	}

	if src.SimCardUsageDetailRecord != nil {
		return json.Marshal(&src.SimCardUsageDetailRecord)
	}

	if src.VerifyDetailRecord != nil {
		return json.Marshal(&src.VerifyDetailRecord)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DetailRecord) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AmdDetailRecord != nil {
		return obj.AmdDetailRecord
	}

	if obj.ConferenceDetailRecord != nil {
		return obj.ConferenceDetailRecord
	}

	if obj.ConferenceParticipantDetailRecord != nil {
		return obj.ConferenceParticipantDetailRecord
	}

	if obj.MediaStorageDetailRecord != nil {
		return obj.MediaStorageDetailRecord
	}

	if obj.MessageDetailRecord != nil {
		return obj.MessageDetailRecord
	}

	if obj.SimCardUsageDetailRecord != nil {
		return obj.SimCardUsageDetailRecord
	}

	if obj.VerifyDetailRecord != nil {
		return obj.VerifyDetailRecord
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj DetailRecord) GetActualInstanceValue() (interface{}) {
	if obj.AmdDetailRecord != nil {
		return *obj.AmdDetailRecord
	}

	if obj.ConferenceDetailRecord != nil {
		return *obj.ConferenceDetailRecord
	}

	if obj.ConferenceParticipantDetailRecord != nil {
		return *obj.ConferenceParticipantDetailRecord
	}

	if obj.MediaStorageDetailRecord != nil {
		return *obj.MediaStorageDetailRecord
	}

	if obj.MessageDetailRecord != nil {
		return *obj.MessageDetailRecord
	}

	if obj.SimCardUsageDetailRecord != nil {
		return *obj.SimCardUsageDetailRecord
	}

	if obj.VerifyDetailRecord != nil {
		return *obj.VerifyDetailRecord
	}

	// all schemas are nil
	return nil
}

type NullableDetailRecord struct {
	value *DetailRecord
	isSet bool
}

func (v NullableDetailRecord) Get() *DetailRecord {
	return v.value
}

func (v *NullableDetailRecord) Set(val *DetailRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailRecord(val *DetailRecord) *NullableDetailRecord {
	return &NullableDetailRecord{value: val, isSet: true}
}

func (v NullableDetailRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



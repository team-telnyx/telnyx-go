// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/internal/apierror"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type ConnectionsPaginationMeta = shared.ConnectionsPaginationMeta

// This is an alias to an internal type.
type DocReqsRequirementType = shared.DocReqsRequirementType

// Specifies objective criteria for acceptance
//
// This is an alias to an internal type.
type DocReqsRequirementTypeAcceptanceCriteria = shared.DocReqsRequirementTypeAcceptanceCriteria

// Defines the type of this requirement type
//
// This is an alias to an internal type.
type DocReqsRequirementTypeType = shared.DocReqsRequirementTypeType

// Equals "document"
const DocReqsRequirementTypeTypeDocument = shared.DocReqsRequirementTypeTypeDocument

// Equals "address"
const DocReqsRequirementTypeTypeAddress = shared.DocReqsRequirementTypeTypeAddress

// Equals "textual"
const DocReqsRequirementTypeTypeTextual = shared.DocReqsRequirementTypeTypeTextual

// This is an alias to an internal type.
type HostedNumber = shared.HostedNumber

// This is an alias to an internal type.
type HostedNumberStatus = shared.HostedNumberStatus

// Equals "deleted"
const HostedNumberStatusDeleted = shared.HostedNumberStatusDeleted

// Equals "failed"
const HostedNumberStatusFailed = shared.HostedNumberStatusFailed

// Equals "failed_activation"
const HostedNumberStatusFailedActivation = shared.HostedNumberStatusFailedActivation

// Equals "failed_carrier_rejected"
const HostedNumberStatusFailedCarrierRejected = shared.HostedNumberStatusFailedCarrierRejected

// Equals "failed_ineligible_carrier"
const HostedNumberStatusFailedIneligibleCarrier = shared.HostedNumberStatusFailedIneligibleCarrier

// Equals "failed_number_already_hosted"
const HostedNumberStatusFailedNumberAlreadyHosted = shared.HostedNumberStatusFailedNumberAlreadyHosted

// Equals "failed_number_not_found"
const HostedNumberStatusFailedNumberNotFound = shared.HostedNumberStatusFailedNumberNotFound

// Equals "failed_ownership_verification"
const HostedNumberStatusFailedOwnershipVerification = shared.HostedNumberStatusFailedOwnershipVerification

// Equals "failed_timeout"
const HostedNumberStatusFailedTimeout = shared.HostedNumberStatusFailedTimeout

// Equals "pending"
const HostedNumberStatusPending = shared.HostedNumberStatusPending

// Equals "provisioning"
const HostedNumberStatusProvisioning = shared.HostedNumberStatusProvisioning

// Equals "successful"
const HostedNumberStatusSuccessful = shared.HostedNumberStatusSuccessful

// The set of features available for a specific messaging use case (SMS or MMS).
// Features can vary depending on the characteristics the phone number, as well as
// its current product configuration.
//
// This is an alias to an internal type.
type MessagingFeatureSet = shared.MessagingFeatureSet

// This is an alias to an internal type.
type MessagingHostedNumberOrder = shared.MessagingHostedNumberOrder

// This is an alias to an internal type.
type MessagingHostedNumberOrderStatus = shared.MessagingHostedNumberOrderStatus

// Equals "carrier_rejected"
const MessagingHostedNumberOrderStatusCarrierRejected = shared.MessagingHostedNumberOrderStatusCarrierRejected

// Equals "compliance_review_failed"
const MessagingHostedNumberOrderStatusComplianceReviewFailed = shared.MessagingHostedNumberOrderStatusComplianceReviewFailed

// Equals "deleted"
const MessagingHostedNumberOrderStatusDeleted = shared.MessagingHostedNumberOrderStatusDeleted

// Equals "failed"
const MessagingHostedNumberOrderStatusFailed = shared.MessagingHostedNumberOrderStatusFailed

// Equals "incomplete_documentation"
const MessagingHostedNumberOrderStatusIncompleteDocumentation = shared.MessagingHostedNumberOrderStatusIncompleteDocumentation

// Equals "incorrect_billing_information"
const MessagingHostedNumberOrderStatusIncorrectBillingInformation = shared.MessagingHostedNumberOrderStatusIncorrectBillingInformation

// Equals "ineligible_carrier"
const MessagingHostedNumberOrderStatusIneligibleCarrier = shared.MessagingHostedNumberOrderStatusIneligibleCarrier

// Equals "loa_file_invalid"
const MessagingHostedNumberOrderStatusLoaFileInvalid = shared.MessagingHostedNumberOrderStatusLoaFileInvalid

// Equals "loa_file_successful"
const MessagingHostedNumberOrderStatusLoaFileSuccessful = shared.MessagingHostedNumberOrderStatusLoaFileSuccessful

// Equals "pending"
const MessagingHostedNumberOrderStatusPending = shared.MessagingHostedNumberOrderStatusPending

// Equals "provisioning"
const MessagingHostedNumberOrderStatusProvisioning = shared.MessagingHostedNumberOrderStatusProvisioning

// Equals "successful"
const MessagingHostedNumberOrderStatusSuccessful = shared.MessagingHostedNumberOrderStatusSuccessful

// This is an alias to an internal type.
type Metadata = shared.Metadata

// High level health metrics about the number and it's messaging sending patterns.
//
// This is an alias to an internal type.
type NumberHealthMetrics = shared.NumberHealthMetrics

// This is an alias to an internal type.
type PhoneNumberWithMessagingSettings = shared.PhoneNumberWithMessagingSettings

// This is an alias to an internal type.
type PhoneNumberWithMessagingSettingsFeatures = shared.PhoneNumberWithMessagingSettingsFeatures

// Identifies the type of the resource.
//
// This is an alias to an internal type.
type PhoneNumberWithMessagingSettingsRecordType = shared.PhoneNumberWithMessagingSettingsRecordType

// Equals "messaging_phone_number"
const PhoneNumberWithMessagingSettingsRecordTypeMessagingPhoneNumber = shared.PhoneNumberWithMessagingSettingsRecordTypeMessagingPhoneNumber

// Equals "messaging_settings"
const PhoneNumberWithMessagingSettingsRecordTypeMessagingSettings = shared.PhoneNumberWithMessagingSettingsRecordTypeMessagingSettings

// The type of the phone number
//
// This is an alias to an internal type.
type PhoneNumberWithMessagingSettingsType = shared.PhoneNumberWithMessagingSettingsType

// Equals "long-code"
const PhoneNumberWithMessagingSettingsTypeLongCode = shared.PhoneNumberWithMessagingSettingsTypeLongCode

// Equals "toll-free"
const PhoneNumberWithMessagingSettingsTypeTollFree = shared.PhoneNumberWithMessagingSettingsTypeTollFree

// Equals "short-code"
const PhoneNumberWithMessagingSettingsTypeShortCode = shared.PhoneNumberWithMessagingSettingsTypeShortCode

// Equals "longcode"
const PhoneNumberWithMessagingSettingsTypeLongcode = shared.PhoneNumberWithMessagingSettingsTypeLongcode

// Equals "tollfree"
const PhoneNumberWithMessagingSettingsTypeTollfree = shared.PhoneNumberWithMessagingSettingsTypeTollfree

// Equals "shortcode"
const PhoneNumberWithMessagingSettingsTypeShortcode = shared.PhoneNumberWithMessagingSettingsTypeShortcode

// Porting order status
//
// This is an alias to an internal type.
type PortingOrderStatus = shared.PortingOrderStatus

// The current status of the porting order
//
// This is an alias to an internal type.
type PortingOrderStatusValue = shared.PortingOrderStatusValue

// Equals "draft"
const PortingOrderStatusValueDraft = shared.PortingOrderStatusValueDraft

// Equals "in-process"
const PortingOrderStatusValueInProcess = shared.PortingOrderStatusValueInProcess

// Equals "submitted"
const PortingOrderStatusValueSubmitted = shared.PortingOrderStatusValueSubmitted

// Equals "exception"
const PortingOrderStatusValueException = shared.PortingOrderStatusValueException

// Equals "foc-date-confirmed"
const PortingOrderStatusValueFocDateConfirmed = shared.PortingOrderStatusValueFocDateConfirmed

// Equals "ported"
const PortingOrderStatusValuePorted = shared.PortingOrderStatusValuePorted

// Equals "cancelled"
const PortingOrderStatusValueCancelled = shared.PortingOrderStatusValueCancelled

// Equals "cancel-pending"
const PortingOrderStatusValueCancelPending = shared.PortingOrderStatusValueCancelPending

// This is an alias to an internal type.
type PortingOrdersExceptionType = shared.PortingOrdersExceptionType

// Identifier of an exception type
//
// This is an alias to an internal type.
type PortingOrdersExceptionTypeCode = shared.PortingOrdersExceptionTypeCode

// Equals "ACCOUNT_NUMBER_MISMATCH"
const PortingOrdersExceptionTypeCodeAccountNumberMismatch = shared.PortingOrdersExceptionTypeCodeAccountNumberMismatch

// Equals "AUTH_PERSON_MISMATCH"
const PortingOrdersExceptionTypeCodeAuthPersonMismatch = shared.PortingOrdersExceptionTypeCodeAuthPersonMismatch

// Equals "BTN_ATN_MISMATCH"
const PortingOrdersExceptionTypeCodeBtnAtnMismatch = shared.PortingOrdersExceptionTypeCodeBtnAtnMismatch

// Equals "ENTITY_NAME_MISMATCH"
const PortingOrdersExceptionTypeCodeEntityNameMismatch = shared.PortingOrdersExceptionTypeCodeEntityNameMismatch

// Equals "FOC_EXPIRED"
const PortingOrdersExceptionTypeCodeFocExpired = shared.PortingOrdersExceptionTypeCodeFocExpired

// Equals "FOC_REJECTED"
const PortingOrdersExceptionTypeCodeFocRejected = shared.PortingOrdersExceptionTypeCodeFocRejected

// Equals "LOCATION_MISMATCH"
const PortingOrdersExceptionTypeCodeLocationMismatch = shared.PortingOrdersExceptionTypeCodeLocationMismatch

// Equals "LSR_PENDING"
const PortingOrdersExceptionTypeCodeLsrPending = shared.PortingOrdersExceptionTypeCodeLsrPending

// Equals "MAIN_BTN_PORTING"
const PortingOrdersExceptionTypeCodeMainBtnPorting = shared.PortingOrdersExceptionTypeCodeMainBtnPorting

// Equals "OSP_IRRESPONSIVE"
const PortingOrdersExceptionTypeCodeOspIrresponsive = shared.PortingOrdersExceptionTypeCodeOspIrresponsive

// Equals "OTHER"
const PortingOrdersExceptionTypeCodeOther = shared.PortingOrdersExceptionTypeCodeOther

// Equals "PASSCODE_PIN_INVALID"
const PortingOrdersExceptionTypeCodePasscodePinInvalid = shared.PortingOrdersExceptionTypeCodePasscodePinInvalid

// Equals "PHONE_NUMBER_HAS_SPECIAL_FEATURE"
const PortingOrdersExceptionTypeCodePhoneNumberHasSpecialFeature = shared.PortingOrdersExceptionTypeCodePhoneNumberHasSpecialFeature

// Equals "PHONE_NUMBER_MISMATCH"
const PortingOrdersExceptionTypeCodePhoneNumberMismatch = shared.PortingOrdersExceptionTypeCodePhoneNumberMismatch

// Equals "PHONE_NUMBER_NOT_PORTABLE"
const PortingOrdersExceptionTypeCodePhoneNumberNotPortable = shared.PortingOrdersExceptionTypeCodePhoneNumberNotPortable

// Equals "PORT_TYPE_INCORRECT"
const PortingOrdersExceptionTypeCodePortTypeIncorrect = shared.PortingOrdersExceptionTypeCodePortTypeIncorrect

// Equals "PORTING_ORDER_SPLIT_REQUIRED"
const PortingOrdersExceptionTypeCodePortingOrderSplitRequired = shared.PortingOrdersExceptionTypeCodePortingOrderSplitRequired

// Equals "POSTAL_CODE_MISMATCH"
const PortingOrdersExceptionTypeCodePostalCodeMismatch = shared.PortingOrdersExceptionTypeCodePostalCodeMismatch

// Equals "RATE_CENTER_NOT_PORTABLE"
const PortingOrdersExceptionTypeCodeRateCenterNotPortable = shared.PortingOrdersExceptionTypeCodeRateCenterNotPortable

// Equals "SV_CONFLICT"
const PortingOrdersExceptionTypeCodeSvConflict = shared.PortingOrdersExceptionTypeCodeSvConflict

// Equals "SV_UNKNOWN_FAILURE"
const PortingOrdersExceptionTypeCodeSvUnknownFailure = shared.PortingOrdersExceptionTypeCodeSvUnknownFailure

// This is an alias to an internal type.
type RoomParticipant = shared.RoomParticipant

// This is an alias to an internal type.
type ShortCode = shared.ShortCode

// Identifies the type of the resource.
//
// This is an alias to an internal type.
type ShortCodeRecordType = shared.ShortCodeRecordType

// Equals "short_code"
const ShortCodeRecordTypeShortCode = shared.ShortCodeRecordTypeShortCode

// This is an alias to an internal type.
type SimCardStatus = shared.SimCardStatus

// The current status of the SIM card. It will be one of the following: <br/>
//
// <ul>
//
//	<li><code>registering</code> - the card is being registered</li>
//	<li><code>enabling</code> - the card is being enabled</li>
//	<li><code>enabled</code> - the card is enabled and ready for use</li>
//	<li><code>disabling</code> - the card is being disabled</li>
//	<li><code>disabled</code> - the card has been disabled and cannot be used</li>
//	<li><code>data_limit_exceeded</code> - the card has exceeded its data consumption limit</li>
//	<li><code>setting_standby</code> - the process to set the card in stand by is in progress</li>
//	<li><code>standby</code> - the card is in stand by</li>
//
// </ul>
// Transitioning between the enabled and disabled states may take a period of time.
//
// This is an alias to an internal type.
type SimCardStatusValue = shared.SimCardStatusValue

// Equals "registering"
const SimCardStatusValueRegistering = shared.SimCardStatusValueRegistering

// Equals "enabling"
const SimCardStatusValueEnabling = shared.SimCardStatusValueEnabling

// Equals "enabled"
const SimCardStatusValueEnabled = shared.SimCardStatusValueEnabled

// Equals "disabling"
const SimCardStatusValueDisabling = shared.SimCardStatusValueDisabling

// Equals "disabled"
const SimCardStatusValueDisabled = shared.SimCardStatusValueDisabled

// Equals "data_limit_exceeded"
const SimCardStatusValueDataLimitExceeded = shared.SimCardStatusValueDataLimitExceeded

// Equals "setting_standby"
const SimCardStatusValueSettingStandby = shared.SimCardStatusValueSettingStandby

// Equals "standby"
const SimCardStatusValueStandby = shared.SimCardStatusValueStandby

// This is an alias to an internal type.
type SimCardStatusParam = shared.SimCardStatusParam

// This is an alias to an internal type.
type SimpleSimCard = shared.SimpleSimCard

// The SIM card consumption so far in the current billing cycle.
//
// This is an alias to an internal type.
type SimpleSimCardCurrentBillingPeriodConsumedData = shared.SimpleSimCardCurrentBillingPeriodConsumedData

// The SIM card individual data limit configuration.
//
// This is an alias to an internal type.
type SimpleSimCardDataLimit = shared.SimpleSimCardDataLimit

// The type of SIM card
//
// This is an alias to an internal type.
type SimpleSimCardType = shared.SimpleSimCardType

// Equals "physical"
const SimpleSimCardTypePhysical = shared.SimpleSimCardTypePhysical

// Equals "esim"
const SimpleSimCardTypeEsim = shared.SimpleSimCardTypeEsim

// This is an alias to an internal type.
type SubNumberOrderRegulatoryRequirementWithValue = shared.SubNumberOrderRegulatoryRequirementWithValue

// This is an alias to an internal type.
type SubNumberOrderRegulatoryRequirementWithValueFieldType = shared.SubNumberOrderRegulatoryRequirementWithValueFieldType

// Equals "textual"
const SubNumberOrderRegulatoryRequirementWithValueFieldTypeTextual = shared.SubNumberOrderRegulatoryRequirementWithValueFieldTypeTextual

// Equals "datetime"
const SubNumberOrderRegulatoryRequirementWithValueFieldTypeDatetime = shared.SubNumberOrderRegulatoryRequirementWithValueFieldTypeDatetime

// Equals "address"
const SubNumberOrderRegulatoryRequirementWithValueFieldTypeAddress = shared.SubNumberOrderRegulatoryRequirementWithValueFieldTypeAddress

// Equals "document"
const SubNumberOrderRegulatoryRequirementWithValueFieldTypeDocument = shared.SubNumberOrderRegulatoryRequirementWithValueFieldTypeDocument

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v3/option"
)

// LegacyReportingBatchDetailRecordService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLegacyReportingBatchDetailRecordService] method instead.
type LegacyReportingBatchDetailRecordService struct {
	Options      []option.RequestOption
	Messaging    LegacyReportingBatchDetailRecordMessagingService
	SpeechToText LegacyReportingBatchDetailRecordSpeechToTextService
	Voice        LegacyReportingBatchDetailRecordVoiceService
}

// NewLegacyReportingBatchDetailRecordService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLegacyReportingBatchDetailRecordService(opts ...option.RequestOption) (r LegacyReportingBatchDetailRecordService) {
	r = LegacyReportingBatchDetailRecordService{}
	r.Options = opts
	r.Messaging = NewLegacyReportingBatchDetailRecordMessagingService(opts...)
	r.SpeechToText = NewLegacyReportingBatchDetailRecordSpeechToTextService(opts...)
	r.Voice = NewLegacyReportingBatchDetailRecordVoiceService(opts...)
	return
}

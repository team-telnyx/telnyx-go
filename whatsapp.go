// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// WhatsappService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappService] method instead.
type WhatsappService struct {
	Options []option.RequestOption
	// Manage Whatsapp business accounts
	BusinessAccounts WhatsappBusinessAccountService
	// Manage Whatsapp message templates
	MessageTemplates WhatsappMessageTemplateService
	// Manage Whatsapp phone numbers
	PhoneNumbers WhatsappPhoneNumberService
}

// NewWhatsappService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWhatsappService(opts ...option.RequestOption) (r WhatsappService) {
	r = WhatsappService{}
	r.Options = opts
	r.BusinessAccounts = NewWhatsappBusinessAccountService(opts...)
	r.MessageTemplates = NewWhatsappMessageTemplateService(opts...)
	r.PhoneNumbers = NewWhatsappPhoneNumberService(opts...)
	return
}

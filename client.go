// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the telnyx API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options      []option.RequestOption
	Legacy       LegacyService
	OAuth        OAuthService
	OAuthClients OAuthClientService
	OAuthGrants  OAuthGrantService
	Webhooks     WebhookService
	// IP Address Operations
	AccessIPAddress AccessIPAddressService
	// IP Range Operations
	AccessIPRanges AccessIPRangeService
	Actions        ActionService
	// Operations to work with Address records. Address records are emergency-validated
	// addresses meant to be associated with phone numbers. They are validated for
	// emergency usage purposes at creation time, although you may validate them
	// separately with a custom workflow using the ValidateAddress operation
	// separately. Address records are not usable for physical orders, such as for
	// Telnyx SIM cards, please use UserAddress for that. It is not possible to
	// entirely skip emergency service validation for Address records; if an emergency
	// provider for a phone number rejects the address then it cannot be used on a
	// phone number. To prevent records from getting out of sync, Address records are
	// immutable and cannot be altered once created. If you realize you need to alter
	// an address, a new record must be created with the differing address.
	Addresses      AddressService
	AdvancedOrders AdvancedOrderService
	// Generate text with LLMs
	AI AIService
	// Audit log operations.
	AuditEvents             AuditEventService
	AuthenticationProviders AuthenticationProviderService
	// Number search
	AvailablePhoneNumberBlocks AvailablePhoneNumberBlockService
	// Number search
	AvailablePhoneNumbers AvailablePhoneNumberService
	// Billing operations
	Balance BalanceService
	// Billing groups operations
	BillingGroups BillingGroupService
	// View SIM card actions, their progress and timestamps using the SIM Card Actions
	// API
	BulkSimCardActions BulkSimCardActionService
	BundlePricing      BundlePricingService
	// Call Control applications operations
	CallControlApplications CallControlApplicationService
	// Call Control debugging
	CallEvents CallEventService
	Calls      CallService
	// Voice Channels
	ChannelZones     ChannelZoneService
	ChargesBreakdown ChargesBreakdownService
	ChargesSummary   ChargesSummaryService
	// Number orders
	Comments CommentService
	// Conference command operations
	Conferences ConferenceService
	Connections ConnectionService
	// Country Coverage
	CountryCoverage CountryCoverageService
	// Credential connection operations
	CredentialConnections CredentialConnectionService
	// Call Recordings operations.
	CustomStorageCredentials CustomStorageCredentialService
	// Customer Service Record operations
	CustomerServiceRecords CustomerServiceRecordService
	// Detail Records operations
	DetailRecords DetailRecordService
	// Dialogflow Connection Operations.
	DialogflowConnections DialogflowConnectionService
	// Documents
	DocumentLinks DocumentLinkService
	// Documents
	Documents DocumentService
	// Dynamic emergency address operations
	DynamicEmergencyAddresses DynamicEmergencyAddressService
	// Dynamic Emergency Endpoints
	DynamicEmergencyEndpoints DynamicEmergencyEndpointService
	// External Connections operations
	ExternalConnections ExternalConnectionService
	// Fax Applications operations
	FaxApplications FaxApplicationService
	// Programmable fax command operations
	Faxes FaxService
	// FQDN connection operations
	FqdnConnections FqdnConnectionService
	// FQDN operations
	Fqdns FqdnService
	// Global IPs
	GlobalIPAllowedPorts GlobalIPAllowedPortService
	// Global IPs
	GlobalIPAssignmentHealth GlobalIPAssignmentHealthService
	// Global IPs
	GlobalIPAssignments GlobalIPAssignmentService
	// Global IPs
	GlobalIPAssignmentsUsage GlobalIPAssignmentsUsageService
	// Global IPs
	GlobalIPHealthCheckTypes GlobalIPHealthCheckTypeService
	// Global IPs
	GlobalIPHealthChecks GlobalIPHealthCheckService
	// Global IPs
	GlobalIPLatency GlobalIPLatencyService
	// Global IPs
	GlobalIPProtocols GlobalIPProtocolService
	// Global IPs
	GlobalIPUsage GlobalIPUsageService
	// Global IPs
	GlobalIPs GlobalIPService
	// Voice Channels
	InboundChannels InboundChannelService
	// Store and retrieve integration secrets
	IntegrationSecrets IntegrationSecretService
	// Inventory Level
	InventoryCoverage InventoryCoverageService
	Invoices          InvoiceService
	// IP connection operations
	IPConnections IPConnectionService
	// IP operations
	IPs IPService
	// Ledger billing reports
	LedgerBillingGroupReports LedgerBillingGroupReportService
	// Voice Channels
	List ListService
	// Managed Accounts operations
	ManagedAccounts ManagedAccountService
	// Media Storage operations
	Media     MediaService
	Messages  MessageService
	Messaging MessagingService
	// Manage your messaging hosted numbers
	MessagingHostedNumberOrders MessagingHostedNumberOrderService
	MessagingHostedNumbers      MessagingHostedNumberService
	// Configure your phone numbers
	MessagingNumbersBulkUpdates MessagingNumbersBulkUpdateService
	// Opt-Out Management
	MessagingOptouts  MessagingOptoutService
	MessagingProfiles MessagingProfileService
	MessagingTollfree MessagingTollfreeService
	// Messaging URL Domains
	MessagingURLDomains MessagingURLDomainService
	// Mobile network operators operations
	MobileNetworkOperators MobileNetworkOperatorService
	// Mobile push credential management
	MobilePushCredentials MobilePushCredentialService
	NetworkCoverage       NetworkCoverageService
	// Network operations
	Networks NetworkService
	// Notification settings operations
	NotificationChannels NotificationChannelService
	// Notification settings operations
	NotificationEventConditions NotificationEventConditionService
	// Notification settings operations
	NotificationEvents NotificationEventService
	// Notification settings operations
	NotificationProfiles NotificationProfileService
	// Notification settings operations
	NotificationSettings NotificationSettingService
	NumberBlockOrders    NumberBlockOrderService
	// Look up phone number data
	NumberLookup            NumberLookupService
	NumberOrderPhoneNumbers NumberOrderPhoneNumberService
	// Number orders
	NumberOrders NumberOrderService
	// Number reservations
	NumberReservations NumberReservationService
	NumbersFeatures    NumbersFeatureService
	OperatorConnect    OperatorConnectService
	// OTA updates operations
	OtaUpdates OtaUpdateService
	// Outbound voice profiles operations
	OutboundVoiceProfiles OutboundVoiceProfileService
	// Operations for managing stored payment transactions.
	Payment           PaymentService
	PhoneNumberBlocks PhoneNumberBlockService
	// Configure your phone numbers
	PhoneNumbers PhoneNumberService
	// Regulatory Requirements
	PhoneNumbersRegulatoryRequirements PhoneNumbersRegulatoryRequirementService
	// Determining portability of phone numbers
	PortabilityChecks PortabilityCheckService
	// Endpoints related to porting orders management.
	Porting PortingService
	// Endpoints related to porting orders management.
	PortingOrders PortingOrderService
	// Endpoints related to porting orders management.
	PortingPhoneNumbers PortingPhoneNumberService
	// Number portout operations
	Portouts PortoutService
	// Private Wireless Gateways operations
	PrivateWirelessGateways PrivateWirelessGatewayService
	// Public Internet Gateway operations
	PublicInternetGateways PublicInternetGatewayService
	// Queue commands operations
	Queues    QueueService
	RcsAgents RcsAgentService
	// Call Recordings operations.
	RecordingTranscriptions RecordingTranscriptionService
	// Call Recordings operations.
	Recordings RecordingService
	// Regions
	Regions RegionService
	// Regulatory Requirements
	RegulatoryRequirements RegulatoryRequirementService
	Reports                ReportService
	// Requirement Groups
	RequirementGroups RequirementGroupService
	// Types of requirements for international numbers and porting orders
	RequirementTypes RequirementTypeService
	// Requirements for international numbers and porting orders
	Requirements RequirementService
	// Rooms Compositions operations.
	RoomCompositions RoomCompositionService
	// Rooms Participants operations.
	RoomParticipants RoomParticipantService
	// Rooms Recordings operations.
	RoomRecordings RoomRecordingService
	// Rooms operations.
	Rooms RoomService
	// Observability into Telnyx platform stability and performance.
	Seti SetiService
	// Short codes
	ShortCodes ShortCodeService
	// SIM Cards operations
	SimCardDataUsageNotifications SimCardDataUsageNotificationService
	// SIM Card Groups operations
	SimCardGroups SimCardGroupService
	// SIM Card Orders operations
	SimCardOrderPreview SimCardOrderPreviewService
	// SIM Card Orders operations
	SimCardOrders SimCardOrderService
	// SIM Cards operations
	SimCards SimCardService
	// SIPREC connectors configuration.
	SiprecConnectors SiprecConnectorService
	// Migrate data from an external provider into Telnyx Cloud Storage
	Storage         StorageService
	SubNumberOrders SubNumberOrderService
	// Number orders
	SubNumberOrdersReport SubNumberOrdersReportService
	TelephonyCredentials  TelephonyCredentialService
	// TeXML REST Commands
	Texml TexmlService
	// TeXML Applications operations
	TexmlApplications TexmlApplicationService
	// Text to speech streaming command operations
	TextToSpeech TextToSpeechService
	// Usage data reporting across Telnyx products
	UsageReports UsageReportService
	// Operations for working with UserAddress records. UserAddress records are stored
	// addresses that users can use for non-emergency-calling purposes, such as for
	// shipping addresses for orders of wireless SIMs (or other physical items). They
	// cannot be used for emergency calling and are distinct from Address records,
	// which are used on phone numbers.
	UserAddresses UserAddressService
	// User-defined tags for Telnyx resources
	UserTags UserTagService
	// Two factor authentication API
	Verifications VerificationService
	// Verified Numbers operations
	VerifiedNumbers VerifiedNumberService
	// Two factor authentication API
	VerifyProfiles VerifyProfileService
	// Virtual Cross Connect operations
	VirtualCrossConnects VirtualCrossConnectService
	// Virtual Cross Connect operations
	VirtualCrossConnectsCoverage VirtualCrossConnectsCoverageService
	// Webhooks operations
	WebhookDeliveries WebhookDeliveryService
	// WireGuard Interface operations
	WireguardInterfaces WireguardInterfaceService
	// WireGuard Interface operations
	WireguardPeers WireguardPeerService
	// Regions for wireless services
	Wireless WirelessService
	// Wireless Blocklists operations
	WirelessBlocklistValues WirelessBlocklistValueService
	// Wireless Blocklists operations
	WirelessBlocklists WirelessBlocklistService
	WellKnown          WellKnownService
	// Inexplicit number orders for bulk purchasing without specifying exact numbers
	InexplicitNumberOrders InexplicitNumberOrderService
	// Mobile phone number operations
	MobilePhoneNumbers MobilePhoneNumberService
	// Mobile voice connection operations
	MobileVoiceConnections MobileVoiceConnectionService
	Messaging10dlc         Messaging10dlcService
	// Speech to text command operations
	SpeechToText            SpeechToTextService
	Organizations           OrganizationService
	AlphanumericSenderIDs   AlphanumericSenderIDService
	MessagingProfileMetrics MessagingProfileMetricService
}

// DefaultClientOptions read from the environment (TELNYX_API_KEY,
// TELNYX_PUBLIC_KEY, TELNYX_CLIENT_ID, TELNYX_CLIENT_SECRET, TELNYX_BASE_URL).
// This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("TELNYX_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("TELNYX_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("TELNYX_PUBLIC_KEY"); ok {
		defaults = append(defaults, option.WithPublicKey(o))
	}
	if o, ok := os.LookupEnv("TELNYX_CLIENT_ID"); ok {
		defaults = append(defaults, option.WithClientID(o))
	}
	if o, ok := os.LookupEnv("TELNYX_CLIENT_SECRET"); ok {
		defaults = append(defaults, option.WithClientSecret(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (TELNYX_API_KEY, TELNYX_PUBLIC_KEY, TELNYX_CLIENT_ID,
// TELNYX_CLIENT_SECRET, TELNYX_BASE_URL). The option passed in as arguments are
// applied after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.Legacy = NewLegacyService(opts...)
	r.OAuth = NewOAuthService(opts...)
	r.OAuthClients = NewOAuthClientService(opts...)
	r.OAuthGrants = NewOAuthGrantService(opts...)
	r.Webhooks = NewWebhookService(opts...)
	r.AccessIPAddress = NewAccessIPAddressService(opts...)
	r.AccessIPRanges = NewAccessIPRangeService(opts...)
	r.Actions = NewActionService(opts...)
	r.Addresses = NewAddressService(opts...)
	r.AdvancedOrders = NewAdvancedOrderService(opts...)
	r.AI = NewAIService(opts...)
	r.AuditEvents = NewAuditEventService(opts...)
	r.AuthenticationProviders = NewAuthenticationProviderService(opts...)
	r.AvailablePhoneNumberBlocks = NewAvailablePhoneNumberBlockService(opts...)
	r.AvailablePhoneNumbers = NewAvailablePhoneNumberService(opts...)
	r.Balance = NewBalanceService(opts...)
	r.BillingGroups = NewBillingGroupService(opts...)
	r.BulkSimCardActions = NewBulkSimCardActionService(opts...)
	r.BundlePricing = NewBundlePricingService(opts...)
	r.CallControlApplications = NewCallControlApplicationService(opts...)
	r.CallEvents = NewCallEventService(opts...)
	r.Calls = NewCallService(opts...)
	r.ChannelZones = NewChannelZoneService(opts...)
	r.ChargesBreakdown = NewChargesBreakdownService(opts...)
	r.ChargesSummary = NewChargesSummaryService(opts...)
	r.Comments = NewCommentService(opts...)
	r.Conferences = NewConferenceService(opts...)
	r.Connections = NewConnectionService(opts...)
	r.CountryCoverage = NewCountryCoverageService(opts...)
	r.CredentialConnections = NewCredentialConnectionService(opts...)
	r.CustomStorageCredentials = NewCustomStorageCredentialService(opts...)
	r.CustomerServiceRecords = NewCustomerServiceRecordService(opts...)
	r.DetailRecords = NewDetailRecordService(opts...)
	r.DialogflowConnections = NewDialogflowConnectionService(opts...)
	r.DocumentLinks = NewDocumentLinkService(opts...)
	r.Documents = NewDocumentService(opts...)
	r.DynamicEmergencyAddresses = NewDynamicEmergencyAddressService(opts...)
	r.DynamicEmergencyEndpoints = NewDynamicEmergencyEndpointService(opts...)
	r.ExternalConnections = NewExternalConnectionService(opts...)
	r.FaxApplications = NewFaxApplicationService(opts...)
	r.Faxes = NewFaxService(opts...)
	r.FqdnConnections = NewFqdnConnectionService(opts...)
	r.Fqdns = NewFqdnService(opts...)
	r.GlobalIPAllowedPorts = NewGlobalIPAllowedPortService(opts...)
	r.GlobalIPAssignmentHealth = NewGlobalIPAssignmentHealthService(opts...)
	r.GlobalIPAssignments = NewGlobalIPAssignmentService(opts...)
	r.GlobalIPAssignmentsUsage = NewGlobalIPAssignmentsUsageService(opts...)
	r.GlobalIPHealthCheckTypes = NewGlobalIPHealthCheckTypeService(opts...)
	r.GlobalIPHealthChecks = NewGlobalIPHealthCheckService(opts...)
	r.GlobalIPLatency = NewGlobalIPLatencyService(opts...)
	r.GlobalIPProtocols = NewGlobalIPProtocolService(opts...)
	r.GlobalIPUsage = NewGlobalIPUsageService(opts...)
	r.GlobalIPs = NewGlobalIPService(opts...)
	r.InboundChannels = NewInboundChannelService(opts...)
	r.IntegrationSecrets = NewIntegrationSecretService(opts...)
	r.InventoryCoverage = NewInventoryCoverageService(opts...)
	r.Invoices = NewInvoiceService(opts...)
	r.IPConnections = NewIPConnectionService(opts...)
	r.IPs = NewIPService(opts...)
	r.LedgerBillingGroupReports = NewLedgerBillingGroupReportService(opts...)
	r.List = NewListService(opts...)
	r.ManagedAccounts = NewManagedAccountService(opts...)
	r.Media = NewMediaService(opts...)
	r.Messages = NewMessageService(opts...)
	r.Messaging = NewMessagingService(opts...)
	r.MessagingHostedNumberOrders = NewMessagingHostedNumberOrderService(opts...)
	r.MessagingHostedNumbers = NewMessagingHostedNumberService(opts...)
	r.MessagingNumbersBulkUpdates = NewMessagingNumbersBulkUpdateService(opts...)
	r.MessagingOptouts = NewMessagingOptoutService(opts...)
	r.MessagingProfiles = NewMessagingProfileService(opts...)
	r.MessagingTollfree = NewMessagingTollfreeService(opts...)
	r.MessagingURLDomains = NewMessagingURLDomainService(opts...)
	r.MobileNetworkOperators = NewMobileNetworkOperatorService(opts...)
	r.MobilePushCredentials = NewMobilePushCredentialService(opts...)
	r.NetworkCoverage = NewNetworkCoverageService(opts...)
	r.Networks = NewNetworkService(opts...)
	r.NotificationChannels = NewNotificationChannelService(opts...)
	r.NotificationEventConditions = NewNotificationEventConditionService(opts...)
	r.NotificationEvents = NewNotificationEventService(opts...)
	r.NotificationProfiles = NewNotificationProfileService(opts...)
	r.NotificationSettings = NewNotificationSettingService(opts...)
	r.NumberBlockOrders = NewNumberBlockOrderService(opts...)
	r.NumberLookup = NewNumberLookupService(opts...)
	r.NumberOrderPhoneNumbers = NewNumberOrderPhoneNumberService(opts...)
	r.NumberOrders = NewNumberOrderService(opts...)
	r.NumberReservations = NewNumberReservationService(opts...)
	r.NumbersFeatures = NewNumbersFeatureService(opts...)
	r.OperatorConnect = NewOperatorConnectService(opts...)
	r.OtaUpdates = NewOtaUpdateService(opts...)
	r.OutboundVoiceProfiles = NewOutboundVoiceProfileService(opts...)
	r.Payment = NewPaymentService(opts...)
	r.PhoneNumberBlocks = NewPhoneNumberBlockService(opts...)
	r.PhoneNumbers = NewPhoneNumberService(opts...)
	r.PhoneNumbersRegulatoryRequirements = NewPhoneNumbersRegulatoryRequirementService(opts...)
	r.PortabilityChecks = NewPortabilityCheckService(opts...)
	r.Porting = NewPortingService(opts...)
	r.PortingOrders = NewPortingOrderService(opts...)
	r.PortingPhoneNumbers = NewPortingPhoneNumberService(opts...)
	r.Portouts = NewPortoutService(opts...)
	r.PrivateWirelessGateways = NewPrivateWirelessGatewayService(opts...)
	r.PublicInternetGateways = NewPublicInternetGatewayService(opts...)
	r.Queues = NewQueueService(opts...)
	r.RcsAgents = NewRcsAgentService(opts...)
	r.RecordingTranscriptions = NewRecordingTranscriptionService(opts...)
	r.Recordings = NewRecordingService(opts...)
	r.Regions = NewRegionService(opts...)
	r.RegulatoryRequirements = NewRegulatoryRequirementService(opts...)
	r.Reports = NewReportService(opts...)
	r.RequirementGroups = NewRequirementGroupService(opts...)
	r.RequirementTypes = NewRequirementTypeService(opts...)
	r.Requirements = NewRequirementService(opts...)
	r.RoomCompositions = NewRoomCompositionService(opts...)
	r.RoomParticipants = NewRoomParticipantService(opts...)
	r.RoomRecordings = NewRoomRecordingService(opts...)
	r.Rooms = NewRoomService(opts...)
	r.Seti = NewSetiService(opts...)
	r.ShortCodes = NewShortCodeService(opts...)
	r.SimCardDataUsageNotifications = NewSimCardDataUsageNotificationService(opts...)
	r.SimCardGroups = NewSimCardGroupService(opts...)
	r.SimCardOrderPreview = NewSimCardOrderPreviewService(opts...)
	r.SimCardOrders = NewSimCardOrderService(opts...)
	r.SimCards = NewSimCardService(opts...)
	r.SiprecConnectors = NewSiprecConnectorService(opts...)
	r.Storage = NewStorageService(opts...)
	r.SubNumberOrders = NewSubNumberOrderService(opts...)
	r.SubNumberOrdersReport = NewSubNumberOrdersReportService(opts...)
	r.TelephonyCredentials = NewTelephonyCredentialService(opts...)
	r.Texml = NewTexmlService(opts...)
	r.TexmlApplications = NewTexmlApplicationService(opts...)
	r.TextToSpeech = NewTextToSpeechService(opts...)
	r.UsageReports = NewUsageReportService(opts...)
	r.UserAddresses = NewUserAddressService(opts...)
	r.UserTags = NewUserTagService(opts...)
	r.Verifications = NewVerificationService(opts...)
	r.VerifiedNumbers = NewVerifiedNumberService(opts...)
	r.VerifyProfiles = NewVerifyProfileService(opts...)
	r.VirtualCrossConnects = NewVirtualCrossConnectService(opts...)
	r.VirtualCrossConnectsCoverage = NewVirtualCrossConnectsCoverageService(opts...)
	r.WebhookDeliveries = NewWebhookDeliveryService(opts...)
	r.WireguardInterfaces = NewWireguardInterfaceService(opts...)
	r.WireguardPeers = NewWireguardPeerService(opts...)
	r.Wireless = NewWirelessService(opts...)
	r.WirelessBlocklistValues = NewWirelessBlocklistValueService(opts...)
	r.WirelessBlocklists = NewWirelessBlocklistService(opts...)
	r.WellKnown = NewWellKnownService(opts...)
	r.InexplicitNumberOrders = NewInexplicitNumberOrderService(opts...)
	r.MobilePhoneNumbers = NewMobilePhoneNumberService(opts...)
	r.MobileVoiceConnections = NewMobileVoiceConnectionService(opts...)
	r.Messaging10dlc = NewMessaging10dlcService(opts...)
	r.SpeechToText = NewSpeechToTextService(opts...)
	r.Organizations = NewOrganizationService(opts...)
	r.AlphanumericSenderIDs = NewAlphanumericSenderIDService(opts...)
	r.MessagingProfileMetrics = NewMessagingProfileMetricService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}

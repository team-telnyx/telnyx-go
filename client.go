// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the telnyx API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                            []option.RequestOption
	Legacy                             LegacyService
	OAuth                              OAuthService
	OAuthClients                       OAuthClientService
	OAuthGrants                        OAuthGrantService
	Webhooks                           WebhookService
	AccessIPAddress                    AccessIPAddressService
	AccessIPRanges                     AccessIPRangeService
	Actions                            ActionService
	Addresses                          AddressService
	AdvancedOrders                     AdvancedOrderService
	AI                                 AIService
	AuditEvents                        AuditEventService
	AuthenticationProviders            AuthenticationProviderService
	AvailablePhoneNumberBlocks         AvailablePhoneNumberBlockService
	AvailablePhoneNumbers              AvailablePhoneNumberService
	Balance                            BalanceService
	BillingGroups                      BillingGroupService
	Brand                              BrandService
	BulkSimCardActions                 BulkSimCardActionService
	BundlePricing                      BundlePricingService
	CallControlApplications            CallControlApplicationService
	CallEvents                         CallEventService
	Calls                              CallService
	Campaign                           CampaignService
	CampaignBuilder                    CampaignBuilderService
	ChannelZones                       ChannelZoneService
	ChargesBreakdown                   ChargesBreakdownService
	ChargesSummary                     ChargesSummaryService
	Comments                           CommentService
	Conferences                        ConferenceService
	Connections                        ConnectionService
	CountryCoverage                    CountryCoverageService
	CredentialConnections              CredentialConnectionService
	CustomStorageCredentials           CustomStorageCredentialService
	CustomerServiceRecords             CustomerServiceRecordService
	DetailRecords                      DetailRecordService
	DialogflowConnections              DialogflowConnectionService
	DocumentLinks                      DocumentLinkService
	Documents                          DocumentService
	DynamicEmergencyAddresses          DynamicEmergencyAddressService
	DynamicEmergencyEndpoints          DynamicEmergencyEndpointService
	Enumeration                        EnumerationService
	ExternalConnections                ExternalConnectionService
	FaxApplications                    FaxApplicationService
	Faxes                              FaxService
	FqdnConnections                    FqdnConnectionService
	Fqdns                              FqdnService
	GlobalIPAllowedPorts               GlobalIPAllowedPortService
	GlobalIPAssignmentHealth           GlobalIPAssignmentHealthService
	GlobalIPAssignments                GlobalIPAssignmentService
	GlobalIPAssignmentsUsage           GlobalIPAssignmentsUsageService
	GlobalIPHealthCheckTypes           GlobalIPHealthCheckTypeService
	GlobalIPHealthChecks               GlobalIPHealthCheckService
	GlobalIPLatency                    GlobalIPLatencyService
	GlobalIPProtocols                  GlobalIPProtocolService
	GlobalIPUsage                      GlobalIPUsageService
	GlobalIPs                          GlobalIPService
	InboundChannels                    InboundChannelService
	IntegrationSecrets                 IntegrationSecretService
	InventoryCoverage                  InventoryCoverageService
	Invoices                           InvoiceService
	IPConnections                      IPConnectionService
	IPs                                IPService
	LedgerBillingGroupReports          LedgerBillingGroupReportService
	List                               ListService
	ManagedAccounts                    ManagedAccountService
	Media                              MediaService
	Messages                           MessageService
	Messaging                          MessagingService
	MessagingHostedNumberOrders        MessagingHostedNumberOrderService
	MessagingHostedNumbers             MessagingHostedNumberService
	MessagingNumbersBulkUpdates        MessagingNumbersBulkUpdateService
	MessagingOptouts                   MessagingOptoutService
	MessagingProfiles                  MessagingProfileService
	MessagingTollfree                  MessagingTollfreeService
	MessagingURLDomains                MessagingURLDomainService
	Messsages                          MesssageService
	MobileNetworkOperators             MobileNetworkOperatorService
	MobilePushCredentials              MobilePushCredentialService
	NetworkCoverage                    NetworkCoverageService
	Networks                           NetworkService
	NotificationChannels               NotificationChannelService
	NotificationEventConditions        NotificationEventConditionService
	NotificationEvents                 NotificationEventService
	NotificationProfiles               NotificationProfileService
	NotificationSettings               NotificationSettingService
	NumberBlockOrders                  NumberBlockOrderService
	NumberLookup                       NumberLookupService
	NumberOrderPhoneNumbers            NumberOrderPhoneNumberService
	NumberOrders                       NumberOrderService
	NumberReservations                 NumberReservationService
	NumbersFeatures                    NumbersFeatureService
	OperatorConnect                    OperatorConnectService
	OtaUpdates                         OtaUpdateService
	OutboundVoiceProfiles              OutboundVoiceProfileService
	Payment                            PaymentService
	PhoneNumberAssignmentByProfile     PhoneNumberAssignmentByProfileService
	PhoneNumberBlocks                  PhoneNumberBlockService
	PhoneNumberCampaigns               PhoneNumberCampaignService
	PhoneNumbers                       PhoneNumberService
	PhoneNumbersRegulatoryRequirements PhoneNumbersRegulatoryRequirementService
	PortabilityChecks                  PortabilityCheckService
	Porting                            PortingService
	PortingOrders                      PortingOrderService
	PortingPhoneNumbers                PortingPhoneNumberService
	Portouts                           PortoutService
	PrivateWirelessGateways            PrivateWirelessGatewayService
	PublicInternetGateways             PublicInternetGatewayService
	Queues                             QueueService
	RcsAgents                          RcsAgentService
	RecordingTranscriptions            RecordingTranscriptionService
	Recordings                         RecordingService
	Regions                            RegionService
	RegulatoryRequirements             RegulatoryRequirementService
	Reports                            ReportService
	RequirementGroups                  RequirementGroupService
	RequirementTypes                   RequirementTypeService
	Requirements                       RequirementService
	RoomCompositions                   RoomCompositionService
	RoomParticipants                   RoomParticipantService
	RoomRecordings                     RoomRecordingService
	Rooms                              RoomService
	Seti                               SetiService
	ShortCodes                         ShortCodeService
	SimCardDataUsageNotifications      SimCardDataUsageNotificationService
	SimCardGroups                      SimCardGroupService
	SimCardOrderPreview                SimCardOrderPreviewService
	SimCardOrders                      SimCardOrderService
	SimCards                           SimCardService
	SiprecConnectors                   SiprecConnectorService
	Storage                            StorageService
	SubNumberOrders                    SubNumberOrderService
	SubNumberOrdersReport              SubNumberOrdersReportService
	TelephonyCredentials               TelephonyCredentialService
	Texml                              TexmlService
	TexmlApplications                  TexmlApplicationService
	TextToSpeech                       TextToSpeechService
	UsageReports                       UsageReportService
	UserAddresses                      UserAddressService
	UserTags                           UserTagService
	Verifications                      VerificationService
	VerifiedNumbers                    VerifiedNumberService
	VerifyProfiles                     VerifyProfileService
	VirtualCrossConnects               VirtualCrossConnectService
	VirtualCrossConnectsCoverage       VirtualCrossConnectsCoverageService
	WebhookDeliveries                  WebhookDeliveryService
	WireguardInterfaces                WireguardInterfaceService
	WireguardPeers                     WireguardPeerService
	Wireless                           WirelessService
	WirelessBlocklistValues            WirelessBlocklistValueService
	WirelessBlocklists                 WirelessBlocklistService
	PartnerCampaigns                   PartnerCampaignService
	WellKnown                          WellKnownService
	InexplicitNumberOrders             InexplicitNumberOrderService
	MobilePhoneNumbers                 MobilePhoneNumberService
	MobileVoiceConnections             MobileVoiceConnectionService
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
	r.Brand = NewBrandService(opts...)
	r.BulkSimCardActions = NewBulkSimCardActionService(opts...)
	r.BundlePricing = NewBundlePricingService(opts...)
	r.CallControlApplications = NewCallControlApplicationService(opts...)
	r.CallEvents = NewCallEventService(opts...)
	r.Calls = NewCallService(opts...)
	r.Campaign = NewCampaignService(opts...)
	r.CampaignBuilder = NewCampaignBuilderService(opts...)
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
	r.Enumeration = NewEnumerationService(opts...)
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
	r.Messsages = NewMesssageService(opts...)
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
	r.PhoneNumberAssignmentByProfile = NewPhoneNumberAssignmentByProfileService(opts...)
	r.PhoneNumberBlocks = NewPhoneNumberBlockService(opts...)
	r.PhoneNumberCampaigns = NewPhoneNumberCampaignService(opts...)
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
	r.PartnerCampaigns = NewPartnerCampaignService(opts...)
	r.WellKnown = NewWellKnownService(opts...)
	r.InexplicitNumberOrders = NewInexplicitNumberOrderService(opts...)
	r.MobilePhoneNumbers = NewMobilePhoneNumberService(opts...)
	r.MobileVoiceConnections = NewMobileVoiceConnectionService(opts...)

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

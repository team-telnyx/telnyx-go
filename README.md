# Go API client for telnyx

SIP trunking, SMS, MMS, Call Control and Telephony Data Services.

## Project Structure

```
telnyx-go-sdk/
├── api/           # API service implementations
├── models/        # Data models and structs
├── docs/          # Auto-generated documentation
├── test/          # Unit tests
├── client.go      # Main API client
├── configuration.go # Configuration management
├── response.go    # Response handling
├── utils.go       # Utility functions
└── go.mod         # Go module definition
```

## Roadmap / Caveats

This SDK is good, but not pefect. This stems from a difference in the iodmatic practices of the language itself which can be described with the following:

✅ **Good Idiomatic Practices:**

1. **Proper Context Usage**
   - All API methods properly accept and use `context.Context` as first parameter
   - Context is threaded through the entire call chain

2. **Error Handling**
   - Functions return `(result, *http.Response, error)` tuples following Go conventions
   - Proper error wrapping and custom error types (`GenericOpenAPIError`)

3. **Package Structure**
   - Single package approach is clean for an SDK
   - Clear separation of concerns with separate files

4. **Pointer Utilities**
   - Provides helpful `Ptr*` functions for creating pointers to literals
   - Common pattern in Go APIs for optional fields

5. **JSON Handling**
   - Proper struct tags for JSON marshaling/unmarshaling
   - Handles nullable/optional fields appropriately

### ❌ **Non-Idiomatic Patterns:**

1. **Method Naming**
   ```go
   // Non-idiomatic - very verbose
   CreateGroupMmsMessage(ctx context.Context) ApiCreateGroupMmsMessageRequest
   CreateGroupMmsMessageExecute(r ApiCreateGroupMmsMessageRequest) (*MessageResponse, *http.Response, error)
   
   // More idiomatic would be:
   CreateGroupMMS(ctx context.Context, req CreateGroupMMSRequest) (*MessageResponse, error)
   ```

2. **Builder Pattern Overuse**
   - The request builder pattern adds unnecessary complexity
   - Go developers typically prefer direct function calls with structs

3. **Verbose Type Names**
   ```go
   // Generated:
   type ApiCreateGroupMmsMessageRequest struct {...}
   
   // More idiomatic:
   type CreateGroupMMSRequest struct {...}
   ```

4. **Constructor Functions**
   ```go
   // Overly verbose:
   func NewAccessIPAddressListResponseSchemaWithDefaults() *AccessIPAddressListResponseSchema
   
   // More idiomatic:
   func NewAccessIPAddressListResponse() *AccessIPAddressListResponse
   ```

5. **Fluent Interface Pattern**
   - The `.Execute()` pattern is more Java-like than Go-like
   - Go prefers direct function calls

### 🤔 **Questionable Patterns:**

1. **Service Struct Pattern**
   - Having separate `*APIService` structs is okay but adds indirection
   - More idiomatic might be functions on the main client

2. **Configuration Approach**
   - The configuration struct is reasonable but quite verbose

### **More Idiomatic Alternative Would Look Like:**

```go
package telnyx

// More Go-idiomatic API design
type Client struct {
    httpClient *http.Client
    baseURL    string
    token      string
}

func NewClient(token string, opts ...Option) *Client { ... }

// Direct, simple method calls
func (c *Client) SendSMS(ctx context.Context, req *SendSMSRequest) (*Message, error) { ... }
func (c *Client) CreateCall(ctx context.Context, req *CreateCallRequest) (*Call, error) { ... }
```

These will be addressed in the near future.

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import telnyx "github.com/telnyx/telnyx-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `telnyx.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), telnyx.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `telnyx.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), telnyx.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `telnyx.ContextOperationServerIndices` and `telnyx.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), telnyx.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), telnyx.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.telnyx.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessTokensAPI* | [**CreateTelephonyCredentialToken**](docs/AccessTokensAPI.md#createtelephonycredentialtoken) | **Post** /telephony_credentials/{id}/token | Create an Access Token.
*AddressesAPI* | [**AcceptAddressSuggestions**](docs/AddressesAPI.md#acceptaddresssuggestions) | **Post** /addresses/{id}/actions/accept_suggestions | Accepts this address suggestion as a new emergency address for Operator Connect and finishes the uploads of the numbers associated with it to Microsoft.
*AddressesAPI* | [**CreateAddress**](docs/AddressesAPI.md#createaddress) | **Post** /addresses | Creates an address
*AddressesAPI* | [**DeleteAddress**](docs/AddressesAPI.md#deleteaddress) | **Delete** /addresses/{id} | Deletes an address
*AddressesAPI* | [**FindAddresses**](docs/AddressesAPI.md#findaddresses) | **Get** /addresses | List all addresses
*AddressesAPI* | [**GetAddress**](docs/AddressesAPI.md#getaddress) | **Get** /addresses/{id} | Retrieve an address
*AddressesAPI* | [**ValidateAddress**](docs/AddressesAPI.md#validateaddress) | **Post** /addresses/actions/validate | Validate an address
*AdvancedNumberOrdersAPI* | [**CreateAdvancedOrderV2**](docs/AdvancedNumberOrdersAPI.md#createadvancedorderv2) | **Post** /advanced_orders | Create Advanced Order
*AdvancedNumberOrdersAPI* | [**GetAdvancedOrderV2**](docs/AdvancedNumberOrdersAPI.md#getadvancedorderv2) | **Get** /advanced_orders/{order_id} | Get Advanced Order
*AdvancedNumberOrdersAPI* | [**ListAdvancedOrdersV2**](docs/AdvancedNumberOrdersAPI.md#listadvancedordersv2) | **Get** /advanced_orders | List Advanced Orders
*AdvancedOptInOptOutAPI* | [**CreateAutorespConfig**](docs/AdvancedOptInOptOutAPI.md#createautorespconfig) | **Post** /messaging_profiles/{profile_id}/autoresp_configs | Create Auto-Reponse Setting
*AdvancedOptInOptOutAPI* | [**DeleteAutorespConfig**](docs/AdvancedOptInOptOutAPI.md#deleteautorespconfig) | **Delete** /messaging_profiles/{profile_id}/autoresp_configs/{autoresp_cfg_id} | Delete Auto-Response Setting
*AdvancedOptInOptOutAPI* | [**GetAutorespConfig**](docs/AdvancedOptInOptOutAPI.md#getautorespconfig) | **Get** /messaging_profiles/{profile_id}/autoresp_configs/{autoresp_cfg_id} | Get Auto-Response Setting
*AdvancedOptInOptOutAPI* | [**GetAutorespConfigs**](docs/AdvancedOptInOptOutAPI.md#getautorespconfigs) | **Get** /messaging_profiles/{profile_id}/autoresp_configs | List Auto-Response Settings
*AdvancedOptInOptOutAPI* | [**ListOptOuts**](docs/AdvancedOptInOptOutAPI.md#listoptouts) | **Get** /messaging_optouts | List opt-outs
*AdvancedOptInOptOutAPI* | [**UpdateAutoRespConfig**](docs/AdvancedOptInOptOutAPI.md#updateautorespconfig) | **Put** /messaging_profiles/{profile_id}/autoresp_configs/{autoresp_cfg_id} | Update Auto-Response Setting
*AssistantsAPI* | [**CreateNewAssistantPublicAssistantsPost**](docs/AssistantsAPI.md#createnewassistantpublicassistantspost) | **Post** /ai/assistants | Create an assistant
*AssistantsAPI* | [**CreateScheduledEvent**](docs/AssistantsAPI.md#createscheduledevent) | **Post** /ai/assistants/{assistant_id}/scheduled_events | Create a scheduled event
*AssistantsAPI* | [**DeleteAssistantPublicAssistantsAssistantIdDelete**](docs/AssistantsAPI.md#deleteassistantpublicassistantsassistantiddelete) | **Delete** /ai/assistants/{assistant_id} | Delete an assistant
*AssistantsAPI* | [**DeleteScheduledEvent**](docs/AssistantsAPI.md#deletescheduledevent) | **Delete** /ai/assistants/{assistant_id}/scheduled_events/{event_id} | Delete a scheduled event
*AssistantsAPI* | [**GetAssistantPublicAssistantsAssistantIdGet**](docs/AssistantsAPI.md#getassistantpublicassistantsassistantidget) | **Get** /ai/assistants/{assistant_id} | Get an assistant
*AssistantsAPI* | [**GetAssistantTexmlPublicAssistantsAssistantIdTexmlGet**](docs/AssistantsAPI.md#getassistanttexmlpublicassistantsassistantidtexmlget) | **Get** /ai/assistants/{assistant_id}/texml | Get assistant texml
*AssistantsAPI* | [**GetAssistantsPublicAssistantsGet**](docs/AssistantsAPI.md#getassistantspublicassistantsget) | **Get** /ai/assistants | List assistants
*AssistantsAPI* | [**GetScheduledEvent**](docs/AssistantsAPI.md#getscheduledevent) | **Get** /ai/assistants/{assistant_id}/scheduled_events/{event_id} | Get a scheduled event
*AssistantsAPI* | [**GetScheduledEvents**](docs/AssistantsAPI.md#getscheduledevents) | **Get** /ai/assistants/{assistant_id}/scheduled_events | List scheduled events
*AssistantsAPI* | [**UpdateAssistantPublicAssistantsAssistantIdPost**](docs/AssistantsAPI.md#updateassistantpublicassistantsassistantidpost) | **Post** /ai/assistants/{assistant_id} | Update an assistant
*AudioAPI* | [**AudioPublicAudioTranscriptionsPost**](docs/AudioAPI.md#audiopublicaudiotranscriptionspost) | **Post** /ai/audio/transcriptions | Transcribe speech to text
*AuditLogsAPI* | [**ListAuditLogs**](docs/AuditLogsAPI.md#listauditlogs) | **Get** /audit_events | List Audit Logs
*AuthenticationProvidersAPI* | [**CreateAuthenticationProvider**](docs/AuthenticationProvidersAPI.md#createauthenticationprovider) | **Post** /authentication_providers | Creates an authentication provider
*AuthenticationProvidersAPI* | [**DeleteAuthenticationProvider**](docs/AuthenticationProvidersAPI.md#deleteauthenticationprovider) | **Delete** /authentication_providers/{id} | Deletes an authentication provider
*AuthenticationProvidersAPI* | [**FindAuthenticationProviders**](docs/AuthenticationProvidersAPI.md#findauthenticationproviders) | **Get** /authentication_providers | List all SSO authentication providers
*AuthenticationProvidersAPI* | [**GetAuthenticationProvider**](docs/AuthenticationProvidersAPI.md#getauthenticationprovider) | **Get** /authentication_providers/{id} | Retrieve an authentication provider
*AuthenticationProvidersAPI* | [**UpdateAuthenticationProvider**](docs/AuthenticationProvidersAPI.md#updateauthenticationprovider) | **Patch** /authentication_providers/{id} | Update an authentication provider
*AutoRechargePreferencesAPI* | [**GetAutoRechargePrefs**](docs/AutoRechargePreferencesAPI.md#getautorechargeprefs) | **Get** /payment/auto_recharge_prefs | List auto recharge preferences
*AutoRechargePreferencesAPI* | [**UpdateAutoRechargePrefs**](docs/AutoRechargePreferencesAPI.md#updateautorechargeprefs) | **Patch** /payment/auto_recharge_prefs | Update auto recharge preferences
*BillingAPI* | [**GetUserBalance**](docs/BillingAPI.md#getuserbalance) | **Get** /balance | Get user balance details
*BillingGroupsAPI* | [**CreateBillingGroup**](docs/BillingGroupsAPI.md#createbillinggroup) | **Post** /billing_groups | Create a billing group
*BillingGroupsAPI* | [**DeleteBillingGroup**](docs/BillingGroupsAPI.md#deletebillinggroup) | **Delete** /billing_groups/{id} | Delete a billing group
*BillingGroupsAPI* | [**GetBillingGroup**](docs/BillingGroupsAPI.md#getbillinggroup) | **Get** /billing_groups/{id} | Get a billing group
*BillingGroupsAPI* | [**ListBillingGroups**](docs/BillingGroupsAPI.md#listbillinggroups) | **Get** /billing_groups | List all billing groups
*BillingGroupsAPI* | [**UpdateBillingGroup**](docs/BillingGroupsAPI.md#updatebillinggroup) | **Patch** /billing_groups/{id} | Update a billing group
*BrandsAPI* | [**CreateBrandPost**](docs/BrandsAPI.md#createbrandpost) | **Post** /brand | Create Brand
*BrandsAPI* | [**DeleteBrand**](docs/BrandsAPI.md#deletebrand) | **Delete** /brand/{brandId} | Delete Brand
*BrandsAPI* | [**GetBrand**](docs/BrandsAPI.md#getbrand) | **Get** /brand/{brandId} | Get Brand
*BrandsAPI* | [**GetBrandFeedbackById**](docs/BrandsAPI.md#getbrandfeedbackbyid) | **Get** /brand/feedback/{brandId} | Get Brand Feedback By Id
*BrandsAPI* | [**GetBrands**](docs/BrandsAPI.md#getbrands) | **Get** /brand | List Brands
*BrandsAPI* | [**ListExternalVettings**](docs/BrandsAPI.md#listexternalvettings) | **Get** /brand/{brandId}/externalVetting | List External Vettings
*BrandsAPI* | [**PostOrderExternalVetting**](docs/BrandsAPI.md#postorderexternalvetting) | **Post** /brand/{brandId}/externalVetting | Order Brand External Vetting
*BrandsAPI* | [**PutExternalVettingRecord**](docs/BrandsAPI.md#putexternalvettingrecord) | **Put** /brand/{brandId}/externalVetting | Import External Vetting Record
*BrandsAPI* | [**ResendBrand2faEmail**](docs/BrandsAPI.md#resendbrand2faemail) | **Post** /brand/{brandId}/2faEmail | Resend brand 2FA email
*BrandsAPI* | [**RevetBrand**](docs/BrandsAPI.md#revetbrand) | **Put** /brand/{brandId}/revet | Revet Brand
*BrandsAPI* | [**UpdateBrand**](docs/BrandsAPI.md#updatebrand) | **Put** /brand/{brandId} | Update Brand
*BucketAPI* | [**CreateBucket**](docs/BucketAPI.md#createbucket) | **Put** /{bucketName} | CreateBucket
*BucketAPI* | [**DeleteBucket**](docs/BucketAPI.md#deletebucket) | **Delete** /{bucketName} | DeleteBucket
*BucketAPI* | [**HeadBucket**](docs/BucketAPI.md#headbucket) | **Head** /{bucketName} | HeadBucket
*BucketAPI* | [**ListBuckets**](docs/BucketAPI.md#listbuckets) | **Get** / | ListBuckets
*BucketSSLCertificateAPI* | [**AddStorageSSLCertificate**](docs/BucketSSLCertificateAPI.md#addstoragesslcertificate) | **Put** /storage/buckets/{bucketName}/ssl_certificate | Add SSL Certificate
*BucketSSLCertificateAPI* | [**GetStorageSSLCertificates**](docs/BucketSSLCertificateAPI.md#getstoragesslcertificates) | **Get** /storage/buckets/{bucketName}/ssl_certificate | Get Bucket SSL Certificate
*BucketSSLCertificateAPI* | [**RemoveStorageSSLCertificate**](docs/BucketSSLCertificateAPI.md#removestoragesslcertificate) | **Delete** /storage/buckets/{bucketName}/ssl_certificate | Remove SSL Certificate
*BucketUsageAPI* | [**GetBucketUsage**](docs/BucketUsageAPI.md#getbucketusage) | **Get** /storage/buckets/{bucketName}/usage/storage | Get Bucket Usage
*BucketUsageAPI* | [**GetStorageAPIUsage**](docs/BucketUsageAPI.md#getstorageapiusage) | **Get** /storage/buckets/{bucketName}/usage/api | Get API Usage
*BulkPhoneNumberCampaignsAPI* | [**GetAssignmentTaskStatus**](docs/BulkPhoneNumberCampaignsAPI.md#getassignmenttaskstatus) | **Get** /phoneNumberAssignmentByProfile/{taskId} | Get Assignment Task Status
*BulkPhoneNumberCampaignsAPI* | [**GetPhoneNumberStatus**](docs/BulkPhoneNumberCampaignsAPI.md#getphonenumberstatus) | **Get** /phoneNumberAssignmentByProfile/{taskId}/phoneNumbers | Get Phone Number Status
*BulkPhoneNumberCampaignsAPI* | [**PostAssignMessagingProfileToCampaign**](docs/BulkPhoneNumberCampaignsAPI.md#postassignmessagingprofiletocampaign) | **Post** /phoneNumberAssignmentByProfile | Assign Messaging Profile To Campaign
*BulkPhoneNumberOperationsAPI* | [**CreateDeletePhoneNumbersJob**](docs/BulkPhoneNumberOperationsAPI.md#createdeletephonenumbersjob) | **Post** /phone_numbers/jobs/delete_phone_numbers | Delete a batch of numbers
*BulkPhoneNumberOperationsAPI* | [**CreatePhoneNumbersJobUpdateEmergencySettings**](docs/BulkPhoneNumberOperationsAPI.md#createphonenumbersjobupdateemergencysettings) | **Post** /phone_numbers/jobs/update_emergency_settings | Update the emergency settings from a batch of numbers
*BulkPhoneNumberOperationsAPI* | [**CreateUpdatePhoneNumbersJob**](docs/BulkPhoneNumberOperationsAPI.md#createupdatephonenumbersjob) | **Post** /phone_numbers/jobs/update_phone_numbers | Update a batch of numbers
*BulkPhoneNumberOperationsAPI* | [**ListPhoneNumbersJobs**](docs/BulkPhoneNumberOperationsAPI.md#listphonenumbersjobs) | **Get** /phone_numbers/jobs | Lists the phone numbers jobs
*BulkPhoneNumberOperationsAPI* | [**RetrievePhoneNumbersJob**](docs/BulkPhoneNumberOperationsAPI.md#retrievephonenumbersjob) | **Get** /phone_numbers/jobs/{id} | Retrieve a phone numbers job
*BundlesAPI* | [**GetBillingBundleById**](docs/BundlesAPI.md#getbillingbundlebyid) | **Get** /bundle_pricing/billing_bundles/{bundle_id} | Get Bundle By Id
*BundlesAPI* | [**GetUserBillingBundles**](docs/BundlesAPI.md#getuserbillingbundles) | **Get** /bundle_pricing/billing_bundles | Retrieve Bundles
*CDRUsageReportsAPI* | [**GetCDRUsageReportSync**](docs/CDRUsageReportsAPI.md#getcdrusagereportsync) | **Get** /reports/cdr_usage_reports/sync | Generates and fetches CDR Usage Reports
*CSVDownloadsAPI* | [**CreateCsvDownload**](docs/CSVDownloadsAPI.md#createcsvdownload) | **Post** /phone_numbers/csv_downloads | Create a CSV download
*CSVDownloadsAPI* | [**GetCsvDownload**](docs/CSVDownloadsAPI.md#getcsvdownload) | **Get** /phone_numbers/csv_downloads/{id} | Retrieve a CSV download
*CSVDownloadsAPI* | [**ListCsvDownloads**](docs/CSVDownloadsAPI.md#listcsvdownloads) | **Get** /phone_numbers/csv_downloads | List CSV downloads
*CallCommandsAPI* | [**AnswerCall**](docs/CallCommandsAPI.md#answercall) | **Post** /calls/{call_control_id}/actions/answer | Answer call
*CallCommandsAPI* | [**BridgeCall**](docs/CallCommandsAPI.md#bridgecall) | **Post** /calls/{call_control_id}/actions/bridge | Bridge calls
*CallCommandsAPI* | [**CallGatherUsingAI**](docs/CallCommandsAPI.md#callgatherusingai) | **Post** /calls/{call_control_id}/actions/gather_using_ai | Gather using AI
*CallCommandsAPI* | [**CallStartAIAssistant**](docs/CallCommandsAPI.md#callstartaiassistant) | **Post** /calls/{call_control_id}/actions/ai_assistant_start | Start AI Assistant
*CallCommandsAPI* | [**CallStopAIAssistant**](docs/CallCommandsAPI.md#callstopaiassistant) | **Post** /calls/{call_control_id}/actions/ai_assistant_stop | Stop AI Assistant
*CallCommandsAPI* | [**DialCall**](docs/CallCommandsAPI.md#dialcall) | **Post** /calls | Dial
*CallCommandsAPI* | [**EnqueueCall**](docs/CallCommandsAPI.md#enqueuecall) | **Post** /calls/{call_control_id}/actions/enqueue | Enqueue call
*CallCommandsAPI* | [**GatherCall**](docs/CallCommandsAPI.md#gathercall) | **Post** /calls/{call_control_id}/actions/gather | Gather
*CallCommandsAPI* | [**GatherUsingAudio**](docs/CallCommandsAPI.md#gatherusingaudio) | **Post** /calls/{call_control_id}/actions/gather_using_audio | Gather using audio
*CallCommandsAPI* | [**GatherUsingSpeak**](docs/CallCommandsAPI.md#gatherusingspeak) | **Post** /calls/{call_control_id}/actions/gather_using_speak | Gather using speak
*CallCommandsAPI* | [**HangupCall**](docs/CallCommandsAPI.md#hangupcall) | **Post** /calls/{call_control_id}/actions/hangup | Hangup call
*CallCommandsAPI* | [**LeaveQueue**](docs/CallCommandsAPI.md#leavequeue) | **Post** /calls/{call_control_id}/actions/leave_queue | Remove call from a queue
*CallCommandsAPI* | [**NoiseSuppressionStart**](docs/CallCommandsAPI.md#noisesuppressionstart) | **Post** /calls/{call_control_id}/actions/suppression_start | Noise Suppression Start (BETA)
*CallCommandsAPI* | [**NoiseSuppressionStop**](docs/CallCommandsAPI.md#noisesuppressionstop) | **Post** /calls/{call_control_id}/actions/suppression_stop | Noise Suppression Stop (BETA)
*CallCommandsAPI* | [**PauseCallRecording**](docs/CallCommandsAPI.md#pausecallrecording) | **Post** /calls/{call_control_id}/actions/record_pause | Record pause
*CallCommandsAPI* | [**ReferCall**](docs/CallCommandsAPI.md#refercall) | **Post** /calls/{call_control_id}/actions/refer | SIP Refer a call
*CallCommandsAPI* | [**RejectCall**](docs/CallCommandsAPI.md#rejectcall) | **Post** /calls/{call_control_id}/actions/reject | Reject a call
*CallCommandsAPI* | [**ResumeCallRecording**](docs/CallCommandsAPI.md#resumecallrecording) | **Post** /calls/{call_control_id}/actions/record_resume | Record resume
*CallCommandsAPI* | [**SendDTMF**](docs/CallCommandsAPI.md#senddtmf) | **Post** /calls/{call_control_id}/actions/send_dtmf | Send DTMF
*CallCommandsAPI* | [**SendSIPInfo**](docs/CallCommandsAPI.md#sendsipinfo) | **Post** /calls/{call_control_id}/actions/send_sip_info | Send SIP info
*CallCommandsAPI* | [**SpeakCall**](docs/CallCommandsAPI.md#speakcall) | **Post** /calls/{call_control_id}/actions/speak | Speak text
*CallCommandsAPI* | [**StartCallFork**](docs/CallCommandsAPI.md#startcallfork) | **Post** /calls/{call_control_id}/actions/fork_start | Forking start
*CallCommandsAPI* | [**StartCallPlayback**](docs/CallCommandsAPI.md#startcallplayback) | **Post** /calls/{call_control_id}/actions/playback_start | Play audio URL
*CallCommandsAPI* | [**StartCallRecord**](docs/CallCommandsAPI.md#startcallrecord) | **Post** /calls/{call_control_id}/actions/record_start | Recording start
*CallCommandsAPI* | [**StartCallStreaming**](docs/CallCommandsAPI.md#startcallstreaming) | **Post** /calls/{call_control_id}/actions/streaming_start | Streaming start
*CallCommandsAPI* | [**StartCallTranscription**](docs/CallCommandsAPI.md#startcalltranscription) | **Post** /calls/{call_control_id}/actions/transcription_start | Transcription start
*CallCommandsAPI* | [**StartSiprecSession**](docs/CallCommandsAPI.md#startsiprecsession) | **Post** /calls/{call_control_id}/actions/siprec_start | SIPREC start
*CallCommandsAPI* | [**StopCallFork**](docs/CallCommandsAPI.md#stopcallfork) | **Post** /calls/{call_control_id}/actions/fork_stop | Forking stop
*CallCommandsAPI* | [**StopCallGather**](docs/CallCommandsAPI.md#stopcallgather) | **Post** /calls/{call_control_id}/actions/gather_stop | Gather stop
*CallCommandsAPI* | [**StopCallPlayback**](docs/CallCommandsAPI.md#stopcallplayback) | **Post** /calls/{call_control_id}/actions/playback_stop | Stop audio playback
*CallCommandsAPI* | [**StopCallRecording**](docs/CallCommandsAPI.md#stopcallrecording) | **Post** /calls/{call_control_id}/actions/record_stop | Recording stop
*CallCommandsAPI* | [**StopCallStreaming**](docs/CallCommandsAPI.md#stopcallstreaming) | **Post** /calls/{call_control_id}/actions/streaming_stop | Streaming stop
*CallCommandsAPI* | [**StopCallTranscription**](docs/CallCommandsAPI.md#stopcalltranscription) | **Post** /calls/{call_control_id}/actions/transcription_stop | Transcription stop
*CallCommandsAPI* | [**StopSiprecSession**](docs/CallCommandsAPI.md#stopsiprecsession) | **Post** /calls/{call_control_id}/actions/siprec_stop | SIPREC stop
*CallCommandsAPI* | [**TransferCall**](docs/CallCommandsAPI.md#transfercall) | **Post** /calls/{call_control_id}/actions/transfer | Transfer call
*CallCommandsAPI* | [**UpdateClientState**](docs/CallCommandsAPI.md#updateclientstate) | **Put** /calls/{call_control_id}/actions/client_state_update | Update client state
*CallControlApplicationsAPI* | [**CreateCallControlApplication**](docs/CallControlApplicationsAPI.md#createcallcontrolapplication) | **Post** /call_control_applications | Create a call control application
*CallControlApplicationsAPI* | [**DeleteCallControlApplication**](docs/CallControlApplicationsAPI.md#deletecallcontrolapplication) | **Delete** /call_control_applications/{id} | Delete a call control application
*CallControlApplicationsAPI* | [**ListCallControlApplications**](docs/CallControlApplicationsAPI.md#listcallcontrolapplications) | **Get** /call_control_applications | List call control applications
*CallControlApplicationsAPI* | [**RetrieveCallControlApplication**](docs/CallControlApplicationsAPI.md#retrievecallcontrolapplication) | **Get** /call_control_applications/{id} | Retrieve a call control application
*CallControlApplicationsAPI* | [**UpdateCallControlApplication**](docs/CallControlApplicationsAPI.md#updatecallcontrolapplication) | **Patch** /call_control_applications/{id} | Update a call control application
*CallInformationAPI* | [**ListConnectionActiveCalls**](docs/CallInformationAPI.md#listconnectionactivecalls) | **Get** /connections/{connection_id}/active_calls | List all active calls for given connection
*CallInformationAPI* | [**RetrieveCallStatus**](docs/CallInformationAPI.md#retrievecallstatus) | **Get** /calls/{call_control_id} | Retrieve a call status
*CallRecordingsAPI* | [**CreateCustomStorageCredentials**](docs/CallRecordingsAPI.md#createcustomstoragecredentials) | **Post** /custom_storage_credentials/{connection_id} | Create a custom storage credential
*CallRecordingsAPI* | [**DeleteCustomStorageCredentials**](docs/CallRecordingsAPI.md#deletecustomstoragecredentials) | **Delete** /custom_storage_credentials/{connection_id} | Delete a stored credential
*CallRecordingsAPI* | [**DeleteRecording**](docs/CallRecordingsAPI.md#deleterecording) | **Delete** /recordings/{recording_id} | Delete a call recording
*CallRecordingsAPI* | [**DeleteRecordingTranscription**](docs/CallRecordingsAPI.md#deleterecordingtranscription) | **Delete** /recording_transcriptions/{recording_transcription_id} | Delete a recording transcription
*CallRecordingsAPI* | [**DeleteRecordings**](docs/CallRecordingsAPI.md#deleterecordings) | **Delete** /recordings/actions/delete | Delete a list of call recordings
*CallRecordingsAPI* | [**GetCustomStorageCredentials**](docs/CallRecordingsAPI.md#getcustomstoragecredentials) | **Get** /custom_storage_credentials/{connection_id} | Retrieve a stored credential
*CallRecordingsAPI* | [**GetRecording**](docs/CallRecordingsAPI.md#getrecording) | **Get** /recordings/{recording_id} | Retrieve a call recording
*CallRecordingsAPI* | [**GetRecordingTranscription**](docs/CallRecordingsAPI.md#getrecordingtranscription) | **Get** /recording_transcriptions/{recording_transcription_id} | Retrieve a recording transcription
*CallRecordingsAPI* | [**GetRecordingTranscriptions**](docs/CallRecordingsAPI.md#getrecordingtranscriptions) | **Get** /recording_transcriptions | List all recording transcriptions
*CallRecordingsAPI* | [**GetRecordings**](docs/CallRecordingsAPI.md#getrecordings) | **Get** /recordings | List all call recordings
*CallRecordingsAPI* | [**UpdateCustomStorageCredentials**](docs/CallRecordingsAPI.md#updatecustomstoragecredentials) | **Put** /custom_storage_credentials/{connection_id} | Update a stored credential
*CampaignAPI* | [**AcceptCampaign**](docs/CampaignAPI.md#acceptcampaign) | **Post** /campaign/acceptSharing/{campaignId} | Accept Shared Campaign
*CampaignAPI* | [**DeactivateCampaign**](docs/CampaignAPI.md#deactivatecampaign) | **Delete** /campaign/{campaignId} | Deactivate My Campaign
*CampaignAPI* | [**GetCampaign**](docs/CampaignAPI.md#getcampaign) | **Get** /campaign/{campaignId} | Get My Campaign
*CampaignAPI* | [**GetCampaignCost**](docs/CampaignAPI.md#getcampaigncost) | **Get** /campaign/usecase/cost | Get Campaign Cost
*CampaignAPI* | [**GetCampaignMnoMetadata**](docs/CampaignAPI.md#getcampaignmnometadata) | **Get** /campaign/{campaignId}/mnoMetadata | Get Campaign Mno Metadata
*CampaignAPI* | [**GetCampaignOperationStatus**](docs/CampaignAPI.md#getcampaignoperationstatus) | **Get** /campaign/{campaignId}/operationStatus | Get My Campaign Operation Status
*CampaignAPI* | [**GetCampaignOsrAttributes**](docs/CampaignAPI.md#getcampaignosrattributes) | **Get** /campaign/{campaignId}/osr/attributes | Get My Osr Campaign Attributes
*CampaignAPI* | [**GetCampaignSharingStatus**](docs/CampaignAPI.md#getcampaignsharingstatus) | **Get** /campaign/{campaignId}/sharing | Get Sharing Status
*CampaignAPI* | [**GetCampaigns**](docs/CampaignAPI.md#getcampaigns) | **Get** /campaign | List Campaigns
*CampaignAPI* | [**GetUsecaseQualification**](docs/CampaignAPI.md#getusecasequalification) | **Get** /campaignBuilder/brand/{brandId}/usecase/{usecase} | Qualify By Usecase
*CampaignAPI* | [**PostCampaign**](docs/CampaignAPI.md#postcampaign) | **Post** /campaignBuilder | Submit Campaign
*CampaignAPI* | [**UpdateCampaign**](docs/CampaignAPI.md#updatecampaign) | **Put** /campaign/{campaignId} | Update My Campaign
*ChatAPI* | [**ChatPublicChatCompletionsPost**](docs/ChatAPI.md#chatpublicchatcompletionspost) | **Post** /ai/chat/completions | Create a chat completion
*ChatAPI* | [**GetModelsPublicModelsGet**](docs/ChatAPI.md#getmodelspublicmodelsget) | **Get** /ai/models | Get available models
*ChatAPI* | [**PostSummary**](docs/ChatAPI.md#postsummary) | **Post** /ai/summarize | Summarize file content
*ClustersAPI* | [**ComputeNewClusterPublicTextClustersPost**](docs/ClustersAPI.md#computenewclusterpublictextclusterspost) | **Post** /ai/clusters | Compute new clusters
*ClustersAPI* | [**DeleteClusterByTaskIdPublicTextClustersTaskIdDelete**](docs/ClustersAPI.md#deleteclusterbytaskidpublictextclusterstaskiddelete) | **Delete** /ai/clusters/{task_id} | Delete a cluster
*ClustersAPI* | [**FetchClusterByTaskIdPublicTextClustersTaskIdGet**](docs/ClustersAPI.md#fetchclusterbytaskidpublictextclusterstaskidget) | **Get** /ai/clusters/{task_id} | Fetch a cluster
*ClustersAPI* | [**FetchClusterImageByTaskIdPublicTextClustersTaskIdImageGet**](docs/ClustersAPI.md#fetchclusterimagebytaskidpublictextclusterstaskidimageget) | **Get** /ai/clusters/{task_id}/graph | Fetch a cluster visualization
*ClustersAPI* | [**ListAllRequestedClustersPublicTextClustersGet**](docs/ClustersAPI.md#listallrequestedclusterspublictextclustersget) | **Get** /ai/clusters | List all clusters
*ConferenceCommandsAPI* | [**CreateConference**](docs/ConferenceCommandsAPI.md#createconference) | **Post** /conferences | Create conference
*ConferenceCommandsAPI* | [**HoldConferenceParticipants**](docs/ConferenceCommandsAPI.md#holdconferenceparticipants) | **Post** /conferences/{id}/actions/hold | Hold conference participants
*ConferenceCommandsAPI* | [**JoinConference**](docs/ConferenceCommandsAPI.md#joinconference) | **Post** /conferences/{id}/actions/join | Join a conference
*ConferenceCommandsAPI* | [**LeaveConference**](docs/ConferenceCommandsAPI.md#leaveconference) | **Post** /conferences/{id}/actions/leave | Leave a conference
*ConferenceCommandsAPI* | [**ListConferenceParticipants**](docs/ConferenceCommandsAPI.md#listconferenceparticipants) | **Get** /conferences/{conference_id}/participants | List conference participants
*ConferenceCommandsAPI* | [**ListConferences**](docs/ConferenceCommandsAPI.md#listconferences) | **Get** /conferences | List conferences
*ConferenceCommandsAPI* | [**MuteConferenceParticipants**](docs/ConferenceCommandsAPI.md#muteconferenceparticipants) | **Post** /conferences/{id}/actions/mute | Mute conference participants
*ConferenceCommandsAPI* | [**PauseConferenceRecording**](docs/ConferenceCommandsAPI.md#pauseconferencerecording) | **Post** /conferences/{id}/actions/record_pause | Conference recording pause
*ConferenceCommandsAPI* | [**PlayConferenceAudio**](docs/ConferenceCommandsAPI.md#playconferenceaudio) | **Post** /conferences/{id}/actions/play | Play audio to conference participants
*ConferenceCommandsAPI* | [**ResumeConferenceRecording**](docs/ConferenceCommandsAPI.md#resumeconferencerecording) | **Post** /conferences/{id}/actions/record_resume | Conference recording resume
*ConferenceCommandsAPI* | [**RetrieveConference**](docs/ConferenceCommandsAPI.md#retrieveconference) | **Get** /conferences/{id} | Retrieve a conference
*ConferenceCommandsAPI* | [**SpeakTextToConference**](docs/ConferenceCommandsAPI.md#speaktexttoconference) | **Post** /conferences/{id}/actions/speak | Speak text to conference participants
*ConferenceCommandsAPI* | [**StartConferenceRecording**](docs/ConferenceCommandsAPI.md#startconferencerecording) | **Post** /conferences/{id}/actions/record_start | Conference recording start
*ConferenceCommandsAPI* | [**StopConferenceAudio**](docs/ConferenceCommandsAPI.md#stopconferenceaudio) | **Post** /conferences/{id}/actions/stop | Stop audio being played on the conference
*ConferenceCommandsAPI* | [**StopConferenceRecording**](docs/ConferenceCommandsAPI.md#stopconferencerecording) | **Post** /conferences/{id}/actions/record_stop | Conference recording stop
*ConferenceCommandsAPI* | [**UnholdConferenceParticipants**](docs/ConferenceCommandsAPI.md#unholdconferenceparticipants) | **Post** /conferences/{id}/actions/unhold | Unhold conference participants
*ConferenceCommandsAPI* | [**UnmuteConferenceParticipants**](docs/ConferenceCommandsAPI.md#unmuteconferenceparticipants) | **Post** /conferences/{id}/actions/unmute | Unmute conference participants
*ConferenceCommandsAPI* | [**UpdateConference**](docs/ConferenceCommandsAPI.md#updateconference) | **Post** /conferences/{id}/actions/update | Update conference participant
*ConnectionsAPI* | [**ListConnections**](docs/ConnectionsAPI.md#listconnections) | **Get** /connections | List connections
*ConnectionsAPI* | [**RetrieveConnection**](docs/ConnectionsAPI.md#retrieveconnection) | **Get** /connections/{id} | Retrieve a connection
*ConversationsAPI* | [**CreateNewConversationPublicConversationsPost**](docs/ConversationsAPI.md#createnewconversationpublicconversationspost) | **Post** /ai/conversations | Create a conversation
*ConversationsAPI* | [**DeleteConversationByIdPublicConversationsDelete**](docs/ConversationsAPI.md#deleteconversationbyidpublicconversationsdelete) | **Delete** /ai/conversations/{conversation_id} | Delete a conversation
*ConversationsAPI* | [**GetConversationByIdPublicConversationsGet**](docs/ConversationsAPI.md#getconversationbyidpublicconversationsget) | **Get** /ai/conversations/{conversation_id} | Get a conversation
*ConversationsAPI* | [**GetConversationsPublicConversationIdInsightsGet**](docs/ConversationsAPI.md#getconversationspublicconversationidinsightsget) | **Get** /ai/conversations/{conversation_id}/conversations-insights | Get insights for a conversation
*ConversationsAPI* | [**GetConversationsPublicConversationIdMessagesGet**](docs/ConversationsAPI.md#getconversationspublicconversationidmessagesget) | **Get** /ai/conversations/{conversation_id}/messages | Get conversation messages
*ConversationsAPI* | [**GetConversationsPublicConversationsGet**](docs/ConversationsAPI.md#getconversationspublicconversationsget) | **Get** /ai/conversations | List conversations
*ConversationsAPI* | [**UpdateConversationByIdPublicConversationsPut**](docs/ConversationsAPI.md#updateconversationbyidpublicconversationsput) | **Put** /ai/conversations/{conversation_id} | Update conversation metadata
*CountryCoverageAPI* | [**RetreiveCountryCoverage**](docs/CountryCoverageAPI.md#retreivecountrycoverage) | **Get** /country_coverage | Get country coverage
*CountryCoverageAPI* | [**RetreiveSpecificCountryCoverage**](docs/CountryCoverageAPI.md#retreivespecificcountrycoverage) | **Get** /country_coverage/countries/{country_code} | Get coverage for a specific country
*CoverageAPI* | [**ListNetworkCoverage**](docs/CoverageAPI.md#listnetworkcoverage) | **Get** /network_coverage | List network coverage locations
*CredentialConnectionsAPI* | [**CheckRegistrationStatus**](docs/CredentialConnectionsAPI.md#checkregistrationstatus) | **Post** /credential_connections/{id}/actions/check_registration_status | Check a Credential Connection Registration Status
*CredentialConnectionsAPI* | [**CreateCredentialConnection**](docs/CredentialConnectionsAPI.md#createcredentialconnection) | **Post** /credential_connections | Create a credential connection
*CredentialConnectionsAPI* | [**DeleteCredentialConnection**](docs/CredentialConnectionsAPI.md#deletecredentialconnection) | **Delete** /credential_connections/{id} | Delete a credential connection
*CredentialConnectionsAPI* | [**ListCredentialConnections**](docs/CredentialConnectionsAPI.md#listcredentialconnections) | **Get** /credential_connections | List credential connections
*CredentialConnectionsAPI* | [**RetrieveCredentialConnection**](docs/CredentialConnectionsAPI.md#retrievecredentialconnection) | **Get** /credential_connections/{id} | Retrieve a credential connection
*CredentialConnectionsAPI* | [**UpdateCredentialConnection**](docs/CredentialConnectionsAPI.md#updatecredentialconnection) | **Patch** /credential_connections/{id} | Update a credential connection
*CredentialsAPI* | [**CreateTelephonyCredential**](docs/CredentialsAPI.md#createtelephonycredential) | **Post** /telephony_credentials | Create a credential
*CredentialsAPI* | [**DeleteTelephonyCredential**](docs/CredentialsAPI.md#deletetelephonycredential) | **Delete** /telephony_credentials/{id} | Delete a credential
*CredentialsAPI* | [**FindTelephonyCredentials**](docs/CredentialsAPI.md#findtelephonycredentials) | **Get** /telephony_credentials | List all credentials
*CredentialsAPI* | [**GetTelephonyCredential**](docs/CredentialsAPI.md#gettelephonycredential) | **Get** /telephony_credentials/{id} | Get a credential
*CredentialsAPI* | [**UpdateTelephonyCredential**](docs/CredentialsAPI.md#updatetelephonycredential) | **Patch** /telephony_credentials/{id} | Update a credential
*CustomerServiceRecordAPI* | [**CreateCustomerServiceRecord**](docs/CustomerServiceRecordAPI.md#createcustomerservicerecord) | **Post** /customer_service_records | Create a customer service record
*CustomerServiceRecordAPI* | [**GetCustomerServiceRecord**](docs/CustomerServiceRecordAPI.md#getcustomerservicerecord) | **Get** /customer_service_records/{customer_service_record_id} | Get a customer service record
*CustomerServiceRecordAPI* | [**ListCustomerServiceRecords**](docs/CustomerServiceRecordAPI.md#listcustomerservicerecords) | **Get** /customer_service_records | List customer service records
*CustomerServiceRecordAPI* | [**VerifyPhoneNumberCoverage**](docs/CustomerServiceRecordAPI.md#verifyphonenumbercoverage) | **Post** /customer_service_records/phone_number_coverages | Verify CSR phone number coverage
*DataMigrationAPI* | [**CreateMigration**](docs/DataMigrationAPI.md#createmigration) | **Post** /storage/migrations | Create a Migration
*DataMigrationAPI* | [**CreateMigrationSource**](docs/DataMigrationAPI.md#createmigrationsource) | **Post** /storage/migration_sources | Create a Migration Source
*DataMigrationAPI* | [**DeleteMigrationSource**](docs/DataMigrationAPI.md#deletemigrationsource) | **Delete** /storage/migration_sources/{id} | Delete a Migration Source
*DataMigrationAPI* | [**GetMigration**](docs/DataMigrationAPI.md#getmigration) | **Get** /storage/migrations/{id} | Get a Migration
*DataMigrationAPI* | [**GetMigrationSource**](docs/DataMigrationAPI.md#getmigrationsource) | **Get** /storage/migration_sources/{id} | Get a Migration Source
*DataMigrationAPI* | [**ListMigrationSourceCoverage**](docs/DataMigrationAPI.md#listmigrationsourcecoverage) | **Get** /storage/migration_source_coverage | List Migration Source coverage
*DataMigrationAPI* | [**ListMigrationSources**](docs/DataMigrationAPI.md#listmigrationsources) | **Get** /storage/migration_sources | List all Migration Sources
*DataMigrationAPI* | [**ListMigrations**](docs/DataMigrationAPI.md#listmigrations) | **Get** /storage/migrations | List all Migrations
*DataMigrationAPI* | [**StopMigration**](docs/DataMigrationAPI.md#stopmigration) | **Post** /storage/migrations/{id}/actions/stop | Stop a Migration
*DebuggingAPI* | [**ListCallEvents**](docs/DebuggingAPI.md#listcallevents) | **Get** /call_events | List call events
*DetailRecordsAPI* | [**SearchDetailRecords**](docs/DetailRecordsAPI.md#searchdetailrecords) | **Get** /detail_records | Search detail records
*DialogflowIntegrationAPI* | [**CreateDialogflowConnection**](docs/DialogflowIntegrationAPI.md#createdialogflowconnection) | **Post** /dialogflow_connections/{connection_id} | Create a Dialogflow Connection
*DialogflowIntegrationAPI* | [**DeleteDialogflowConnection**](docs/DialogflowIntegrationAPI.md#deletedialogflowconnection) | **Delete** /dialogflow_connections/{connection_id} | Delete stored Dialogflow Connection
*DialogflowIntegrationAPI* | [**GetDialogflowConnection**](docs/DialogflowIntegrationAPI.md#getdialogflowconnection) | **Get** /dialogflow_connections/{connection_id} | Retrieve stored Dialogflow Connection
*DialogflowIntegrationAPI* | [**UpdateDialogflowConnection**](docs/DialogflowIntegrationAPI.md#updatedialogflowconnection) | **Put** /dialogflow_connections/{connection_id} | Update stored Dialogflow Connection
*DocumentsAPI* | [**CreateDocument**](docs/DocumentsAPI.md#createdocument) | **Post** /documents | Upload a document
*DocumentsAPI* | [**DeleteDocument**](docs/DocumentsAPI.md#deletedocument) | **Delete** /documents/{id} | Delete a document
*DocumentsAPI* | [**DownloadDocument**](docs/DocumentsAPI.md#downloaddocument) | **Get** /documents/{id}/download | Download a document
*DocumentsAPI* | [**ListDocumentLinks**](docs/DocumentsAPI.md#listdocumentlinks) | **Get** /document_links | List all document links
*DocumentsAPI* | [**ListDocuments**](docs/DocumentsAPI.md#listdocuments) | **Get** /documents | List all documents
*DocumentsAPI* | [**RetrieveDocument**](docs/DocumentsAPI.md#retrievedocument) | **Get** /documents/{id} | Retrieve a document
*DocumentsAPI* | [**UpdateDocument**](docs/DocumentsAPI.md#updatedocument) | **Patch** /documents/{id} | Update a document
*DynamicEmergencyAddressesAPI* | [**CreateDynamicEmergencyAddress**](docs/DynamicEmergencyAddressesAPI.md#createdynamicemergencyaddress) | **Post** /dynamic_emergency_addresses | Create a dynamic emergency address.
*DynamicEmergencyAddressesAPI* | [**DeleteDynamicEmergencyAddress**](docs/DynamicEmergencyAddressesAPI.md#deletedynamicemergencyaddress) | **Delete** /dynamic_emergency_addresses/{id} | Delete a dynamic emergency address
*DynamicEmergencyAddressesAPI* | [**GetDynamicEmergencyAddress**](docs/DynamicEmergencyAddressesAPI.md#getdynamicemergencyaddress) | **Get** /dynamic_emergency_addresses/{id} | Get a dynamic emergency address
*DynamicEmergencyAddressesAPI* | [**ListDynamicEmergencyAddresses**](docs/DynamicEmergencyAddressesAPI.md#listdynamicemergencyaddresses) | **Get** /dynamic_emergency_addresses | List dynamic emergency addresses
*DynamicEmergencyEndpointsAPI* | [**CreateDynamicEmergencyEndpoint**](docs/DynamicEmergencyEndpointsAPI.md#createdynamicemergencyendpoint) | **Post** /dynamic_emergency_endpoints | Create a dynamic emergency endpoint.
*DynamicEmergencyEndpointsAPI* | [**DeleteDynamicEmergencyEndpoint**](docs/DynamicEmergencyEndpointsAPI.md#deletedynamicemergencyendpoint) | **Delete** /dynamic_emergency_endpoints/{id} | Delete a dynamic emergency endpoint
*DynamicEmergencyEndpointsAPI* | [**GetDynamicEmergencyEndpoint**](docs/DynamicEmergencyEndpointsAPI.md#getdynamicemergencyendpoint) | **Get** /dynamic_emergency_endpoints/{id} | Get a dynamic emergency endpoint
*DynamicEmergencyEndpointsAPI* | [**ListDynamicEmergencyEndpoints**](docs/DynamicEmergencyEndpointsAPI.md#listdynamicemergencyendpoints) | **Get** /dynamic_emergency_endpoints | List dynamic emergency endpoints
*EmbeddingsAPI* | [**EmbeddingBucketFilesPublicEmbeddingBucketsBucketNameDelete**](docs/EmbeddingsAPI.md#embeddingbucketfilespublicembeddingbucketsbucketnamedelete) | **Delete** /ai/embeddings/buckets/{bucket_name} | Disable AI for an Embedded Bucket
*EmbeddingsAPI* | [**GetBucketName**](docs/EmbeddingsAPI.md#getbucketname) | **Get** /ai/embeddings/buckets/{bucket_name} | Get file-level embedding statuses for a bucket
*EmbeddingsAPI* | [**GetEmbeddingBuckets**](docs/EmbeddingsAPI.md#getembeddingbuckets) | **Get** /ai/embeddings/buckets | List embedded buckets
*EmbeddingsAPI* | [**GetEmbeddingTask**](docs/EmbeddingsAPI.md#getembeddingtask) | **Get** /ai/embeddings/{task_id} | Get an embedding task&#39;s status
*EmbeddingsAPI* | [**GetTasksByStatus**](docs/EmbeddingsAPI.md#gettasksbystatus) | **Get** /ai/embeddings | Get Tasks by Status
*EmbeddingsAPI* | [**PostEmbedding**](docs/EmbeddingsAPI.md#postembedding) | **Post** /ai/embeddings | Embed documents
*EmbeddingsAPI* | [**PostEmbeddingSimilaritySearch**](docs/EmbeddingsAPI.md#postembeddingsimilaritysearch) | **Post** /ai/embeddings/similarity-search | Search for documents
*EmbeddingsAPI* | [**PostEmbeddingUrl**](docs/EmbeddingsAPI.md#postembeddingurl) | **Post** /ai/embeddings/url | Embed URL content
*EnumAPI* | [**GetEnumEndpoint**](docs/EnumAPI.md#getenumendpoint) | **Get** /enum/{endpoint} | Get Enum
*ExternalConnectionsAPI* | [**CreateExternalConnection**](docs/ExternalConnectionsAPI.md#createexternalconnection) | **Post** /external_connections | Creates an External Connection
*ExternalConnectionsAPI* | [**CreateExternalConnectionUpload**](docs/ExternalConnectionsAPI.md#createexternalconnectionupload) | **Post** /external_connections/{id}/uploads | Creates an Upload request
*ExternalConnectionsAPI* | [**DeleteExternalConnection**](docs/ExternalConnectionsAPI.md#deleteexternalconnection) | **Delete** /external_connections/{id} | Deletes an External Connection
*ExternalConnectionsAPI* | [**DeleteExternalConnectionLogMessage**](docs/ExternalConnectionsAPI.md#deleteexternalconnectionlogmessage) | **Delete** /external_connections/log_messages/{id} | Dismiss a log message
*ExternalConnectionsAPI* | [**GetExternalConnection**](docs/ExternalConnectionsAPI.md#getexternalconnection) | **Get** /external_connections/{id} | Retrieve an External Connection
*ExternalConnectionsAPI* | [**GetExternalConnectionCivicAddress**](docs/ExternalConnectionsAPI.md#getexternalconnectioncivicaddress) | **Get** /external_connections/{id}/civic_addresses/{address_id} | Retrieve a Civic Address
*ExternalConnectionsAPI* | [**GetExternalConnectionLogMessage**](docs/ExternalConnectionsAPI.md#getexternalconnectionlogmessage) | **Get** /external_connections/log_messages/{id} | Retrieve a log message
*ExternalConnectionsAPI* | [**GetExternalConnectionPhoneNumber**](docs/ExternalConnectionsAPI.md#getexternalconnectionphonenumber) | **Get** /external_connections/{id}/phone_numbers/{phone_number_id} | Retrieve a phone number
*ExternalConnectionsAPI* | [**GetExternalConnectionRelease**](docs/ExternalConnectionsAPI.md#getexternalconnectionrelease) | **Get** /external_connections/{id}/releases/{release_id} | Retrieve a Release request
*ExternalConnectionsAPI* | [**GetExternalConnectionUpload**](docs/ExternalConnectionsAPI.md#getexternalconnectionupload) | **Get** /external_connections/{id}/uploads/{ticket_id} | Retrieve an Upload request
*ExternalConnectionsAPI* | [**GetExternalConnectionUploadsStatus**](docs/ExternalConnectionsAPI.md#getexternalconnectionuploadsstatus) | **Get** /external_connections/{id}/uploads/status | Get the count of pending upload requests
*ExternalConnectionsAPI* | [**ListCivicAddresses**](docs/ExternalConnectionsAPI.md#listcivicaddresses) | **Get** /external_connections/{id}/civic_addresses | List all civic addresses and locations
*ExternalConnectionsAPI* | [**ListExternalConnectionLogMessages**](docs/ExternalConnectionsAPI.md#listexternalconnectionlogmessages) | **Get** /external_connections/log_messages | List all log messages
*ExternalConnectionsAPI* | [**ListExternalConnectionPhoneNumbers**](docs/ExternalConnectionsAPI.md#listexternalconnectionphonenumbers) | **Get** /external_connections/{id}/phone_numbers | List all phone numbers
*ExternalConnectionsAPI* | [**ListExternalConnectionReleases**](docs/ExternalConnectionsAPI.md#listexternalconnectionreleases) | **Get** /external_connections/{id}/releases | List all Releases
*ExternalConnectionsAPI* | [**ListExternalConnectionUploads**](docs/ExternalConnectionsAPI.md#listexternalconnectionuploads) | **Get** /external_connections/{id}/uploads | List all Upload requests
*ExternalConnectionsAPI* | [**ListExternalConnections**](docs/ExternalConnectionsAPI.md#listexternalconnections) | **Get** /external_connections | List all External Connections
*ExternalConnectionsAPI* | [**OperatorConnectRefresh**](docs/ExternalConnectionsAPI.md#operatorconnectrefresh) | **Post** /operator_connect/actions/refresh | Refresh Operator Connect integration
*ExternalConnectionsAPI* | [**RefreshExternalConnectionUploads**](docs/ExternalConnectionsAPI.md#refreshexternalconnectionuploads) | **Post** /external_connections/{id}/uploads/refresh | Refresh the status of all Upload requests
*ExternalConnectionsAPI* | [**RetryUpload**](docs/ExternalConnectionsAPI.md#retryupload) | **Post** /external_connections/{id}/uploads/{ticket_id}/retry | Retry an Upload request
*ExternalConnectionsAPI* | [**UpdateExternalConnection**](docs/ExternalConnectionsAPI.md#updateexternalconnection) | **Patch** /external_connections/{id} | Update an External Connection
*ExternalConnectionsAPI* | [**UpdateExternalConnectionPhoneNumber**](docs/ExternalConnectionsAPI.md#updateexternalconnectionphonenumber) | **Patch** /external_connections/{id}/phone_numbers/{phone_number_id} | Update a phone number
*ExternalConnectionsAPI* | [**UpdateLocation**](docs/ExternalConnectionsAPI.md#updatelocation) | **Patch** /external_connections/{id}/locations/{location_id} | Update a location&#39;s static emergency address
*FQDNConnectionsAPI* | [**CreateFqdnConnection**](docs/FQDNConnectionsAPI.md#createfqdnconnection) | **Post** /fqdn_connections | Create an FQDN connection
*FQDNConnectionsAPI* | [**DeleteFqdnConnection**](docs/FQDNConnectionsAPI.md#deletefqdnconnection) | **Delete** /fqdn_connections/{id} | Delete an FQDN connection
*FQDNConnectionsAPI* | [**ListFqdnConnections**](docs/FQDNConnectionsAPI.md#listfqdnconnections) | **Get** /fqdn_connections | List FQDN connections
*FQDNConnectionsAPI* | [**RetrieveFqdnConnection**](docs/FQDNConnectionsAPI.md#retrievefqdnconnection) | **Get** /fqdn_connections/{id} | Retrieve an FQDN connection
*FQDNConnectionsAPI* | [**UpdateFqdnConnection**](docs/FQDNConnectionsAPI.md#updatefqdnconnection) | **Patch** /fqdn_connections/{id} | Update an FQDN connection
*FQDNsAPI* | [**CreateFqdn**](docs/FQDNsAPI.md#createfqdn) | **Post** /fqdns | Create an FQDN
*FQDNsAPI* | [**DeleteFqdn**](docs/FQDNsAPI.md#deletefqdn) | **Delete** /fqdns/{id} | Delete an FQDN
*FQDNsAPI* | [**ListFqdns**](docs/FQDNsAPI.md#listfqdns) | **Get** /fqdns | List FQDNs
*FQDNsAPI* | [**RetrieveFqdn**](docs/FQDNsAPI.md#retrievefqdn) | **Get** /fqdns/{id} | Retrieve an FQDN
*FQDNsAPI* | [**UpdateFqdn**](docs/FQDNsAPI.md#updatefqdn) | **Patch** /fqdns/{id} | Update an FQDN
*FineTuningAPI* | [**CancelNewFinetuningjobPublicFinetuningPost**](docs/FineTuningAPI.md#cancelnewfinetuningjobpublicfinetuningpost) | **Post** /ai/fine_tuning/jobs/{job_id}/cancel | Cancel a fine tuning job
*FineTuningAPI* | [**CreateNewFinetuningjobPublicFinetuningPost**](docs/FineTuningAPI.md#createnewfinetuningjobpublicfinetuningpost) | **Post** /ai/fine_tuning/jobs | Create a fine tuning job
*FineTuningAPI* | [**GetFinetuningjobPublicFinetuningGet**](docs/FineTuningAPI.md#getfinetuningjobpublicfinetuningget) | **Get** /ai/fine_tuning/jobs | List fine tuning jobs
*FineTuningAPI* | [**GetFinetuningjobPublicFinetuningJobIdGet**](docs/FineTuningAPI.md#getfinetuningjobpublicfinetuningjobidget) | **Get** /ai/fine_tuning/jobs/{job_id} | Get a fine tuning job
*GlobalIPsAPI* | [**CreateGlobalIp**](docs/GlobalIPsAPI.md#createglobalip) | **Post** /global_ips | Create a Global IP
*GlobalIPsAPI* | [**CreateGlobalIpAssignment**](docs/GlobalIPsAPI.md#createglobalipassignment) | **Post** /global_ip_assignments | Create a Global IP assignment
*GlobalIPsAPI* | [**CreateGlobalIpHealthCheck**](docs/GlobalIPsAPI.md#createglobaliphealthcheck) | **Post** /global_ip_health_checks | Create a Global IP health check
*GlobalIPsAPI* | [**DeleteGlobalIp**](docs/GlobalIPsAPI.md#deleteglobalip) | **Delete** /global_ips/{id} | Delete a Global IP
*GlobalIPsAPI* | [**DeleteGlobalIpAssignment**](docs/GlobalIPsAPI.md#deleteglobalipassignment) | **Delete** /global_ip_assignments/{id} | Delete a Global IP assignment
*GlobalIPsAPI* | [**DeleteGlobalIpHealthCheck**](docs/GlobalIPsAPI.md#deleteglobaliphealthcheck) | **Delete** /global_ip_health_checks/{id} | Delete a Global IP health check
*GlobalIPsAPI* | [**GetGlobalIp**](docs/GlobalIPsAPI.md#getglobalip) | **Get** /global_ips/{id} | Retrieve a Global IP
*GlobalIPsAPI* | [**GetGlobalIpAssignment**](docs/GlobalIPsAPI.md#getglobalipassignment) | **Get** /global_ip_assignments/{id} | Retrieve a Global IP
*GlobalIPsAPI* | [**GetGlobalIpAssignmentHealth**](docs/GlobalIPsAPI.md#getglobalipassignmenthealth) | **Get** /global_ip_assignment_health | Global IP Assignment Health Check Metrics
*GlobalIPsAPI* | [**GetGlobalIpAssignmentUsage**](docs/GlobalIPsAPI.md#getglobalipassignmentusage) | **Get** /global_ip_assignments_usage | Global IP Assignment Usage Metrics
*GlobalIPsAPI* | [**GetGlobalIpHealthCheck**](docs/GlobalIPsAPI.md#getglobaliphealthcheck) | **Get** /global_ip_health_checks/{id} | Retrieve a Global IP health check
*GlobalIPsAPI* | [**GetGlobalIpLatency**](docs/GlobalIPsAPI.md#getglobaliplatency) | **Get** /global_ip_latency | Global IP Latency Metrics
*GlobalIPsAPI* | [**GetGlobalIpUsage**](docs/GlobalIPsAPI.md#getglobalipusage) | **Get** /global_ip_usage | Global IP Usage Metrics
*GlobalIPsAPI* | [**ListGlobalIpAllowedPorts**](docs/GlobalIPsAPI.md#listglobalipallowedports) | **Get** /global_ip_allowed_ports | List all Global IP Allowed Ports
*GlobalIPsAPI* | [**ListGlobalIpAssignments**](docs/GlobalIPsAPI.md#listglobalipassignments) | **Get** /global_ip_assignments | List all Global IP assignments
*GlobalIPsAPI* | [**ListGlobalIpHealthCheckTypes**](docs/GlobalIPsAPI.md#listglobaliphealthchecktypes) | **Get** /global_ip_health_check_types | List all Global IP Health check types
*GlobalIPsAPI* | [**ListGlobalIpHealthChecks**](docs/GlobalIPsAPI.md#listglobaliphealthchecks) | **Get** /global_ip_health_checks | List all Global IP health checks
*GlobalIPsAPI* | [**ListGlobalIpProtocols**](docs/GlobalIPsAPI.md#listglobalipprotocols) | **Get** /global_ip_protocols | List all Global IP Protocols
*GlobalIPsAPI* | [**ListGlobalIps**](docs/GlobalIPsAPI.md#listglobalips) | **Get** /global_ips | List all Global IPs
*GlobalIPsAPI* | [**UpdateGlobalIpAssignment**](docs/GlobalIPsAPI.md#updateglobalipassignment) | **Patch** /global_ip_assignments/{id} | Update a Global IP assignment
*IPAddressesAPI* | [**CreateAccessIpAddress**](docs/IPAddressesAPI.md#createaccessipaddress) | **Post** /access_ip_address | Create new Access IP Address
*IPAddressesAPI* | [**DeleteAccessIpAddress**](docs/IPAddressesAPI.md#deleteaccessipaddress) | **Delete** /access_ip_address/{access_ip_address_id} | Delete access IP address
*IPAddressesAPI* | [**GetAccessIpAddress**](docs/IPAddressesAPI.md#getaccessipaddress) | **Get** /access_ip_address/{access_ip_address_id} | Retrieve an access IP address
*IPAddressesAPI* | [**ListAccessIpAddresses**](docs/IPAddressesAPI.md#listaccessipaddresses) | **Get** /access_ip_address | List all Access IP Addresses
*IPConnectionsAPI* | [**CreateIpConnection**](docs/IPConnectionsAPI.md#createipconnection) | **Post** /ip_connections | Create an Ip connection
*IPConnectionsAPI* | [**DeleteIpConnection**](docs/IPConnectionsAPI.md#deleteipconnection) | **Delete** /ip_connections/{id} | Delete an Ip connection
*IPConnectionsAPI* | [**ListIpConnections**](docs/IPConnectionsAPI.md#listipconnections) | **Get** /ip_connections | List Ip connections
*IPConnectionsAPI* | [**RetrieveIpConnection**](docs/IPConnectionsAPI.md#retrieveipconnection) | **Get** /ip_connections/{id} | Retrieve an Ip connection
*IPConnectionsAPI* | [**UpdateIpConnection**](docs/IPConnectionsAPI.md#updateipconnection) | **Patch** /ip_connections/{id} | Update an Ip connection
*IPRangesAPI* | [**AccessIpRangesAccessIpRangeIdDelete**](docs/IPRangesAPI.md#accessiprangesaccessiprangeiddelete) | **Delete** /access_ip_ranges/{access_ip_range_id} | Delete access IP ranges
*IPRangesAPI* | [**CreateAccessIPRange**](docs/IPRangesAPI.md#createaccessiprange) | **Post** /access_ip_ranges | Create new Access IP Range
*IPRangesAPI* | [**ListAccessIpRanges**](docs/IPRangesAPI.md#listaccessipranges) | **Get** /access_ip_ranges | List all Access IP Ranges
*IPsAPI* | [**CreateIp**](docs/IPsAPI.md#createip) | **Post** /ips | Create an Ip
*IPsAPI* | [**DeleteIp**](docs/IPsAPI.md#deleteip) | **Delete** /ips/{id} | Delete an Ip
*IPsAPI* | [**ListIps**](docs/IPsAPI.md#listips) | **Get** /ips | List Ips
*IPsAPI* | [**RetrieveIp**](docs/IPsAPI.md#retrieveip) | **Get** /ips/{id} | Retrieve an Ip
*IPsAPI* | [**UpdateIp**](docs/IPsAPI.md#updateip) | **Patch** /ips/{id} | Update an Ip
*IntegrationSecretsAPI* | [**CreateIntegrationSecret**](docs/IntegrationSecretsAPI.md#createintegrationsecret) | **Post** /integration_secrets | Create a secret
*IntegrationSecretsAPI* | [**DeleteIntegrationSecret**](docs/IntegrationSecretsAPI.md#deleteintegrationsecret) | **Delete** /integration_secrets/{id} | Delete an integration secret
*IntegrationSecretsAPI* | [**ListIntegrationSecrets**](docs/IntegrationSecretsAPI.md#listintegrationsecrets) | **Get** /integration_secrets | List integration secrets
*InventoryLevelAPI* | [**CreateInventoryCoverage**](docs/InventoryLevelAPI.md#createinventorycoverage) | **Get** /inventory_coverage | Create an inventory coverage request
*MDRDetailReportsAPI* | [**GetPaginatedMdrs**](docs/MDRDetailReportsAPI.md#getpaginatedmdrs) | **Get** /reports/mdrs | Fetch all Mdr records
*MDRUsageReportsAPI* | [**DeleteUsageReport**](docs/MDRUsageReportsAPI.md#deleteusagereport) | **Delete** /reports/mdr_usage_reports/{id} | Delete MDR Usage Report
*MDRUsageReportsAPI* | [**GetMDRUsageReportSync**](docs/MDRUsageReportsAPI.md#getmdrusagereportsync) | **Get** /reports/mdr_usage_reports/sync | Generate and fetch MDR Usage Report
*MDRUsageReportsAPI* | [**GetMdrUsageReports**](docs/MDRUsageReportsAPI.md#getmdrusagereports) | **Get** /reports/mdr_usage_reports | Fetch all Messaging usage reports
*MDRUsageReportsAPI* | [**GetUsageReport**](docs/MDRUsageReportsAPI.md#getusagereport) | **Get** /reports/mdr_usage_reports/{id} | Retrieve messaging report
*MDRUsageReportsAPI* | [**SubmitUsageReport**](docs/MDRUsageReportsAPI.md#submitusagereport) | **Post** /reports/mdr_usage_reports | Create MDR Usage Report
*ManagedAccountsAPI* | [**CreateManagedAccount**](docs/ManagedAccountsAPI.md#createmanagedaccount) | **Post** /managed_accounts | Create a new managed account.
*ManagedAccountsAPI* | [**DisableManagedAccount**](docs/ManagedAccountsAPI.md#disablemanagedaccount) | **Post** /managed_accounts/{id}/actions/disable | Disables a managed account
*ManagedAccountsAPI* | [**EnableManagedAccount**](docs/ManagedAccountsAPI.md#enablemanagedaccount) | **Post** /managed_accounts/{id}/actions/enable | Enables a managed account
*ManagedAccountsAPI* | [**ListAllocatableGlobalOutboundChannels**](docs/ManagedAccountsAPI.md#listallocatableglobaloutboundchannels) | **Get** /managed_accounts/allocatable_global_outbound_channels | Display information about allocatable global outbound channels for the current user.
*ManagedAccountsAPI* | [**ListManagedAccounts**](docs/ManagedAccountsAPI.md#listmanagedaccounts) | **Get** /managed_accounts | Lists accounts managed by the current user.
*ManagedAccountsAPI* | [**RetrieveManagedAccount**](docs/ManagedAccountsAPI.md#retrievemanagedaccount) | **Get** /managed_accounts/{id} | Retrieve a managed account
*ManagedAccountsAPI* | [**UpdateManagedAccount**](docs/ManagedAccountsAPI.md#updatemanagedaccount) | **Patch** /managed_accounts/{id} | Update a managed account
*ManagedAccountsAPI* | [**UpdateManagedAccountGlobalChannelLimit**](docs/ManagedAccountsAPI.md#updatemanagedaccountglobalchannellimit) | **Patch** /managed_accounts/{id}/update_global_channel_limit | Update the amount of allocatable global outbound channels allocated to a specific managed account.
*MediaStorageAPIAPI* | [**CreateMediaStorage**](docs/MediaStorageAPIAPI.md#createmediastorage) | **Post** /media | Upload media
*MediaStorageAPIAPI* | [**DeleteMediaStorage**](docs/MediaStorageAPIAPI.md#deletemediastorage) | **Delete** /media/{media_name} | Deletes stored media
*MediaStorageAPIAPI* | [**DownloadMedia**](docs/MediaStorageAPIAPI.md#downloadmedia) | **Get** /media/{media_name}/download | Download stored media
*MediaStorageAPIAPI* | [**GetMediaStorage**](docs/MediaStorageAPIAPI.md#getmediastorage) | **Get** /media/{media_name} | Retrieve stored media
*MediaStorageAPIAPI* | [**ListMediaStorage**](docs/MediaStorageAPIAPI.md#listmediastorage) | **Get** /media | List uploaded media
*MediaStorageAPIAPI* | [**UpdateMediaStorage**](docs/MediaStorageAPIAPI.md#updatemediastorage) | **Put** /media/{media_name} | Update stored media
*MessagesAPI* | [**CreateGroupMmsMessage**](docs/MessagesAPI.md#creategroupmmsmessage) | **Post** /messages/group_mms | Send a group MMS message
*MessagesAPI* | [**CreateLongCodeMessage**](docs/MessagesAPI.md#createlongcodemessage) | **Post** /messages/long_code | Send a long code message
*MessagesAPI* | [**CreateNumberPoolMessage**](docs/MessagesAPI.md#createnumberpoolmessage) | **Post** /messages/number_pool | Send a message using number pool
*MessagesAPI* | [**CreateShortCodeMessage**](docs/MessagesAPI.md#createshortcodemessage) | **Post** /messages/short_code | Send a short code message
*MessagesAPI* | [**GetMessage**](docs/MessagesAPI.md#getmessage) | **Get** /messages/{id} | Retrieve a message
*MessagesAPI* | [**SendMessage**](docs/MessagesAPI.md#sendmessage) | **Post** /messages | Send a message
*MessagingHostedNumberAPI* | [**CheckEligibilityNumbers**](docs/MessagingHostedNumberAPI.md#checkeligibilitynumbers) | **Get** /messaging_hosted_number_orders/eligibility_numbers_check | Check eligibility of phone numbers for hosted messaging
*MessagingHostedNumberAPI* | [**CreateMessagingHostedNumberOrder**](docs/MessagingHostedNumberAPI.md#createmessaginghostednumberorder) | **Post** /messaging_hosted_number_orders | Create a messaging hosted number order
*MessagingHostedNumberAPI* | [**CreateVerificationCodesForMessagingHostedNumberOrder**](docs/MessagingHostedNumberAPI.md#createverificationcodesformessaginghostednumberorder) | **Post** /messaging_hosted_number_orders/{id}/verification_codes | Create verification codes for the hosted numbers order
*MessagingHostedNumberAPI* | [**DeleteMessagingHostedNumber**](docs/MessagingHostedNumberAPI.md#deletemessaginghostednumber) | **Delete** /messaging_hosted_numbers/{id} | Delete a messaging hosted number
*MessagingHostedNumberAPI* | [**GetMessagingHostedNumberOrder**](docs/MessagingHostedNumberAPI.md#getmessaginghostednumberorder) | **Get** /messaging_hosted_number_orders/{id} | Retrieve a messaging hosted number order
*MessagingHostedNumberAPI* | [**ListMessagingHostedNumberOrders**](docs/MessagingHostedNumberAPI.md#listmessaginghostednumberorders) | **Get** /messaging_hosted_number_orders | List messaging hosted number orders
*MessagingHostedNumberAPI* | [**UploadMessagingHostedNumberOrderFile**](docs/MessagingHostedNumberAPI.md#uploadmessaginghostednumberorderfile) | **Post** /messaging_hosted_number_orders/{id}/actions/file_upload | Upload file required for a messaging hosted number order
*MessagingHostedNumberAPI* | [**ValidateVerificationCodesForMessagingHostedNumberOrder**](docs/MessagingHostedNumberAPI.md#validateverificationcodesformessaginghostednumberorder) | **Post** /messaging_hosted_number_orders/{id}/validation_codes | Validate the verification codes for the hosted numbers order
*MessagingProfilesAPI* | [**CreateMessagingProfile**](docs/MessagingProfilesAPI.md#createmessagingprofile) | **Post** /messaging_profiles | Create a messaging profile
*MessagingProfilesAPI* | [**DeleteMessagingProfile**](docs/MessagingProfilesAPI.md#deletemessagingprofile) | **Delete** /messaging_profiles/{id} | Delete a messaging profile
*MessagingProfilesAPI* | [**GetMessagingProfileMetrics**](docs/MessagingProfilesAPI.md#getmessagingprofilemetrics) | **Get** /messaging_profiles/{id}/metrics | Retrieve messaging profile metrics
*MessagingProfilesAPI* | [**ListMessagingProfiles**](docs/MessagingProfilesAPI.md#listmessagingprofiles) | **Get** /messaging_profiles | List messaging profiles
*MessagingProfilesAPI* | [**ListProfileMetrics**](docs/MessagingProfilesAPI.md#listprofilemetrics) | **Get** /messaging_profile_metrics | List messaging profile metrics
*MessagingProfilesAPI* | [**ListProfilePhoneNumbers**](docs/MessagingProfilesAPI.md#listprofilephonenumbers) | **Get** /messaging_profiles/{id}/phone_numbers | List phone numbers associated with a messaging profile
*MessagingProfilesAPI* | [**ListProfileShortCodes**](docs/MessagingProfilesAPI.md#listprofileshortcodes) | **Get** /messaging_profiles/{id}/short_codes | List short codes associated with a messaging profile
*MessagingProfilesAPI* | [**RetrieveMessagingProfile**](docs/MessagingProfilesAPI.md#retrievemessagingprofile) | **Get** /messaging_profiles/{id} | Retrieve a messaging profile
*MessagingProfilesAPI* | [**UpdateMessagingProfile**](docs/MessagingProfilesAPI.md#updatemessagingprofile) | **Patch** /messaging_profiles/{id} | Update a messaging profile
*MessagingTollfreeVerificationAPI* | [**DeleteVerificationRequest**](docs/MessagingTollfreeVerificationAPI.md#deleteverificationrequest) | **Delete** /messaging_tollfree/verification/requests/{id} | Delete Verification Request
*MessagingTollfreeVerificationAPI* | [**GetVerificationRequest**](docs/MessagingTollfreeVerificationAPI.md#getverificationrequest) | **Get** /messaging_tollfree/verification/requests/{id} | Get Verification Request
*MessagingTollfreeVerificationAPI* | [**ListVerificationRequests**](docs/MessagingTollfreeVerificationAPI.md#listverificationrequests) | **Get** /messaging_tollfree/verification/requests | List Verification Requests
*MessagingTollfreeVerificationAPI* | [**SubmitVerificationRequest**](docs/MessagingTollfreeVerificationAPI.md#submitverificationrequest) | **Post** /messaging_tollfree/verification/requests | Submit Verification Request
*MessagingTollfreeVerificationAPI* | [**UpdateVerificationRequest**](docs/MessagingTollfreeVerificationAPI.md#updateverificationrequest) | **Patch** /messaging_tollfree/verification/requests/{id} | Update Verification Request
*MessagingURLDomainsAPI* | [**ListMessagingUrlDomains**](docs/MessagingURLDomainsAPI.md#listmessagingurldomains) | **Get** /messaging_url_domains | List messaging URL domains
*MobileNetworkOperatorsAPI* | [**GetMobileNetworkOperators**](docs/MobileNetworkOperatorsAPI.md#getmobilenetworkoperators) | **Get** /mobile_network_operators | List mobile network operators
*NetworksAPI* | [**CreateDefaultGateway**](docs/NetworksAPI.md#createdefaultgateway) | **Post** /networks/{id}/default_gateway | Create Default Gateway.
*NetworksAPI* | [**CreateNetwork**](docs/NetworksAPI.md#createnetwork) | **Post** /networks | Create a Network
*NetworksAPI* | [**DeleteDefaultGateway**](docs/NetworksAPI.md#deletedefaultgateway) | **Delete** /networks/{id}/default_gateway | Delete Default Gateway.
*NetworksAPI* | [**DeleteNetwork**](docs/NetworksAPI.md#deletenetwork) | **Delete** /networks/{id} | Delete a Network
*NetworksAPI* | [**GetDefaultGateway**](docs/NetworksAPI.md#getdefaultgateway) | **Get** /networks/{id}/default_gateway | Get Default Gateway status.
*NetworksAPI* | [**GetNetwork**](docs/NetworksAPI.md#getnetwork) | **Get** /networks/{id} | Retrieve a Network
*NetworksAPI* | [**ListNetworkInterfaces**](docs/NetworksAPI.md#listnetworkinterfaces) | **Get** /networks/{id}/network_interfaces | List all Interfaces for a Network.
*NetworksAPI* | [**ListNetworks**](docs/NetworksAPI.md#listnetworks) | **Get** /networks | List all Networks
*NetworksAPI* | [**UpdateNetwork**](docs/NetworksAPI.md#updatenetwork) | **Patch** /networks/{id} | Update a Network
*NotificationsAPI* | [**CreateNotificationChannels**](docs/NotificationsAPI.md#createnotificationchannels) | **Post** /notification_channels | Create a notification channel
*NotificationsAPI* | [**CreateNotificationProfile**](docs/NotificationsAPI.md#createnotificationprofile) | **Post** /notification_profiles | Create a notification profile
*NotificationsAPI* | [**CreateNotificationSetting**](docs/NotificationsAPI.md#createnotificationsetting) | **Post** /notification_settings | Add a Notification Setting
*NotificationsAPI* | [**DeleteNotificationChannel**](docs/NotificationsAPI.md#deletenotificationchannel) | **Delete** /notification_channels/{id} | Delete a notification channel
*NotificationsAPI* | [**DeleteNotificationProfile**](docs/NotificationsAPI.md#deletenotificationprofile) | **Delete** /notification_profiles/{id} | Delete a notification profile
*NotificationsAPI* | [**DeleteNotificationSetting**](docs/NotificationsAPI.md#deletenotificationsetting) | **Delete** /notification_settings/{id} | Delete a notification setting
*NotificationsAPI* | [**FindNotificationsEvents**](docs/NotificationsAPI.md#findnotificationsevents) | **Get** /notification_events | List all Notifications Events
*NotificationsAPI* | [**FindNotificationsEventsConditions**](docs/NotificationsAPI.md#findnotificationseventsconditions) | **Get** /notification_event_conditions | List all Notifications Events Conditions
*NotificationsAPI* | [**FindNotificationsProfiles**](docs/NotificationsAPI.md#findnotificationsprofiles) | **Get** /notification_profiles | List all Notifications Profiles
*NotificationsAPI* | [**GetNotificationChannel**](docs/NotificationsAPI.md#getnotificationchannel) | **Get** /notification_channels/{id} | Get a notification channel
*NotificationsAPI* | [**GetNotificationProfile**](docs/NotificationsAPI.md#getnotificationprofile) | **Get** /notification_profiles/{id} | Get a notification profile
*NotificationsAPI* | [**GetNotificationSetting**](docs/NotificationsAPI.md#getnotificationsetting) | **Get** /notification_settings/{id} | Get a notification setting
*NotificationsAPI* | [**ListNotificationChannels**](docs/NotificationsAPI.md#listnotificationchannels) | **Get** /notification_channels | List notification channels
*NotificationsAPI* | [**ListNotificationSettings**](docs/NotificationsAPI.md#listnotificationsettings) | **Get** /notification_settings | List notification settings
*NotificationsAPI* | [**UpdateNotificationChannel**](docs/NotificationsAPI.md#updatenotificationchannel) | **Patch** /notification_channels/{id} | Update a notification channel
*NotificationsAPI* | [**UpdateNotificationProfile**](docs/NotificationsAPI.md#updatenotificationprofile) | **Patch** /notification_profiles/{id} | Update a notification profile
*NumberConfigurationsAPI* | [**BulkUpdateMessagingSettingsOnPhoneNumbers**](docs/NumberConfigurationsAPI.md#bulkupdatemessagingsettingsonphonenumbers) | **Post** /messaging_numbers_bulk_updates | Update the messaging profile of multiple phone numbers
*NumberConfigurationsAPI* | [**GetBulkUpdateMessagingSettingsOnPhoneNumbersStatus**](docs/NumberConfigurationsAPI.md#getbulkupdatemessagingsettingsonphonenumbersstatus) | **Get** /messaging_numbers_bulk_updates/{order_id} | Retrieve bulk update status
*NumberConfigurationsAPI* | [**GetPhoneNumberMessagingSettings**](docs/NumberConfigurationsAPI.md#getphonenumbermessagingsettings) | **Get** /phone_numbers/{id}/messaging | Retrieve a phone number with messaging settings
*NumberConfigurationsAPI* | [**ListPhoneNumbersWithMessagingSettings**](docs/NumberConfigurationsAPI.md#listphonenumberswithmessagingsettings) | **Get** /phone_numbers/messaging | List phone numbers with messaging settings
*NumberConfigurationsAPI* | [**UpdatePhoneNumberMessagingSettings**](docs/NumberConfigurationsAPI.md#updatephonenumbermessagingsettings) | **Patch** /phone_numbers/{id}/messaging | Update the messaging profile and/or messaging product of a phone number
*NumberLookupAPI* | [**LookupNumber**](docs/NumberLookupAPI.md#lookupnumber) | **Get** /number_lookup/{phone_number} | Lookup phone number data
*NumberPortoutAPI* | [**CreatePortoutReport**](docs/NumberPortoutAPI.md#createportoutreport) | **Post** /portouts/reports | Create a port-out related report
*NumberPortoutAPI* | [**FindPortoutComments**](docs/NumberPortoutAPI.md#findportoutcomments) | **Get** /portouts/{id}/comments | List all comments for a portout request
*NumberPortoutAPI* | [**FindPortoutRequest**](docs/NumberPortoutAPI.md#findportoutrequest) | **Get** /portouts/{id} | Get a portout request
*NumberPortoutAPI* | [**GetPortRequestSupportingDocuments**](docs/NumberPortoutAPI.md#getportrequestsupportingdocuments) | **Get** /portouts/{id}/supporting_documents | List supporting documents on a portout request
*NumberPortoutAPI* | [**GetPortoutReport**](docs/NumberPortoutAPI.md#getportoutreport) | **Get** /portouts/reports/{id} | Retrieve a report
*NumberPortoutAPI* | [**ListPortoutEvents**](docs/NumberPortoutAPI.md#listportoutevents) | **Get** /portouts/events | List all port-out events
*NumberPortoutAPI* | [**ListPortoutRejections**](docs/NumberPortoutAPI.md#listportoutrejections) | **Get** /portouts/rejections/{portout_id} | List eligible port-out rejection codes for a specific order
*NumberPortoutAPI* | [**ListPortoutReports**](docs/NumberPortoutAPI.md#listportoutreports) | **Get** /portouts/reports | List port-out related reports
*NumberPortoutAPI* | [**ListPortoutRequest**](docs/NumberPortoutAPI.md#listportoutrequest) | **Get** /portouts | List portout requests
*NumberPortoutAPI* | [**PostPortRequestComment**](docs/NumberPortoutAPI.md#postportrequestcomment) | **Post** /portouts/{id}/comments | Create a comment on a portout request
*NumberPortoutAPI* | [**PostPortRequestSupportingDocuments**](docs/NumberPortoutAPI.md#postportrequestsupportingdocuments) | **Post** /portouts/{id}/supporting_documents | Create a list of supporting documents on a portout request
*NumberPortoutAPI* | [**RepublishPortoutEvent**](docs/NumberPortoutAPI.md#republishportoutevent) | **Post** /portouts/events/{id}/republish | Republish a port-out event
*NumberPortoutAPI* | [**ShowPortoutEvent**](docs/NumberPortoutAPI.md#showportoutevent) | **Get** /portouts/events/{id} | Show a port-out event
*NumberPortoutAPI* | [**UpdatePortoutStatus**](docs/NumberPortoutAPI.md#updateportoutstatus) | **Patch** /portouts/{id}/{status} | Update Status
*NumbersFeaturesAPI* | [**PostNumbersFeatures**](docs/NumbersFeaturesAPI.md#postnumbersfeatures) | **Post** /numbers_features | Retrieve the features for a list of numbers
*OTAUpdatesAPI* | [**GetOtaUpdate**](docs/OTAUpdatesAPI.md#getotaupdate) | **Get** /ota_updates/{id} | Get OTA update
*OTAUpdatesAPI* | [**ListOtaUpdates**](docs/OTAUpdatesAPI.md#listotaupdates) | **Get** /ota_updates | List OTA updates
*ObjectAPI* | [**DeleteObject**](docs/ObjectAPI.md#deleteobject) | **Delete** /{bucketName}/{objectName} | DeleteObject
*ObjectAPI* | [**DeleteObjects**](docs/ObjectAPI.md#deleteobjects) | **Post** /{bucketName} | DeleteObjects
*ObjectAPI* | [**GetObject**](docs/ObjectAPI.md#getobject) | **Get** /{bucketName}/{objectName} | GetObject
*ObjectAPI* | [**HeadObject**](docs/ObjectAPI.md#headobject) | **Head** /{bucketName}/{objectName} | HeadObject
*ObjectAPI* | [**ListObjects**](docs/ObjectAPI.md#listobjects) | **Get** /{bucketName} | ListObjectsV2
*ObjectAPI* | [**PutObject**](docs/ObjectAPI.md#putobject) | **Put** /{bucketName}/{objectName} | PutObject
*OutboundVoiceProfilesAPI* | [**CreateVoiceProfile**](docs/OutboundVoiceProfilesAPI.md#createvoiceprofile) | **Post** /outbound_voice_profiles | Create an outbound voice profile
*OutboundVoiceProfilesAPI* | [**DeleteOutboundVoiceProfile**](docs/OutboundVoiceProfilesAPI.md#deleteoutboundvoiceprofile) | **Delete** /outbound_voice_profiles/{id} | Delete an outbound voice profile
*OutboundVoiceProfilesAPI* | [**GetOutboundVoiceProfile**](docs/OutboundVoiceProfilesAPI.md#getoutboundvoiceprofile) | **Get** /outbound_voice_profiles/{id} | Retrieve an outbound voice profile
*OutboundVoiceProfilesAPI* | [**ListOutboundVoiceProfiles**](docs/OutboundVoiceProfilesAPI.md#listoutboundvoiceprofiles) | **Get** /outbound_voice_profiles | Get all outbound voice profiles
*OutboundVoiceProfilesAPI* | [**UpdateOutboundVoiceProfile**](docs/OutboundVoiceProfilesAPI.md#updateoutboundvoiceprofile) | **Patch** /outbound_voice_profiles/{id} | Updates an existing outbound voice profile.
*PhoneNumberBlockOrdersAPI* | [**CreateNumberBlockOrder**](docs/PhoneNumberBlockOrdersAPI.md#createnumberblockorder) | **Post** /number_block_orders | Create a number block order
*PhoneNumberBlockOrdersAPI* | [**ListNumberBlockOrders**](docs/PhoneNumberBlockOrdersAPI.md#listnumberblockorders) | **Get** /number_block_orders | List number block orders
*PhoneNumberBlockOrdersAPI* | [**RetrieveNumberBlockOrder**](docs/PhoneNumberBlockOrdersAPI.md#retrievenumberblockorder) | **Get** /number_block_orders/{number_block_order_id} | Retrieve a number block order
*PhoneNumberBlocksBackgroundJobsAPI* | [**CreatePhoneNumberBlockDeletionJob**](docs/PhoneNumberBlocksBackgroundJobsAPI.md#createphonenumberblockdeletionjob) | **Post** /phone_number_blocks/jobs/delete_phone_number_block | Deletes all numbers associated with a phone number block
*PhoneNumberBlocksBackgroundJobsAPI* | [**GetPhoneNumberBlocksJob**](docs/PhoneNumberBlocksBackgroundJobsAPI.md#getphonenumberblocksjob) | **Get** /phone_number_blocks/jobs/{id} | Retrieves a phone number blocks job
*PhoneNumberBlocksBackgroundJobsAPI* | [**ListPhoneNumberBlocksJobs**](docs/PhoneNumberBlocksBackgroundJobsAPI.md#listphonenumberblocksjobs) | **Get** /phone_number_blocks/jobs | Lists the phone number blocks jobs
*PhoneNumberCampaignsAPI* | [**CreatePhoneNumberCampaign**](docs/PhoneNumberCampaignsAPI.md#createphonenumbercampaign) | **Post** /phone_number_campaigns | Create New Phone Number Campaign
*PhoneNumberCampaignsAPI* | [**DeletePhoneNumberCampaign**](docs/PhoneNumberCampaignsAPI.md#deletephonenumbercampaign) | **Delete** /phone_number_campaigns/{phoneNumber} | Delete Phone Number Campaign
*PhoneNumberCampaignsAPI* | [**GetAllPhoneNumberCampaigns**](docs/PhoneNumberCampaignsAPI.md#getallphonenumbercampaigns) | **Get** /phone_number_campaigns | Retrieve All Phone Number Campaigns
*PhoneNumberCampaignsAPI* | [**GetSinglePhoneNumberCampaign**](docs/PhoneNumberCampaignsAPI.md#getsinglephonenumbercampaign) | **Get** /phone_number_campaigns/{phoneNumber} | Get Single Phone Number Campaign
*PhoneNumberCampaignsAPI* | [**PutPhoneNumberCampaign**](docs/PhoneNumberCampaignsAPI.md#putphonenumbercampaign) | **Put** /phone_number_campaigns/{phoneNumber} | Create New Phone Number Campaign
*PhoneNumberConfigurationsAPI* | [**DeletePhoneNumber**](docs/PhoneNumberConfigurationsAPI.md#deletephonenumber) | **Delete** /phone_numbers/{id} | Delete a phone number
*PhoneNumberConfigurationsAPI* | [**EnablePhoneNumberEmergency**](docs/PhoneNumberConfigurationsAPI.md#enablephonenumberemergency) | **Post** /phone_numbers/{id}/actions/enable_emergency | Enable emergency for a phone number
*PhoneNumberConfigurationsAPI* | [**GetPhoneNumberVoiceSettings**](docs/PhoneNumberConfigurationsAPI.md#getphonenumbervoicesettings) | **Get** /phone_numbers/{id}/voice | Retrieve a phone number with voice settings
*PhoneNumberConfigurationsAPI* | [**ListPhoneNumbers**](docs/PhoneNumberConfigurationsAPI.md#listphonenumbers) | **Get** /phone_numbers | List phone numbers
*PhoneNumberConfigurationsAPI* | [**ListPhoneNumbersWithVoiceSettings**](docs/PhoneNumberConfigurationsAPI.md#listphonenumberswithvoicesettings) | **Get** /phone_numbers/voice | List phone numbers with voice settings
*PhoneNumberConfigurationsAPI* | [**PhoneNumberBundleStatusChange**](docs/PhoneNumberConfigurationsAPI.md#phonenumberbundlestatuschange) | **Patch** /phone_numbers/{id}/actions/bundle_status_change | Change the bundle status for a phone number (set to being in a bundle or remove from a bundle)
*PhoneNumberConfigurationsAPI* | [**RetrievePhoneNumber**](docs/PhoneNumberConfigurationsAPI.md#retrievephonenumber) | **Get** /phone_numbers/{id} | Retrieve a phone number
*PhoneNumberConfigurationsAPI* | [**SlimListPhoneNumbers**](docs/PhoneNumberConfigurationsAPI.md#slimlistphonenumbers) | **Get** /phone_numbers/slim | Slim List phone numbers
*PhoneNumberConfigurationsAPI* | [**UpdatePhoneNumber**](docs/PhoneNumberConfigurationsAPI.md#updatephonenumber) | **Patch** /phone_numbers/{id} | Update a phone number
*PhoneNumberConfigurationsAPI* | [**UpdatePhoneNumberVoiceSettings**](docs/PhoneNumberConfigurationsAPI.md#updatephonenumbervoicesettings) | **Patch** /phone_numbers/{id}/voice | Update a phone number with voice settings
*PhoneNumberOrdersAPI* | [**CancelSubNumberOrder**](docs/PhoneNumberOrdersAPI.md#cancelsubnumberorder) | **Patch** /sub_number_orders/{sub_number_order_id}/cancel | Cancel a sub number order
*PhoneNumberOrdersAPI* | [**CreateComment**](docs/PhoneNumberOrdersAPI.md#createcomment) | **Post** /comments | Create a comment
*PhoneNumberOrdersAPI* | [**CreateNumberOrder**](docs/PhoneNumberOrdersAPI.md#createnumberorder) | **Post** /number_orders | Create a number order
*PhoneNumberOrdersAPI* | [**GetNumberOrderPhoneNumber**](docs/PhoneNumberOrdersAPI.md#getnumberorderphonenumber) | **Get** /number_order_phone_numbers/{number_order_phone_number_id} | Retrieve a single phone number within a number order.
*PhoneNumberOrdersAPI* | [**GetSubNumberOrder**](docs/PhoneNumberOrdersAPI.md#getsubnumberorder) | **Get** /sub_number_orders/{sub_number_order_id} | Retrieve a sub number order
*PhoneNumberOrdersAPI* | [**ListComments**](docs/PhoneNumberOrdersAPI.md#listcomments) | **Get** /comments | Retrieve all comments
*PhoneNumberOrdersAPI* | [**ListNumberOrders**](docs/PhoneNumberOrdersAPI.md#listnumberorders) | **Get** /number_orders | List number orders
*PhoneNumberOrdersAPI* | [**ListSubNumberOrders**](docs/PhoneNumberOrdersAPI.md#listsubnumberorders) | **Get** /sub_number_orders | List sub number orders
*PhoneNumberOrdersAPI* | [**MarkCommentRead**](docs/PhoneNumberOrdersAPI.md#markcommentread) | **Patch** /comments/{id}/read | Mark a comment as read
*PhoneNumberOrdersAPI* | [**RetrieveComment**](docs/PhoneNumberOrdersAPI.md#retrievecomment) | **Get** /comments/{id} | Retrieve a comment
*PhoneNumberOrdersAPI* | [**RetrieveNumberOrder**](docs/PhoneNumberOrdersAPI.md#retrievenumberorder) | **Get** /number_orders/{number_order_id} | Retrieve a number order
*PhoneNumberOrdersAPI* | [**RetrieveOrderPhoneNumbers**](docs/PhoneNumberOrdersAPI.md#retrieveorderphonenumbers) | **Get** /number_order_phone_numbers | Retrieve a list of phone numbers associated to orders
*PhoneNumberOrdersAPI* | [**UpdateNumberOrder**](docs/PhoneNumberOrdersAPI.md#updatenumberorder) | **Patch** /number_orders/{number_order_id} | Update a number order
*PhoneNumberOrdersAPI* | [**UpdateNumberOrderPhoneNumber**](docs/PhoneNumberOrdersAPI.md#updatenumberorderphonenumber) | **Patch** /number_order_phone_numbers/{number_order_phone_number_id} | Update requirements for a single phone number within a number order.
*PhoneNumberOrdersAPI* | [**UpdateSubNumberOrder**](docs/PhoneNumberOrdersAPI.md#updatesubnumberorder) | **Patch** /sub_number_orders/{sub_number_order_id} | Update a sub number order&#39;s requirements
*PhoneNumberPortingAPI* | [**PostPortabilityCheck**](docs/PhoneNumberPortingAPI.md#postportabilitycheck) | **Post** /portability_checks | Run a portability check
*PhoneNumberReservationsAPI* | [**CreateNumberReservation**](docs/PhoneNumberReservationsAPI.md#createnumberreservation) | **Post** /number_reservations | Create a number reservation
*PhoneNumberReservationsAPI* | [**ExtendNumberReservationExpiryTime**](docs/PhoneNumberReservationsAPI.md#extendnumberreservationexpirytime) | **Post** /number_reservations/{number_reservation_id}/actions/extend | Extend a number reservation
*PhoneNumberReservationsAPI* | [**ListNumberReservations**](docs/PhoneNumberReservationsAPI.md#listnumberreservations) | **Get** /number_reservations | List number reservations
*PhoneNumberReservationsAPI* | [**RetrieveNumberReservation**](docs/PhoneNumberReservationsAPI.md#retrievenumberreservation) | **Get** /number_reservations/{number_reservation_id} | Retrieve a number reservation
*PhoneNumberSearchAPI* | [**ListAvailablePhoneNumberBlocks**](docs/PhoneNumberSearchAPI.md#listavailablephonenumberblocks) | **Get** /available_phone_number_blocks | List available phone number blocks
*PhoneNumberSearchAPI* | [**ListAvailablePhoneNumbers**](docs/PhoneNumberSearchAPI.md#listavailablephonenumbers) | **Get** /available_phone_numbers | List available phone numbers
*PortingOrdersAPI* | [**ActivatePortingOrder**](docs/PortingOrdersAPI.md#activateportingorder) | **Post** /porting_orders/{id}/actions/activate | Activate every number in a porting order asynchronously.
*PortingOrdersAPI* | [**CancelPortingOrder**](docs/PortingOrdersAPI.md#cancelportingorder) | **Post** /porting_orders/{id}/actions/cancel | Cancel a porting order
*PortingOrdersAPI* | [**ConfirmPortingOrder**](docs/PortingOrdersAPI.md#confirmportingorder) | **Post** /porting_orders/{id}/actions/confirm | Submit a porting order.
*PortingOrdersAPI* | [**CreateAdditionalDocuments**](docs/PortingOrdersAPI.md#createadditionaldocuments) | **Post** /porting_orders/{id}/additional_documents | Create a list of additional documents
*PortingOrdersAPI* | [**CreateLoaConfiguration**](docs/PortingOrdersAPI.md#createloaconfiguration) | **Post** /porting/loa_configurations | Create a LOA configuration
*PortingOrdersAPI* | [**CreatePhoneNumberConfigurations**](docs/PortingOrdersAPI.md#createphonenumberconfigurations) | **Post** /porting_orders/phone_number_configurations | Create a list of phone number configurations
*PortingOrdersAPI* | [**CreatePortingOrder**](docs/PortingOrdersAPI.md#createportingorder) | **Post** /porting_orders | Create a porting order
*PortingOrdersAPI* | [**CreatePortingOrderComment**](docs/PortingOrdersAPI.md#createportingordercomment) | **Post** /porting_orders/{id}/comments | Create a comment for a porting order
*PortingOrdersAPI* | [**CreatePortingPhoneNumberBlock**](docs/PortingOrdersAPI.md#createportingphonenumberblock) | **Post** /porting_orders/{porting_order_id}/phone_number_blocks | Create a phone number block
*PortingOrdersAPI* | [**CreatePortingPhoneNumberExtension**](docs/PortingOrdersAPI.md#createportingphonenumberextension) | **Post** /porting_orders/{porting_order_id}/phone_number_extensions | Create a phone number extension
*PortingOrdersAPI* | [**CreatePortingReport**](docs/PortingOrdersAPI.md#createportingreport) | **Post** /porting/reports | Create a porting related report
*PortingOrdersAPI* | [**DeleteAdditionalDocument**](docs/PortingOrdersAPI.md#deleteadditionaldocument) | **Delete** /porting_orders/{id}/additional_documents/{additional_document_id} | Delete an additional document
*PortingOrdersAPI* | [**DeleteLoaConfiguration**](docs/PortingOrdersAPI.md#deleteloaconfiguration) | **Delete** /porting/loa_configurations/{id} | Delete a LOA configuration
*PortingOrdersAPI* | [**DeletePortingOrder**](docs/PortingOrdersAPI.md#deleteportingorder) | **Delete** /porting_orders/{id} | Delete a porting order
*PortingOrdersAPI* | [**DeletePortingPhoneNumberBlock**](docs/PortingOrdersAPI.md#deleteportingphonenumberblock) | **Delete** /porting_orders/{porting_order_id}/phone_number_blocks/{id} | Delete a phone number block
*PortingOrdersAPI* | [**DeletePortingPhoneNumberExtension**](docs/PortingOrdersAPI.md#deleteportingphonenumberextension) | **Delete** /porting_orders/{porting_order_id}/phone_number_extensions/{id} | Delete a phone number extension
*PortingOrdersAPI* | [**GetLoaConfiguration**](docs/PortingOrdersAPI.md#getloaconfiguration) | **Get** /porting/loa_configurations/{id} | Retrieve a LOA configuration
*PortingOrdersAPI* | [**GetPortingOrder**](docs/PortingOrdersAPI.md#getportingorder) | **Get** /porting_orders/{id} | Retrieve a porting order
*PortingOrdersAPI* | [**GetPortingOrderLoaTemplate**](docs/PortingOrdersAPI.md#getportingorderloatemplate) | **Get** /porting_orders/{id}/loa_template | Download a porting order loa template
*PortingOrdersAPI* | [**GetPortingOrderSubRequest**](docs/PortingOrdersAPI.md#getportingordersubrequest) | **Get** /porting_orders/{id}/sub_request | Retrieve the associated V1 sub_request_id and port_request_id
*PortingOrdersAPI* | [**GetPortingOrdersActivationJob**](docs/PortingOrdersAPI.md#getportingordersactivationjob) | **Get** /porting_orders/{id}/activation_jobs/{activationJobId} | Retrieve a porting activation job
*PortingOrdersAPI* | [**GetPortingReport**](docs/PortingOrdersAPI.md#getportingreport) | **Get** /porting/reports/{id} | Retrieve a report
*PortingOrdersAPI* | [**ListAdditionalDocuments**](docs/PortingOrdersAPI.md#listadditionaldocuments) | **Get** /porting_orders/{id}/additional_documents | List additional documents
*PortingOrdersAPI* | [**ListAllowedFocWindows**](docs/PortingOrdersAPI.md#listallowedfocwindows) | **Get** /porting_orders/{id}/allowed_foc_windows | List allowed FOC dates
*PortingOrdersAPI* | [**ListExceptionTypes**](docs/PortingOrdersAPI.md#listexceptiontypes) | **Get** /porting_orders/exception_types | List all exception types
*PortingOrdersAPI* | [**ListLoaConfigurations**](docs/PortingOrdersAPI.md#listloaconfigurations) | **Get** /porting/loa_configurations | List LOA configurations
*PortingOrdersAPI* | [**ListPhoneNumberConfigurations**](docs/PortingOrdersAPI.md#listphonenumberconfigurations) | **Get** /porting_orders/phone_number_configurations | List all phone number configurations
*PortingOrdersAPI* | [**ListPortingEvents**](docs/PortingOrdersAPI.md#listportingevents) | **Get** /porting/events | List all porting events
*PortingOrdersAPI* | [**ListPortingOrderActivationJobs**](docs/PortingOrdersAPI.md#listportingorderactivationjobs) | **Get** /porting_orders/{id}/activation_jobs | List all porting activation jobs
*PortingOrdersAPI* | [**ListPortingOrderComments**](docs/PortingOrdersAPI.md#listportingordercomments) | **Get** /porting_orders/{id}/comments | List all comments of a porting order
*PortingOrdersAPI* | [**ListPortingOrderRequirements**](docs/PortingOrdersAPI.md#listportingorderrequirements) | **Get** /porting_orders/{id}/requirements | List porting order requirements
*PortingOrdersAPI* | [**ListPortingOrders**](docs/PortingOrdersAPI.md#listportingorders) | **Get** /porting_orders | List all porting orders
*PortingOrdersAPI* | [**ListPortingPhoneNumberBlocks**](docs/PortingOrdersAPI.md#listportingphonenumberblocks) | **Get** /porting_orders/{porting_order_id}/phone_number_blocks | List all phone number blocks
*PortingOrdersAPI* | [**ListPortingPhoneNumberExtensions**](docs/PortingOrdersAPI.md#listportingphonenumberextensions) | **Get** /porting_orders/{porting_order_id}/phone_number_extensions | List all phone number extensions
*PortingOrdersAPI* | [**ListPortingPhoneNumbers**](docs/PortingOrdersAPI.md#listportingphonenumbers) | **Get** /porting_phone_numbers | List all porting phone numbers
*PortingOrdersAPI* | [**ListPortingReports**](docs/PortingOrdersAPI.md#listportingreports) | **Get** /porting/reports | List porting related reports
*PortingOrdersAPI* | [**ListVerificationCodes**](docs/PortingOrdersAPI.md#listverificationcodes) | **Get** /porting_orders/{id}/verification_codes | List verification codes
*PortingOrdersAPI* | [**PreviewLoaConfiguration**](docs/PortingOrdersAPI.md#previewloaconfiguration) | **Get** /porting/loa_configurations/{id}/preview | Preview a LOA configuration
*PortingOrdersAPI* | [**PreviewLoaConfigurationParams**](docs/PortingOrdersAPI.md#previewloaconfigurationparams) | **Post** /porting/loa_configuration/preview | Preview the LOA configuration parameters
*PortingOrdersAPI* | [**RepublishPortingEvent**](docs/PortingOrdersAPI.md#republishportingevent) | **Post** /porting/events/{id}/republish | Republish a porting event
*PortingOrdersAPI* | [**SendPortingVerificationCodes**](docs/PortingOrdersAPI.md#sendportingverificationcodes) | **Post** /porting_orders/{id}/verification_codes/send | Send the verification codes
*PortingOrdersAPI* | [**SharePortingOrder**](docs/PortingOrdersAPI.md#shareportingorder) | **Post** /porting_orders/{id}/actions/share | Share a porting order
*PortingOrdersAPI* | [**ShowPortingEvent**](docs/PortingOrdersAPI.md#showportingevent) | **Get** /porting/events/{id} | Show a porting event
*PortingOrdersAPI* | [**UpdateLoaConfiguration**](docs/PortingOrdersAPI.md#updateloaconfiguration) | **Patch** /porting/loa_configurations/{id} | Update a LOA configuration
*PortingOrdersAPI* | [**UpdatePortingOrder**](docs/PortingOrdersAPI.md#updateportingorder) | **Patch** /porting_orders/{id} | Edit a porting order
*PortingOrdersAPI* | [**UpdatePortingOrdersActivationJob**](docs/PortingOrdersAPI.md#updateportingordersactivationjob) | **Patch** /porting_orders/{id}/activation_jobs/{activationJobId} | Update a porting activation job
*PortingOrdersAPI* | [**VerifyPortingVerificationCodes**](docs/PortingOrdersAPI.md#verifyportingverificationcodes) | **Post** /porting_orders/{id}/verification_codes/verify | Verify the verification code for a list of phone numbers
*PresignedObjectURLsAPI* | [**CreatePresignedObjectUrl**](docs/PresignedObjectURLsAPI.md#createpresignedobjecturl) | **Post** /storage/buckets/{bucketName}/{objectName}/presigned_url | Create Presigned Object URL
*PrivateWirelessGatewaysAPI* | [**CreatePrivateWirelessGateway**](docs/PrivateWirelessGatewaysAPI.md#createprivatewirelessgateway) | **Post** /private_wireless_gateways | Create a Private Wireless Gateway
*PrivateWirelessGatewaysAPI* | [**DeleteWirelessGateway**](docs/PrivateWirelessGatewaysAPI.md#deletewirelessgateway) | **Delete** /private_wireless_gateways/{id} | Delete a Private Wireless Gateway
*PrivateWirelessGatewaysAPI* | [**GetPrivateWirelessGateway**](docs/PrivateWirelessGatewaysAPI.md#getprivatewirelessgateway) | **Get** /private_wireless_gateways/{id} | Get a Private Wireless Gateway
*PrivateWirelessGatewaysAPI* | [**GetPrivateWirelessGateways**](docs/PrivateWirelessGatewaysAPI.md#getprivatewirelessgateways) | **Get** /private_wireless_gateways | Get all Private Wireless Gateways
*ProgrammableFaxApplicationsAPI* | [**CreateFaxApplication**](docs/ProgrammableFaxApplicationsAPI.md#createfaxapplication) | **Post** /fax_applications | Creates a Fax Application
*ProgrammableFaxApplicationsAPI* | [**DeleteFaxApplication**](docs/ProgrammableFaxApplicationsAPI.md#deletefaxapplication) | **Delete** /fax_applications/{id} | Deletes a Fax Application
*ProgrammableFaxApplicationsAPI* | [**GetFaxApplication**](docs/ProgrammableFaxApplicationsAPI.md#getfaxapplication) | **Get** /fax_applications/{id} | Retrieve a Fax Application
*ProgrammableFaxApplicationsAPI* | [**ListFaxApplications**](docs/ProgrammableFaxApplicationsAPI.md#listfaxapplications) | **Get** /fax_applications | List all Fax Applications
*ProgrammableFaxApplicationsAPI* | [**UpdateFaxApplication**](docs/ProgrammableFaxApplicationsAPI.md#updatefaxapplication) | **Patch** /fax_applications/{id} | Update a Fax Application
*ProgrammableFaxCommandsAPI* | [**CancelFax**](docs/ProgrammableFaxCommandsAPI.md#cancelfax) | **Post** /faxes/{id}/actions/cancel | Cancel a fax
*ProgrammableFaxCommandsAPI* | [**DeleteFax**](docs/ProgrammableFaxCommandsAPI.md#deletefax) | **Delete** /faxes/{id} | Delete a fax
*ProgrammableFaxCommandsAPI* | [**ListFaxes**](docs/ProgrammableFaxCommandsAPI.md#listfaxes) | **Get** /faxes | View a list of faxes
*ProgrammableFaxCommandsAPI* | [**RefreshFax**](docs/ProgrammableFaxCommandsAPI.md#refreshfax) | **Post** /faxes/{id}/actions/refresh | Refresh a fax
*ProgrammableFaxCommandsAPI* | [**SendFax**](docs/ProgrammableFaxCommandsAPI.md#sendfax) | **Post** /faxes | Send a fax
*ProgrammableFaxCommandsAPI* | [**ViewFax**](docs/ProgrammableFaxCommandsAPI.md#viewfax) | **Get** /faxes/{id} | View a fax
*PublicInternetGatewaysAPI* | [**CreatePublicInternetGateway**](docs/PublicInternetGatewaysAPI.md#createpublicinternetgateway) | **Post** /public_internet_gateways | Create a Public Internet Gateway
*PublicInternetGatewaysAPI* | [**DeletePublicInternetGateway**](docs/PublicInternetGatewaysAPI.md#deletepublicinternetgateway) | **Delete** /public_internet_gateways/{id} | Delete a Public Internet Gateway
*PublicInternetGatewaysAPI* | [**GetPublicInternetGateway**](docs/PublicInternetGatewaysAPI.md#getpublicinternetgateway) | **Get** /public_internet_gateways/{id} | Retrieve a Public Internet Gateway
*PublicInternetGatewaysAPI* | [**ListPublicInternetGateways**](docs/PublicInternetGatewaysAPI.md#listpublicinternetgateways) | **Get** /public_internet_gateways | List all Public Internet Gateways
*PushCredentialsAPI* | [**CreatePushCredential**](docs/PushCredentialsAPI.md#createpushcredential) | **Post** /mobile_push_credentials | Creates a new mobile push credential
*PushCredentialsAPI* | [**DeletePushCredentialById**](docs/PushCredentialsAPI.md#deletepushcredentialbyid) | **Delete** /mobile_push_credentials/{push_credential_id} | Deletes a mobile push credential
*PushCredentialsAPI* | [**GetPushCredentialById**](docs/PushCredentialsAPI.md#getpushcredentialbyid) | **Get** /mobile_push_credentials/{push_credential_id} | Retrieves a mobile push credential
*PushCredentialsAPI* | [**ListPushCredentials**](docs/PushCredentialsAPI.md#listpushcredentials) | **Get** /mobile_push_credentials | List mobile push credentials
*QueueCommandsAPI* | [**ListQueueCalls**](docs/QueueCommandsAPI.md#listqueuecalls) | **Get** /queues/{queue_name}/calls | Retrieve calls from a queue
*QueueCommandsAPI* | [**RetrieveCallFromQueue**](docs/QueueCommandsAPI.md#retrievecallfromqueue) | **Get** /queues/{queue_name}/calls/{call_control_id} | Retrieve a call from a queue
*QueueCommandsAPI* | [**RetrieveCallQueue**](docs/QueueCommandsAPI.md#retrievecallqueue) | **Get** /queues/{queue_name} | Retrieve a call queue
*RCSMessagingAPI* | [**MesssagesRcsPost**](docs/RCSMessagingAPI.md#messsagesrcspost) | **Post** /messsages/rcs | Send an RCS message
*RegionsAPI* | [**ListRegions**](docs/RegionsAPI.md#listregions) | **Get** /regions | List all Regions
*RegulatoryRequirementsAPI* | [**ListRegulatoryRequirements**](docs/RegulatoryRequirementsAPI.md#listregulatoryrequirements) | **Get** /regulatory_requirements | Retrieve regulatory requirements
*RegulatoryRequirementsAPI* | [**ListRegulatoryRequirementsPhoneNumbers**](docs/RegulatoryRequirementsAPI.md#listregulatoryrequirementsphonenumbers) | **Get** /phone_numbers_regulatory_requirements | Retrieve regulatory requirements for a list of phone numbers
*ReportingAPI* | [**CreateWdrReport**](docs/ReportingAPI.md#createwdrreport) | **Post** /wireless/detail_records_reports | Create a Wireless Detail Records (WDRs) Report
*ReportingAPI* | [**DeleteWdrReport**](docs/ReportingAPI.md#deletewdrreport) | **Delete** /wireless/detail_records_reports/{id} | Delete a Wireless Detail Record (WDR) Report
*ReportingAPI* | [**GetWdrReport**](docs/ReportingAPI.md#getwdrreport) | **Get** /wireless/detail_records_reports/{id} | Get a Wireless Detail Record (WDR) Report
*ReportingAPI* | [**GetWdrReports**](docs/ReportingAPI.md#getwdrreports) | **Get** /wireless/detail_records_reports | Get all Wireless Detail Records (WDRs) Reports
*ReportsAPI* | [**CreateBillingGroupReport**](docs/ReportsAPI.md#createbillinggroupreport) | **Post** /ledger_billing_group_reports | Create a ledger billing group report
*ReportsAPI* | [**GetBillingGroupReport**](docs/ReportsAPI.md#getbillinggroupreport) | **Get** /ledger_billing_group_reports/{id} | Get a ledger billing group report
*RequirementGroupsAPI* | [**CreateRequirementGroup**](docs/RequirementGroupsAPI.md#createrequirementgroup) | **Post** /requirement_groups | Create a new requirement group
*RequirementGroupsAPI* | [**DeleteRequirementGroup**](docs/RequirementGroupsAPI.md#deleterequirementgroup) | **Delete** /requirement_groups/{id} | Delete a requirement group by ID
*RequirementGroupsAPI* | [**GetRequirementGroup**](docs/RequirementGroupsAPI.md#getrequirementgroup) | **Get** /requirement_groups/{id} | Get a single requirement group by ID
*RequirementGroupsAPI* | [**GetRequirementGroups**](docs/RequirementGroupsAPI.md#getrequirementgroups) | **Get** /requirement_groups | List requirement groups
*RequirementGroupsAPI* | [**SubmitRequirementGroup**](docs/RequirementGroupsAPI.md#submitrequirementgroup) | **Post** /requirement_groups/{id}/submit_for_approval | Submit a Requirement Group for Approval
*RequirementGroupsAPI* | [**UpdateNumberOrderPhoneNumberRequirementGroup**](docs/RequirementGroupsAPI.md#updatenumberorderphonenumberrequirementgroup) | **Post** /number_order_phone_numbers/{id}/requirement_group | Update requirement group for a phone number order
*RequirementGroupsAPI* | [**UpdateRequirementGroup**](docs/RequirementGroupsAPI.md#updaterequirementgroup) | **Patch** /requirement_groups/{id} | Update requirement values in requirement group
*RequirementGroupsAPI* | [**UpdateSubNumberOrderRequirementGroup**](docs/RequirementGroupsAPI.md#updatesubnumberorderrequirementgroup) | **Post** /sub_number_orders/{id}/requirement_group | Update requirement group for a sub number order
*RequirementTypesAPI* | [**ListRequirementTypes**](docs/RequirementTypesAPI.md#listrequirementtypes) | **Get** /requirement_types | List all requirement types
*RequirementTypesAPI* | [**RetrieveRequirementType**](docs/RequirementTypesAPI.md#retrieverequirementtype) | **Get** /requirement_types/{id} | Retrieve a requirement type
*RequirementsAPI* | [**ListRequirements**](docs/RequirementsAPI.md#listrequirements) | **Get** /requirements | List all requirements
*RequirementsAPI* | [**RetrieveDocumentRequirements**](docs/RequirementsAPI.md#retrievedocumentrequirements) | **Get** /requirements/{id} | Retrieve a document requirement
*RoomCompositionsAPI* | [**CreateRoomComposition**](docs/RoomCompositionsAPI.md#createroomcomposition) | **Post** /room_compositions | Create a room composition.
*RoomCompositionsAPI* | [**DeleteRoomComposition**](docs/RoomCompositionsAPI.md#deleteroomcomposition) | **Delete** /room_compositions/{room_composition_id} | Delete a room composition.
*RoomCompositionsAPI* | [**ListRoomCompositions**](docs/RoomCompositionsAPI.md#listroomcompositions) | **Get** /room_compositions | View a list of room compositions.
*RoomCompositionsAPI* | [**ViewRoomComposition**](docs/RoomCompositionsAPI.md#viewroomcomposition) | **Get** /room_compositions/{room_composition_id} | View a room composition.
*RoomParticipantsAPI* | [**ListRoomParticipants**](docs/RoomParticipantsAPI.md#listroomparticipants) | **Get** /room_participants | View a list of room participants.
*RoomParticipantsAPI* | [**ViewRoomParticipant**](docs/RoomParticipantsAPI.md#viewroomparticipant) | **Get** /room_participants/{room_participant_id} | View a room participant.
*RoomRecordingsAPI* | [**DeleteRoomRecording**](docs/RoomRecordingsAPI.md#deleteroomrecording) | **Delete** /room_recordings/{room_recording_id} | Delete a room recording.
*RoomRecordingsAPI* | [**DeleteRoomRecordings**](docs/RoomRecordingsAPI.md#deleteroomrecordings) | **Delete** /room_recordings | Delete several room recordings in a bulk.
*RoomRecordingsAPI* | [**ListRoomRecordings**](docs/RoomRecordingsAPI.md#listroomrecordings) | **Get** /room_recordings | View a list of room recordings.
*RoomRecordingsAPI* | [**ViewRoomRecording**](docs/RoomRecordingsAPI.md#viewroomrecording) | **Get** /room_recordings/{room_recording_id} | View a room recording.
*RoomSessionsAPI* | [**EndSession**](docs/RoomSessionsAPI.md#endsession) | **Post** /room_sessions/{room_session_id}/actions/end | End a room session.
*RoomSessionsAPI* | [**KickParticipantInSession**](docs/RoomSessionsAPI.md#kickparticipantinsession) | **Post** /room_sessions/{room_session_id}/actions/kick | Kick participants from a room session.
*RoomSessionsAPI* | [**ListRoomSessions**](docs/RoomSessionsAPI.md#listroomsessions) | **Get** /room_sessions | View a list of room sessions.
*RoomSessionsAPI* | [**MuteParticipantInSession**](docs/RoomSessionsAPI.md#muteparticipantinsession) | **Post** /room_sessions/{room_session_id}/actions/mute | Mute participants in room session.
*RoomSessionsAPI* | [**RetrieveListRoomParticipants**](docs/RoomSessionsAPI.md#retrievelistroomparticipants) | **Get** /room_sessions/{room_session_id}/participants | View a list of room participants.
*RoomSessionsAPI* | [**UnmuteParticipantInSession**](docs/RoomSessionsAPI.md#unmuteparticipantinsession) | **Post** /room_sessions/{room_session_id}/actions/unmute | Unmute participants in room session.
*RoomSessionsAPI* | [**ViewRoomSession**](docs/RoomSessionsAPI.md#viewroomsession) | **Get** /room_sessions/{room_session_id} | View a room session.
*RoomsAPI* | [**CreateRoom**](docs/RoomsAPI.md#createroom) | **Post** /rooms | Create a room.
*RoomsAPI* | [**DeleteRoom**](docs/RoomsAPI.md#deleteroom) | **Delete** /rooms/{room_id} | Delete a room.
*RoomsAPI* | [**ListRooms**](docs/RoomsAPI.md#listrooms) | **Get** /rooms | View a list of rooms.
*RoomsAPI* | [**RetrieveListRoomSessions**](docs/RoomsAPI.md#retrievelistroomsessions) | **Get** /rooms/{room_id}/sessions | View a list of room sessions.
*RoomsAPI* | [**UpdateRoom**](docs/RoomsAPI.md#updateroom) | **Patch** /rooms/{room_id} | Update a room.
*RoomsAPI* | [**ViewRoom**](docs/RoomsAPI.md#viewroom) | **Get** /rooms/{room_id} | View a room.
*RoomsClientTokensAPI* | [**CreateRoomClientToken**](docs/RoomsClientTokensAPI.md#createroomclienttoken) | **Post** /rooms/{room_id}/actions/generate_join_client_token | Create Client Token to join a room.
*RoomsClientTokensAPI* | [**RefreshRoomClientToken**](docs/RoomsClientTokensAPI.md#refreshroomclienttoken) | **Post** /rooms/{room_id}/actions/refresh_client_token | Refresh Client Token to join a room.
*SIMCardActionsAPI* | [**GetBulkSimCardAction**](docs/SIMCardActionsAPI.md#getbulksimcardaction) | **Get** /bulk_sim_card_actions/{id} | Get bulk SIM card action details
*SIMCardActionsAPI* | [**GetSimCardAction**](docs/SIMCardActionsAPI.md#getsimcardaction) | **Get** /sim_card_actions/{id} | Get SIM card action details
*SIMCardActionsAPI* | [**ListBulkSimCardActions**](docs/SIMCardActionsAPI.md#listbulksimcardactions) | **Get** /bulk_sim_card_actions | List bulk SIM card actions
*SIMCardActionsAPI* | [**ListSimCardActions**](docs/SIMCardActionsAPI.md#listsimcardactions) | **Get** /sim_card_actions | List SIM card actions
*SIMCardGroupActionsAPI* | [**GetSimCardGroupAction**](docs/SIMCardGroupActionsAPI.md#getsimcardgroupaction) | **Get** /sim_card_group_actions/{id} | Get SIM card group action details
*SIMCardGroupActionsAPI* | [**GetSimCardGroupActions**](docs/SIMCardGroupActionsAPI.md#getsimcardgroupactions) | **Get** /sim_card_group_actions | List SIM card group actions
*SIMCardGroupsAPI* | [**CreateSimCardGroup**](docs/SIMCardGroupsAPI.md#createsimcardgroup) | **Post** /sim_card_groups | Create a SIM card group
*SIMCardGroupsAPI* | [**DeleteSimCardGroup**](docs/SIMCardGroupsAPI.md#deletesimcardgroup) | **Delete** /sim_card_groups/{id} | Delete a SIM card group
*SIMCardGroupsAPI* | [**GetAllSimCardGroups**](docs/SIMCardGroupsAPI.md#getallsimcardgroups) | **Get** /sim_card_groups | Get all SIM card groups
*SIMCardGroupsAPI* | [**GetSimCardGroup**](docs/SIMCardGroupsAPI.md#getsimcardgroup) | **Get** /sim_card_groups/{id} | Get SIM card group
*SIMCardGroupsAPI* | [**RemoveSimCardGroupPrivateWirelessGateway**](docs/SIMCardGroupsAPI.md#removesimcardgroupprivatewirelessgateway) | **Post** /sim_card_groups/{id}/actions/remove_private_wireless_gateway | Request Private Wireless Gateway removal from SIM card group
*SIMCardGroupsAPI* | [**SetPrivateWirelessGatewayForSimCardGroup**](docs/SIMCardGroupsAPI.md#setprivatewirelessgatewayforsimcardgroup) | **Post** /sim_card_groups/{id}/actions/set_private_wireless_gateway | Request Private Wireless Gateway assignment for SIM card group
*SIMCardGroupsAPI* | [**UpdateSimCardGroup**](docs/SIMCardGroupsAPI.md#updatesimcardgroup) | **Patch** /sim_card_groups/{id} | Update a SIM card group
*SIMCardOrdersAPI* | [**CreateSimCardOrder**](docs/SIMCardOrdersAPI.md#createsimcardorder) | **Post** /sim_card_orders | Create a SIM card order
*SIMCardOrdersAPI* | [**GetSimCardOrder**](docs/SIMCardOrdersAPI.md#getsimcardorder) | **Get** /sim_card_orders/{id} | Get a single SIM card order
*SIMCardOrdersAPI* | [**GetSimCardOrders**](docs/SIMCardOrdersAPI.md#getsimcardorders) | **Get** /sim_card_orders | Get all SIM card orders
*SIMCardOrdersAPI* | [**PreviewSimCardOrders**](docs/SIMCardOrdersAPI.md#previewsimcardorders) | **Post** /sim_card_order_preview | Preview SIM card orders
*SIMCardsAPI* | [**DeleteSimCard**](docs/SIMCardsAPI.md#deletesimcard) | **Delete** /sim_cards/{id} | Deletes a SIM card
*SIMCardsAPI* | [**DeleteSimCardDataUsageNotifications**](docs/SIMCardsAPI.md#deletesimcarddatausagenotifications) | **Delete** /sim_card_data_usage_notifications/{id} | Delete SIM card data usage notifications
*SIMCardsAPI* | [**DisableSimCard**](docs/SIMCardsAPI.md#disablesimcard) | **Post** /sim_cards/{id}/actions/disable | Request a SIM card disable
*SIMCardsAPI* | [**EnableSimCard**](docs/SIMCardsAPI.md#enablesimcard) | **Post** /sim_cards/{id}/actions/enable | Request a SIM card enable
*SIMCardsAPI* | [**GetSimCard**](docs/SIMCardsAPI.md#getsimcard) | **Get** /sim_cards/{id} | Get SIM card
*SIMCardsAPI* | [**GetSimCardActivationCode**](docs/SIMCardsAPI.md#getsimcardactivationcode) | **Get** /sim_cards/{id}/activation_code | Get activation code for an eSIM
*SIMCardsAPI* | [**GetSimCardDataUsageNotification**](docs/SIMCardsAPI.md#getsimcarddatausagenotification) | **Get** /sim_card_data_usage_notifications/{id} | Get a single SIM card data usage notification
*SIMCardsAPI* | [**GetSimCardDeviceDetails**](docs/SIMCardsAPI.md#getsimcarddevicedetails) | **Get** /sim_cards/{id}/device_details | Get SIM card device details
*SIMCardsAPI* | [**GetSimCardPublicIp**](docs/SIMCardsAPI.md#getsimcardpublicip) | **Get** /sim_cards/{id}/public_ip | Get SIM card public IP definition
*SIMCardsAPI* | [**GetSimCards**](docs/SIMCardsAPI.md#getsimcards) | **Get** /sim_cards | Get all SIM cards
*SIMCardsAPI* | [**GetWirelessConnectivityLogs**](docs/SIMCardsAPI.md#getwirelessconnectivitylogs) | **Get** /sim_cards/{id}/wireless_connectivity_logs | List wireless connectivity logs
*SIMCardsAPI* | [**ListDataUsageNotifications**](docs/SIMCardsAPI.md#listdatausagenotifications) | **Get** /sim_card_data_usage_notifications | List SIM card data usage notifications
*SIMCardsAPI* | [**PatchSimCardDataUsageNotification**](docs/SIMCardsAPI.md#patchsimcarddatausagenotification) | **Patch** /sim_card_data_usage_notifications/{id} | Updates information for a SIM Card Data Usage Notification
*SIMCardsAPI* | [**PostSimCardDataUsageNotification**](docs/SIMCardsAPI.md#postsimcarddatausagenotification) | **Post** /sim_card_data_usage_notifications | Create a new SIM card data usage notification
*SIMCardsAPI* | [**PurchaseESim**](docs/SIMCardsAPI.md#purchaseesim) | **Post** /actions/purchase/esims | Purchase eSIMs
*SIMCardsAPI* | [**RegisterSimCards**](docs/SIMCardsAPI.md#registersimcards) | **Post** /actions/register/sim_cards | Register SIM cards
*SIMCardsAPI* | [**RemoveSimCardPublicIp**](docs/SIMCardsAPI.md#removesimcardpublicip) | **Post** /sim_cards/{id}/actions/remove_public_ip | Request removing a SIM card public IP
*SIMCardsAPI* | [**SetPublicIPsBulk**](docs/SIMCardsAPI.md#setpublicipsbulk) | **Post** /sim_cards/actions/bulk_set_public_ips | Request bulk setting SIM card public IPs.
*SIMCardsAPI* | [**SetSimCardPublicIp**](docs/SIMCardsAPI.md#setsimcardpublicip) | **Post** /sim_cards/{id}/actions/set_public_ip | Request setting a SIM card public IP
*SIMCardsAPI* | [**SetSimCardStandby**](docs/SIMCardsAPI.md#setsimcardstandby) | **Post** /sim_cards/{id}/actions/set_standby | Request setting a SIM card to standby
*SIMCardsAPI* | [**UpdateSimCard**](docs/SIMCardsAPI.md#updatesimcard) | **Patch** /sim_cards/{id} | Update a SIM card
*SIMCardsAPI* | [**ValidateRegistrationCodes**](docs/SIMCardsAPI.md#validateregistrationcodes) | **Post** /sim_cards/actions/validate_registration_codes | Validate SIM cards registration codes
*SIPRECConnectorsAPI* | [**CreateSiprecConnector**](docs/SIPRECConnectorsAPI.md#createsiprecconnector) | **Post** /siprec_connectors | Create a SIPREC connector
*SIPRECConnectorsAPI* | [**DeleteSiprecConnector**](docs/SIPRECConnectorsAPI.md#deletesiprecconnector) | **Delete** /siprec_connectors/{connector_name} | Delete a SIPREC connector
*SIPRECConnectorsAPI* | [**GetSiprecConnector**](docs/SIPRECConnectorsAPI.md#getsiprecconnector) | **Get** /siprec_connectors/{connector_name} | Retrieve a SIPREC connector
*SIPRECConnectorsAPI* | [**UpdateSiprecConnector**](docs/SIPRECConnectorsAPI.md#updatesiprecconnector) | **Put** /siprec_connectors/{connector_name} | Update a SIPREC connector
*SharedCampaignsAPI* | [**GetPartnerCampaignSharingStatus**](docs/SharedCampaignsAPI.md#getpartnercampaignsharingstatus) | **Get** /partnerCampaign/{campaignId}/sharing | Get Sharing Status
*SharedCampaignsAPI* | [**GetPartnerCampaignsSharedByUser**](docs/SharedCampaignsAPI.md#getpartnercampaignssharedbyuser) | **Get** /partnerCampaign/sharedByMe | Get Partner Campaigns Shared By User
*SharedCampaignsAPI* | [**GetSharedCampaign**](docs/SharedCampaignsAPI.md#getsharedcampaign) | **Get** /partner_campaigns/{campaignId} | Get Single Shared Campaign
*SharedCampaignsAPI* | [**GetSharedCampaigns**](docs/SharedCampaignsAPI.md#getsharedcampaigns) | **Get** /partner_campaigns | List Shared Campaigns
*SharedCampaignsAPI* | [**UpdateSharedCampaign**](docs/SharedCampaignsAPI.md#updatesharedcampaign) | **Patch** /partner_campaigns/{campaignId} | Update Single Shared Campaign
*ShortCodesAPI* | [**ListShortCodes**](docs/ShortCodesAPI.md#listshortcodes) | **Get** /short_codes | List short codes
*ShortCodesAPI* | [**RetrieveShortCode**](docs/ShortCodesAPI.md#retrieveshortcode) | **Get** /short_codes/{id} | Retrieve a short code
*ShortCodesAPI* | [**UpdateShortCode**](docs/ShortCodesAPI.md#updateshortcode) | **Patch** /short_codes/{id} | Update short code
*TeXMLApplicationsAPI* | [**CreateTexmlApplication**](docs/TeXMLApplicationsAPI.md#createtexmlapplication) | **Post** /texml_applications | Creates a TeXML Application
*TeXMLApplicationsAPI* | [**DeleteTexmlApplication**](docs/TeXMLApplicationsAPI.md#deletetexmlapplication) | **Delete** /texml_applications/{id} | Deletes a TeXML Application
*TeXMLApplicationsAPI* | [**FindTexmlApplications**](docs/TeXMLApplicationsAPI.md#findtexmlapplications) | **Get** /texml_applications | List all TeXML Applications
*TeXMLApplicationsAPI* | [**GetTexmlApplication**](docs/TeXMLApplicationsAPI.md#gettexmlapplication) | **Get** /texml_applications/{id} | Retrieve a TeXML Application
*TeXMLApplicationsAPI* | [**UpdateTexmlApplication**](docs/TeXMLApplicationsAPI.md#updatetexmlapplication) | **Patch** /texml_applications/{id} | Update a TeXML Application
*TeXMLRESTCommandsAPI* | [**CreateTexmlSecret**](docs/TeXMLRESTCommandsAPI.md#createtexmlsecret) | **Post** /texml/secrets | Create a TeXML secret
*TeXMLRESTCommandsAPI* | [**DeleteTeXMLCallRecording**](docs/TeXMLRESTCommandsAPI.md#deletetexmlcallrecording) | **Delete** /texml/Accounts/{account_sid}/Recordings/{recording_sid}.json | Delete recording resource
*TeXMLRESTCommandsAPI* | [**DeleteTeXMLRecordingTranscription**](docs/TeXMLRESTCommandsAPI.md#deletetexmlrecordingtranscription) | **Delete** /texml/Accounts/{account_sid}/Transcriptions/{recording_transcription_sid}.json | Delete a recording transcription
*TeXMLRESTCommandsAPI* | [**DeleteTexmlConferenceParticipant**](docs/TeXMLRESTCommandsAPI.md#deletetexmlconferenceparticipant) | **Delete** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Participants/{call_sid} | Delete a conference participant
*TeXMLRESTCommandsAPI* | [**DeprecatedInitiateTexmlCall**](docs/TeXMLRESTCommandsAPI.md#deprecatedinitiatetexmlcall) | **Post** /texml/calls/{application_id} | (Deprecated) Initiate an outbound call
*TeXMLRESTCommandsAPI* | [**DeprecatedUpdateTexmlCall**](docs/TeXMLRESTCommandsAPI.md#deprecatedupdatetexmlcall) | **Post** /texml/calls/{call_sid}/update | (Deprecated) Update call
*TeXMLRESTCommandsAPI* | [**DialTexmlConferenceParticipant**](docs/TeXMLRESTCommandsAPI.md#dialtexmlconferenceparticipant) | **Post** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Participants | Dial a new conference participant
*TeXMLRESTCommandsAPI* | [**FetchTeXMLCallRecordings**](docs/TeXMLRESTCommandsAPI.md#fetchtexmlcallrecordings) | **Get** /texml/Accounts/{account_sid}/Calls/{call_sid}/Recordings.json | Fetch recordings for a call
*TeXMLRESTCommandsAPI* | [**FetchTeXMLConferenceRecordings**](docs/TeXMLRESTCommandsAPI.md#fetchtexmlconferencerecordings) | **Get** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Recordings.json | Fetch recordings for a conference
*TeXMLRESTCommandsAPI* | [**GetTeXMLCallRecording**](docs/TeXMLRESTCommandsAPI.md#gettexmlcallrecording) | **Get** /texml/Accounts/{account_sid}/Recordings/{recording_sid}.json | Fetch recording resource
*TeXMLRESTCommandsAPI* | [**GetTeXMLCallRecordings**](docs/TeXMLRESTCommandsAPI.md#gettexmlcallrecordings) | **Get** /texml/Accounts/{account_sid}/Recordings.json | Fetch multiple recording resources
*TeXMLRESTCommandsAPI* | [**GetTeXMLRecordingTranscription**](docs/TeXMLRESTCommandsAPI.md#gettexmlrecordingtranscription) | **Get** /texml/Accounts/{account_sid}/Transcriptions/{recording_transcription_sid}.json | Fetch a recording transcription resource
*TeXMLRESTCommandsAPI* | [**GetTeXMLRecordingTranscriptions**](docs/TeXMLRESTCommandsAPI.md#gettexmlrecordingtranscriptions) | **Get** /texml/Accounts/{account_sid}/Transcriptions.json | List recording transcriptions
*TeXMLRESTCommandsAPI* | [**GetTexmlCall**](docs/TeXMLRESTCommandsAPI.md#gettexmlcall) | **Get** /texml/Accounts/{account_sid}/Calls/{call_sid} | Fetch a call
*TeXMLRESTCommandsAPI* | [**GetTexmlCalls**](docs/TeXMLRESTCommandsAPI.md#gettexmlcalls) | **Get** /texml/Accounts/{account_sid}/Calls | Fetch multiple call resources
*TeXMLRESTCommandsAPI* | [**GetTexmlConference**](docs/TeXMLRESTCommandsAPI.md#gettexmlconference) | **Get** /texml/Accounts/{account_sid}/Conferences/{conference_sid} | Fetch a conference resource
*TeXMLRESTCommandsAPI* | [**GetTexmlConferenceParticipant**](docs/TeXMLRESTCommandsAPI.md#gettexmlconferenceparticipant) | **Get** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Participants/{call_sid} | Get conference participant resource
*TeXMLRESTCommandsAPI* | [**GetTexmlConferenceParticipants**](docs/TeXMLRESTCommandsAPI.md#gettexmlconferenceparticipants) | **Get** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Participants | List conference participants
*TeXMLRESTCommandsAPI* | [**GetTexmlConferenceRecordings**](docs/TeXMLRESTCommandsAPI.md#gettexmlconferencerecordings) | **Get** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Recordings | List conference recordings
*TeXMLRESTCommandsAPI* | [**GetTexmlConferences**](docs/TeXMLRESTCommandsAPI.md#gettexmlconferences) | **Get** /texml/Accounts/{account_sid}/Conferences | List conference resources
*TeXMLRESTCommandsAPI* | [**InitiateTexmlCall**](docs/TeXMLRESTCommandsAPI.md#initiatetexmlcall) | **Post** /texml/Accounts/{account_sid}/Calls | Initiate an outbound call
*TeXMLRESTCommandsAPI* | [**StartTeXMLCallRecording**](docs/TeXMLRESTCommandsAPI.md#starttexmlcallrecording) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid}/Recordings.json | Request recording for a call
*TeXMLRESTCommandsAPI* | [**StartTeXMLCallStreaming**](docs/TeXMLRESTCommandsAPI.md#starttexmlcallstreaming) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid}/Streams.json | Start streaming media from a call.
*TeXMLRESTCommandsAPI* | [**StartTeXMLSiprecSession**](docs/TeXMLRESTCommandsAPI.md#starttexmlsiprecsession) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid}/Siprec.json | Request siprec session for a call
*TeXMLRESTCommandsAPI* | [**UpdateTeXMLCallRecording**](docs/TeXMLRESTCommandsAPI.md#updatetexmlcallrecording) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid}/Recordings/{recording_sid}.json | Update recording on a call
*TeXMLRESTCommandsAPI* | [**UpdateTeXMLCallStreaming**](docs/TeXMLRESTCommandsAPI.md#updatetexmlcallstreaming) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid}/Streams/{streaming_sid}.json | Update streaming on a call
*TeXMLRESTCommandsAPI* | [**UpdateTeXMLSiprecSession**](docs/TeXMLRESTCommandsAPI.md#updatetexmlsiprecsession) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid}/Siprec/{siprec_sid}.json | Updates siprec session for a call
*TeXMLRESTCommandsAPI* | [**UpdateTexmlCall**](docs/TeXMLRESTCommandsAPI.md#updatetexmlcall) | **Post** /texml/Accounts/{account_sid}/Calls/{call_sid} | Update call
*TeXMLRESTCommandsAPI* | [**UpdateTexmlConference**](docs/TeXMLRESTCommandsAPI.md#updatetexmlconference) | **Post** /texml/Accounts/{account_sid}/Conferences/{conference_sid} | Update a conference resource
*TeXMLRESTCommandsAPI* | [**UpdateTexmlConferenceParticipant**](docs/TeXMLRESTCommandsAPI.md#updatetexmlconferenceparticipant) | **Post** /texml/Accounts/{account_sid}/Conferences/{conference_sid}/Participants/{call_sid} | Update a conference participant
*TextToSpeechCommandsAPI* | [**GenerateTextToSpeech**](docs/TextToSpeechCommandsAPI.md#generatetexttospeech) | **Post** /text-to-speech/speech | Generate speech from text
*TextToSpeechCommandsAPI* | [**ListTextToSpeechVoices**](docs/TextToSpeechCommandsAPI.md#listtexttospeechvoices) | **Get** /text-to-speech/voices | List available text to speech voices
*UsageReportsBETAAPI* | [**GetUsageReports**](docs/UsageReportsBETAAPI.md#getusagereports) | **Get** /usage_reports | Get Telnyx product usage data (BETA)
*UsageReportsBETAAPI* | [**ListUsageReportsOptions**](docs/UsageReportsBETAAPI.md#listusagereportsoptions) | **Get** /usage_reports/options | Get Usage Reports query options (BETA)
*UserAddressesAPI* | [**CreateUserAddress**](docs/UserAddressesAPI.md#createuseraddress) | **Post** /user_addresses | Creates a user address
*UserAddressesAPI* | [**FindUserAddress**](docs/UserAddressesAPI.md#finduseraddress) | **Get** /user_addresses | List all user addresses
*UserAddressesAPI* | [**GetUserAddress**](docs/UserAddressesAPI.md#getuseraddress) | **Get** /user_addresses/{id} | Retrieve a user address
*UserBundlesAPI* | [**CreateUserBundlesBulk**](docs/UserBundlesAPI.md#createuserbundlesbulk) | **Post** /bundle_pricing/user_bundles/bulk | Create User Bundles
*UserBundlesAPI* | [**DeactivateUserBundle**](docs/UserBundlesAPI.md#deactivateuserbundle) | **Delete** /bundle_pricing/user_bundles/{user_bundle_id} | Deactivate User Bundle
*UserBundlesAPI* | [**GetUnusedUserBundles**](docs/UserBundlesAPI.md#getunuseduserbundles) | **Get** /bundle_pricing/user_bundles/unused | Get Unused User Bundles
*UserBundlesAPI* | [**GetUserBundleById**](docs/UserBundlesAPI.md#getuserbundlebyid) | **Get** /bundle_pricing/user_bundles/{user_bundle_id} | Get User Bundle by Id
*UserBundlesAPI* | [**GetUserBundleResources**](docs/UserBundlesAPI.md#getuserbundleresources) | **Get** /bundle_pricing/user_bundles/{user_bundle_id}/resources | Get User Bundle Resources
*UserBundlesAPI* | [**GetUserBundles**](docs/UserBundlesAPI.md#getuserbundles) | **Get** /bundle_pricing/user_bundles | Get User Bundles
*UserTagsAPI* | [**GetUserTags**](docs/UserTagsAPI.md#getusertags) | **Get** /user_tags | List User Tags
*VerifiedNumbersAPI* | [**CreateVerifiedNumber**](docs/VerifiedNumbersAPI.md#createverifiednumber) | **Post** /verified_numbers | Request phone number verification
*VerifiedNumbersAPI* | [**DeleteVerifiedNumber**](docs/VerifiedNumbersAPI.md#deleteverifiednumber) | **Delete** /verified_numbers/{phone_number} | Delete a verified number
*VerifiedNumbersAPI* | [**GetVerifiedNumber**](docs/VerifiedNumbersAPI.md#getverifiednumber) | **Get** /verified_numbers/{phone_number} | Retrieve a verified number
*VerifiedNumbersAPI* | [**ListVerifiedNumbers**](docs/VerifiedNumbersAPI.md#listverifiednumbers) | **Get** /verified_numbers | List all Verified Numbers
*VerifiedNumbersAPI* | [**VerifyVerificationCode**](docs/VerifiedNumbersAPI.md#verifyverificationcode) | **Post** /verified_numbers/{phone_number}/actions/verify | Submit verification code
*VerifyAPI* | [**CreateFlashcallVerification**](docs/VerifyAPI.md#createflashcallverification) | **Post** /verifications/flashcall | Trigger Flash call verification
*VerifyAPI* | [**CreateVerificationCall**](docs/VerifyAPI.md#createverificationcall) | **Post** /verifications/call | Trigger Call verification
*VerifyAPI* | [**CreateVerificationSms**](docs/VerifyAPI.md#createverificationsms) | **Post** /verifications/sms | Trigger SMS verification
*VerifyAPI* | [**CreateVerifyProfile**](docs/VerifyAPI.md#createverifyprofile) | **Post** /verify_profiles | Create a Verify profile
*VerifyAPI* | [**DeleteProfile**](docs/VerifyAPI.md#deleteprofile) | **Delete** /verify_profiles/{verify_profile_id} | Delete Verify profile
*VerifyAPI* | [**GetVerifyProfile**](docs/VerifyAPI.md#getverifyprofile) | **Get** /verify_profiles/{verify_profile_id} | Retrieve Verify profile
*VerifyAPI* | [**ListProfileMessageTemplates**](docs/VerifyAPI.md#listprofilemessagetemplates) | **Get** /verify_profiles/templates | Retrieve Verify profile message templates
*VerifyAPI* | [**ListProfiles**](docs/VerifyAPI.md#listprofiles) | **Get** /verify_profiles | List all Verify profiles
*VerifyAPI* | [**ListVerifications**](docs/VerifyAPI.md#listverifications) | **Get** /verifications/by_phone_number/{phone_number} | List verifications by phone number
*VerifyAPI* | [**RetrieveVerification**](docs/VerifyAPI.md#retrieveverification) | **Get** /verifications/{verification_id} | Retrieve verification
*VerifyAPI* | [**UpdateVerifyProfile**](docs/VerifyAPI.md#updateverifyprofile) | **Patch** /verify_profiles/{verify_profile_id} | Update Verify profile
*VerifyAPI* | [**VerifyVerificationCodeById**](docs/VerifyAPI.md#verifyverificationcodebyid) | **Post** /verifications/{verification_id}/actions/verify | Verify verification code by ID
*VerifyAPI* | [**VerifyVerificationCodeByPhoneNumber**](docs/VerifyAPI.md#verifyverificationcodebyphonenumber) | **Post** /verifications/by_phone_number/{phone_number}/actions/verify | Verify verification code by phone number
*VirtualCrossConnectsAPI* | [**CreateVirtualCrossConnect**](docs/VirtualCrossConnectsAPI.md#createvirtualcrossconnect) | **Post** /virtual_cross_connects | Create a Virtual Cross Connect
*VirtualCrossConnectsAPI* | [**DeleteVirtualCrossConnect**](docs/VirtualCrossConnectsAPI.md#deletevirtualcrossconnect) | **Delete** /virtual_cross_connects/{id} | Delete a Virtual Cross Connect
*VirtualCrossConnectsAPI* | [**GetVirtualCrossConnect**](docs/VirtualCrossConnectsAPI.md#getvirtualcrossconnect) | **Get** /virtual_cross_connects/{id} | Retrieve a Virtual Cross Connect
*VirtualCrossConnectsAPI* | [**ListVirtualCrossConnectCoverage**](docs/VirtualCrossConnectsAPI.md#listvirtualcrossconnectcoverage) | **Get** /virtual_cross_connects_coverage | List Virtual Cross Connect Cloud Coverage
*VirtualCrossConnectsAPI* | [**ListVirtualCrossConnects**](docs/VirtualCrossConnectsAPI.md#listvirtualcrossconnects) | **Get** /virtual_cross_connects | List all Virtual Cross Connects
*VirtualCrossConnectsAPI* | [**UpdateVirtualCrossConnect**](docs/VirtualCrossConnectsAPI.md#updatevirtualcrossconnect) | **Patch** /virtual_cross_connects/{id} | Update the Virtual Cross Connect
*VoiceChannelsAPI* | [**GetAllNumbersChannelZones**](docs/VoiceChannelsAPI.md#getallnumberschannelzones) | **Get** /list | List All Numbers using Channel Billing
*VoiceChannelsAPI* | [**GetChannelZones**](docs/VoiceChannelsAPI.md#getchannelzones) | **Get** /channel_zones | List your voice channels for non-US zones
*VoiceChannelsAPI* | [**GetNumbersChannelZones**](docs/VoiceChannelsAPI.md#getnumberschannelzones) | **Get** /list/{channel_zone_id} | List Numbers using Channel Billing for a specific Zone
*VoiceChannelsAPI* | [**ListInboundChannels**](docs/VoiceChannelsAPI.md#listinboundchannels) | **Get** /inbound_channels | List your voice channels for US Zone
*VoiceChannelsAPI* | [**PatchChannelZone**](docs/VoiceChannelsAPI.md#patchchannelzone) | **Put** /channel_zones/{channel_zone_id} | Update voice channels for non-US Zones
*VoiceChannelsAPI* | [**UpdateOutboundChannels**](docs/VoiceChannelsAPI.md#updateoutboundchannels) | **Patch** /inbound_channels | Update voice channels for US Zone
*VoicemailAPI* | [**CreateVoicemail**](docs/VoicemailAPI.md#createvoicemail) | **Post** /phone_numbers/{phone_number_id}/voicemail | Create voicemail
*VoicemailAPI* | [**GetVoicemail**](docs/VoicemailAPI.md#getvoicemail) | **Get** /phone_numbers/{phone_number_id}/voicemail | Get voicemail
*VoicemailAPI* | [**UpdateVoicemail**](docs/VoicemailAPI.md#updatevoicemail) | **Patch** /phone_numbers/{phone_number_id}/voicemail | Update voicemail
*WDRDetailReportsAPI* | [**GetPaginatedWdrs**](docs/WDRDetailReportsAPI.md#getpaginatedwdrs) | **Get** /reports/wdrs | Fetches all Wdr records
*WebhooksAPI* | [**GetWebhookDeliveries**](docs/WebhooksAPI.md#getwebhookdeliveries) | **Get** /webhook_deliveries | List webhook deliveries
*WebhooksAPI* | [**GetWebhookDelivery**](docs/WebhooksAPI.md#getwebhookdelivery) | **Get** /webhook_deliveries/{id} | Find webhook_delivery details by ID
*WireGuardInterfacesAPI* | [**CreateWireguardInterface**](docs/WireGuardInterfacesAPI.md#createwireguardinterface) | **Post** /wireguard_interfaces | Create a WireGuard Interface
*WireGuardInterfacesAPI* | [**CreateWireguardPeer**](docs/WireGuardInterfacesAPI.md#createwireguardpeer) | **Post** /wireguard_peers | Create a WireGuard Peer
*WireGuardInterfacesAPI* | [**DeleteWireguardInterface**](docs/WireGuardInterfacesAPI.md#deletewireguardinterface) | **Delete** /wireguard_interfaces/{id} | Delete a WireGuard Interface
*WireGuardInterfacesAPI* | [**DeleteWireguardPeer**](docs/WireGuardInterfacesAPI.md#deletewireguardpeer) | **Delete** /wireguard_peers/{id} | Delete the WireGuard Peer
*WireGuardInterfacesAPI* | [**GetWireguardInterface**](docs/WireGuardInterfacesAPI.md#getwireguardinterface) | **Get** /wireguard_interfaces/{id} | Retrieve a WireGuard Interfaces
*WireGuardInterfacesAPI* | [**GetWireguardPeer**](docs/WireGuardInterfacesAPI.md#getwireguardpeer) | **Get** /wireguard_peers/{id} | Retrieve the WireGuard Peer
*WireGuardInterfacesAPI* | [**GetWireguardPeerConfig**](docs/WireGuardInterfacesAPI.md#getwireguardpeerconfig) | **Get** /wireguard_peers/{id}/config | Retrieve Wireguard config template for Peer
*WireGuardInterfacesAPI* | [**ListWireguardInterfaces**](docs/WireGuardInterfacesAPI.md#listwireguardinterfaces) | **Get** /wireguard_interfaces | List all WireGuard Interfaces
*WireGuardInterfacesAPI* | [**ListWireguardPeers**](docs/WireGuardInterfacesAPI.md#listwireguardpeers) | **Get** /wireguard_peers | List all WireGuard Peers
*WireGuardInterfacesAPI* | [**UpdateWireguardPeer**](docs/WireGuardInterfacesAPI.md#updatewireguardpeer) | **Patch** /wireguard_peers/{id} | Update the WireGuard Peer
*WirelessRegionsAPI* | [**WirelessRegionsGetAll**](docs/WirelessRegionsAPI.md#wirelessregionsgetall) | **Get** /wireless/regions | Get all wireless regions


## Documentation For Models

 - [AIAssistantStartRequest](docs/AIAssistantStartRequest.md)
 - [AIAssistantStartRequestAssistant](docs/AIAssistantStartRequestAssistant.md)
 - [AIAssistantStartRequestVoiceSettings](docs/AIAssistantStartRequestVoiceSettings.md)
 - [AIAssistantStopRequest](docs/AIAssistantStopRequest.md)
 - [AcceptSuggestionsRequest](docs/AcceptSuggestionsRequest.md)
 - [AccessIPAddressListResponseSchema](docs/AccessIPAddressListResponseSchema.md)
 - [AccessIPAddressPOST](docs/AccessIPAddressPOST.md)
 - [AccessIPAddressResponseSchema](docs/AccessIPAddressResponseSchema.md)
 - [AccessIPRangeListResponseSchema](docs/AccessIPRangeListResponseSchema.md)
 - [AccessIPRangePOST](docs/AccessIPRangePOST.md)
 - [AccessIPRangeResponseSchema](docs/AccessIPRangeResponseSchema.md)
 - [ActionsParticipantsRequest](docs/ActionsParticipantsRequest.md)
 - [ActionsParticipantsRequestParticipants](docs/ActionsParticipantsRequestParticipants.md)
 - [ActivatePortingOrder202Response](docs/ActivatePortingOrder202Response.md)
 - [ActiveCall](docs/ActiveCall.md)
 - [ActiveCallsResponse](docs/ActiveCallsResponse.md)
 - [Address](docs/Address.md)
 - [AddressCreate](docs/AddressCreate.md)
 - [AddressSuggestionResponse](docs/AddressSuggestionResponse.md)
 - [AddressSuggestionResponseData](docs/AddressSuggestionResponseData.md)
 - [AdvancedOrderRequest](docs/AdvancedOrderRequest.md)
 - [AdvancedOrderResponse](docs/AdvancedOrderResponse.md)
 - [AltBusinessIdType](docs/AltBusinessIdType.md)
 - [AmdDetailRecord](docs/AmdDetailRecord.md)
 - [AnchorsiteOverride](docs/AnchorsiteOverride.md)
 - [AnswerRequest](docs/AnswerRequest.md)
 - [AssignProfileToCampaignRequest](docs/AssignProfileToCampaignRequest.md)
 - [AssignProfileToCampaignResponse](docs/AssignProfileToCampaignResponse.md)
 - [AssignmentTaskStatusResponse](docs/AssignmentTaskStatusResponse.md)
 - [Assistant](docs/Assistant.md)
 - [AssistantDeletedResponse](docs/AssistantDeletedResponse.md)
 - [AssistantToolsInner](docs/AssistantToolsInner.md)
 - [AssistantsListData](docs/AssistantsListData.md)
 - [Attempt](docs/Attempt.md)
 - [AudioTranscriptionResponse](docs/AudioTranscriptionResponse.md)
 - [AudioTranscriptionResponseSegments](docs/AudioTranscriptionResponseSegments.md)
 - [AuditEventChanges](docs/AuditEventChanges.md)
 - [AuditEventChangesFrom](docs/AuditEventChangesFrom.md)
 - [AuditEventChangesTo](docs/AuditEventChangesTo.md)
 - [AuditLog](docs/AuditLog.md)
 - [AuditLogList](docs/AuditLogList.md)
 - [AuthenticationProvider](docs/AuthenticationProvider.md)
 - [AuthenticationProviderCreate](docs/AuthenticationProviderCreate.md)
 - [AuthenticationProviderSettings](docs/AuthenticationProviderSettings.md)
 - [AutoRechargePref](docs/AutoRechargePref.md)
 - [AutoRechargePrefRequest](docs/AutoRechargePrefRequest.md)
 - [AutoRespConfigCreateSchema](docs/AutoRespConfigCreateSchema.md)
 - [AutorespConfigResponseSchema](docs/AutorespConfigResponseSchema.md)
 - [AutorespConfigSchema](docs/AutorespConfigSchema.md)
 - [AutorespConfigsResponseSchema](docs/AutorespConfigsResponseSchema.md)
 - [AvailablePhoneNumber](docs/AvailablePhoneNumber.md)
 - [AvailablePhoneNumberBlock](docs/AvailablePhoneNumberBlock.md)
 - [AvailablePhoneNumbersMetadata](docs/AvailablePhoneNumbersMetadata.md)
 - [AvailableService](docs/AvailableService.md)
 - [AzureConfigurationData](docs/AzureConfigurationData.md)
 - [BackgroundTaskStatus](docs/BackgroundTaskStatus.md)
 - [BackgroundTasksQueryResponse](docs/BackgroundTasksQueryResponse.md)
 - [BackgroundTasksQueryResponseData](docs/BackgroundTasksQueryResponseData.md)
 - [BillingBundleResponse](docs/BillingBundleResponse.md)
 - [BillingBundleSchema](docs/BillingBundleSchema.md)
 - [BillingBundleSummary](docs/BillingBundleSummary.md)
 - [BillingGroup](docs/BillingGroup.md)
 - [BookAppointmentTool](docs/BookAppointmentTool.md)
 - [BookAppointmentToolParams](docs/BookAppointmentToolParams.md)
 - [BrandBasic](docs/BrandBasic.md)
 - [BrandFeedback](docs/BrandFeedback.md)
 - [BrandFeedbackCategory](docs/BrandFeedbackCategory.md)
 - [BrandIdentityStatus](docs/BrandIdentityStatus.md)
 - [BrandOptionalAttributes](docs/BrandOptionalAttributes.md)
 - [BrandRecordSetCSP](docs/BrandRecordSetCSP.md)
 - [BrandRelationship](docs/BrandRelationship.md)
 - [BridgeRequest](docs/BridgeRequest.md)
 - [BucketAPIUsageResponse](docs/BucketAPIUsageResponse.md)
 - [BucketIds](docs/BucketIds.md)
 - [BucketNotFoundError](docs/BucketNotFoundError.md)
 - [BucketOps](docs/BucketOps.md)
 - [BucketOpsTotal](docs/BucketOpsTotal.md)
 - [BucketUsage](docs/BucketUsage.md)
 - [BulkMessagingSettingsUpdatePhoneNumbers](docs/BulkMessagingSettingsUpdatePhoneNumbers.md)
 - [BulkMessagingSettingsUpdatePhoneNumbersRequest](docs/BulkMessagingSettingsUpdatePhoneNumbersRequest.md)
 - [BulkRoomRecordingsDeleteResponse](docs/BulkRoomRecordingsDeleteResponse.md)
 - [BulkRoomRecordingsDeleteResponseData](docs/BulkRoomRecordingsDeleteResponseData.md)
 - [BulkSIMCardAction](docs/BulkSIMCardAction.md)
 - [BulkSIMCardActionDetailed](docs/BulkSIMCardActionDetailed.md)
 - [BundleLimitDirection](docs/BundleLimitDirection.md)
 - [BundleLimitSchema](docs/BundleLimitSchema.md)
 - [CSVDownloadResponse](docs/CSVDownloadResponse.md)
 - [Call](docs/Call.md)
 - [CallAIGatherEnded](docs/CallAIGatherEnded.md)
 - [CallAIGatherEndedEvent](docs/CallAIGatherEndedEvent.md)
 - [CallAIGatherEndedPayload](docs/CallAIGatherEndedPayload.md)
 - [CallAIGatherEndedPayloadMessageHistoryInner](docs/CallAIGatherEndedPayloadMessageHistoryInner.md)
 - [CallAIGatherMessageHistoryUpdated](docs/CallAIGatherMessageHistoryUpdated.md)
 - [CallAIGatherMessageHistoryUpdatedEvent](docs/CallAIGatherMessageHistoryUpdatedEvent.md)
 - [CallAIGatherMessageHistoryUpdatedPayload](docs/CallAIGatherMessageHistoryUpdatedPayload.md)
 - [CallAIGatherPartialResults](docs/CallAIGatherPartialResults.md)
 - [CallAIGatherPartialResultsEvent](docs/CallAIGatherPartialResultsEvent.md)
 - [CallAIGatherPartialResultsPayload](docs/CallAIGatherPartialResultsPayload.md)
 - [CallAnswered](docs/CallAnswered.md)
 - [CallAnsweredEvent](docs/CallAnsweredEvent.md)
 - [CallAnsweredPayload](docs/CallAnsweredPayload.md)
 - [CallBridged](docs/CallBridged.md)
 - [CallBridgedEvent](docs/CallBridgedEvent.md)
 - [CallBridgedPayload](docs/CallBridgedPayload.md)
 - [CallControlApplication](docs/CallControlApplication.md)
 - [CallControlApplicationInbound](docs/CallControlApplicationInbound.md)
 - [CallControlApplicationOutbound](docs/CallControlApplicationOutbound.md)
 - [CallControlApplicationResponse](docs/CallControlApplicationResponse.md)
 - [CallControlCommandResponse](docs/CallControlCommandResponse.md)
 - [CallControlCommandResponseWithRecordingID](docs/CallControlCommandResponseWithRecordingID.md)
 - [CallControlCommandResult](docs/CallControlCommandResult.md)
 - [CallControlCommandResultWithRecordingId](docs/CallControlCommandResultWithRecordingId.md)
 - [CallConversationInsightsGenerated](docs/CallConversationInsightsGenerated.md)
 - [CallConversationInsightsGeneratedEvent](docs/CallConversationInsightsGeneratedEvent.md)
 - [CallConversationInsightsGeneratedPayload](docs/CallConversationInsightsGeneratedPayload.md)
 - [CallConversationInsightsGeneratedPayloadResultsInner](docs/CallConversationInsightsGeneratedPayloadResultsInner.md)
 - [CallConversationInsightsGeneratedPayloadResultsInnerResult](docs/CallConversationInsightsGeneratedPayloadResultsInnerResult.md)
 - [CallCost](docs/CallCost.md)
 - [CallCostMeta](docs/CallCostMeta.md)
 - [CallCostMetaMeta](docs/CallCostMetaMeta.md)
 - [CallCostPayload](docs/CallCostPayload.md)
 - [CallCostPayloadCostPartsInner](docs/CallCostPayloadCostPartsInner.md)
 - [CallDtmfReceived](docs/CallDtmfReceived.md)
 - [CallDtmfReceivedEvent](docs/CallDtmfReceivedEvent.md)
 - [CallDtmfReceivedPayload](docs/CallDtmfReceivedPayload.md)
 - [CallEnqueued](docs/CallEnqueued.md)
 - [CallEnqueuedEvent](docs/CallEnqueuedEvent.md)
 - [CallEnqueuedPayload](docs/CallEnqueuedPayload.md)
 - [CallEvent](docs/CallEvent.md)
 - [CallForkStarted](docs/CallForkStarted.md)
 - [CallForkStartedEvent](docs/CallForkStartedEvent.md)
 - [CallForkStartedPayload](docs/CallForkStartedPayload.md)
 - [CallForkStopped](docs/CallForkStopped.md)
 - [CallForkStoppedEvent](docs/CallForkStoppedEvent.md)
 - [CallForwarding](docs/CallForwarding.md)
 - [CallGatherEnded](docs/CallGatherEnded.md)
 - [CallGatherEndedEvent](docs/CallGatherEndedEvent.md)
 - [CallGatherEndedPayload](docs/CallGatherEndedPayload.md)
 - [CallHangup](docs/CallHangup.md)
 - [CallHangupEvent](docs/CallHangupEvent.md)
 - [CallHangupPayload](docs/CallHangupPayload.md)
 - [CallInitiated](docs/CallInitiated.md)
 - [CallInitiatedEvent](docs/CallInitiatedEvent.md)
 - [CallInitiatedPayload](docs/CallInitiatedPayload.md)
 - [CallLeftQueue](docs/CallLeftQueue.md)
 - [CallLeftQueueEvent](docs/CallLeftQueueEvent.md)
 - [CallLeftQueuePayload](docs/CallLeftQueuePayload.md)
 - [CallMachineDetectionEnded](docs/CallMachineDetectionEnded.md)
 - [CallMachineDetectionEndedEvent](docs/CallMachineDetectionEndedEvent.md)
 - [CallMachineDetectionEndedPayload](docs/CallMachineDetectionEndedPayload.md)
 - [CallMachineGreetingEnded](docs/CallMachineGreetingEnded.md)
 - [CallMachineGreetingEndedEvent](docs/CallMachineGreetingEndedEvent.md)
 - [CallMachineGreetingEndedPayload](docs/CallMachineGreetingEndedPayload.md)
 - [CallMachinePremiumDetectionEnded](docs/CallMachinePremiumDetectionEnded.md)
 - [CallMachinePremiumDetectionEndedEvent](docs/CallMachinePremiumDetectionEndedEvent.md)
 - [CallMachinePremiumDetectionEndedPayload](docs/CallMachinePremiumDetectionEndedPayload.md)
 - [CallMachinePremiumGreetingEnded](docs/CallMachinePremiumGreetingEnded.md)
 - [CallMachinePremiumGreetingEndedEvent](docs/CallMachinePremiumGreetingEndedEvent.md)
 - [CallMachinePremiumGreetingEndedPayload](docs/CallMachinePremiumGreetingEndedPayload.md)
 - [CallPlaybackEnded](docs/CallPlaybackEnded.md)
 - [CallPlaybackEndedEvent](docs/CallPlaybackEndedEvent.md)
 - [CallPlaybackEndedPayload](docs/CallPlaybackEndedPayload.md)
 - [CallPlaybackStarted](docs/CallPlaybackStarted.md)
 - [CallPlaybackStartedEvent](docs/CallPlaybackStartedEvent.md)
 - [CallPlaybackStartedPayload](docs/CallPlaybackStartedPayload.md)
 - [CallRecording](docs/CallRecording.md)
 - [CallRecordingError](docs/CallRecordingError.md)
 - [CallRecordingErrorEvent](docs/CallRecordingErrorEvent.md)
 - [CallRecordingErrorPayload](docs/CallRecordingErrorPayload.md)
 - [CallRecordingSaved](docs/CallRecordingSaved.md)
 - [CallRecordingSavedEvent](docs/CallRecordingSavedEvent.md)
 - [CallRecordingSavedPayload](docs/CallRecordingSavedPayload.md)
 - [CallRecordingSavedPayloadPublicRecordingUrls](docs/CallRecordingSavedPayloadPublicRecordingUrls.md)
 - [CallRecordingSavedPayloadRecordingUrls](docs/CallRecordingSavedPayloadRecordingUrls.md)
 - [CallRecordingTranscriptionSaved](docs/CallRecordingTranscriptionSaved.md)
 - [CallRecordingTranscriptionSavedEvent](docs/CallRecordingTranscriptionSavedEvent.md)
 - [CallRecordingTranscriptionSavedPayload](docs/CallRecordingTranscriptionSavedPayload.md)
 - [CallReferCompleted](docs/CallReferCompleted.md)
 - [CallReferCompletedEvent](docs/CallReferCompletedEvent.md)
 - [CallReferCompletedPayload](docs/CallReferCompletedPayload.md)
 - [CallReferFailed](docs/CallReferFailed.md)
 - [CallReferFailedEvent](docs/CallReferFailedEvent.md)
 - [CallReferFailedPayload](docs/CallReferFailedPayload.md)
 - [CallReferStarted](docs/CallReferStarted.md)
 - [CallReferStartedEvent](docs/CallReferStartedEvent.md)
 - [CallReferStartedPayload](docs/CallReferStartedPayload.md)
 - [CallRequest](docs/CallRequest.md)
 - [CallRequestAnsweringMachineDetectionConfig](docs/CallRequestAnsweringMachineDetectionConfig.md)
 - [CallRequestConferenceConfig](docs/CallRequestConferenceConfig.md)
 - [CallRequestTo](docs/CallRequestTo.md)
 - [CallResource](docs/CallResource.md)
 - [CallResourceIndex](docs/CallResourceIndex.md)
 - [CallSiprecFailed](docs/CallSiprecFailed.md)
 - [CallSiprecFailedEvent](docs/CallSiprecFailedEvent.md)
 - [CallSiprecFailedPayload](docs/CallSiprecFailedPayload.md)
 - [CallSiprecStarted](docs/CallSiprecStarted.md)
 - [CallSiprecStartedEvent](docs/CallSiprecStartedEvent.md)
 - [CallSiprecStartedPayload](docs/CallSiprecStartedPayload.md)
 - [CallSiprecStopped](docs/CallSiprecStopped.md)
 - [CallSiprecStoppedEvent](docs/CallSiprecStoppedEvent.md)
 - [CallSiprecStoppedPayload](docs/CallSiprecStoppedPayload.md)
 - [CallSpeakEnded](docs/CallSpeakEnded.md)
 - [CallSpeakEndedEvent](docs/CallSpeakEndedEvent.md)
 - [CallSpeakEndedPayload](docs/CallSpeakEndedPayload.md)
 - [CallSpeakStarted](docs/CallSpeakStarted.md)
 - [CallSpeakStartedEvent](docs/CallSpeakStartedEvent.md)
 - [CallSpeakStartedPayload](docs/CallSpeakStartedPayload.md)
 - [CallStreamingFailed](docs/CallStreamingFailed.md)
 - [CallStreamingFailedEvent](docs/CallStreamingFailedEvent.md)
 - [CallStreamingFailedPayload](docs/CallStreamingFailedPayload.md)
 - [CallStreamingFailedPayloadStreamParams](docs/CallStreamingFailedPayloadStreamParams.md)
 - [CallStreamingStarted](docs/CallStreamingStarted.md)
 - [CallStreamingStartedEvent](docs/CallStreamingStartedEvent.md)
 - [CallStreamingStartedPayload](docs/CallStreamingStartedPayload.md)
 - [CallStreamingStopped](docs/CallStreamingStopped.md)
 - [CallStreamingStoppedEvent](docs/CallStreamingStoppedEvent.md)
 - [CallWithRecordingId](docs/CallWithRecordingId.md)
 - [CallbackWebhookMeta](docs/CallbackWebhookMeta.md)
 - [CallerName](docs/CallerName.md)
 - [CampaignCost](docs/CampaignCost.md)
 - [CampaignDeletionResponse](docs/CampaignDeletionResponse.md)
 - [CampaignRecordSetCSP](docs/CampaignRecordSetCSP.md)
 - [CampaignRequest](docs/CampaignRequest.md)
 - [CampaignSharingChain](docs/CampaignSharingChain.md)
 - [CampaignSharingStatus](docs/CampaignSharingStatus.md)
 - [CampaignStatusUpdateEvent](docs/CampaignStatusUpdateEvent.md)
 - [CancelPortingOrder200Response](docs/CancelPortingOrder200Response.md)
 - [Carrier](docs/Carrier.md)
 - [CdrGetSyncUsageReportResponse](docs/CdrGetSyncUsageReportResponse.md)
 - [CdrUsageReportResponse](docs/CdrUsageReportResponse.md)
 - [ChatCompletionRequest](docs/ChatCompletionRequest.md)
 - [ChatCompletionRequestToolsInner](docs/ChatCompletionRequestToolsInner.md)
 - [ChatCompletionResponseFormatParam](docs/ChatCompletionResponseFormatParam.md)
 - [ChatCompletionSystemMessageParam](docs/ChatCompletionSystemMessageParam.md)
 - [ChatCompletionSystemMessageParamContent](docs/ChatCompletionSystemMessageParamContent.md)
 - [ChatCompletionToolParam](docs/ChatCompletionToolParam.md)
 - [CheckAvailabilityTool](docs/CheckAvailabilityTool.md)
 - [CheckAvailabilityToolParams](docs/CheckAvailabilityToolParams.md)
 - [CivicAddress](docs/CivicAddress.md)
 - [ClientStateUpdateRequest](docs/ClientStateUpdateRequest.md)
 - [CloudflareSyncStatus](docs/CloudflareSyncStatus.md)
 - [ClusterNode](docs/ClusterNode.md)
 - [ClusteringRequestInfo](docs/ClusteringRequestInfo.md)
 - [ClusteringRequestInfoData](docs/ClusteringRequestInfoData.md)
 - [ClusteringStatusResponse](docs/ClusteringStatusResponse.md)
 - [ClusteringStatusResponseData](docs/ClusteringStatusResponseData.md)
 - [CnamListing](docs/CnamListing.md)
 - [Comment](docs/Comment.md)
 - [CompleteOTAUpdate](docs/CompleteOTAUpdate.md)
 - [CompleteOTAUpdateSettings](docs/CompleteOTAUpdateSettings.md)
 - [Conference](docs/Conference.md)
 - [ConferenceCommandResponse](docs/ConferenceCommandResponse.md)
 - [ConferenceCommandResult](docs/ConferenceCommandResult.md)
 - [ConferenceCreated](docs/ConferenceCreated.md)
 - [ConferenceCreatedEvent](docs/ConferenceCreatedEvent.md)
 - [ConferenceCreatedPayload](docs/ConferenceCreatedPayload.md)
 - [ConferenceDetailRecord](docs/ConferenceDetailRecord.md)
 - [ConferenceEnded](docs/ConferenceEnded.md)
 - [ConferenceEndedBy](docs/ConferenceEndedBy.md)
 - [ConferenceEndedEvent](docs/ConferenceEndedEvent.md)
 - [ConferenceEndedPayload](docs/ConferenceEndedPayload.md)
 - [ConferenceFloorChangedEvent](docs/ConferenceFloorChangedEvent.md)
 - [ConferenceFloorChangedEventPayload](docs/ConferenceFloorChangedEventPayload.md)
 - [ConferenceHoldRequest](docs/ConferenceHoldRequest.md)
 - [ConferenceMuteRequest](docs/ConferenceMuteRequest.md)
 - [ConferenceParticipantDetailRecord](docs/ConferenceParticipantDetailRecord.md)
 - [ConferenceParticipantJoined](docs/ConferenceParticipantJoined.md)
 - [ConferenceParticipantJoinedEvent](docs/ConferenceParticipantJoinedEvent.md)
 - [ConferenceParticipantJoinedPayload](docs/ConferenceParticipantJoinedPayload.md)
 - [ConferenceParticipantLeft](docs/ConferenceParticipantLeft.md)
 - [ConferenceParticipantLeftEvent](docs/ConferenceParticipantLeftEvent.md)
 - [ConferenceParticipantPlaybackEnded](docs/ConferenceParticipantPlaybackEnded.md)
 - [ConferenceParticipantPlaybackEndedEvent](docs/ConferenceParticipantPlaybackEndedEvent.md)
 - [ConferenceParticipantPlaybackEndedPayload](docs/ConferenceParticipantPlaybackEndedPayload.md)
 - [ConferenceParticipantPlaybackStarted](docs/ConferenceParticipantPlaybackStarted.md)
 - [ConferenceParticipantPlaybackStartedEvent](docs/ConferenceParticipantPlaybackStartedEvent.md)
 - [ConferenceParticipantSpeakEnded](docs/ConferenceParticipantSpeakEnded.md)
 - [ConferenceParticipantSpeakEndedEvent](docs/ConferenceParticipantSpeakEndedEvent.md)
 - [ConferenceParticipantSpeakEndedPayload](docs/ConferenceParticipantSpeakEndedPayload.md)
 - [ConferenceParticipantSpeakStarted](docs/ConferenceParticipantSpeakStarted.md)
 - [ConferenceParticipantSpeakStartedEvent](docs/ConferenceParticipantSpeakStartedEvent.md)
 - [ConferencePlayRequest](docs/ConferencePlayRequest.md)
 - [ConferencePlaybackEnded](docs/ConferencePlaybackEnded.md)
 - [ConferencePlaybackEndedEvent](docs/ConferencePlaybackEndedEvent.md)
 - [ConferencePlaybackEndedPayload](docs/ConferencePlaybackEndedPayload.md)
 - [ConferencePlaybackStarted](docs/ConferencePlaybackStarted.md)
 - [ConferencePlaybackStartedEvent](docs/ConferencePlaybackStartedEvent.md)
 - [ConferenceRecordingResource](docs/ConferenceRecordingResource.md)
 - [ConferenceRecordingResourceIndex](docs/ConferenceRecordingResourceIndex.md)
 - [ConferenceRecordingSaved](docs/ConferenceRecordingSaved.md)
 - [ConferenceRecordingSavedEvent](docs/ConferenceRecordingSavedEvent.md)
 - [ConferenceRecordingSavedPayload](docs/ConferenceRecordingSavedPayload.md)
 - [ConferenceResource](docs/ConferenceResource.md)
 - [ConferenceResourceIndex](docs/ConferenceResourceIndex.md)
 - [ConferenceResponse](docs/ConferenceResponse.md)
 - [ConferenceSpeakEnded](docs/ConferenceSpeakEnded.md)
 - [ConferenceSpeakEndedEvent](docs/ConferenceSpeakEndedEvent.md)
 - [ConferenceSpeakEndedPayload](docs/ConferenceSpeakEndedPayload.md)
 - [ConferenceSpeakRequest](docs/ConferenceSpeakRequest.md)
 - [ConferenceSpeakStarted](docs/ConferenceSpeakStarted.md)
 - [ConferenceSpeakStartedEvent](docs/ConferenceSpeakStartedEvent.md)
 - [ConferenceStopRequest](docs/ConferenceStopRequest.md)
 - [ConferenceUnholdRequest](docs/ConferenceUnholdRequest.md)
 - [ConferenceUnmuteRequest](docs/ConferenceUnmuteRequest.md)
 - [ConfirmPortingOrder200Response](docs/ConfirmPortingOrder200Response.md)
 - [Connection](docs/Connection.md)
 - [ConnectionResponse](docs/ConnectionResponse.md)
 - [ConnectionRtcpSettings](docs/ConnectionRtcpSettings.md)
 - [ConsumedData](docs/ConsumedData.md)
 - [Conversation](docs/Conversation.md)
 - [ConversationChannelType](docs/ConversationChannelType.md)
 - [ConversationInsight](docs/ConversationInsight.md)
 - [ConversationInsightConversationInsightsInner](docs/ConversationInsightConversationInsightsInner.md)
 - [ConversationInsightListData](docs/ConversationInsightListData.md)
 - [ConversationMessage](docs/ConversationMessage.md)
 - [ConversationMessageListData](docs/ConversationMessageListData.md)
 - [ConversationMessageToolCallsInner](docs/ConversationMessageToolCallsInner.md)
 - [ConversationMessageToolCallsInnerFunction](docs/ConversationMessageToolCallsInnerFunction.md)
 - [ConversationMetadataValue](docs/ConversationMetadataValue.md)
 - [ConversationsListData](docs/ConversationsListData.md)
 - [CostInformation](docs/CostInformation.md)
 - [CountryCoverage](docs/CountryCoverage.md)
 - [CountryCoverageLocal](docs/CountryCoverageLocal.md)
 - [CreateAdditionalDocuments201Response](docs/CreateAdditionalDocuments201Response.md)
 - [CreateAdditionalDocumentsRequest](docs/CreateAdditionalDocumentsRequest.md)
 - [CreateAdditionalDocumentsRequestAdditionalDocumentsInner](docs/CreateAdditionalDocumentsRequestAdditionalDocumentsInner.md)
 - [CreateAddress200Response](docs/CreateAddress200Response.md)
 - [CreateAndroidPushCredentialRequest](docs/CreateAndroidPushCredentialRequest.md)
 - [CreateAssistantRequest](docs/CreateAssistantRequest.md)
 - [CreateAuthenticationProvider200Response](docs/CreateAuthenticationProvider200Response.md)
 - [CreateBillingGroup200Response](docs/CreateBillingGroup200Response.md)
 - [CreateBillingGroupReport200Response](docs/CreateBillingGroupReport200Response.md)
 - [CreateBrand](docs/CreateBrand.md)
 - [CreateBucketRequest](docs/CreateBucketRequest.md)
 - [CreateCallControlApplicationRequest](docs/CreateCallControlApplicationRequest.md)
 - [CreateComment200Response](docs/CreateComment200Response.md)
 - [CreateComment200ResponseData](docs/CreateComment200ResponseData.md)
 - [CreateConferenceRequest](docs/CreateConferenceRequest.md)
 - [CreateConversationRequest](docs/CreateConversationRequest.md)
 - [CreateCredentialConnectionRequest](docs/CreateCredentialConnectionRequest.md)
 - [CreateCustomerServiceRecord201Response](docs/CreateCustomerServiceRecord201Response.md)
 - [CreateCustomerServiceRecordRequest](docs/CreateCustomerServiceRecordRequest.md)
 - [CreateDocument200Response](docs/CreateDocument200Response.md)
 - [CreateDocumentRequest](docs/CreateDocumentRequest.md)
 - [CreateDocumentRequestOneOf](docs/CreateDocumentRequestOneOf.md)
 - [CreateDocumentRequestOneOf1](docs/CreateDocumentRequestOneOf1.md)
 - [CreateDynamicEmergencyAddress201Response](docs/CreateDynamicEmergencyAddress201Response.md)
 - [CreateDynamicEmergencyEndpoint201Response](docs/CreateDynamicEmergencyEndpoint201Response.md)
 - [CreateExternalConnectionRequest](docs/CreateExternalConnectionRequest.md)
 - [CreateExternalConnectionRequestInbound](docs/CreateExternalConnectionRequestInbound.md)
 - [CreateExternalConnectionRequestOutbound](docs/CreateExternalConnectionRequestOutbound.md)
 - [CreateExternalConnectionUploadRequest](docs/CreateExternalConnectionUploadRequest.md)
 - [CreateFaxApplicationRequest](docs/CreateFaxApplicationRequest.md)
 - [CreateFaxApplicationRequestInbound](docs/CreateFaxApplicationRequestInbound.md)
 - [CreateFineTuningJobRequest](docs/CreateFineTuningJobRequest.md)
 - [CreateFineTuningJobRequestHyperparameters](docs/CreateFineTuningJobRequestHyperparameters.md)
 - [CreateFqdnConnectionRequest](docs/CreateFqdnConnectionRequest.md)
 - [CreateFqdnRequest](docs/CreateFqdnRequest.md)
 - [CreateGlobalIp202Response](docs/CreateGlobalIp202Response.md)
 - [CreateGlobalIpAssignment202Response](docs/CreateGlobalIpAssignment202Response.md)
 - [CreateGlobalIpHealthCheck202Response](docs/CreateGlobalIpHealthCheck202Response.md)
 - [CreateGroupMMSMessageRequest](docs/CreateGroupMMSMessageRequest.md)
 - [CreateInboundIpRequest](docs/CreateInboundIpRequest.md)
 - [CreateIntegrationSecretRequest](docs/CreateIntegrationSecretRequest.md)
 - [CreateInventoryCoverage200Response](docs/CreateInventoryCoverage200Response.md)
 - [CreateIosPushCredentialRequest](docs/CreateIosPushCredentialRequest.md)
 - [CreateIpConnectionRequest](docs/CreateIpConnectionRequest.md)
 - [CreateIpRequest](docs/CreateIpRequest.md)
 - [CreateLoaConfiguration201Response](docs/CreateLoaConfiguration201Response.md)
 - [CreateLongCodeMessageRequest](docs/CreateLongCodeMessageRequest.md)
 - [CreateManagedAccount200Response](docs/CreateManagedAccount200Response.md)
 - [CreateManagedAccount422Response](docs/CreateManagedAccount422Response.md)
 - [CreateManagedAccountRequest](docs/CreateManagedAccountRequest.md)
 - [CreateMessageRequest](docs/CreateMessageRequest.md)
 - [CreateMessagingHostedNumberOrderRequest](docs/CreateMessagingHostedNumberOrderRequest.md)
 - [CreateMessagingProfileRequest](docs/CreateMessagingProfileRequest.md)
 - [CreateMigration200Response](docs/CreateMigration200Response.md)
 - [CreateMigrationSource200Response](docs/CreateMigrationSource200Response.md)
 - [CreateNetwork200Response](docs/CreateNetwork200Response.md)
 - [CreateNotificationChannels200Response](docs/CreateNotificationChannels200Response.md)
 - [CreateNotificationProfile200Response](docs/CreateNotificationProfile200Response.md)
 - [CreateNotificationSetting200Response](docs/CreateNotificationSetting200Response.md)
 - [CreateNumberBlockOrderRequest](docs/CreateNumberBlockOrderRequest.md)
 - [CreateNumberOrderRequest](docs/CreateNumberOrderRequest.md)
 - [CreateNumberOrderRequestPhoneNumbersInner](docs/CreateNumberOrderRequestPhoneNumbersInner.md)
 - [CreateNumberPoolMessageRequest](docs/CreateNumberPoolMessageRequest.md)
 - [CreateNumberReservationRequest](docs/CreateNumberReservationRequest.md)
 - [CreateOutboundVoiceProfileRequest](docs/CreateOutboundVoiceProfileRequest.md)
 - [CreatePhoneNumberConfigurations201Response](docs/CreatePhoneNumberConfigurations201Response.md)
 - [CreatePhoneNumberConfigurationsRequest](docs/CreatePhoneNumberConfigurationsRequest.md)
 - [CreatePhoneNumberConfigurationsRequestPhoneNumberConfigurationsInner](docs/CreatePhoneNumberConfigurationsRequestPhoneNumberConfigurationsInner.md)
 - [CreatePortingOrder](docs/CreatePortingOrder.md)
 - [CreatePortingOrder201Response](docs/CreatePortingOrder201Response.md)
 - [CreatePortingOrderComment](docs/CreatePortingOrderComment.md)
 - [CreatePortingOrderComment201Response](docs/CreatePortingOrderComment201Response.md)
 - [CreatePortingPhoneNumberBlock201Response](docs/CreatePortingPhoneNumberBlock201Response.md)
 - [CreatePortingPhoneNumberBlockRequest](docs/CreatePortingPhoneNumberBlockRequest.md)
 - [CreatePortingPhoneNumberBlockRequestActivationRangesInner](docs/CreatePortingPhoneNumberBlockRequestActivationRangesInner.md)
 - [CreatePortingPhoneNumberBlockRequestPhoneNumberRange](docs/CreatePortingPhoneNumberBlockRequestPhoneNumberRange.md)
 - [CreatePortingPhoneNumberExtension201Response](docs/CreatePortingPhoneNumberExtension201Response.md)
 - [CreatePortingPhoneNumberExtensionRequest](docs/CreatePortingPhoneNumberExtensionRequest.md)
 - [CreatePortingPhoneNumberExtensionRequestActivationRangesInner](docs/CreatePortingPhoneNumberExtensionRequestActivationRangesInner.md)
 - [CreatePortingPhoneNumberExtensionRequestExtensionRange](docs/CreatePortingPhoneNumberExtensionRequestExtensionRange.md)
 - [CreatePortingReport201Response](docs/CreatePortingReport201Response.md)
 - [CreatePortingReportRequest](docs/CreatePortingReportRequest.md)
 - [CreatePortoutReport201Response](docs/CreatePortoutReport201Response.md)
 - [CreatePortoutReportRequest](docs/CreatePortoutReportRequest.md)
 - [CreatePrivateWirelessGateway202Response](docs/CreatePrivateWirelessGateway202Response.md)
 - [CreatePrivateWirelessGatewayRequest](docs/CreatePrivateWirelessGatewayRequest.md)
 - [CreatePublicInternetGateway202Response](docs/CreatePublicInternetGateway202Response.md)
 - [CreatePushCredentialRequest](docs/CreatePushCredentialRequest.md)
 - [CreateRequirementGroupRequest](docs/CreateRequirementGroupRequest.md)
 - [CreateRequirementGroupRequestRegulatoryRequirementsInner](docs/CreateRequirementGroupRequestRegulatoryRequirementsInner.md)
 - [CreateRoom201Response](docs/CreateRoom201Response.md)
 - [CreateRoomClientToken201Response](docs/CreateRoomClientToken201Response.md)
 - [CreateRoomClientToken201ResponseData](docs/CreateRoomClientToken201ResponseData.md)
 - [CreateRoomClientTokenRequest](docs/CreateRoomClientTokenRequest.md)
 - [CreateRoomComposition202Response](docs/CreateRoomComposition202Response.md)
 - [CreateRoomCompositionRequest](docs/CreateRoomCompositionRequest.md)
 - [CreateRoomRequest](docs/CreateRoomRequest.md)
 - [CreateScheduledEventRequest](docs/CreateScheduledEventRequest.md)
 - [CreateShortCodeMessageRequest](docs/CreateShortCodeMessageRequest.md)
 - [CreateSimCardGroup200Response](docs/CreateSimCardGroup200Response.md)
 - [CreateSimCardOrder200Response](docs/CreateSimCardOrder200Response.md)
 - [CreateTeXMLSecretRequest](docs/CreateTeXMLSecretRequest.md)
 - [CreateTeXMLSecretResult](docs/CreateTeXMLSecretResult.md)
 - [CreateTexmlApplicationRequest](docs/CreateTexmlApplicationRequest.md)
 - [CreateTexmlApplicationRequestInbound](docs/CreateTexmlApplicationRequestInbound.md)
 - [CreateTexmlApplicationRequestOutbound](docs/CreateTexmlApplicationRequestOutbound.md)
 - [CreateUploadRequestResponse](docs/CreateUploadRequestResponse.md)
 - [CreateUploadRequestResponse1](docs/CreateUploadRequestResponse1.md)
 - [CreateUserAddress200Response](docs/CreateUserAddress200Response.md)
 - [CreateUserBundlesBulkRequest](docs/CreateUserBundlesBulkRequest.md)
 - [CreateUserBundlesBulkRequestItemsInner](docs/CreateUserBundlesBulkRequestItemsInner.md)
 - [CreateVerificationRequestCall](docs/CreateVerificationRequestCall.md)
 - [CreateVerificationRequestFlashcall](docs/CreateVerificationRequestFlashcall.md)
 - [CreateVerificationRequestSMS](docs/CreateVerificationRequestSMS.md)
 - [CreateVerificationResponse](docs/CreateVerificationResponse.md)
 - [CreateVerifiedNumberRequest](docs/CreateVerifiedNumberRequest.md)
 - [CreateVerifiedNumberResponse](docs/CreateVerifiedNumberResponse.md)
 - [CreateVerifyProfileCallRequest](docs/CreateVerifyProfileCallRequest.md)
 - [CreateVerifyProfileFlashcallRequest](docs/CreateVerifyProfileFlashcallRequest.md)
 - [CreateVerifyProfileRequest](docs/CreateVerifyProfileRequest.md)
 - [CreateVerifyProfileSMSRequest](docs/CreateVerifyProfileSMSRequest.md)
 - [CreateVirtualCrossConnect200Response](docs/CreateVirtualCrossConnect200Response.md)
 - [CreateWdrReport201Response](docs/CreateWdrReport201Response.md)
 - [CreateWireguardInterface202Response](docs/CreateWireguardInterface202Response.md)
 - [CreateWireguardPeer202Response](docs/CreateWireguardPeer202Response.md)
 - [CreatedUserBundlesResponse](docs/CreatedUserBundlesResponse.md)
 - [CredentialConnection](docs/CredentialConnection.md)
 - [CredentialConnectionResponse](docs/CredentialConnectionResponse.md)
 - [CredentialInbound](docs/CredentialInbound.md)
 - [CredentialOutbound](docs/CredentialOutbound.md)
 - [CredentialsResponse](docs/CredentialsResponse.md)
 - [CsvDownload](docs/CsvDownload.md)
 - [Cursor](docs/Cursor.md)
 - [CursorPaginationMeta](docs/CursorPaginationMeta.md)
 - [CustomSipHeader](docs/CustomSipHeader.md)
 - [CustomStorageConfiguration](docs/CustomStorageConfiguration.md)
 - [CustomStorageConfigurationConfiguration](docs/CustomStorageConfigurationConfiguration.md)
 - [CustomerServiceRecord](docs/CustomerServiceRecord.md)
 - [CustomerServiceRecordAdditionalData](docs/CustomerServiceRecordAdditionalData.md)
 - [CustomerServiceRecordPhoneNumberCoverage](docs/CustomerServiceRecordPhoneNumberCoverage.md)
 - [CustomerServiceRecordResult](docs/CustomerServiceRecordResult.md)
 - [CustomerServiceRecordResultAddress](docs/CustomerServiceRecordResultAddress.md)
 - [CustomerServiceRecordResultAdmin](docs/CustomerServiceRecordResultAdmin.md)
 - [CustomerServiceRecordStatusChangedEvent](docs/CustomerServiceRecordStatusChangedEvent.md)
 - [CustomerServiceRecordStatusChangedEventPayload](docs/CustomerServiceRecordStatusChangedEventPayload.md)
 - [CustomerServiceRecordsPostRequest](docs/CustomerServiceRecordsPostRequest.md)
 - [DTMFTool](docs/DTMFTool.md)
 - [DataInner](docs/DataInner.md)
 - [DefaultGateway](docs/DefaultGateway.md)
 - [DeleteObjectsRequestInner](docs/DeleteObjectsRequestInner.md)
 - [DeprecatedInitiateCallRequest](docs/DeprecatedInitiateCallRequest.md)
 - [DetailRecord](docs/DetailRecord.md)
 - [DetailRecordsSearchResponse](docs/DetailRecordsSearchResponse.md)
 - [DialogflowConfig](docs/DialogflowConfig.md)
 - [DialogflowConnection](docs/DialogflowConnection.md)
 - [DialogflowConnectionResponse](docs/DialogflowConnectionResponse.md)
 - [Direction](docs/Direction.md)
 - [DismissRequestWasSuccessful](docs/DismissRequestWasSuccessful.md)
 - [DocReqsRequirement](docs/DocReqsRequirement.md)
 - [DocReqsRequirementType](docs/DocReqsRequirementType.md)
 - [DocReqsRequirementTypeAcceptanceCriteria](docs/DocReqsRequirementTypeAcceptanceCriteria.md)
 - [DocServiceDocument](docs/DocServiceDocument.md)
 - [DocServiceDocumentAllOfSize](docs/DocServiceDocumentAllOfSize.md)
 - [DocServiceDocumentLink](docs/DocServiceDocumentLink.md)
 - [DocServiceRecord](docs/DocServiceRecord.md)
 - [DownlinkData](docs/DownlinkData.md)
 - [DtmfType](docs/DtmfType.md)
 - [DynamicEmergencyAddress](docs/DynamicEmergencyAddress.md)
 - [DynamicEmergencyEndpoint](docs/DynamicEmergencyEndpoint.md)
 - [ESimPurchase](docs/ESimPurchase.md)
 - [ElevenLabsVoiceSettings](docs/ElevenLabsVoiceSettings.md)
 - [EligibilityNumberResponse](docs/EligibilityNumberResponse.md)
 - [EligibilityNumbersResponse](docs/EligibilityNumbersResponse.md)
 - [EmbeddingBucketRequest](docs/EmbeddingBucketRequest.md)
 - [EmbeddingMetadata](docs/EmbeddingMetadata.md)
 - [EmbeddingResponse](docs/EmbeddingResponse.md)
 - [EmbeddingResponseData](docs/EmbeddingResponseData.md)
 - [EmbeddingSimilaritySearchDocument](docs/EmbeddingSimilaritySearchDocument.md)
 - [EmbeddingSimilaritySearchRequest](docs/EmbeddingSimilaritySearchRequest.md)
 - [EmbeddingSimilaritySearchResponse](docs/EmbeddingSimilaritySearchResponse.md)
 - [EmbeddingUrlRequest](docs/EmbeddingUrlRequest.md)
 - [EmbeddingsBucketFiles](docs/EmbeddingsBucketFiles.md)
 - [EmbeddingsBucketFilesData](docs/EmbeddingsBucketFilesData.md)
 - [EmergencySettings](docs/EmergencySettings.md)
 - [EnableManagedAccountRequest](docs/EnableManagedAccountRequest.md)
 - [EnabledFeatures](docs/EnabledFeatures.md)
 - [EncryptedMedia](docs/EncryptedMedia.md)
 - [EndSession200Response](docs/EndSession200Response.md)
 - [EndSession200ResponseData](docs/EndSession200ResponseData.md)
 - [EnqueueRequest](docs/EnqueueRequest.md)
 - [EntityType](docs/EntityType.md)
 - [EnumListResponseInner](docs/EnumListResponseInner.md)
 - [Error](docs/Error.md)
 - [ErrorRecord](docs/ErrorRecord.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ErrorResponseErrorsInner](docs/ErrorResponseErrorsInner.md)
 - [ErrorResponseErrorsInnerMeta](docs/ErrorResponseErrorsInnerMeta.md)
 - [ErrorResponseErrorsInnerSource](docs/ErrorResponseErrorsInnerSource.md)
 - [ErrorSource](docs/ErrorSource.md)
 - [Errors](docs/Errors.md)
 - [EventStatus](docs/EventStatus.md)
 - [ExportPortingOrdersCSVReport](docs/ExportPortingOrdersCSVReport.md)
 - [ExportPortingOrdersCSVReportFilters](docs/ExportPortingOrdersCSVReportFilters.md)
 - [ExportPortoutsCSVReport](docs/ExportPortoutsCSVReport.md)
 - [ExportPortoutsCSVReportFilters](docs/ExportPortoutsCSVReportFilters.md)
 - [ExternalConnection](docs/ExternalConnection.md)
 - [ExternalConnectionPhoneNumber](docs/ExternalConnectionPhoneNumber.md)
 - [ExternalConnectionResponse](docs/ExternalConnectionResponse.md)
 - [ExternalSipConnection](docs/ExternalSipConnection.md)
 - [ExternalSipConnectionZoomOnly](docs/ExternalSipConnectionZoomOnly.md)
 - [ExternalVetting](docs/ExternalVetting.md)
 - [ExternalWdrDetailRecordDto](docs/ExternalWdrDetailRecordDto.md)
 - [ExternalWdrGetDetailResponse](docs/ExternalWdrGetDetailResponse.md)
 - [FQDNConnectionResponse](docs/FQDNConnectionResponse.md)
 - [FQDNResponse](docs/FQDNResponse.md)
 - [FailClusteringProcessRequest](docs/FailClusteringProcessRequest.md)
 - [Fax](docs/Fax.md)
 - [FaxApplication](docs/FaxApplication.md)
 - [FaxApplicationResponse](docs/FaxApplicationResponse.md)
 - [FaxDelivered](docs/FaxDelivered.md)
 - [FaxDeliveredPayload](docs/FaxDeliveredPayload.md)
 - [FaxFailed](docs/FaxFailed.md)
 - [FaxFailedPayload](docs/FaxFailedPayload.md)
 - [FaxMediaProcessed](docs/FaxMediaProcessed.md)
 - [FaxMediaProcessedPayload](docs/FaxMediaProcessedPayload.md)
 - [FaxQueued](docs/FaxQueued.md)
 - [FaxQueuedPayload](docs/FaxQueuedPayload.md)
 - [FaxSendingStarted](docs/FaxSendingStarted.md)
 - [FaxSendingStartedPayload](docs/FaxSendingStartedPayload.md)
 - [Feature](docs/Feature.md)
 - [FindAddresses200Response](docs/FindAddresses200Response.md)
 - [FindAuthenticationProviders200Response](docs/FindAuthenticationProviders200Response.md)
 - [FindNotificationsEvents200Response](docs/FindNotificationsEvents200Response.md)
 - [FindNotificationsEventsConditions200Response](docs/FindNotificationsEventsConditions200Response.md)
 - [FindNotificationsProfiles200Response](docs/FindNotificationsProfiles200Response.md)
 - [FindPortoutComments200Response](docs/FindPortoutComments200Response.md)
 - [FindPortoutRequest200Response](docs/FindPortoutRequest200Response.md)
 - [FindUserAddress200Response](docs/FindUserAddress200Response.md)
 - [FineTuningJob](docs/FineTuningJob.md)
 - [FineTuningJobHyperparameters](docs/FineTuningJobHyperparameters.md)
 - [FineTuningJobsListData](docs/FineTuningJobsListData.md)
 - [ForbiddenError](docs/ForbiddenError.md)
 - [ForbiddenErrorAllOfMeta](docs/ForbiddenErrorAllOfMeta.md)
 - [Fqdn](docs/Fqdn.md)
 - [FqdnConnection](docs/FqdnConnection.md)
 - [FqdnConnectionTransportProtocol](docs/FqdnConnectionTransportProtocol.md)
 - [FunctionDefinition](docs/FunctionDefinition.md)
 - [GCSConfigurationData](docs/GCSConfigurationData.md)
 - [GatherRequest](docs/GatherRequest.md)
 - [GatherUsingAIRequest](docs/GatherUsingAIRequest.md)
 - [GatherUsingAudioRequest](docs/GatherUsingAudioRequest.md)
 - [GatherUsingSpeakRequest](docs/GatherUsingSpeakRequest.md)
 - [GcbChannelZone](docs/GcbChannelZone.md)
 - [GcbPhoneNumber](docs/GcbPhoneNumber.md)
 - [GenerateTextToSpeechRequest](docs/GenerateTextToSpeechRequest.md)
 - [GenericError](docs/GenericError.md)
 - [GetAllCivicAddressesResponse](docs/GetAllCivicAddressesResponse.md)
 - [GetAllExternalConnectionsResponse](docs/GetAllExternalConnectionsResponse.md)
 - [GetAllFaxApplicationsResponse](docs/GetAllFaxApplicationsResponse.md)
 - [GetAllNumbersChannelZones200Response](docs/GetAllNumbersChannelZones200Response.md)
 - [GetAllNumbersChannelZones200ResponseDataInner](docs/GetAllNumbersChannelZones200ResponseDataInner.md)
 - [GetAllNumbersChannelZones200ResponseDataInnerNumbersInner](docs/GetAllNumbersChannelZones200ResponseDataInnerNumbersInner.md)
 - [GetAllSimCardGroups200Response](docs/GetAllSimCardGroups200Response.md)
 - [GetAllTelephonyCredentialResponse](docs/GetAllTelephonyCredentialResponse.md)
 - [GetAllTexmlApplicationsResponse](docs/GetAllTexmlApplicationsResponse.md)
 - [GetAutoRechargePrefs200Response](docs/GetAutoRechargePrefs200Response.md)
 - [GetBucketUsage200Response](docs/GetBucketUsage200Response.md)
 - [GetBulkSimCardAction200Response](docs/GetBulkSimCardAction200Response.md)
 - [GetChannelZones200Response](docs/GetChannelZones200Response.md)
 - [GetCivicAddressResponse](docs/GetCivicAddressResponse.md)
 - [GetConversationByIdPublicConversationsGet200Response](docs/GetConversationByIdPublicConversationsGet200Response.md)
 - [GetCustomerServiceRecord201Response](docs/GetCustomerServiceRecord201Response.md)
 - [GetCustomerServiceRecord404Response](docs/GetCustomerServiceRecord404Response.md)
 - [GetDefaultGateway200Response](docs/GetDefaultGateway200Response.md)
 - [GetEnumEndpoint200Response](docs/GetEnumEndpoint200Response.md)
 - [GetExternalConnectionPhoneNumberResponse](docs/GetExternalConnectionPhoneNumberResponse.md)
 - [GetFaxResponse](docs/GetFaxResponse.md)
 - [GetGlobalIpAssignmentHealth200Response](docs/GetGlobalIpAssignmentHealth200Response.md)
 - [GetGlobalIpAssignmentUsage200Response](docs/GetGlobalIpAssignmentUsage200Response.md)
 - [GetGlobalIpLatency200Response](docs/GetGlobalIpLatency200Response.md)
 - [GetGlobalIpUsage200Response](docs/GetGlobalIpUsage200Response.md)
 - [GetLogMessageResponse](docs/GetLogMessageResponse.md)
 - [GetMessage200Response](docs/GetMessage200Response.md)
 - [GetMessage200ResponseData](docs/GetMessage200ResponseData.md)
 - [GetMobileNetworkOperators200Response](docs/GetMobileNetworkOperators200Response.md)
 - [GetOtaUpdate200Response](docs/GetOtaUpdate200Response.md)
 - [GetPortRequestSupportingDocuments201Response](docs/GetPortRequestSupportingDocuments201Response.md)
 - [GetPortingOrder200Response](docs/GetPortingOrder200Response.md)
 - [GetPortingOrder200ResponseMeta](docs/GetPortingOrder200ResponseMeta.md)
 - [GetPortingOrderSubRequest200Response](docs/GetPortingOrderSubRequest200Response.md)
 - [GetPrivateWirelessGateway200Response](docs/GetPrivateWirelessGateway200Response.md)
 - [GetPrivateWirelessGateways200Response](docs/GetPrivateWirelessGateways200Response.md)
 - [GetRecordingTranscription200Response](docs/GetRecordingTranscription200Response.md)
 - [GetRecordingTranscriptions200Response](docs/GetRecordingTranscriptions200Response.md)
 - [GetRecordings200Response](docs/GetRecordings200Response.md)
 - [GetReleaseResponse](docs/GetReleaseResponse.md)
 - [GetSimCard200Response](docs/GetSimCard200Response.md)
 - [GetSimCardAction200Response](docs/GetSimCardAction200Response.md)
 - [GetSimCardActivationCode200Response](docs/GetSimCardActivationCode200Response.md)
 - [GetSimCardDeviceDetails200Response](docs/GetSimCardDeviceDetails200Response.md)
 - [GetSimCardGroupAction200Response](docs/GetSimCardGroupAction200Response.md)
 - [GetSimCardGroupActions200Response](docs/GetSimCardGroupActions200Response.md)
 - [GetSimCardOrders200Response](docs/GetSimCardOrders200Response.md)
 - [GetSimCardPublicIp200Response](docs/GetSimCardPublicIp200Response.md)
 - [GetSimCards200Response](docs/GetSimCards200Response.md)
 - [GetStorageAPIUsage200Response](docs/GetStorageAPIUsage200Response.md)
 - [GetStorageSSLCertificates200Response](docs/GetStorageSSLCertificates200Response.md)
 - [GetSubRequestByPortingOrder](docs/GetSubRequestByPortingOrder.md)
 - [GetUploadResponse](docs/GetUploadResponse.md)
 - [GetUploadsStatusResponse](docs/GetUploadsStatusResponse.md)
 - [GetUploadsStatusResponseData](docs/GetUploadsStatusResponseData.md)
 - [GetUserBalance200Response](docs/GetUserBalance200Response.md)
 - [GetUserTags200Response](docs/GetUserTags200Response.md)
 - [GetUserTags200ResponseData](docs/GetUserTags200ResponseData.md)
 - [GetVoicemail200Response](docs/GetVoicemail200Response.md)
 - [GetWdrReports200Response](docs/GetWdrReports200Response.md)
 - [GetWebhookDeliveries200Response](docs/GetWebhookDeliveries200Response.md)
 - [GetWebhookDelivery200Response](docs/GetWebhookDelivery200Response.md)
 - [GetWirelessConnectivityLogs200Response](docs/GetWirelessConnectivityLogs200Response.md)
 - [GlobalIP](docs/GlobalIP.md)
 - [GlobalIPAllowedPort](docs/GlobalIPAllowedPort.md)
 - [GlobalIPHealthCheck](docs/GlobalIPHealthCheck.md)
 - [GlobalIPProtocol](docs/GlobalIPProtocol.md)
 - [GlobalIpAssignment](docs/GlobalIpAssignment.md)
 - [GlobalIpAssignmentHealthMetric](docs/GlobalIpAssignmentHealthMetric.md)
 - [GlobalIpAssignmentHealthMetricGlobalIp](docs/GlobalIpAssignmentHealthMetricGlobalIp.md)
 - [GlobalIpAssignmentHealthMetricGlobalIpAssignment](docs/GlobalIpAssignmentHealthMetricGlobalIpAssignment.md)
 - [GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer](docs/GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer.md)
 - [GlobalIpAssignmentHealthMetricHealth](docs/GlobalIpAssignmentHealthMetricHealth.md)
 - [GlobalIpAssignmentUpdate](docs/GlobalIpAssignmentUpdate.md)
 - [GlobalIpAssignmentUsageMetric](docs/GlobalIpAssignmentUsageMetric.md)
 - [GlobalIpAssignmentUsageMetricGlobalIp](docs/GlobalIpAssignmentUsageMetricGlobalIp.md)
 - [GlobalIpAssignmentUsageMetricGlobalIpAssignment](docs/GlobalIpAssignmentUsageMetricGlobalIpAssignment.md)
 - [GlobalIpAssignmentUsageMetricGlobalIpAssignmentWireguardPeer](docs/GlobalIpAssignmentUsageMetricGlobalIpAssignmentWireguardPeer.md)
 - [GlobalIpAssignmentUsageMetricReceived](docs/GlobalIpAssignmentUsageMetricReceived.md)
 - [GlobalIpAssignmentUsageMetricTransmitted](docs/GlobalIpAssignmentUsageMetricTransmitted.md)
 - [GlobalIpHealthCheckType](docs/GlobalIpHealthCheckType.md)
 - [GlobalIpLatencyMetric](docs/GlobalIpLatencyMetric.md)
 - [GlobalIpLatencyMetricMeanLatency](docs/GlobalIpLatencyMetricMeanLatency.md)
 - [GlobalIpLatencyMetricPercentileLatency](docs/GlobalIpLatencyMetricPercentileLatency.md)
 - [GlobalIpLatencyMetricPercentileLatency0](docs/GlobalIpLatencyMetricPercentileLatency0.md)
 - [GlobalIpLatencyMetricPercentileLatency100](docs/GlobalIpLatencyMetricPercentileLatency100.md)
 - [GlobalIpLatencyMetricPercentileLatency25](docs/GlobalIpLatencyMetricPercentileLatency25.md)
 - [GlobalIpLatencyMetricPercentileLatency50](docs/GlobalIpLatencyMetricPercentileLatency50.md)
 - [GlobalIpLatencyMetricPercentileLatency75](docs/GlobalIpLatencyMetricPercentileLatency75.md)
 - [GlobalIpLatencyMetricPercentileLatency90](docs/GlobalIpLatencyMetricPercentileLatency90.md)
 - [GlobalIpLatencyMetricPercentileLatency99](docs/GlobalIpLatencyMetricPercentileLatency99.md)
 - [GlobalIpLatencyMetricProberLocation](docs/GlobalIpLatencyMetricProberLocation.md)
 - [GlobalIpUsageMetric](docs/GlobalIpUsageMetric.md)
 - [GlobalIpUsageMetricGlobalIp](docs/GlobalIpUsageMetricGlobalIp.md)
 - [GoogleTranscriptionLanguage](docs/GoogleTranscriptionLanguage.md)
 - [GoogleTranscriptionLanguageLong](docs/GoogleTranscriptionLanguageLong.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [HangupRequest](docs/HangupRequest.md)
 - [HangupTool](docs/HangupTool.md)
 - [HangupToolParams](docs/HangupToolParams.md)
 - [HostedNumber](docs/HostedNumber.md)
 - [Http](docs/Http.md)
 - [HttpRequest](docs/HttpRequest.md)
 - [HttpResponse](docs/HttpResponse.md)
 - [ImportExternalVetting](docs/ImportExternalVetting.md)
 - [InboundFqdn](docs/InboundFqdn.md)
 - [InboundIp](docs/InboundIp.md)
 - [InboundMessage](docs/InboundMessage.md)
 - [InboundMessageEvent](docs/InboundMessageEvent.md)
 - [InboundMessagePayload](docs/InboundMessagePayload.md)
 - [InboundMessagePayloadCcInner](docs/InboundMessagePayloadCcInner.md)
 - [InboundMessagePayloadCost](docs/InboundMessagePayloadCost.md)
 - [InboundMessagePayloadCostBreakdown](docs/InboundMessagePayloadCostBreakdown.md)
 - [InboundMessagePayloadCostBreakdownCarrierFee](docs/InboundMessagePayloadCostBreakdownCarrierFee.md)
 - [InboundMessagePayloadCostBreakdownRate](docs/InboundMessagePayloadCostBreakdownRate.md)
 - [InboundMessagePayloadFrom](docs/InboundMessagePayloadFrom.md)
 - [InboundMessagePayloadMediaInner](docs/InboundMessagePayloadMediaInner.md)
 - [InboundMessagePayloadToInner](docs/InboundMessagePayloadToInner.md)
 - [InitiateCallRequest](docs/InitiateCallRequest.md)
 - [InitiateCallResult](docs/InitiateCallResult.md)
 - [InitiateTeXMLCallResponse](docs/InitiateTeXMLCallResponse.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject1Meta](docs/InlineObject1Meta.md)
 - [InsightSettings](docs/InsightSettings.md)
 - [IntegrationSecret](docs/IntegrationSecret.md)
 - [IntegrationSecretCreatedResponse](docs/IntegrationSecretCreatedResponse.md)
 - [IntegrationSecretsListData](docs/IntegrationSecretsListData.md)
 - [Interface](docs/Interface.md)
 - [InterfaceStatus](docs/InterfaceStatus.md)
 - [InterruptionSettings](docs/InterruptionSettings.md)
 - [InventoryCoverage](docs/InventoryCoverage.md)
 - [InventoryCoverageMetadata](docs/InventoryCoverageMetadata.md)
 - [Ip](docs/Ip.md)
 - [IpConnection](docs/IpConnection.md)
 - [IpConnectionResponse](docs/IpConnectionResponse.md)
 - [IpResponse](docs/IpResponse.md)
 - [JoinConferenceRequest](docs/JoinConferenceRequest.md)
 - [LeaveConferenceRequest](docs/LeaveConferenceRequest.md)
 - [LeaveQueueRequest](docs/LeaveQueueRequest.md)
 - [LedgerBillingGroupReport](docs/LedgerBillingGroupReport.md)
 - [ListAdditionalDocuments200Response](docs/ListAdditionalDocuments200Response.md)
 - [ListAdvancedOrderResponse](docs/ListAdvancedOrderResponse.md)
 - [ListAllocatableGlobalOutboundChannels200Response](docs/ListAllocatableGlobalOutboundChannels200Response.md)
 - [ListAllowedFocWindows200Response](docs/ListAllowedFocWindows200Response.md)
 - [ListAvailablePhoneNumbersBlocksResponse](docs/ListAvailablePhoneNumbersBlocksResponse.md)
 - [ListAvailablePhoneNumbersResponse](docs/ListAvailablePhoneNumbersResponse.md)
 - [ListBillingGroups200Response](docs/ListBillingGroups200Response.md)
 - [ListBucketsResponse](docs/ListBucketsResponse.md)
 - [ListBucketsResponseBucketsInner](docs/ListBucketsResponseBucketsInner.md)
 - [ListBulkSimCardActions200Response](docs/ListBulkSimCardActions200Response.md)
 - [ListCSVDownloadsResponse](docs/ListCSVDownloadsResponse.md)
 - [ListCallControlApplicationsResponse](docs/ListCallControlApplicationsResponse.md)
 - [ListCallEventsResponse](docs/ListCallEventsResponse.md)
 - [ListComments200Response](docs/ListComments200Response.md)
 - [ListConferencesResponse](docs/ListConferencesResponse.md)
 - [ListConnectionsResponse](docs/ListConnectionsResponse.md)
 - [ListCredentialConnectionsResponse](docs/ListCredentialConnectionsResponse.md)
 - [ListCustomerServiceRecords200Response](docs/ListCustomerServiceRecords200Response.md)
 - [ListCustomerServiceRecords401Response](docs/ListCustomerServiceRecords401Response.md)
 - [ListCustomerServiceRecords403Response](docs/ListCustomerServiceRecords403Response.md)
 - [ListCustomerServiceRecords422Response](docs/ListCustomerServiceRecords422Response.md)
 - [ListCustomerServiceRecords500Response](docs/ListCustomerServiceRecords500Response.md)
 - [ListDataUsageNotifications200Response](docs/ListDataUsageNotifications200Response.md)
 - [ListDocumentLinks200Response](docs/ListDocumentLinks200Response.md)
 - [ListDocuments200Response](docs/ListDocuments200Response.md)
 - [ListDynamicEmergencyAddresses200Response](docs/ListDynamicEmergencyAddresses200Response.md)
 - [ListDynamicEmergencyEndpoints200Response](docs/ListDynamicEmergencyEndpoints200Response.md)
 - [ListExceptionTypes200Response](docs/ListExceptionTypes200Response.md)
 - [ListExternalConnectionPhoneNumbersResponse](docs/ListExternalConnectionPhoneNumbersResponse.md)
 - [ListFQDNConnectionsResponse](docs/ListFQDNConnectionsResponse.md)
 - [ListFQDNsResponse](docs/ListFQDNsResponse.md)
 - [ListFaxesResponse](docs/ListFaxesResponse.md)
 - [ListGlobalIpAllowedPorts200Response](docs/ListGlobalIpAllowedPorts200Response.md)
 - [ListGlobalIpAssignments200Response](docs/ListGlobalIpAssignments200Response.md)
 - [ListGlobalIpHealthCheckTypes200Response](docs/ListGlobalIpHealthCheckTypes200Response.md)
 - [ListGlobalIpHealthChecks200Response](docs/ListGlobalIpHealthChecks200Response.md)
 - [ListGlobalIpProtocols200Response](docs/ListGlobalIpProtocols200Response.md)
 - [ListGlobalIps200Response](docs/ListGlobalIps200Response.md)
 - [ListInboundChannels200Response](docs/ListInboundChannels200Response.md)
 - [ListInboundChannels200ResponseData](docs/ListInboundChannels200ResponseData.md)
 - [ListIpConnectionsResponse](docs/ListIpConnectionsResponse.md)
 - [ListIpsResponse](docs/ListIpsResponse.md)
 - [ListLoaConfigurations200Response](docs/ListLoaConfigurations200Response.md)
 - [ListLogMessagesResponse](docs/ListLogMessagesResponse.md)
 - [ListManagedAccounts200Response](docs/ListManagedAccounts200Response.md)
 - [ListMessagingHostedNumberOrderResponse](docs/ListMessagingHostedNumberOrderResponse.md)
 - [ListMessagingProfileMetricsResponse](docs/ListMessagingProfileMetricsResponse.md)
 - [ListMessagingProfilePhoneNumbersResponse](docs/ListMessagingProfilePhoneNumbersResponse.md)
 - [ListMessagingProfileShortCodesResponse](docs/ListMessagingProfileShortCodesResponse.md)
 - [ListMessagingProfileUrlDomainsResponse](docs/ListMessagingProfileUrlDomainsResponse.md)
 - [ListMessagingProfilesResponse](docs/ListMessagingProfilesResponse.md)
 - [ListMessagingSettingsResponse](docs/ListMessagingSettingsResponse.md)
 - [ListMigrationSourceCoverage200Response](docs/ListMigrationSourceCoverage200Response.md)
 - [ListMigrationSources200Response](docs/ListMigrationSources200Response.md)
 - [ListMigrations200Response](docs/ListMigrations200Response.md)
 - [ListNetworkCoverage200Response](docs/ListNetworkCoverage200Response.md)
 - [ListNetworkInterfaces200Response](docs/ListNetworkInterfaces200Response.md)
 - [ListNetworks200Response](docs/ListNetworks200Response.md)
 - [ListNotificationChannels200Response](docs/ListNotificationChannels200Response.md)
 - [ListNotificationSettings200Response](docs/ListNotificationSettings200Response.md)
 - [ListNumberBlockOrdersResponse](docs/ListNumberBlockOrdersResponse.md)
 - [ListNumberOrderPhoneNumbersResponse](docs/ListNumberOrderPhoneNumbersResponse.md)
 - [ListNumberOrdersResponse](docs/ListNumberOrdersResponse.md)
 - [ListNumberReservationsResponse](docs/ListNumberReservationsResponse.md)
 - [ListObjectsResponse](docs/ListObjectsResponse.md)
 - [ListObjectsResponseContentsInner](docs/ListObjectsResponseContentsInner.md)
 - [ListOfMediaResourcesResponse](docs/ListOfMediaResourcesResponse.md)
 - [ListOtaUpdates200Response](docs/ListOtaUpdates200Response.md)
 - [ListOutboundVoiceProfilesResponse](docs/ListOutboundVoiceProfilesResponse.md)
 - [ListParticipantsResponse](docs/ListParticipantsResponse.md)
 - [ListPhoneNumberBlocksBackgroundJobsResponse](docs/ListPhoneNumberBlocksBackgroundJobsResponse.md)
 - [ListPhoneNumberConfigurations200Response](docs/ListPhoneNumberConfigurations200Response.md)
 - [ListPhoneNumbersBackgroundJobsResponse](docs/ListPhoneNumbersBackgroundJobsResponse.md)
 - [ListPhoneNumbersFilterCountryIsoAlpha2Parameter](docs/ListPhoneNumbersFilterCountryIsoAlpha2Parameter.md)
 - [ListPhoneNumbersResponse](docs/ListPhoneNumbersResponse.md)
 - [ListPhoneNumbersResponse1](docs/ListPhoneNumbersResponse1.md)
 - [ListPhoneNumbersWithVoiceSettingsResponse](docs/ListPhoneNumbersWithVoiceSettingsResponse.md)
 - [ListPortingEvents200Response](docs/ListPortingEvents200Response.md)
 - [ListPortingOrderActivationJobs200Response](docs/ListPortingOrderActivationJobs200Response.md)
 - [ListPortingOrderComments200Response](docs/ListPortingOrderComments200Response.md)
 - [ListPortingOrderRequirements200Response](docs/ListPortingOrderRequirements200Response.md)
 - [ListPortingOrders200Response](docs/ListPortingOrders200Response.md)
 - [ListPortingPhoneNumberBlocks200Response](docs/ListPortingPhoneNumberBlocks200Response.md)
 - [ListPortingPhoneNumberExtensions200Response](docs/ListPortingPhoneNumberExtensions200Response.md)
 - [ListPortingPhoneNumbers200Response](docs/ListPortingPhoneNumbers200Response.md)
 - [ListPortingReports200Response](docs/ListPortingReports200Response.md)
 - [ListPortoutEvents200Response](docs/ListPortoutEvents200Response.md)
 - [ListPortoutRejections200Response](docs/ListPortoutRejections200Response.md)
 - [ListPortoutReports200Response](docs/ListPortoutReports200Response.md)
 - [ListPortoutRequest200Response](docs/ListPortoutRequest200Response.md)
 - [ListPublicInternetGateways200Response](docs/ListPublicInternetGateways200Response.md)
 - [ListPushCredentialsResponse](docs/ListPushCredentialsResponse.md)
 - [ListQueueCallsResponse](docs/ListQueueCallsResponse.md)
 - [ListRecordingTranscriptionsResponse](docs/ListRecordingTranscriptionsResponse.md)
 - [ListRegions200Response](docs/ListRegions200Response.md)
 - [ListRegulatoryRequirements200Response](docs/ListRegulatoryRequirements200Response.md)
 - [ListRegulatoryRequirementsPhoneNumbers200Response](docs/ListRegulatoryRequirementsPhoneNumbers200Response.md)
 - [ListReleasesResponse](docs/ListReleasesResponse.md)
 - [ListRequirementTypes200Response](docs/ListRequirementTypes200Response.md)
 - [ListRequirements200Response](docs/ListRequirements200Response.md)
 - [ListRoomCompositions200Response](docs/ListRoomCompositions200Response.md)
 - [ListRoomParticipants200Response](docs/ListRoomParticipants200Response.md)
 - [ListRoomRecordings200Response](docs/ListRoomRecordings200Response.md)
 - [ListRoomSessions200Response](docs/ListRoomSessions200Response.md)
 - [ListRooms200Response](docs/ListRooms200Response.md)
 - [ListShortCodesResponse](docs/ListShortCodesResponse.md)
 - [ListSimCardActions200Response](docs/ListSimCardActions200Response.md)
 - [ListSubNumberOrdersResponse](docs/ListSubNumberOrdersResponse.md)
 - [ListTextToSpeechVoices200Response](docs/ListTextToSpeechVoices200Response.md)
 - [ListTextToSpeechVoices200ResponseVoicesInner](docs/ListTextToSpeechVoices200ResponseVoicesInner.md)
 - [ListUploadsResponse](docs/ListUploadsResponse.md)
 - [ListVerificationCodes200Response](docs/ListVerificationCodes200Response.md)
 - [ListVerificationsResponse](docs/ListVerificationsResponse.md)
 - [ListVerifiedNumbersResponse](docs/ListVerifiedNumbersResponse.md)
 - [ListVerifyProfileMessageTemplateResponse](docs/ListVerifyProfileMessageTemplateResponse.md)
 - [ListVerifyProfilesResponse](docs/ListVerifyProfilesResponse.md)
 - [ListVirtualCrossConnectCoverage200Response](docs/ListVirtualCrossConnectCoverage200Response.md)
 - [ListVirtualCrossConnects200Response](docs/ListVirtualCrossConnects200Response.md)
 - [ListWireguardInterfaces200Response](docs/ListWireguardInterfaces200Response.md)
 - [ListWireguardPeers200Response](docs/ListWireguardPeers200Response.md)
 - [Location](docs/Location.md)
 - [Location1Inner](docs/Location1Inner.md)
 - [LocationResponse](docs/LocationResponse.md)
 - [LocationResponseData](docs/LocationResponseData.md)
 - [LogMessage](docs/LogMessage.md)
 - [LogMessageMeta](docs/LogMessageMeta.md)
 - [LogMessageSource](docs/LogMessageSource.md)
 - [Loopcount](docs/Loopcount.md)
 - [ManagedAccount](docs/ManagedAccount.md)
 - [ManagedAccountBalance](docs/ManagedAccountBalance.md)
 - [ManagedAccountMultiListing](docs/ManagedAccountMultiListing.md)
 - [ManagedAccountsGlobalOutboundChannels](docs/ManagedAccountsGlobalOutboundChannels.md)
 - [MdrDeleteUsageReportsResponse](docs/MdrDeleteUsageReportsResponse.md)
 - [MdrDetailResponse](docs/MdrDetailResponse.md)
 - [MdrGetDetailResponse](docs/MdrGetDetailResponse.md)
 - [MdrGetSyncUsageReportResponse](docs/MdrGetSyncUsageReportResponse.md)
 - [MdrGetUsageReportsByIdResponse](docs/MdrGetUsageReportsByIdResponse.md)
 - [MdrGetUsageReportsResponse](docs/MdrGetUsageReportsResponse.md)
 - [MdrPostUsageReportRequest](docs/MdrPostUsageReportRequest.md)
 - [MdrPostUsageReportsResponse](docs/MdrPostUsageReportsResponse.md)
 - [MdrUsageRecord](docs/MdrUsageRecord.md)
 - [MdrUsageReportResponse](docs/MdrUsageReportResponse.md)
 - [MediaFeatures](docs/MediaFeatures.md)
 - [MediaResource](docs/MediaResource.md)
 - [MediaResourceResponse](docs/MediaResourceResponse.md)
 - [MediaStorageDetailRecord](docs/MediaStorageDetailRecord.md)
 - [MessageDetailRecord](docs/MessageDetailRecord.md)
 - [MessageResponse](docs/MessageResponse.md)
 - [MessagingFeatureSet](docs/MessagingFeatureSet.md)
 - [MessagingHostedNumberOrder](docs/MessagingHostedNumberOrder.md)
 - [MessagingProfile](docs/MessagingProfile.md)
 - [MessagingProfileDetailedMetric](docs/MessagingProfileDetailedMetric.md)
 - [MessagingProfileDetailedMetrics](docs/MessagingProfileDetailedMetrics.md)
 - [MessagingProfileHighLevelMetrics](docs/MessagingProfileHighLevelMetrics.md)
 - [MessagingProfileHighLevelMetricsInbound](docs/MessagingProfileHighLevelMetricsInbound.md)
 - [MessagingProfileHighLevelMetricsOutbound](docs/MessagingProfileHighLevelMetricsOutbound.md)
 - [MessagingProfileMessageTypeMetrics](docs/MessagingProfileMessageTypeMetrics.md)
 - [MessagingProfileResponse](docs/MessagingProfileResponse.md)
 - [MessagingSettings](docs/MessagingSettings.md)
 - [MessagingUrlDomain](docs/MessagingUrlDomain.md)
 - [Meta](docs/Meta.md)
 - [MetaResponse](docs/MetaResponse.md)
 - [Metadata](docs/Metadata.md)
 - [MigrationParams](docs/MigrationParams.md)
 - [MigrationSourceCoverageParams](docs/MigrationSourceCoverageParams.md)
 - [MigrationSourceParams](docs/MigrationSourceParams.md)
 - [MigrationSourceParamsProviderAuth](docs/MigrationSourceParamsProviderAuth.md)
 - [MnoMetadata](docs/MnoMetadata.md)
 - [MnoMetadataItem](docs/MnoMetadataItem.md)
 - [MobileNetworkOperator](docs/MobileNetworkOperator.md)
 - [MobileNetworkOperatorPreferencesResponse](docs/MobileNetworkOperatorPreferencesResponse.md)
 - [ModelMetadata](docs/ModelMetadata.md)
 - [ModelsResponse](docs/ModelsResponse.md)
 - [Network](docs/Network.md)
 - [NetworkCoverage](docs/NetworkCoverage.md)
 - [NetworkCoverageAvailableServicesInner](docs/NetworkCoverageAvailableServicesInner.md)
 - [NetworkCreate](docs/NetworkCreate.md)
 - [NetworkInterface](docs/NetworkInterface.md)
 - [NewBillingGroup](docs/NewBillingGroup.md)
 - [NewLedgerBillingGroupReport](docs/NewLedgerBillingGroupReport.md)
 - [NewParticipantResource](docs/NewParticipantResource.md)
 - [NoiseSuppressionDirection](docs/NoiseSuppressionDirection.md)
 - [NoiseSuppressionStart](docs/NoiseSuppressionStart.md)
 - [NoiseSuppressionStop](docs/NoiseSuppressionStop.md)
 - [NotFoundError](docs/NotFoundError.md)
 - [NotificationChannel](docs/NotificationChannel.md)
 - [NotificationEvent](docs/NotificationEvent.md)
 - [NotificationEventCondition](docs/NotificationEventCondition.md)
 - [NotificationEventConditionParametersInner](docs/NotificationEventConditionParametersInner.md)
 - [NotificationProfile](docs/NotificationProfile.md)
 - [NotificationSetting](docs/NotificationSetting.md)
 - [NotificationSettingParametersInner](docs/NotificationSettingParametersInner.md)
 - [NumberBlockOrder](docs/NumberBlockOrder.md)
 - [NumberBlockOrderResponse](docs/NumberBlockOrderResponse.md)
 - [NumberHealthMetrics](docs/NumberHealthMetrics.md)
 - [NumberLookupRecord](docs/NumberLookupRecord.md)
 - [NumberLookupResponse](docs/NumberLookupResponse.md)
 - [NumberOrder](docs/NumberOrder.md)
 - [NumberOrderBlockEvent](docs/NumberOrderBlockEvent.md)
 - [NumberOrderPhoneNumber](docs/NumberOrderPhoneNumber.md)
 - [NumberOrderPhoneNumberRequirementGroupResponse](docs/NumberOrderPhoneNumberRequirementGroupResponse.md)
 - [NumberOrderPhoneNumberRequirementGroupResponseRegulatoryRequirementsInner](docs/NumberOrderPhoneNumberRequirementGroupResponseRegulatoryRequirementsInner.md)
 - [NumberOrderPhoneNumberResponse](docs/NumberOrderPhoneNumberResponse.md)
 - [NumberOrderResponse](docs/NumberOrderResponse.md)
 - [NumberOrderWithPhoneNumbers](docs/NumberOrderWithPhoneNumbers.md)
 - [NumberPoolSettings](docs/NumberPoolSettings.md)
 - [NumberReservation](docs/NumberReservation.md)
 - [NumberReservationResponse](docs/NumberReservationResponse.md)
 - [OperatorConnectRefreshResponse](docs/OperatorConnectRefreshResponse.md)
 - [OperatorConnectRefreshResponse1](docs/OperatorConnectRefreshResponse1.md)
 - [OptOutItem](docs/OptOutItem.md)
 - [OptOutListResponse](docs/OptOutListResponse.md)
 - [OrderExternalVetting](docs/OrderExternalVetting.md)
 - [OutboundCallRecording](docs/OutboundCallRecording.md)
 - [OutboundFqdn](docs/OutboundFqdn.md)
 - [OutboundIp](docs/OutboundIp.md)
 - [OutboundMessage](docs/OutboundMessage.md)
 - [OutboundMessageEvent](docs/OutboundMessageEvent.md)
 - [OutboundMessageEventMeta](docs/OutboundMessageEventMeta.md)
 - [OutboundMessagePayload](docs/OutboundMessagePayload.md)
 - [OutboundMessagePayloadFrom](docs/OutboundMessagePayloadFrom.md)
 - [OutboundMessagePayloadMediaInner](docs/OutboundMessagePayloadMediaInner.md)
 - [OutboundMessagePayloadToInner](docs/OutboundMessagePayloadToInner.md)
 - [OutboundVoiceProfile](docs/OutboundVoiceProfile.md)
 - [OutboundVoiceProfileResponse](docs/OutboundVoiceProfileResponse.md)
 - [PWGAssignedResourcesSummary](docs/PWGAssignedResourcesSummary.md)
 - [PaginatedBillingBundlesResponse](docs/PaginatedBillingBundlesResponse.md)
 - [PaginatedScheduledEventList](docs/PaginatedScheduledEventList.md)
 - [PaginatedUserBundlesResponse](docs/PaginatedUserBundlesResponse.md)
 - [PaginatedVerificationRequestStatus](docs/PaginatedVerificationRequestStatus.md)
 - [PaginationData](docs/PaginationData.md)
 - [PaginationMeta](docs/PaginationMeta.md)
 - [PaginationMetaSimple](docs/PaginationMetaSimple.md)
 - [PaginationResponse](docs/PaginationResponse.md)
 - [Participant](docs/Participant.md)
 - [ParticipantConference](docs/ParticipantConference.md)
 - [ParticipantResource](docs/ParticipantResource.md)
 - [ParticipantResourceIndex](docs/ParticipantResourceIndex.md)
 - [PatchChannelZoneRequest](docs/PatchChannelZoneRequest.md)
 - [PatchRoomRequest](docs/PatchRoomRequest.md)
 - [PauseConferenceRecordingRequest](docs/PauseConferenceRecordingRequest.md)
 - [PauseRecordingRequest](docs/PauseRecordingRequest.md)
 - [PhoneNumberBlocksJob](docs/PhoneNumberBlocksJob.md)
 - [PhoneNumberBlocksJobDeletePhoneNumberBlock](docs/PhoneNumberBlocksJobDeletePhoneNumberBlock.md)
 - [PhoneNumberBlocksJobDeletePhoneNumberBlockRequest](docs/PhoneNumberBlocksJobDeletePhoneNumberBlockRequest.md)
 - [PhoneNumberBlocksJobFailedOperation](docs/PhoneNumberBlocksJobFailedOperation.md)
 - [PhoneNumberBlocksJobSuccessfulOperation](docs/PhoneNumberBlocksJobSuccessfulOperation.md)
 - [PhoneNumberBundleStatusChange](docs/PhoneNumberBundleStatusChange.md)
 - [PhoneNumberBundleStatusChangeRequest](docs/PhoneNumberBundleStatusChangeRequest.md)
 - [PhoneNumberCampaign](docs/PhoneNumberCampaign.md)
 - [PhoneNumberCampaignCreate](docs/PhoneNumberCampaignCreate.md)
 - [PhoneNumberCampaignPaginated](docs/PhoneNumberCampaignPaginated.md)
 - [PhoneNumberDeletedDetailed](docs/PhoneNumberDeletedDetailed.md)
 - [PhoneNumberDetailed](docs/PhoneNumberDetailed.md)
 - [PhoneNumberEnableEmergency](docs/PhoneNumberEnableEmergency.md)
 - [PhoneNumberEnableEmergencyRequest](docs/PhoneNumberEnableEmergencyRequest.md)
 - [PhoneNumberResponse](docs/PhoneNumberResponse.md)
 - [PhoneNumberResponse1](docs/PhoneNumberResponse1.md)
 - [PhoneNumberStatusResponsePaginated](docs/PhoneNumberStatusResponsePaginated.md)
 - [PhoneNumberWithMessagingSettings](docs/PhoneNumberWithMessagingSettings.md)
 - [PhoneNumberWithMessagingSettingsFeatures](docs/PhoneNumberWithMessagingSettingsFeatures.md)
 - [PhoneNumberWithVoiceSettings](docs/PhoneNumberWithVoiceSettings.md)
 - [PhoneNumbersEnableEmergency](docs/PhoneNumbersEnableEmergency.md)
 - [PhoneNumbersJob](docs/PhoneNumbersJob.md)
 - [PhoneNumbersJobDeletePhoneNumbers](docs/PhoneNumbersJobDeletePhoneNumbers.md)
 - [PhoneNumbersJobDeletePhoneNumbersRequest](docs/PhoneNumbersJobDeletePhoneNumbersRequest.md)
 - [PhoneNumbersJobFailedOperation](docs/PhoneNumbersJobFailedOperation.md)
 - [PhoneNumbersJobPendingOperation](docs/PhoneNumbersJobPendingOperation.md)
 - [PhoneNumbersJobPhoneNumber](docs/PhoneNumbersJobPhoneNumber.md)
 - [PhoneNumbersJobSuccessfulOperation](docs/PhoneNumbersJobSuccessfulOperation.md)
 - [PhoneNumbersJobUpdateEmergencySettingsRequest](docs/PhoneNumbersJobUpdateEmergencySettingsRequest.md)
 - [PhoneNumbersJobUpdatePhoneNumbers](docs/PhoneNumbersJobUpdatePhoneNumbers.md)
 - [PhoneNumbersJobUpdatePhoneNumbersRequest](docs/PhoneNumbersJobUpdatePhoneNumbersRequest.md)
 - [PlayAudioUrlRequest](docs/PlayAudioUrlRequest.md)
 - [PlaybackStopRequest](docs/PlaybackStopRequest.md)
 - [PortOutSupportingDocument](docs/PortOutSupportingDocument.md)
 - [Portability](docs/Portability.md)
 - [PortabilityCheckDetails](docs/PortabilityCheckDetails.md)
 - [PortabilityStatus](docs/PortabilityStatus.md)
 - [PortingAdditionalDocument](docs/PortingAdditionalDocument.md)
 - [PortingEvent](docs/PortingEvent.md)
 - [PortingEventPayload](docs/PortingEventPayload.md)
 - [PortingLOAConfiguration](docs/PortingLOAConfiguration.md)
 - [PortingLOAConfigurationAddress](docs/PortingLOAConfigurationAddress.md)
 - [PortingLOAConfigurationContact](docs/PortingLOAConfigurationContact.md)
 - [PortingLOAConfigurationLogo](docs/PortingLOAConfigurationLogo.md)
 - [PortingOrder](docs/PortingOrder.md)
 - [PortingOrderActivationSettings](docs/PortingOrderActivationSettings.md)
 - [PortingOrderActivationStatus](docs/PortingOrderActivationStatus.md)
 - [PortingOrderDocuments](docs/PortingOrderDocuments.md)
 - [PortingOrderEndUser](docs/PortingOrderEndUser.md)
 - [PortingOrderEndUserAdmin](docs/PortingOrderEndUserAdmin.md)
 - [PortingOrderEndUserLocation](docs/PortingOrderEndUserLocation.md)
 - [PortingOrderMessaging](docs/PortingOrderMessaging.md)
 - [PortingOrderMisc](docs/PortingOrderMisc.md)
 - [PortingOrderPhoneNumberConfiguration](docs/PortingOrderPhoneNumberConfiguration.md)
 - [PortingOrderRequirement](docs/PortingOrderRequirement.md)
 - [PortingOrderRequirementDetail](docs/PortingOrderRequirementDetail.md)
 - [PortingOrderRequirementDetailRequirementType](docs/PortingOrderRequirementDetailRequirementType.md)
 - [PortingOrderSharingToken](docs/PortingOrderSharingToken.md)
 - [PortingOrderStatus](docs/PortingOrderStatus.md)
 - [PortingOrderType](docs/PortingOrderType.md)
 - [PortingOrderUserFeedback](docs/PortingOrderUserFeedback.md)
 - [PortingOrdersActivationJob](docs/PortingOrdersActivationJob.md)
 - [PortingOrdersActivationJobActivationWindowsInner](docs/PortingOrdersActivationJobActivationWindowsInner.md)
 - [PortingOrdersAllowedFocWindow](docs/PortingOrdersAllowedFocWindow.md)
 - [PortingOrdersComment](docs/PortingOrdersComment.md)
 - [PortingOrdersExceptionType](docs/PortingOrdersExceptionType.md)
 - [PortingPhoneNumber](docs/PortingPhoneNumber.md)
 - [PortingPhoneNumberBlock](docs/PortingPhoneNumberBlock.md)
 - [PortingPhoneNumberBlockActivationRangesInner](docs/PortingPhoneNumberBlockActivationRangesInner.md)
 - [PortingPhoneNumberBlockPhoneNumberRange](docs/PortingPhoneNumberBlockPhoneNumberRange.md)
 - [PortingPhoneNumberConfiguration](docs/PortingPhoneNumberConfiguration.md)
 - [PortingPhoneNumberExtension](docs/PortingPhoneNumberExtension.md)
 - [PortingPhoneNumberExtensionActivationRangesInner](docs/PortingPhoneNumberExtensionActivationRangesInner.md)
 - [PortingPhoneNumberExtensionExtensionRange](docs/PortingPhoneNumberExtensionExtensionRange.md)
 - [PortingReport](docs/PortingReport.md)
 - [PortingVerificationCode](docs/PortingVerificationCode.md)
 - [PortoutComment](docs/PortoutComment.md)
 - [PortoutDetails](docs/PortoutDetails.md)
 - [PortoutEvent](docs/PortoutEvent.md)
 - [PortoutEventPayload](docs/PortoutEventPayload.md)
 - [PortoutRejection](docs/PortoutRejection.md)
 - [PortoutReport](docs/PortoutReport.md)
 - [PostNumbersFeatures200Response](docs/PostNumbersFeatures200Response.md)
 - [PostNumbersFeatures200ResponseDataInner](docs/PostNumbersFeatures200ResponseDataInner.md)
 - [PostNumbersFeaturesRequest](docs/PostNumbersFeaturesRequest.md)
 - [PostPortRequestComment201Response](docs/PostPortRequestComment201Response.md)
 - [PostPortRequestCommentRequest](docs/PostPortRequestCommentRequest.md)
 - [PostPortRequestSupportingDocumentsRequest](docs/PostPortRequestSupportingDocumentsRequest.md)
 - [PostPortRequestSupportingDocumentsRequestDocumentsInner](docs/PostPortRequestSupportingDocumentsRequestDocumentsInner.md)
 - [PostPortabilityCheck201Response](docs/PostPortabilityCheck201Response.md)
 - [PostPortabilityCheckRequest](docs/PostPortabilityCheckRequest.md)
 - [PostSimCardDataUsageNotification201Response](docs/PostSimCardDataUsageNotification201Response.md)
 - [PostSimCardDataUsageNotificationRequest](docs/PostSimCardDataUsageNotificationRequest.md)
 - [PostSimCardDataUsageNotificationRequestThreshold](docs/PostSimCardDataUsageNotificationRequestThreshold.md)
 - [PresignedObjectUrl](docs/PresignedObjectUrl.md)
 - [PresignedObjectUrlContent](docs/PresignedObjectUrlContent.md)
 - [PresignedObjectUrlParams](docs/PresignedObjectUrlParams.md)
 - [PreviewLoaConfigurationParamsRequest](docs/PreviewLoaConfigurationParamsRequest.md)
 - [PreviewLoaConfigurationParamsRequestAddress](docs/PreviewLoaConfigurationParamsRequestAddress.md)
 - [PreviewLoaConfigurationParamsRequestContact](docs/PreviewLoaConfigurationParamsRequestContact.md)
 - [PreviewLoaConfigurationParamsRequestLogo](docs/PreviewLoaConfigurationParamsRequestLogo.md)
 - [PreviewSimCardOrders202Response](docs/PreviewSimCardOrders202Response.md)
 - [PreviewSimCardOrdersRequest](docs/PreviewSimCardOrdersRequest.md)
 - [PrivacySettings](docs/PrivacySettings.md)
 - [PrivateWirelessGateway](docs/PrivateWirelessGateway.md)
 - [PrivateWirelessGatewayStatus](docs/PrivateWirelessGatewayStatus.md)
 - [ProfileAssignmentPhoneNumbers](docs/ProfileAssignmentPhoneNumbers.md)
 - [PublicInternetGateway](docs/PublicInternetGateway.md)
 - [PublicInternetGatewayCreate](docs/PublicInternetGatewayCreate.md)
 - [PublicInternetGatewayRead](docs/PublicInternetGatewayRead.md)
 - [PublicTextClusteringRequest](docs/PublicTextClusteringRequest.md)
 - [PurchaseESim202Response](docs/PurchaseESim202Response.md)
 - [PushCredential](docs/PushCredential.md)
 - [PushCredentialResponse](docs/PushCredentialResponse.md)
 - [Quality](docs/Quality.md)
 - [Queue](docs/Queue.md)
 - [QueueCall](docs/QueueCall.md)
 - [QueueCallResponse](docs/QueueCallResponse.md)
 - [QueueResponse](docs/QueueResponse.md)
 - [RCSAction](docs/RCSAction.md)
 - [RCSAgentMessage](docs/RCSAgentMessage.md)
 - [RCSCardContent](docs/RCSCardContent.md)
 - [RCSCarouselCard](docs/RCSCarouselCard.md)
 - [RCSComposeAction](docs/RCSComposeAction.md)
 - [RCSComposeRecordingMessage](docs/RCSComposeRecordingMessage.md)
 - [RCSComposeTextMessage](docs/RCSComposeTextMessage.md)
 - [RCSContentInfo](docs/RCSContentInfo.md)
 - [RCSContentMessage](docs/RCSContentMessage.md)
 - [RCSCreateCalendarEventAction](docs/RCSCreateCalendarEventAction.md)
 - [RCSDialAction](docs/RCSDialAction.md)
 - [RCSEvent](docs/RCSEvent.md)
 - [RCSFrom](docs/RCSFrom.md)
 - [RCSLatLng](docs/RCSLatLng.md)
 - [RCSMedia](docs/RCSMedia.md)
 - [RCSMessage](docs/RCSMessage.md)
 - [RCSOpenUrlAction](docs/RCSOpenUrlAction.md)
 - [RCSReply](docs/RCSReply.md)
 - [RCSResponse](docs/RCSResponse.md)
 - [RCSResponseData](docs/RCSResponseData.md)
 - [RCSRichCard](docs/RCSRichCard.md)
 - [RCSStandaloneCard](docs/RCSStandaloneCard.md)
 - [RCSSuggestion](docs/RCSSuggestion.md)
 - [RCSToItem](docs/RCSToItem.md)
 - [RCSViewLocationAction](docs/RCSViewLocationAction.md)
 - [Record](docs/Record.md)
 - [RecordType](docs/RecordType.md)
 - [RecordingResponse](docs/RecordingResponse.md)
 - [RecordingResponseData](docs/RecordingResponseData.md)
 - [RecordingResponseDataDownloadUrls](docs/RecordingResponseDataDownloadUrls.md)
 - [RecordingSource](docs/RecordingSource.md)
 - [RecordingTrack](docs/RecordingTrack.md)
 - [RecordingTranscription](docs/RecordingTranscription.md)
 - [RecursiveCluster](docs/RecursiveCluster.md)
 - [ReferRequest](docs/ReferRequest.md)
 - [RefreshFaxResponse](docs/RefreshFaxResponse.md)
 - [RefreshRoomClientToken201Response](docs/RefreshRoomClientToken201Response.md)
 - [RefreshRoomClientToken201ResponseData](docs/RefreshRoomClientToken201ResponseData.md)
 - [RefreshRoomClientTokenRequest](docs/RefreshRoomClientTokenRequest.md)
 - [Region](docs/Region.md)
 - [RegionIn](docs/RegionIn.md)
 - [RegionInformation](docs/RegionInformation.md)
 - [RegionOut](docs/RegionOut.md)
 - [RegionOutRegion](docs/RegionOutRegion.md)
 - [RegistrationStatus](docs/RegistrationStatus.md)
 - [RegistrationStatusResponse](docs/RegistrationStatusResponse.md)
 - [RegulatoryRequirement](docs/RegulatoryRequirement.md)
 - [RegulatoryRequirements](docs/RegulatoryRequirements.md)
 - [RegulatoryRequirementsPhoneNumbers](docs/RegulatoryRequirementsPhoneNumbers.md)
 - [RegulatoryRequirementsPhoneNumbersRegionInformationInner](docs/RegulatoryRequirementsPhoneNumbersRegionInformationInner.md)
 - [RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInner](docs/RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInner.md)
 - [RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria](docs/RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria.md)
 - [RegulatoryRequirementsRegulatoryRequirementsInner](docs/RegulatoryRequirementsRegulatoryRequirementsInner.md)
 - [RegulatoryRequirementsRegulatoryRequirementsInnerAcceptanceCriteria](docs/RegulatoryRequirementsRegulatoryRequirementsInnerAcceptanceCriteria.md)
 - [RejectRequest](docs/RejectRequest.md)
 - [Release](docs/Release.md)
 - [ReplacedLinkClick](docs/ReplacedLinkClick.md)
 - [ReplacedLinkClickEvent](docs/ReplacedLinkClickEvent.md)
 - [RequirementGroup](docs/RequirementGroup.md)
 - [ReservedPhoneNumber](docs/ReservedPhoneNumber.md)
 - [ResourceNotFoundError](docs/ResourceNotFoundError.md)
 - [ResourceNotFoundErrorErrorsInner](docs/ResourceNotFoundErrorErrorsInner.md)
 - [ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost](docs/ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost.md)
 - [ResponseSubmitCampaignPublicCampaignbuilderPost](docs/ResponseSubmitCampaignPublicCampaignbuilderPost.md)
 - [ResumeConferenceRecordingRequest](docs/ResumeConferenceRecordingRequest.md)
 - [ResumeRecordingRequest](docs/ResumeRecordingRequest.md)
 - [RetreiveCountryCoverage200Response](docs/RetreiveCountryCoverage200Response.md)
 - [RetreiveSpecificCountryCoverage200Response](docs/RetreiveSpecificCountryCoverage200Response.md)
 - [Retrieval](docs/Retrieval.md)
 - [RetrievalTool](docs/RetrievalTool.md)
 - [RetrieveBulkUpdateMessagingSettingsResponse](docs/RetrieveBulkUpdateMessagingSettingsResponse.md)
 - [RetrieveCallStatusResponse](docs/RetrieveCallStatusResponse.md)
 - [RetrieveCallStatusResponseWithRecordingID](docs/RetrieveCallStatusResponseWithRecordingID.md)
 - [RetrieveDocumentRequirements200Response](docs/RetrieveDocumentRequirements200Response.md)
 - [RetrieveMessagingHostedNumberOrderResponse](docs/RetrieveMessagingHostedNumberOrderResponse.md)
 - [RetrieveMessagingHostedNumberResponse](docs/RetrieveMessagingHostedNumberResponse.md)
 - [RetrieveMessagingHostedNumberResponse1](docs/RetrieveMessagingHostedNumberResponse1.md)
 - [RetrieveMessagingProfileMetricsResponse](docs/RetrieveMessagingProfileMetricsResponse.md)
 - [RetrieveMessagingSettingsResponse](docs/RetrieveMessagingSettingsResponse.md)
 - [RetrievePhoneNumberVoiceResponse](docs/RetrievePhoneNumberVoiceResponse.md)
 - [RetrieveRequirementType200Response](docs/RetrieveRequirementType200Response.md)
 - [RetrieveVerificationResponse](docs/RetrieveVerificationResponse.md)
 - [Room](docs/Room.md)
 - [RoomComposition](docs/RoomComposition.md)
 - [RoomParticipant](docs/RoomParticipant.md)
 - [RoomRecording](docs/RoomRecording.md)
 - [RoomSession](docs/RoomSession.md)
 - [S3ConfigurationData](docs/S3ConfigurationData.md)
 - [SIMCard](docs/SIMCard.md)
 - [SIMCardAction](docs/SIMCardAction.md)
 - [SIMCardActionStatus](docs/SIMCardActionStatus.md)
 - [SIMCardActionsSummary](docs/SIMCardActionsSummary.md)
 - [SIMCardActivationCode](docs/SIMCardActivationCode.md)
 - [SIMCardCurrentBillingPeriodConsumedData](docs/SIMCardCurrentBillingPeriodConsumedData.md)
 - [SIMCardCurrentDeviceLocation](docs/SIMCardCurrentDeviceLocation.md)
 - [SIMCardDataLimit](docs/SIMCardDataLimit.md)
 - [SIMCardDeviceDetails](docs/SIMCardDeviceDetails.md)
 - [SIMCardGroup](docs/SIMCardGroup.md)
 - [SIMCardGroupAction](docs/SIMCardGroupAction.md)
 - [SIMCardGroupActionSettings](docs/SIMCardGroupActionSettings.md)
 - [SIMCardGroupCreate](docs/SIMCardGroupCreate.md)
 - [SIMCardGroupDataLimit](docs/SIMCardGroupDataLimit.md)
 - [SIMCardGroupPatch](docs/SIMCardGroupPatch.md)
 - [SIMCardOrder](docs/SIMCardOrder.md)
 - [SIMCardOrderCost](docs/SIMCardOrderCost.md)
 - [SIMCardOrderOrderAddress](docs/SIMCardOrderOrderAddress.md)
 - [SIMCardOrderPreview](docs/SIMCardOrderPreview.md)
 - [SIMCardOrderPreviewTotalCost](docs/SIMCardOrderPreviewTotalCost.md)
 - [SIMCardPublicIP](docs/SIMCardPublicIP.md)
 - [SIMCardRegistration](docs/SIMCardRegistration.md)
 - [SIMCardRegistrationCodeValidation](docs/SIMCardRegistrationCodeValidation.md)
 - [SIMCardRegistrationCodeValidations](docs/SIMCardRegistrationCodeValidations.md)
 - [SIMCardStatus](docs/SIMCardStatus.md)
 - [SIPRECConnector](docs/SIPRECConnector.md)
 - [SMSFallback](docs/SMSFallback.md)
 - [SSLCertificate](docs/SSLCertificate.md)
 - [SSLCertificateIssuedBy](docs/SSLCertificateIssuedBy.md)
 - [SSLCertificateIssuedTo](docs/SSLCertificateIssuedTo.md)
 - [ScheduledEventResponse](docs/ScheduledEventResponse.md)
 - [ScheduledPhoneCallEventResponse](docs/ScheduledPhoneCallEventResponse.md)
 - [ScheduledSmsEventResponse](docs/ScheduledSmsEventResponse.md)
 - [SearchedSIMCardGroup](docs/SearchedSIMCardGroup.md)
 - [SendDTMFRequest](docs/SendDTMFRequest.md)
 - [SendFaxRequest](docs/SendFaxRequest.md)
 - [SendFaxResponse](docs/SendFaxResponse.md)
 - [SendPortingVerificationCodesRequest](docs/SendPortingVerificationCodesRequest.md)
 - [SendSIPInfoRequest](docs/SendSIPInfoRequest.md)
 - [ServicePlan](docs/ServicePlan.md)
 - [SetPrivateWirelessGatewayForSimCardGroupRequest](docs/SetPrivateWirelessGatewayForSimCardGroupRequest.md)
 - [SetPublicIPsBulk202Response](docs/SetPublicIPsBulk202Response.md)
 - [SetPublicIPsBulkRequest](docs/SetPublicIPsBulkRequest.md)
 - [Settings](docs/Settings.md)
 - [SettingsDataErrorMessage](docs/SettingsDataErrorMessage.md)
 - [SharePortingOrder201Response](docs/SharePortingOrder201Response.md)
 - [SharePortingOrderRequest](docs/SharePortingOrderRequest.md)
 - [SharedCampaign](docs/SharedCampaign.md)
 - [SharedCampaignRecordSet](docs/SharedCampaignRecordSet.md)
 - [ShortCode](docs/ShortCode.md)
 - [ShortCodeResponse](docs/ShortCodeResponse.md)
 - [ShowPortingEvent200Response](docs/ShowPortingEvent200Response.md)
 - [ShowPortoutEvent200Response](docs/ShowPortoutEvent200Response.md)
 - [SimCardDataUsageNotification](docs/SimCardDataUsageNotification.md)
 - [SimCardOrderCreate](docs/SimCardOrderCreate.md)
 - [SimCardUsageDetailRecord](docs/SimCardUsageDetailRecord.md)
 - [SimpleSIMCard](docs/SimpleSIMCard.md)
 - [SimpleSIMCardDataLimit](docs/SimpleSIMCardDataLimit.md)
 - [SimplifiedOTAUpdate](docs/SimplifiedOTAUpdate.md)
 - [SingleManagedAccountGlobalOutboundChannels](docs/SingleManagedAccountGlobalOutboundChannels.md)
 - [SipHeader](docs/SipHeader.md)
 - [SiprecConnector](docs/SiprecConnector.md)
 - [SiprecConnectorResponse](docs/SiprecConnectorResponse.md)
 - [SlimPhoneNumberDetailed](docs/SlimPhoneNumberDetailed.md)
 - [SoundModifications](docs/SoundModifications.md)
 - [SourceResponse](docs/SourceResponse.md)
 - [SpeakRequest](docs/SpeakRequest.md)
 - [StartConferenceRecordingRequest](docs/StartConferenceRecordingRequest.md)
 - [StartForkingRequest](docs/StartForkingRequest.md)
 - [StartRecordingRequest](docs/StartRecordingRequest.md)
 - [StartSiprecRequest](docs/StartSiprecRequest.md)
 - [StartStreamingRequest](docs/StartStreamingRequest.md)
 - [StockExchange](docs/StockExchange.md)
 - [StopForkingRequest](docs/StopForkingRequest.md)
 - [StopGatherRequest](docs/StopGatherRequest.md)
 - [StopRecordingRequest](docs/StopRecordingRequest.md)
 - [StopSiprecRequest](docs/StopSiprecRequest.md)
 - [StopStreamingRequest](docs/StopStreamingRequest.md)
 - [StreamBidirectionalCodec](docs/StreamBidirectionalCodec.md)
 - [StreamBidirectionalMode](docs/StreamBidirectionalMode.md)
 - [StreamBidirectionalSamplingRate](docs/StreamBidirectionalSamplingRate.md)
 - [StreamBidirectionalTargetLegs](docs/StreamBidirectionalTargetLegs.md)
 - [StreamStatus](docs/StreamStatus.md)
 - [StreamTrack](docs/StreamTrack.md)
 - [SubNumberOrder](docs/SubNumberOrder.md)
 - [SubNumberOrderPhoneNumber](docs/SubNumberOrderPhoneNumber.md)
 - [SubNumberOrderPhoneNumberRegulatoryRequirementsInner](docs/SubNumberOrderPhoneNumberRegulatoryRequirementsInner.md)
 - [SubNumberOrderRegulatoryRequirement](docs/SubNumberOrderRegulatoryRequirement.md)
 - [SubNumberOrderRegulatoryRequirementWithValue](docs/SubNumberOrderRegulatoryRequirementWithValue.md)
 - [SubNumberOrderRequirementGroupResponse](docs/SubNumberOrderRequirementGroupResponse.md)
 - [SubNumberOrderRequirementGroupResponseData](docs/SubNumberOrderRequirementGroupResponseData.md)
 - [SubNumberOrderResponse](docs/SubNumberOrderResponse.md)
 - [SuccessfulResponseUponAcceptingCancelFaxCommand](docs/SuccessfulResponseUponAcceptingCancelFaxCommand.md)
 - [SummaryRequest](docs/SummaryRequest.md)
 - [SummaryResponse](docs/SummaryResponse.md)
 - [SummaryResponseData](docs/SummaryResponseData.md)
 - [SupportedEmbeddingLoaders](docs/SupportedEmbeddingLoaders.md)
 - [SupportedEmbeddingModels](docs/SupportedEmbeddingModels.md)
 - [TFPhoneNumber](docs/TFPhoneNumber.md)
 - [TFVerificationRequest](docs/TFVerificationRequest.md)
 - [TFVerificationStatus](docs/TFVerificationStatus.md)
 - [TaskStatus](docs/TaskStatus.md)
 - [TaskStatusResponse](docs/TaskStatusResponse.md)
 - [TaskStatusResponseData](docs/TaskStatusResponseData.md)
 - [TeXMLRESTCommandResponse](docs/TeXMLRESTCommandResponse.md)
 - [TelephonyCredential](docs/TelephonyCredential.md)
 - [TelephonyCredentialCreateRequest](docs/TelephonyCredentialCreateRequest.md)
 - [TelephonyCredentialResponse](docs/TelephonyCredentialResponse.md)
 - [TelephonyCredentialUpdateRequest](docs/TelephonyCredentialUpdateRequest.md)
 - [TelephonySettings](docs/TelephonySettings.md)
 - [TelnyxBrand](docs/TelnyxBrand.md)
 - [TelnyxBrandWithCampaignsCount](docs/TelnyxBrandWithCampaignsCount.md)
 - [TelnyxCampaignCSP](docs/TelnyxCampaignCSP.md)
 - [TelnyxCampaignWithAssignedCountCSP](docs/TelnyxCampaignWithAssignedCountCSP.md)
 - [TelnyxDownstreamCampaign](docs/TelnyxDownstreamCampaign.md)
 - [TelnyxDownstreamCampaignRecordSet](docs/TelnyxDownstreamCampaignRecordSet.md)
 - [TelnyxTranscriptionLanguage](docs/TelnyxTranscriptionLanguage.md)
 - [TelnyxVoiceSettings](docs/TelnyxVoiceSettings.md)
 - [TexmlApplication](docs/TexmlApplication.md)
 - [TexmlApplicationResponse](docs/TexmlApplicationResponse.md)
 - [TexmlBidirectionalStreamCodec](docs/TexmlBidirectionalStreamCodec.md)
 - [TexmlBidirectionalStreamMode](docs/TexmlBidirectionalStreamMode.md)
 - [TexmlCreateCallRecordingResponseBody](docs/TexmlCreateCallRecordingResponseBody.md)
 - [TexmlCreateCallStreamingResponseBody](docs/TexmlCreateCallStreamingResponseBody.md)
 - [TexmlCreateSiprecSessionResponseBody](docs/TexmlCreateSiprecSessionResponseBody.md)
 - [TexmlGetCallRecordingResponseBody](docs/TexmlGetCallRecordingResponseBody.md)
 - [TexmlGetCallRecordingsResponseBody](docs/TexmlGetCallRecordingsResponseBody.md)
 - [TexmlRecordingChannels](docs/TexmlRecordingChannels.md)
 - [TexmlRecordingStatus](docs/TexmlRecordingStatus.md)
 - [TexmlRecordingSubresourcesUris](docs/TexmlRecordingSubresourcesUris.md)
 - [TexmlRecordingTranscription](docs/TexmlRecordingTranscription.md)
 - [TexmlStatusCallbackMethod](docs/TexmlStatusCallbackMethod.md)
 - [TexmlUpdateCallStreamingResponseBody](docs/TexmlUpdateCallStreamingResponseBody.md)
 - [TexmlUpdateSiprecSessionResponseBody](docs/TexmlUpdateSiprecSessionResponseBody.md)
 - [TextAndImageArrayInner](docs/TextAndImageArrayInner.md)
 - [TextClusteringResponse](docs/TextClusteringResponse.md)
 - [TextClusteringResponseData](docs/TextClusteringResponseData.md)
 - [TnReleaseEntry](docs/TnReleaseEntry.md)
 - [TnUploadEntry](docs/TnUploadEntry.md)
 - [TrafficType](docs/TrafficType.md)
 - [Transcription](docs/Transcription.md)
 - [TranscriptionConfig](docs/TranscriptionConfig.md)
 - [TranscriptionEngineAConfig](docs/TranscriptionEngineAConfig.md)
 - [TranscriptionEngineAConfigSpeechContextInner](docs/TranscriptionEngineAConfigSpeechContextInner.md)
 - [TranscriptionEngineBConfig](docs/TranscriptionEngineBConfig.md)
 - [TranscriptionEvent](docs/TranscriptionEvent.md)
 - [TranscriptionPayload](docs/TranscriptionPayload.md)
 - [TranscriptionPayloadTranscriptionData](docs/TranscriptionPayloadTranscriptionData.md)
 - [TranscriptionSettings](docs/TranscriptionSettings.md)
 - [TranscriptionStartRequest](docs/TranscriptionStartRequest.md)
 - [TranscriptionStartRequestTranscriptionEngineConfig](docs/TranscriptionStartRequestTranscriptionEngineConfig.md)
 - [TranscriptionStopRequest](docs/TranscriptionStopRequest.md)
 - [TransferCallRequest](docs/TransferCallRequest.md)
 - [TransferTool](docs/TransferTool.md)
 - [TransferToolParams](docs/TransferToolParams.md)
 - [TransferToolParamsCustomHeadersInner](docs/TransferToolParamsCustomHeadersInner.md)
 - [TransferToolParamsTargetsInner](docs/TransferToolParamsTargetsInner.md)
 - [TwimlRecordingChannels](docs/TwimlRecordingChannels.md)
 - [UnauthorizedError](docs/UnauthorizedError.md)
 - [UnauthorizedErrorAllOfMeta](docs/UnauthorizedErrorAllOfMeta.md)
 - [UnexpectedError](docs/UnexpectedError.md)
 - [UnexpectedErrorAllOfMeta](docs/UnexpectedErrorAllOfMeta.md)
 - [UnprocessableEntityError](docs/UnprocessableEntityError.md)
 - [UnusedUserBundlesResponse](docs/UnusedUserBundlesResponse.md)
 - [UnusedUserBundlesSchema](docs/UnusedUserBundlesSchema.md)
 - [UpdateAssistantRequest](docs/UpdateAssistantRequest.md)
 - [UpdateAuthenticationProviderRequest](docs/UpdateAuthenticationProviderRequest.md)
 - [UpdateBillingGroup](docs/UpdateBillingGroup.md)
 - [UpdateBrand](docs/UpdateBrand.md)
 - [UpdateCallControlApplicationRequest](docs/UpdateCallControlApplicationRequest.md)
 - [UpdateCallRequest](docs/UpdateCallRequest.md)
 - [UpdateCampaignRequest](docs/UpdateCampaignRequest.md)
 - [UpdateCommandResult](docs/UpdateCommandResult.md)
 - [UpdateConferenceRequest](docs/UpdateConferenceRequest.md)
 - [UpdateConversationRequest](docs/UpdateConversationRequest.md)
 - [UpdateCredentialConnectionRequest](docs/UpdateCredentialConnectionRequest.md)
 - [UpdateExternalConnectionPhoneNumberRequest](docs/UpdateExternalConnectionPhoneNumberRequest.md)
 - [UpdateExternalConnectionRequest](docs/UpdateExternalConnectionRequest.md)
 - [UpdateFaxApplicationRequest](docs/UpdateFaxApplicationRequest.md)
 - [UpdateFqdnConnectionRequest](docs/UpdateFqdnConnectionRequest.md)
 - [UpdateFqdnRequest](docs/UpdateFqdnRequest.md)
 - [UpdateIpConnectionRequest](docs/UpdateIpConnectionRequest.md)
 - [UpdateIpRequest](docs/UpdateIpRequest.md)
 - [UpdateLocationRequest](docs/UpdateLocationRequest.md)
 - [UpdateManagedAccountGlobalChannelLimit200Response](docs/UpdateManagedAccountGlobalChannelLimit200Response.md)
 - [UpdateManagedAccountGlobalChannelLimitRequest](docs/UpdateManagedAccountGlobalChannelLimitRequest.md)
 - [UpdateManagedAccountRequest](docs/UpdateManagedAccountRequest.md)
 - [UpdateMediaRequest](docs/UpdateMediaRequest.md)
 - [UpdateMessagingProfileRequest](docs/UpdateMessagingProfileRequest.md)
 - [UpdateNumberOrderPhoneNumberRequest](docs/UpdateNumberOrderPhoneNumberRequest.md)
 - [UpdateNumberOrderPhoneNumberRequirementGroup200Response](docs/UpdateNumberOrderPhoneNumberRequirementGroup200Response.md)
 - [UpdateNumberOrderPhoneNumberRequirementGroupRequest](docs/UpdateNumberOrderPhoneNumberRequirementGroupRequest.md)
 - [UpdateNumberOrderRequest](docs/UpdateNumberOrderRequest.md)
 - [UpdateOutboundChannels200Response](docs/UpdateOutboundChannels200Response.md)
 - [UpdateOutboundChannels200ResponseData](docs/UpdateOutboundChannels200ResponseData.md)
 - [UpdateOutboundChannelsDefaultResponse](docs/UpdateOutboundChannelsDefaultResponse.md)
 - [UpdateOutboundChannelsDefaultResponseErrorsInner](docs/UpdateOutboundChannelsDefaultResponseErrorsInner.md)
 - [UpdateOutboundChannelsDefaultResponseErrorsInnerSource](docs/UpdateOutboundChannelsDefaultResponseErrorsInnerSource.md)
 - [UpdateOutboundChannelsRequest](docs/UpdateOutboundChannelsRequest.md)
 - [UpdateOutboundVoiceProfileRequest](docs/UpdateOutboundVoiceProfileRequest.md)
 - [UpdatePartnerCampaignRequest](docs/UpdatePartnerCampaignRequest.md)
 - [UpdatePhoneNumberMessagingSettingsRequest](docs/UpdatePhoneNumberMessagingSettingsRequest.md)
 - [UpdatePhoneNumberRequest](docs/UpdatePhoneNumberRequest.md)
 - [UpdatePhoneNumberVoiceSettingsRequest](docs/UpdatePhoneNumberVoiceSettingsRequest.md)
 - [UpdatePortingOrder](docs/UpdatePortingOrder.md)
 - [UpdatePortingOrder200Response](docs/UpdatePortingOrder200Response.md)
 - [UpdatePortingOrder200ResponseMeta](docs/UpdatePortingOrder200ResponseMeta.md)
 - [UpdatePortingOrderActivationSettings](docs/UpdatePortingOrderActivationSettings.md)
 - [UpdatePortingOrderMessaging](docs/UpdatePortingOrderMessaging.md)
 - [UpdatePortingOrderRequirement](docs/UpdatePortingOrderRequirement.md)
 - [UpdatePortingOrdersActivationJobRequest](docs/UpdatePortingOrdersActivationJobRequest.md)
 - [UpdatePortoutStatusRequest](docs/UpdatePortoutStatusRequest.md)
 - [UpdateRegulatoryRequirement](docs/UpdateRegulatoryRequirement.md)
 - [UpdateRequirementGroupRequest](docs/UpdateRequirementGroupRequest.md)
 - [UpdateRequirementGroupRequestRegulatoryRequirementsInner](docs/UpdateRequirementGroupRequestRegulatoryRequirementsInner.md)
 - [UpdateShortCodeRequest](docs/UpdateShortCodeRequest.md)
 - [UpdateSubNumberOrderRequest](docs/UpdateSubNumberOrderRequest.md)
 - [UpdateTexmlApplicationRequest](docs/UpdateTexmlApplicationRequest.md)
 - [UpdateVerifyProfileCallRequest](docs/UpdateVerifyProfileCallRequest.md)
 - [UpdateVerifyProfileFlashcallRequest](docs/UpdateVerifyProfileFlashcallRequest.md)
 - [UpdateVerifyProfileRequest](docs/UpdateVerifyProfileRequest.md)
 - [UpdateVerifyProfileSMSRequest](docs/UpdateVerifyProfileSMSRequest.md)
 - [UplinkData](docs/UplinkData.md)
 - [Upload](docs/Upload.md)
 - [UploadMediaRequest](docs/UploadMediaRequest.md)
 - [Url](docs/Url.md)
 - [UrlShortenerSettings](docs/UrlShortenerSettings.md)
 - [UsagePaymentMethod](docs/UsagePaymentMethod.md)
 - [UsageReportsOptionsRecord](docs/UsageReportsOptionsRecord.md)
 - [UsageReportsOptionsResponse](docs/UsageReportsOptionsResponse.md)
 - [UsageReportsResponse](docs/UsageReportsResponse.md)
 - [UseCaseCategories](docs/UseCaseCategories.md)
 - [UsecaseMetadata](docs/UsecaseMetadata.md)
 - [UserAddress](docs/UserAddress.md)
 - [UserAddressCreate](docs/UserAddressCreate.md)
 - [UserBalance](docs/UserBalance.md)
 - [UserBundle](docs/UserBundle.md)
 - [UserBundleCreateRequest](docs/UserBundleCreateRequest.md)
 - [UserBundleCreateResponse](docs/UserBundleCreateResponse.md)
 - [UserBundleResourceSchema](docs/UserBundleResourceSchema.md)
 - [UserBundleResourcesResponse](docs/UserBundleResourcesResponse.md)
 - [UserBundleResponse](docs/UserBundleResponse.md)
 - [UserBundleSummary](docs/UserBundleSummary.md)
 - [UserEmbeddedBuckets](docs/UserEmbeddedBuckets.md)
 - [UserEmbeddedBucketsData](docs/UserEmbeddedBucketsData.md)
 - [UserRequirement](docs/UserRequirement.md)
 - [ValidateAddressActionResponse](docs/ValidateAddressActionResponse.md)
 - [ValidateAddressField](docs/ValidateAddressField.md)
 - [ValidateAddressRequest](docs/ValidateAddressRequest.md)
 - [ValidateAddressResult](docs/ValidateAddressResult.md)
 - [ValidateRegistrationCodesRequest](docs/ValidateRegistrationCodesRequest.md)
 - [ValidationCodes](docs/ValidationCodes.md)
 - [ValidationCodesPhoneNumbersInner](docs/ValidationCodesPhoneNumbersInner.md)
 - [ValidationCodesRequest](docs/ValidationCodesRequest.md)
 - [ValidationCodesRequestVerificationCodesInner](docs/ValidationCodesRequestVerificationCodesInner.md)
 - [ValidationError](docs/ValidationError.md)
 - [Verification](docs/Verification.md)
 - [VerificationCodesRequest](docs/VerificationCodesRequest.md)
 - [VerificationProfileRecordType](docs/VerificationProfileRecordType.md)
 - [VerificationRecordType](docs/VerificationRecordType.md)
 - [VerificationRequestEgress](docs/VerificationRequestEgress.md)
 - [VerificationRequestStatus](docs/VerificationRequestStatus.md)
 - [VerificationStatus](docs/VerificationStatus.md)
 - [VerificationType](docs/VerificationType.md)
 - [VerifiedNumberRecordType](docs/VerifiedNumberRecordType.md)
 - [VerifiedNumberResponse](docs/VerifiedNumberResponse.md)
 - [VerifiedNumberResponseDataWrapper](docs/VerifiedNumberResponseDataWrapper.md)
 - [VerifyDetailRecord](docs/VerifyDetailRecord.md)
 - [VerifyPhoneNumberCoverage201Response](docs/VerifyPhoneNumberCoverage201Response.md)
 - [VerifyPhoneNumberCoverageRequest](docs/VerifyPhoneNumberCoverageRequest.md)
 - [VerifyPortingVerificationCodes200Response](docs/VerifyPortingVerificationCodes200Response.md)
 - [VerifyPortingVerificationCodesRequest](docs/VerifyPortingVerificationCodesRequest.md)
 - [VerifyPortingVerificationCodesRequestVerificationCodesInner](docs/VerifyPortingVerificationCodesRequestVerificationCodesInner.md)
 - [VerifyProfileCallResponse](docs/VerifyProfileCallResponse.md)
 - [VerifyProfileFlashcallResponse](docs/VerifyProfileFlashcallResponse.md)
 - [VerifyProfileMessageTemplateResponse](docs/VerifyProfileMessageTemplateResponse.md)
 - [VerifyProfileResponse](docs/VerifyProfileResponse.md)
 - [VerifyProfileResponseDataWrapper](docs/VerifyProfileResponseDataWrapper.md)
 - [VerifyProfileSMSResponse](docs/VerifyProfileSMSResponse.md)
 - [VerifyVerificationCodeRequest](docs/VerifyVerificationCodeRequest.md)
 - [VerifyVerificationCodeRequestById](docs/VerifyVerificationCodeRequestById.md)
 - [VerifyVerificationCodeRequestByPhoneNumber](docs/VerifyVerificationCodeRequestByPhoneNumber.md)
 - [VerifyVerificationCodeResponse](docs/VerifyVerificationCodeResponse.md)
 - [VerifyVerificationCodeResponseData](docs/VerifyVerificationCodeResponseData.md)
 - [Vertical](docs/Vertical.md)
 - [VideoRegion](docs/VideoRegion.md)
 - [ViewRoomParticipant200Response](docs/ViewRoomParticipant200Response.md)
 - [ViewRoomRecording200Response](docs/ViewRoomRecording200Response.md)
 - [ViewRoomSession200Response](docs/ViewRoomSession200Response.md)
 - [VirtualCrossConnect](docs/VirtualCrossConnect.md)
 - [VirtualCrossConnectCombined](docs/VirtualCrossConnectCombined.md)
 - [VirtualCrossConnectCoverage](docs/VirtualCrossConnectCoverage.md)
 - [VirtualCrossConnectCreate](docs/VirtualCrossConnectCreate.md)
 - [VirtualCrossConnectPatch](docs/VirtualCrossConnectPatch.md)
 - [VoiceSettings](docs/VoiceSettings.md)
 - [VoicemailPrefResponse](docs/VoicemailPrefResponse.md)
 - [VoicemailRequest](docs/VoicemailRequest.md)
 - [Volume](docs/Volume.md)
 - [WdrReport](docs/WdrReport.md)
 - [WdrReportRequest](docs/WdrReportRequest.md)
 - [WebhookApiVersion](docs/WebhookApiVersion.md)
 - [WebhookDelivery](docs/WebhookDelivery.md)
 - [WebhookDeliveryWebhook](docs/WebhookDeliveryWebhook.md)
 - [WebhookPortingOrderDeletedPayload](docs/WebhookPortingOrderDeletedPayload.md)
 - [WebhookPortingOrderMessagingChangedPayload](docs/WebhookPortingOrderMessagingChangedPayload.md)
 - [WebhookPortingOrderMessagingChangedPayloadMessaging](docs/WebhookPortingOrderMessagingChangedPayloadMessaging.md)
 - [WebhookPortingOrderNewCommentPayload](docs/WebhookPortingOrderNewCommentPayload.md)
 - [WebhookPortingOrderNewCommentPayloadComment](docs/WebhookPortingOrderNewCommentPayloadComment.md)
 - [WebhookPortingOrderSplitPayload](docs/WebhookPortingOrderSplitPayload.md)
 - [WebhookPortingOrderSplitPayloadFrom](docs/WebhookPortingOrderSplitPayloadFrom.md)
 - [WebhookPortingOrderSplitPayloadPortingPhoneNumbersInner](docs/WebhookPortingOrderSplitPayloadPortingPhoneNumbersInner.md)
 - [WebhookPortingOrderSplitPayloadTo](docs/WebhookPortingOrderSplitPayloadTo.md)
 - [WebhookPortingOrderStatusChangedPayload](docs/WebhookPortingOrderStatusChangedPayload.md)
 - [WebhookPortoutFocDateChangedPayload](docs/WebhookPortoutFocDateChangedPayload.md)
 - [WebhookPortoutNewCommentPayload](docs/WebhookPortoutNewCommentPayload.md)
 - [WebhookPortoutStatusChangedPayload](docs/WebhookPortoutStatusChangedPayload.md)
 - [WebhookTool](docs/WebhookTool.md)
 - [WebhookToolParams](docs/WebhookToolParams.md)
 - [WebhookToolParamsBodyParameters](docs/WebhookToolParamsBodyParameters.md)
 - [WebhookToolParamsHeadersInner](docs/WebhookToolParamsHeadersInner.md)
 - [WebhookToolParamsPathParameters](docs/WebhookToolParamsPathParameters.md)
 - [WebhookToolParamsQueryParameters](docs/WebhookToolParamsQueryParameters.md)
 - [WireguardInterface](docs/WireguardInterface.md)
 - [WireguardInterfaceCreate](docs/WireguardInterfaceCreate.md)
 - [WireguardInterfaceRead](docs/WireguardInterfaceRead.md)
 - [WireguardPeer](docs/WireguardPeer.md)
 - [WireguardPeerCreate](docs/WireguardPeerCreate.md)
 - [WireguardPeerPatch](docs/WireguardPeerPatch.md)
 - [WirelessConnectivityLog](docs/WirelessConnectivityLog.md)
 - [WirelessCost](docs/WirelessCost.md)
 - [WirelessRate](docs/WirelessRate.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), telnyx.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@telnyx.com


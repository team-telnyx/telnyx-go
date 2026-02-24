// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// SimCardService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimCardService] method instead.
type SimCardService struct {
	Options []option.RequestOption
	Actions SimCardActionService
}

// NewSimCardService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSimCardService(opts ...option.RequestOption) (r SimCardService) {
	r = SimCardService{}
	r.Options = opts
	r.Actions = NewSimCardActionService(opts...)
	return
}

// Returns the details regarding a specific SIM card.
func (r *SimCardService) Get(ctx context.Context, id string, query SimCardGetParams, opts ...option.RequestOption) (res *SimCardGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_cards/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates SIM card data
func (r *SimCardService) Update(ctx context.Context, simCardID string, body SimCardUpdateParams, opts ...option.RequestOption) (res *SimCardUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if simCardID == "" {
		err = errors.New("missing required sim_card_id parameter")
		return
	}
	path := fmt.Sprintf("sim_cards/%s", simCardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get all SIM cards belonging to the user that match the given filters.
func (r *SimCardService) List(ctx context.Context, query SimCardListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[shared.SimpleSimCard], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "sim_cards"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get all SIM cards belonging to the user that match the given filters.
func (r *SimCardService) ListAutoPaging(ctx context.Context, query SimCardListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[shared.SimpleSimCard] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// The SIM card will be decommissioned, removed from your account and you will stop
// being charged.<br />The SIM card won't be able to connect to the network after
// the deletion is completed, thus making it impossible to consume data.<br/>
// Transitioning to the disabled state may take a period of time. Until the
// transition is completed, the SIM card status will be disabling
// <code>disabling</code>.<br />In order to re-enable the SIM card, you will need
// to re-register it.
func (r *SimCardService) Delete(ctx context.Context, id string, body SimCardDeleteParams, opts ...option.RequestOption) (res *SimCardDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_cards/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// It returns the activation code for an eSIM.<br/><br/> This API is only available
// for eSIMs. If the given SIM is a physical SIM card, or has already been
// installed, an error will be returned.
func (r *SimCardService) GetActivationCode(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardGetActivationCodeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_cards/%s/activation_code", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// It returns the device details where a SIM card is currently being used.
func (r *SimCardService) GetDeviceDetails(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardGetDeviceDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_cards/%s/device_details", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// It returns the public IP requested for a SIM card.
func (r *SimCardService) GetPublicIP(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardGetPublicIPResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_cards/%s/public_ip", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This API allows listing a paginated collection of Wireless Connectivity Logs
// associated with a SIM Card, for troubleshooting purposes.
func (r *SimCardService) ListWirelessConnectivityLogs(ctx context.Context, id string, query SimCardListWirelessConnectivityLogsParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[SimCardListWirelessConnectivityLogsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_cards/%s/wireless_connectivity_logs", id)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// This API allows listing a paginated collection of Wireless Connectivity Logs
// associated with a SIM Card, for troubleshooting purposes.
func (r *SimCardService) ListWirelessConnectivityLogsAutoPaging(ctx context.Context, id string, query SimCardListWirelessConnectivityLogsParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[SimCardListWirelessConnectivityLogsResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.ListWirelessConnectivityLogs(ctx, id, query, opts...))
}

type SimCard struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// Indicate whether the SIM card has any pending (in-progress) actions.
	ActionsInProgress bool `json:"actions_in_progress"`
	// List of IMEIs authorized to use a given SIM card.
	AuthorizedImeis []string `json:"authorized_imeis" api:"nullable"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The SIM card consumption so far in the current billing cycle.
	CurrentBillingPeriodConsumedData SimCardCurrentBillingPeriodConsumedData `json:"current_billing_period_consumed_data"`
	// Current physical location data of a given SIM card. Accuracy is given in meters.
	CurrentDeviceLocation SimCardCurrentDeviceLocation `json:"current_device_location"`
	// IMEI of the device where a given SIM card is currently being used.
	CurrentImei string `json:"current_imei"`
	// Mobile Country Code of the current network to which the SIM card is connected.
	// It's a three decimal digit that identifies a country.<br/><br/> This code is
	// commonly seen joined with a Mobile Network Code (MNC) in a tuple that allows
	// identifying a carrier known as PLMN (Public Land Mobile Network) code.
	CurrentMcc string `json:"current_mcc"`
	// Mobile Network Code of the current network to which the SIM card is connected.
	// It's a two to three decimal digits that identify a network.<br/><br/> This code
	// is commonly seen joined with a Mobile Country Code (MCC) in a tuple that allows
	// identifying a carrier known as PLMN (Public Land Mobile Network) code.
	CurrentMnc string `json:"current_mnc"`
	// The SIM card individual data limit configuration.
	DataLimit SimCardDataLimit `json:"data_limit"`
	// The Embedded Identity Document (eID) for eSIM cards.
	Eid string `json:"eid" api:"nullable"`
	// The installation status of the eSIM. Only applicable for eSIM cards.
	//
	// Any of "released", "disabled".
	EsimInstallationStatus SimCardEsimInstallationStatus `json:"esim_installation_status" api:"nullable"`
	// The ICCID is the identifier of the specific SIM card/chip. Each SIM is
	// internationally identified by its integrated circuit card identifier (ICCID).
	// ICCIDs are stored in the SIM card's memory and are also engraved or printed on
	// the SIM card body during a process called personalization.
	Iccid string `json:"iccid"`
	// SIM cards are identified on their individual network operators by a unique
	// International Mobile Subscriber Identity (IMSI). <br/> Mobile network operators
	// connect mobile phone calls and communicate with their market SIM cards using
	// their IMSIs. The IMSI is stored in the Subscriber Identity Module (SIM) inside
	// the device and is sent by the device to the appropriate network. It is used to
	// acquire the details of the device in the Home Location Register (HLR) or the
	// Visitor Location Register (VLR).
	Imsi string `json:"imsi"`
	// The SIM's address in the currently connected network. This IPv4 address is
	// usually obtained dynamically, so it may vary according to the location or new
	// connections.
	Ipv4 string `json:"ipv4"`
	// The SIM's address in the currently connected network. This IPv6 address is
	// usually obtained dynamically, so it may vary according to the location or new
	// connections.
	Ipv6 string `json:"ipv6"`
	// Indicates whether the device is actively connected to a network and able to run
	// data.
	//
	// Any of "connected", "disconnected", "unknown".
	LiveDataSession SimCardLiveDataSession `json:"live_data_session"`
	// Mobile Station International Subscriber Directory Number (MSISDN) is a number
	// used to identify a mobile phone number internationally. <br/> MSISDN is defined
	// by the E.164 numbering plan. It includes a country code and a National
	// Destination Code which identifies the subscriber's operator.
	Msisdn string `json:"msisdn"`
	// PIN and PUK codes for the SIM card. Only available when
	// include_pin_puk_codes=true is set in the request.
	PinPukCodes SimCardPinPukCodes `json:"pin_puk_codes"`
	RecordType  string             `json:"record_type"`
	// List of resources with actions in progress.
	ResourcesWithInProgressActions []map[string]any `json:"resources_with_in_progress_actions"`
	// The group SIMCardGroup identification. This attribute can be <code>null</code>
	// when it's present in an associated resource.
	SimCardGroupID string               `json:"sim_card_group_id" format:"uuid"`
	Status         shared.SimCardStatus `json:"status"`
	// Searchable tags associated with the SIM card
	Tags []string `json:"tags"`
	// The type of SIM card
	//
	// Any of "physical", "esim".
	Type SimCardType `json:"type"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// The version of the SIM card.
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                               respjson.Field
		ActionsInProgress                respjson.Field
		AuthorizedImeis                  respjson.Field
		CreatedAt                        respjson.Field
		CurrentBillingPeriodConsumedData respjson.Field
		CurrentDeviceLocation            respjson.Field
		CurrentImei                      respjson.Field
		CurrentMcc                       respjson.Field
		CurrentMnc                       respjson.Field
		DataLimit                        respjson.Field
		Eid                              respjson.Field
		EsimInstallationStatus           respjson.Field
		Iccid                            respjson.Field
		Imsi                             respjson.Field
		Ipv4                             respjson.Field
		Ipv6                             respjson.Field
		LiveDataSession                  respjson.Field
		Msisdn                           respjson.Field
		PinPukCodes                      respjson.Field
		RecordType                       respjson.Field
		ResourcesWithInProgressActions   respjson.Field
		SimCardGroupID                   respjson.Field
		Status                           respjson.Field
		Tags                             respjson.Field
		Type                             respjson.Field
		UpdatedAt                        respjson.Field
		Version                          respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCard) RawJSON() string { return r.JSON.raw }
func (r *SimCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SimCard to a SimCardParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SimCardParam.Overrides()
func (r SimCard) ToParam() SimCardParam {
	return param.Override[SimCardParam](json.RawMessage(r.RawJSON()))
}

// The SIM card consumption so far in the current billing cycle.
type SimCardCurrentBillingPeriodConsumedData struct {
	Amount string `json:"amount"`
	Unit   string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardCurrentBillingPeriodConsumedData) RawJSON() string { return r.JSON.raw }
func (r *SimCardCurrentBillingPeriodConsumedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current physical location data of a given SIM card. Accuracy is given in meters.
type SimCardCurrentDeviceLocation struct {
	Accuracy     int64  `json:"accuracy"`
	AccuracyUnit string `json:"accuracy_unit"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accuracy     respjson.Field
		AccuracyUnit respjson.Field
		Latitude     respjson.Field
		Longitude    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardCurrentDeviceLocation) RawJSON() string { return r.JSON.raw }
func (r *SimCardCurrentDeviceLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The SIM card individual data limit configuration.
type SimCardDataLimit struct {
	Amount string `json:"amount"`
	// Any of "MB", "GB".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardDataLimit) RawJSON() string { return r.JSON.raw }
func (r *SimCardDataLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The installation status of the eSIM. Only applicable for eSIM cards.
type SimCardEsimInstallationStatus string

const (
	SimCardEsimInstallationStatusReleased SimCardEsimInstallationStatus = "released"
	SimCardEsimInstallationStatusDisabled SimCardEsimInstallationStatus = "disabled"
)

// Indicates whether the device is actively connected to a network and able to run
// data.
type SimCardLiveDataSession string

const (
	SimCardLiveDataSessionConnected    SimCardLiveDataSession = "connected"
	SimCardLiveDataSessionDisconnected SimCardLiveDataSession = "disconnected"
	SimCardLiveDataSessionUnknown      SimCardLiveDataSession = "unknown"
)

// PIN and PUK codes for the SIM card. Only available when
// include_pin_puk_codes=true is set in the request.
type SimCardPinPukCodes struct {
	// The primary Personal Identification Number (PIN) for the SIM card. This is a
	// 4-digit code used to protect the SIM card from unauthorized use.
	Pin1 string `json:"pin1"`
	// The secondary Personal Identification Number (PIN2) for the SIM card. This is a
	// 4-digit code used for additional security features.
	Pin2 string `json:"pin2"`
	// The primary Personal Unblocking Key (PUK1) for the SIM card. This is an 8-digit
	// code used to unlock the SIM card if PIN1 is entered incorrectly multiple times.
	Puk1 string `json:"puk1"`
	// The secondary Personal Unblocking Key (PUK2) for the SIM card. This is an
	// 8-digit code used to unlock the SIM card if PIN2 is entered incorrectly multiple
	// times.
	Puk2 string `json:"puk2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pin1        respjson.Field
		Pin2        respjson.Field
		Puk1        respjson.Field
		Puk2        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardPinPukCodes) RawJSON() string { return r.JSON.raw }
func (r *SimCardPinPukCodes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of SIM card
type SimCardType string

const (
	SimCardTypePhysical SimCardType = "physical"
	SimCardTypeEsim     SimCardType = "esim"
)

type SimCardParam struct {
	// The group SIMCardGroup identification. This attribute can be <code>null</code>
	// when it's present in an associated resource.
	SimCardGroupID param.Opt[string] `json:"sim_card_group_id,omitzero" format:"uuid"`
	// List of IMEIs authorized to use a given SIM card.
	AuthorizedImeis []string `json:"authorized_imeis,omitzero"`
	// The SIM card individual data limit configuration.
	DataLimit SimCardDataLimitParam     `json:"data_limit,omitzero"`
	Status    shared.SimCardStatusParam `json:"status,omitzero"`
	// Searchable tags associated with the SIM card
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r SimCardParam) MarshalJSON() (data []byte, err error) {
	type shadow SimCardParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The SIM card consumption so far in the current billing cycle.
type SimCardCurrentBillingPeriodConsumedDataParam struct {
	Amount param.Opt[string] `json:"amount,omitzero"`
	Unit   param.Opt[string] `json:"unit,omitzero"`
	paramObj
}

func (r SimCardCurrentBillingPeriodConsumedDataParam) MarshalJSON() (data []byte, err error) {
	type shadow SimCardCurrentBillingPeriodConsumedDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardCurrentBillingPeriodConsumedDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current physical location data of a given SIM card. Accuracy is given in meters.
type SimCardCurrentDeviceLocationParam struct {
	Accuracy     param.Opt[int64]  `json:"accuracy,omitzero"`
	AccuracyUnit param.Opt[string] `json:"accuracy_unit,omitzero"`
	Latitude     param.Opt[string] `json:"latitude,omitzero"`
	Longitude    param.Opt[string] `json:"longitude,omitzero"`
	paramObj
}

func (r SimCardCurrentDeviceLocationParam) MarshalJSON() (data []byte, err error) {
	type shadow SimCardCurrentDeviceLocationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardCurrentDeviceLocationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The SIM card individual data limit configuration.
type SimCardDataLimitParam struct {
	Amount param.Opt[string] `json:"amount,omitzero"`
	// Any of "MB", "GB".
	Unit string `json:"unit,omitzero"`
	paramObj
}

func (r SimCardDataLimitParam) MarshalJSON() (data []byte, err error) {
	type shadow SimCardDataLimitParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardDataLimitParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SimCardDataLimitParam](
		"unit", "MB", "GB",
	)
}

// PIN and PUK codes for the SIM card. Only available when
// include_pin_puk_codes=true is set in the request.
type SimCardPinPukCodesParam struct {
	paramObj
}

func (r SimCardPinPukCodesParam) MarshalJSON() (data []byte, err error) {
	type shadow SimCardPinPukCodesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardPinPukCodesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGetResponse struct {
	Data SimCard `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardUpdateResponse struct {
	Data SimCard `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardDeleteResponse struct {
	Data SimCard `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGetActivationCodeResponse struct {
	Data SimCardGetActivationCodeResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGetActivationCodeResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGetActivationCodeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGetActivationCodeResponseData struct {
	// Contents of the eSIM activation QR code.
	ActivationCode string `json:"activation_code"`
	RecordType     string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActivationCode respjson.Field
		RecordType     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGetActivationCodeResponseData) RawJSON() string { return r.JSON.raw }
func (r *SimCardGetActivationCodeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGetDeviceDetailsResponse struct {
	Data SimCardGetDeviceDetailsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGetDeviceDetailsResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGetDeviceDetailsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGetDeviceDetailsResponseData struct {
	// Brand of the device where the SIM card is being used in.
	BrandName string `json:"brand_name"`
	// Type of the device where the SIM card is being used in.
	DeviceType string `json:"device_type"`
	// IMEI of the device where the SIM card is being used in.
	Imei string `json:"imei"`
	// Brand of the device where the SIM card is being used in.
	ModelName string `json:"model_name"`
	// Operating system of the device where the SIM card is being used in.
	OperatingSystem string `json:"operating_system"`
	RecordType      string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandName       respjson.Field
		DeviceType      respjson.Field
		Imei            respjson.Field
		ModelName       respjson.Field
		OperatingSystem respjson.Field
		RecordType      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGetDeviceDetailsResponseData) RawJSON() string { return r.JSON.raw }
func (r *SimCardGetDeviceDetailsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGetPublicIPResponse struct {
	Data SimCardGetPublicIPResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGetPublicIPResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGetPublicIPResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGetPublicIPResponseData struct {
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The provisioned IP address. This attribute will only be available when
	// underlying resource status is in a "provisioned" status.
	IP         string `json:"ip"`
	RecordType string `json:"record_type"`
	RegionCode string `json:"region_code"`
	SimCardID  string `json:"sim_card_id" format:"uuid"`
	// Any of "ipv4".
	Type string `json:"type"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		IP          respjson.Field
		RecordType  respjson.Field
		RegionCode  respjson.Field
		SimCardID   respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGetPublicIPResponseData) RawJSON() string { return r.JSON.raw }
func (r *SimCardGetPublicIPResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This object represents a wireless connectivity session log that happened through
// a SIM card. It aids in finding out potential problems when the SIM is not able
// to attach properly.
type SimCardListWirelessConnectivityLogsResponse struct {
	// Uniquely identifies the session.
	ID int64 `json:"id"`
	// The Access Point Name (APN) identifies the packet data network that a mobile
	// data user wants to communicate with.
	Apn string `json:"apn"`
	// The cell ID to which the SIM connected.
	CellID string `json:"cell_id"`
	// ISO 8601 formatted date-time indicating when the record was created.
	CreatedAt string `json:"created_at"`
	// The International Mobile Equipment Identity (or IMEI) is a number, usually
	// unique, that identifies the device currently being used connect to the network.
	Imei string `json:"imei"`
	// SIM cards are identified on their individual network operators by a unique
	// International Mobile Subscriber Identity (IMSI). <br/> Mobile network operators
	// connect mobile phone calls and communicate with their market SIM cards using
	// their IMSIs. The IMSI is stored in the Subscriber Identity Module (SIM) inside
	// the device and is sent by the device to the appropriate network. It is used to
	// acquire the details of the device in the Home Location Register (HLR) or the
	// Visitor Location Register (VLR).
	Imsi string `json:"imsi"`
	// The SIM's address in the currently connected network. This IPv4 address is
	// usually obtained dynamically, so it may vary according to the location or new
	// connections.
	Ipv4 string `json:"ipv4"`
	// The SIM's address in the currently connected network. This IPv6 address is
	// usually obtained dynamically, so it may vary according to the location or new
	// connections.
	Ipv6 string `json:"ipv6"`
	// ISO 8601 formatted date-time indicating when the last heartbeat to the device
	// was successfully recorded.
	LastSeen string `json:"last_seen"`
	// The type of the session, 'registration' being the initial authentication session
	// and 'data' the actual data transfer sessions.
	//
	// Any of "registration", "data".
	LogType SimCardListWirelessConnectivityLogsResponseLogType `json:"log_type"`
	// It's a three decimal digit that identifies a country.<br/><br/> This code is
	// commonly seen joined with a Mobile Network Code (MNC) in a tuple that allows
	// identifying a carrier known as PLMN (Public Land Mobile Network) code.
	MobileCountryCode string `json:"mobile_country_code"`
	// It's a two to three decimal digits that identify a network.<br/><br/> This code
	// is commonly seen joined with a Mobile Country Code (MCC) in a tuple that allows
	// identifying a carrier known as PLMN (Public Land Mobile Network) code.
	MobileNetworkCode string `json:"mobile_network_code"`
	// The radio technology the SIM card used during the session.
	RadioAccessTechnology string `json:"radio_access_technology"`
	RecordType            string `json:"record_type"`
	// The identification UUID of the related SIM card resource.
	SimCardID string `json:"sim_card_id" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the session started.
	StartTime string `json:"start_time"`
	// The state of the SIM card after when the session happened.
	State string `json:"state"`
	// ISO 8601 formatted date-time indicating when the session ended.
	StopTime string `json:"stop_time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Apn                   respjson.Field
		CellID                respjson.Field
		CreatedAt             respjson.Field
		Imei                  respjson.Field
		Imsi                  respjson.Field
		Ipv4                  respjson.Field
		Ipv6                  respjson.Field
		LastSeen              respjson.Field
		LogType               respjson.Field
		MobileCountryCode     respjson.Field
		MobileNetworkCode     respjson.Field
		RadioAccessTechnology respjson.Field
		RecordType            respjson.Field
		SimCardID             respjson.Field
		StartTime             respjson.Field
		State                 respjson.Field
		StopTime              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardListWirelessConnectivityLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardListWirelessConnectivityLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the session, 'registration' being the initial authentication session
// and 'data' the actual data transfer sessions.
type SimCardListWirelessConnectivityLogsResponseLogType string

const (
	SimCardListWirelessConnectivityLogsResponseLogTypeRegistration SimCardListWirelessConnectivityLogsResponseLogType = "registration"
	SimCardListWirelessConnectivityLogsResponseLogTypeData         SimCardListWirelessConnectivityLogsResponseLogType = "data"
)

type SimCardGetParams struct {
	// When set to true, includes the PIN and PUK codes in the response. These codes
	// are used for SIM card security and unlocking purposes. Available for both
	// physical SIM cards and eSIMs.
	IncludePinPukCodes param.Opt[bool] `query:"include_pin_puk_codes,omitzero" json:"-"`
	// It includes the associated SIM card group object in the response when present.
	IncludeSimCardGroup param.Opt[bool] `query:"include_sim_card_group,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardGetParams]'s query parameters as `url.Values`.
func (r SimCardGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SimCardUpdateParams struct {
	SimCard SimCardParam
	paramObj
}

func (r SimCardUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimCard)
}
func (r *SimCardUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SimCard)
}

type SimCardListParams struct {
	// A valid SIM card group ID.
	FilterSimCardGroupID param.Opt[string] `query:"filter[sim_card_group_id],omitzero" format:"uuid" json:"-"`
	// It includes the associated SIM card group object in the response when present.
	IncludeSimCardGroup param.Opt[bool]  `query:"include_sim_card_group,omitzero" json:"-"`
	PageNumber          param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize            param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter for SIM cards (deepObject style). Originally:
	// filter[iccid], filter[msisdn], filter[status], filter[tags]
	Filter SimCardListParamsFilter `query:"filter,omitzero" json:"-"`
	// Sorts SIM cards by the given field. Defaults to ascending order unless field is
	// prefixed with a minus sign.
	//
	// Any of "current_billing_period_consumed_data.amount",
	// "-current_billing_period_consumed_data.amount".
	Sort SimCardListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardListParams]'s query parameters as `url.Values`.
func (r SimCardListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter for SIM cards (deepObject style). Originally:
// filter[iccid], filter[msisdn], filter[status], filter[tags]
type SimCardListParamsFilter struct {
	// A search string to partially match for the SIM card's ICCID.
	Iccid param.Opt[string] `query:"iccid,omitzero" json:"-"`
	// A search string to match for the SIM card's MSISDN.
	Msisdn param.Opt[string] `query:"msisdn,omitzero" json:"-"`
	// Filter by a SIM card's status.
	//
	// Any of "enabled", "disabled", "standby", "data_limit_exceeded",
	// "unauthorized_imei".
	Status []string `query:"status,omitzero" json:"-"`
	// A list of SIM card tags to filter on.<br/><br/> If the SIM card contains
	// <b><i>all</i></b> of the given <code>tags</code> they will be found.<br/><br/>
	// For example, if the SIM cards have the following tags: <ul>
	//
	//	<li><code>['customers', 'staff', 'test']</code>
	//	<li><code>['test']</code></li>
	//	<li><code>['customers']</code></li>
	//
	// </ul>
	// Searching for <code>['customers', 'test']</code> returns only the first because it's the only one with both tags.<br/> Searching for <code>test</code> returns the first two SIMs, because both of them have such tag.<br/> Searching for <code>customers</code> returns the first and last SIMs.<br/>
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardListParamsFilter]'s query parameters as
// `url.Values`.
func (r SimCardListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorts SIM cards by the given field. Defaults to ascending order unless field is
// prefixed with a minus sign.
type SimCardListParamsSort string

const (
	SimCardListParamsSortCurrentBillingPeriodConsumedDataAmount      SimCardListParamsSort = "current_billing_period_consumed_data.amount"
	SimCardListParamsSortMinusCurrentBillingPeriodConsumedDataAmount SimCardListParamsSort = "-current_billing_period_consumed_data.amount"
)

type SimCardDeleteParams struct {
	// Enables deletion of disabled eSIMs that can't be uninstalled from a device. This
	// is irreversible and the eSIM cannot be re-registered.
	ReportLost param.Opt[bool] `query:"report_lost,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardDeleteParams]'s query parameters as `url.Values`.
func (r SimCardDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SimCardListWirelessConnectivityLogsParams struct {
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardListWirelessConnectivityLogsParams]'s query
// parameters as `url.Values`.
func (r SimCardListWirelessConnectivityLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

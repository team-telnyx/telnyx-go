# WirelessConnectivityLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **int32** | Uniquely identifies the session. | [optional] [readonly] 
**SimCardId** | Pointer to **string** | The identification UUID of the related SIM card resource. | [optional] 
**LogType** | Pointer to **string** | The type of the session, &#39;registration&#39; being the initial authentication session and &#39;data&#39; the actual data transfer sessions. | [optional] [readonly] 
**Imsi** | Pointer to **string** | SIM cards are identified on their individual network operators by a unique International Mobile Subscriber Identity (IMSI). &lt;br/&gt; Mobile network operators connect mobile phone calls and communicate with their market SIM cards using their IMSIs. The IMSI is stored in the Subscriber  Identity Module (SIM) inside the device and is sent by the device to the appropriate network. It is used to acquire the details of the device in the Home  Location Register (HLR) or the Visitor Location Register (VLR).  | [optional] [readonly] 
**Imei** | Pointer to **string** | The International Mobile Equipment Identity (or IMEI) is a number, usually unique, that identifies the device currently being used connect to the network. | [optional] [readonly] 
**MobileCountryCode** | Pointer to **string** | It&#39;s a three decimal digit that identifies a country.&lt;br/&gt;&lt;br/&gt; This code is commonly seen joined with a Mobile Network Code (MNC) in a tuple that allows identifying a carrier known as PLMN (Public Land Mobile Network) code. | [optional] [readonly] 
**MobileNetworkCode** | Pointer to **string** | It&#39;s a two to three decimal digits that identify a network.&lt;br/&gt;&lt;br/&gt;  This code is commonly seen joined with a Mobile Country Code (MCC) in a tuple that allows identifying a carrier known as PLMN (Public Land Mobile Network) code. | [optional] [readonly] 
**StartTime** | Pointer to **string** | ISO 8601 formatted date-time indicating when the session started. | [optional] [readonly] 
**StopTime** | Pointer to **string** | ISO 8601 formatted date-time indicating when the session ended. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the record was created. | [optional] [readonly] 
**LastSeen** | Pointer to **string** | ISO 8601 formatted date-time indicating when the last heartbeat to the device was successfully recorded. | [optional] [readonly] 
**Apn** | Pointer to **string** | The Access Point Name (APN) identifies the packet data network that a mobile data user wants to communicate with. | [optional] [readonly] 
**Ipv4** | Pointer to **string** | The SIM&#39;s address in the currently connected network. This IPv4 address is usually obtained dynamically, so it may vary according to the location or new connections.  | [optional] [readonly] 
**Ipv6** | Pointer to **string** | The SIM&#39;s address in the currently connected network. This IPv6 address is usually obtained dynamically, so it may vary according to the location or new connections.  | [optional] [readonly] 
**RadioAccessTechnology** | Pointer to **string** | The radio technology the SIM card used during the session. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the SIM card after when the session happened. | [optional] [readonly] 
**CellId** | Pointer to **string** | The cell ID to which the SIM connected. | [optional] [readonly] 

## Methods

### NewWirelessConnectivityLog

`func NewWirelessConnectivityLog() *WirelessConnectivityLog`

NewWirelessConnectivityLog instantiates a new WirelessConnectivityLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessConnectivityLogWithDefaults

`func NewWirelessConnectivityLogWithDefaults() *WirelessConnectivityLog`

NewWirelessConnectivityLogWithDefaults instantiates a new WirelessConnectivityLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *WirelessConnectivityLog) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *WirelessConnectivityLog) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *WirelessConnectivityLog) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *WirelessConnectivityLog) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *WirelessConnectivityLog) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WirelessConnectivityLog) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WirelessConnectivityLog) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WirelessConnectivityLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSimCardId

`func (o *WirelessConnectivityLog) GetSimCardId() string`

GetSimCardId returns the SimCardId field if non-nil, zero value otherwise.

### GetSimCardIdOk

`func (o *WirelessConnectivityLog) GetSimCardIdOk() (*string, bool)`

GetSimCardIdOk returns a tuple with the SimCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardId

`func (o *WirelessConnectivityLog) SetSimCardId(v string)`

SetSimCardId sets SimCardId field to given value.

### HasSimCardId

`func (o *WirelessConnectivityLog) HasSimCardId() bool`

HasSimCardId returns a boolean if a field has been set.

### GetLogType

`func (o *WirelessConnectivityLog) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *WirelessConnectivityLog) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *WirelessConnectivityLog) SetLogType(v string)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *WirelessConnectivityLog) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### GetImsi

`func (o *WirelessConnectivityLog) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *WirelessConnectivityLog) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *WirelessConnectivityLog) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *WirelessConnectivityLog) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetImei

`func (o *WirelessConnectivityLog) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *WirelessConnectivityLog) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *WirelessConnectivityLog) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *WirelessConnectivityLog) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetMobileCountryCode

`func (o *WirelessConnectivityLog) GetMobileCountryCode() string`

GetMobileCountryCode returns the MobileCountryCode field if non-nil, zero value otherwise.

### GetMobileCountryCodeOk

`func (o *WirelessConnectivityLog) GetMobileCountryCodeOk() (*string, bool)`

GetMobileCountryCodeOk returns a tuple with the MobileCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileCountryCode

`func (o *WirelessConnectivityLog) SetMobileCountryCode(v string)`

SetMobileCountryCode sets MobileCountryCode field to given value.

### HasMobileCountryCode

`func (o *WirelessConnectivityLog) HasMobileCountryCode() bool`

HasMobileCountryCode returns a boolean if a field has been set.

### GetMobileNetworkCode

`func (o *WirelessConnectivityLog) GetMobileNetworkCode() string`

GetMobileNetworkCode returns the MobileNetworkCode field if non-nil, zero value otherwise.

### GetMobileNetworkCodeOk

`func (o *WirelessConnectivityLog) GetMobileNetworkCodeOk() (*string, bool)`

GetMobileNetworkCodeOk returns a tuple with the MobileNetworkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNetworkCode

`func (o *WirelessConnectivityLog) SetMobileNetworkCode(v string)`

SetMobileNetworkCode sets MobileNetworkCode field to given value.

### HasMobileNetworkCode

`func (o *WirelessConnectivityLog) HasMobileNetworkCode() bool`

HasMobileNetworkCode returns a boolean if a field has been set.

### GetStartTime

`func (o *WirelessConnectivityLog) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WirelessConnectivityLog) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WirelessConnectivityLog) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WirelessConnectivityLog) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStopTime

`func (o *WirelessConnectivityLog) GetStopTime() string`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *WirelessConnectivityLog) GetStopTimeOk() (*string, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *WirelessConnectivityLog) SetStopTime(v string)`

SetStopTime sets StopTime field to given value.

### HasStopTime

`func (o *WirelessConnectivityLog) HasStopTime() bool`

HasStopTime returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WirelessConnectivityLog) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WirelessConnectivityLog) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WirelessConnectivityLog) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WirelessConnectivityLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastSeen

`func (o *WirelessConnectivityLog) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *WirelessConnectivityLog) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *WirelessConnectivityLog) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *WirelessConnectivityLog) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetApn

`func (o *WirelessConnectivityLog) GetApn() string`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *WirelessConnectivityLog) GetApnOk() (*string, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *WirelessConnectivityLog) SetApn(v string)`

SetApn sets Apn field to given value.

### HasApn

`func (o *WirelessConnectivityLog) HasApn() bool`

HasApn returns a boolean if a field has been set.

### GetIpv4

`func (o *WirelessConnectivityLog) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *WirelessConnectivityLog) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *WirelessConnectivityLog) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *WirelessConnectivityLog) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *WirelessConnectivityLog) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *WirelessConnectivityLog) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *WirelessConnectivityLog) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *WirelessConnectivityLog) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRadioAccessTechnology

`func (o *WirelessConnectivityLog) GetRadioAccessTechnology() string`

GetRadioAccessTechnology returns the RadioAccessTechnology field if non-nil, zero value otherwise.

### GetRadioAccessTechnologyOk

`func (o *WirelessConnectivityLog) GetRadioAccessTechnologyOk() (*string, bool)`

GetRadioAccessTechnologyOk returns a tuple with the RadioAccessTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioAccessTechnology

`func (o *WirelessConnectivityLog) SetRadioAccessTechnology(v string)`

SetRadioAccessTechnology sets RadioAccessTechnology field to given value.

### HasRadioAccessTechnology

`func (o *WirelessConnectivityLog) HasRadioAccessTechnology() bool`

HasRadioAccessTechnology returns a boolean if a field has been set.

### GetState

`func (o *WirelessConnectivityLog) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WirelessConnectivityLog) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WirelessConnectivityLog) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *WirelessConnectivityLog) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCellId

`func (o *WirelessConnectivityLog) GetCellId() string`

GetCellId returns the CellId field if non-nil, zero value otherwise.

### GetCellIdOk

`func (o *WirelessConnectivityLog) GetCellIdOk() (*string, bool)`

GetCellIdOk returns a tuple with the CellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellId

`func (o *WirelessConnectivityLog) SetCellId(v string)`

SetCellId sets CellId field to given value.

### HasCellId

`func (o *WirelessConnectivityLog) HasCellId() bool`

HasCellId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



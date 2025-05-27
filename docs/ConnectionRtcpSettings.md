# ConnectionRtcpSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | RTCP port by default is rtp+1, it can also be set to rtcp-mux | [optional] [default to "rtp+1"]
**CaptureEnabled** | Pointer to **bool** | BETA - Enable the capture and storage of RTCP messages to create QoS reports on the Telnyx Mission Control Portal. | [optional] [default to false]
**ReportFrequencySecs** | Pointer to **int32** | RTCP reports are sent to customers based on the frequency set. Frequency is in seconds and it can be set to values from 5 to 3000 seconds. | [optional] [default to 5]

## Methods

### NewConnectionRtcpSettings

`func NewConnectionRtcpSettings() *ConnectionRtcpSettings`

NewConnectionRtcpSettings instantiates a new ConnectionRtcpSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionRtcpSettingsWithDefaults

`func NewConnectionRtcpSettingsWithDefaults() *ConnectionRtcpSettings`

NewConnectionRtcpSettingsWithDefaults instantiates a new ConnectionRtcpSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ConnectionRtcpSettings) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ConnectionRtcpSettings) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ConnectionRtcpSettings) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ConnectionRtcpSettings) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCaptureEnabled

`func (o *ConnectionRtcpSettings) GetCaptureEnabled() bool`

GetCaptureEnabled returns the CaptureEnabled field if non-nil, zero value otherwise.

### GetCaptureEnabledOk

`func (o *ConnectionRtcpSettings) GetCaptureEnabledOk() (*bool, bool)`

GetCaptureEnabledOk returns a tuple with the CaptureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureEnabled

`func (o *ConnectionRtcpSettings) SetCaptureEnabled(v bool)`

SetCaptureEnabled sets CaptureEnabled field to given value.

### HasCaptureEnabled

`func (o *ConnectionRtcpSettings) HasCaptureEnabled() bool`

HasCaptureEnabled returns a boolean if a field has been set.

### GetReportFrequencySecs

`func (o *ConnectionRtcpSettings) GetReportFrequencySecs() int32`

GetReportFrequencySecs returns the ReportFrequencySecs field if non-nil, zero value otherwise.

### GetReportFrequencySecsOk

`func (o *ConnectionRtcpSettings) GetReportFrequencySecsOk() (*int32, bool)`

GetReportFrequencySecsOk returns a tuple with the ReportFrequencySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFrequencySecs

`func (o *ConnectionRtcpSettings) SetReportFrequencySecs(v int32)`

SetReportFrequencySecs sets ReportFrequencySecs field to given value.

### HasReportFrequencySecs

`func (o *ConnectionRtcpSettings) HasReportFrequencySecs() bool`

HasReportFrequencySecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



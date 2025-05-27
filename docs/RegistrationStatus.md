# RegistrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Status** | Pointer to **string** | The current registration status of your SIP connection | [optional] 
**SipUsername** | Pointer to **string** | The user name of the SIP connection | [optional] 
**IpAddress** | Pointer to **string** | The ip used during the SIP connection | [optional] 
**Transport** | Pointer to **string** | The protocol of the SIP connection | [optional] 
**Port** | Pointer to **int32** | The port of the SIP connection | [optional] 
**UserAgent** | Pointer to **string** | The user agent of the SIP connection | [optional] 
**LastRegistration** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was last updated. | [optional] 

## Methods

### NewRegistrationStatus

`func NewRegistrationStatus() *RegistrationStatus`

NewRegistrationStatus instantiates a new RegistrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationStatusWithDefaults

`func NewRegistrationStatusWithDefaults() *RegistrationStatus`

NewRegistrationStatusWithDefaults instantiates a new RegistrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *RegistrationStatus) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *RegistrationStatus) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *RegistrationStatus) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *RegistrationStatus) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetStatus

`func (o *RegistrationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegistrationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegistrationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegistrationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSipUsername

`func (o *RegistrationStatus) GetSipUsername() string`

GetSipUsername returns the SipUsername field if non-nil, zero value otherwise.

### GetSipUsernameOk

`func (o *RegistrationStatus) GetSipUsernameOk() (*string, bool)`

GetSipUsernameOk returns a tuple with the SipUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipUsername

`func (o *RegistrationStatus) SetSipUsername(v string)`

SetSipUsername sets SipUsername field to given value.

### HasSipUsername

`func (o *RegistrationStatus) HasSipUsername() bool`

HasSipUsername returns a boolean if a field has been set.

### GetIpAddress

`func (o *RegistrationStatus) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *RegistrationStatus) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *RegistrationStatus) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *RegistrationStatus) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetTransport

`func (o *RegistrationStatus) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *RegistrationStatus) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *RegistrationStatus) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *RegistrationStatus) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetPort

`func (o *RegistrationStatus) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RegistrationStatus) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RegistrationStatus) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RegistrationStatus) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUserAgent

`func (o *RegistrationStatus) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *RegistrationStatus) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *RegistrationStatus) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *RegistrationStatus) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetLastRegistration

`func (o *RegistrationStatus) GetLastRegistration() string`

GetLastRegistration returns the LastRegistration field if non-nil, zero value otherwise.

### GetLastRegistrationOk

`func (o *RegistrationStatus) GetLastRegistrationOk() (*string, bool)`

GetLastRegistrationOk returns a tuple with the LastRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRegistration

`func (o *RegistrationStatus) SetLastRegistration(v string)`

SetLastRegistration sets LastRegistration field to given value.

### HasLastRegistration

`func (o *RegistrationStatus) HasLastRegistration() bool`

HasLastRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



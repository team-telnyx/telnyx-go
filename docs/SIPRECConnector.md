# SiprecConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname/IPv4 address of the SIPREC SRS. | 
**Port** | **int32** | Port for the SIPREC SRS. | 
**Name** | **string** | Name for the SIPREC connector resource. | 
**AppSubdomain** | Pointer to **string** | Subdomain to route the call when using Telnyx SRS (optional for non-Telnyx SRS). | [optional] 

## Methods

### NewSiprecConnector

`func NewSiprecConnector(host string, port int32, name string, ) *SiprecConnector`

NewSiprecConnector instantiates a new SiprecConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiprecConnectorWithDefaults

`func NewSiprecConnectorWithDefaults() *SiprecConnector`

NewSiprecConnectorWithDefaults instantiates a new SiprecConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SiprecConnector) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SiprecConnector) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SiprecConnector) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *SiprecConnector) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SiprecConnector) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SiprecConnector) SetPort(v int32)`

SetPort sets Port field to given value.


### GetName

`func (o *SiprecConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiprecConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiprecConnector) SetName(v string)`

SetName sets Name field to given value.


### GetAppSubdomain

`func (o *SiprecConnector) GetAppSubdomain() string`

GetAppSubdomain returns the AppSubdomain field if non-nil, zero value otherwise.

### GetAppSubdomainOk

`func (o *SiprecConnector) GetAppSubdomainOk() (*string, bool)`

GetAppSubdomainOk returns a tuple with the AppSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSubdomain

`func (o *SiprecConnector) SetAppSubdomain(v string)`

SetAppSubdomain sets AppSubdomain field to given value.

### HasAppSubdomain

`func (o *SiprecConnector) HasAppSubdomain() bool`

HasAppSubdomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



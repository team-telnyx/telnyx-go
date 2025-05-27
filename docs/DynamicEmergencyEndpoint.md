# DynamicEmergencyEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**DynamicEmergencyAddressId** | **string** | An id of a currently active dynamic emergency location. | 
**Status** | Pointer to **string** | Status of dynamic emergency address | [optional] [readonly] 
**SipFromId** | Pointer to **string** |  | [optional] [readonly] 
**CallbackNumber** | **string** |  | 
**CallerName** | **string** |  | 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date of when the resource was created | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date of when the resource was last updated | [optional] [readonly] 

## Methods

### NewDynamicEmergencyEndpoint

`func NewDynamicEmergencyEndpoint(dynamicEmergencyAddressId string, callbackNumber string, callerName string, ) *DynamicEmergencyEndpoint`

NewDynamicEmergencyEndpoint instantiates a new DynamicEmergencyEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicEmergencyEndpointWithDefaults

`func NewDynamicEmergencyEndpointWithDefaults() *DynamicEmergencyEndpoint`

NewDynamicEmergencyEndpointWithDefaults instantiates a new DynamicEmergencyEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DynamicEmergencyEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DynamicEmergencyEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DynamicEmergencyEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DynamicEmergencyEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *DynamicEmergencyEndpoint) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *DynamicEmergencyEndpoint) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *DynamicEmergencyEndpoint) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *DynamicEmergencyEndpoint) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetDynamicEmergencyAddressId

`func (o *DynamicEmergencyEndpoint) GetDynamicEmergencyAddressId() string`

GetDynamicEmergencyAddressId returns the DynamicEmergencyAddressId field if non-nil, zero value otherwise.

### GetDynamicEmergencyAddressIdOk

`func (o *DynamicEmergencyEndpoint) GetDynamicEmergencyAddressIdOk() (*string, bool)`

GetDynamicEmergencyAddressIdOk returns a tuple with the DynamicEmergencyAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicEmergencyAddressId

`func (o *DynamicEmergencyEndpoint) SetDynamicEmergencyAddressId(v string)`

SetDynamicEmergencyAddressId sets DynamicEmergencyAddressId field to given value.


### GetStatus

`func (o *DynamicEmergencyEndpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DynamicEmergencyEndpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DynamicEmergencyEndpoint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DynamicEmergencyEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSipFromId

`func (o *DynamicEmergencyEndpoint) GetSipFromId() string`

GetSipFromId returns the SipFromId field if non-nil, zero value otherwise.

### GetSipFromIdOk

`func (o *DynamicEmergencyEndpoint) GetSipFromIdOk() (*string, bool)`

GetSipFromIdOk returns a tuple with the SipFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipFromId

`func (o *DynamicEmergencyEndpoint) SetSipFromId(v string)`

SetSipFromId sets SipFromId field to given value.

### HasSipFromId

`func (o *DynamicEmergencyEndpoint) HasSipFromId() bool`

HasSipFromId returns a boolean if a field has been set.

### GetCallbackNumber

`func (o *DynamicEmergencyEndpoint) GetCallbackNumber() string`

GetCallbackNumber returns the CallbackNumber field if non-nil, zero value otherwise.

### GetCallbackNumberOk

`func (o *DynamicEmergencyEndpoint) GetCallbackNumberOk() (*string, bool)`

GetCallbackNumberOk returns a tuple with the CallbackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackNumber

`func (o *DynamicEmergencyEndpoint) SetCallbackNumber(v string)`

SetCallbackNumber sets CallbackNumber field to given value.


### GetCallerName

`func (o *DynamicEmergencyEndpoint) GetCallerName() string`

GetCallerName returns the CallerName field if non-nil, zero value otherwise.

### GetCallerNameOk

`func (o *DynamicEmergencyEndpoint) GetCallerNameOk() (*string, bool)`

GetCallerNameOk returns a tuple with the CallerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerName

`func (o *DynamicEmergencyEndpoint) SetCallerName(v string)`

SetCallerName sets CallerName field to given value.


### GetCreatedAt

`func (o *DynamicEmergencyEndpoint) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DynamicEmergencyEndpoint) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DynamicEmergencyEndpoint) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DynamicEmergencyEndpoint) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DynamicEmergencyEndpoint) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DynamicEmergencyEndpoint) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DynamicEmergencyEndpoint) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DynamicEmergencyEndpoint) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



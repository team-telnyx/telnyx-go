# BookAppointmentToolParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTypeId** | **int32** | Event Type ID for which slots are being fetched. [cal.com](https://cal.com/docs/api-reference/v2/bookings/create-a-booking#body-event-type-id) | 
**ApiKeyRef** | **string** | Reference to an integration secret that contains your Cal.com API key. You would pass the &#x60;identifier&#x60; for an integration secret [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) that refers to your Cal.com API key. | 
**AttendeeName** | Pointer to **string** | The name of the attendee [cal.com](https://cal.com/docs/api-reference/v2/bookings/create-a-booking#body-attendee-name). If not provided, the assistant will ask for the attendee&#39;s name. | [optional] 
**AttendeeTimezone** | Pointer to **string** | The timezone of the attendee [cal.com](https://cal.com/docs/api-reference/v2/bookings/create-a-booking#body-attendee-timezone). If not provided, the assistant will ask for the attendee&#39;s timezone. | [optional] 

## Methods

### NewBookAppointmentToolParams

`func NewBookAppointmentToolParams(eventTypeId int32, apiKeyRef string, ) *BookAppointmentToolParams`

NewBookAppointmentToolParams instantiates a new BookAppointmentToolParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookAppointmentToolParamsWithDefaults

`func NewBookAppointmentToolParamsWithDefaults() *BookAppointmentToolParams`

NewBookAppointmentToolParamsWithDefaults instantiates a new BookAppointmentToolParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTypeId

`func (o *BookAppointmentToolParams) GetEventTypeId() int32`

GetEventTypeId returns the EventTypeId field if non-nil, zero value otherwise.

### GetEventTypeIdOk

`func (o *BookAppointmentToolParams) GetEventTypeIdOk() (*int32, bool)`

GetEventTypeIdOk returns a tuple with the EventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeId

`func (o *BookAppointmentToolParams) SetEventTypeId(v int32)`

SetEventTypeId sets EventTypeId field to given value.


### GetApiKeyRef

`func (o *BookAppointmentToolParams) GetApiKeyRef() string`

GetApiKeyRef returns the ApiKeyRef field if non-nil, zero value otherwise.

### GetApiKeyRefOk

`func (o *BookAppointmentToolParams) GetApiKeyRefOk() (*string, bool)`

GetApiKeyRefOk returns a tuple with the ApiKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyRef

`func (o *BookAppointmentToolParams) SetApiKeyRef(v string)`

SetApiKeyRef sets ApiKeyRef field to given value.


### GetAttendeeName

`func (o *BookAppointmentToolParams) GetAttendeeName() string`

GetAttendeeName returns the AttendeeName field if non-nil, zero value otherwise.

### GetAttendeeNameOk

`func (o *BookAppointmentToolParams) GetAttendeeNameOk() (*string, bool)`

GetAttendeeNameOk returns a tuple with the AttendeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendeeName

`func (o *BookAppointmentToolParams) SetAttendeeName(v string)`

SetAttendeeName sets AttendeeName field to given value.

### HasAttendeeName

`func (o *BookAppointmentToolParams) HasAttendeeName() bool`

HasAttendeeName returns a boolean if a field has been set.

### GetAttendeeTimezone

`func (o *BookAppointmentToolParams) GetAttendeeTimezone() string`

GetAttendeeTimezone returns the AttendeeTimezone field if non-nil, zero value otherwise.

### GetAttendeeTimezoneOk

`func (o *BookAppointmentToolParams) GetAttendeeTimezoneOk() (*string, bool)`

GetAttendeeTimezoneOk returns a tuple with the AttendeeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendeeTimezone

`func (o *BookAppointmentToolParams) SetAttendeeTimezone(v string)`

SetAttendeeTimezone sets AttendeeTimezone field to given value.

### HasAttendeeTimezone

`func (o *BookAppointmentToolParams) HasAttendeeTimezone() bool`

HasAttendeeTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



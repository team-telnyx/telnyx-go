# CheckAvailabilityToolParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTypeId** | **int32** | Event Type ID for which slots are being fetched. [cal.com](https://cal.com/docs/api-reference/v2/slots/get-available-slots#parameter-event-type-id) | 
**ApiKeyRef** | **string** | Reference to an integration secret that contains your Cal.com API key. You would pass the &#x60;identifier&#x60; for an integration secret [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) that refers to your Cal.com API key. | 

## Methods

### NewCheckAvailabilityToolParams

`func NewCheckAvailabilityToolParams(eventTypeId int32, apiKeyRef string, ) *CheckAvailabilityToolParams`

NewCheckAvailabilityToolParams instantiates a new CheckAvailabilityToolParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAvailabilityToolParamsWithDefaults

`func NewCheckAvailabilityToolParamsWithDefaults() *CheckAvailabilityToolParams`

NewCheckAvailabilityToolParamsWithDefaults instantiates a new CheckAvailabilityToolParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTypeId

`func (o *CheckAvailabilityToolParams) GetEventTypeId() int32`

GetEventTypeId returns the EventTypeId field if non-nil, zero value otherwise.

### GetEventTypeIdOk

`func (o *CheckAvailabilityToolParams) GetEventTypeIdOk() (*int32, bool)`

GetEventTypeIdOk returns a tuple with the EventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeId

`func (o *CheckAvailabilityToolParams) SetEventTypeId(v int32)`

SetEventTypeId sets EventTypeId field to given value.


### GetApiKeyRef

`func (o *CheckAvailabilityToolParams) GetApiKeyRef() string`

GetApiKeyRef returns the ApiKeyRef field if non-nil, zero value otherwise.

### GetApiKeyRefOk

`func (o *CheckAvailabilityToolParams) GetApiKeyRefOk() (*string, bool)`

GetApiKeyRefOk returns a tuple with the ApiKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyRef

`func (o *CheckAvailabilityToolParams) SetApiKeyRef(v string)`

SetApiKeyRef sets ApiKeyRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



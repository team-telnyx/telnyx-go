# ActionsParticipantsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Participants** | Pointer to [**ActionsParticipantsRequestParticipants**](ActionsParticipantsRequestParticipants.md) |  | [optional] 
**Exclude** | Pointer to **[]string** | List of participant id to exclude from the action. | [optional] 

## Methods

### NewActionsParticipantsRequest

`func NewActionsParticipantsRequest() *ActionsParticipantsRequest`

NewActionsParticipantsRequest instantiates a new ActionsParticipantsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsParticipantsRequestWithDefaults

`func NewActionsParticipantsRequestWithDefaults() *ActionsParticipantsRequest`

NewActionsParticipantsRequestWithDefaults instantiates a new ActionsParticipantsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParticipants

`func (o *ActionsParticipantsRequest) GetParticipants() ActionsParticipantsRequestParticipants`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ActionsParticipantsRequest) GetParticipantsOk() (*ActionsParticipantsRequestParticipants, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ActionsParticipantsRequest) SetParticipants(v ActionsParticipantsRequestParticipants)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *ActionsParticipantsRequest) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetExclude

`func (o *ActionsParticipantsRequest) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *ActionsParticipantsRequest) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *ActionsParticipantsRequest) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *ActionsParticipantsRequest) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



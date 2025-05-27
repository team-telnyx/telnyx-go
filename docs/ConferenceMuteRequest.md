# ConferenceMuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlIds** | Pointer to **[]string** | Array of unique identifiers and tokens for controlling the call. When empty all participants will be muted. | [optional] 

## Methods

### NewConferenceMuteRequest

`func NewConferenceMuteRequest() *ConferenceMuteRequest`

NewConferenceMuteRequest instantiates a new ConferenceMuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceMuteRequestWithDefaults

`func NewConferenceMuteRequestWithDefaults() *ConferenceMuteRequest`

NewConferenceMuteRequestWithDefaults instantiates a new ConferenceMuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlIds

`func (o *ConferenceMuteRequest) GetCallControlIds() []string`

GetCallControlIds returns the CallControlIds field if non-nil, zero value otherwise.

### GetCallControlIdsOk

`func (o *ConferenceMuteRequest) GetCallControlIdsOk() (*[]string, bool)`

GetCallControlIdsOk returns a tuple with the CallControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlIds

`func (o *ConferenceMuteRequest) SetCallControlIds(v []string)`

SetCallControlIds sets CallControlIds field to given value.

### HasCallControlIds

`func (o *ConferenceMuteRequest) HasCallControlIds() bool`

HasCallControlIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



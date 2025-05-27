# Attempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** | ISO 8601 timestamp indicating when the attempt was initiated. | [optional] 
**FinishedAt** | Pointer to **time.Time** | ISO 8601 timestamp indicating when the attempt has finished. | [optional] 
**Http** | Pointer to [**Http**](Http.md) |  | [optional] 
**Errors** | Pointer to **[]int32** | Webhook delivery error codes. | [optional] 

## Methods

### NewAttempt

`func NewAttempt() *Attempt`

NewAttempt instantiates a new Attempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttemptWithDefaults

`func NewAttemptWithDefaults() *Attempt`

NewAttemptWithDefaults instantiates a new Attempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Attempt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Attempt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Attempt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Attempt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *Attempt) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Attempt) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Attempt) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Attempt) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *Attempt) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *Attempt) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *Attempt) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *Attempt) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetHttp

`func (o *Attempt) GetHttp() Http`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *Attempt) GetHttpOk() (*Http, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *Attempt) SetHttp(v Http)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *Attempt) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetErrors

`func (o *Attempt) GetErrors() []int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Attempt) GetErrorsOk() (*[]int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Attempt) SetErrors(v []int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Attempt) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



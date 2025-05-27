# AutoRespConfigCreateSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Keywords** | **[]string** |  | 
**RespText** | Pointer to **string** |  | [optional] 
**CountryCode** | **string** |  | 

## Methods

### NewAutoRespConfigCreateSchema

`func NewAutoRespConfigCreateSchema(op string, keywords []string, countryCode string, ) *AutoRespConfigCreateSchema`

NewAutoRespConfigCreateSchema instantiates a new AutoRespConfigCreateSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRespConfigCreateSchemaWithDefaults

`func NewAutoRespConfigCreateSchemaWithDefaults() *AutoRespConfigCreateSchema`

NewAutoRespConfigCreateSchemaWithDefaults instantiates a new AutoRespConfigCreateSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *AutoRespConfigCreateSchema) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AutoRespConfigCreateSchema) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AutoRespConfigCreateSchema) SetOp(v string)`

SetOp sets Op field to given value.


### GetKeywords

`func (o *AutoRespConfigCreateSchema) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *AutoRespConfigCreateSchema) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *AutoRespConfigCreateSchema) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.


### GetRespText

`func (o *AutoRespConfigCreateSchema) GetRespText() string`

GetRespText returns the RespText field if non-nil, zero value otherwise.

### GetRespTextOk

`func (o *AutoRespConfigCreateSchema) GetRespTextOk() (*string, bool)`

GetRespTextOk returns a tuple with the RespText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespText

`func (o *AutoRespConfigCreateSchema) SetRespText(v string)`

SetRespText sets RespText field to given value.

### HasRespText

`func (o *AutoRespConfigCreateSchema) HasRespText() bool`

HasRespText returns a boolean if a field has been set.

### GetCountryCode

`func (o *AutoRespConfigCreateSchema) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AutoRespConfigCreateSchema) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AutoRespConfigCreateSchema) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



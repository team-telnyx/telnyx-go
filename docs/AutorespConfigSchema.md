# AutorespConfigSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Keywords** | **[]string** |  | 
**RespText** | Pointer to **string** |  | [optional] 
**CountryCode** | **string** |  | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAutorespConfigSchema

`func NewAutorespConfigSchema(op string, keywords []string, countryCode string, id string, createdAt time.Time, updatedAt time.Time, ) *AutorespConfigSchema`

NewAutorespConfigSchema instantiates a new AutorespConfigSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutorespConfigSchemaWithDefaults

`func NewAutorespConfigSchemaWithDefaults() *AutorespConfigSchema`

NewAutorespConfigSchemaWithDefaults instantiates a new AutorespConfigSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *AutorespConfigSchema) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AutorespConfigSchema) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AutorespConfigSchema) SetOp(v string)`

SetOp sets Op field to given value.


### GetKeywords

`func (o *AutorespConfigSchema) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *AutorespConfigSchema) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *AutorespConfigSchema) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.


### GetRespText

`func (o *AutorespConfigSchema) GetRespText() string`

GetRespText returns the RespText field if non-nil, zero value otherwise.

### GetRespTextOk

`func (o *AutorespConfigSchema) GetRespTextOk() (*string, bool)`

GetRespTextOk returns a tuple with the RespText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespText

`func (o *AutorespConfigSchema) SetRespText(v string)`

SetRespText sets RespText field to given value.

### HasRespText

`func (o *AutorespConfigSchema) HasRespText() bool`

HasRespText returns a boolean if a field has been set.

### GetCountryCode

`func (o *AutorespConfigSchema) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AutorespConfigSchema) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AutorespConfigSchema) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetId

`func (o *AutorespConfigSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutorespConfigSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutorespConfigSchema) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *AutorespConfigSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutorespConfigSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutorespConfigSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AutorespConfigSchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AutorespConfigSchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AutorespConfigSchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



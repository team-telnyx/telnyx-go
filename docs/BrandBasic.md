# BrandBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandId** | Pointer to **string** | Unique identifier assigned to the brand. | [optional] 
**TcrBrandId** | Pointer to **string** | Unique identifier assigned to the brand by the registry. | [optional] 
**EntityType** | Pointer to [**EntityType**](EntityType.md) |  | [optional] 
**IdentityStatus** | Pointer to [**BrandIdentityStatus**](BrandIdentityStatus.md) |  | [optional] 
**CompanyName** | Pointer to **string** | (Required for Non-profit/private/public) Legal company name. | [optional] 
**DisplayName** | Pointer to **string** | Display or marketing name of the brand. | [optional] 
**Email** | Pointer to **string** | Valid email address of brand support contact. | [optional] 
**Website** | Pointer to **string** | Brand website URL. | [optional] 
**FailureReasons** | Pointer to **interface{}** | Failure reasons for brand | [optional] 
**Status** | Pointer to **string** | Status of the brand | [optional] 
**CreatedAt** | Pointer to **string** | Date and time that the brand was created at. | [optional] 
**UpdatedAt** | Pointer to **string** | Date and time that the brand was last updated at. | [optional] 
**AssignedCampaingsCount** | Pointer to **int32** | Number of campaigns associated with the brand | [optional] 

## Methods

### NewBrandBasic

`func NewBrandBasic() *BrandBasic`

NewBrandBasic instantiates a new BrandBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandBasicWithDefaults

`func NewBrandBasicWithDefaults() *BrandBasic`

NewBrandBasicWithDefaults instantiates a new BrandBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandId

`func (o *BrandBasic) GetBrandId() string`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *BrandBasic) GetBrandIdOk() (*string, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *BrandBasic) SetBrandId(v string)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *BrandBasic) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetTcrBrandId

`func (o *BrandBasic) GetTcrBrandId() string`

GetTcrBrandId returns the TcrBrandId field if non-nil, zero value otherwise.

### GetTcrBrandIdOk

`func (o *BrandBasic) GetTcrBrandIdOk() (*string, bool)`

GetTcrBrandIdOk returns a tuple with the TcrBrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrBrandId

`func (o *BrandBasic) SetTcrBrandId(v string)`

SetTcrBrandId sets TcrBrandId field to given value.

### HasTcrBrandId

`func (o *BrandBasic) HasTcrBrandId() bool`

HasTcrBrandId returns a boolean if a field has been set.

### GetEntityType

`func (o *BrandBasic) GetEntityType() EntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *BrandBasic) GetEntityTypeOk() (*EntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *BrandBasic) SetEntityType(v EntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *BrandBasic) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetIdentityStatus

`func (o *BrandBasic) GetIdentityStatus() BrandIdentityStatus`

GetIdentityStatus returns the IdentityStatus field if non-nil, zero value otherwise.

### GetIdentityStatusOk

`func (o *BrandBasic) GetIdentityStatusOk() (*BrandIdentityStatus, bool)`

GetIdentityStatusOk returns a tuple with the IdentityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStatus

`func (o *BrandBasic) SetIdentityStatus(v BrandIdentityStatus)`

SetIdentityStatus sets IdentityStatus field to given value.

### HasIdentityStatus

`func (o *BrandBasic) HasIdentityStatus() bool`

HasIdentityStatus returns a boolean if a field has been set.

### GetCompanyName

`func (o *BrandBasic) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *BrandBasic) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *BrandBasic) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *BrandBasic) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetDisplayName

`func (o *BrandBasic) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BrandBasic) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BrandBasic) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BrandBasic) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *BrandBasic) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BrandBasic) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BrandBasic) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BrandBasic) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetWebsite

`func (o *BrandBasic) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *BrandBasic) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *BrandBasic) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *BrandBasic) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetFailureReasons

`func (o *BrandBasic) GetFailureReasons() interface{}`

GetFailureReasons returns the FailureReasons field if non-nil, zero value otherwise.

### GetFailureReasonsOk

`func (o *BrandBasic) GetFailureReasonsOk() (*interface{}, bool)`

GetFailureReasonsOk returns a tuple with the FailureReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReasons

`func (o *BrandBasic) SetFailureReasons(v interface{})`

SetFailureReasons sets FailureReasons field to given value.

### HasFailureReasons

`func (o *BrandBasic) HasFailureReasons() bool`

HasFailureReasons returns a boolean if a field has been set.

### SetFailureReasonsNil

`func (o *BrandBasic) SetFailureReasonsNil(b bool)`

 SetFailureReasonsNil sets the value for FailureReasons to be an explicit nil

### UnsetFailureReasons
`func (o *BrandBasic) UnsetFailureReasons()`

UnsetFailureReasons ensures that no value is present for FailureReasons, not even an explicit nil
### GetStatus

`func (o *BrandBasic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BrandBasic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BrandBasic) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BrandBasic) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BrandBasic) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BrandBasic) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BrandBasic) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BrandBasic) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BrandBasic) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BrandBasic) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BrandBasic) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BrandBasic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAssignedCampaingsCount

`func (o *BrandBasic) GetAssignedCampaingsCount() int32`

GetAssignedCampaingsCount returns the AssignedCampaingsCount field if non-nil, zero value otherwise.

### GetAssignedCampaingsCountOk

`func (o *BrandBasic) GetAssignedCampaingsCountOk() (*int32, bool)`

GetAssignedCampaingsCountOk returns a tuple with the AssignedCampaingsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedCampaingsCount

`func (o *BrandBasic) SetAssignedCampaingsCount(v int32)`

SetAssignedCampaingsCount sets AssignedCampaingsCount field to given value.

### HasAssignedCampaingsCount

`func (o *BrandBasic) HasAssignedCampaingsCount() bool`

HasAssignedCampaingsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



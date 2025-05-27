# CampaignStatusUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandId** | Pointer to **string** | Brand ID associated with the campaign. | [optional] 
**CampaignId** | Pointer to **string** | The ID of the campaign. | [optional] 
**CreateDate** | Pointer to **string** | Unix timestamp when campaign was created. | [optional] 
**CspId** | Pointer to **string** | Alphanumeric identifier of the CSP associated with this campaign. | [optional] 
**IsTMobileRegistered** | Pointer to **bool** | Indicates whether the campaign is registered with T-Mobile. | [optional] 

## Methods

### NewCampaignStatusUpdateEvent

`func NewCampaignStatusUpdateEvent() *CampaignStatusUpdateEvent`

NewCampaignStatusUpdateEvent instantiates a new CampaignStatusUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignStatusUpdateEventWithDefaults

`func NewCampaignStatusUpdateEventWithDefaults() *CampaignStatusUpdateEvent`

NewCampaignStatusUpdateEventWithDefaults instantiates a new CampaignStatusUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandId

`func (o *CampaignStatusUpdateEvent) GetBrandId() string`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *CampaignStatusUpdateEvent) GetBrandIdOk() (*string, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *CampaignStatusUpdateEvent) SetBrandId(v string)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *CampaignStatusUpdateEvent) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetCampaignId

`func (o *CampaignStatusUpdateEvent) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignStatusUpdateEvent) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CampaignStatusUpdateEvent) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *CampaignStatusUpdateEvent) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetCreateDate

`func (o *CampaignStatusUpdateEvent) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *CampaignStatusUpdateEvent) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *CampaignStatusUpdateEvent) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *CampaignStatusUpdateEvent) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetCspId

`func (o *CampaignStatusUpdateEvent) GetCspId() string`

GetCspId returns the CspId field if non-nil, zero value otherwise.

### GetCspIdOk

`func (o *CampaignStatusUpdateEvent) GetCspIdOk() (*string, bool)`

GetCspIdOk returns a tuple with the CspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspId

`func (o *CampaignStatusUpdateEvent) SetCspId(v string)`

SetCspId sets CspId field to given value.

### HasCspId

`func (o *CampaignStatusUpdateEvent) HasCspId() bool`

HasCspId returns a boolean if a field has been set.

### GetIsTMobileRegistered

`func (o *CampaignStatusUpdateEvent) GetIsTMobileRegistered() bool`

GetIsTMobileRegistered returns the IsTMobileRegistered field if non-nil, zero value otherwise.

### GetIsTMobileRegisteredOk

`func (o *CampaignStatusUpdateEvent) GetIsTMobileRegisteredOk() (*bool, bool)`

GetIsTMobileRegisteredOk returns a tuple with the IsTMobileRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTMobileRegistered

`func (o *CampaignStatusUpdateEvent) SetIsTMobileRegistered(v bool)`

SetIsTMobileRegistered sets IsTMobileRegistered field to given value.

### HasIsTMobileRegistered

`func (o *CampaignStatusUpdateEvent) HasIsTMobileRegistered() bool`

HasIsTMobileRegistered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



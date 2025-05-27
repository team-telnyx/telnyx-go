# PhoneNumberBundleStatusChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | **string** | The new bundle_id setting for the number. If you are assigning the number to a bundle, this is the unique ID of the bundle you wish to use. If you are removing the number from a bundle, this must be null. You cannot assign a number from one bundle to another directly. You must first remove it from a bundle, and then assign it to a new bundle. | 

## Methods

### NewPhoneNumberBundleStatusChangeRequest

`func NewPhoneNumberBundleStatusChangeRequest(bundleId string, ) *PhoneNumberBundleStatusChangeRequest`

NewPhoneNumberBundleStatusChangeRequest instantiates a new PhoneNumberBundleStatusChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberBundleStatusChangeRequestWithDefaults

`func NewPhoneNumberBundleStatusChangeRequestWithDefaults() *PhoneNumberBundleStatusChangeRequest`

NewPhoneNumberBundleStatusChangeRequestWithDefaults instantiates a new PhoneNumberBundleStatusChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *PhoneNumberBundleStatusChangeRequest) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *PhoneNumberBundleStatusChangeRequest) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *PhoneNumberBundleStatusChangeRequest) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



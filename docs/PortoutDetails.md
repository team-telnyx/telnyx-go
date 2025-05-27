# PortoutDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**PhoneNumbers** | Pointer to **[]string** | Phone numbers associated with this portout | [optional] 
**AuthorizedName** | Pointer to **string** | Name of person authorizing the porting order | [optional] 
**CarrierName** | Pointer to **string** | Carrier the number will be ported out to | [optional] 
**CurrentCarrier** | Pointer to **string** | The current carrier | [optional] 
**EndUserName** | Pointer to **string** | Person name or company name requesting the port | [optional] 
**City** | Pointer to **string** | City or municipality of billing address | [optional] 
**State** | Pointer to **string** | State, province, or similar of billing address | [optional] 
**Zip** | Pointer to **string** | Postal Code of billing address | [optional] 
**Lsr** | Pointer to **[]string** | The Local Service Request | [optional] 
**Pon** | Pointer to **string** | Port order number assigned by the carrier the number will be ported out to | [optional] 
**Reason** | Pointer to **string** | The reason why the order is being rejected by the user. If the order is authorized, this field can be left null | [optional] 
**RejectionCode** | Pointer to **int32** | The rejection code for one of the valid rejections to reject a port out order | [optional] 
**ServiceAddress** | Pointer to **string** | First line of billing address (street address) | [optional] 
**FocDate** | Pointer to **string** | ISO 8601 formatted Date/Time of the FOC date | [optional] 
**RequestedFocDate** | Pointer to **string** | ISO 8601 formatted Date/Time of the user requested FOC date | [optional] 
**Spid** | Pointer to **string** | New service provider spid | [optional] 
**SupportKey** | Pointer to **string** | A key to reference this port out request when contacting Telnyx customer support | [optional] 
**Status** | Pointer to **string** | Status of portout request | [optional] 
**AlreadyPorted** | Pointer to **bool** | Is true when the number is already ported | [optional] 
**UserId** | Pointer to **string** | Identifies the user (or organization) who requested the port out | [optional] 
**Vendor** | Pointer to **string** | Telnyx partner providing network coverage | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date of when the portout was created | [optional] 
**InsertedAt** | Pointer to **string** | ISO 8601 formatted date of when the portout was created | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date of when the portout was last updated | [optional] 

## Methods

### NewPortoutDetails

`func NewPortoutDetails() *PortoutDetails`

NewPortoutDetails instantiates a new PortoutDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortoutDetailsWithDefaults

`func NewPortoutDetailsWithDefaults() *PortoutDetails`

NewPortoutDetailsWithDefaults instantiates a new PortoutDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortoutDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortoutDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortoutDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortoutDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *PortoutDetails) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortoutDetails) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortoutDetails) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortoutDetails) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *PortoutDetails) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *PortoutDetails) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *PortoutDetails) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *PortoutDetails) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetAuthorizedName

`func (o *PortoutDetails) GetAuthorizedName() string`

GetAuthorizedName returns the AuthorizedName field if non-nil, zero value otherwise.

### GetAuthorizedNameOk

`func (o *PortoutDetails) GetAuthorizedNameOk() (*string, bool)`

GetAuthorizedNameOk returns a tuple with the AuthorizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedName

`func (o *PortoutDetails) SetAuthorizedName(v string)`

SetAuthorizedName sets AuthorizedName field to given value.

### HasAuthorizedName

`func (o *PortoutDetails) HasAuthorizedName() bool`

HasAuthorizedName returns a boolean if a field has been set.

### GetCarrierName

`func (o *PortoutDetails) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *PortoutDetails) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *PortoutDetails) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *PortoutDetails) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetCurrentCarrier

`func (o *PortoutDetails) GetCurrentCarrier() string`

GetCurrentCarrier returns the CurrentCarrier field if non-nil, zero value otherwise.

### GetCurrentCarrierOk

`func (o *PortoutDetails) GetCurrentCarrierOk() (*string, bool)`

GetCurrentCarrierOk returns a tuple with the CurrentCarrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCarrier

`func (o *PortoutDetails) SetCurrentCarrier(v string)`

SetCurrentCarrier sets CurrentCarrier field to given value.

### HasCurrentCarrier

`func (o *PortoutDetails) HasCurrentCarrier() bool`

HasCurrentCarrier returns a boolean if a field has been set.

### GetEndUserName

`func (o *PortoutDetails) GetEndUserName() string`

GetEndUserName returns the EndUserName field if non-nil, zero value otherwise.

### GetEndUserNameOk

`func (o *PortoutDetails) GetEndUserNameOk() (*string, bool)`

GetEndUserNameOk returns a tuple with the EndUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserName

`func (o *PortoutDetails) SetEndUserName(v string)`

SetEndUserName sets EndUserName field to given value.

### HasEndUserName

`func (o *PortoutDetails) HasEndUserName() bool`

HasEndUserName returns a boolean if a field has been set.

### GetCity

`func (o *PortoutDetails) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PortoutDetails) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PortoutDetails) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PortoutDetails) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *PortoutDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PortoutDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PortoutDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PortoutDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *PortoutDetails) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *PortoutDetails) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *PortoutDetails) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *PortoutDetails) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetLsr

`func (o *PortoutDetails) GetLsr() []string`

GetLsr returns the Lsr field if non-nil, zero value otherwise.

### GetLsrOk

`func (o *PortoutDetails) GetLsrOk() (*[]string, bool)`

GetLsrOk returns a tuple with the Lsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLsr

`func (o *PortoutDetails) SetLsr(v []string)`

SetLsr sets Lsr field to given value.

### HasLsr

`func (o *PortoutDetails) HasLsr() bool`

HasLsr returns a boolean if a field has been set.

### GetPon

`func (o *PortoutDetails) GetPon() string`

GetPon returns the Pon field if non-nil, zero value otherwise.

### GetPonOk

`func (o *PortoutDetails) GetPonOk() (*string, bool)`

GetPonOk returns a tuple with the Pon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPon

`func (o *PortoutDetails) SetPon(v string)`

SetPon sets Pon field to given value.

### HasPon

`func (o *PortoutDetails) HasPon() bool`

HasPon returns a boolean if a field has been set.

### GetReason

`func (o *PortoutDetails) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PortoutDetails) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PortoutDetails) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PortoutDetails) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRejectionCode

`func (o *PortoutDetails) GetRejectionCode() int32`

GetRejectionCode returns the RejectionCode field if non-nil, zero value otherwise.

### GetRejectionCodeOk

`func (o *PortoutDetails) GetRejectionCodeOk() (*int32, bool)`

GetRejectionCodeOk returns a tuple with the RejectionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionCode

`func (o *PortoutDetails) SetRejectionCode(v int32)`

SetRejectionCode sets RejectionCode field to given value.

### HasRejectionCode

`func (o *PortoutDetails) HasRejectionCode() bool`

HasRejectionCode returns a boolean if a field has been set.

### GetServiceAddress

`func (o *PortoutDetails) GetServiceAddress() string`

GetServiceAddress returns the ServiceAddress field if non-nil, zero value otherwise.

### GetServiceAddressOk

`func (o *PortoutDetails) GetServiceAddressOk() (*string, bool)`

GetServiceAddressOk returns a tuple with the ServiceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAddress

`func (o *PortoutDetails) SetServiceAddress(v string)`

SetServiceAddress sets ServiceAddress field to given value.

### HasServiceAddress

`func (o *PortoutDetails) HasServiceAddress() bool`

HasServiceAddress returns a boolean if a field has been set.

### GetFocDate

`func (o *PortoutDetails) GetFocDate() string`

GetFocDate returns the FocDate field if non-nil, zero value otherwise.

### GetFocDateOk

`func (o *PortoutDetails) GetFocDateOk() (*string, bool)`

GetFocDateOk returns a tuple with the FocDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocDate

`func (o *PortoutDetails) SetFocDate(v string)`

SetFocDate sets FocDate field to given value.

### HasFocDate

`func (o *PortoutDetails) HasFocDate() bool`

HasFocDate returns a boolean if a field has been set.

### GetRequestedFocDate

`func (o *PortoutDetails) GetRequestedFocDate() string`

GetRequestedFocDate returns the RequestedFocDate field if non-nil, zero value otherwise.

### GetRequestedFocDateOk

`func (o *PortoutDetails) GetRequestedFocDateOk() (*string, bool)`

GetRequestedFocDateOk returns a tuple with the RequestedFocDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFocDate

`func (o *PortoutDetails) SetRequestedFocDate(v string)`

SetRequestedFocDate sets RequestedFocDate field to given value.

### HasRequestedFocDate

`func (o *PortoutDetails) HasRequestedFocDate() bool`

HasRequestedFocDate returns a boolean if a field has been set.

### GetSpid

`func (o *PortoutDetails) GetSpid() string`

GetSpid returns the Spid field if non-nil, zero value otherwise.

### GetSpidOk

`func (o *PortoutDetails) GetSpidOk() (*string, bool)`

GetSpidOk returns a tuple with the Spid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpid

`func (o *PortoutDetails) SetSpid(v string)`

SetSpid sets Spid field to given value.

### HasSpid

`func (o *PortoutDetails) HasSpid() bool`

HasSpid returns a boolean if a field has been set.

### GetSupportKey

`func (o *PortoutDetails) GetSupportKey() string`

GetSupportKey returns the SupportKey field if non-nil, zero value otherwise.

### GetSupportKeyOk

`func (o *PortoutDetails) GetSupportKeyOk() (*string, bool)`

GetSupportKeyOk returns a tuple with the SupportKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportKey

`func (o *PortoutDetails) SetSupportKey(v string)`

SetSupportKey sets SupportKey field to given value.

### HasSupportKey

`func (o *PortoutDetails) HasSupportKey() bool`

HasSupportKey returns a boolean if a field has been set.

### GetStatus

`func (o *PortoutDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortoutDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortoutDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PortoutDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAlreadyPorted

`func (o *PortoutDetails) GetAlreadyPorted() bool`

GetAlreadyPorted returns the AlreadyPorted field if non-nil, zero value otherwise.

### GetAlreadyPortedOk

`func (o *PortoutDetails) GetAlreadyPortedOk() (*bool, bool)`

GetAlreadyPortedOk returns a tuple with the AlreadyPorted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyPorted

`func (o *PortoutDetails) SetAlreadyPorted(v bool)`

SetAlreadyPorted sets AlreadyPorted field to given value.

### HasAlreadyPorted

`func (o *PortoutDetails) HasAlreadyPorted() bool`

HasAlreadyPorted returns a boolean if a field has been set.

### GetUserId

`func (o *PortoutDetails) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PortoutDetails) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PortoutDetails) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PortoutDetails) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVendor

`func (o *PortoutDetails) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *PortoutDetails) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *PortoutDetails) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *PortoutDetails) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortoutDetails) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortoutDetails) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortoutDetails) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortoutDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetInsertedAt

`func (o *PortoutDetails) GetInsertedAt() string`

GetInsertedAt returns the InsertedAt field if non-nil, zero value otherwise.

### GetInsertedAtOk

`func (o *PortoutDetails) GetInsertedAtOk() (*string, bool)`

GetInsertedAtOk returns a tuple with the InsertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedAt

`func (o *PortoutDetails) SetInsertedAt(v string)`

SetInsertedAt sets InsertedAt field to given value.

### HasInsertedAt

`func (o *PortoutDetails) HasInsertedAt() bool`

HasInsertedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortoutDetails) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortoutDetails) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortoutDetails) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortoutDetails) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



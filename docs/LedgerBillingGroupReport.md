# LedgerBillingGroupReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**OrganizationId** | Pointer to **string** | Uniquely identifies the organization that owns the resource. | [optional] 
**Status** | Pointer to **string** | Status of the ledger billing group report | [optional] 
**ReportUrl** | Pointer to **NullableString** | External url of the ledger billing group report, if the status is complete | [optional] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewLedgerBillingGroupReport

`func NewLedgerBillingGroupReport() *LedgerBillingGroupReport`

NewLedgerBillingGroupReport instantiates a new LedgerBillingGroupReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerBillingGroupReportWithDefaults

`func NewLedgerBillingGroupReportWithDefaults() *LedgerBillingGroupReport`

NewLedgerBillingGroupReportWithDefaults instantiates a new LedgerBillingGroupReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *LedgerBillingGroupReport) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *LedgerBillingGroupReport) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *LedgerBillingGroupReport) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *LedgerBillingGroupReport) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *LedgerBillingGroupReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerBillingGroupReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerBillingGroupReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LedgerBillingGroupReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *LedgerBillingGroupReport) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *LedgerBillingGroupReport) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *LedgerBillingGroupReport) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *LedgerBillingGroupReport) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetStatus

`func (o *LedgerBillingGroupReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LedgerBillingGroupReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LedgerBillingGroupReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LedgerBillingGroupReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReportUrl

`func (o *LedgerBillingGroupReport) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *LedgerBillingGroupReport) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *LedgerBillingGroupReport) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *LedgerBillingGroupReport) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.

### SetReportUrlNil

`func (o *LedgerBillingGroupReport) SetReportUrlNil(b bool)`

 SetReportUrlNil sets the value for ReportUrl to be an explicit nil

### UnsetReportUrl
`func (o *LedgerBillingGroupReport) UnsetReportUrl()`

UnsetReportUrl ensures that no value is present for ReportUrl, not even an explicit nil
### GetCreatedAt

`func (o *LedgerBillingGroupReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LedgerBillingGroupReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LedgerBillingGroupReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LedgerBillingGroupReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LedgerBillingGroupReport) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LedgerBillingGroupReport) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LedgerBillingGroupReport) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LedgerBillingGroupReport) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



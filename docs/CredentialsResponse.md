# CredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CustomStorageConfiguration**](CustomStorageConfiguration.md) |  | 
**ConnectionId** | **string** | Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection resource. | 
**RecordType** | [**RecordType**](RecordType.md) |  | 

## Methods

### NewCredentialsResponse

`func NewCredentialsResponse(data CustomStorageConfiguration, connectionId string, recordType RecordType, ) *CredentialsResponse`

NewCredentialsResponse instantiates a new CredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsResponseWithDefaults

`func NewCredentialsResponseWithDefaults() *CredentialsResponse`

NewCredentialsResponseWithDefaults instantiates a new CredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CredentialsResponse) GetData() CustomStorageConfiguration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CredentialsResponse) GetDataOk() (*CustomStorageConfiguration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CredentialsResponse) SetData(v CustomStorageConfiguration)`

SetData sets Data field to given value.


### GetConnectionId

`func (o *CredentialsResponse) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CredentialsResponse) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CredentialsResponse) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetRecordType

`func (o *CredentialsResponse) GetRecordType() RecordType`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CredentialsResponse) GetRecordTypeOk() (*RecordType, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CredentialsResponse) SetRecordType(v RecordType)`

SetRecordType sets RecordType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



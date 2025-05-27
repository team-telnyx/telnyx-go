# \AdvancedOptInOptOutAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutorespConfig**](AdvancedOptInOptOutAPI.md#CreateAutorespConfig) | **Post** /messaging_profiles/{profile_id}/autoresp_configs | Create Auto-Reponse Setting
[**DeleteAutorespConfig**](AdvancedOptInOptOutAPI.md#DeleteAutorespConfig) | **Delete** /messaging_profiles/{profile_id}/autoresp_configs/{autoresp_cfg_id} | Delete Auto-Response Setting
[**GetAutorespConfig**](AdvancedOptInOptOutAPI.md#GetAutorespConfig) | **Get** /messaging_profiles/{profile_id}/autoresp_configs/{autoresp_cfg_id} | Get Auto-Response Setting
[**GetAutorespConfigs**](AdvancedOptInOptOutAPI.md#GetAutorespConfigs) | **Get** /messaging_profiles/{profile_id}/autoresp_configs | List Auto-Response Settings
[**ListOptOuts**](AdvancedOptInOptOutAPI.md#ListOptOuts) | **Get** /messaging_optouts | List opt-outs
[**UpdateAutoRespConfig**](AdvancedOptInOptOutAPI.md#UpdateAutoRespConfig) | **Put** /messaging_profiles/{profile_id}/autoresp_configs/{autoresp_cfg_id} | Update Auto-Response Setting



## CreateAutorespConfig

> AutorespConfigResponseSchema CreateAutorespConfig(ctx, profileId).AutoRespConfigCreateSchema(autoRespConfigCreateSchema).Execute()

Create Auto-Reponse Setting

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	profileId := "profileId_example" // string | 
	autoRespConfigCreateSchema := *openapiclient.NewAutoRespConfigCreateSchema("Op_example", []string{"Keywords_example"}, "US") // AutoRespConfigCreateSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedOptInOptOutAPI.CreateAutorespConfig(context.Background(), profileId).AutoRespConfigCreateSchema(autoRespConfigCreateSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedOptInOptOutAPI.CreateAutorespConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutorespConfig`: AutorespConfigResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `AdvancedOptInOptOutAPI.CreateAutorespConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutorespConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoRespConfigCreateSchema** | [**AutoRespConfigCreateSchema**](AutoRespConfigCreateSchema.md) |  | 

### Return type

[**AutorespConfigResponseSchema**](AutorespConfigResponseSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutorespConfig

> interface{} DeleteAutorespConfig(ctx, profileId, autorespCfgId).Execute()

Delete Auto-Response Setting

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	autorespCfgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedOptInOptOutAPI.DeleteAutorespConfig(context.Background(), profileId, autorespCfgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedOptInOptOutAPI.DeleteAutorespConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAutorespConfig`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdvancedOptInOptOutAPI.DeleteAutorespConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 
**autorespCfgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutorespConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutorespConfig

> AutorespConfigResponseSchema GetAutorespConfig(ctx, profileId, autorespCfgId).Execute()

Get Auto-Response Setting

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	autorespCfgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedOptInOptOutAPI.GetAutorespConfig(context.Background(), profileId, autorespCfgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedOptInOptOutAPI.GetAutorespConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutorespConfig`: AutorespConfigResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `AdvancedOptInOptOutAPI.GetAutorespConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 
**autorespCfgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutorespConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AutorespConfigResponseSchema**](AutorespConfigResponseSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutorespConfigs

> AutorespConfigsResponseSchema GetAutorespConfigs(ctx, profileId).CountryCode(countryCode).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).UpdatedAtGte(updatedAtGte).UpdatedAtLte(updatedAtLte).Execute()

List Auto-Response Settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	countryCode := "countryCode_example" // string |  (optional)
	createdAtGte := "createdAtGte_example" // string |  (optional)
	createdAtLte := "createdAtLte_example" // string |  (optional)
	updatedAtGte := "updatedAtGte_example" // string |  (optional)
	updatedAtLte := "updatedAtLte_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedOptInOptOutAPI.GetAutorespConfigs(context.Background(), profileId).CountryCode(countryCode).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).UpdatedAtGte(updatedAtGte).UpdatedAtLte(updatedAtLte).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedOptInOptOutAPI.GetAutorespConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutorespConfigs`: AutorespConfigsResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `AdvancedOptInOptOutAPI.GetAutorespConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutorespConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **countryCode** | **string** |  | 
 **createdAtGte** | **string** |  | 
 **createdAtLte** | **string** |  | 
 **updatedAtGte** | **string** |  | 
 **updatedAtLte** | **string** |  | 

### Return type

[**AutorespConfigsResponseSchema**](AutorespConfigsResponseSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOptOuts

> OptOutListResponse ListOptOuts(ctx).FilterMessagingProfileId(filterMessagingProfileId).CreatedAtGte(createdAtGte).FilterFrom(filterFrom).CreatedAtLte(createdAtLte).RedactionEnabled(redactionEnabled).PageNumber(pageNumber).PageSize(pageSize).Execute()

List opt-outs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	filterMessagingProfileId := "filterMessagingProfileId_example" // string | The ID of the messaging profile to retrieve opt-outs for (optional)
	createdAtGte := time.Now() // time.Time | Filter opt-outs created after this date (ISO-8601 format) (optional)
	filterFrom := "filterFrom_example" // string | The sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code) to retrieve opt-outs for (optional)
	createdAtLte := time.Now() // time.Time | Filter opt-outs created before this date (ISO-8601 format) (optional)
	redactionEnabled := "+447766****" // bool | If receiving address (+E.164 formatted phone number) should be redacted (optional)
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedOptInOptOutAPI.ListOptOuts(context.Background()).FilterMessagingProfileId(filterMessagingProfileId).CreatedAtGte(createdAtGte).FilterFrom(filterFrom).CreatedAtLte(createdAtLte).RedactionEnabled(redactionEnabled).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedOptInOptOutAPI.ListOptOuts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOptOuts`: OptOutListResponse
	fmt.Fprintf(os.Stdout, "Response from `AdvancedOptInOptOutAPI.ListOptOuts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOptOutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterMessagingProfileId** | **string** | The ID of the messaging profile to retrieve opt-outs for | 
 **createdAtGte** | **time.Time** | Filter opt-outs created after this date (ISO-8601 format) | 
 **filterFrom** | **string** | The sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code) to retrieve opt-outs for | 
 **createdAtLte** | **time.Time** | Filter opt-outs created before this date (ISO-8601 format) | 
 **redactionEnabled** | **bool** | If receiving address (+E.164 formatted phone number) should be redacted | 
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**OptOutListResponse**](OptOutListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoRespConfig

> AutorespConfigResponseSchema UpdateAutoRespConfig(ctx, profileId, autorespCfgId).AutoRespConfigCreateSchema(autoRespConfigCreateSchema).Execute()

Update Auto-Response Setting

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	autorespCfgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	autoRespConfigCreateSchema := *openapiclient.NewAutoRespConfigCreateSchema("Op_example", []string{"Keywords_example"}, "US") // AutoRespConfigCreateSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedOptInOptOutAPI.UpdateAutoRespConfig(context.Background(), profileId, autorespCfgId).AutoRespConfigCreateSchema(autoRespConfigCreateSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedOptInOptOutAPI.UpdateAutoRespConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoRespConfig`: AutorespConfigResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `AdvancedOptInOptOutAPI.UpdateAutoRespConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 
**autorespCfgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoRespConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoRespConfigCreateSchema** | [**AutoRespConfigCreateSchema**](AutoRespConfigCreateSchema.md) |  | 

### Return type

[**AutorespConfigResponseSchema**](AutorespConfigResponseSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


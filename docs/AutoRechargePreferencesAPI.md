# \AutoRechargePreferencesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutoRechargePrefs**](AutoRechargePreferencesAPI.md#GetAutoRechargePrefs) | **Get** /payment/auto_recharge_prefs | List auto recharge preferences
[**UpdateAutoRechargePrefs**](AutoRechargePreferencesAPI.md#UpdateAutoRechargePrefs) | **Patch** /payment/auto_recharge_prefs | Update auto recharge preferences



## GetAutoRechargePrefs

> GetAutoRechargePrefs200Response GetAutoRechargePrefs(ctx).Execute()

List auto recharge preferences



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoRechargePreferencesAPI.GetAutoRechargePrefs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoRechargePreferencesAPI.GetAutoRechargePrefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoRechargePrefs`: GetAutoRechargePrefs200Response
	fmt.Fprintf(os.Stdout, "Response from `AutoRechargePreferencesAPI.GetAutoRechargePrefs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoRechargePrefsRequest struct via the builder pattern


### Return type

[**GetAutoRechargePrefs200Response**](GetAutoRechargePrefs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoRechargePrefs

> GetAutoRechargePrefs200Response UpdateAutoRechargePrefs(ctx).AutoRechargePrefRequest(autoRechargePrefRequest).Execute()

Update auto recharge preferences



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
	autoRechargePrefRequest := *openapiclient.NewAutoRechargePrefRequest() // AutoRechargePrefRequest | Details to update auto recharge preferences

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoRechargePreferencesAPI.UpdateAutoRechargePrefs(context.Background()).AutoRechargePrefRequest(autoRechargePrefRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoRechargePreferencesAPI.UpdateAutoRechargePrefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoRechargePrefs`: GetAutoRechargePrefs200Response
	fmt.Fprintf(os.Stdout, "Response from `AutoRechargePreferencesAPI.UpdateAutoRechargePrefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoRechargePrefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoRechargePrefRequest** | [**AutoRechargePrefRequest**](AutoRechargePrefRequest.md) | Details to update auto recharge preferences | 

### Return type

[**GetAutoRechargePrefs200Response**](GetAutoRechargePrefs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


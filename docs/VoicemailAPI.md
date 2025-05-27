# \VoicemailAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVoicemail**](VoicemailAPI.md#CreateVoicemail) | **Post** /phone_numbers/{phone_number_id}/voicemail | Create voicemail
[**GetVoicemail**](VoicemailAPI.md#GetVoicemail) | **Get** /phone_numbers/{phone_number_id}/voicemail | Get voicemail
[**UpdateVoicemail**](VoicemailAPI.md#UpdateVoicemail) | **Patch** /phone_numbers/{phone_number_id}/voicemail | Update voicemail



## CreateVoicemail

> GetVoicemail200Response CreateVoicemail(ctx, phoneNumberId).VoicemailRequest(voicemailRequest).Execute()

Create voicemail



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
	phoneNumberId := "123455678900" // string | The ID of the phone number.
	voicemailRequest := *openapiclient.NewVoicemailRequest() // VoicemailRequest | Details to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.CreateVoicemail(context.Background(), phoneNumberId).VoicemailRequest(voicemailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.CreateVoicemail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVoicemail`: GetVoicemail200Response
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.CreateVoicemail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | **string** | The ID of the phone number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoicemailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **voicemailRequest** | [**VoicemailRequest**](VoicemailRequest.md) | Details to create | 

### Return type

[**GetVoicemail200Response**](GetVoicemail200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVoicemail

> GetVoicemail200Response GetVoicemail(ctx, phoneNumberId).Execute()

Get voicemail



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
	phoneNumberId := "123455678900" // string | The ID of the phone number.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.GetVoicemail(context.Background(), phoneNumberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.GetVoicemail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVoicemail`: GetVoicemail200Response
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.GetVoicemail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | **string** | The ID of the phone number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoicemailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetVoicemail200Response**](GetVoicemail200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVoicemail

> GetVoicemail200Response UpdateVoicemail(ctx, phoneNumberId).VoicemailRequest(voicemailRequest).Execute()

Update voicemail



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
	phoneNumberId := "123455678900" // string | The ID of the phone number.
	voicemailRequest := *openapiclient.NewVoicemailRequest() // VoicemailRequest | Details to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicemailAPI.UpdateVoicemail(context.Background(), phoneNumberId).VoicemailRequest(voicemailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicemailAPI.UpdateVoicemail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVoicemail`: GetVoicemail200Response
	fmt.Fprintf(os.Stdout, "Response from `VoicemailAPI.UpdateVoicemail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | **string** | The ID of the phone number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVoicemailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **voicemailRequest** | [**VoicemailRequest**](VoicemailRequest.md) | Details to update | 

### Return type

[**GetVoicemail200Response**](GetVoicemail200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


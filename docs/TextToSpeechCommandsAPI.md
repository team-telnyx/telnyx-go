# \TextToSpeechCommandsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateTextToSpeech**](TextToSpeechCommandsAPI.md#GenerateTextToSpeech) | **Post** /text-to-speech/speech | Generate speech from text
[**ListTextToSpeechVoices**](TextToSpeechCommandsAPI.md#ListTextToSpeechVoices) | **Get** /text-to-speech/voices | List available text to speech voices



## GenerateTextToSpeech

> map[string]interface{} GenerateTextToSpeech(ctx).GenerateTextToSpeechRequest(generateTextToSpeechRequest).Execute()

Generate speech from text



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
	generateTextToSpeechRequest := *openapiclient.NewGenerateTextToSpeechRequest("Voice_example", "Text_example") // GenerateTextToSpeechRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextToSpeechCommandsAPI.GenerateTextToSpeech(context.Background()).GenerateTextToSpeechRequest(generateTextToSpeechRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextToSpeechCommandsAPI.GenerateTextToSpeech``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateTextToSpeech`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TextToSpeechCommandsAPI.GenerateTextToSpeech`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTextToSpeechRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateTextToSpeechRequest** | [**GenerateTextToSpeechRequest**](GenerateTextToSpeechRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: audio/mpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTextToSpeechVoices

> ListTextToSpeechVoices200Response ListTextToSpeechVoices(ctx).Provider(provider).ElevenlabsApiKeyRef(elevenlabsApiKeyRef).Execute()

List available text to speech voices



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
	provider := "provider_example" // string | Filter voices by provider (optional)
	elevenlabsApiKeyRef := "elevenlabsApiKeyRef_example" // string | Reference to your ElevenLabs API key stored in the Telnyx Portal (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextToSpeechCommandsAPI.ListTextToSpeechVoices(context.Background()).Provider(provider).ElevenlabsApiKeyRef(elevenlabsApiKeyRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextToSpeechCommandsAPI.ListTextToSpeechVoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTextToSpeechVoices`: ListTextToSpeechVoices200Response
	fmt.Fprintf(os.Stdout, "Response from `TextToSpeechCommandsAPI.ListTextToSpeechVoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTextToSpeechVoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Filter voices by provider | 
 **elevenlabsApiKeyRef** | **string** | Reference to your ElevenLabs API key stored in the Telnyx Portal | 

### Return type

[**ListTextToSpeechVoices200Response**](ListTextToSpeechVoices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \AudioAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AudioPublicAudioTranscriptionsPost**](AudioAPI.md#AudioPublicAudioTranscriptionsPost) | **Post** /ai/audio/transcriptions | Transcribe speech to text



## AudioPublicAudioTranscriptionsPost

> AudioTranscriptionResponse AudioPublicAudioTranscriptionsPost(ctx).Model(model).File(file).FileUrl(fileUrl).ResponseFormat(responseFormat).TimestampGranularities(timestampGranularities).Execute()

Transcribe speech to text



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
	model := "model_example" // string | ID of the model to use. `distil-whisper/distil-large-v2` is lower latency but English-only. `openai/whisper-large-v3-turbo` is multi-lingual but slightly higher latency. (default to "distil-whisper/distil-large-v2")
	file := os.NewFile(1234, "some_file") // *os.File | The audio file object to transcribe, in one of these formats: flac, mp3, mp4, mpeg, mpga, m4a, ogg, wav, or webm. File uploads are limited to 100 MB. Cannot be used together with `file_url` (optional)
	fileUrl := "fileUrl_example" // string | Link to audio file in one of these formats: flac, mp3, mp4, mpeg, mpga, m4a, ogg, wav, or webm. Support for hosted files is limited to 100MB. Cannot be used together with `file` (optional)
	responseFormat := "responseFormat_example" // string | The format of the transcript output. Use `verbose_json` to take advantage of timestamps. (optional) (default to "json")
	timestampGranularities := "timestampGranularities_example" // string | The timestamp granularities to populate for this transcription. `response_format` must be set verbose_json to use timestamp granularities. Currently `segment` is supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.AudioPublicAudioTranscriptionsPost(context.Background()).Model(model).File(file).FileUrl(fileUrl).ResponseFormat(responseFormat).TimestampGranularities(timestampGranularities).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.AudioPublicAudioTranscriptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AudioPublicAudioTranscriptionsPost`: AudioTranscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.AudioPublicAudioTranscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAudioPublicAudioTranscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** | ID of the model to use. &#x60;distil-whisper/distil-large-v2&#x60; is lower latency but English-only. &#x60;openai/whisper-large-v3-turbo&#x60; is multi-lingual but slightly higher latency. | [default to &quot;distil-whisper/distil-large-v2&quot;]
 **file** | ***os.File** | The audio file object to transcribe, in one of these formats: flac, mp3, mp4, mpeg, mpga, m4a, ogg, wav, or webm. File uploads are limited to 100 MB. Cannot be used together with &#x60;file_url&#x60; | 
 **fileUrl** | **string** | Link to audio file in one of these formats: flac, mp3, mp4, mpeg, mpga, m4a, ogg, wav, or webm. Support for hosted files is limited to 100MB. Cannot be used together with &#x60;file&#x60; | 
 **responseFormat** | **string** | The format of the transcript output. Use &#x60;verbose_json&#x60; to take advantage of timestamps. | [default to &quot;json&quot;]
 **timestampGranularities** | **string** | The timestamp granularities to populate for this transcription. &#x60;response_format&#x60; must be set verbose_json to use timestamp granularities. Currently &#x60;segment&#x60; is supported. | 

### Return type

[**AudioTranscriptionResponse**](AudioTranscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


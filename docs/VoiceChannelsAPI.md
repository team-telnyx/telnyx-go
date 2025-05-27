# \VoiceChannelsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllNumbersChannelZones**](VoiceChannelsAPI.md#GetAllNumbersChannelZones) | **Get** /list | List All Numbers using Channel Billing
[**GetChannelZones**](VoiceChannelsAPI.md#GetChannelZones) | **Get** /channel_zones | List your voice channels for non-US zones
[**GetNumbersChannelZones**](VoiceChannelsAPI.md#GetNumbersChannelZones) | **Get** /list/{channel_zone_id} | List Numbers using Channel Billing for a specific Zone
[**ListInboundChannels**](VoiceChannelsAPI.md#ListInboundChannels) | **Get** /inbound_channels | List your voice channels for US Zone
[**PatchChannelZone**](VoiceChannelsAPI.md#PatchChannelZone) | **Put** /channel_zones/{channel_zone_id} | Update voice channels for non-US Zones
[**UpdateOutboundChannels**](VoiceChannelsAPI.md#UpdateOutboundChannels) | **Patch** /inbound_channels | Update voice channels for US Zone



## GetAllNumbersChannelZones

> GetAllNumbersChannelZones200Response GetAllNumbersChannelZones(ctx).Execute()

List All Numbers using Channel Billing



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
	resp, r, err := apiClient.VoiceChannelsAPI.GetAllNumbersChannelZones(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoiceChannelsAPI.GetAllNumbersChannelZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllNumbersChannelZones`: GetAllNumbersChannelZones200Response
	fmt.Fprintf(os.Stdout, "Response from `VoiceChannelsAPI.GetAllNumbersChannelZones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNumbersChannelZonesRequest struct via the builder pattern


### Return type

[**GetAllNumbersChannelZones200Response**](GetAllNumbersChannelZones200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelZones

> GetChannelZones200Response GetChannelZones(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

List your voice channels for non-US zones



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
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoiceChannelsAPI.GetChannelZones(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoiceChannelsAPI.GetChannelZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelZones`: GetChannelZones200Response
	fmt.Fprintf(os.Stdout, "Response from `VoiceChannelsAPI.GetChannelZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]

### Return type

[**GetChannelZones200Response**](GetChannelZones200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNumbersChannelZones

> GetAllNumbersChannelZones200Response GetNumbersChannelZones(ctx, channelZoneId).Execute()

List Numbers using Channel Billing for a specific Zone



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
	channelZoneId := "channelZoneId_example" // string | Channel zone identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoiceChannelsAPI.GetNumbersChannelZones(context.Background(), channelZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoiceChannelsAPI.GetNumbersChannelZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNumbersChannelZones`: GetAllNumbersChannelZones200Response
	fmt.Fprintf(os.Stdout, "Response from `VoiceChannelsAPI.GetNumbersChannelZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelZoneId** | **string** | Channel zone identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNumbersChannelZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAllNumbersChannelZones200Response**](GetAllNumbersChannelZones200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInboundChannels

> ListInboundChannels200Response ListInboundChannels(ctx).Execute()

List your voice channels for US Zone



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
	resp, r, err := apiClient.VoiceChannelsAPI.ListInboundChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoiceChannelsAPI.ListInboundChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInboundChannels`: ListInboundChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `VoiceChannelsAPI.ListInboundChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListInboundChannelsRequest struct via the builder pattern


### Return type

[**ListInboundChannels200Response**](ListInboundChannels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchChannelZone

> GcbChannelZone PatchChannelZone(ctx, channelZoneId).Body(body).Execute()

Update voice channels for non-US Zones



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
	channelZoneId := "channelZoneId_example" // string | Channel zone identifier
	body := *openapiclient.NewPatchChannelZoneRequest(int64(123)) // PatchChannelZoneRequest | Quantity of reserved channels and organizational update option

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoiceChannelsAPI.PatchChannelZone(context.Background(), channelZoneId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoiceChannelsAPI.PatchChannelZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchChannelZone`: GcbChannelZone
	fmt.Fprintf(os.Stdout, "Response from `VoiceChannelsAPI.PatchChannelZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelZoneId** | **string** | Channel zone identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchChannelZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PatchChannelZoneRequest**](PatchChannelZoneRequest.md) | Quantity of reserved channels and organizational update option | 

### Return type

[**GcbChannelZone**](GcbChannelZone.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOutboundChannels

> UpdateOutboundChannels200Response UpdateOutboundChannels(ctx).UpdateOutboundChannelsRequest(updateOutboundChannelsRequest).Execute()

Update voice channels for US Zone



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
	updateOutboundChannelsRequest := *openapiclient.NewUpdateOutboundChannelsRequest(int32(7)) // UpdateOutboundChannelsRequest | Voice channels update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoiceChannelsAPI.UpdateOutboundChannels(context.Background()).UpdateOutboundChannelsRequest(updateOutboundChannelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoiceChannelsAPI.UpdateOutboundChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOutboundChannels`: UpdateOutboundChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `VoiceChannelsAPI.UpdateOutboundChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOutboundChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOutboundChannelsRequest** | [**UpdateOutboundChannelsRequest**](UpdateOutboundChannelsRequest.md) | Voice channels update | 

### Return type

[**UpdateOutboundChannels200Response**](UpdateOutboundChannels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


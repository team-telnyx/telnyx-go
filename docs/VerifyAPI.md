# \VerifyAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlashcallVerification**](VerifyAPI.md#CreateFlashcallVerification) | **Post** /verifications/flashcall | Trigger Flash call verification
[**CreateVerificationCall**](VerifyAPI.md#CreateVerificationCall) | **Post** /verifications/call | Trigger Call verification
[**CreateVerificationSms**](VerifyAPI.md#CreateVerificationSms) | **Post** /verifications/sms | Trigger SMS verification
[**CreateVerifyProfile**](VerifyAPI.md#CreateVerifyProfile) | **Post** /verify_profiles | Create a Verify profile
[**DeleteProfile**](VerifyAPI.md#DeleteProfile) | **Delete** /verify_profiles/{verify_profile_id} | Delete Verify profile
[**GetVerifyProfile**](VerifyAPI.md#GetVerifyProfile) | **Get** /verify_profiles/{verify_profile_id} | Retrieve Verify profile
[**ListProfileMessageTemplates**](VerifyAPI.md#ListProfileMessageTemplates) | **Get** /verify_profiles/templates | Retrieve Verify profile message templates
[**ListProfiles**](VerifyAPI.md#ListProfiles) | **Get** /verify_profiles | List all Verify profiles
[**ListVerifications**](VerifyAPI.md#ListVerifications) | **Get** /verifications/by_phone_number/{phone_number} | List verifications by phone number
[**RetrieveVerification**](VerifyAPI.md#RetrieveVerification) | **Get** /verifications/{verification_id} | Retrieve verification
[**UpdateVerifyProfile**](VerifyAPI.md#UpdateVerifyProfile) | **Patch** /verify_profiles/{verify_profile_id} | Update Verify profile
[**VerifyVerificationCodeById**](VerifyAPI.md#VerifyVerificationCodeById) | **Post** /verifications/{verification_id}/actions/verify | Verify verification code by ID
[**VerifyVerificationCodeByPhoneNumber**](VerifyAPI.md#VerifyVerificationCodeByPhoneNumber) | **Post** /verifications/by_phone_number/{phone_number}/actions/verify | Verify verification code by phone number



## CreateFlashcallVerification

> CreateVerificationResponse CreateFlashcallVerification(ctx).CreateVerificationRequestFlashcall(createVerificationRequestFlashcall).Execute()

Trigger Flash call verification

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
	createVerificationRequestFlashcall := *openapiclient.NewCreateVerificationRequestFlashcall("+13035551234", "12ade33a-21c0-473b-b055-b3c836e1c292") // CreateVerificationRequestFlashcall | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.CreateFlashcallVerification(context.Background()).CreateVerificationRequestFlashcall(createVerificationRequestFlashcall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.CreateFlashcallVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFlashcallVerification`: CreateVerificationResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.CreateFlashcallVerification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlashcallVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVerificationRequestFlashcall** | [**CreateVerificationRequestFlashcall**](CreateVerificationRequestFlashcall.md) |  | 

### Return type

[**CreateVerificationResponse**](CreateVerificationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVerificationCall

> CreateVerificationResponse CreateVerificationCall(ctx).CreateVerificationRequestCall(createVerificationRequestCall).Execute()

Trigger Call verification

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
	createVerificationRequestCall := *openapiclient.NewCreateVerificationRequestCall("+13035551234", "12ade33a-21c0-473b-b055-b3c836e1c292") // CreateVerificationRequestCall | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.CreateVerificationCall(context.Background()).CreateVerificationRequestCall(createVerificationRequestCall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.CreateVerificationCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVerificationCall`: CreateVerificationResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.CreateVerificationCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerificationCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVerificationRequestCall** | [**CreateVerificationRequestCall**](CreateVerificationRequestCall.md) |  | 

### Return type

[**CreateVerificationResponse**](CreateVerificationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVerificationSms

> CreateVerificationResponse CreateVerificationSms(ctx).CreateVerificationRequestSMS(createVerificationRequestSMS).Execute()

Trigger SMS verification

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
	createVerificationRequestSMS := *openapiclient.NewCreateVerificationRequestSMS("+13035551234", "12ade33a-21c0-473b-b055-b3c836e1c292") // CreateVerificationRequestSMS | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.CreateVerificationSms(context.Background()).CreateVerificationRequestSMS(createVerificationRequestSMS).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.CreateVerificationSms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVerificationSms`: CreateVerificationResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.CreateVerificationSms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerificationSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVerificationRequestSMS** | [**CreateVerificationRequestSMS**](CreateVerificationRequestSMS.md) |  | 

### Return type

[**CreateVerificationResponse**](CreateVerificationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVerifyProfile

> VerifyProfileResponseDataWrapper CreateVerifyProfile(ctx).CreateVerifyProfileRequest(createVerifyProfileRequest).Execute()

Create a Verify profile



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
	createVerifyProfileRequest := *openapiclient.NewCreateVerifyProfileRequest("Test Profile") // CreateVerifyProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.CreateVerifyProfile(context.Background()).CreateVerifyProfileRequest(createVerifyProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.CreateVerifyProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVerifyProfile`: VerifyProfileResponseDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.CreateVerifyProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerifyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVerifyProfileRequest** | [**CreateVerifyProfileRequest**](CreateVerifyProfileRequest.md) |  | 

### Return type

[**VerifyProfileResponseDataWrapper**](VerifyProfileResponseDataWrapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProfile

> VerifyProfileResponseDataWrapper DeleteProfile(ctx, verifyProfileId).Execute()

Delete Verify profile

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
	verifyProfileId := "12ade33a-21c0-473b-b055-b3c836e1c292" // string | The identifier of the Verify profile to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.DeleteProfile(context.Background(), verifyProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.DeleteProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProfile`: VerifyProfileResponseDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.DeleteProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**verifyProfileId** | **string** | The identifier of the Verify profile to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VerifyProfileResponseDataWrapper**](VerifyProfileResponseDataWrapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVerifyProfile

> VerifyProfileResponseDataWrapper GetVerifyProfile(ctx, verifyProfileId).Execute()

Retrieve Verify profile



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
	verifyProfileId := "12ade33a-21c0-473b-b055-b3c836e1c292" // string | The identifier of the Verify profile to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.GetVerifyProfile(context.Background(), verifyProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.GetVerifyProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVerifyProfile`: VerifyProfileResponseDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.GetVerifyProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**verifyProfileId** | **string** | The identifier of the Verify profile to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVerifyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VerifyProfileResponseDataWrapper**](VerifyProfileResponseDataWrapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfileMessageTemplates

> ListVerifyProfileMessageTemplateResponse ListProfileMessageTemplates(ctx).Execute()

Retrieve Verify profile message templates



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
	resp, r, err := apiClient.VerifyAPI.ListProfileMessageTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.ListProfileMessageTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProfileMessageTemplates`: ListVerifyProfileMessageTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.ListProfileMessageTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListProfileMessageTemplatesRequest struct via the builder pattern


### Return type

[**ListVerifyProfileMessageTemplateResponse**](ListVerifyProfileMessageTemplateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfiles

> ListVerifyProfilesResponse ListProfiles(ctx).FilterName(filterName).PageSize(pageSize).PageNumber(pageNumber).Execute()

List all Verify profiles



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
	filterName := "filterName_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 25)
	pageNumber := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.ListProfiles(context.Background()).FilterName(filterName).PageSize(pageSize).PageNumber(pageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.ListProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProfiles`: ListVerifyProfilesResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.ListProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **string** |  | 
 **pageSize** | **int32** |  | [default to 25]
 **pageNumber** | **int32** |  | [default to 1]

### Return type

[**ListVerifyProfilesResponse**](ListVerifyProfilesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVerifications

> ListVerificationsResponse ListVerifications(ctx, phoneNumber).Execute()

List verifications by phone number

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
	phoneNumber := "+13035551234" // string | The phone number associated with the verifications to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.ListVerifications(context.Background(), phoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.ListVerifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVerifications`: ListVerificationsResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.ListVerifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** | The phone number associated with the verifications to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListVerificationsResponse**](ListVerificationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVerification

> RetrieveVerificationResponse RetrieveVerification(ctx, verificationId).Execute()

Retrieve verification

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
	verificationId := "12ade33a-21c0-473b-b055-b3c836e1c292" // string | The identifier of the verification to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.RetrieveVerification(context.Background(), verificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.RetrieveVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVerification`: RetrieveVerificationResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.RetrieveVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**verificationId** | **string** | The identifier of the verification to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetrieveVerificationResponse**](RetrieveVerificationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVerifyProfile

> VerifyProfileResponseDataWrapper UpdateVerifyProfile(ctx, verifyProfileId).UpdateVerifyProfileRequest(updateVerifyProfileRequest).Execute()

Update Verify profile

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
	verifyProfileId := "12ade33a-21c0-473b-b055-b3c836e1c292" // string | The identifier of the Verify profile to update.
	updateVerifyProfileRequest := *openapiclient.NewUpdateVerifyProfileRequest() // UpdateVerifyProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.UpdateVerifyProfile(context.Background(), verifyProfileId).UpdateVerifyProfileRequest(updateVerifyProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.UpdateVerifyProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVerifyProfile`: VerifyProfileResponseDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.UpdateVerifyProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**verifyProfileId** | **string** | The identifier of the Verify profile to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVerifyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVerifyProfileRequest** | [**UpdateVerifyProfileRequest**](UpdateVerifyProfileRequest.md) |  | 

### Return type

[**VerifyProfileResponseDataWrapper**](VerifyProfileResponseDataWrapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyVerificationCodeById

> VerifyVerificationCodeResponse VerifyVerificationCodeById(ctx, verificationId).VerifyVerificationCodeRequestById(verifyVerificationCodeRequestById).Execute()

Verify verification code by ID

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
	verificationId := "12ade33a-21c0-473b-b055-b3c836e1c292" // string | The identifier of the verification to retrieve.
	verifyVerificationCodeRequestById := *openapiclient.NewVerifyVerificationCodeRequestById() // VerifyVerificationCodeRequestById | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.VerifyVerificationCodeById(context.Background(), verificationId).VerifyVerificationCodeRequestById(verifyVerificationCodeRequestById).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.VerifyVerificationCodeById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyVerificationCodeById`: VerifyVerificationCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.VerifyVerificationCodeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**verificationId** | **string** | The identifier of the verification to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyVerificationCodeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyVerificationCodeRequestById** | [**VerifyVerificationCodeRequestById**](VerifyVerificationCodeRequestById.md) |  | 

### Return type

[**VerifyVerificationCodeResponse**](VerifyVerificationCodeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyVerificationCodeByPhoneNumber

> VerifyVerificationCodeResponse VerifyVerificationCodeByPhoneNumber(ctx, phoneNumber).VerifyVerificationCodeRequestByPhoneNumber(verifyVerificationCodeRequestByPhoneNumber).Execute()

Verify verification code by phone number

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
	phoneNumber := "+13035551234" // string | The phone number associated with the verification code being verified.
	verifyVerificationCodeRequestByPhoneNumber := *openapiclient.NewVerifyVerificationCodeRequestByPhoneNumber("17686", "12ade33a-21c0-473b-b055-b3c836e1c292") // VerifyVerificationCodeRequestByPhoneNumber | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.VerifyVerificationCodeByPhoneNumber(context.Background(), phoneNumber).VerifyVerificationCodeRequestByPhoneNumber(verifyVerificationCodeRequestByPhoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.VerifyVerificationCodeByPhoneNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyVerificationCodeByPhoneNumber`: VerifyVerificationCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.VerifyVerificationCodeByPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** | The phone number associated with the verification code being verified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyVerificationCodeByPhoneNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyVerificationCodeRequestByPhoneNumber** | [**VerifyVerificationCodeRequestByPhoneNumber**](VerifyVerificationCodeRequestByPhoneNumber.md) |  | 

### Return type

[**VerifyVerificationCodeResponse**](VerifyVerificationCodeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


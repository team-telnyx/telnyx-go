# \DataMigrationAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMigration**](DataMigrationAPI.md#CreateMigration) | **Post** /storage/migrations | Create a Migration
[**CreateMigrationSource**](DataMigrationAPI.md#CreateMigrationSource) | **Post** /storage/migration_sources | Create a Migration Source
[**DeleteMigrationSource**](DataMigrationAPI.md#DeleteMigrationSource) | **Delete** /storage/migration_sources/{id} | Delete a Migration Source
[**GetMigration**](DataMigrationAPI.md#GetMigration) | **Get** /storage/migrations/{id} | Get a Migration
[**GetMigrationSource**](DataMigrationAPI.md#GetMigrationSource) | **Get** /storage/migration_sources/{id} | Get a Migration Source
[**ListMigrationSourceCoverage**](DataMigrationAPI.md#ListMigrationSourceCoverage) | **Get** /storage/migration_source_coverage | List Migration Source coverage
[**ListMigrationSources**](DataMigrationAPI.md#ListMigrationSources) | **Get** /storage/migration_sources | List all Migration Sources
[**ListMigrations**](DataMigrationAPI.md#ListMigrations) | **Get** /storage/migrations | List all Migrations
[**StopMigration**](DataMigrationAPI.md#StopMigration) | **Post** /storage/migrations/{id}/actions/stop | Stop a Migration



## CreateMigration

> CreateMigration200Response CreateMigration(ctx).MigrationParams(migrationParams).Execute()

Create a Migration



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
	migrationParams := *openapiclient.NewMigrationParams("SourceId_example", "TargetBucketName_example", "TargetRegion_example") // MigrationParams |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataMigrationAPI.CreateMigration(context.Background()).MigrationParams(migrationParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataMigrationAPI.CreateMigration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMigration`: CreateMigration200Response
	fmt.Fprintf(os.Stdout, "Response from `DataMigrationAPI.CreateMigration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationParams** | [**MigrationParams**](MigrationParams.md) |  | 

### Return type

[**CreateMigration200Response**](CreateMigration200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMigrationSource

> CreateMigrationSource200Response CreateMigrationSource(ctx).MigrationSourceParams(migrationSourceParams).Execute()

Create a Migration Source



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
	migrationSourceParams := *openapiclient.NewMigrationSourceParams("Provider_example", *openapiclient.NewMigrationSourceParamsProviderAuth(), "BucketName_example") // MigrationSourceParams |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataMigrationAPI.CreateMigrationSource(context.Background()).MigrationSourceParams(migrationSourceParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataMigrationAPI.CreateMigrationSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMigrationSource`: CreateMigrationSource200Response
	fmt.Fprintf(os.Stdout, "Response from `DataMigrationAPI.CreateMigrationSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMigrationSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationSourceParams** | [**MigrationSourceParams**](MigrationSourceParams.md) |  | 

### Return type

[**CreateMigrationSource200Response**](CreateMigrationSource200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMigrationSource

> CreateMigrationSource200Response DeleteMigrationSource(ctx, id).Execute()

Delete a Migration Source



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
	id := "id_example" // string | Unique identifier for the data migration source.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataMigrationAPI.DeleteMigrationSource(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataMigrationAPI.DeleteMigrationSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMigrationSource`: CreateMigrationSource200Response
	fmt.Fprintf(os.Stdout, "Response from `DataMigrationAPI.DeleteMigrationSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the data migration source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMigrationSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateMigrationSource200Response**](CreateMigrationSource200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMigration

> CreateMigration200Response GetMigration(ctx, id).Execute()

Get a Migration



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
	id := "id_example" // string | Unique identifier for the data migration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataMigrationAPI.GetMigration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataMigrationAPI.GetMigration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMigration`: CreateMigration200Response
	fmt.Fprintf(os.Stdout, "Response from `DataMigrationAPI.GetMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the data migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateMigration200Response**](CreateMigration200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMigrationSource

> CreateMigrationSource200Response GetMigrationSource(ctx, id).Execute()

Get a Migration Source



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
	id := "id_example" // string | Unique identifier for the data migration source.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataMigrationAPI.GetMigrationSource(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataMigrationAPI.GetMigrationSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMigrationSource`: CreateMigrationSource200Response
	fmt.Fprintf(os.Stdout, "Response from `DataMigrationAPI.GetMigrationSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the data migration source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMigrationSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateMigrationSource200Response**](CreateMigrationSource200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMigrationSourceCoverage

> ListMigrationSourceCoverage200Response ListMigrationSourceCoverage(ctx).Execute()

List Migration Source coverage



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
	resp, r, err := apiClient.DataMigrationAPI.ListMigrationSourceCoverage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataMigrationAPI.ListMigrationSourceCoverage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMigrationSourceCoverage`: ListMigrationSourceCoverage200Response
	fmt.Fprintf(os.Stdout, "Response from `DataMigrationAPI.ListMigrationSourceCoverage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMigrationSourceCoverageRequest struct via the builder pattern


### Return type

[**ListMigrationSourceCoverage200Response**](ListMigrationSourceCoverage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMigrationSources

> ListMigrationSources200Response ListMigrationSources(ctx).Execute()

List all Migration Sources



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
	resp, r, err := apiClient.DataMigrationAPI.ListMigrationSources(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataMigrationAPI.ListMigrationSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMigrationSources`: ListMigrationSources200Response
	fmt.Fprintf(os.Stdout, "Response from `DataMigrationAPI.ListMigrationSources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMigrationSourcesRequest struct via the builder pattern


### Return type

[**ListMigrationSources200Response**](ListMigrationSources200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMigrations

> ListMigrations200Response ListMigrations(ctx).Execute()

List all Migrations



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
	resp, r, err := apiClient.DataMigrationAPI.ListMigrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataMigrationAPI.ListMigrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMigrations`: ListMigrations200Response
	fmt.Fprintf(os.Stdout, "Response from `DataMigrationAPI.ListMigrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMigrationsRequest struct via the builder pattern


### Return type

[**ListMigrations200Response**](ListMigrations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopMigration

> CreateMigration200Response StopMigration(ctx, id).Execute()

Stop a Migration



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
	id := "id_example" // string | Unique identifier for the data migration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataMigrationAPI.StopMigration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataMigrationAPI.StopMigration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopMigration`: CreateMigration200Response
	fmt.Fprintf(os.Stdout, "Response from `DataMigrationAPI.StopMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the data migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateMigration200Response**](CreateMigration200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \FastlyIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFastlyIntegrationService**](FastlyIntegrationApi.md#CreateFastlyIntegrationService) | **Post** /api/v1/integration/fastly/configuration/services | Create a new service object
[**DeleteFastlyIntegrationService**](FastlyIntegrationApi.md#DeleteFastlyIntegrationService) | **Delete** /api/v1/integration/fastly/configuration/services/{service_id} | Delete a single service object
[**GetFastlyIntegrationService**](FastlyIntegrationApi.md#GetFastlyIntegrationService) | **Get** /api/v1/integration/fastly/configuration/services/{service_id} | Get a single service object
[**UpdateFastlyIntegrationService**](FastlyIntegrationApi.md#UpdateFastlyIntegrationService) | **Put** /api/v1/integration/fastly/configuration/services/{service_id} | Update a single service object



## CreateFastlyIntegrationService

> FastlyServiceId CreateFastlyIntegrationService(ctx).Body(body).Execute()

Create a new service object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewFastlyService("ServiceId_example", "ServiceTags_example", "ServiceAccount_example") // FastlyService | Create a new service object request body.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FastlyIntegrationApi.CreateFastlyIntegrationService(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FastlyIntegrationApi.CreateFastlyIntegrationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFastlyIntegrationService`: FastlyServiceId
    fmt.Fprintf(os.Stdout, "Response from `FastlyIntegrationApi.CreateFastlyIntegrationService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFastlyIntegrationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FastlyService**](FastlyService.md) | Create a new service object request body. | 

### Return type

[**FastlyServiceId**](FastlyServiceId.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFastlyIntegrationService

> DeleteFastlyIntegrationService(ctx, serviceId).Execute()

Delete a single service object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | The service id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FastlyIntegrationApi.DeleteFastlyIntegrationService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FastlyIntegrationApi.DeleteFastlyIntegrationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFastlyIntegrationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFastlyIntegrationService

> FastlyServiceId GetFastlyIntegrationService(ctx, serviceId).Execute()

Get a single service object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | The service id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FastlyIntegrationApi.GetFastlyIntegrationService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FastlyIntegrationApi.GetFastlyIntegrationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFastlyIntegrationService`: FastlyServiceId
    fmt.Fprintf(os.Stdout, "Response from `FastlyIntegrationApi.GetFastlyIntegrationService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFastlyIntegrationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FastlyServiceId**](FastlyServiceId.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFastlyIntegrationService

> UpdateFastlyIntegrationService(ctx, serviceId, serviceAccount, serviceTags).Execute()

Update a single service object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | The service id
    serviceAccount := "serviceAccount_example" // string | The service account
    serviceTags := "serviceTags_example" // string | The service tags

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FastlyIntegrationApi.UpdateFastlyIntegrationService(context.Background(), serviceId, serviceAccount, serviceTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FastlyIntegrationApi.UpdateFastlyIntegrationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The service id | 
**serviceAccount** | **string** | The service account | 
**serviceTags** | **string** | The service tags | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFastlyIntegrationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


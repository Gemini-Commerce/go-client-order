# \OrderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderApproveOrder**](OrderApi.md#OrderApproveOrder) | **Post** /order.Order/ApproveOrder | 
[**OrderAssignShipment**](OrderApi.md#OrderAssignShipment) | **Post** /order.Order/AssignShipment | 
[**OrderCalculateRefund**](OrderApi.md#OrderCalculateRefund) | **Post** /order.Order/CalculateRefund | Refund custom methods
[**OrderCancelFulfillment**](OrderApi.md#OrderCancelFulfillment) | **Post** /order.Order/CancelFulfillment | 
[**OrderCancelOrder**](OrderApi.md#OrderCancelOrder) | **Post** /order.Order/CancelOrder | 
[**OrderCancelShipment**](OrderApi.md#OrderCancelShipment) | **Post** /order.Order/CancelShipment | 
[**OrderCompleteShipmentPacking**](OrderApi.md#OrderCompleteShipmentPacking) | **Post** /order.Order/CompleteShipmentPacking | 
[**OrderCreateFulfillment**](OrderApi.md#OrderCreateFulfillment) | **Post** /order.Order/CreateFulfillment | 
[**OrderCreateHistory**](OrderApi.md#OrderCreateHistory) | **Post** /order.Order/CreateHistory | 
[**OrderCreateOrder**](OrderApi.md#OrderCreateOrder) | **Post** /order.Order/CreateOrder | 
[**OrderCreatePayment**](OrderApi.md#OrderCreatePayment) | **Post** /order.Order/CreatePayment | 
[**OrderCreatePaymentTransaction**](OrderApi.md#OrderCreatePaymentTransaction) | **Post** /order.Order/CreatePaymentTransaction | 
[**OrderCreateRefund**](OrderApi.md#OrderCreateRefund) | **Post** /order.Order/CreateRefund | 
[**OrderCreateRefundTransaction**](OrderApi.md#OrderCreateRefundTransaction) | **Post** /order.Order/CreateRefundTransaction | 
[**OrderCreateShipment**](OrderApi.md#OrderCreateShipment) | **Post** /order.Order/CreateShipment | 
[**OrderDeleteOrder**](OrderApi.md#OrderDeleteOrder) | **Post** /order.Order/DeleteOrder | 
[**OrderGetFulfillment**](OrderApi.md#OrderGetFulfillment) | **Post** /order.Order/GetFulfillment | 
[**OrderGetOrder**](OrderApi.md#OrderGetOrder) | **Post** /order.Order/GetOrder | 
[**OrderGetOrderByCartId**](OrderApi.md#OrderGetOrderByCartId) | **Post** /order.Order/GetOrderByCartId | cart method
[**OrderGetOrderByOrderNumber**](OrderApi.md#OrderGetOrderByOrderNumber) | **Post** /order.Order/GetOrderByOrderNumber | 
[**OrderGetPayment**](OrderApi.md#OrderGetPayment) | **Post** /order.Order/GetPayment | 
[**OrderGetShipment**](OrderApi.md#OrderGetShipment) | **Post** /order.Order/GetShipment | 
[**OrderGetTransaction**](OrderApi.md#OrderGetTransaction) | **Post** /order.Order/GetTransaction | 
[**OrderHoldOrder**](OrderApi.md#OrderHoldOrder) | **Post** /order.Order/HoldOrder | Order custom methods
[**OrderImportOrder**](OrderApi.md#OrderImportOrder) | **Post** /order.Order/ImportOrder | Import orders API
[**OrderListFulfillments**](OrderApi.md#OrderListFulfillments) | **Post** /order.Order/ListFulfillments | 
[**OrderListOrders**](OrderApi.md#OrderListOrders) | **Post** /order.Order/ListOrders | 
[**OrderListOrdersByCustomer**](OrderApi.md#OrderListOrdersByCustomer) | **Post** /order.Order/ListOrdersByCustomer | 
[**OrderListOrdersByNumbers**](OrderApi.md#OrderListOrdersByNumbers) | **Post** /order.Order/ListOrdersByNumbers | 
[**OrderListShipments**](OrderApi.md#OrderListShipments) | **Post** /order.Order/ListShipments | 
[**OrderPrintOrdersLabels**](OrderApi.md#OrderPrintOrdersLabels) | **Post** /order.Order/PrintOrdersLabels | Labels custom methods
[**OrderQuashFulfillment**](OrderApi.md#OrderQuashFulfillment) | **Post** /order.Order/QuashFulfillment | 
[**OrderQuashShipment**](OrderApi.md#OrderQuashShipment) | **Post** /order.Order/QuashShipment | 
[**OrderReceiveFulfillment**](OrderApi.md#OrderReceiveFulfillment) | **Post** /order.Order/ReceiveFulfillment | 
[**OrderReportFulfillmentError**](OrderApi.md#OrderReportFulfillmentError) | **Post** /order.Order/ReportFulfillmentError | 
[**OrderReportFulfillmentNotResolvable**](OrderApi.md#OrderReportFulfillmentNotResolvable) | **Post** /order.Order/ReportFulfillmentNotResolvable | 
[**OrderReportFulfillmentReady**](OrderApi.md#OrderReportFulfillmentReady) | **Post** /order.Order/ReportFulfillmentReady | 
[**OrderReportShipmentDelivery**](OrderApi.md#OrderReportShipmentDelivery) | **Post** /order.Order/ReportShipmentDelivery | 
[**OrderReportShipmentMissingStock**](OrderApi.md#OrderReportShipmentMissingStock) | **Post** /order.Order/ReportShipmentMissingStock | 
[**OrderResolveShipmentMissingStock**](OrderApi.md#OrderResolveShipmentMissingStock) | **Post** /order.Order/ResolveShipmentMissingStock | 
[**OrderRetryFulfillment**](OrderApi.md#OrderRetryFulfillment) | **Post** /order.Order/RetryFulfillment | 
[**OrderSearchOrders**](OrderApi.md#OrderSearchOrders) | **Post** /order.Order/SearchOrders | 
[**OrderSendFulfillment**](OrderApi.md#OrderSendFulfillment) | **Post** /order.Order/SendFulfillment | 
[**OrderSendOrderNotification**](OrderApi.md#OrderSendOrderNotification) | **Post** /order.Order/SendOrderNotification | 
[**OrderStartFulfillmentProcessing**](OrderApi.md#OrderStartFulfillmentProcessing) | **Post** /order.Order/StartFulfillmentProcessing | Fulfillment custom methods
[**OrderStartShipmentProcessing**](OrderApi.md#OrderStartShipmentProcessing) | **Post** /order.Order/StartShipmentProcessing | Shipment custom methods
[**OrderUnholdOrder**](OrderApi.md#OrderUnholdOrder) | **Post** /order.Order/UnholdOrder | 
[**OrderUpdateOrder**](OrderApi.md#OrderUpdateOrder) | **Post** /order.Order/UpdateOrder | 
[**OrderUpdatePayment**](OrderApi.md#OrderUpdatePayment) | **Post** /order.Order/UpdatePayment | 



## OrderApproveOrder

> map[string]interface{} OrderApproveOrder(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderApproveOrderRequest() // OrderApproveOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderApproveOrder(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderApproveOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderApproveOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderApproveOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderApproveOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderApproveOrderRequest**](OrderApproveOrderRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderAssignShipment

> map[string]interface{} OrderAssignShipment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderAssignShipmentRequest() // OrderAssignShipmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderAssignShipment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderAssignShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderAssignShipment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderAssignShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderAssignShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderAssignShipmentRequest**](OrderAssignShipmentRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCalculateRefund

> OrderCalculateRefundResponse OrderCalculateRefund(ctx).Body(body).Execute()

Refund custom methods

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
    body := *openapiclient.NewOrderCalculateRefundRequest() // OrderCalculateRefundRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCalculateRefund(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCalculateRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCalculateRefund`: OrderCalculateRefundResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCalculateRefund`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCalculateRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCalculateRefundRequest**](OrderCalculateRefundRequest.md) |  | 

### Return type

[**OrderCalculateRefundResponse**](OrderCalculateRefundResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCancelFulfillment

> map[string]interface{} OrderCancelFulfillment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderCancelFulfillmentRequest() // OrderCancelFulfillmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCancelFulfillment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCancelFulfillment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCancelFulfillment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCancelFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCancelFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCancelFulfillmentRequest**](OrderCancelFulfillmentRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCancelOrder

> map[string]interface{} OrderCancelOrder(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderCancelOrderRequest() // OrderCancelOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCancelOrder(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCancelOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCancelOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCancelOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCancelOrderRequest**](OrderCancelOrderRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCancelShipment

> map[string]interface{} OrderCancelShipment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderCancelShipmentRequest() // OrderCancelShipmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCancelShipment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCancelShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCancelShipment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCancelShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCancelShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCancelShipmentRequest**](OrderCancelShipmentRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCompleteShipmentPacking

> map[string]interface{} OrderCompleteShipmentPacking(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderCompleteShipmentPackingRequest() // OrderCompleteShipmentPackingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCompleteShipmentPacking(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCompleteShipmentPacking``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCompleteShipmentPacking`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCompleteShipmentPacking`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCompleteShipmentPackingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCompleteShipmentPackingRequest**](OrderCompleteShipmentPackingRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCreateFulfillment

> OrderFulfillment OrderCreateFulfillment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderCreateFulfillmentRequest() // OrderCreateFulfillmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCreateFulfillment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCreateFulfillment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCreateFulfillment`: OrderFulfillment
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCreateFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCreateFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCreateFulfillmentRequest**](OrderCreateFulfillmentRequest.md) |  | 

### Return type

[**OrderFulfillment**](OrderFulfillment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCreateHistory

> OrderDataHistory OrderCreateHistory(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderCreateHistoryRequest() // OrderCreateHistoryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCreateHistory(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCreateHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCreateHistory`: OrderDataHistory
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCreateHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCreateHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCreateHistoryRequest**](OrderCreateHistoryRequest.md) |  | 

### Return type

[**OrderDataHistory**](OrderDataHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCreateOrder

> OrderOrderData OrderCreateOrder(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderCreateOrderRequest() // OrderCreateOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCreateOrder(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCreateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCreateOrder`: OrderOrderData
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCreateOrderRequest**](OrderCreateOrderRequest.md) |  | 

### Return type

[**OrderOrderData**](OrderOrderData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCreatePayment

> OrderPayment OrderCreatePayment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderCreatePaymentRequest() // OrderCreatePaymentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCreatePayment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCreatePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCreatePayment`: OrderPayment
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCreatePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCreatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCreatePaymentRequest**](OrderCreatePaymentRequest.md) |  | 

### Return type

[**OrderPayment**](OrderPayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCreatePaymentTransaction

> OrderTransaction OrderCreatePaymentTransaction(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderCreatePaymentTransactionRequest() // OrderCreatePaymentTransactionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCreatePaymentTransaction(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCreatePaymentTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCreatePaymentTransaction`: OrderTransaction
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCreatePaymentTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCreatePaymentTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCreatePaymentTransactionRequest**](OrderCreatePaymentTransactionRequest.md) |  | 

### Return type

[**OrderTransaction**](OrderTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCreateRefund

> OrderRefund OrderCreateRefund(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderCreateRefundRequest() // OrderCreateRefundRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCreateRefund(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCreateRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCreateRefund`: OrderRefund
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCreateRefund`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCreateRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCreateRefundRequest**](OrderCreateRefundRequest.md) |  | 

### Return type

[**OrderRefund**](OrderRefund.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCreateRefundTransaction

> OrderTransaction OrderCreateRefundTransaction(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderCreateRefundTransactionRequest() // OrderCreateRefundTransactionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCreateRefundTransaction(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCreateRefundTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCreateRefundTransaction`: OrderTransaction
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCreateRefundTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCreateRefundTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCreateRefundTransactionRequest**](OrderCreateRefundTransactionRequest.md) |  | 

### Return type

[**OrderTransaction**](OrderTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCreateShipment

> OrderShipment OrderCreateShipment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderCreateShipmentRequest() // OrderCreateShipmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderCreateShipment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderCreateShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderCreateShipment`: OrderShipment
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderCreateShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCreateShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderCreateShipmentRequest**](OrderCreateShipmentRequest.md) |  | 

### Return type

[**OrderShipment**](OrderShipment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderDeleteOrder

> map[string]interface{} OrderDeleteOrder(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderDeleteOrderRequest() // OrderDeleteOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderDeleteOrder(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderDeleteOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderDeleteOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderDeleteOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderDeleteOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderDeleteOrderRequest**](OrderDeleteOrderRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGetFulfillment

> OrderFulfillment OrderGetFulfillment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderGetFulfillmentRequest() // OrderGetFulfillmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderGetFulfillment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderGetFulfillment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderGetFulfillment`: OrderFulfillment
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderGetFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGetFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderGetFulfillmentRequest**](OrderGetFulfillmentRequest.md) |  | 

### Return type

[**OrderFulfillment**](OrderFulfillment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGetOrder

> OrderOrderData OrderGetOrder(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderGetOrderRequest() // OrderGetOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderGetOrder(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderGetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderGetOrder`: OrderOrderData
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderGetOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderGetOrderRequest**](OrderGetOrderRequest.md) |  | 

### Return type

[**OrderOrderData**](OrderOrderData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGetOrderByCartId

> OrderOrderData OrderGetOrderByCartId(ctx).Body(body).Execute()

cart method

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
    body := *openapiclient.NewOrderGetOrderByCartIdRequest() // OrderGetOrderByCartIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderGetOrderByCartId(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderGetOrderByCartId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderGetOrderByCartId`: OrderOrderData
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderGetOrderByCartId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGetOrderByCartIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderGetOrderByCartIdRequest**](OrderGetOrderByCartIdRequest.md) |  | 

### Return type

[**OrderOrderData**](OrderOrderData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGetOrderByOrderNumber

> OrderOrderData OrderGetOrderByOrderNumber(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderGetOrderByOrderNumberRequest() // OrderGetOrderByOrderNumberRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderGetOrderByOrderNumber(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderGetOrderByOrderNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderGetOrderByOrderNumber`: OrderOrderData
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderGetOrderByOrderNumber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGetOrderByOrderNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderGetOrderByOrderNumberRequest**](OrderGetOrderByOrderNumberRequest.md) |  | 

### Return type

[**OrderOrderData**](OrderOrderData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGetPayment

> OrderPayment OrderGetPayment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderGetPaymentRequest() // OrderGetPaymentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderGetPayment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderGetPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderGetPayment`: OrderPayment
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderGetPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGetPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderGetPaymentRequest**](OrderGetPaymentRequest.md) |  | 

### Return type

[**OrderPayment**](OrderPayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGetShipment

> OrderShipment OrderGetShipment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderGetShipmentRequest() // OrderGetShipmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderGetShipment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderGetShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderGetShipment`: OrderShipment
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderGetShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGetShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderGetShipmentRequest**](OrderGetShipmentRequest.md) |  | 

### Return type

[**OrderShipment**](OrderShipment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGetTransaction

> OrderTransaction OrderGetTransaction(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderGetTransactionRequest() // OrderGetTransactionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderGetTransaction(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderGetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderGetTransaction`: OrderTransaction
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderGetTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderGetTransactionRequest**](OrderGetTransactionRequest.md) |  | 

### Return type

[**OrderTransaction**](OrderTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderHoldOrder

> map[string]interface{} OrderHoldOrder(ctx).Body(body).Execute()

Order custom methods

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
    body := *openapiclient.NewOrderHoldOrderRequest() // OrderHoldOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderHoldOrder(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderHoldOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderHoldOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderHoldOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderHoldOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderHoldOrderRequest**](OrderHoldOrderRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderImportOrder

> OrderOrderData OrderImportOrder(ctx).Body(body).Execute()

Import orders API

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
    body := *openapiclient.NewOrderImportOrderRequest() // OrderImportOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderImportOrder(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderImportOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderImportOrder`: OrderOrderData
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderImportOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderImportOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderImportOrderRequest**](OrderImportOrderRequest.md) |  | 

### Return type

[**OrderOrderData**](OrderOrderData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderListFulfillments

> OrderListFulfillmentsResponse OrderListFulfillments(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderListFulfillmentsRequest() // OrderListFulfillmentsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderListFulfillments(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderListFulfillments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderListFulfillments`: OrderListFulfillmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderListFulfillments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderListFulfillmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderListFulfillmentsRequest**](OrderListFulfillmentsRequest.md) |  | 

### Return type

[**OrderListFulfillmentsResponse**](OrderListFulfillmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderListOrders

> OrderListOrdersResponse OrderListOrders(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderListOrdersRequest() // OrderListOrdersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderListOrders(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderListOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderListOrders`: OrderListOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderListOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderListOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderListOrdersRequest**](OrderListOrdersRequest.md) |  | 

### Return type

[**OrderListOrdersResponse**](OrderListOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderListOrdersByCustomer

> OrderListOrdersByCustomerResponse OrderListOrdersByCustomer(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderListOrdersByCustomerRequest() // OrderListOrdersByCustomerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderListOrdersByCustomer(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderListOrdersByCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderListOrdersByCustomer`: OrderListOrdersByCustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderListOrdersByCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderListOrdersByCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderListOrdersByCustomerRequest**](OrderListOrdersByCustomerRequest.md) |  | 

### Return type

[**OrderListOrdersByCustomerResponse**](OrderListOrdersByCustomerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderListOrdersByNumbers

> OrderListOrdersByNumbersResponse OrderListOrdersByNumbers(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderListOrdersByNumbersRequest() // OrderListOrdersByNumbersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderListOrdersByNumbers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderListOrdersByNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderListOrdersByNumbers`: OrderListOrdersByNumbersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderListOrdersByNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderListOrdersByNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderListOrdersByNumbersRequest**](OrderListOrdersByNumbersRequest.md) |  | 

### Return type

[**OrderListOrdersByNumbersResponse**](OrderListOrdersByNumbersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderListShipments

> OrderListShipmentsResponse OrderListShipments(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderListShipmentsRequest() // OrderListShipmentsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderListShipments(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderListShipments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderListShipments`: OrderListShipmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderListShipments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderListShipmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderListShipmentsRequest**](OrderListShipmentsRequest.md) |  | 

### Return type

[**OrderListShipmentsResponse**](OrderListShipmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderPrintOrdersLabels

> OrderPrintOrdersLabelsResponse OrderPrintOrdersLabels(ctx).Body(body).Execute()

Labels custom methods

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
    body := *openapiclient.NewOrderPrintOrdersLabelsRequest() // OrderPrintOrdersLabelsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderPrintOrdersLabels(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderPrintOrdersLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderPrintOrdersLabels`: OrderPrintOrdersLabelsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderPrintOrdersLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderPrintOrdersLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderPrintOrdersLabelsRequest**](OrderPrintOrdersLabelsRequest.md) |  | 

### Return type

[**OrderPrintOrdersLabelsResponse**](OrderPrintOrdersLabelsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderQuashFulfillment

> map[string]interface{} OrderQuashFulfillment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderQuashFulfillmentRequest() // OrderQuashFulfillmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderQuashFulfillment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderQuashFulfillment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderQuashFulfillment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderQuashFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderQuashFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderQuashFulfillmentRequest**](OrderQuashFulfillmentRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderQuashShipment

> map[string]interface{} OrderQuashShipment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderQuashShipmentRequest() // OrderQuashShipmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderQuashShipment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderQuashShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderQuashShipment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderQuashShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderQuashShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderQuashShipmentRequest**](OrderQuashShipmentRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderReceiveFulfillment

> map[string]interface{} OrderReceiveFulfillment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderReceiveFulfillmentRequest() // OrderReceiveFulfillmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderReceiveFulfillment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderReceiveFulfillment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderReceiveFulfillment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderReceiveFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderReceiveFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderReceiveFulfillmentRequest**](OrderReceiveFulfillmentRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderReportFulfillmentError

> map[string]interface{} OrderReportFulfillmentError(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderReportFulfillmentErrorRequest() // OrderReportFulfillmentErrorRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderReportFulfillmentError(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderReportFulfillmentError``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderReportFulfillmentError`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderReportFulfillmentError`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderReportFulfillmentErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderReportFulfillmentErrorRequest**](OrderReportFulfillmentErrorRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderReportFulfillmentNotResolvable

> map[string]interface{} OrderReportFulfillmentNotResolvable(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderReportFulfillmentNotResolvableRequest() // OrderReportFulfillmentNotResolvableRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderReportFulfillmentNotResolvable(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderReportFulfillmentNotResolvable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderReportFulfillmentNotResolvable`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderReportFulfillmentNotResolvable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderReportFulfillmentNotResolvableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderReportFulfillmentNotResolvableRequest**](OrderReportFulfillmentNotResolvableRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderReportFulfillmentReady

> map[string]interface{} OrderReportFulfillmentReady(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderReportFulfillmentReadyRequest() // OrderReportFulfillmentReadyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderReportFulfillmentReady(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderReportFulfillmentReady``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderReportFulfillmentReady`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderReportFulfillmentReady`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderReportFulfillmentReadyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderReportFulfillmentReadyRequest**](OrderReportFulfillmentReadyRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderReportShipmentDelivery

> map[string]interface{} OrderReportShipmentDelivery(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderReportShipmentDeliveryRequest() // OrderReportShipmentDeliveryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderReportShipmentDelivery(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderReportShipmentDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderReportShipmentDelivery`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderReportShipmentDelivery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderReportShipmentDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderReportShipmentDeliveryRequest**](OrderReportShipmentDeliveryRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderReportShipmentMissingStock

> map[string]interface{} OrderReportShipmentMissingStock(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderReportShipmentMissingStockRequest() // OrderReportShipmentMissingStockRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderReportShipmentMissingStock(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderReportShipmentMissingStock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderReportShipmentMissingStock`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderReportShipmentMissingStock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderReportShipmentMissingStockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderReportShipmentMissingStockRequest**](OrderReportShipmentMissingStockRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderResolveShipmentMissingStock

> map[string]interface{} OrderResolveShipmentMissingStock(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderResolveShipmentMissingStockRequest() // OrderResolveShipmentMissingStockRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderResolveShipmentMissingStock(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderResolveShipmentMissingStock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderResolveShipmentMissingStock`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderResolveShipmentMissingStock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderResolveShipmentMissingStockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderResolveShipmentMissingStockRequest**](OrderResolveShipmentMissingStockRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderRetryFulfillment

> map[string]interface{} OrderRetryFulfillment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderRetryFulfillmentRequest() // OrderRetryFulfillmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderRetryFulfillment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderRetryFulfillment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderRetryFulfillment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderRetryFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderRetryFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderRetryFulfillmentRequest**](OrderRetryFulfillmentRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderSearchOrders

> OrderSearchOrdersResponse OrderSearchOrders(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderSearchOrdersRequest() // OrderSearchOrdersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderSearchOrders(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderSearchOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderSearchOrders`: OrderSearchOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderSearchOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderSearchOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderSearchOrdersRequest**](OrderSearchOrdersRequest.md) |  | 

### Return type

[**OrderSearchOrdersResponse**](OrderSearchOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderSendFulfillment

> map[string]interface{} OrderSendFulfillment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderSendFulfillmentRequest() // OrderSendFulfillmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderSendFulfillment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderSendFulfillment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderSendFulfillment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderSendFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderSendFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderSendFulfillmentRequest**](OrderSendFulfillmentRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderSendOrderNotification

> map[string]interface{} OrderSendOrderNotification(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderSendOrderNotificationRequest() // OrderSendOrderNotificationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderSendOrderNotification(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderSendOrderNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderSendOrderNotification`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderSendOrderNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderSendOrderNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderSendOrderNotificationRequest**](OrderSendOrderNotificationRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderStartFulfillmentProcessing

> map[string]interface{} OrderStartFulfillmentProcessing(ctx).Body(body).Execute()

Fulfillment custom methods

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
    body := *openapiclient.NewOrderStartFulfillmentProcessingRequest() // OrderStartFulfillmentProcessingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderStartFulfillmentProcessing(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderStartFulfillmentProcessing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderStartFulfillmentProcessing`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderStartFulfillmentProcessing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderStartFulfillmentProcessingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderStartFulfillmentProcessingRequest**](OrderStartFulfillmentProcessingRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderStartShipmentProcessing

> map[string]interface{} OrderStartShipmentProcessing(ctx).Body(body).Execute()

Shipment custom methods

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
    body := *openapiclient.NewOrderStartShipmentProcessingRequest() // OrderStartShipmentProcessingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderStartShipmentProcessing(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderStartShipmentProcessing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderStartShipmentProcessing`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderStartShipmentProcessing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderStartShipmentProcessingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderStartShipmentProcessingRequest**](OrderStartShipmentProcessingRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderUnholdOrder

> map[string]interface{} OrderUnholdOrder(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderUnholdOrderRequest() // OrderUnholdOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderUnholdOrder(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderUnholdOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderUnholdOrder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderUnholdOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderUnholdOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderUnholdOrderRequest**](OrderUnholdOrderRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderUpdateOrder

> OrderOrderData OrderUpdateOrder(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderUpdateOrderRequest() // OrderUpdateOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderUpdateOrder(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderUpdateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderUpdateOrder`: OrderOrderData
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderUpdateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderUpdateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderUpdateOrderRequest**](OrderUpdateOrderRequest.md) |  | 

### Return type

[**OrderOrderData**](OrderOrderData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderUpdatePayment

> OrderPayment OrderUpdatePayment(ctx).Body(body).Execute()



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
    body := *openapiclient.NewOrderUpdatePaymentRequest() // OrderUpdatePaymentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.OrderUpdatePayment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.OrderUpdatePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderUpdatePayment`: OrderPayment
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.OrderUpdatePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderUpdatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrderUpdatePaymentRequest**](OrderUpdatePaymentRequest.md) |  | 

### Return type

[**OrderPayment**](OrderPayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


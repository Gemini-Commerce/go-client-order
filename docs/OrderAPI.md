# \OrderAPI

All URIs are relative to *https://dom.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveOrder**](OrderAPI.md#ApproveOrder) | **Post** /order.Order/ApproveOrder | Approve Order
[**AssignShipment**](OrderAPI.md#AssignShipment) | **Post** /order.Order/AssignShipment | Assign Shipment
[**CalculateRefund**](OrderAPI.md#CalculateRefund) | **Post** /order.Order/CalculateRefund | Calculate Refund
[**CancelFulfillment**](OrderAPI.md#CancelFulfillment) | **Post** /order.Order/CancelFulfillment | Cancel Fulfillment
[**CancelOrder**](OrderAPI.md#CancelOrder) | **Post** /order.Order/CancelOrder | Cancel Order
[**CancelShipment**](OrderAPI.md#CancelShipment) | **Post** /order.Order/CancelShipment | Cancel Shipment
[**CompleteShipmentPacking**](OrderAPI.md#CompleteShipmentPacking) | **Post** /order.Order/CompleteShipmentPacking | Complete Shipment Packing
[**CreateFulfillment**](OrderAPI.md#CreateFulfillment) | **Post** /order.Order/CreateFulfillment | Create Fulfillment
[**CreateOrder**](OrderAPI.md#CreateOrder) | **Post** /order.Order/CreateOrder | Create Order
[**CreateOrderHistory**](OrderAPI.md#CreateOrderHistory) | **Post** /order.Order/CreateHistory | Create Order History
[**CreatePayment**](OrderAPI.md#CreatePayment) | **Post** /order.Order/CreatePayment | Create Payment
[**CreatePaymentTransaction**](OrderAPI.md#CreatePaymentTransaction) | **Post** /order.Order/CreatePaymentTransaction | Create Payment Transaction
[**CreateRefund**](OrderAPI.md#CreateRefund) | **Post** /order.Order/CreateRefund | Create Refund
[**CreateRefundTransaction**](OrderAPI.md#CreateRefundTransaction) | **Post** /order.Order/CreateRefundTransaction | Create Refund Transaction
[**CreateShipment**](OrderAPI.md#CreateShipment) | **Post** /order.Order/CreateShipment | Create Shipment
[**DeleteOrder**](OrderAPI.md#DeleteOrder) | **Post** /order.Order/DeleteOrder | Delete Order
[**GetFulfillment**](OrderAPI.md#GetFulfillment) | **Post** /order.Order/GetFulfillment | Get Fulfillment
[**GetOrder**](OrderAPI.md#GetOrder) | **Post** /order.Order/GetOrder | Get Order
[**GetOrderByCartId**](OrderAPI.md#GetOrderByCartId) | **Post** /order.Order/GetOrderByCartId | Get Order by Cart ID
[**GetOrderByOrderNumber**](OrderAPI.md#GetOrderByOrderNumber) | **Post** /order.Order/GetOrderByOrderNumber | Get Order by Order Number
[**GetPayment**](OrderAPI.md#GetPayment) | **Post** /order.Order/GetPayment | Get Payment
[**GetShipment**](OrderAPI.md#GetShipment) | **Post** /order.Order/GetShipment | Get Shipment
[**GetTransaction**](OrderAPI.md#GetTransaction) | **Post** /order.Order/GetTransaction | Get Transaction
[**HoldOrder**](OrderAPI.md#HoldOrder) | **Post** /order.Order/HoldOrder | Hold Order
[**ImportOrder**](OrderAPI.md#ImportOrder) | **Post** /order.Order/ImportOrder | Import Order
[**ListFulfillments**](OrderAPI.md#ListFulfillments) | **Post** /order.Order/ListFulfillments | List Fulfillments
[**ListOrders**](OrderAPI.md#ListOrders) | **Post** /order.Order/ListOrders | List Orders
[**ListOrdersByCustomer**](OrderAPI.md#ListOrdersByCustomer) | **Post** /order.Order/ListOrdersByCustomer | List Orders by Customer
[**ListOrdersByNumbers**](OrderAPI.md#ListOrdersByNumbers) | **Post** /order.Order/ListOrdersByNumbers | List Orders by Numbers
[**ListShipments**](OrderAPI.md#ListShipments) | **Post** /order.Order/ListShipments | List Shipments
[**PrintOrdersLabels**](OrderAPI.md#PrintOrdersLabels) | **Post** /order.Order/PrintOrdersLabels | Print Orders Labels
[**QuashFulfillment**](OrderAPI.md#QuashFulfillment) | **Post** /order.Order/QuashFulfillment | Quash Fulfillment
[**QuashShipment**](OrderAPI.md#QuashShipment) | **Post** /order.Order/QuashShipment | Quash Shipment
[**ReceiveFulfillment**](OrderAPI.md#ReceiveFulfillment) | **Post** /order.Order/ReceiveFulfillment | Receive Fulfillment
[**ReportFulfillmentError**](OrderAPI.md#ReportFulfillmentError) | **Post** /order.Order/ReportFulfillmentError | Report Fulfillment Error
[**ReportFulfillmentNotResolvable**](OrderAPI.md#ReportFulfillmentNotResolvable) | **Post** /order.Order/ReportFulfillmentNotResolvable | Report Fulfillment Not Resolvable
[**ReportFulfillmentReady**](OrderAPI.md#ReportFulfillmentReady) | **Post** /order.Order/ReportFulfillmentReady | Report Fulfillment Ready
[**ReportShipmentDelivery**](OrderAPI.md#ReportShipmentDelivery) | **Post** /order.Order/ReportShipmentDelivery | Report Shipment Delivery
[**ReportShipmentMissingStock**](OrderAPI.md#ReportShipmentMissingStock) | **Post** /order.Order/ReportShipmentMissingStock | Report Shipment Missing Stock
[**ResolveShipmentMissingStock**](OrderAPI.md#ResolveShipmentMissingStock) | **Post** /order.Order/ResolveShipmentMissingStock | Resolve Shipment Missing Stock
[**RetryFulfillment**](OrderAPI.md#RetryFulfillment) | **Post** /order.Order/RetryFulfillment | Retry Fulfillment
[**SearchOrders**](OrderAPI.md#SearchOrders) | **Post** /order.Order/SearchOrders | Search Orders
[**SendFulfillment**](OrderAPI.md#SendFulfillment) | **Post** /order.Order/SendFulfillment | Send Fulfillment
[**SendOrderNotification**](OrderAPI.md#SendOrderNotification) | **Post** /order.Order/SendOrderNotification | Send Order Notification
[**StartFulfillmentProcessing**](OrderAPI.md#StartFulfillmentProcessing) | **Post** /order.Order/StartFulfillmentProcessing | Start Fulfillment Processing
[**StartShipmentProcessing**](OrderAPI.md#StartShipmentProcessing) | **Post** /order.Order/StartShipmentProcessing | Start Shipment Processing
[**UnholdOrder**](OrderAPI.md#UnholdOrder) | **Post** /order.Order/UnholdOrder | Unhold Order
[**UpdateOrder**](OrderAPI.md#UpdateOrder) | **Post** /order.Order/UpdateOrder | Update Order
[**UpdatePayment**](OrderAPI.md#UpdatePayment) | **Post** /order.Order/UpdatePayment | Update Payment



## ApproveOrder

> map[string]interface{} ApproveOrder(ctx).Body(body).Execute()

Approve Order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderApproveOrderRequest() // OrderApproveOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ApproveOrder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ApproveOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveOrder`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ApproveOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApproveOrderRequest struct via the builder pattern


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


## AssignShipment

> map[string]interface{} AssignShipment(ctx).Body(body).Execute()

Assign Shipment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderAssignShipmentRequest() // OrderAssignShipmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.AssignShipment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.AssignShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignShipment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.AssignShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignShipmentRequest struct via the builder pattern


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


## CalculateRefund

> OrderCalculateRefundResponse CalculateRefund(ctx).Body(body).Execute()

Calculate Refund

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCalculateRefundRequest() // OrderCalculateRefundRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CalculateRefund(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CalculateRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateRefund`: OrderCalculateRefundResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CalculateRefund`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalculateRefundRequest struct via the builder pattern


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


## CancelFulfillment

> map[string]interface{} CancelFulfillment(ctx).Body(body).Execute()

Cancel Fulfillment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCancelFulfillmentRequest() // OrderCancelFulfillmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CancelFulfillment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CancelFulfillment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelFulfillment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CancelFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelFulfillmentRequest struct via the builder pattern


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


## CancelOrder

> map[string]interface{} CancelOrder(ctx).Body(body).Execute()

Cancel Order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCancelOrderRequest() // OrderCancelOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CancelOrder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CancelOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelOrder`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CancelOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderRequest struct via the builder pattern


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


## CancelShipment

> map[string]interface{} CancelShipment(ctx).Body(body).Execute()

Cancel Shipment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCancelShipmentRequest() // OrderCancelShipmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CancelShipment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CancelShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelShipment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CancelShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelShipmentRequest struct via the builder pattern


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


## CompleteShipmentPacking

> map[string]interface{} CompleteShipmentPacking(ctx).Body(body).Execute()

Complete Shipment Packing

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCompleteShipmentPackingRequest() // OrderCompleteShipmentPackingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CompleteShipmentPacking(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CompleteShipmentPacking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteShipmentPacking`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CompleteShipmentPacking`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteShipmentPackingRequest struct via the builder pattern


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


## CreateFulfillment

> OrderFulfillment CreateFulfillment(ctx).Body(body).Execute()

Create Fulfillment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCreateFulfillmentRequest() // OrderCreateFulfillmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CreateFulfillment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CreateFulfillment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFulfillment`: OrderFulfillment
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CreateFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFulfillmentRequest struct via the builder pattern


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


## CreateOrder

> OrderOrderData CreateOrder(ctx).Body(body).Execute()

Create Order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCreateOrderRequest() // OrderCreateOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CreateOrder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CreateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrder`: OrderOrderData
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


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


## CreateOrderHistory

> OrderDataHistory CreateOrderHistory(ctx).Body(body).Execute()

Create Order History

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCreateHistoryRequest() // OrderCreateHistoryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CreateOrderHistory(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CreateOrderHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrderHistory`: OrderDataHistory
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CreateOrderHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderHistoryRequest struct via the builder pattern


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


## CreatePayment

> OrderPayment CreatePayment(ctx).Body(body).Execute()

Create Payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCreatePaymentRequest() // OrderCreatePaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CreatePayment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CreatePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePayment`: OrderPayment
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CreatePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentRequest struct via the builder pattern


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


## CreatePaymentTransaction

> OrderTransaction CreatePaymentTransaction(ctx).Body(body).Execute()

Create Payment Transaction

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCreatePaymentTransactionRequest() // OrderCreatePaymentTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CreatePaymentTransaction(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CreatePaymentTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentTransaction`: OrderTransaction
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CreatePaymentTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentTransactionRequest struct via the builder pattern


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


## CreateRefund

> OrderRefund CreateRefund(ctx).Body(body).Execute()

Create Refund

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCreateRefundRequest() // OrderCreateRefundRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CreateRefund(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CreateRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRefund`: OrderRefund
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CreateRefund`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRefundRequest struct via the builder pattern


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


## CreateRefundTransaction

> OrderTransaction CreateRefundTransaction(ctx).Body(body).Execute()

Create Refund Transaction

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCreateRefundTransactionRequest() // OrderCreateRefundTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CreateRefundTransaction(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CreateRefundTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRefundTransaction`: OrderTransaction
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CreateRefundTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRefundTransactionRequest struct via the builder pattern


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


## CreateShipment

> OrderShipment CreateShipment(ctx).Body(body).Execute()

Create Shipment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderCreateShipmentRequest() // OrderCreateShipmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.CreateShipment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.CreateShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShipment`: OrderShipment
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.CreateShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShipmentRequest struct via the builder pattern


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


## DeleteOrder

> map[string]interface{} DeleteOrder(ctx).Body(body).Execute()

Delete Order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderDeleteOrderRequest() // OrderDeleteOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.DeleteOrder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.DeleteOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrder`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.DeleteOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrderRequest struct via the builder pattern


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


## GetFulfillment

> OrderFulfillment GetFulfillment(ctx).Body(body).Execute()

Get Fulfillment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderGetFulfillmentRequest() // OrderGetFulfillmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.GetFulfillment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetFulfillment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFulfillment`: OrderFulfillment
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFulfillmentRequest struct via the builder pattern


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


## GetOrder

> OrderOrderData GetOrder(ctx).Body(body).Execute()

Get Order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderGetOrderRequest() // OrderGetOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.GetOrder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrder`: OrderOrderData
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


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


## GetOrderByCartId

> OrderOrderData GetOrderByCartId(ctx).Body(body).Execute()

Get Order by Cart ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderGetOrderByCartIdRequest() // OrderGetOrderByCartIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.GetOrderByCartId(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetOrderByCartId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderByCartId`: OrderOrderData
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetOrderByCartId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderByCartIdRequest struct via the builder pattern


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


## GetOrderByOrderNumber

> OrderOrderData GetOrderByOrderNumber(ctx).Body(body).Execute()

Get Order by Order Number

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderGetOrderByOrderNumberRequest() // OrderGetOrderByOrderNumberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.GetOrderByOrderNumber(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetOrderByOrderNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderByOrderNumber`: OrderOrderData
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetOrderByOrderNumber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderByOrderNumberRequest struct via the builder pattern


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


## GetPayment

> OrderPayment GetPayment(ctx).Body(body).Execute()

Get Payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderGetPaymentRequest() // OrderGetPaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.GetPayment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPayment`: OrderPayment
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRequest struct via the builder pattern


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


## GetShipment

> OrderShipment GetShipment(ctx).Body(body).Execute()

Get Shipment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderGetShipmentRequest() // OrderGetShipmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.GetShipment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShipment`: OrderShipment
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetShipmentRequest struct via the builder pattern


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


## GetTransaction

> OrderTransaction GetTransaction(ctx).Body(body).Execute()

Get Transaction

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderGetTransactionRequest() // OrderGetTransactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.GetTransaction(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransaction`: OrderTransaction
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


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


## HoldOrder

> map[string]interface{} HoldOrder(ctx).Body(body).Execute()

Hold Order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderHoldOrderRequest() // OrderHoldOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.HoldOrder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.HoldOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HoldOrder`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.HoldOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHoldOrderRequest struct via the builder pattern


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


## ImportOrder

> OrderOrderData ImportOrder(ctx).Body(body).Execute()

Import Order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderImportOrderRequest() // OrderImportOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ImportOrder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ImportOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportOrder`: OrderOrderData
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ImportOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportOrderRequest struct via the builder pattern


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


## ListFulfillments

> OrderListFulfillmentsResponse ListFulfillments(ctx).Body(body).Execute()

List Fulfillments

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderListFulfillmentsRequest() // OrderListFulfillmentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ListFulfillments(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ListFulfillments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFulfillments`: OrderListFulfillmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ListFulfillments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFulfillmentsRequest struct via the builder pattern


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


## ListOrders

> OrderListOrdersResponse ListOrders(ctx).Body(body).Execute()

List Orders

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderListOrdersRequest() // OrderListOrdersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ListOrders(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ListOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrders`: OrderListOrdersResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ListOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrdersRequest struct via the builder pattern


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


## ListOrdersByCustomer

> OrderListOrdersByCustomerResponse ListOrdersByCustomer(ctx).Body(body).Execute()

List Orders by Customer

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderListOrdersByCustomerRequest() // OrderListOrdersByCustomerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ListOrdersByCustomer(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ListOrdersByCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrdersByCustomer`: OrderListOrdersByCustomerResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ListOrdersByCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrdersByCustomerRequest struct via the builder pattern


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


## ListOrdersByNumbers

> OrderListOrdersByNumbersResponse ListOrdersByNumbers(ctx).Body(body).Execute()

List Orders by Numbers

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderListOrdersByNumbersRequest() // OrderListOrdersByNumbersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ListOrdersByNumbers(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ListOrdersByNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrdersByNumbers`: OrderListOrdersByNumbersResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ListOrdersByNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrdersByNumbersRequest struct via the builder pattern


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


## ListShipments

> OrderListShipmentsResponse ListShipments(ctx).Body(body).Execute()

List Shipments

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderListShipmentsRequest() // OrderListShipmentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ListShipments(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ListShipments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListShipments`: OrderListShipmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ListShipments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListShipmentsRequest struct via the builder pattern


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


## PrintOrdersLabels

> OrderPrintOrdersLabelsResponse PrintOrdersLabels(ctx).Body(body).Execute()

Print Orders Labels

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderPrintOrdersLabelsRequest() // OrderPrintOrdersLabelsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.PrintOrdersLabels(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.PrintOrdersLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrintOrdersLabels`: OrderPrintOrdersLabelsResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.PrintOrdersLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintOrdersLabelsRequest struct via the builder pattern


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


## QuashFulfillment

> map[string]interface{} QuashFulfillment(ctx).Body(body).Execute()

Quash Fulfillment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderQuashFulfillmentRequest() // OrderQuashFulfillmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.QuashFulfillment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.QuashFulfillment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuashFulfillment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.QuashFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuashFulfillmentRequest struct via the builder pattern


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


## QuashShipment

> map[string]interface{} QuashShipment(ctx).Body(body).Execute()

Quash Shipment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderQuashShipmentRequest() // OrderQuashShipmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.QuashShipment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.QuashShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuashShipment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.QuashShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuashShipmentRequest struct via the builder pattern


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


## ReceiveFulfillment

> map[string]interface{} ReceiveFulfillment(ctx).Body(body).Execute()

Receive Fulfillment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderReceiveFulfillmentRequest() // OrderReceiveFulfillmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ReceiveFulfillment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ReceiveFulfillment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceiveFulfillment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ReceiveFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReceiveFulfillmentRequest struct via the builder pattern


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


## ReportFulfillmentError

> map[string]interface{} ReportFulfillmentError(ctx).Body(body).Execute()

Report Fulfillment Error

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderReportFulfillmentErrorRequest() // OrderReportFulfillmentErrorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ReportFulfillmentError(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ReportFulfillmentError``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFulfillmentError`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ReportFulfillmentError`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportFulfillmentErrorRequest struct via the builder pattern


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


## ReportFulfillmentNotResolvable

> map[string]interface{} ReportFulfillmentNotResolvable(ctx).Body(body).Execute()

Report Fulfillment Not Resolvable

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderReportFulfillmentNotResolvableRequest() // OrderReportFulfillmentNotResolvableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ReportFulfillmentNotResolvable(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ReportFulfillmentNotResolvable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFulfillmentNotResolvable`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ReportFulfillmentNotResolvable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportFulfillmentNotResolvableRequest struct via the builder pattern


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


## ReportFulfillmentReady

> map[string]interface{} ReportFulfillmentReady(ctx).Body(body).Execute()

Report Fulfillment Ready

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderReportFulfillmentReadyRequest() // OrderReportFulfillmentReadyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ReportFulfillmentReady(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ReportFulfillmentReady``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFulfillmentReady`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ReportFulfillmentReady`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportFulfillmentReadyRequest struct via the builder pattern


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


## ReportShipmentDelivery

> map[string]interface{} ReportShipmentDelivery(ctx).Body(body).Execute()

Report Shipment Delivery

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderReportShipmentDeliveryRequest() // OrderReportShipmentDeliveryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ReportShipmentDelivery(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ReportShipmentDelivery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportShipmentDelivery`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ReportShipmentDelivery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportShipmentDeliveryRequest struct via the builder pattern


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


## ReportShipmentMissingStock

> map[string]interface{} ReportShipmentMissingStock(ctx).Body(body).Execute()

Report Shipment Missing Stock

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderReportShipmentMissingStockRequest() // OrderReportShipmentMissingStockRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ReportShipmentMissingStock(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ReportShipmentMissingStock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportShipmentMissingStock`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ReportShipmentMissingStock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportShipmentMissingStockRequest struct via the builder pattern


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


## ResolveShipmentMissingStock

> map[string]interface{} ResolveShipmentMissingStock(ctx).Body(body).Execute()

Resolve Shipment Missing Stock

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderResolveShipmentMissingStockRequest() // OrderResolveShipmentMissingStockRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.ResolveShipmentMissingStock(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.ResolveShipmentMissingStock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveShipmentMissingStock`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.ResolveShipmentMissingStock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveShipmentMissingStockRequest struct via the builder pattern


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


## RetryFulfillment

> map[string]interface{} RetryFulfillment(ctx).Body(body).Execute()

Retry Fulfillment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderRetryFulfillmentRequest() // OrderRetryFulfillmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.RetryFulfillment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.RetryFulfillment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryFulfillment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.RetryFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetryFulfillmentRequest struct via the builder pattern


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


## SearchOrders

> OrderSearchOrdersResponse SearchOrders(ctx).Body(body).Execute()

Search Orders

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderSearchOrdersRequest() // OrderSearchOrdersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.SearchOrders(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.SearchOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrders`: OrderSearchOrdersResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.SearchOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrdersRequest struct via the builder pattern


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


## SendFulfillment

> map[string]interface{} SendFulfillment(ctx).Body(body).Execute()

Send Fulfillment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderSendFulfillmentRequest() // OrderSendFulfillmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.SendFulfillment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.SendFulfillment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendFulfillment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.SendFulfillment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendFulfillmentRequest struct via the builder pattern


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


## SendOrderNotification

> map[string]interface{} SendOrderNotification(ctx).Body(body).Execute()

Send Order Notification

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderSendOrderNotificationRequest() // OrderSendOrderNotificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.SendOrderNotification(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.SendOrderNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendOrderNotification`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.SendOrderNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendOrderNotificationRequest struct via the builder pattern


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


## StartFulfillmentProcessing

> map[string]interface{} StartFulfillmentProcessing(ctx).Body(body).Execute()

Start Fulfillment Processing

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderStartFulfillmentProcessingRequest() // OrderStartFulfillmentProcessingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.StartFulfillmentProcessing(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.StartFulfillmentProcessing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartFulfillmentProcessing`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.StartFulfillmentProcessing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartFulfillmentProcessingRequest struct via the builder pattern


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


## StartShipmentProcessing

> map[string]interface{} StartShipmentProcessing(ctx).Body(body).Execute()

Start Shipment Processing

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderStartShipmentProcessingRequest() // OrderStartShipmentProcessingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.StartShipmentProcessing(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.StartShipmentProcessing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartShipmentProcessing`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.StartShipmentProcessing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartShipmentProcessingRequest struct via the builder pattern


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


## UnholdOrder

> map[string]interface{} UnholdOrder(ctx).Body(body).Execute()

Unhold Order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderUnholdOrderRequest() // OrderUnholdOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.UnholdOrder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.UnholdOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnholdOrder`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.UnholdOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnholdOrderRequest struct via the builder pattern


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


## UpdateOrder

> OrderOrderData UpdateOrder(ctx).Body(body).Execute()

Update Order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderUpdateOrderRequest() // OrderUpdateOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.UpdateOrder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.UpdateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrder`: OrderOrderData
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.UpdateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderRequest struct via the builder pattern


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


## UpdatePayment

> OrderPayment UpdatePayment(ctx).Body(body).Execute()

Update Payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/gemini-commerce/go-client-order"
)

func main() {
	body := *openapiclient.NewOrderUpdatePaymentRequest() // OrderUpdatePaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.UpdatePayment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.UpdatePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePayment`: OrderPayment
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.UpdatePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentRequest struct via the builder pattern


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


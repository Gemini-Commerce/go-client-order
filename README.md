# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: version not set
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "bitbucket.org/gogemini/go-client-order"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OrderApi* | [**OrderApproveOrder**](docs/OrderApi.md#orderapproveorder) | **Post** /order.Order/ApproveOrder | 
*OrderApi* | [**OrderAssignShipment**](docs/OrderApi.md#orderassignshipment) | **Post** /order.Order/AssignShipment | 
*OrderApi* | [**OrderCalculateRefund**](docs/OrderApi.md#ordercalculaterefund) | **Post** /order.Order/CalculateRefund | Refund custom methods
*OrderApi* | [**OrderCancelFulfillment**](docs/OrderApi.md#ordercancelfulfillment) | **Post** /order.Order/CancelFulfillment | 
*OrderApi* | [**OrderCancelOrder**](docs/OrderApi.md#ordercancelorder) | **Post** /order.Order/CancelOrder | 
*OrderApi* | [**OrderCancelShipment**](docs/OrderApi.md#ordercancelshipment) | **Post** /order.Order/CancelShipment | 
*OrderApi* | [**OrderCompleteShipmentPacking**](docs/OrderApi.md#ordercompleteshipmentpacking) | **Post** /order.Order/CompleteShipmentPacking | 
*OrderApi* | [**OrderCreateFulfillment**](docs/OrderApi.md#ordercreatefulfillment) | **Post** /order.Order/CreateFulfillment | 
*OrderApi* | [**OrderCreateHistory**](docs/OrderApi.md#ordercreatehistory) | **Post** /order.Order/CreateHistory | 
*OrderApi* | [**OrderCreateOrder**](docs/OrderApi.md#ordercreateorder) | **Post** /order.Order/CreateOrder | 
*OrderApi* | [**OrderCreatePayment**](docs/OrderApi.md#ordercreatepayment) | **Post** /order.Order/CreatePayment | 
*OrderApi* | [**OrderCreatePaymentTransaction**](docs/OrderApi.md#ordercreatepaymenttransaction) | **Post** /order.Order/CreatePaymentTransaction | 
*OrderApi* | [**OrderCreateRefund**](docs/OrderApi.md#ordercreaterefund) | **Post** /order.Order/CreateRefund | 
*OrderApi* | [**OrderCreateRefundTransaction**](docs/OrderApi.md#ordercreaterefundtransaction) | **Post** /order.Order/CreateRefundTransaction | 
*OrderApi* | [**OrderCreateShipment**](docs/OrderApi.md#ordercreateshipment) | **Post** /order.Order/CreateShipment | 
*OrderApi* | [**OrderDeleteOrder**](docs/OrderApi.md#orderdeleteorder) | **Post** /order.Order/DeleteOrder | 
*OrderApi* | [**OrderGetFulfillment**](docs/OrderApi.md#ordergetfulfillment) | **Post** /order.Order/GetFulfillment | 
*OrderApi* | [**OrderGetOrder**](docs/OrderApi.md#ordergetorder) | **Post** /order.Order/GetOrder | 
*OrderApi* | [**OrderGetOrderByCartId**](docs/OrderApi.md#ordergetorderbycartid) | **Post** /order.Order/GetOrderByCartId | cart method
*OrderApi* | [**OrderGetOrderByOrderNumber**](docs/OrderApi.md#ordergetorderbyordernumber) | **Post** /order.Order/GetOrderByOrderNumber | 
*OrderApi* | [**OrderGetPayment**](docs/OrderApi.md#ordergetpayment) | **Post** /order.Order/GetPayment | 
*OrderApi* | [**OrderGetShipment**](docs/OrderApi.md#ordergetshipment) | **Post** /order.Order/GetShipment | 
*OrderApi* | [**OrderGetTransaction**](docs/OrderApi.md#ordergettransaction) | **Post** /order.Order/GetTransaction | 
*OrderApi* | [**OrderHoldOrder**](docs/OrderApi.md#orderholdorder) | **Post** /order.Order/HoldOrder | Order custom methods
*OrderApi* | [**OrderImportOrder**](docs/OrderApi.md#orderimportorder) | **Post** /order.Order/ImportOrder | Import orders API
*OrderApi* | [**OrderListFulfillments**](docs/OrderApi.md#orderlistfulfillments) | **Post** /order.Order/ListFulfillments | 
*OrderApi* | [**OrderListOrders**](docs/OrderApi.md#orderlistorders) | **Post** /order.Order/ListOrders | 
*OrderApi* | [**OrderListOrdersByCustomer**](docs/OrderApi.md#orderlistordersbycustomer) | **Post** /order.Order/ListOrdersByCustomer | 
*OrderApi* | [**OrderListOrdersByNumbers**](docs/OrderApi.md#orderlistordersbynumbers) | **Post** /order.Order/ListOrdersByNumbers | 
*OrderApi* | [**OrderListShipments**](docs/OrderApi.md#orderlistshipments) | **Post** /order.Order/ListShipments | 
*OrderApi* | [**OrderPrintOrdersLabels**](docs/OrderApi.md#orderprintorderslabels) | **Post** /order.Order/PrintOrdersLabels | Labels custom methods
*OrderApi* | [**OrderQuashFulfillment**](docs/OrderApi.md#orderquashfulfillment) | **Post** /order.Order/QuashFulfillment | 
*OrderApi* | [**OrderQuashShipment**](docs/OrderApi.md#orderquashshipment) | **Post** /order.Order/QuashShipment | 
*OrderApi* | [**OrderReceiveFulfillment**](docs/OrderApi.md#orderreceivefulfillment) | **Post** /order.Order/ReceiveFulfillment | 
*OrderApi* | [**OrderReportFulfillmentError**](docs/OrderApi.md#orderreportfulfillmenterror) | **Post** /order.Order/ReportFulfillmentError | 
*OrderApi* | [**OrderReportFulfillmentNotResolvable**](docs/OrderApi.md#orderreportfulfillmentnotresolvable) | **Post** /order.Order/ReportFulfillmentNotResolvable | 
*OrderApi* | [**OrderReportFulfillmentReady**](docs/OrderApi.md#orderreportfulfillmentready) | **Post** /order.Order/ReportFulfillmentReady | 
*OrderApi* | [**OrderReportShipmentDelivery**](docs/OrderApi.md#orderreportshipmentdelivery) | **Post** /order.Order/ReportShipmentDelivery | 
*OrderApi* | [**OrderReportShipmentMissingStock**](docs/OrderApi.md#orderreportshipmentmissingstock) | **Post** /order.Order/ReportShipmentMissingStock | 
*OrderApi* | [**OrderResolveShipmentMissingStock**](docs/OrderApi.md#orderresolveshipmentmissingstock) | **Post** /order.Order/ResolveShipmentMissingStock | 
*OrderApi* | [**OrderRetryFulfillment**](docs/OrderApi.md#orderretryfulfillment) | **Post** /order.Order/RetryFulfillment | 
*OrderApi* | [**OrderSearchOrders**](docs/OrderApi.md#ordersearchorders) | **Post** /order.Order/SearchOrders | 
*OrderApi* | [**OrderSendFulfillment**](docs/OrderApi.md#ordersendfulfillment) | **Post** /order.Order/SendFulfillment | 
*OrderApi* | [**OrderSendOrderNotification**](docs/OrderApi.md#ordersendordernotification) | **Post** /order.Order/SendOrderNotification | 
*OrderApi* | [**OrderStartFulfillmentProcessing**](docs/OrderApi.md#orderstartfulfillmentprocessing) | **Post** /order.Order/StartFulfillmentProcessing | Fulfillment custom methods
*OrderApi* | [**OrderStartShipmentProcessing**](docs/OrderApi.md#orderstartshipmentprocessing) | **Post** /order.Order/StartShipmentProcessing | Shipment custom methods
*OrderApi* | [**OrderUnholdOrder**](docs/OrderApi.md#orderunholdorder) | **Post** /order.Order/UnholdOrder | 
*OrderApi* | [**OrderUpdateOrder**](docs/OrderApi.md#orderupdateorder) | **Post** /order.Order/UpdateOrder | 
*OrderApi* | [**OrderUpdatePayment**](docs/OrderApi.md#orderupdatepayment) | **Post** /order.Order/UpdatePayment | 


## Documentation For Models

 - [CreateOrderRequestInitialPayment](docs/CreateOrderRequestInitialPayment.md)
 - [ImportOrderRequestImportedPayment](docs/ImportOrderRequestImportedPayment.md)
 - [InitialPaymentInitialTransaction](docs/InitialPaymentInitialTransaction.md)
 - [OrderApproveOrderRequest](docs/OrderApproveOrderRequest.md)
 - [OrderAssignShipmentRequest](docs/OrderAssignShipmentRequest.md)
 - [OrderByDirection](docs/OrderByDirection.md)
 - [OrderCalculateRefundRequest](docs/OrderCalculateRefundRequest.md)
 - [OrderCalculateRefundResponse](docs/OrderCalculateRefundResponse.md)
 - [OrderCancelFulfillmentRequest](docs/OrderCancelFulfillmentRequest.md)
 - [OrderCancelOrderRequest](docs/OrderCancelOrderRequest.md)
 - [OrderCancelShipmentRequest](docs/OrderCancelShipmentRequest.md)
 - [OrderCompleteShipmentPackingRequest](docs/OrderCompleteShipmentPackingRequest.md)
 - [OrderCreateFulfillmentRequest](docs/OrderCreateFulfillmentRequest.md)
 - [OrderCreateHistoryRequest](docs/OrderCreateHistoryRequest.md)
 - [OrderCreateOrderRequest](docs/OrderCreateOrderRequest.md)
 - [OrderCreatePaymentRequest](docs/OrderCreatePaymentRequest.md)
 - [OrderCreatePaymentTransactionRequest](docs/OrderCreatePaymentTransactionRequest.md)
 - [OrderCreateRefundRequest](docs/OrderCreateRefundRequest.md)
 - [OrderCreateRefundTransactionRequest](docs/OrderCreateRefundTransactionRequest.md)
 - [OrderCreateShipmentRequest](docs/OrderCreateShipmentRequest.md)
 - [OrderCurrency](docs/OrderCurrency.md)
 - [OrderDataCustomerInfo](docs/OrderDataCustomerInfo.md)
 - [OrderDataHistory](docs/OrderDataHistory.md)
 - [OrderDataPaymentInfo](docs/OrderDataPaymentInfo.md)
 - [OrderDataPromotionInfo](docs/OrderDataPromotionInfo.md)
 - [OrderDataShipmentInfo](docs/OrderDataShipmentInfo.md)
 - [OrderDataSubtotal](docs/OrderDataSubtotal.md)
 - [OrderDataSubtotalCode](docs/OrderDataSubtotalCode.md)
 - [OrderDataTotal](docs/OrderDataTotal.md)
 - [OrderDataTotalCode](docs/OrderDataTotalCode.md)
 - [OrderDeleteOrderRequest](docs/OrderDeleteOrderRequest.md)
 - [OrderFulfillment](docs/OrderFulfillment.md)
 - [OrderFulfillmentItem](docs/OrderFulfillmentItem.md)
 - [OrderGetFulfillmentRequest](docs/OrderGetFulfillmentRequest.md)
 - [OrderGetOrderByCartIdRequest](docs/OrderGetOrderByCartIdRequest.md)
 - [OrderGetOrderByOrderNumberRequest](docs/OrderGetOrderByOrderNumberRequest.md)
 - [OrderGetOrderRequest](docs/OrderGetOrderRequest.md)
 - [OrderGetPaymentRequest](docs/OrderGetPaymentRequest.md)
 - [OrderGetShipmentRequest](docs/OrderGetShipmentRequest.md)
 - [OrderGetTransactionRequest](docs/OrderGetTransactionRequest.md)
 - [OrderHoldOrderRequest](docs/OrderHoldOrderRequest.md)
 - [OrderImportOrderRequest](docs/OrderImportOrderRequest.md)
 - [OrderListFulfillmentsRequest](docs/OrderListFulfillmentsRequest.md)
 - [OrderListFulfillmentsResponse](docs/OrderListFulfillmentsResponse.md)
 - [OrderListOrdersByCustomerRequest](docs/OrderListOrdersByCustomerRequest.md)
 - [OrderListOrdersByCustomerResponse](docs/OrderListOrdersByCustomerResponse.md)
 - [OrderListOrdersByNumbersRequest](docs/OrderListOrdersByNumbersRequest.md)
 - [OrderListOrdersByNumbersResponse](docs/OrderListOrdersByNumbersResponse.md)
 - [OrderListOrdersRequest](docs/OrderListOrdersRequest.md)
 - [OrderListOrdersResponse](docs/OrderListOrdersResponse.md)
 - [OrderListShipmentsRequest](docs/OrderListShipmentsRequest.md)
 - [OrderListShipmentsResponse](docs/OrderListShipmentsResponse.md)
 - [OrderLocalizedText](docs/OrderLocalizedText.md)
 - [OrderMoney](docs/OrderMoney.md)
 - [OrderOrderBy](docs/OrderOrderBy.md)
 - [OrderOrderData](docs/OrderOrderData.md)
 - [OrderOrderDataItem](docs/OrderOrderDataItem.md)
 - [OrderPayment](docs/OrderPayment.md)
 - [OrderPaymentAmount](docs/OrderPaymentAmount.md)
 - [OrderPaymentAmountCode](docs/OrderPaymentAmountCode.md)
 - [OrderPaymentFilter](docs/OrderPaymentFilter.md)
 - [OrderPaymentFilterCondition](docs/OrderPaymentFilterCondition.md)
 - [OrderPostalAddress](docs/OrderPostalAddress.md)
 - [OrderPrintOrdersLabelsRequest](docs/OrderPrintOrdersLabelsRequest.md)
 - [OrderPrintOrdersLabelsResponse](docs/OrderPrintOrdersLabelsResponse.md)
 - [OrderQuashFulfillmentRequest](docs/OrderQuashFulfillmentRequest.md)
 - [OrderQuashShipmentRequest](docs/OrderQuashShipmentRequest.md)
 - [OrderReceiveFulfillmentRequest](docs/OrderReceiveFulfillmentRequest.md)
 - [OrderRefund](docs/OrderRefund.md)
 - [OrderRefundAmount](docs/OrderRefundAmount.md)
 - [OrderRefundAmountCode](docs/OrderRefundAmountCode.md)
 - [OrderRefundItem](docs/OrderRefundItem.md)
 - [OrderReportFulfillmentErrorRequest](docs/OrderReportFulfillmentErrorRequest.md)
 - [OrderReportFulfillmentNotResolvableRequest](docs/OrderReportFulfillmentNotResolvableRequest.md)
 - [OrderReportFulfillmentReadyRequest](docs/OrderReportFulfillmentReadyRequest.md)
 - [OrderReportShipmentDeliveryRequest](docs/OrderReportShipmentDeliveryRequest.md)
 - [OrderReportShipmentMissingStockRequest](docs/OrderReportShipmentMissingStockRequest.md)
 - [OrderResolveShipmentMissingStockRequest](docs/OrderResolveShipmentMissingStockRequest.md)
 - [OrderRetryFulfillmentRequest](docs/OrderRetryFulfillmentRequest.md)
 - [OrderSearchOrdersRequest](docs/OrderSearchOrdersRequest.md)
 - [OrderSearchOrdersResponse](docs/OrderSearchOrdersResponse.md)
 - [OrderSendFulfillmentRequest](docs/OrderSendFulfillmentRequest.md)
 - [OrderSendOrderNotificationRequest](docs/OrderSendOrderNotificationRequest.md)
 - [OrderShipment](docs/OrderShipment.md)
 - [OrderShipmentItem](docs/OrderShipmentItem.md)
 - [OrderStartFulfillmentProcessingRequest](docs/OrderStartFulfillmentProcessingRequest.md)
 - [OrderStartShipmentProcessingRequest](docs/OrderStartShipmentProcessingRequest.md)
 - [OrderStatusFilter](docs/OrderStatusFilter.md)
 - [OrderStatusFilterCondition](docs/OrderStatusFilterCondition.md)
 - [OrderTransaction](docs/OrderTransaction.md)
 - [OrderTransactionType](docs/OrderTransactionType.md)
 - [OrderUnholdOrderRequest](docs/OrderUnholdOrderRequest.md)
 - [OrderUpdateOrderRequest](docs/OrderUpdateOrderRequest.md)
 - [OrderUpdatePaymentRequest](docs/OrderUpdatePaymentRequest.md)
 - [PaymentCcInfo](docs/PaymentCcInfo.md)
 - [PrintOrdersLabelsResponseFailedOrder](docs/PrintOrdersLabelsResponseFailedOrder.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [ProtobufNullValue](docs/ProtobufNullValue.md)
 - [RpcStatus](docs/RpcStatus.md)
 - [ShipmentTracking](docs/ShipmentTracking.md)
 - [UpdateOrderRequestPayload](docs/UpdateOrderRequestPayload.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author




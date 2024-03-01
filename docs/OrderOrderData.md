# # OrderOrderData


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt**| [**time.Time**](time.Time.md) |   | [optional] [readonly]
**UpdatedAt**| [**time.Time**](time.Time.md) |   | [optional] [readonly]
**Id**| **string** |   | [optional] [readonly]
**Grn**| **string** |   | [optional] [readonly]
**Number**| **string** |   | [optional]
**Status**| **string** |   | [optional] [readonly]
**Channel**| **string** |   | [optional]
**Market**| **string** |   | [optional]
**Locale**| **string** |   |
**AdditionalInfo**| **map[string]interface{}** |   | [optional]
**Documents**| [**[]OrderDataDocument**](OrderDataDocument.md) |   | [optional]
**Items**| [**[]OrderOrderDataItem**](OrderOrderDataItem.md) |   | [optional]
**Payments**| [**[]OrderPayment**](OrderPayment.md) |   | [optional]
**Shipments**| [**[]OrderShipment**](OrderShipment.md) |   | [optional]
**PaymentsInfo**| [**[]OrderDataPaymentInfo**](OrderDataPaymentInfo.md) |   | [optional]
**ShipmentsInfo**| [**[]OrderDataShipmentInfo**](OrderDataShipmentInfo.md) |   | [optional]
**Promotions**| [**[]OrderDataPromotionInfo**](OrderDataPromotionInfo.md) |   | [optional]
**Currency**| [**OrderCurrency**](OrderCurrency.md) |  for more information please, see Model/OrderCurrency.php  | [optional] [default to ORDERCURRENCY_XXX]
**Subtotals**| [**map[string]OrderDataSubtotal**](OrderDataSubtotal.md) |   | [optional]
**Totals**| [**map[string]OrderDataTotal**](OrderDataTotal.md) |   | [optional]
**VatIncluded**| **bool** |   | [optional]
**BillingAddress**| [**OrderPostalAddress**](OrderPostalAddress.md) |   | [optional]
**ShippingAddress**| [**OrderPostalAddress**](OrderPostalAddress.md) |   | [optional]
**CustomerInfo**| [**OrderDataCustomerInfo**](OrderDataCustomerInfo.md) |   | [optional]
**CartGrn**| **string** |   | [optional]
**OnHold**| **bool** |   | [optional]
**HistoryEvents**| [**[]OrderDataHistory**](OrderDataHistory.md) |   | [optional]
**Fulfillments**| [**[]OrderFulfillment**](OrderFulfillment.md) |   | [optional] [readonly]
**Notes**| **string** |   | [optional]
**IsDeleted**| **bool** | this field is used to delete an order in \&quot;soft-delete mode\&quot;. This field must be used from get/list endpoint to exclude these orders.  | [optional]
**InsertedAt**| [**time.Time**](time.Time.md) | this field is used to save the original created_at order date. The created_at field is used to filter data.  | [optional] [readonly]
**DeletedAt**| [**time.Time**](time.Time.md) |   | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)


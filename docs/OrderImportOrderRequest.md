# # OrderImportOrderRequest


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** |   | [optional]
**CreatedAt**| [**time.Time**](time.Time.md) |   | [optional]
**Number**| **string** |   | [optional]
**Channel**| **string** |   | [optional]
**Market**| **string** |   | [optional]
**Locale**| **string** |   | [optional]
**CustomerInfo**| [**OrderDataCustomerInfo**](OrderDataCustomerInfo.md) |   | [optional]
**ShippingAddress**| [**OrderPostalAddress**](OrderPostalAddress.md) |   | [optional]
**BillingAddress**| [**OrderPostalAddress**](OrderPostalAddress.md) |   | [optional]
**Payments**| [**[]ImportOrderRequestImportedPayment**](ImportOrderRequestImportedPayment.md) |   | [optional]
**PaymentsInfo**| [**[]OrderDataPaymentInfo**](OrderDataPaymentInfo.md) |   | [optional]
**ShipmentsInfo**| [**[]OrderDataShipmentInfo**](OrderDataShipmentInfo.md) |   | [optional]
**Items**| [**[]OrderOrderDataItem**](OrderOrderDataItem.md) |   | [optional]
**Subtotals**| [**map[string]OrderDataSubtotal**](OrderDataSubtotal.md) |   | [optional]
**Totals**| [**map[string]OrderDataTotal**](OrderDataTotal.md) |   | [optional]
**Status**| **string** |   | [optional]
**Currency**| [**OrderCurrency**](OrderCurrency.md) |  for more information please, see Model/OrderCurrency.php  | [optional] [default to ORDERCURRENCY_XXX]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)


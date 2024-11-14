# # OrderImportOrderRequest


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** |   |
**CreatedAt**| [**time.Time**](time.Time.md) |   | [optional]
**Number**| **string** |   |
**Channel**| **string** |   | [optional]
**Market**| **string** |   |
**Locale**| **string** |   |
**CustomerInfo**| [**OrderDataCustomerInfo**](OrderDataCustomerInfo.md) |   |
**ShippingAddress**| [**OrderPostalAddress**](OrderPostalAddress.md) |   |
**BillingAddress**| [**OrderPostalAddress**](OrderPostalAddress.md) |   |
**Payments**| [**[]ImportOrderRequestImportedPayment**](ImportOrderRequestImportedPayment.md) |   |
**PaymentsInfo**| [**[]OrderDataPaymentInfo**](OrderDataPaymentInfo.md) |   |
**ShipmentsInfo**| [**[]OrderDataShipmentInfo**](OrderDataShipmentInfo.md) |   |
**Items**| [**[]OrderOrderDataItem**](OrderOrderDataItem.md) |   |
**Subtotals**| [**map[string]OrderDataSubtotal**](OrderDataSubtotal.md) |   |
**Totals**| [**map[string]OrderDataTotal**](OrderDataTotal.md) |   |
**Status**| **string** |   |
**Currency**| [**OrderCurrency**](OrderCurrency.md) |  for more information please, see Model/OrderCurrency.php  | [default to XXX]
**VatIncluded**| **bool** |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)


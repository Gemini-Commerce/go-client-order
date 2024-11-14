# # OrderCreateOrderRequest


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** |   |
**Number**| **string** |   |
**Channel**| **string** |   | [optional]
**Market**| **string** |   |
**Locale**| **string** |   |
**Items**| [**[]OrderOrderDataItem**](OrderOrderDataItem.md) |   |
**PaymentsInfo**| [**[]OrderDataPaymentInfo**](OrderDataPaymentInfo.md) |   | [optional]
**ShipmentsInfo**| [**[]OrderDataShipmentInfo**](OrderDataShipmentInfo.md) |   | [optional]
**Promotions**| [**[]OrderDataPromotionInfo**](OrderDataPromotionInfo.md) |   | [optional]
**Payments**| [**[]CreateOrderRequestInitialPayment**](CreateOrderRequestInitialPayment.md) |   | [optional]
**Currency**| [**OrderCurrency**](OrderCurrency.md) |  for more information please, see Model/OrderCurrency.php  | [default to ORDERCURRENCY_XXX]
**Subtotals**| [**map[string]OrderDataSubtotal**](OrderDataSubtotal.md) |   |
**Totals**| [**map[string]OrderDataTotal**](OrderDataTotal.md) |   |
**VatIncluded**| **bool** |   |
**BillingAddress**| [**OrderPostalAddress**](OrderPostalAddress.md) |   |
**ShippingAddress**| [**OrderPostalAddress**](OrderPostalAddress.md) |   |
**CustomerInfo**| [**OrderDataCustomerInfo**](OrderDataCustomerInfo.md) |   |
**CartGrn**| **string** |   | [optional]
**OnHold**| **bool** |   | [optional]
**Notes**| **string** |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)


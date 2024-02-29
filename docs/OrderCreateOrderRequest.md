# # OrderCreateOrderRequest


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** |   | [optional]
**Number**| **string** |   | [optional]
**Channel**| **string** |   | [optional]
**Market**| **string** |   | [optional]
**Locale**| **string** |   | [optional]
**Items**| [**[]OrderOrderDataItem**](OrderOrderDataItem.md) |   | [optional]
**PaymentsInfo**| [**[]OrderDataPaymentInfo**](OrderDataPaymentInfo.md) |   | [optional]
**ShipmentsInfo**| [**[]OrderDataShipmentInfo**](OrderDataShipmentInfo.md) |   | [optional]
**Promotions**| [**[]OrderDataPromotionInfo**](OrderDataPromotionInfo.md) |   | [optional]
**Payments**| [**[]CreateOrderRequestInitialPayment**](CreateOrderRequestInitialPayment.md) |   | [optional]
**Currency**| [**OrderCurrency**](OrderCurrency.md) |  for more information please, see Model/OrderCurrency.php  | [optional] [default to XXX]
**Subtotals**| [**map[string]OrderDataSubtotal**](OrderDataSubtotal.md) |   | [optional]
**Totals**| [**map[string]OrderDataTotal**](OrderDataTotal.md) |   | [optional]
**VatIncluded**| **bool** |   | [optional]
**BillingAddress**| [**OrderPostalAddress**](OrderPostalAddress.md) |   | [optional]
**ShippingAddress**| [**OrderPostalAddress**](OrderPostalAddress.md) |   | [optional]
**CustomerInfo**| [**OrderDataCustomerInfo**](OrderDataCustomerInfo.md) |   | [optional]
**CartGrn**| **string** |   | [optional]
**OnHold**| **bool** |   | [optional]
**Notes**| **string** |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)


# # OrderTransaction


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt**| [**time.Time**](time.Time.md) |   | [optional] [readonly]
**PaymentId**| **string** |   | [optional]
**Id**| **string** |   | [optional]
**Type**| [**OrderTransactionType**](OrderTransactionType.md) |  for more information please, see Model/OrderTransactionType.php  | [optional] [default to ORDERTRANSACTIONTYPE_UNKNOWN]
**AdditionalInfo**| **string** |   | [optional]
**ChildTransactions**| [**[]OrderTransaction**](OrderTransaction.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)


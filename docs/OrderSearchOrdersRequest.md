# # OrderSearchOrdersRequest


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** |   |
**SearchQuery**| **string** |   | [optional]
**PageSize**| **int64** | The maximum number of orders to return. The service may return fewer than this value. If unspecified, at most 10 orders will be returned. The maximum value is 100; values above 100 will be coerced to 100.  | [optional]
**PageToken**| **string** | A page token, received from a previous &#x60;ListOrders&#x60; call. Provide this to retrieve the subsequent page.   When paginating, all other parameters provided to &#x60;ListOrders&#x60; must match the call that provided the page token.  | [optional]
**OrderBy**| [**[]OrderOrderBy**](OrderOrderBy.md) |   | [optional]
**StatusFilter**| [**OrderStatusFilter**](OrderStatusFilter.md) |   | [optional]
**FromDate**| [**time.Time**](time.Time.md) |   | [optional]
**ToDate**| [**time.Time**](time.Time.md) |   | [optional]
**PaymentFilter**| [**OrderPaymentFilter**](OrderPaymentFilter.md) |   | [optional]
**AgentGrn**| **string** |   | [optional]
**UpdatedAtFrom**| [**time.Time**](time.Time.md) |   | [optional]
**UpdatedAtTo**| [**time.Time**](time.Time.md) |   | [optional]
**OnHold**| **bool** |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)


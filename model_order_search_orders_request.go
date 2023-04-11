/*
order/order.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// OrderSearchOrdersRequest struct for OrderSearchOrdersRequest
type OrderSearchOrdersRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	SearchQuery *string `json:"searchQuery,omitempty"`
	// The maximum number of orders to return. The service may return fewer than this value. If unspecified, at most 10 orders will be returned. The maximum value is 100; values above 100 will be coerced to 100.
	PageSize *int64 `json:"pageSize,omitempty"`
	// A page token, received from a previous `ListOrders` call. Provide this to retrieve the subsequent page.   When paginating, all other parameters provided to `ListOrders` must match the call that provided the page token.
	PageToken *string `json:"pageToken,omitempty"`
	OrderBy []OrderOrderBy `json:"orderBy,omitempty"`
	StatusFilter *OrderStatusFilter `json:"statusFilter,omitempty"`
	FromDate *time.Time `json:"fromDate,omitempty"`
	ToDate *time.Time `json:"toDate,omitempty"`
	PaymentFilter *OrderPaymentFilter `json:"paymentFilter,omitempty"`
}

// NewOrderSearchOrdersRequest instantiates a new OrderSearchOrdersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchOrdersRequest() *OrderSearchOrdersRequest {
	this := OrderSearchOrdersRequest{}
	return &this
}

// NewOrderSearchOrdersRequestWithDefaults instantiates a new OrderSearchOrdersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchOrdersRequestWithDefaults() *OrderSearchOrdersRequest {
	this := OrderSearchOrdersRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderSearchOrdersRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetSearchQuery returns the SearchQuery field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetSearchQuery() string {
	if o == nil || isNil(o.SearchQuery) {
		var ret string
		return ret
	}
	return *o.SearchQuery
}

// GetSearchQueryOk returns a tuple with the SearchQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetSearchQueryOk() (*string, bool) {
	if o == nil || isNil(o.SearchQuery) {
    return nil, false
	}
	return o.SearchQuery, true
}

// HasSearchQuery returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) HasSearchQuery() bool {
	if o != nil && !isNil(o.SearchQuery) {
		return true
	}

	return false
}

// SetSearchQuery gets a reference to the given string and assigns it to the SearchQuery field.
func (o *OrderSearchOrdersRequest) SetSearchQuery(v string) {
	o.SearchQuery = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetPageSize() int64 {
	if o == nil || isNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *OrderSearchOrdersRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetPageToken() string {
	if o == nil || isNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.PageToken) {
    return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) HasPageToken() bool {
	if o != nil && !isNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *OrderSearchOrdersRequest) SetPageToken(v string) {
	o.PageToken = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetOrderBy() []OrderOrderBy {
	if o == nil || isNil(o.OrderBy) {
		var ret []OrderOrderBy
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetOrderByOk() ([]OrderOrderBy, bool) {
	if o == nil || isNil(o.OrderBy) {
    return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) HasOrderBy() bool {
	if o != nil && !isNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given []OrderOrderBy and assigns it to the OrderBy field.
func (o *OrderSearchOrdersRequest) SetOrderBy(v []OrderOrderBy) {
	o.OrderBy = v
}

// GetStatusFilter returns the StatusFilter field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetStatusFilter() OrderStatusFilter {
	if o == nil || isNil(o.StatusFilter) {
		var ret OrderStatusFilter
		return ret
	}
	return *o.StatusFilter
}

// GetStatusFilterOk returns a tuple with the StatusFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetStatusFilterOk() (*OrderStatusFilter, bool) {
	if o == nil || isNil(o.StatusFilter) {
    return nil, false
	}
	return o.StatusFilter, true
}

// HasStatusFilter returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) HasStatusFilter() bool {
	if o != nil && !isNil(o.StatusFilter) {
		return true
	}

	return false
}

// SetStatusFilter gets a reference to the given OrderStatusFilter and assigns it to the StatusFilter field.
func (o *OrderSearchOrdersRequest) SetStatusFilter(v OrderStatusFilter) {
	o.StatusFilter = &v
}

// GetFromDate returns the FromDate field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetFromDate() time.Time {
	if o == nil || isNil(o.FromDate) {
		var ret time.Time
		return ret
	}
	return *o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetFromDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.FromDate) {
    return nil, false
	}
	return o.FromDate, true
}

// HasFromDate returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) HasFromDate() bool {
	if o != nil && !isNil(o.FromDate) {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given time.Time and assigns it to the FromDate field.
func (o *OrderSearchOrdersRequest) SetFromDate(v time.Time) {
	o.FromDate = &v
}

// GetToDate returns the ToDate field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetToDate() time.Time {
	if o == nil || isNil(o.ToDate) {
		var ret time.Time
		return ret
	}
	return *o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetToDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.ToDate) {
    return nil, false
	}
	return o.ToDate, true
}

// HasToDate returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) HasToDate() bool {
	if o != nil && !isNil(o.ToDate) {
		return true
	}

	return false
}

// SetToDate gets a reference to the given time.Time and assigns it to the ToDate field.
func (o *OrderSearchOrdersRequest) SetToDate(v time.Time) {
	o.ToDate = &v
}

// GetPaymentFilter returns the PaymentFilter field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetPaymentFilter() OrderPaymentFilter {
	if o == nil || isNil(o.PaymentFilter) {
		var ret OrderPaymentFilter
		return ret
	}
	return *o.PaymentFilter
}

// GetPaymentFilterOk returns a tuple with the PaymentFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetPaymentFilterOk() (*OrderPaymentFilter, bool) {
	if o == nil || isNil(o.PaymentFilter) {
    return nil, false
	}
	return o.PaymentFilter, true
}

// HasPaymentFilter returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) HasPaymentFilter() bool {
	if o != nil && !isNil(o.PaymentFilter) {
		return true
	}

	return false
}

// SetPaymentFilter gets a reference to the given OrderPaymentFilter and assigns it to the PaymentFilter field.
func (o *OrderSearchOrdersRequest) SetPaymentFilter(v OrderPaymentFilter) {
	o.PaymentFilter = &v
}

func (o OrderSearchOrdersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.SearchQuery) {
		toSerialize["searchQuery"] = o.SearchQuery
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !isNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	if !isNil(o.OrderBy) {
		toSerialize["orderBy"] = o.OrderBy
	}
	if !isNil(o.StatusFilter) {
		toSerialize["statusFilter"] = o.StatusFilter
	}
	if !isNil(o.FromDate) {
		toSerialize["fromDate"] = o.FromDate
	}
	if !isNil(o.ToDate) {
		toSerialize["toDate"] = o.ToDate
	}
	if !isNil(o.PaymentFilter) {
		toSerialize["paymentFilter"] = o.PaymentFilter
	}
	return json.Marshal(toSerialize)
}

type NullableOrderSearchOrdersRequest struct {
	value *OrderSearchOrdersRequest
	isSet bool
}

func (v NullableOrderSearchOrdersRequest) Get() *OrderSearchOrdersRequest {
	return v.value
}

func (v *NullableOrderSearchOrdersRequest) Set(val *OrderSearchOrdersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchOrdersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchOrdersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchOrdersRequest(val *OrderSearchOrdersRequest) *NullableOrderSearchOrdersRequest {
	return &NullableOrderSearchOrdersRequest{value: val, isSet: true}
}

func (v NullableOrderSearchOrdersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchOrdersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


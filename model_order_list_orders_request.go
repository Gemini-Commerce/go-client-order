/*
order/order.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OrderListOrdersRequest struct for OrderListOrdersRequest
type OrderListOrdersRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	// The maximum number of orders to return. The service may return fewer than this value. If unspecified, at most 10 orders will be returned. The maximum value is 100; values above 100 will be coerced to 100.
	PageSize *int64 `json:"pageSize,omitempty"`
	// A page token, received from a previous `ListOrders` call. Provide this to retrieve the subsequent page.   When paginating, all other parameters provided to `ListOrders` must match the call that provided the page token.
	PageToken *string `json:"pageToken,omitempty"`
	OrderBy []OrderOrderBy `json:"orderBy,omitempty"`
	StatusFilter *OrderStatusFilter `json:"statusFilter,omitempty"`
}

// NewOrderListOrdersRequest instantiates a new OrderListOrdersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListOrdersRequest() *OrderListOrdersRequest {
	this := OrderListOrdersRequest{}
	return &this
}

// NewOrderListOrdersRequestWithDefaults instantiates a new OrderListOrdersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListOrdersRequestWithDefaults() *OrderListOrdersRequest {
	this := OrderListOrdersRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderListOrdersRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOrdersRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderListOrdersRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderListOrdersRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *OrderListOrdersRequest) GetPageSize() int64 {
	if o == nil || isNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOrdersRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *OrderListOrdersRequest) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *OrderListOrdersRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *OrderListOrdersRequest) GetPageToken() string {
	if o == nil || isNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOrdersRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.PageToken) {
    return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *OrderListOrdersRequest) HasPageToken() bool {
	if o != nil && !isNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *OrderListOrdersRequest) SetPageToken(v string) {
	o.PageToken = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *OrderListOrdersRequest) GetOrderBy() []OrderOrderBy {
	if o == nil || isNil(o.OrderBy) {
		var ret []OrderOrderBy
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOrdersRequest) GetOrderByOk() ([]OrderOrderBy, bool) {
	if o == nil || isNil(o.OrderBy) {
    return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *OrderListOrdersRequest) HasOrderBy() bool {
	if o != nil && !isNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given []OrderOrderBy and assigns it to the OrderBy field.
func (o *OrderListOrdersRequest) SetOrderBy(v []OrderOrderBy) {
	o.OrderBy = v
}

// GetStatusFilter returns the StatusFilter field value if set, zero value otherwise.
func (o *OrderListOrdersRequest) GetStatusFilter() OrderStatusFilter {
	if o == nil || isNil(o.StatusFilter) {
		var ret OrderStatusFilter
		return ret
	}
	return *o.StatusFilter
}

// GetStatusFilterOk returns a tuple with the StatusFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOrdersRequest) GetStatusFilterOk() (*OrderStatusFilter, bool) {
	if o == nil || isNil(o.StatusFilter) {
    return nil, false
	}
	return o.StatusFilter, true
}

// HasStatusFilter returns a boolean if a field has been set.
func (o *OrderListOrdersRequest) HasStatusFilter() bool {
	if o != nil && !isNil(o.StatusFilter) {
		return true
	}

	return false
}

// SetStatusFilter gets a reference to the given OrderStatusFilter and assigns it to the StatusFilter field.
func (o *OrderListOrdersRequest) SetStatusFilter(v OrderStatusFilter) {
	o.StatusFilter = &v
}

func (o OrderListOrdersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
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
	return json.Marshal(toSerialize)
}

type NullableOrderListOrdersRequest struct {
	value *OrderListOrdersRequest
	isSet bool
}

func (v NullableOrderListOrdersRequest) Get() *OrderListOrdersRequest {
	return v.value
}

func (v *NullableOrderListOrdersRequest) Set(val *OrderListOrdersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListOrdersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListOrdersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListOrdersRequest(val *OrderListOrdersRequest) *NullableOrderListOrdersRequest {
	return &NullableOrderListOrdersRequest{value: val, isSet: true}
}

func (v NullableOrderListOrdersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListOrdersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// OrderListFulfillmentsRequest struct for OrderListFulfillmentsRequest
type OrderListFulfillmentsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	OrderId *string `json:"orderId,omitempty"`
}

// NewOrderListFulfillmentsRequest instantiates a new OrderListFulfillmentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListFulfillmentsRequest() *OrderListFulfillmentsRequest {
	this := OrderListFulfillmentsRequest{}
	return &this
}

// NewOrderListFulfillmentsRequestWithDefaults instantiates a new OrderListFulfillmentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListFulfillmentsRequestWithDefaults() *OrderListFulfillmentsRequest {
	this := OrderListFulfillmentsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderListFulfillmentsRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListFulfillmentsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderListFulfillmentsRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderListFulfillmentsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderListFulfillmentsRequest) GetOrderId() string {
	if o == nil || isNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListFulfillmentsRequest) GetOrderIdOk() (*string, bool) {
	if o == nil || isNil(o.OrderId) {
    return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderListFulfillmentsRequest) HasOrderId() bool {
	if o != nil && !isNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderListFulfillmentsRequest) SetOrderId(v string) {
	o.OrderId = &v
}

func (o OrderListFulfillmentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	return json.Marshal(toSerialize)
}

type NullableOrderListFulfillmentsRequest struct {
	value *OrderListFulfillmentsRequest
	isSet bool
}

func (v NullableOrderListFulfillmentsRequest) Get() *OrderListFulfillmentsRequest {
	return v.value
}

func (v *NullableOrderListFulfillmentsRequest) Set(val *OrderListFulfillmentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListFulfillmentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListFulfillmentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListFulfillmentsRequest(val *OrderListFulfillmentsRequest) *NullableOrderListFulfillmentsRequest {
	return &NullableOrderListFulfillmentsRequest{value: val, isSet: true}
}

func (v NullableOrderListFulfillmentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListFulfillmentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


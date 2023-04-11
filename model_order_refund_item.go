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

// OrderRefundItem struct for OrderRefundItem
type OrderRefundItem struct {
	OrderItemId *string `json:"orderItemId,omitempty"`
	Qty *int64 `json:"qty,omitempty"`
}

// NewOrderRefundItem instantiates a new OrderRefundItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderRefundItem() *OrderRefundItem {
	this := OrderRefundItem{}
	return &this
}

// NewOrderRefundItemWithDefaults instantiates a new OrderRefundItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderRefundItemWithDefaults() *OrderRefundItem {
	this := OrderRefundItem{}
	return &this
}

// GetOrderItemId returns the OrderItemId field value if set, zero value otherwise.
func (o *OrderRefundItem) GetOrderItemId() string {
	if o == nil || isNil(o.OrderItemId) {
		var ret string
		return ret
	}
	return *o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefundItem) GetOrderItemIdOk() (*string, bool) {
	if o == nil || isNil(o.OrderItemId) {
    return nil, false
	}
	return o.OrderItemId, true
}

// HasOrderItemId returns a boolean if a field has been set.
func (o *OrderRefundItem) HasOrderItemId() bool {
	if o != nil && !isNil(o.OrderItemId) {
		return true
	}

	return false
}

// SetOrderItemId gets a reference to the given string and assigns it to the OrderItemId field.
func (o *OrderRefundItem) SetOrderItemId(v string) {
	o.OrderItemId = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *OrderRefundItem) GetQty() int64 {
	if o == nil || isNil(o.Qty) {
		var ret int64
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefundItem) GetQtyOk() (*int64, bool) {
	if o == nil || isNil(o.Qty) {
    return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *OrderRefundItem) HasQty() bool {
	if o != nil && !isNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given int64 and assigns it to the Qty field.
func (o *OrderRefundItem) SetQty(v int64) {
	o.Qty = &v
}

func (o OrderRefundItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OrderItemId) {
		toSerialize["orderItemId"] = o.OrderItemId
	}
	if !isNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	return json.Marshal(toSerialize)
}

type NullableOrderRefundItem struct {
	value *OrderRefundItem
	isSet bool
}

func (v NullableOrderRefundItem) Get() *OrderRefundItem {
	return v.value
}

func (v *NullableOrderRefundItem) Set(val *OrderRefundItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderRefundItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderRefundItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderRefundItem(val *OrderRefundItem) *NullableOrderRefundItem {
	return &NullableOrderRefundItem{value: val, isSet: true}
}

func (v NullableOrderRefundItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderRefundItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


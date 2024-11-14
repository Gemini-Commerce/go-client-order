/*
order Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OrderShipmentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderShipmentItem{}

// OrderShipmentItem struct for OrderShipmentItem
type OrderShipmentItem struct {
	OrderItemId *string `json:"orderItemId,omitempty"`
	Qty *int64 `json:"qty,omitempty"`
	RowWeight *string `json:"rowWeight,omitempty"`
}

// NewOrderShipmentItem instantiates a new OrderShipmentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderShipmentItem() *OrderShipmentItem {
	this := OrderShipmentItem{}
	return &this
}

// NewOrderShipmentItemWithDefaults instantiates a new OrderShipmentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderShipmentItemWithDefaults() *OrderShipmentItem {
	this := OrderShipmentItem{}
	return &this
}

// GetOrderItemId returns the OrderItemId field value if set, zero value otherwise.
func (o *OrderShipmentItem) GetOrderItemId() string {
	if o == nil || IsNil(o.OrderItemId) {
		var ret string
		return ret
	}
	return *o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentItem) GetOrderItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderItemId) {
		return nil, false
	}
	return o.OrderItemId, true
}

// OrderItemId returns a boolean if a field has been set.
func (o *OrderShipmentItem) OrderItemId() bool {
	if o != nil && !IsNil(o.OrderItemId) {
		return true
	}

	return false
}

// SetOrderItemId gets a reference to the given string and assigns it to the OrderItemId field.
func (o *OrderShipmentItem) SetOrderItemId(v string) {
	o.OrderItemId = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *OrderShipmentItem) GetQty() int64 {
	if o == nil || IsNil(o.Qty) {
		var ret int64
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentItem) GetQtyOk() (*int64, bool) {
	if o == nil || IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// Qty returns a boolean if a field has been set.
func (o *OrderShipmentItem) Qty() bool {
	if o != nil && !IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given int64 and assigns it to the Qty field.
func (o *OrderShipmentItem) SetQty(v int64) {
	o.Qty = &v
}

// GetRowWeight returns the RowWeight field value if set, zero value otherwise.
func (o *OrderShipmentItem) GetRowWeight() string {
	if o == nil || IsNil(o.RowWeight) {
		var ret string
		return ret
	}
	return *o.RowWeight
}

// GetRowWeightOk returns a tuple with the RowWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentItem) GetRowWeightOk() (*string, bool) {
	if o == nil || IsNil(o.RowWeight) {
		return nil, false
	}
	return o.RowWeight, true
}

// RowWeight returns a boolean if a field has been set.
func (o *OrderShipmentItem) RowWeight() bool {
	if o != nil && !IsNil(o.RowWeight) {
		return true
	}

	return false
}

// SetRowWeight gets a reference to the given string and assigns it to the RowWeight field.
func (o *OrderShipmentItem) SetRowWeight(v string) {
	o.RowWeight = &v
}

func (o OrderShipmentItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderShipmentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderItemId) {
		toSerialize["orderItemId"] = o.OrderItemId
	}
	if !IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !IsNil(o.RowWeight) {
		toSerialize["rowWeight"] = o.RowWeight
	}
	return toSerialize, nil
}

type NullableOrderShipmentItem struct {
	value *OrderShipmentItem
	isSet bool
}

func (v NullableOrderShipmentItem) Get() *OrderShipmentItem {
	return v.value
}

func (v *NullableOrderShipmentItem) Set(val *OrderShipmentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderShipmentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderShipmentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderShipmentItem(val *OrderShipmentItem) *NullableOrderShipmentItem {
	return &NullableOrderShipmentItem{value: val, isSet: true}
}

func (v NullableOrderShipmentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderShipmentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



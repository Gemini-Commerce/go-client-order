/*
order Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package order

import (
	"encoding/json"
)

// checks if the OrderFulfillmentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderFulfillmentItem{}

// OrderFulfillmentItem struct for OrderFulfillmentItem
type OrderFulfillmentItem struct {
	OrderItemId *string `json:"orderItemId,omitempty"`
	Qty *int64 `json:"qty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderFulfillmentItem OrderFulfillmentItem

// NewOrderFulfillmentItem instantiates a new OrderFulfillmentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderFulfillmentItem() *OrderFulfillmentItem {
	this := OrderFulfillmentItem{}
	return &this
}

// NewOrderFulfillmentItemWithDefaults instantiates a new OrderFulfillmentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderFulfillmentItemWithDefaults() *OrderFulfillmentItem {
	this := OrderFulfillmentItem{}
	return &this
}

// GetOrderItemId returns the OrderItemId field value if set, zero value otherwise.
func (o *OrderFulfillmentItem) GetOrderItemId() string {
	if o == nil || IsNil(o.OrderItemId) {
		var ret string
		return ret
	}
	return *o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFulfillmentItem) GetOrderItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderItemId) {
		return nil, false
	}
	return o.OrderItemId, true
}

// &#39;Has&#39;OrderItemId returns a boolean if a field has been set.
func (o *OrderFulfillmentItem) &#39;Has&#39;OrderItemId() bool {
	if o != nil && !IsNil(o.OrderItemId) {
		return true
	}

	return false
}

// SetOrderItemId gets a reference to the given string and assigns it to the OrderItemId field.
func (o *OrderFulfillmentItem) SetOrderItemId(v string) {
	o.OrderItemId = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *OrderFulfillmentItem) GetQty() int64 {
	if o == nil || IsNil(o.Qty) {
		var ret int64
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFulfillmentItem) GetQtyOk() (*int64, bool) {
	if o == nil || IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// &#39;Has&#39;Qty returns a boolean if a field has been set.
func (o *OrderFulfillmentItem) &#39;Has&#39;Qty() bool {
	if o != nil && !IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given int64 and assigns it to the Qty field.
func (o *OrderFulfillmentItem) SetQty(v int64) {
	o.Qty = &v
}

func (o OrderFulfillmentItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderFulfillmentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderItemId) {
		toSerialize["orderItemId"] = o.OrderItemId
	}
	if !IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderFulfillmentItem) UnmarshalJSON(data []byte) (err error) {
	varOrderFulfillmentItem := _OrderFulfillmentItem{}

	err = json.Unmarshal(data, &varOrderFulfillmentItem)

	if err != nil {
		return err
	}

	*o = OrderFulfillmentItem(varOrderFulfillmentItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderItemId")
		delete(additionalProperties, "qty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderFulfillmentItem) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderFulfillmentItem) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderFulfillmentItem struct {
	value *OrderFulfillmentItem
	isSet bool
}

func (v NullableOrderFulfillmentItem) Get() *OrderFulfillmentItem {
	return v.value
}

func (v *NullableOrderFulfillmentItem) Set(val *OrderFulfillmentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderFulfillmentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderFulfillmentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderFulfillmentItem(val *OrderFulfillmentItem) *NullableOrderFulfillmentItem {
	return &NullableOrderFulfillmentItem{value: val, isSet: true}
}

func (v NullableOrderFulfillmentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderFulfillmentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



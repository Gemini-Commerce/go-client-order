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

// OrderRefundAmount struct for OrderRefundAmount
type OrderRefundAmount struct {
	Code *OrderRefundAmountCode `json:"code,omitempty"`
	Value *OrderMoney `json:"value,omitempty"`
}

// NewOrderRefundAmount instantiates a new OrderRefundAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderRefundAmount() *OrderRefundAmount {
	this := OrderRefundAmount{}
	var code OrderRefundAmountCode = ORDERREFUNDAMOUNTCODE_UNKNOWN
	this.Code = &code
	return &this
}

// NewOrderRefundAmountWithDefaults instantiates a new OrderRefundAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderRefundAmountWithDefaults() *OrderRefundAmount {
	this := OrderRefundAmount{}
	var code OrderRefundAmountCode = ORDERREFUNDAMOUNTCODE_UNKNOWN
	this.Code = &code
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OrderRefundAmount) GetCode() OrderRefundAmountCode {
	if o == nil || isNil(o.Code) {
		var ret OrderRefundAmountCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefundAmount) GetCodeOk() (*OrderRefundAmountCode, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OrderRefundAmount) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given OrderRefundAmountCode and assigns it to the Code field.
func (o *OrderRefundAmount) SetCode(v OrderRefundAmountCode) {
	o.Code = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OrderRefundAmount) GetValue() OrderMoney {
	if o == nil || isNil(o.Value) {
		var ret OrderMoney
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefundAmount) GetValueOk() (*OrderMoney, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OrderRefundAmount) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given OrderMoney and assigns it to the Value field.
func (o *OrderRefundAmount) SetValue(v OrderMoney) {
	o.Value = &v
}

func (o OrderRefundAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableOrderRefundAmount struct {
	value *OrderRefundAmount
	isSet bool
}

func (v NullableOrderRefundAmount) Get() *OrderRefundAmount {
	return v.value
}

func (v *NullableOrderRefundAmount) Set(val *OrderRefundAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderRefundAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderRefundAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderRefundAmount(val *OrderRefundAmount) *NullableOrderRefundAmount {
	return &NullableOrderRefundAmount{value: val, isSet: true}
}

func (v NullableOrderRefundAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderRefundAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


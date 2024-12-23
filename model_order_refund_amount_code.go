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
	"fmt"
)

// OrderRefundAmountCode the model 'OrderRefundAmountCode'
type OrderRefundAmountCode string

// List of orderRefundAmountCode
const (
	ORDERREFUNDAMOUNTCODE_UNKNOWN  OrderRefundAmountCode = "UNKNOWN"
	ORDERREFUNDAMOUNTCODE_ORDERED  OrderRefundAmountCode = "ORDERED"
	ORDERREFUNDAMOUNTCODE_SHIPPING OrderRefundAmountCode = "SHIPPING"
)

// All allowed values of OrderRefundAmountCode enum
var AllowedOrderRefundAmountCodeEnumValues = []OrderRefundAmountCode{
	"UNKNOWN",
	"ORDERED",
	"SHIPPING",
}

func (v *OrderRefundAmountCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderRefundAmountCode(value)
	for _, existing := range AllowedOrderRefundAmountCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderRefundAmountCode", value)
}

// NewOrderRefundAmountCodeFromValue returns a pointer to a valid OrderRefundAmountCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderRefundAmountCodeFromValue(v string) (*OrderRefundAmountCode, error) {
	ev := OrderRefundAmountCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderRefundAmountCode: valid values are %v", v, AllowedOrderRefundAmountCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderRefundAmountCode) IsValid() bool {
	for _, existing := range AllowedOrderRefundAmountCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderRefundAmountCode value
func (v OrderRefundAmountCode) Ptr() *OrderRefundAmountCode {
	return &v
}

type NullableOrderRefundAmountCode struct {
	value *OrderRefundAmountCode
	isSet bool
}

func (v NullableOrderRefundAmountCode) Get() *OrderRefundAmountCode {
	return v.value
}

func (v *NullableOrderRefundAmountCode) Set(val *OrderRefundAmountCode) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderRefundAmountCode) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderRefundAmountCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderRefundAmountCode(val *OrderRefundAmountCode) *NullableOrderRefundAmountCode {
	return &NullableOrderRefundAmountCode{value: val, isSet: true}
}

func (v NullableOrderRefundAmountCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderRefundAmountCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

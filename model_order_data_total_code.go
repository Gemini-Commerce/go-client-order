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

// OrderDataTotalCode the model 'OrderDataTotalCode'
type OrderDataTotalCode string

// List of OrderDataTotalCode
const (
	ORDERDATATOTALCODE_UNKNOWN  OrderDataTotalCode = "UNKNOWN"
	ORDERDATATOTALCODE_ORDERED  OrderDataTotalCode = "ORDERED"
	ORDERDATATOTALCODE_PAID     OrderDataTotalCode = "PAID"
	ORDERDATATOTALCODE_REFUNDED OrderDataTotalCode = "REFUNDED"
)

// All allowed values of OrderDataTotalCode enum
var AllowedOrderDataTotalCodeEnumValues = []OrderDataTotalCode{
	"UNKNOWN",
	"ORDERED",
	"PAID",
	"REFUNDED",
}

func (v *OrderDataTotalCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderDataTotalCode(value)
	for _, existing := range AllowedOrderDataTotalCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderDataTotalCode", value)
}

// NewOrderDataTotalCodeFromValue returns a pointer to a valid OrderDataTotalCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderDataTotalCodeFromValue(v string) (*OrderDataTotalCode, error) {
	ev := OrderDataTotalCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderDataTotalCode: valid values are %v", v, AllowedOrderDataTotalCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderDataTotalCode) IsValid() bool {
	for _, existing := range AllowedOrderDataTotalCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderDataTotalCode value
func (v OrderDataTotalCode) Ptr() *OrderDataTotalCode {
	return &v
}

type NullableOrderDataTotalCode struct {
	value *OrderDataTotalCode
	isSet bool
}

func (v NullableOrderDataTotalCode) Get() *OrderDataTotalCode {
	return v.value
}

func (v *NullableOrderDataTotalCode) Set(val *OrderDataTotalCode) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataTotalCode) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataTotalCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataTotalCode(val *OrderDataTotalCode) *NullableOrderDataTotalCode {
	return &NullableOrderDataTotalCode{value: val, isSet: true}
}

func (v NullableOrderDataTotalCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataTotalCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

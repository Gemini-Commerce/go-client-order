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

// OrderTransactionType the model 'OrderTransactionType'
type OrderTransactionType string

// List of orderTransactionType
const (
	ORDERTRANSACTIONTYPE_UNKNOWN       OrderTransactionType = "UNKNOWN"
	ORDERTRANSACTIONTYPE_AUTHORIZATION OrderTransactionType = "AUTHORIZATION"
	ORDERTRANSACTIONTYPE_CAPTURE       OrderTransactionType = "CAPTURE"
	ORDERTRANSACTIONTYPE_SALE          OrderTransactionType = "SALE"
	ORDERTRANSACTIONTYPE_REFUND        OrderTransactionType = "REFUND"
	ORDERTRANSACTIONTYPE_VOID          OrderTransactionType = "VOID"
	ORDERTRANSACTIONTYPE_FAILED        OrderTransactionType = "FAILED"
	ORDERTRANSACTIONTYPE_PENDING       OrderTransactionType = "PENDING"
	ORDERTRANSACTIONTYPE_FRAUD         OrderTransactionType = "FRAUD"
	ORDERTRANSACTIONTYPE_NOOP          OrderTransactionType = "NOOP"
)

// All allowed values of OrderTransactionType enum
var AllowedOrderTransactionTypeEnumValues = []OrderTransactionType{
	"UNKNOWN",
	"AUTHORIZATION",
	"CAPTURE",
	"SALE",
	"REFUND",
	"VOID",
	"FAILED",
	"PENDING",
	"FRAUD",
	"NOOP",
}

func (v *OrderTransactionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderTransactionType(value)
	for _, existing := range AllowedOrderTransactionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderTransactionType", value)
}

// NewOrderTransactionTypeFromValue returns a pointer to a valid OrderTransactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderTransactionTypeFromValue(v string) (*OrderTransactionType, error) {
	ev := OrderTransactionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderTransactionType: valid values are %v", v, AllowedOrderTransactionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderTransactionType) IsValid() bool {
	for _, existing := range AllowedOrderTransactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderTransactionType value
func (v OrderTransactionType) Ptr() *OrderTransactionType {
	return &v
}

type NullableOrderTransactionType struct {
	value *OrderTransactionType
	isSet bool
}

func (v NullableOrderTransactionType) Get() *OrderTransactionType {
	return v.value
}

func (v *NullableOrderTransactionType) Set(val *OrderTransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTransactionType(val *OrderTransactionType) *NullableOrderTransactionType {
	return &NullableOrderTransactionType{value: val, isSet: true}
}

func (v NullableOrderTransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

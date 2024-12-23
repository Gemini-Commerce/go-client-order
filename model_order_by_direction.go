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

// OrderByDirection the model 'OrderByDirection'
type OrderByDirection string

// List of OrderByDirection
const (
	ORDERBYDIRECTION_DEFAULT OrderByDirection = "DEFAULT"
	ORDERBYDIRECTION_ASC     OrderByDirection = "ASC"
	ORDERBYDIRECTION_DESC    OrderByDirection = "DESC"
)

// All allowed values of OrderByDirection enum
var AllowedOrderByDirectionEnumValues = []OrderByDirection{
	"DEFAULT",
	"ASC",
	"DESC",
}

func (v *OrderByDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderByDirection(value)
	for _, existing := range AllowedOrderByDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderByDirection", value)
}

// NewOrderByDirectionFromValue returns a pointer to a valid OrderByDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderByDirectionFromValue(v string) (*OrderByDirection, error) {
	ev := OrderByDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderByDirection: valid values are %v", v, AllowedOrderByDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderByDirection) IsValid() bool {
	for _, existing := range AllowedOrderByDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderByDirection value
func (v OrderByDirection) Ptr() *OrderByDirection {
	return &v
}

type NullableOrderByDirection struct {
	value *OrderByDirection
	isSet bool
}

func (v NullableOrderByDirection) Get() *OrderByDirection {
	return v.value
}

func (v *NullableOrderByDirection) Set(val *OrderByDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderByDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderByDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderByDirection(val *OrderByDirection) *NullableOrderByDirection {
	return &NullableOrderByDirection{value: val, isSet: true}
}

func (v NullableOrderByDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderByDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

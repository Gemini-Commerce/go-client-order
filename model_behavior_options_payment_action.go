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
	"fmt"
)

// BehaviorOptionsPaymentAction the model 'BehaviorOptionsPaymentAction'
type BehaviorOptionsPaymentAction string

// List of BehaviorOptionsPaymentAction
const (
	UNKNOWN BehaviorOptionsPaymentAction = "UNKNOWN"
	DO_NOT_HANDLE BehaviorOptionsPaymentAction = "DO_NOT_HANDLE"
)

// All allowed values of BehaviorOptionsPaymentAction enum
var AllowedBehaviorOptionsPaymentActionEnumValues = []BehaviorOptionsPaymentAction{
	"UNKNOWN",
	"DO_NOT_HANDLE",
}

func (v *BehaviorOptionsPaymentAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BehaviorOptionsPaymentAction(value)
	for _, existing := range AllowedBehaviorOptionsPaymentActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BehaviorOptionsPaymentAction", value)
}

// NewBehaviorOptionsPaymentActionFromValue returns a pointer to a valid BehaviorOptionsPaymentAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBehaviorOptionsPaymentActionFromValue(v string) (*BehaviorOptionsPaymentAction, error) {
	ev := BehaviorOptionsPaymentAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BehaviorOptionsPaymentAction: valid values are %v", v, AllowedBehaviorOptionsPaymentActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BehaviorOptionsPaymentAction) IsValid() bool {
	for _, existing := range AllowedBehaviorOptionsPaymentActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BehaviorOptionsPaymentAction value
func (v BehaviorOptionsPaymentAction) Ptr() *BehaviorOptionsPaymentAction {
	return &v
}

type NullableBehaviorOptionsPaymentAction struct {
	value *BehaviorOptionsPaymentAction
	isSet bool
}

func (v NullableBehaviorOptionsPaymentAction) Get() *BehaviorOptionsPaymentAction {
	return v.value
}

func (v *NullableBehaviorOptionsPaymentAction) Set(val *BehaviorOptionsPaymentAction) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorOptionsPaymentAction) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorOptionsPaymentAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorOptionsPaymentAction(val *BehaviorOptionsPaymentAction) *NullableBehaviorOptionsPaymentAction {
	return &NullableBehaviorOptionsPaymentAction{value: val, isSet: true}
}

func (v NullableBehaviorOptionsPaymentAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorOptionsPaymentAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


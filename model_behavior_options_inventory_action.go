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

// BehaviorOptionsInventoryAction the model 'BehaviorOptionsInventoryAction'
type BehaviorOptionsInventoryAction string

// List of BehaviorOptionsInventoryAction
const (
	BEHAVIOROPTIONSINVENTORYACTION_UNKNOWN BehaviorOptionsInventoryAction = "UNKNOWN"
	BEHAVIOROPTIONSINVENTORYACTION_DO_NOT_HANDLE BehaviorOptionsInventoryAction = "DO_NOT_HANDLE"
)

// All allowed values of BehaviorOptionsInventoryAction enum
var AllowedBehaviorOptionsInventoryActionEnumValues = []BehaviorOptionsInventoryAction{
	"UNKNOWN",
	"DO_NOT_HANDLE",
}

func (v *BehaviorOptionsInventoryAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BehaviorOptionsInventoryAction(value)
	for _, existing := range AllowedBehaviorOptionsInventoryActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BehaviorOptionsInventoryAction", value)
}

// NewBehaviorOptionsInventoryActionFromValue returns a pointer to a valid BehaviorOptionsInventoryAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBehaviorOptionsInventoryActionFromValue(v string) (*BehaviorOptionsInventoryAction, error) {
	ev := BehaviorOptionsInventoryAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BehaviorOptionsInventoryAction: valid values are %v", v, AllowedBehaviorOptionsInventoryActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BehaviorOptionsInventoryAction) IsValid() bool {
	for _, existing := range AllowedBehaviorOptionsInventoryActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BehaviorOptionsInventoryAction value
func (v BehaviorOptionsInventoryAction) Ptr() *BehaviorOptionsInventoryAction {
	return &v
}

type NullableBehaviorOptionsInventoryAction struct {
	value *BehaviorOptionsInventoryAction
	isSet bool
}

func (v NullableBehaviorOptionsInventoryAction) Get() *BehaviorOptionsInventoryAction {
	return v.value
}

func (v *NullableBehaviorOptionsInventoryAction) Set(val *BehaviorOptionsInventoryAction) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorOptionsInventoryAction) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorOptionsInventoryAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorOptionsInventoryAction(val *BehaviorOptionsInventoryAction) *NullableBehaviorOptionsInventoryAction {
	return &NullableBehaviorOptionsInventoryAction{value: val, isSet: true}
}

func (v NullableBehaviorOptionsInventoryAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorOptionsInventoryAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


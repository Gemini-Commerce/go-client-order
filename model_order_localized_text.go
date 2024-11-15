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

// checks if the OrderLocalizedText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderLocalizedText{}

// OrderLocalizedText struct for OrderLocalizedText
type OrderLocalizedText struct {
	Value *map[string]string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderLocalizedText OrderLocalizedText

// NewOrderLocalizedText instantiates a new OrderLocalizedText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderLocalizedText() *OrderLocalizedText {
	this := OrderLocalizedText{}
	return &this
}

// NewOrderLocalizedTextWithDefaults instantiates a new OrderLocalizedText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderLocalizedTextWithDefaults() *OrderLocalizedText {
	this := OrderLocalizedText{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OrderLocalizedText) GetValue() map[string]string {
	if o == nil || IsNil(o.Value) {
		var ret map[string]string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLocalizedText) GetValueOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OrderLocalizedText) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]string and assigns it to the Value field.
func (o *OrderLocalizedText) SetValue(v map[string]string) {
	o.Value = &v
}

func (o OrderLocalizedText) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderLocalizedText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderLocalizedText) UnmarshalJSON(data []byte) (err error) {
	varOrderLocalizedText := _OrderLocalizedText{}

	err = json.Unmarshal(data, &varOrderLocalizedText)

	if err != nil {
		return err
	}

	*o = OrderLocalizedText(varOrderLocalizedText)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderLocalizedText) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderLocalizedText) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderLocalizedText struct {
	value *OrderLocalizedText
	isSet bool
}

func (v NullableOrderLocalizedText) Get() *OrderLocalizedText {
	return v.value
}

func (v *NullableOrderLocalizedText) Set(val *OrderLocalizedText) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderLocalizedText) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderLocalizedText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderLocalizedText(val *OrderLocalizedText) *NullableOrderLocalizedText {
	return &NullableOrderLocalizedText{value: val, isSet: true}
}

func (v NullableOrderLocalizedText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderLocalizedText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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
	"bytes"
	"fmt"
)

// checks if the OrderDataSubtotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDataSubtotal{}

// OrderDataSubtotal struct for OrderDataSubtotal
type OrderDataSubtotal struct {
	Code OrderDataSubtotalCode `json:"code"`
	Value OrderMoney `json:"value"`
}

type _OrderDataSubtotal OrderDataSubtotal

// NewOrderDataSubtotal instantiates a new OrderDataSubtotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataSubtotal(code OrderDataSubtotalCode, value OrderMoney) *OrderDataSubtotal {
	this := OrderDataSubtotal{}
	this.Code = code
	this.Value = value
	return &this
}

// NewOrderDataSubtotalWithDefaults instantiates a new OrderDataSubtotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataSubtotalWithDefaults() *OrderDataSubtotal {
	this := OrderDataSubtotal{}
	var code OrderDataSubtotalCode = UNKNOWN
	this.Code = code
	return &this
}

// GetCode returns the Code field value
func (o *OrderDataSubtotal) GetCode() OrderDataSubtotalCode {
	if o == nil {
		var ret OrderDataSubtotalCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OrderDataSubtotal) GetCodeOk() (*OrderDataSubtotalCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OrderDataSubtotal) SetCode(v OrderDataSubtotalCode) {
	o.Code = v
}

// GetValue returns the Value field value
func (o *OrderDataSubtotal) GetValue() OrderMoney {
	if o == nil {
		var ret OrderMoney
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OrderDataSubtotal) GetValueOk() (*OrderMoney, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *OrderDataSubtotal) SetValue(v OrderMoney) {
	o.Value = v
}

func (o OrderDataSubtotal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDataSubtotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *OrderDataSubtotal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderDataSubtotal := _OrderDataSubtotal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderDataSubtotal)

	if err != nil {
		return err
	}

	*o = OrderDataSubtotal(varOrderDataSubtotal)

	return err
}

type NullableOrderDataSubtotal struct {
	value *OrderDataSubtotal
	isSet bool
}

func (v NullableOrderDataSubtotal) Get() *OrderDataSubtotal {
	return v.value
}

func (v *NullableOrderDataSubtotal) Set(val *OrderDataSubtotal) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataSubtotal) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataSubtotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataSubtotal(val *OrderDataSubtotal) *NullableOrderDataSubtotal {
	return &NullableOrderDataSubtotal{value: val, isSet: true}
}

func (v NullableOrderDataSubtotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataSubtotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



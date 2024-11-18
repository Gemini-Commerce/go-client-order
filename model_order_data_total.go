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

// checks if the OrderDataTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDataTotal{}

// OrderDataTotal struct for OrderDataTotal
type OrderDataTotal struct {
	Code                 OrderDataTotalCode `json:"code"`
	Value                OrderMoney         `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _OrderDataTotal OrderDataTotal

// NewOrderDataTotal instantiates a new OrderDataTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataTotal(code OrderDataTotalCode, value OrderMoney) *OrderDataTotal {
	this := OrderDataTotal{}
	this.Code = code
	this.Value = value
	return &this
}

// NewOrderDataTotalWithDefaults instantiates a new OrderDataTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataTotalWithDefaults() *OrderDataTotal {
	this := OrderDataTotal{}
	var code OrderDataTotalCode = ORDERDATATOTALCODE_UNKNOWN
	this.Code = code
	return &this
}

// GetCode returns the Code field value
func (o *OrderDataTotal) GetCode() OrderDataTotalCode {
	if o == nil {
		var ret OrderDataTotalCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OrderDataTotal) GetCodeOk() (*OrderDataTotalCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OrderDataTotal) SetCode(v OrderDataTotalCode) {
	o.Code = v
}

// GetValue returns the Value field value
func (o *OrderDataTotal) GetValue() OrderMoney {
	if o == nil {
		var ret OrderMoney
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OrderDataTotal) GetValueOk() (*OrderMoney, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *OrderDataTotal) SetValue(v OrderMoney) {
	o.Value = v
}

func (o OrderDataTotal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDataTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderDataTotal) UnmarshalJSON(data []byte) (err error) {
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
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderDataTotal := _OrderDataTotal{}

	err = json.Unmarshal(data, &varOrderDataTotal)

	if err != nil {
		return err
	}

	*o = OrderDataTotal(varOrderDataTotal)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderDataTotal struct {
	value *OrderDataTotal
	isSet bool
}

func (v NullableOrderDataTotal) Get() *OrderDataTotal {
	return v.value
}

func (v *NullableOrderDataTotal) Set(val *OrderDataTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataTotal(val *OrderDataTotal) *NullableOrderDataTotal {
	return &NullableOrderDataTotal{value: val, isSet: true}
}

func (v NullableOrderDataTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

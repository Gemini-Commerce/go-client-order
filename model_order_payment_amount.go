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

// checks if the OrderPaymentAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderPaymentAmount{}

// OrderPaymentAmount struct for OrderPaymentAmount
type OrderPaymentAmount struct {
	Code                 *OrderPaymentAmountCode `json:"code,omitempty"`
	Value                *OrderMoney             `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderPaymentAmount OrderPaymentAmount

// NewOrderPaymentAmount instantiates a new OrderPaymentAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderPaymentAmount() *OrderPaymentAmount {
	this := OrderPaymentAmount{}
	var code OrderPaymentAmountCode = ORDERPAYMENTAMOUNTCODE_UNKNOWN
	this.Code = &code
	return &this
}

// NewOrderPaymentAmountWithDefaults instantiates a new OrderPaymentAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderPaymentAmountWithDefaults() *OrderPaymentAmount {
	this := OrderPaymentAmount{}
	var code OrderPaymentAmountCode = ORDERPAYMENTAMOUNTCODE_UNKNOWN
	this.Code = &code
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OrderPaymentAmount) GetCode() OrderPaymentAmountCode {
	if o == nil || IsNil(o.Code) {
		var ret OrderPaymentAmountCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPaymentAmount) GetCodeOk() (*OrderPaymentAmountCode, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OrderPaymentAmount) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given OrderPaymentAmountCode and assigns it to the Code field.
func (o *OrderPaymentAmount) SetCode(v OrderPaymentAmountCode) {
	o.Code = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OrderPaymentAmount) GetValue() OrderMoney {
	if o == nil || IsNil(o.Value) {
		var ret OrderMoney
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPaymentAmount) GetValueOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OrderPaymentAmount) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given OrderMoney and assigns it to the Value field.
func (o *OrderPaymentAmount) SetValue(v OrderMoney) {
	o.Value = &v
}

func (o OrderPaymentAmount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderPaymentAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderPaymentAmount) UnmarshalJSON(data []byte) (err error) {
	varOrderPaymentAmount := _OrderPaymentAmount{}

	err = json.Unmarshal(data, &varOrderPaymentAmount)

	if err != nil {
		return err
	}

	*o = OrderPaymentAmount(varOrderPaymentAmount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderPaymentAmount struct {
	value *OrderPaymentAmount
	isSet bool
}

func (v NullableOrderPaymentAmount) Get() *OrderPaymentAmount {
	return v.value
}

func (v *NullableOrderPaymentAmount) Set(val *OrderPaymentAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderPaymentAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderPaymentAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderPaymentAmount(val *OrderPaymentAmount) *NullableOrderPaymentAmount {
	return &NullableOrderPaymentAmount{value: val, isSet: true}
}

func (v NullableOrderPaymentAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderPaymentAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the OrderPaymentFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderPaymentFilter{}

// OrderPaymentFilter struct for OrderPaymentFilter
type OrderPaymentFilter struct {
	Codes                []string                     `json:"codes,omitempty"`
	Condition            *OrderPaymentFilterCondition `json:"condition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderPaymentFilter OrderPaymentFilter

// NewOrderPaymentFilter instantiates a new OrderPaymentFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderPaymentFilter() *OrderPaymentFilter {
	this := OrderPaymentFilter{}
	var condition OrderPaymentFilterCondition = ORDERPAYMENTFILTERCONDITION_IN
	this.Condition = &condition
	return &this
}

// NewOrderPaymentFilterWithDefaults instantiates a new OrderPaymentFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderPaymentFilterWithDefaults() *OrderPaymentFilter {
	this := OrderPaymentFilter{}
	var condition OrderPaymentFilterCondition = ORDERPAYMENTFILTERCONDITION_IN
	this.Condition = &condition
	return &this
}

// GetCodes returns the Codes field value if set, zero value otherwise.
func (o *OrderPaymentFilter) GetCodes() []string {
	if o == nil || IsNil(o.Codes) {
		var ret []string
		return ret
	}
	return o.Codes
}

// GetCodesOk returns a tuple with the Codes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPaymentFilter) GetCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.Codes) {
		return nil, false
	}
	return o.Codes, true
}

// HasCodes returns a boolean if a field has been set.
func (o *OrderPaymentFilter) IsSetCodes() bool {
	if o != nil && !IsNil(o.Codes) {
		return true
	}

	return false
}

// SetCodes gets a reference to the given []string and assigns it to the Codes field.
func (o *OrderPaymentFilter) SetCodes(v []string) {
	o.Codes = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *OrderPaymentFilter) GetCondition() OrderPaymentFilterCondition {
	if o == nil || IsNil(o.Condition) {
		var ret OrderPaymentFilterCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPaymentFilter) GetConditionOk() (*OrderPaymentFilterCondition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *OrderPaymentFilter) IsSetCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given OrderPaymentFilterCondition and assigns it to the Condition field.
func (o *OrderPaymentFilter) SetCondition(v OrderPaymentFilterCondition) {
	o.Condition = &v
}

func (o OrderPaymentFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderPaymentFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Codes) {
		toSerialize["codes"] = o.Codes
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderPaymentFilter) UnmarshalJSON(data []byte) (err error) {
	varOrderPaymentFilter := _OrderPaymentFilter{}

	err = json.Unmarshal(data, &varOrderPaymentFilter)

	if err != nil {
		return err
	}

	*o = OrderPaymentFilter(varOrderPaymentFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "codes")
		delete(additionalProperties, "condition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderPaymentFilter) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderPaymentFilter) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderPaymentFilter struct {
	value *OrderPaymentFilter
	isSet bool
}

func (v NullableOrderPaymentFilter) Get() *OrderPaymentFilter {
	return v.value
}

func (v *NullableOrderPaymentFilter) Set(val *OrderPaymentFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderPaymentFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderPaymentFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderPaymentFilter(val *OrderPaymentFilter) *NullableOrderPaymentFilter {
	return &NullableOrderPaymentFilter{value: val, isSet: true}
}

func (v NullableOrderPaymentFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderPaymentFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

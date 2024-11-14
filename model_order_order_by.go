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

// checks if the OrderOrderBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderOrderBy{}

// OrderOrderBy struct for OrderOrderBy
type OrderOrderBy struct {
	Field *string `json:"field,omitempty"`
	Direction *OrderByDirection `json:"direction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderOrderBy OrderOrderBy

// NewOrderOrderBy instantiates a new OrderOrderBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderOrderBy() *OrderOrderBy {
	this := OrderOrderBy{}
	var direction OrderByDirection = ORDERBYDIRECTION_DEFAULT
	this.Direction = &direction
	return &this
}

// NewOrderOrderByWithDefaults instantiates a new OrderOrderBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderOrderByWithDefaults() *OrderOrderBy {
	this := OrderOrderBy{}
	var direction OrderByDirection = ORDERBYDIRECTION_DEFAULT
	this.Direction = &direction
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *OrderOrderBy) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderBy) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// IsSetField returns a boolean if a field has been set.
func (o *OrderOrderBy) IsSetField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *OrderOrderBy) SetField(v string) {
	o.Field = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *OrderOrderBy) GetDirection() OrderByDirection {
	if o == nil || IsNil(o.Direction) {
		var ret OrderByDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderBy) GetDirectionOk() (*OrderByDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// IsSetDirection returns a boolean if a field has been set.
func (o *OrderOrderBy) IsSetDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given OrderByDirection and assigns it to the Direction field.
func (o *OrderOrderBy) SetDirection(v OrderByDirection) {
	o.Direction = &v
}

func (o OrderOrderBy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderOrderBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderOrderBy) UnmarshalJSON(data []byte) (err error) {
	varOrderOrderBy := _OrderOrderBy{}

	err = json.Unmarshal(data, &varOrderOrderBy)

	if err != nil {
		return err
	}

	*o = OrderOrderBy(varOrderOrderBy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "field")
		delete(additionalProperties, "direction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderOrderBy) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderOrderBy) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderOrderBy struct {
	value *OrderOrderBy
	isSet bool
}

func (v NullableOrderOrderBy) Get() *OrderOrderBy {
	return v.value
}

func (v *NullableOrderOrderBy) Set(val *OrderOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderOrderBy(val *OrderOrderBy) *NullableOrderOrderBy {
	return &NullableOrderOrderBy{value: val, isSet: true}
}

func (v NullableOrderOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



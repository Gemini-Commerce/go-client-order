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

// checks if the OrderGetOrderByCartIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderGetOrderByCartIdRequest{}

// OrderGetOrderByCartIdRequest struct for OrderGetOrderByCartIdRequest
type OrderGetOrderByCartIdRequest struct {
	TenantId string `json:"tenantId"`
	CartId string `json:"cartId"`
}

type _OrderGetOrderByCartIdRequest OrderGetOrderByCartIdRequest

// NewOrderGetOrderByCartIdRequest instantiates a new OrderGetOrderByCartIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderGetOrderByCartIdRequest(tenantId string, cartId string) *OrderGetOrderByCartIdRequest {
	this := OrderGetOrderByCartIdRequest{}
	this.TenantId = tenantId
	this.CartId = cartId
	return &this
}

// NewOrderGetOrderByCartIdRequestWithDefaults instantiates a new OrderGetOrderByCartIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderGetOrderByCartIdRequestWithDefaults() *OrderGetOrderByCartIdRequest {
	this := OrderGetOrderByCartIdRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderGetOrderByCartIdRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderGetOrderByCartIdRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderGetOrderByCartIdRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetCartId returns the CartId field value
func (o *OrderGetOrderByCartIdRequest) GetCartId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CartId
}

// GetCartIdOk returns a tuple with the CartId field value
// and a boolean to check if the value has been set.
func (o *OrderGetOrderByCartIdRequest) GetCartIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CartId, true
}

// SetCartId sets field value
func (o *OrderGetOrderByCartIdRequest) SetCartId(v string) {
	o.CartId = v
}

func (o OrderGetOrderByCartIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderGetOrderByCartIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["cartId"] = o.CartId
	return toSerialize, nil
}

func (o *OrderGetOrderByCartIdRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"cartId",
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

	varOrderGetOrderByCartIdRequest := _OrderGetOrderByCartIdRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderGetOrderByCartIdRequest)

	if err != nil {
		return err
	}

	*o = OrderGetOrderByCartIdRequest(varOrderGetOrderByCartIdRequest)

	return err
}

type NullableOrderGetOrderByCartIdRequest struct {
	value *OrderGetOrderByCartIdRequest
	isSet bool
}

func (v NullableOrderGetOrderByCartIdRequest) Get() *OrderGetOrderByCartIdRequest {
	return v.value
}

func (v *NullableOrderGetOrderByCartIdRequest) Set(val *OrderGetOrderByCartIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderGetOrderByCartIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderGetOrderByCartIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderGetOrderByCartIdRequest(val *OrderGetOrderByCartIdRequest) *NullableOrderGetOrderByCartIdRequest {
	return &NullableOrderGetOrderByCartIdRequest{value: val, isSet: true}
}

func (v NullableOrderGetOrderByCartIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderGetOrderByCartIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



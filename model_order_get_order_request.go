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
	"bytes"
	"fmt"
)

// checks if the OrderGetOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderGetOrderRequest{}

// OrderGetOrderRequest struct for OrderGetOrderRequest
type OrderGetOrderRequest struct {
	TenantId string `json:"tenantId"`
	Id string `json:"id"`
}

type _OrderGetOrderRequest OrderGetOrderRequest

// NewOrderGetOrderRequest instantiates a new OrderGetOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderGetOrderRequest(tenantId string, id string) *OrderGetOrderRequest {
	this := OrderGetOrderRequest{}
	this.TenantId = tenantId
	this.Id = id
	return &this
}

// NewOrderGetOrderRequestWithDefaults instantiates a new OrderGetOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderGetOrderRequestWithDefaults() *OrderGetOrderRequest {
	this := OrderGetOrderRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderGetOrderRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderGetOrderRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderGetOrderRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *OrderGetOrderRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderGetOrderRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderGetOrderRequest) SetId(v string) {
	o.Id = v
}

func (o OrderGetOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderGetOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *OrderGetOrderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
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

	varOrderGetOrderRequest := _OrderGetOrderRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderGetOrderRequest)

	if err != nil {
		return err
	}

	*o = OrderGetOrderRequest(varOrderGetOrderRequest)

	return err
}

type NullableOrderGetOrderRequest struct {
	value *OrderGetOrderRequest
	isSet bool
}

func (v NullableOrderGetOrderRequest) Get() *OrderGetOrderRequest {
	return v.value
}

func (v *NullableOrderGetOrderRequest) Set(val *OrderGetOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderGetOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderGetOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderGetOrderRequest(val *OrderGetOrderRequest) *NullableOrderGetOrderRequest {
	return &NullableOrderGetOrderRequest{value: val, isSet: true}
}

func (v NullableOrderGetOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderGetOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



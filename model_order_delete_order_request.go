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

// checks if the OrderDeleteOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDeleteOrderRequest{}

// OrderDeleteOrderRequest struct for OrderDeleteOrderRequest
type OrderDeleteOrderRequest struct {
	TenantId string `json:"tenantId"`
	OrderId string `json:"orderId"`
}

type _OrderDeleteOrderRequest OrderDeleteOrderRequest

// NewOrderDeleteOrderRequest instantiates a new OrderDeleteOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDeleteOrderRequest(tenantId string, orderId string) *OrderDeleteOrderRequest {
	this := OrderDeleteOrderRequest{}
	this.TenantId = tenantId
	this.OrderId = orderId
	return &this
}

// NewOrderDeleteOrderRequestWithDefaults instantiates a new OrderDeleteOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDeleteOrderRequestWithDefaults() *OrderDeleteOrderRequest {
	this := OrderDeleteOrderRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderDeleteOrderRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderDeleteOrderRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderDeleteOrderRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetOrderId returns the OrderId field value
func (o *OrderDeleteOrderRequest) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *OrderDeleteOrderRequest) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *OrderDeleteOrderRequest) SetOrderId(v string) {
	o.OrderId = v
}

func (o OrderDeleteOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDeleteOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["orderId"] = o.OrderId
	return toSerialize, nil
}

func (o *OrderDeleteOrderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"orderId",
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

	varOrderDeleteOrderRequest := _OrderDeleteOrderRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderDeleteOrderRequest)

	if err != nil {
		return err
	}

	*o = OrderDeleteOrderRequest(varOrderDeleteOrderRequest)

	return err
}

type NullableOrderDeleteOrderRequest struct {
	value *OrderDeleteOrderRequest
	isSet bool
}

func (v NullableOrderDeleteOrderRequest) Get() *OrderDeleteOrderRequest {
	return v.value
}

func (v *NullableOrderDeleteOrderRequest) Set(val *OrderDeleteOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDeleteOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDeleteOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDeleteOrderRequest(val *OrderDeleteOrderRequest) *NullableOrderDeleteOrderRequest {
	return &NullableOrderDeleteOrderRequest{value: val, isSet: true}
}

func (v NullableOrderDeleteOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDeleteOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



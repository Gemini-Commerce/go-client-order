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

// checks if the OrderUnholdOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderUnholdOrderRequest{}

// OrderUnholdOrderRequest struct for OrderUnholdOrderRequest
type OrderUnholdOrderRequest struct {
	TenantId string `json:"tenantId"`
	OrderId string `json:"orderId"`
}

type _OrderUnholdOrderRequest OrderUnholdOrderRequest

// NewOrderUnholdOrderRequest instantiates a new OrderUnholdOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderUnholdOrderRequest(tenantId string, orderId string) *OrderUnholdOrderRequest {
	this := OrderUnholdOrderRequest{}
	this.TenantId = tenantId
	this.OrderId = orderId
	return &this
}

// NewOrderUnholdOrderRequestWithDefaults instantiates a new OrderUnholdOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderUnholdOrderRequestWithDefaults() *OrderUnholdOrderRequest {
	this := OrderUnholdOrderRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderUnholdOrderRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderUnholdOrderRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderUnholdOrderRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetOrderId returns the OrderId field value
func (o *OrderUnholdOrderRequest) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *OrderUnholdOrderRequest) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *OrderUnholdOrderRequest) SetOrderId(v string) {
	o.OrderId = v
}

func (o OrderUnholdOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderUnholdOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["orderId"] = o.OrderId
	return toSerialize, nil
}

func (o *OrderUnholdOrderRequest) UnmarshalJSON(data []byte) (err error) {
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

	varOrderUnholdOrderRequest := _OrderUnholdOrderRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderUnholdOrderRequest)

	if err != nil {
		return err
	}

	*o = OrderUnholdOrderRequest(varOrderUnholdOrderRequest)

	return err
}

type NullableOrderUnholdOrderRequest struct {
	value *OrderUnholdOrderRequest
	isSet bool
}

func (v NullableOrderUnholdOrderRequest) Get() *OrderUnholdOrderRequest {
	return v.value
}

func (v *NullableOrderUnholdOrderRequest) Set(val *OrderUnholdOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderUnholdOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderUnholdOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderUnholdOrderRequest(val *OrderUnholdOrderRequest) *NullableOrderUnholdOrderRequest {
	return &NullableOrderUnholdOrderRequest{value: val, isSet: true}
}

func (v NullableOrderUnholdOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderUnholdOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
order/order.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OrderUnholdOrderRequest struct for OrderUnholdOrderRequest
type OrderUnholdOrderRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	OrderId *string `json:"orderId,omitempty"`
}

// NewOrderUnholdOrderRequest instantiates a new OrderUnholdOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderUnholdOrderRequest() *OrderUnholdOrderRequest {
	this := OrderUnholdOrderRequest{}
	return &this
}

// NewOrderUnholdOrderRequestWithDefaults instantiates a new OrderUnholdOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderUnholdOrderRequestWithDefaults() *OrderUnholdOrderRequest {
	this := OrderUnholdOrderRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderUnholdOrderRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUnholdOrderRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderUnholdOrderRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderUnholdOrderRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderUnholdOrderRequest) GetOrderId() string {
	if o == nil || isNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUnholdOrderRequest) GetOrderIdOk() (*string, bool) {
	if o == nil || isNil(o.OrderId) {
    return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderUnholdOrderRequest) HasOrderId() bool {
	if o != nil && !isNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderUnholdOrderRequest) SetOrderId(v string) {
	o.OrderId = &v
}

func (o OrderUnholdOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	return json.Marshal(toSerialize)
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



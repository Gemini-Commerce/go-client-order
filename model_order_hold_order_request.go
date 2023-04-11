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

// OrderHoldOrderRequest struct for OrderHoldOrderRequest
type OrderHoldOrderRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	OrderId *string `json:"orderId,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// NewOrderHoldOrderRequest instantiates a new OrderHoldOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderHoldOrderRequest() *OrderHoldOrderRequest {
	this := OrderHoldOrderRequest{}
	return &this
}

// NewOrderHoldOrderRequestWithDefaults instantiates a new OrderHoldOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderHoldOrderRequestWithDefaults() *OrderHoldOrderRequest {
	this := OrderHoldOrderRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderHoldOrderRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderHoldOrderRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderHoldOrderRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderHoldOrderRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderHoldOrderRequest) GetOrderId() string {
	if o == nil || isNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderHoldOrderRequest) GetOrderIdOk() (*string, bool) {
	if o == nil || isNil(o.OrderId) {
    return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderHoldOrderRequest) HasOrderId() bool {
	if o != nil && !isNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderHoldOrderRequest) SetOrderId(v string) {
	o.OrderId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *OrderHoldOrderRequest) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderHoldOrderRequest) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *OrderHoldOrderRequest) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *OrderHoldOrderRequest) SetReason(v string) {
	o.Reason = &v
}

func (o OrderHoldOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableOrderHoldOrderRequest struct {
	value *OrderHoldOrderRequest
	isSet bool
}

func (v NullableOrderHoldOrderRequest) Get() *OrderHoldOrderRequest {
	return v.value
}

func (v *NullableOrderHoldOrderRequest) Set(val *OrderHoldOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderHoldOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderHoldOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderHoldOrderRequest(val *OrderHoldOrderRequest) *NullableOrderHoldOrderRequest {
	return &NullableOrderHoldOrderRequest{value: val, isSet: true}
}

func (v NullableOrderHoldOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderHoldOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



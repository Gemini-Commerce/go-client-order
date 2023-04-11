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

// OrderCancelShipmentRequest struct for OrderCancelShipmentRequest
type OrderCancelShipmentRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	ShipmentId *string `json:"shipmentId,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// NewOrderCancelShipmentRequest instantiates a new OrderCancelShipmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCancelShipmentRequest() *OrderCancelShipmentRequest {
	this := OrderCancelShipmentRequest{}
	return &this
}

// NewOrderCancelShipmentRequestWithDefaults instantiates a new OrderCancelShipmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCancelShipmentRequestWithDefaults() *OrderCancelShipmentRequest {
	this := OrderCancelShipmentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderCancelShipmentRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelShipmentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderCancelShipmentRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderCancelShipmentRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetShipmentId returns the ShipmentId field value if set, zero value otherwise.
func (o *OrderCancelShipmentRequest) GetShipmentId() string {
	if o == nil || isNil(o.ShipmentId) {
		var ret string
		return ret
	}
	return *o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelShipmentRequest) GetShipmentIdOk() (*string, bool) {
	if o == nil || isNil(o.ShipmentId) {
    return nil, false
	}
	return o.ShipmentId, true
}

// HasShipmentId returns a boolean if a field has been set.
func (o *OrderCancelShipmentRequest) HasShipmentId() bool {
	if o != nil && !isNil(o.ShipmentId) {
		return true
	}

	return false
}

// SetShipmentId gets a reference to the given string and assigns it to the ShipmentId field.
func (o *OrderCancelShipmentRequest) SetShipmentId(v string) {
	o.ShipmentId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *OrderCancelShipmentRequest) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelShipmentRequest) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *OrderCancelShipmentRequest) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *OrderCancelShipmentRequest) SetReason(v string) {
	o.Reason = &v
}

func (o OrderCancelShipmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.ShipmentId) {
		toSerialize["shipmentId"] = o.ShipmentId
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableOrderCancelShipmentRequest struct {
	value *OrderCancelShipmentRequest
	isSet bool
}

func (v NullableOrderCancelShipmentRequest) Get() *OrderCancelShipmentRequest {
	return v.value
}

func (v *NullableOrderCancelShipmentRequest) Set(val *OrderCancelShipmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelShipmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelShipmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelShipmentRequest(val *OrderCancelShipmentRequest) *NullableOrderCancelShipmentRequest {
	return &NullableOrderCancelShipmentRequest{value: val, isSet: true}
}

func (v NullableOrderCancelShipmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelShipmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



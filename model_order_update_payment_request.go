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

// OrderUpdatePaymentRequest struct for OrderUpdatePaymentRequest
type OrderUpdatePaymentRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	PaymentId *string `json:"paymentId,omitempty"`
	CcInfo *PaymentCcInfo `json:"ccInfo,omitempty"`
}

// NewOrderUpdatePaymentRequest instantiates a new OrderUpdatePaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderUpdatePaymentRequest() *OrderUpdatePaymentRequest {
	this := OrderUpdatePaymentRequest{}
	return &this
}

// NewOrderUpdatePaymentRequestWithDefaults instantiates a new OrderUpdatePaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderUpdatePaymentRequestWithDefaults() *OrderUpdatePaymentRequest {
	this := OrderUpdatePaymentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderUpdatePaymentRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdatePaymentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderUpdatePaymentRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderUpdatePaymentRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *OrderUpdatePaymentRequest) GetPaymentId() string {
	if o == nil || isNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdatePaymentRequest) GetPaymentIdOk() (*string, bool) {
	if o == nil || isNil(o.PaymentId) {
    return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *OrderUpdatePaymentRequest) HasPaymentId() bool {
	if o != nil && !isNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *OrderUpdatePaymentRequest) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetCcInfo returns the CcInfo field value if set, zero value otherwise.
func (o *OrderUpdatePaymentRequest) GetCcInfo() PaymentCcInfo {
	if o == nil || isNil(o.CcInfo) {
		var ret PaymentCcInfo
		return ret
	}
	return *o.CcInfo
}

// GetCcInfoOk returns a tuple with the CcInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdatePaymentRequest) GetCcInfoOk() (*PaymentCcInfo, bool) {
	if o == nil || isNil(o.CcInfo) {
    return nil, false
	}
	return o.CcInfo, true
}

// HasCcInfo returns a boolean if a field has been set.
func (o *OrderUpdatePaymentRequest) HasCcInfo() bool {
	if o != nil && !isNil(o.CcInfo) {
		return true
	}

	return false
}

// SetCcInfo gets a reference to the given PaymentCcInfo and assigns it to the CcInfo field.
func (o *OrderUpdatePaymentRequest) SetCcInfo(v PaymentCcInfo) {
	o.CcInfo = &v
}

func (o OrderUpdatePaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !isNil(o.CcInfo) {
		toSerialize["ccInfo"] = o.CcInfo
	}
	return json.Marshal(toSerialize)
}

type NullableOrderUpdatePaymentRequest struct {
	value *OrderUpdatePaymentRequest
	isSet bool
}

func (v NullableOrderUpdatePaymentRequest) Get() *OrderUpdatePaymentRequest {
	return v.value
}

func (v *NullableOrderUpdatePaymentRequest) Set(val *OrderUpdatePaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderUpdatePaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderUpdatePaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderUpdatePaymentRequest(val *OrderUpdatePaymentRequest) *NullableOrderUpdatePaymentRequest {
	return &NullableOrderUpdatePaymentRequest{value: val, isSet: true}
}

func (v NullableOrderUpdatePaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderUpdatePaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


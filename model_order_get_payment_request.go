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

// OrderGetPaymentRequest struct for OrderGetPaymentRequest
type OrderGetPaymentRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewOrderGetPaymentRequest instantiates a new OrderGetPaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderGetPaymentRequest() *OrderGetPaymentRequest {
	this := OrderGetPaymentRequest{}
	return &this
}

// NewOrderGetPaymentRequestWithDefaults instantiates a new OrderGetPaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderGetPaymentRequestWithDefaults() *OrderGetPaymentRequest {
	this := OrderGetPaymentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderGetPaymentRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetPaymentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderGetPaymentRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderGetPaymentRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderGetPaymentRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetPaymentRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderGetPaymentRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderGetPaymentRequest) SetId(v string) {
	o.Id = &v
}

func (o OrderGetPaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrderGetPaymentRequest struct {
	value *OrderGetPaymentRequest
	isSet bool
}

func (v NullableOrderGetPaymentRequest) Get() *OrderGetPaymentRequest {
	return v.value
}

func (v *NullableOrderGetPaymentRequest) Set(val *OrderGetPaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderGetPaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderGetPaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderGetPaymentRequest(val *OrderGetPaymentRequest) *NullableOrderGetPaymentRequest {
	return &NullableOrderGetPaymentRequest{value: val, isSet: true}
}

func (v NullableOrderGetPaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderGetPaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


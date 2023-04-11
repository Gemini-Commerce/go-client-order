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

// OrderReportFulfillmentErrorRequest struct for OrderReportFulfillmentErrorRequest
type OrderReportFulfillmentErrorRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	FulfillmentId *string `json:"fulfillmentId,omitempty"`
}

// NewOrderReportFulfillmentErrorRequest instantiates a new OrderReportFulfillmentErrorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderReportFulfillmentErrorRequest() *OrderReportFulfillmentErrorRequest {
	this := OrderReportFulfillmentErrorRequest{}
	return &this
}

// NewOrderReportFulfillmentErrorRequestWithDefaults instantiates a new OrderReportFulfillmentErrorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderReportFulfillmentErrorRequestWithDefaults() *OrderReportFulfillmentErrorRequest {
	this := OrderReportFulfillmentErrorRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderReportFulfillmentErrorRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportFulfillmentErrorRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderReportFulfillmentErrorRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderReportFulfillmentErrorRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetFulfillmentId returns the FulfillmentId field value if set, zero value otherwise.
func (o *OrderReportFulfillmentErrorRequest) GetFulfillmentId() string {
	if o == nil || isNil(o.FulfillmentId) {
		var ret string
		return ret
	}
	return *o.FulfillmentId
}

// GetFulfillmentIdOk returns a tuple with the FulfillmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportFulfillmentErrorRequest) GetFulfillmentIdOk() (*string, bool) {
	if o == nil || isNil(o.FulfillmentId) {
    return nil, false
	}
	return o.FulfillmentId, true
}

// HasFulfillmentId returns a boolean if a field has been set.
func (o *OrderReportFulfillmentErrorRequest) HasFulfillmentId() bool {
	if o != nil && !isNil(o.FulfillmentId) {
		return true
	}

	return false
}

// SetFulfillmentId gets a reference to the given string and assigns it to the FulfillmentId field.
func (o *OrderReportFulfillmentErrorRequest) SetFulfillmentId(v string) {
	o.FulfillmentId = &v
}

func (o OrderReportFulfillmentErrorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.FulfillmentId) {
		toSerialize["fulfillmentId"] = o.FulfillmentId
	}
	return json.Marshal(toSerialize)
}

type NullableOrderReportFulfillmentErrorRequest struct {
	value *OrderReportFulfillmentErrorRequest
	isSet bool
}

func (v NullableOrderReportFulfillmentErrorRequest) Get() *OrderReportFulfillmentErrorRequest {
	return v.value
}

func (v *NullableOrderReportFulfillmentErrorRequest) Set(val *OrderReportFulfillmentErrorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderReportFulfillmentErrorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderReportFulfillmentErrorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderReportFulfillmentErrorRequest(val *OrderReportFulfillmentErrorRequest) *NullableOrderReportFulfillmentErrorRequest {
	return &NullableOrderReportFulfillmentErrorRequest{value: val, isSet: true}
}

func (v NullableOrderReportFulfillmentErrorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderReportFulfillmentErrorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



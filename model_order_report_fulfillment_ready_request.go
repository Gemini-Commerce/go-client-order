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

// OrderReportFulfillmentReadyRequest struct for OrderReportFulfillmentReadyRequest
type OrderReportFulfillmentReadyRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	FulfillmentId *string `json:"fulfillmentId,omitempty"`
}

// NewOrderReportFulfillmentReadyRequest instantiates a new OrderReportFulfillmentReadyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderReportFulfillmentReadyRequest() *OrderReportFulfillmentReadyRequest {
	this := OrderReportFulfillmentReadyRequest{}
	return &this
}

// NewOrderReportFulfillmentReadyRequestWithDefaults instantiates a new OrderReportFulfillmentReadyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderReportFulfillmentReadyRequestWithDefaults() *OrderReportFulfillmentReadyRequest {
	this := OrderReportFulfillmentReadyRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderReportFulfillmentReadyRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportFulfillmentReadyRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderReportFulfillmentReadyRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderReportFulfillmentReadyRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetFulfillmentId returns the FulfillmentId field value if set, zero value otherwise.
func (o *OrderReportFulfillmentReadyRequest) GetFulfillmentId() string {
	if o == nil || isNil(o.FulfillmentId) {
		var ret string
		return ret
	}
	return *o.FulfillmentId
}

// GetFulfillmentIdOk returns a tuple with the FulfillmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportFulfillmentReadyRequest) GetFulfillmentIdOk() (*string, bool) {
	if o == nil || isNil(o.FulfillmentId) {
    return nil, false
	}
	return o.FulfillmentId, true
}

// HasFulfillmentId returns a boolean if a field has been set.
func (o *OrderReportFulfillmentReadyRequest) HasFulfillmentId() bool {
	if o != nil && !isNil(o.FulfillmentId) {
		return true
	}

	return false
}

// SetFulfillmentId gets a reference to the given string and assigns it to the FulfillmentId field.
func (o *OrderReportFulfillmentReadyRequest) SetFulfillmentId(v string) {
	o.FulfillmentId = &v
}

func (o OrderReportFulfillmentReadyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.FulfillmentId) {
		toSerialize["fulfillmentId"] = o.FulfillmentId
	}
	return json.Marshal(toSerialize)
}

type NullableOrderReportFulfillmentReadyRequest struct {
	value *OrderReportFulfillmentReadyRequest
	isSet bool
}

func (v NullableOrderReportFulfillmentReadyRequest) Get() *OrderReportFulfillmentReadyRequest {
	return v.value
}

func (v *NullableOrderReportFulfillmentReadyRequest) Set(val *OrderReportFulfillmentReadyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderReportFulfillmentReadyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderReportFulfillmentReadyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderReportFulfillmentReadyRequest(val *OrderReportFulfillmentReadyRequest) *NullableOrderReportFulfillmentReadyRequest {
	return &NullableOrderReportFulfillmentReadyRequest{value: val, isSet: true}
}

func (v NullableOrderReportFulfillmentReadyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderReportFulfillmentReadyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



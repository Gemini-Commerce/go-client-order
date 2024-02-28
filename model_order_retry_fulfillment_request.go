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
)

// checks if the OrderRetryFulfillmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderRetryFulfillmentRequest{}

// OrderRetryFulfillmentRequest struct for OrderRetryFulfillmentRequest
type OrderRetryFulfillmentRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	FulfillmentId *string `json:"fulfillmentId,omitempty"`
}

// NewOrderRetryFulfillmentRequest instantiates a new OrderRetryFulfillmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderRetryFulfillmentRequest() *OrderRetryFulfillmentRequest {
	this := OrderRetryFulfillmentRequest{}
	return &this
}

// NewOrderRetryFulfillmentRequestWithDefaults instantiates a new OrderRetryFulfillmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderRetryFulfillmentRequestWithDefaults() *OrderRetryFulfillmentRequest {
	this := OrderRetryFulfillmentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderRetryFulfillmentRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRetryFulfillmentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderRetryFulfillmentRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderRetryFulfillmentRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetFulfillmentId returns the FulfillmentId field value if set, zero value otherwise.
func (o *OrderRetryFulfillmentRequest) GetFulfillmentId() string {
	if o == nil || IsNil(o.FulfillmentId) {
		var ret string
		return ret
	}
	return *o.FulfillmentId
}

// GetFulfillmentIdOk returns a tuple with the FulfillmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRetryFulfillmentRequest) GetFulfillmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentId) {
		return nil, false
	}
	return o.FulfillmentId, true
}

// HasFulfillmentId returns a boolean if a field has been set.
func (o *OrderRetryFulfillmentRequest) HasFulfillmentId() bool {
	if o != nil && !IsNil(o.FulfillmentId) {
		return true
	}

	return false
}

// SetFulfillmentId gets a reference to the given string and assigns it to the FulfillmentId field.
func (o *OrderRetryFulfillmentRequest) SetFulfillmentId(v string) {
	o.FulfillmentId = &v
}

func (o OrderRetryFulfillmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderRetryFulfillmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.FulfillmentId) {
		toSerialize["fulfillmentId"] = o.FulfillmentId
	}
	return toSerialize, nil
}

type NullableOrderRetryFulfillmentRequest struct {
	value *OrderRetryFulfillmentRequest
	isSet bool
}

func (v NullableOrderRetryFulfillmentRequest) Get() *OrderRetryFulfillmentRequest {
	return v.value
}

func (v *NullableOrderRetryFulfillmentRequest) Set(val *OrderRetryFulfillmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderRetryFulfillmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderRetryFulfillmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderRetryFulfillmentRequest(val *OrderRetryFulfillmentRequest) *NullableOrderRetryFulfillmentRequest {
	return &NullableOrderRetryFulfillmentRequest{value: val, isSet: true}
}

func (v NullableOrderRetryFulfillmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderRetryFulfillmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



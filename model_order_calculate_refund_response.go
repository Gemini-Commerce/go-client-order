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
)

// checks if the OrderCalculateRefundResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCalculateRefundResponse{}

// OrderCalculateRefundResponse struct for OrderCalculateRefundResponse
type OrderCalculateRefundResponse struct {
	CreateRefundRequest *OrderCreateRefundRequest `json:"createRefundRequest,omitempty"`
}

// NewOrderCalculateRefundResponse instantiates a new OrderCalculateRefundResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCalculateRefundResponse() *OrderCalculateRefundResponse {
	this := OrderCalculateRefundResponse{}
	return &this
}

// NewOrderCalculateRefundResponseWithDefaults instantiates a new OrderCalculateRefundResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCalculateRefundResponseWithDefaults() *OrderCalculateRefundResponse {
	this := OrderCalculateRefundResponse{}
	return &this
}

// GetCreateRefundRequest returns the CreateRefundRequest field value if set, zero value otherwise.
func (o *OrderCalculateRefundResponse) GetCreateRefundRequest() OrderCreateRefundRequest {
	if o == nil || IsNil(o.CreateRefundRequest) {
		var ret OrderCreateRefundRequest
		return ret
	}
	return *o.CreateRefundRequest
}

// GetCreateRefundRequestOk returns a tuple with the CreateRefundRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCalculateRefundResponse) GetCreateRefundRequestOk() (*OrderCreateRefundRequest, bool) {
	if o == nil || IsNil(o.CreateRefundRequest) {
		return nil, false
	}
	return o.CreateRefundRequest, true
}

// CreateRefundRequest returns a boolean if a field has been set.
func (o *OrderCalculateRefundResponse) CreateRefundRequest() bool {
	if o != nil && !IsNil(o.CreateRefundRequest) {
		return true
	}

	return false
}

// SetCreateRefundRequest gets a reference to the given OrderCreateRefundRequest and assigns it to the CreateRefundRequest field.
func (o *OrderCalculateRefundResponse) SetCreateRefundRequest(v OrderCreateRefundRequest) {
	o.CreateRefundRequest = &v
}

func (o OrderCalculateRefundResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCalculateRefundResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateRefundRequest) {
		toSerialize["createRefundRequest"] = o.CreateRefundRequest
	}
	return toSerialize, nil
}

type NullableOrderCalculateRefundResponse struct {
	value *OrderCalculateRefundResponse
	isSet bool
}

func (v NullableOrderCalculateRefundResponse) Get() *OrderCalculateRefundResponse {
	return v.value
}

func (v *NullableOrderCalculateRefundResponse) Set(val *OrderCalculateRefundResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCalculateRefundResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCalculateRefundResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCalculateRefundResponse(val *OrderCalculateRefundResponse) *NullableOrderCalculateRefundResponse {
	return &NullableOrderCalculateRefundResponse{value: val, isSet: true}
}

func (v NullableOrderCalculateRefundResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCalculateRefundResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



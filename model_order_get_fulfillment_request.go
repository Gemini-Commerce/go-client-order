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

// checks if the OrderGetFulfillmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderGetFulfillmentRequest{}

// OrderGetFulfillmentRequest struct for OrderGetFulfillmentRequest
type OrderGetFulfillmentRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewOrderGetFulfillmentRequest instantiates a new OrderGetFulfillmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderGetFulfillmentRequest() *OrderGetFulfillmentRequest {
	this := OrderGetFulfillmentRequest{}
	return &this
}

// NewOrderGetFulfillmentRequestWithDefaults instantiates a new OrderGetFulfillmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderGetFulfillmentRequestWithDefaults() *OrderGetFulfillmentRequest {
	this := OrderGetFulfillmentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderGetFulfillmentRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetFulfillmentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderGetFulfillmentRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderGetFulfillmentRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderGetFulfillmentRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetFulfillmentRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderGetFulfillmentRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderGetFulfillmentRequest) SetId(v string) {
	o.Id = &v
}

func (o OrderGetFulfillmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderGetFulfillmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableOrderGetFulfillmentRequest struct {
	value *OrderGetFulfillmentRequest
	isSet bool
}

func (v NullableOrderGetFulfillmentRequest) Get() *OrderGetFulfillmentRequest {
	return v.value
}

func (v *NullableOrderGetFulfillmentRequest) Set(val *OrderGetFulfillmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderGetFulfillmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderGetFulfillmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderGetFulfillmentRequest(val *OrderGetFulfillmentRequest) *NullableOrderGetFulfillmentRequest {
	return &NullableOrderGetFulfillmentRequest{value: val, isSet: true}
}

func (v NullableOrderGetFulfillmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderGetFulfillmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the OrderGetShipmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderGetShipmentRequest{}

// OrderGetShipmentRequest struct for OrderGetShipmentRequest
type OrderGetShipmentRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewOrderGetShipmentRequest instantiates a new OrderGetShipmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderGetShipmentRequest() *OrderGetShipmentRequest {
	this := OrderGetShipmentRequest{}
	return &this
}

// NewOrderGetShipmentRequestWithDefaults instantiates a new OrderGetShipmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderGetShipmentRequestWithDefaults() *OrderGetShipmentRequest {
	this := OrderGetShipmentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderGetShipmentRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetShipmentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderGetShipmentRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderGetShipmentRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderGetShipmentRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderGetShipmentRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderGetShipmentRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderGetShipmentRequest) SetId(v string) {
	o.Id = &v
}

func (o OrderGetShipmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderGetShipmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableOrderGetShipmentRequest struct {
	value *OrderGetShipmentRequest
	isSet bool
}

func (v NullableOrderGetShipmentRequest) Get() *OrderGetShipmentRequest {
	return v.value
}

func (v *NullableOrderGetShipmentRequest) Set(val *OrderGetShipmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderGetShipmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderGetShipmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderGetShipmentRequest(val *OrderGetShipmentRequest) *NullableOrderGetShipmentRequest {
	return &NullableOrderGetShipmentRequest{value: val, isSet: true}
}

func (v NullableOrderGetShipmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderGetShipmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



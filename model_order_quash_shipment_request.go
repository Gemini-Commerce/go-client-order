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

// checks if the OrderQuashShipmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderQuashShipmentRequest{}

// OrderQuashShipmentRequest struct for OrderQuashShipmentRequest
type OrderQuashShipmentRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	ShipmentId *string `json:"shipmentId,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// NewOrderQuashShipmentRequest instantiates a new OrderQuashShipmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderQuashShipmentRequest() *OrderQuashShipmentRequest {
	this := OrderQuashShipmentRequest{}
	return &this
}

// NewOrderQuashShipmentRequestWithDefaults instantiates a new OrderQuashShipmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderQuashShipmentRequestWithDefaults() *OrderQuashShipmentRequest {
	this := OrderQuashShipmentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderQuashShipmentRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderQuashShipmentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderQuashShipmentRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderQuashShipmentRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetShipmentId returns the ShipmentId field value if set, zero value otherwise.
func (o *OrderQuashShipmentRequest) GetShipmentId() string {
	if o == nil || IsNil(o.ShipmentId) {
		var ret string
		return ret
	}
	return *o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderQuashShipmentRequest) GetShipmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentId) {
		return nil, false
	}
	return o.ShipmentId, true
}

// HasShipmentId returns a boolean if a field has been set.
func (o *OrderQuashShipmentRequest) HasShipmentId() bool {
	if o != nil && !IsNil(o.ShipmentId) {
		return true
	}

	return false
}

// SetShipmentId gets a reference to the given string and assigns it to the ShipmentId field.
func (o *OrderQuashShipmentRequest) SetShipmentId(v string) {
	o.ShipmentId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *OrderQuashShipmentRequest) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderQuashShipmentRequest) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *OrderQuashShipmentRequest) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *OrderQuashShipmentRequest) SetReason(v string) {
	o.Reason = &v
}

func (o OrderQuashShipmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderQuashShipmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.ShipmentId) {
		toSerialize["shipmentId"] = o.ShipmentId
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableOrderQuashShipmentRequest struct {
	value *OrderQuashShipmentRequest
	isSet bool
}

func (v NullableOrderQuashShipmentRequest) Get() *OrderQuashShipmentRequest {
	return v.value
}

func (v *NullableOrderQuashShipmentRequest) Set(val *OrderQuashShipmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderQuashShipmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderQuashShipmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderQuashShipmentRequest(val *OrderQuashShipmentRequest) *NullableOrderQuashShipmentRequest {
	return &NullableOrderQuashShipmentRequest{value: val, isSet: true}
}

func (v NullableOrderQuashShipmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderQuashShipmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



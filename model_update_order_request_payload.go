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

// checks if the UpdateOrderRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrderRequestPayload{}

// UpdateOrderRequestPayload struct for UpdateOrderRequestPayload
type UpdateOrderRequestPayload struct {
	BillingAddress *OrderPostalAddress `json:"billingAddress,omitempty"`
	ShippingAddress *OrderPostalAddress `json:"shippingAddress,omitempty"`
	AdditionalInfo map[string]interface{} `json:"additionalInfo,omitempty"`
}

// NewUpdateOrderRequestPayload instantiates a new UpdateOrderRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrderRequestPayload() *UpdateOrderRequestPayload {
	this := UpdateOrderRequestPayload{}
	return &this
}

// NewUpdateOrderRequestPayloadWithDefaults instantiates a new UpdateOrderRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrderRequestPayloadWithDefaults() *UpdateOrderRequestPayload {
	this := UpdateOrderRequestPayload{}
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *UpdateOrderRequestPayload) GetBillingAddress() OrderPostalAddress {
	if o == nil || IsNil(o.BillingAddress) {
		var ret OrderPostalAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderRequestPayload) GetBillingAddressOk() (*OrderPostalAddress, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// BillingAddress returns a boolean if a field has been set.
func (o *UpdateOrderRequestPayload) BillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given OrderPostalAddress and assigns it to the BillingAddress field.
func (o *UpdateOrderRequestPayload) SetBillingAddress(v OrderPostalAddress) {
	o.BillingAddress = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *UpdateOrderRequestPayload) GetShippingAddress() OrderPostalAddress {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret OrderPostalAddress
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderRequestPayload) GetShippingAddressOk() (*OrderPostalAddress, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// ShippingAddress returns a boolean if a field has been set.
func (o *UpdateOrderRequestPayload) ShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given OrderPostalAddress and assigns it to the ShippingAddress field.
func (o *UpdateOrderRequestPayload) SetShippingAddress(v OrderPostalAddress) {
	o.ShippingAddress = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *UpdateOrderRequestPayload) GetAdditionalInfo() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderRequestPayload) GetAdditionalInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalInfo, true
}

// AdditionalInfo returns a boolean if a field has been set.
func (o *UpdateOrderRequestPayload) AdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given map[string]interface{} and assigns it to the AdditionalInfo field.
func (o *UpdateOrderRequestPayload) SetAdditionalInfo(v map[string]interface{}) {
	o.AdditionalInfo = v
}

func (o UpdateOrderRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrderRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !IsNil(o.ShippingAddress) {
		toSerialize["shippingAddress"] = o.ShippingAddress
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	return toSerialize, nil
}

type NullableUpdateOrderRequestPayload struct {
	value *UpdateOrderRequestPayload
	isSet bool
}

func (v NullableUpdateOrderRequestPayload) Get() *UpdateOrderRequestPayload {
	return v.value
}

func (v *NullableUpdateOrderRequestPayload) Set(val *UpdateOrderRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrderRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrderRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrderRequestPayload(val *UpdateOrderRequestPayload) *NullableUpdateOrderRequestPayload {
	return &NullableUpdateOrderRequestPayload{value: val, isSet: true}
}

func (v NullableUpdateOrderRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrderRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



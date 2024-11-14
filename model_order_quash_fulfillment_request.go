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
	"bytes"
	"fmt"
)

// checks if the OrderQuashFulfillmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderQuashFulfillmentRequest{}

// OrderQuashFulfillmentRequest struct for OrderQuashFulfillmentRequest
type OrderQuashFulfillmentRequest struct {
	TenantId string `json:"tenantId"`
	FulfillmentId string `json:"fulfillmentId"`
	Reason *string `json:"reason,omitempty"`
}

type _OrderQuashFulfillmentRequest OrderQuashFulfillmentRequest

// NewOrderQuashFulfillmentRequest instantiates a new OrderQuashFulfillmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderQuashFulfillmentRequest(tenantId string, fulfillmentId string) *OrderQuashFulfillmentRequest {
	this := OrderQuashFulfillmentRequest{}
	this.TenantId = tenantId
	this.FulfillmentId = fulfillmentId
	return &this
}

// NewOrderQuashFulfillmentRequestWithDefaults instantiates a new OrderQuashFulfillmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderQuashFulfillmentRequestWithDefaults() *OrderQuashFulfillmentRequest {
	this := OrderQuashFulfillmentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderQuashFulfillmentRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderQuashFulfillmentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderQuashFulfillmentRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetFulfillmentId returns the FulfillmentId field value
func (o *OrderQuashFulfillmentRequest) GetFulfillmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentId
}

// GetFulfillmentIdOk returns a tuple with the FulfillmentId field value
// and a boolean to check if the value has been set.
func (o *OrderQuashFulfillmentRequest) GetFulfillmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentId, true
}

// SetFulfillmentId sets field value
func (o *OrderQuashFulfillmentRequest) SetFulfillmentId(v string) {
	o.FulfillmentId = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *OrderQuashFulfillmentRequest) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderQuashFulfillmentRequest) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// Reason returns a boolean if a field has been set.
func (o *OrderQuashFulfillmentRequest) Reason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *OrderQuashFulfillmentRequest) SetReason(v string) {
	o.Reason = &v
}

func (o OrderQuashFulfillmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderQuashFulfillmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["fulfillmentId"] = o.FulfillmentId
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

func (o *OrderQuashFulfillmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"fulfillmentId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderQuashFulfillmentRequest := _OrderQuashFulfillmentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderQuashFulfillmentRequest)

	if err != nil {
		return err
	}

	*o = OrderQuashFulfillmentRequest(varOrderQuashFulfillmentRequest)

	return err
}

type NullableOrderQuashFulfillmentRequest struct {
	value *OrderQuashFulfillmentRequest
	isSet bool
}

func (v NullableOrderQuashFulfillmentRequest) Get() *OrderQuashFulfillmentRequest {
	return v.value
}

func (v *NullableOrderQuashFulfillmentRequest) Set(val *OrderQuashFulfillmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderQuashFulfillmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderQuashFulfillmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderQuashFulfillmentRequest(val *OrderQuashFulfillmentRequest) *NullableOrderQuashFulfillmentRequest {
	return &NullableOrderQuashFulfillmentRequest{value: val, isSet: true}
}

func (v NullableOrderQuashFulfillmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderQuashFulfillmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



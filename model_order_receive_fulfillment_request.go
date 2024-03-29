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
	"bytes"
	"fmt"
)

// checks if the OrderReceiveFulfillmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderReceiveFulfillmentRequest{}

// OrderReceiveFulfillmentRequest struct for OrderReceiveFulfillmentRequest
type OrderReceiveFulfillmentRequest struct {
	TenantId string `json:"tenantId"`
	FulfillmentId string `json:"fulfillmentId"`
}

type _OrderReceiveFulfillmentRequest OrderReceiveFulfillmentRequest

// NewOrderReceiveFulfillmentRequest instantiates a new OrderReceiveFulfillmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderReceiveFulfillmentRequest(tenantId string, fulfillmentId string) *OrderReceiveFulfillmentRequest {
	this := OrderReceiveFulfillmentRequest{}
	this.TenantId = tenantId
	this.FulfillmentId = fulfillmentId
	return &this
}

// NewOrderReceiveFulfillmentRequestWithDefaults instantiates a new OrderReceiveFulfillmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderReceiveFulfillmentRequestWithDefaults() *OrderReceiveFulfillmentRequest {
	this := OrderReceiveFulfillmentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderReceiveFulfillmentRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderReceiveFulfillmentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderReceiveFulfillmentRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetFulfillmentId returns the FulfillmentId field value
func (o *OrderReceiveFulfillmentRequest) GetFulfillmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentId
}

// GetFulfillmentIdOk returns a tuple with the FulfillmentId field value
// and a boolean to check if the value has been set.
func (o *OrderReceiveFulfillmentRequest) GetFulfillmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentId, true
}

// SetFulfillmentId sets field value
func (o *OrderReceiveFulfillmentRequest) SetFulfillmentId(v string) {
	o.FulfillmentId = v
}

func (o OrderReceiveFulfillmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderReceiveFulfillmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["fulfillmentId"] = o.FulfillmentId
	return toSerialize, nil
}

func (o *OrderReceiveFulfillmentRequest) UnmarshalJSON(data []byte) (err error) {
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

	varOrderReceiveFulfillmentRequest := _OrderReceiveFulfillmentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderReceiveFulfillmentRequest)

	if err != nil {
		return err
	}

	*o = OrderReceiveFulfillmentRequest(varOrderReceiveFulfillmentRequest)

	return err
}

type NullableOrderReceiveFulfillmentRequest struct {
	value *OrderReceiveFulfillmentRequest
	isSet bool
}

func (v NullableOrderReceiveFulfillmentRequest) Get() *OrderReceiveFulfillmentRequest {
	return v.value
}

func (v *NullableOrderReceiveFulfillmentRequest) Set(val *OrderReceiveFulfillmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderReceiveFulfillmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderReceiveFulfillmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderReceiveFulfillmentRequest(val *OrderReceiveFulfillmentRequest) *NullableOrderReceiveFulfillmentRequest {
	return &NullableOrderReceiveFulfillmentRequest{value: val, isSet: true}
}

func (v NullableOrderReceiveFulfillmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderReceiveFulfillmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



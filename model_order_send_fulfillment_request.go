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
	"fmt"
)

// checks if the OrderSendFulfillmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSendFulfillmentRequest{}

// OrderSendFulfillmentRequest struct for OrderSendFulfillmentRequest
type OrderSendFulfillmentRequest struct {
	TenantId             string `json:"tenantId"`
	FulfillmentId        string `json:"fulfillmentId"`
	AdditionalProperties map[string]interface{}
}

type _OrderSendFulfillmentRequest OrderSendFulfillmentRequest

// NewOrderSendFulfillmentRequest instantiates a new OrderSendFulfillmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSendFulfillmentRequest(tenantId string, fulfillmentId string) *OrderSendFulfillmentRequest {
	this := OrderSendFulfillmentRequest{}
	this.TenantId = tenantId
	this.FulfillmentId = fulfillmentId
	return &this
}

// NewOrderSendFulfillmentRequestWithDefaults instantiates a new OrderSendFulfillmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSendFulfillmentRequestWithDefaults() *OrderSendFulfillmentRequest {
	this := OrderSendFulfillmentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderSendFulfillmentRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderSendFulfillmentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderSendFulfillmentRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetFulfillmentId returns the FulfillmentId field value
func (o *OrderSendFulfillmentRequest) GetFulfillmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentId
}

// GetFulfillmentIdOk returns a tuple with the FulfillmentId field value
// and a boolean to check if the value has been set.
func (o *OrderSendFulfillmentRequest) GetFulfillmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentId, true
}

// SetFulfillmentId sets field value
func (o *OrderSendFulfillmentRequest) SetFulfillmentId(v string) {
	o.FulfillmentId = v
}

func (o OrderSendFulfillmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSendFulfillmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["fulfillmentId"] = o.FulfillmentId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderSendFulfillmentRequest) UnmarshalJSON(data []byte) (err error) {
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
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderSendFulfillmentRequest := _OrderSendFulfillmentRequest{}

	err = json.Unmarshal(data, &varOrderSendFulfillmentRequest)

	if err != nil {
		return err
	}

	*o = OrderSendFulfillmentRequest(varOrderSendFulfillmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "fulfillmentId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderSendFulfillmentRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderSendFulfillmentRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderSendFulfillmentRequest struct {
	value *OrderSendFulfillmentRequest
	isSet bool
}

func (v NullableOrderSendFulfillmentRequest) Get() *OrderSendFulfillmentRequest {
	return v.value
}

func (v *NullableOrderSendFulfillmentRequest) Set(val *OrderSendFulfillmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSendFulfillmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSendFulfillmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSendFulfillmentRequest(val *OrderSendFulfillmentRequest) *NullableOrderSendFulfillmentRequest {
	return &NullableOrderSendFulfillmentRequest{value: val, isSet: true}
}

func (v NullableOrderSendFulfillmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSendFulfillmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

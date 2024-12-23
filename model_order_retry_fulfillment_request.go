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

// checks if the OrderRetryFulfillmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderRetryFulfillmentRequest{}

// OrderRetryFulfillmentRequest struct for OrderRetryFulfillmentRequest
type OrderRetryFulfillmentRequest struct {
	TenantId             string `json:"tenantId"`
	FulfillmentId        string `json:"fulfillmentId"`
	AdditionalProperties map[string]interface{}
}

type _OrderRetryFulfillmentRequest OrderRetryFulfillmentRequest

// NewOrderRetryFulfillmentRequest instantiates a new OrderRetryFulfillmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderRetryFulfillmentRequest(tenantId string, fulfillmentId string) *OrderRetryFulfillmentRequest {
	this := OrderRetryFulfillmentRequest{}
	this.TenantId = tenantId
	this.FulfillmentId = fulfillmentId
	return &this
}

// NewOrderRetryFulfillmentRequestWithDefaults instantiates a new OrderRetryFulfillmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderRetryFulfillmentRequestWithDefaults() *OrderRetryFulfillmentRequest {
	this := OrderRetryFulfillmentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderRetryFulfillmentRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderRetryFulfillmentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderRetryFulfillmentRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetFulfillmentId returns the FulfillmentId field value
func (o *OrderRetryFulfillmentRequest) GetFulfillmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentId
}

// GetFulfillmentIdOk returns a tuple with the FulfillmentId field value
// and a boolean to check if the value has been set.
func (o *OrderRetryFulfillmentRequest) GetFulfillmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentId, true
}

// SetFulfillmentId sets field value
func (o *OrderRetryFulfillmentRequest) SetFulfillmentId(v string) {
	o.FulfillmentId = v
}

func (o OrderRetryFulfillmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderRetryFulfillmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["fulfillmentId"] = o.FulfillmentId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderRetryFulfillmentRequest) UnmarshalJSON(data []byte) (err error) {
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

	varOrderRetryFulfillmentRequest := _OrderRetryFulfillmentRequest{}

	err = json.Unmarshal(data, &varOrderRetryFulfillmentRequest)

	if err != nil {
		return err
	}

	*o = OrderRetryFulfillmentRequest(varOrderRetryFulfillmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "fulfillmentId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderRetryFulfillmentRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderRetryFulfillmentRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
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

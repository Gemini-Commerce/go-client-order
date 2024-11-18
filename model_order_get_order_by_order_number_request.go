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

// checks if the OrderGetOrderByOrderNumberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderGetOrderByOrderNumberRequest{}

// OrderGetOrderByOrderNumberRequest struct for OrderGetOrderByOrderNumberRequest
type OrderGetOrderByOrderNumberRequest struct {
	TenantId             string `json:"tenantId"`
	OrderNumber          string `json:"orderNumber"`
	AdditionalProperties map[string]interface{}
}

type _OrderGetOrderByOrderNumberRequest OrderGetOrderByOrderNumberRequest

// NewOrderGetOrderByOrderNumberRequest instantiates a new OrderGetOrderByOrderNumberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderGetOrderByOrderNumberRequest(tenantId string, orderNumber string) *OrderGetOrderByOrderNumberRequest {
	this := OrderGetOrderByOrderNumberRequest{}
	this.TenantId = tenantId
	this.OrderNumber = orderNumber
	return &this
}

// NewOrderGetOrderByOrderNumberRequestWithDefaults instantiates a new OrderGetOrderByOrderNumberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderGetOrderByOrderNumberRequestWithDefaults() *OrderGetOrderByOrderNumberRequest {
	this := OrderGetOrderByOrderNumberRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderGetOrderByOrderNumberRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderGetOrderByOrderNumberRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderGetOrderByOrderNumberRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetOrderNumber returns the OrderNumber field value
func (o *OrderGetOrderByOrderNumberRequest) GetOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value
// and a boolean to check if the value has been set.
func (o *OrderGetOrderByOrderNumberRequest) GetOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderNumber, true
}

// SetOrderNumber sets field value
func (o *OrderGetOrderByOrderNumberRequest) SetOrderNumber(v string) {
	o.OrderNumber = v
}

func (o OrderGetOrderByOrderNumberRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderGetOrderByOrderNumberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["orderNumber"] = o.OrderNumber

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderGetOrderByOrderNumberRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"orderNumber",
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

	varOrderGetOrderByOrderNumberRequest := _OrderGetOrderByOrderNumberRequest{}

	err = json.Unmarshal(data, &varOrderGetOrderByOrderNumberRequest)

	if err != nil {
		return err
	}

	*o = OrderGetOrderByOrderNumberRequest(varOrderGetOrderByOrderNumberRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "orderNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderGetOrderByOrderNumberRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderGetOrderByOrderNumberRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderGetOrderByOrderNumberRequest struct {
	value *OrderGetOrderByOrderNumberRequest
	isSet bool
}

func (v NullableOrderGetOrderByOrderNumberRequest) Get() *OrderGetOrderByOrderNumberRequest {
	return v.value
}

func (v *NullableOrderGetOrderByOrderNumberRequest) Set(val *OrderGetOrderByOrderNumberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderGetOrderByOrderNumberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderGetOrderByOrderNumberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderGetOrderByOrderNumberRequest(val *OrderGetOrderByOrderNumberRequest) *NullableOrderGetOrderByOrderNumberRequest {
	return &NullableOrderGetOrderByOrderNumberRequest{value: val, isSet: true}
}

func (v NullableOrderGetOrderByOrderNumberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderGetOrderByOrderNumberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

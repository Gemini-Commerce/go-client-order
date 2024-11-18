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

// checks if the OrderPrintOrdersLabelsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderPrintOrdersLabelsRequest{}

// OrderPrintOrdersLabelsRequest struct for OrderPrintOrdersLabelsRequest
type OrderPrintOrdersLabelsRequest struct {
	TenantId             string   `json:"tenantId"`
	OrderNumbers         []string `json:"orderNumbers"`
	AdditionalProperties map[string]interface{}
}

type _OrderPrintOrdersLabelsRequest OrderPrintOrdersLabelsRequest

// NewOrderPrintOrdersLabelsRequest instantiates a new OrderPrintOrdersLabelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderPrintOrdersLabelsRequest(tenantId string, orderNumbers []string) *OrderPrintOrdersLabelsRequest {
	this := OrderPrintOrdersLabelsRequest{}
	this.TenantId = tenantId
	this.OrderNumbers = orderNumbers
	return &this
}

// NewOrderPrintOrdersLabelsRequestWithDefaults instantiates a new OrderPrintOrdersLabelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderPrintOrdersLabelsRequestWithDefaults() *OrderPrintOrdersLabelsRequest {
	this := OrderPrintOrdersLabelsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderPrintOrdersLabelsRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderPrintOrdersLabelsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderPrintOrdersLabelsRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetOrderNumbers returns the OrderNumbers field value
func (o *OrderPrintOrdersLabelsRequest) GetOrderNumbers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OrderNumbers
}

// GetOrderNumbersOk returns a tuple with the OrderNumbers field value
// and a boolean to check if the value has been set.
func (o *OrderPrintOrdersLabelsRequest) GetOrderNumbersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderNumbers, true
}

// SetOrderNumbers sets field value
func (o *OrderPrintOrdersLabelsRequest) SetOrderNumbers(v []string) {
	o.OrderNumbers = v
}

func (o OrderPrintOrdersLabelsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderPrintOrdersLabelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["orderNumbers"] = o.OrderNumbers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderPrintOrdersLabelsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"orderNumbers",
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

	varOrderPrintOrdersLabelsRequest := _OrderPrintOrdersLabelsRequest{}

	err = json.Unmarshal(data, &varOrderPrintOrdersLabelsRequest)

	if err != nil {
		return err
	}

	*o = OrderPrintOrdersLabelsRequest(varOrderPrintOrdersLabelsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "orderNumbers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderPrintOrdersLabelsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderPrintOrdersLabelsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderPrintOrdersLabelsRequest struct {
	value *OrderPrintOrdersLabelsRequest
	isSet bool
}

func (v NullableOrderPrintOrdersLabelsRequest) Get() *OrderPrintOrdersLabelsRequest {
	return v.value
}

func (v *NullableOrderPrintOrdersLabelsRequest) Set(val *OrderPrintOrdersLabelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderPrintOrdersLabelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderPrintOrdersLabelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderPrintOrdersLabelsRequest(val *OrderPrintOrdersLabelsRequest) *NullableOrderPrintOrdersLabelsRequest {
	return &NullableOrderPrintOrdersLabelsRequest{value: val, isSet: true}
}

func (v NullableOrderPrintOrdersLabelsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderPrintOrdersLabelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

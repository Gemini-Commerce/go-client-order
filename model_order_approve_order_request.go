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

// checks if the OrderApproveOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderApproveOrderRequest{}

// OrderApproveOrderRequest struct for OrderApproveOrderRequest
type OrderApproveOrderRequest struct {
	TenantId             string `json:"tenantId"`
	OrderId              string `json:"orderId"`
	AdditionalProperties map[string]interface{}
}

type _OrderApproveOrderRequest OrderApproveOrderRequest

// NewOrderApproveOrderRequest instantiates a new OrderApproveOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderApproveOrderRequest(tenantId string, orderId string) *OrderApproveOrderRequest {
	this := OrderApproveOrderRequest{}
	this.TenantId = tenantId
	this.OrderId = orderId
	return &this
}

// NewOrderApproveOrderRequestWithDefaults instantiates a new OrderApproveOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderApproveOrderRequestWithDefaults() *OrderApproveOrderRequest {
	this := OrderApproveOrderRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderApproveOrderRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderApproveOrderRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderApproveOrderRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetOrderId returns the OrderId field value
func (o *OrderApproveOrderRequest) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *OrderApproveOrderRequest) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *OrderApproveOrderRequest) SetOrderId(v string) {
	o.OrderId = v
}

func (o OrderApproveOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderApproveOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["orderId"] = o.OrderId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderApproveOrderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"orderId",
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

	varOrderApproveOrderRequest := _OrderApproveOrderRequest{}

	err = json.Unmarshal(data, &varOrderApproveOrderRequest)

	if err != nil {
		return err
	}

	*o = OrderApproveOrderRequest(varOrderApproveOrderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderApproveOrderRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderApproveOrderRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderApproveOrderRequest struct {
	value *OrderApproveOrderRequest
	isSet bool
}

func (v NullableOrderApproveOrderRequest) Get() *OrderApproveOrderRequest {
	return v.value
}

func (v *NullableOrderApproveOrderRequest) Set(val *OrderApproveOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderApproveOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderApproveOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderApproveOrderRequest(val *OrderApproveOrderRequest) *NullableOrderApproveOrderRequest {
	return &NullableOrderApproveOrderRequest{value: val, isSet: true}
}

func (v NullableOrderApproveOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderApproveOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

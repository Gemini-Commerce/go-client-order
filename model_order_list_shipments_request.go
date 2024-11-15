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

// checks if the OrderListShipmentsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderListShipmentsRequest{}

// OrderListShipmentsRequest struct for OrderListShipmentsRequest
type OrderListShipmentsRequest struct {
	TenantId string `json:"tenantId"`
	OrderId *string `json:"orderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListShipmentsRequest OrderListShipmentsRequest

// NewOrderListShipmentsRequest instantiates a new OrderListShipmentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListShipmentsRequest(tenantId string) *OrderListShipmentsRequest {
	this := OrderListShipmentsRequest{}
	this.TenantId = tenantId
	return &this
}

// NewOrderListShipmentsRequestWithDefaults instantiates a new OrderListShipmentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListShipmentsRequestWithDefaults() *OrderListShipmentsRequest {
	this := OrderListShipmentsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderListShipmentsRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderListShipmentsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderListShipmentsRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderListShipmentsRequest) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListShipmentsRequest) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderListShipmentsRequest) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderListShipmentsRequest) SetOrderId(v string) {
	o.OrderId = &v
}

func (o OrderListShipmentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListShipmentsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderListShipmentsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
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

	varOrderListShipmentsRequest := _OrderListShipmentsRequest{}

	err = json.Unmarshal(data, &varOrderListShipmentsRequest)

	if err != nil {
		return err
	}

	*o = OrderListShipmentsRequest(varOrderListShipmentsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderListShipmentsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderListShipmentsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderListShipmentsRequest struct {
	value *OrderListShipmentsRequest
	isSet bool
}

func (v NullableOrderListShipmentsRequest) Get() *OrderListShipmentsRequest {
	return v.value
}

func (v *NullableOrderListShipmentsRequest) Set(val *OrderListShipmentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListShipmentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListShipmentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListShipmentsRequest(val *OrderListShipmentsRequest) *NullableOrderListShipmentsRequest {
	return &NullableOrderListShipmentsRequest{value: val, isSet: true}
}

func (v NullableOrderListShipmentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListShipmentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



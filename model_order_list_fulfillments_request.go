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

// checks if the OrderListFulfillmentsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderListFulfillmentsRequest{}

// OrderListFulfillmentsRequest struct for OrderListFulfillmentsRequest
type OrderListFulfillmentsRequest struct {
	TenantId string `json:"tenantId"`
	OrderId *string `json:"orderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListFulfillmentsRequest OrderListFulfillmentsRequest

// NewOrderListFulfillmentsRequest instantiates a new OrderListFulfillmentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListFulfillmentsRequest(tenantId string) *OrderListFulfillmentsRequest {
	this := OrderListFulfillmentsRequest{}
	this.TenantId = tenantId
	return &this
}

// NewOrderListFulfillmentsRequestWithDefaults instantiates a new OrderListFulfillmentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListFulfillmentsRequestWithDefaults() *OrderListFulfillmentsRequest {
	this := OrderListFulfillmentsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderListFulfillmentsRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderListFulfillmentsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderListFulfillmentsRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderListFulfillmentsRequest) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListFulfillmentsRequest) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// IsSetOrderId returns a boolean if a field has been set.
func (o *OrderListFulfillmentsRequest) IsSetOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderListFulfillmentsRequest) SetOrderId(v string) {
	o.OrderId = &v
}

func (o OrderListFulfillmentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListFulfillmentsRequest) ToMap() (map[string]interface{}, error) {
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

func (o *OrderListFulfillmentsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varOrderListFulfillmentsRequest := _OrderListFulfillmentsRequest{}

	err = json.Unmarshal(data, &varOrderListFulfillmentsRequest)

	if err != nil {
		return err
	}

	*o = OrderListFulfillmentsRequest(varOrderListFulfillmentsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "orderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderListFulfillmentsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderListFulfillmentsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderListFulfillmentsRequest struct {
	value *OrderListFulfillmentsRequest
	isSet bool
}

func (v NullableOrderListFulfillmentsRequest) Get() *OrderListFulfillmentsRequest {
	return v.value
}

func (v *NullableOrderListFulfillmentsRequest) Set(val *OrderListFulfillmentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListFulfillmentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListFulfillmentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListFulfillmentsRequest(val *OrderListFulfillmentsRequest) *NullableOrderListFulfillmentsRequest {
	return &NullableOrderListFulfillmentsRequest{value: val, isSet: true}
}

func (v NullableOrderListFulfillmentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListFulfillmentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



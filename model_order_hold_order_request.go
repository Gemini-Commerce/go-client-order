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

// checks if the OrderHoldOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderHoldOrderRequest{}

// OrderHoldOrderRequest struct for OrderHoldOrderRequest
type OrderHoldOrderRequest struct {
	TenantId             *string `json:"tenantId,omitempty"`
	OrderId              string  `json:"orderId"`
	Reason               *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderHoldOrderRequest OrderHoldOrderRequest

// NewOrderHoldOrderRequest instantiates a new OrderHoldOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderHoldOrderRequest(orderId string) *OrderHoldOrderRequest {
	this := OrderHoldOrderRequest{}
	this.OrderId = orderId
	return &this
}

// NewOrderHoldOrderRequestWithDefaults instantiates a new OrderHoldOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderHoldOrderRequestWithDefaults() *OrderHoldOrderRequest {
	this := OrderHoldOrderRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderHoldOrderRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderHoldOrderRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderHoldOrderRequest) IsSetTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderHoldOrderRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetOrderId returns the OrderId field value
func (o *OrderHoldOrderRequest) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *OrderHoldOrderRequest) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *OrderHoldOrderRequest) SetOrderId(v string) {
	o.OrderId = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *OrderHoldOrderRequest) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderHoldOrderRequest) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *OrderHoldOrderRequest) IsSetReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *OrderHoldOrderRequest) SetReason(v string) {
	o.Reason = &v
}

func (o OrderHoldOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderHoldOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	toSerialize["orderId"] = o.OrderId
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderHoldOrderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varOrderHoldOrderRequest := _OrderHoldOrderRequest{}

	err = json.Unmarshal(data, &varOrderHoldOrderRequest)

	if err != nil {
		return err
	}

	*o = OrderHoldOrderRequest(varOrderHoldOrderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderHoldOrderRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderHoldOrderRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderHoldOrderRequest struct {
	value *OrderHoldOrderRequest
	isSet bool
}

func (v NullableOrderHoldOrderRequest) Get() *OrderHoldOrderRequest {
	return v.value
}

func (v *NullableOrderHoldOrderRequest) Set(val *OrderHoldOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderHoldOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderHoldOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderHoldOrderRequest(val *OrderHoldOrderRequest) *NullableOrderHoldOrderRequest {
	return &NullableOrderHoldOrderRequest{value: val, isSet: true}
}

func (v NullableOrderHoldOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderHoldOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

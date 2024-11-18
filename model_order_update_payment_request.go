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

// checks if the OrderUpdatePaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderUpdatePaymentRequest{}

// OrderUpdatePaymentRequest struct for OrderUpdatePaymentRequest
type OrderUpdatePaymentRequest struct {
	TenantId             string         `json:"tenantId"`
	PaymentId            string         `json:"paymentId"`
	CcInfo               *PaymentCcInfo `json:"ccInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderUpdatePaymentRequest OrderUpdatePaymentRequest

// NewOrderUpdatePaymentRequest instantiates a new OrderUpdatePaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderUpdatePaymentRequest(tenantId string, paymentId string) *OrderUpdatePaymentRequest {
	this := OrderUpdatePaymentRequest{}
	this.TenantId = tenantId
	this.PaymentId = paymentId
	return &this
}

// NewOrderUpdatePaymentRequestWithDefaults instantiates a new OrderUpdatePaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderUpdatePaymentRequestWithDefaults() *OrderUpdatePaymentRequest {
	this := OrderUpdatePaymentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderUpdatePaymentRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderUpdatePaymentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderUpdatePaymentRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetPaymentId returns the PaymentId field value
func (o *OrderUpdatePaymentRequest) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *OrderUpdatePaymentRequest) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *OrderUpdatePaymentRequest) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetCcInfo returns the CcInfo field value if set, zero value otherwise.
func (o *OrderUpdatePaymentRequest) GetCcInfo() PaymentCcInfo {
	if o == nil || IsNil(o.CcInfo) {
		var ret PaymentCcInfo
		return ret
	}
	return *o.CcInfo
}

// GetCcInfoOk returns a tuple with the CcInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdatePaymentRequest) GetCcInfoOk() (*PaymentCcInfo, bool) {
	if o == nil || IsNil(o.CcInfo) {
		return nil, false
	}
	return o.CcInfo, true
}

// HasCcInfo returns a boolean if a field has been set.
func (o *OrderUpdatePaymentRequest) HasCcInfo() bool {
	if o != nil && !IsNil(o.CcInfo) {
		return true
	}

	return false
}

// SetCcInfo gets a reference to the given PaymentCcInfo and assigns it to the CcInfo field.
func (o *OrderUpdatePaymentRequest) SetCcInfo(v PaymentCcInfo) {
	o.CcInfo = &v
}

func (o OrderUpdatePaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderUpdatePaymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["paymentId"] = o.PaymentId
	if !IsNil(o.CcInfo) {
		toSerialize["ccInfo"] = o.CcInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderUpdatePaymentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"paymentId",
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

	varOrderUpdatePaymentRequest := _OrderUpdatePaymentRequest{}

	err = json.Unmarshal(data, &varOrderUpdatePaymentRequest)

	if err != nil {
		return err
	}

	*o = OrderUpdatePaymentRequest(varOrderUpdatePaymentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "paymentId")
		delete(additionalProperties, "ccInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderUpdatePaymentRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderUpdatePaymentRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderUpdatePaymentRequest struct {
	value *OrderUpdatePaymentRequest
	isSet bool
}

func (v NullableOrderUpdatePaymentRequest) Get() *OrderUpdatePaymentRequest {
	return v.value
}

func (v *NullableOrderUpdatePaymentRequest) Set(val *OrderUpdatePaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderUpdatePaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderUpdatePaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderUpdatePaymentRequest(val *OrderUpdatePaymentRequest) *NullableOrderUpdatePaymentRequest {
	return &NullableOrderUpdatePaymentRequest{value: val, isSet: true}
}

func (v NullableOrderUpdatePaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderUpdatePaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the OrderCalculateRefundRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCalculateRefundRequest{}

// OrderCalculateRefundRequest struct for OrderCalculateRefundRequest
type OrderCalculateRefundRequest struct {
	TenantId string `json:"tenantId"`
	PaymentId string `json:"paymentId"`
	Items []OrderRefundItem `json:"items,omitempty"`
	// Boolean indicating whether to calculate refund for shipping.
	Shipping *bool `json:"shipping,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderCalculateRefundRequest OrderCalculateRefundRequest

// NewOrderCalculateRefundRequest instantiates a new OrderCalculateRefundRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCalculateRefundRequest(tenantId string, paymentId string) *OrderCalculateRefundRequest {
	this := OrderCalculateRefundRequest{}
	this.TenantId = tenantId
	this.PaymentId = paymentId
	return &this
}

// NewOrderCalculateRefundRequestWithDefaults instantiates a new OrderCalculateRefundRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCalculateRefundRequestWithDefaults() *OrderCalculateRefundRequest {
	this := OrderCalculateRefundRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderCalculateRefundRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderCalculateRefundRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderCalculateRefundRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetPaymentId returns the PaymentId field value
func (o *OrderCalculateRefundRequest) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *OrderCalculateRefundRequest) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *OrderCalculateRefundRequest) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *OrderCalculateRefundRequest) GetItems() []OrderRefundItem {
	if o == nil || IsNil(o.Items) {
		var ret []OrderRefundItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCalculateRefundRequest) GetItemsOk() ([]OrderRefundItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *OrderCalculateRefundRequest) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrderRefundItem and assigns it to the Items field.
func (o *OrderCalculateRefundRequest) SetItems(v []OrderRefundItem) {
	o.Items = v
}

// GetShipping returns the Shipping field value if set, zero value otherwise.
func (o *OrderCalculateRefundRequest) GetShipping() bool {
	if o == nil || IsNil(o.Shipping) {
		var ret bool
		return ret
	}
	return *o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCalculateRefundRequest) GetShippingOk() (*bool, bool) {
	if o == nil || IsNil(o.Shipping) {
		return nil, false
	}
	return o.Shipping, true
}

// HasShipping returns a boolean if a field has been set.
func (o *OrderCalculateRefundRequest) HasShipping() bool {
	if o != nil && !IsNil(o.Shipping) {
		return true
	}

	return false
}

// SetShipping gets a reference to the given bool and assigns it to the Shipping field.
func (o *OrderCalculateRefundRequest) SetShipping(v bool) {
	o.Shipping = &v
}

func (o OrderCalculateRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCalculateRefundRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["paymentId"] = o.PaymentId
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Shipping) {
		toSerialize["shipping"] = o.Shipping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderCalculateRefundRequest) UnmarshalJSON(data []byte) (err error) {
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
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderCalculateRefundRequest := _OrderCalculateRefundRequest{}

	err = json.Unmarshal(data, &varOrderCalculateRefundRequest)

	if err != nil {
		return err
	}

	*o = OrderCalculateRefundRequest(varOrderCalculateRefundRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "paymentId")
		delete(additionalProperties, "items")
		delete(additionalProperties, "shipping")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderCalculateRefundRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderCalculateRefundRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderCalculateRefundRequest struct {
	value *OrderCalculateRefundRequest
	isSet bool
}

func (v NullableOrderCalculateRefundRequest) Get() *OrderCalculateRefundRequest {
	return v.value
}

func (v *NullableOrderCalculateRefundRequest) Set(val *OrderCalculateRefundRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCalculateRefundRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCalculateRefundRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCalculateRefundRequest(val *OrderCalculateRefundRequest) *NullableOrderCalculateRefundRequest {
	return &NullableOrderCalculateRefundRequest{value: val, isSet: true}
}

func (v NullableOrderCalculateRefundRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCalculateRefundRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



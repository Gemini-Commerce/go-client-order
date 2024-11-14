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
)

// checks if the OrderSendOrderNotificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSendOrderNotificationRequest{}

// OrderSendOrderNotificationRequest struct for OrderSendOrderNotificationRequest
type OrderSendOrderNotificationRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderSendOrderNotificationRequest OrderSendOrderNotificationRequest

// NewOrderSendOrderNotificationRequest instantiates a new OrderSendOrderNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSendOrderNotificationRequest() *OrderSendOrderNotificationRequest {
	this := OrderSendOrderNotificationRequest{}
	return &this
}

// NewOrderSendOrderNotificationRequestWithDefaults instantiates a new OrderSendOrderNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSendOrderNotificationRequestWithDefaults() *OrderSendOrderNotificationRequest {
	this := OrderSendOrderNotificationRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderSendOrderNotificationRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSendOrderNotificationRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// &#39;Has&#39;TenantId returns a boolean if a field has been set.
func (o *OrderSendOrderNotificationRequest) &#39;Has&#39;TenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderSendOrderNotificationRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderSendOrderNotificationRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSendOrderNotificationRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// &#39;Has&#39;Id returns a boolean if a field has been set.
func (o *OrderSendOrderNotificationRequest) &#39;Has&#39;Id() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderSendOrderNotificationRequest) SetId(v string) {
	o.Id = &v
}

func (o OrderSendOrderNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSendOrderNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderSendOrderNotificationRequest) UnmarshalJSON(data []byte) (err error) {
	varOrderSendOrderNotificationRequest := _OrderSendOrderNotificationRequest{}

	err = json.Unmarshal(data, &varOrderSendOrderNotificationRequest)

	if err != nil {
		return err
	}

	*o = OrderSendOrderNotificationRequest(varOrderSendOrderNotificationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderSendOrderNotificationRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderSendOrderNotificationRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderSendOrderNotificationRequest struct {
	value *OrderSendOrderNotificationRequest
	isSet bool
}

func (v NullableOrderSendOrderNotificationRequest) Get() *OrderSendOrderNotificationRequest {
	return v.value
}

func (v *NullableOrderSendOrderNotificationRequest) Set(val *OrderSendOrderNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSendOrderNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSendOrderNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSendOrderNotificationRequest(val *OrderSendOrderNotificationRequest) *NullableOrderSendOrderNotificationRequest {
	return &NullableOrderSendOrderNotificationRequest{value: val, isSet: true}
}

func (v NullableOrderSendOrderNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSendOrderNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



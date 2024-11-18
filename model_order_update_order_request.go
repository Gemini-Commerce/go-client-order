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

// checks if the OrderUpdateOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderUpdateOrderRequest{}

// OrderUpdateOrderRequest struct for OrderUpdateOrderRequest
type OrderUpdateOrderRequest struct {
	TenantId             string                     `json:"tenantId"`
	Id                   string                     `json:"id"`
	Payload              *UpdateOrderRequestPayload `json:"payload,omitempty"`
	FieldMask            *string                    `json:"fieldMask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderUpdateOrderRequest OrderUpdateOrderRequest

// NewOrderUpdateOrderRequest instantiates a new OrderUpdateOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderUpdateOrderRequest(tenantId string, id string) *OrderUpdateOrderRequest {
	this := OrderUpdateOrderRequest{}
	this.TenantId = tenantId
	this.Id = id
	return &this
}

// NewOrderUpdateOrderRequestWithDefaults instantiates a new OrderUpdateOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderUpdateOrderRequestWithDefaults() *OrderUpdateOrderRequest {
	this := OrderUpdateOrderRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderUpdateOrderRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderUpdateOrderRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderUpdateOrderRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *OrderUpdateOrderRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderUpdateOrderRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderUpdateOrderRequest) SetId(v string) {
	o.Id = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *OrderUpdateOrderRequest) GetPayload() UpdateOrderRequestPayload {
	if o == nil || IsNil(o.Payload) {
		var ret UpdateOrderRequestPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateOrderRequest) GetPayloadOk() (*UpdateOrderRequestPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *OrderUpdateOrderRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given UpdateOrderRequestPayload and assigns it to the Payload field.
func (o *OrderUpdateOrderRequest) SetPayload(v UpdateOrderRequestPayload) {
	o.Payload = &v
}

// GetFieldMask returns the FieldMask field value if set, zero value otherwise.
func (o *OrderUpdateOrderRequest) GetFieldMask() string {
	if o == nil || IsNil(o.FieldMask) {
		var ret string
		return ret
	}
	return *o.FieldMask
}

// GetFieldMaskOk returns a tuple with the FieldMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateOrderRequest) GetFieldMaskOk() (*string, bool) {
	if o == nil || IsNil(o.FieldMask) {
		return nil, false
	}
	return o.FieldMask, true
}

// HasFieldMask returns a boolean if a field has been set.
func (o *OrderUpdateOrderRequest) HasFieldMask() bool {
	if o != nil && !IsNil(o.FieldMask) {
		return true
	}

	return false
}

// SetFieldMask gets a reference to the given string and assigns it to the FieldMask field.
func (o *OrderUpdateOrderRequest) SetFieldMask(v string) {
	o.FieldMask = &v
}

func (o OrderUpdateOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderUpdateOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["id"] = o.Id
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.FieldMask) {
		toSerialize["fieldMask"] = o.FieldMask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderUpdateOrderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"id",
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

	varOrderUpdateOrderRequest := _OrderUpdateOrderRequest{}

	err = json.Unmarshal(data, &varOrderUpdateOrderRequest)

	if err != nil {
		return err
	}

	*o = OrderUpdateOrderRequest(varOrderUpdateOrderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "payload")
		delete(additionalProperties, "fieldMask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderUpdateOrderRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderUpdateOrderRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderUpdateOrderRequest struct {
	value *OrderUpdateOrderRequest
	isSet bool
}

func (v NullableOrderUpdateOrderRequest) Get() *OrderUpdateOrderRequest {
	return v.value
}

func (v *NullableOrderUpdateOrderRequest) Set(val *OrderUpdateOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderUpdateOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderUpdateOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderUpdateOrderRequest(val *OrderUpdateOrderRequest) *NullableOrderUpdateOrderRequest {
	return &NullableOrderUpdateOrderRequest{value: val, isSet: true}
}

func (v NullableOrderUpdateOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderUpdateOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

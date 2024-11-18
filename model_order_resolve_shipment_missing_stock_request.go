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

// checks if the OrderResolveShipmentMissingStockRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResolveShipmentMissingStockRequest{}

// OrderResolveShipmentMissingStockRequest struct for OrderResolveShipmentMissingStockRequest
type OrderResolveShipmentMissingStockRequest struct {
	TenantId             string `json:"tenantId"`
	ShipmentId           string `json:"shipmentId"`
	AdditionalProperties map[string]interface{}
}

type _OrderResolveShipmentMissingStockRequest OrderResolveShipmentMissingStockRequest

// NewOrderResolveShipmentMissingStockRequest instantiates a new OrderResolveShipmentMissingStockRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResolveShipmentMissingStockRequest(tenantId string, shipmentId string) *OrderResolveShipmentMissingStockRequest {
	this := OrderResolveShipmentMissingStockRequest{}
	this.TenantId = tenantId
	this.ShipmentId = shipmentId
	return &this
}

// NewOrderResolveShipmentMissingStockRequestWithDefaults instantiates a new OrderResolveShipmentMissingStockRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResolveShipmentMissingStockRequestWithDefaults() *OrderResolveShipmentMissingStockRequest {
	this := OrderResolveShipmentMissingStockRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderResolveShipmentMissingStockRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderResolveShipmentMissingStockRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderResolveShipmentMissingStockRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetShipmentId returns the ShipmentId field value
func (o *OrderResolveShipmentMissingStockRequest) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *OrderResolveShipmentMissingStockRequest) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *OrderResolveShipmentMissingStockRequest) SetShipmentId(v string) {
	o.ShipmentId = v
}

func (o OrderResolveShipmentMissingStockRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResolveShipmentMissingStockRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["shipmentId"] = o.ShipmentId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderResolveShipmentMissingStockRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"shipmentId",
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

	varOrderResolveShipmentMissingStockRequest := _OrderResolveShipmentMissingStockRequest{}

	err = json.Unmarshal(data, &varOrderResolveShipmentMissingStockRequest)

	if err != nil {
		return err
	}

	*o = OrderResolveShipmentMissingStockRequest(varOrderResolveShipmentMissingStockRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "shipmentId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderResolveShipmentMissingStockRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderResolveShipmentMissingStockRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderResolveShipmentMissingStockRequest struct {
	value *OrderResolveShipmentMissingStockRequest
	isSet bool
}

func (v NullableOrderResolveShipmentMissingStockRequest) Get() *OrderResolveShipmentMissingStockRequest {
	return v.value
}

func (v *NullableOrderResolveShipmentMissingStockRequest) Set(val *OrderResolveShipmentMissingStockRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResolveShipmentMissingStockRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResolveShipmentMissingStockRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResolveShipmentMissingStockRequest(val *OrderResolveShipmentMissingStockRequest) *NullableOrderResolveShipmentMissingStockRequest {
	return &NullableOrderResolveShipmentMissingStockRequest{value: val, isSet: true}
}

func (v NullableOrderResolveShipmentMissingStockRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResolveShipmentMissingStockRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the OrderReportShipmentMissingStockRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderReportShipmentMissingStockRequest{}

// OrderReportShipmentMissingStockRequest struct for OrderReportShipmentMissingStockRequest
type OrderReportShipmentMissingStockRequest struct {
	TenantId             string  `json:"tenantId"`
	ShipmentId           string  `json:"shipmentId"`
	Reason               *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderReportShipmentMissingStockRequest OrderReportShipmentMissingStockRequest

// NewOrderReportShipmentMissingStockRequest instantiates a new OrderReportShipmentMissingStockRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderReportShipmentMissingStockRequest(tenantId string, shipmentId string) *OrderReportShipmentMissingStockRequest {
	this := OrderReportShipmentMissingStockRequest{}
	this.TenantId = tenantId
	this.ShipmentId = shipmentId
	return &this
}

// NewOrderReportShipmentMissingStockRequestWithDefaults instantiates a new OrderReportShipmentMissingStockRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderReportShipmentMissingStockRequestWithDefaults() *OrderReportShipmentMissingStockRequest {
	this := OrderReportShipmentMissingStockRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderReportShipmentMissingStockRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderReportShipmentMissingStockRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderReportShipmentMissingStockRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetShipmentId returns the ShipmentId field value
func (o *OrderReportShipmentMissingStockRequest) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *OrderReportShipmentMissingStockRequest) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *OrderReportShipmentMissingStockRequest) SetShipmentId(v string) {
	o.ShipmentId = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *OrderReportShipmentMissingStockRequest) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReportShipmentMissingStockRequest) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *OrderReportShipmentMissingStockRequest) IsSetReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *OrderReportShipmentMissingStockRequest) SetReason(v string) {
	o.Reason = &v
}

func (o OrderReportShipmentMissingStockRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderReportShipmentMissingStockRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["shipmentId"] = o.ShipmentId
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderReportShipmentMissingStockRequest) UnmarshalJSON(data []byte) (err error) {
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

	varOrderReportShipmentMissingStockRequest := _OrderReportShipmentMissingStockRequest{}

	err = json.Unmarshal(data, &varOrderReportShipmentMissingStockRequest)

	if err != nil {
		return err
	}

	*o = OrderReportShipmentMissingStockRequest(varOrderReportShipmentMissingStockRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "shipmentId")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderReportShipmentMissingStockRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderReportShipmentMissingStockRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderReportShipmentMissingStockRequest struct {
	value *OrderReportShipmentMissingStockRequest
	isSet bool
}

func (v NullableOrderReportShipmentMissingStockRequest) Get() *OrderReportShipmentMissingStockRequest {
	return v.value
}

func (v *NullableOrderReportShipmentMissingStockRequest) Set(val *OrderReportShipmentMissingStockRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderReportShipmentMissingStockRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderReportShipmentMissingStockRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderReportShipmentMissingStockRequest(val *OrderReportShipmentMissingStockRequest) *NullableOrderReportShipmentMissingStockRequest {
	return &NullableOrderReportShipmentMissingStockRequest{value: val, isSet: true}
}

func (v NullableOrderReportShipmentMissingStockRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderReportShipmentMissingStockRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

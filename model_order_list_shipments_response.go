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

// checks if the OrderListShipmentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderListShipmentsResponse{}

// OrderListShipmentsResponse struct for OrderListShipmentsResponse
type OrderListShipmentsResponse struct {
	Shipments            []OrderShipment `json:"shipments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListShipmentsResponse OrderListShipmentsResponse

// NewOrderListShipmentsResponse instantiates a new OrderListShipmentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListShipmentsResponse() *OrderListShipmentsResponse {
	this := OrderListShipmentsResponse{}
	return &this
}

// NewOrderListShipmentsResponseWithDefaults instantiates a new OrderListShipmentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListShipmentsResponseWithDefaults() *OrderListShipmentsResponse {
	this := OrderListShipmentsResponse{}
	return &this
}

// GetShipments returns the Shipments field value if set, zero value otherwise.
func (o *OrderListShipmentsResponse) GetShipments() []OrderShipment {
	if o == nil || IsNil(o.Shipments) {
		var ret []OrderShipment
		return ret
	}
	return o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListShipmentsResponse) GetShipmentsOk() ([]OrderShipment, bool) {
	if o == nil || IsNil(o.Shipments) {
		return nil, false
	}
	return o.Shipments, true
}

// HasShipments returns a boolean if a field has been set.
func (o *OrderListShipmentsResponse) IsSetShipments() bool {
	if o != nil && !IsNil(o.Shipments) {
		return true
	}

	return false
}

// SetShipments gets a reference to the given []OrderShipment and assigns it to the Shipments field.
func (o *OrderListShipmentsResponse) SetShipments(v []OrderShipment) {
	o.Shipments = v
}

func (o OrderListShipmentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListShipmentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Shipments) {
		toSerialize["shipments"] = o.Shipments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderListShipmentsResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderListShipmentsResponse := _OrderListShipmentsResponse{}

	err = json.Unmarshal(data, &varOrderListShipmentsResponse)

	if err != nil {
		return err
	}

	*o = OrderListShipmentsResponse(varOrderListShipmentsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "shipments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderListShipmentsResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderListShipmentsResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderListShipmentsResponse struct {
	value *OrderListShipmentsResponse
	isSet bool
}

func (v NullableOrderListShipmentsResponse) Get() *OrderListShipmentsResponse {
	return v.value
}

func (v *NullableOrderListShipmentsResponse) Set(val *OrderListShipmentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListShipmentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListShipmentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListShipmentsResponse(val *OrderListShipmentsResponse) *NullableOrderListShipmentsResponse {
	return &NullableOrderListShipmentsResponse{value: val, isSet: true}
}

func (v NullableOrderListShipmentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListShipmentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

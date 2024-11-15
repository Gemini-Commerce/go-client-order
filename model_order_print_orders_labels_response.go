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

// checks if the OrderPrintOrdersLabelsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderPrintOrdersLabelsResponse{}

// OrderPrintOrdersLabelsResponse struct for OrderPrintOrdersLabelsResponse
type OrderPrintOrdersLabelsResponse struct {
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	FailedOrders []PrintOrdersLabelsResponseFailedOrder `json:"failedOrders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderPrintOrdersLabelsResponse OrderPrintOrdersLabelsResponse

// NewOrderPrintOrdersLabelsResponse instantiates a new OrderPrintOrdersLabelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderPrintOrdersLabelsResponse() *OrderPrintOrdersLabelsResponse {
	this := OrderPrintOrdersLabelsResponse{}
	return &this
}

// NewOrderPrintOrdersLabelsResponseWithDefaults instantiates a new OrderPrintOrdersLabelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderPrintOrdersLabelsResponseWithDefaults() *OrderPrintOrdersLabelsResponse {
	this := OrderPrintOrdersLabelsResponse{}
	return &this
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *OrderPrintOrdersLabelsResponse) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPrintOrdersLabelsResponse) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *OrderPrintOrdersLabelsResponse) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *OrderPrintOrdersLabelsResponse) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetFailedOrders returns the FailedOrders field value if set, zero value otherwise.
func (o *OrderPrintOrdersLabelsResponse) GetFailedOrders() []PrintOrdersLabelsResponseFailedOrder {
	if o == nil || IsNil(o.FailedOrders) {
		var ret []PrintOrdersLabelsResponseFailedOrder
		return ret
	}
	return o.FailedOrders
}

// GetFailedOrdersOk returns a tuple with the FailedOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPrintOrdersLabelsResponse) GetFailedOrdersOk() ([]PrintOrdersLabelsResponseFailedOrder, bool) {
	if o == nil || IsNil(o.FailedOrders) {
		return nil, false
	}
	return o.FailedOrders, true
}

// HasFailedOrders returns a boolean if a field has been set.
func (o *OrderPrintOrdersLabelsResponse) HasFailedOrders() bool {
	if o != nil && !IsNil(o.FailedOrders) {
		return true
	}

	return false
}

// SetFailedOrders gets a reference to the given []PrintOrdersLabelsResponseFailedOrder and assigns it to the FailedOrders field.
func (o *OrderPrintOrdersLabelsResponse) SetFailedOrders(v []PrintOrdersLabelsResponseFailedOrder) {
	o.FailedOrders = v
}

func (o OrderPrintOrdersLabelsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderPrintOrdersLabelsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownloadUrl) {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	if !IsNil(o.FailedOrders) {
		toSerialize["failedOrders"] = o.FailedOrders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderPrintOrdersLabelsResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderPrintOrdersLabelsResponse := _OrderPrintOrdersLabelsResponse{}

	err = json.Unmarshal(data, &varOrderPrintOrdersLabelsResponse)

	if err != nil {
		return err
	}

	*o = OrderPrintOrdersLabelsResponse(varOrderPrintOrdersLabelsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "downloadUrl")
		delete(additionalProperties, "failedOrders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderPrintOrdersLabelsResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderPrintOrdersLabelsResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderPrintOrdersLabelsResponse struct {
	value *OrderPrintOrdersLabelsResponse
	isSet bool
}

func (v NullableOrderPrintOrdersLabelsResponse) Get() *OrderPrintOrdersLabelsResponse {
	return v.value
}

func (v *NullableOrderPrintOrdersLabelsResponse) Set(val *OrderPrintOrdersLabelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderPrintOrdersLabelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderPrintOrdersLabelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderPrintOrdersLabelsResponse(val *OrderPrintOrdersLabelsResponse) *NullableOrderPrintOrdersLabelsResponse {
	return &NullableOrderPrintOrdersLabelsResponse{value: val, isSet: true}
}

func (v NullableOrderPrintOrdersLabelsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderPrintOrdersLabelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



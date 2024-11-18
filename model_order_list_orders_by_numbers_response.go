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

// checks if the OrderListOrdersByNumbersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderListOrdersByNumbersResponse{}

// OrderListOrdersByNumbersResponse struct for OrderListOrdersByNumbersResponse
type OrderListOrdersByNumbersResponse struct {
	Orders []OrderOrderData `json:"orders,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
	NextPageToken        *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListOrdersByNumbersResponse OrderListOrdersByNumbersResponse

// NewOrderListOrdersByNumbersResponse instantiates a new OrderListOrdersByNumbersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListOrdersByNumbersResponse() *OrderListOrdersByNumbersResponse {
	this := OrderListOrdersByNumbersResponse{}
	return &this
}

// NewOrderListOrdersByNumbersResponseWithDefaults instantiates a new OrderListOrdersByNumbersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListOrdersByNumbersResponseWithDefaults() *OrderListOrdersByNumbersResponse {
	this := OrderListOrdersByNumbersResponse{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *OrderListOrdersByNumbersResponse) GetOrders() []OrderOrderData {
	if o == nil || IsNil(o.Orders) {
		var ret []OrderOrderData
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOrdersByNumbersResponse) GetOrdersOk() ([]OrderOrderData, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *OrderListOrdersByNumbersResponse) IsSetOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderOrderData and assigns it to the Orders field.
func (o *OrderListOrdersByNumbersResponse) SetOrders(v []OrderOrderData) {
	o.Orders = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *OrderListOrdersByNumbersResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOrdersByNumbersResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *OrderListOrdersByNumbersResponse) IsSetNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *OrderListOrdersByNumbersResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o OrderListOrdersByNumbersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListOrdersByNumbersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderListOrdersByNumbersResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderListOrdersByNumbersResponse := _OrderListOrdersByNumbersResponse{}

	err = json.Unmarshal(data, &varOrderListOrdersByNumbersResponse)

	if err != nil {
		return err
	}

	*o = OrderListOrdersByNumbersResponse(varOrderListOrdersByNumbersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orders")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderListOrdersByNumbersResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderListOrdersByNumbersResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderListOrdersByNumbersResponse struct {
	value *OrderListOrdersByNumbersResponse
	isSet bool
}

func (v NullableOrderListOrdersByNumbersResponse) Get() *OrderListOrdersByNumbersResponse {
	return v.value
}

func (v *NullableOrderListOrdersByNumbersResponse) Set(val *OrderListOrdersByNumbersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListOrdersByNumbersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListOrdersByNumbersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListOrdersByNumbersResponse(val *OrderListOrdersByNumbersResponse) *NullableOrderListOrdersByNumbersResponse {
	return &NullableOrderListOrdersByNumbersResponse{value: val, isSet: true}
}

func (v NullableOrderListOrdersByNumbersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListOrdersByNumbersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

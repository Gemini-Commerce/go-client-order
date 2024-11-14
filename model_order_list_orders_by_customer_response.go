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

// checks if the OrderListOrdersByCustomerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderListOrdersByCustomerResponse{}

// OrderListOrdersByCustomerResponse struct for OrderListOrdersByCustomerResponse
type OrderListOrdersByCustomerResponse struct {
	Orders []OrderOrderData `json:"orders,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListOrdersByCustomerResponse OrderListOrdersByCustomerResponse

// NewOrderListOrdersByCustomerResponse instantiates a new OrderListOrdersByCustomerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListOrdersByCustomerResponse() *OrderListOrdersByCustomerResponse {
	this := OrderListOrdersByCustomerResponse{}
	return &this
}

// NewOrderListOrdersByCustomerResponseWithDefaults instantiates a new OrderListOrdersByCustomerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListOrdersByCustomerResponseWithDefaults() *OrderListOrdersByCustomerResponse {
	this := OrderListOrdersByCustomerResponse{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *OrderListOrdersByCustomerResponse) GetOrders() []OrderOrderData {
	if o == nil || IsNil(o.Orders) {
		var ret []OrderOrderData
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOrdersByCustomerResponse) GetOrdersOk() ([]OrderOrderData, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// IsSetOrders returns a boolean if a field has been set.
func (o *OrderListOrdersByCustomerResponse) IsSetOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderOrderData and assigns it to the Orders field.
func (o *OrderListOrdersByCustomerResponse) SetOrders(v []OrderOrderData) {
	o.Orders = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *OrderListOrdersByCustomerResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOrdersByCustomerResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// IsSetNextPageToken returns a boolean if a field has been set.
func (o *OrderListOrdersByCustomerResponse) IsSetNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *OrderListOrdersByCustomerResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o OrderListOrdersByCustomerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListOrdersByCustomerResponse) ToMap() (map[string]interface{}, error) {
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

func (o *OrderListOrdersByCustomerResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderListOrdersByCustomerResponse := _OrderListOrdersByCustomerResponse{}

	err = json.Unmarshal(data, &varOrderListOrdersByCustomerResponse)

	if err != nil {
		return err
	}

	*o = OrderListOrdersByCustomerResponse(varOrderListOrdersByCustomerResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orders")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderListOrdersByCustomerResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderListOrdersByCustomerResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderListOrdersByCustomerResponse struct {
	value *OrderListOrdersByCustomerResponse
	isSet bool
}

func (v NullableOrderListOrdersByCustomerResponse) Get() *OrderListOrdersByCustomerResponse {
	return v.value
}

func (v *NullableOrderListOrdersByCustomerResponse) Set(val *OrderListOrdersByCustomerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListOrdersByCustomerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListOrdersByCustomerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListOrdersByCustomerResponse(val *OrderListOrdersByCustomerResponse) *NullableOrderListOrdersByCustomerResponse {
	return &NullableOrderListOrdersByCustomerResponse{value: val, isSet: true}
}

func (v NullableOrderListOrdersByCustomerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListOrdersByCustomerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the OrderSearchOrdersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSearchOrdersResponse{}

// OrderSearchOrdersResponse struct for OrderSearchOrdersResponse
type OrderSearchOrdersResponse struct {
	Orders []OrderOrderData `json:"orders,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderSearchOrdersResponse OrderSearchOrdersResponse

// NewOrderSearchOrdersResponse instantiates a new OrderSearchOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchOrdersResponse() *OrderSearchOrdersResponse {
	this := OrderSearchOrdersResponse{}
	return &this
}

// NewOrderSearchOrdersResponseWithDefaults instantiates a new OrderSearchOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchOrdersResponseWithDefaults() *OrderSearchOrdersResponse {
	this := OrderSearchOrdersResponse{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *OrderSearchOrdersResponse) GetOrders() []OrderOrderData {
	if o == nil || IsNil(o.Orders) {
		var ret []OrderOrderData
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersResponse) GetOrdersOk() ([]OrderOrderData, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// &#39;Has&#39;Orders returns a boolean if a field has been set.
func (o *OrderSearchOrdersResponse) &#39;Has&#39;Orders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderOrderData and assigns it to the Orders field.
func (o *OrderSearchOrdersResponse) SetOrders(v []OrderOrderData) {
	o.Orders = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *OrderSearchOrdersResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// &#39;Has&#39;NextPageToken returns a boolean if a field has been set.
func (o *OrderSearchOrdersResponse) &#39;Has&#39;NextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *OrderSearchOrdersResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o OrderSearchOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSearchOrdersResponse) ToMap() (map[string]interface{}, error) {
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

func (o *OrderSearchOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderSearchOrdersResponse := _OrderSearchOrdersResponse{}

	err = json.Unmarshal(data, &varOrderSearchOrdersResponse)

	if err != nil {
		return err
	}

	*o = OrderSearchOrdersResponse(varOrderSearchOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orders")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderSearchOrdersResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderSearchOrdersResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderSearchOrdersResponse struct {
	value *OrderSearchOrdersResponse
	isSet bool
}

func (v NullableOrderSearchOrdersResponse) Get() *OrderSearchOrdersResponse {
	return v.value
}

func (v *NullableOrderSearchOrdersResponse) Set(val *OrderSearchOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchOrdersResponse(val *OrderSearchOrdersResponse) *NullableOrderSearchOrdersResponse {
	return &NullableOrderSearchOrdersResponse{value: val, isSet: true}
}

func (v NullableOrderSearchOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



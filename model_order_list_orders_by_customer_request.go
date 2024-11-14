/*
order Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OrderListOrdersByCustomerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderListOrdersByCustomerRequest{}

// OrderListOrdersByCustomerRequest struct for OrderListOrdersByCustomerRequest
type OrderListOrdersByCustomerRequest struct {
	TenantId string `json:"tenantId"`
	CustomerGrn string `json:"customerGrn"`
	// The maximum number of orders to return. The service may return fewer than this value. If unspecified, at most 10 orders will be returned. The maximum value is 100; values above 100 will be coerced to 100.
	PageSize *int64 `json:"pageSize,omitempty"`
	// A page token, received from a previous `ListOrders` call. Provide this to retrieve the subsequent page.   When paginating, all other parameters provided to `ListOrders` must match the call that provided the page token.
	PageToken *string `json:"pageToken,omitempty"`
	OrderBy []OrderOrderBy `json:"orderBy,omitempty"`
}

type _OrderListOrdersByCustomerRequest OrderListOrdersByCustomerRequest

// NewOrderListOrdersByCustomerRequest instantiates a new OrderListOrdersByCustomerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListOrdersByCustomerRequest(tenantId string, customerGrn string) *OrderListOrdersByCustomerRequest {
	this := OrderListOrdersByCustomerRequest{}
	this.TenantId = tenantId
	this.CustomerGrn = customerGrn
	return &this
}

// NewOrderListOrdersByCustomerRequestWithDefaults instantiates a new OrderListOrdersByCustomerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListOrdersByCustomerRequestWithDefaults() *OrderListOrdersByCustomerRequest {
	this := OrderListOrdersByCustomerRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderListOrdersByCustomerRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderListOrdersByCustomerRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderListOrdersByCustomerRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerGrn returns the CustomerGrn field value
func (o *OrderListOrdersByCustomerRequest) GetCustomerGrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerGrn
}

// GetCustomerGrnOk returns a tuple with the CustomerGrn field value
// and a boolean to check if the value has been set.
func (o *OrderListOrdersByCustomerRequest) GetCustomerGrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerGrn, true
}

// SetCustomerGrn sets field value
func (o *OrderListOrdersByCustomerRequest) SetCustomerGrn(v string) {
	o.CustomerGrn = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *OrderListOrdersByCustomerRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOrdersByCustomerRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// PageSize returns a boolean if a field has been set.
func (o *OrderListOrdersByCustomerRequest) PageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *OrderListOrdersByCustomerRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *OrderListOrdersByCustomerRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOrdersByCustomerRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// PageToken returns a boolean if a field has been set.
func (o *OrderListOrdersByCustomerRequest) PageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *OrderListOrdersByCustomerRequest) SetPageToken(v string) {
	o.PageToken = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *OrderListOrdersByCustomerRequest) GetOrderBy() []OrderOrderBy {
	if o == nil || IsNil(o.OrderBy) {
		var ret []OrderOrderBy
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListOrdersByCustomerRequest) GetOrderByOk() ([]OrderOrderBy, bool) {
	if o == nil || IsNil(o.OrderBy) {
		return nil, false
	}
	return o.OrderBy, true
}

// OrderBy returns a boolean if a field has been set.
func (o *OrderListOrdersByCustomerRequest) OrderBy() bool {
	if o != nil && !IsNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given []OrderOrderBy and assigns it to the OrderBy field.
func (o *OrderListOrdersByCustomerRequest) SetOrderBy(v []OrderOrderBy) {
	o.OrderBy = v
}

func (o OrderListOrdersByCustomerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListOrdersByCustomerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["customerGrn"] = o.CustomerGrn
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	if !IsNil(o.OrderBy) {
		toSerialize["orderBy"] = o.OrderBy
	}
	return toSerialize, nil
}

func (o *OrderListOrdersByCustomerRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"customerGrn",
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

	varOrderListOrdersByCustomerRequest := _OrderListOrdersByCustomerRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderListOrdersByCustomerRequest)

	if err != nil {
		return err
	}

	*o = OrderListOrdersByCustomerRequest(varOrderListOrdersByCustomerRequest)

	return err
}

type NullableOrderListOrdersByCustomerRequest struct {
	value *OrderListOrdersByCustomerRequest
	isSet bool
}

func (v NullableOrderListOrdersByCustomerRequest) Get() *OrderListOrdersByCustomerRequest {
	return v.value
}

func (v *NullableOrderListOrdersByCustomerRequest) Set(val *OrderListOrdersByCustomerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListOrdersByCustomerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListOrdersByCustomerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListOrdersByCustomerRequest(val *OrderListOrdersByCustomerRequest) *NullableOrderListOrdersByCustomerRequest {
	return &NullableOrderListOrdersByCustomerRequest{value: val, isSet: true}
}

func (v NullableOrderListOrdersByCustomerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListOrdersByCustomerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



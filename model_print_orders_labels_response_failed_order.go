/*
order/order.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PrintOrdersLabelsResponseFailedOrder struct for PrintOrdersLabelsResponseFailedOrder
type PrintOrdersLabelsResponseFailedOrder struct {
	OrderNumber *string `json:"orderNumber,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty"`
}

// NewPrintOrdersLabelsResponseFailedOrder instantiates a new PrintOrdersLabelsResponseFailedOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrintOrdersLabelsResponseFailedOrder() *PrintOrdersLabelsResponseFailedOrder {
	this := PrintOrdersLabelsResponseFailedOrder{}
	return &this
}

// NewPrintOrdersLabelsResponseFailedOrderWithDefaults instantiates a new PrintOrdersLabelsResponseFailedOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrintOrdersLabelsResponseFailedOrderWithDefaults() *PrintOrdersLabelsResponseFailedOrder {
	this := PrintOrdersLabelsResponseFailedOrder{}
	return &this
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *PrintOrdersLabelsResponseFailedOrder) GetOrderNumber() string {
	if o == nil || isNil(o.OrderNumber) {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintOrdersLabelsResponseFailedOrder) GetOrderNumberOk() (*string, bool) {
	if o == nil || isNil(o.OrderNumber) {
    return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *PrintOrdersLabelsResponseFailedOrder) HasOrderNumber() bool {
	if o != nil && !isNil(o.OrderNumber) {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *PrintOrdersLabelsResponseFailedOrder) SetOrderNumber(v string) {
	o.OrderNumber = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *PrintOrdersLabelsResponseFailedOrder) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintOrdersLabelsResponseFailedOrder) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
    return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *PrintOrdersLabelsResponseFailedOrder) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *PrintOrdersLabelsResponseFailedOrder) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *PrintOrdersLabelsResponseFailedOrder) GetErrorCode() string {
	if o == nil || isNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintOrdersLabelsResponseFailedOrder) GetErrorCodeOk() (*string, bool) {
	if o == nil || isNil(o.ErrorCode) {
    return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *PrintOrdersLabelsResponseFailedOrder) HasErrorCode() bool {
	if o != nil && !isNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *PrintOrdersLabelsResponseFailedOrder) SetErrorCode(v string) {
	o.ErrorCode = &v
}

func (o PrintOrdersLabelsResponseFailedOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OrderNumber) {
		toSerialize["orderNumber"] = o.OrderNumber
	}
	if !isNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !isNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	return json.Marshal(toSerialize)
}

type NullablePrintOrdersLabelsResponseFailedOrder struct {
	value *PrintOrdersLabelsResponseFailedOrder
	isSet bool
}

func (v NullablePrintOrdersLabelsResponseFailedOrder) Get() *PrintOrdersLabelsResponseFailedOrder {
	return v.value
}

func (v *NullablePrintOrdersLabelsResponseFailedOrder) Set(val *PrintOrdersLabelsResponseFailedOrder) {
	v.value = val
	v.isSet = true
}

func (v NullablePrintOrdersLabelsResponseFailedOrder) IsSet() bool {
	return v.isSet
}

func (v *NullablePrintOrdersLabelsResponseFailedOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrintOrdersLabelsResponseFailedOrder(val *PrintOrdersLabelsResponseFailedOrder) *NullablePrintOrdersLabelsResponseFailedOrder {
	return &NullablePrintOrdersLabelsResponseFailedOrder{value: val, isSet: true}
}

func (v NullablePrintOrdersLabelsResponseFailedOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrintOrdersLabelsResponseFailedOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// OrderCreatePaymentRequest struct for OrderCreatePaymentRequest
type OrderCreatePaymentRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	OrderId *string `json:"orderId,omitempty"`
	Code *string `json:"code,omitempty"`
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
	Amount *OrderMoney `json:"amount,omitempty"`
	CcInfo *PaymentCcInfo `json:"ccInfo,omitempty"`
}

// NewOrderCreatePaymentRequest instantiates a new OrderCreatePaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreatePaymentRequest() *OrderCreatePaymentRequest {
	this := OrderCreatePaymentRequest{}
	return &this
}

// NewOrderCreatePaymentRequestWithDefaults instantiates a new OrderCreatePaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreatePaymentRequestWithDefaults() *OrderCreatePaymentRequest {
	this := OrderCreatePaymentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderCreatePaymentRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderCreatePaymentRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderCreatePaymentRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderCreatePaymentRequest) GetOrderId() string {
	if o == nil || isNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentRequest) GetOrderIdOk() (*string, bool) {
	if o == nil || isNil(o.OrderId) {
    return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderCreatePaymentRequest) HasOrderId() bool {
	if o != nil && !isNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderCreatePaymentRequest) SetOrderId(v string) {
	o.OrderId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OrderCreatePaymentRequest) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentRequest) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OrderCreatePaymentRequest) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OrderCreatePaymentRequest) SetCode(v string) {
	o.Code = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *OrderCreatePaymentRequest) GetAdditionalInfo() string {
	if o == nil || isNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentRequest) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalInfo) {
    return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *OrderCreatePaymentRequest) HasAdditionalInfo() bool {
	if o != nil && !isNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *OrderCreatePaymentRequest) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *OrderCreatePaymentRequest) GetAmount() OrderMoney {
	if o == nil || isNil(o.Amount) {
		var ret OrderMoney
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentRequest) GetAmountOk() (*OrderMoney, bool) {
	if o == nil || isNil(o.Amount) {
    return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *OrderCreatePaymentRequest) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given OrderMoney and assigns it to the Amount field.
func (o *OrderCreatePaymentRequest) SetAmount(v OrderMoney) {
	o.Amount = &v
}

// GetCcInfo returns the CcInfo field value if set, zero value otherwise.
func (o *OrderCreatePaymentRequest) GetCcInfo() PaymentCcInfo {
	if o == nil || isNil(o.CcInfo) {
		var ret PaymentCcInfo
		return ret
	}
	return *o.CcInfo
}

// GetCcInfoOk returns a tuple with the CcInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentRequest) GetCcInfoOk() (*PaymentCcInfo, bool) {
	if o == nil || isNil(o.CcInfo) {
    return nil, false
	}
	return o.CcInfo, true
}

// HasCcInfo returns a boolean if a field has been set.
func (o *OrderCreatePaymentRequest) HasCcInfo() bool {
	if o != nil && !isNil(o.CcInfo) {
		return true
	}

	return false
}

// SetCcInfo gets a reference to the given PaymentCcInfo and assigns it to the CcInfo field.
func (o *OrderCreatePaymentRequest) SetCcInfo(v PaymentCcInfo) {
	o.CcInfo = &v
}

func (o OrderCreatePaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.CcInfo) {
		toSerialize["ccInfo"] = o.CcInfo
	}
	return json.Marshal(toSerialize)
}

type NullableOrderCreatePaymentRequest struct {
	value *OrderCreatePaymentRequest
	isSet bool
}

func (v NullableOrderCreatePaymentRequest) Get() *OrderCreatePaymentRequest {
	return v.value
}

func (v *NullableOrderCreatePaymentRequest) Set(val *OrderCreatePaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreatePaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreatePaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreatePaymentRequest(val *OrderCreatePaymentRequest) *NullableOrderCreatePaymentRequest {
	return &NullableOrderCreatePaymentRequest{value: val, isSet: true}
}

func (v NullableOrderCreatePaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreatePaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


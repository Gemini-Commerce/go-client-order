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

// OrderCreatePaymentTransactionRequest struct for OrderCreatePaymentTransactionRequest
type OrderCreatePaymentTransactionRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	PaymentId *string `json:"paymentId,omitempty"`
	Type *OrderTransactionType `json:"type,omitempty"`
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
}

// NewOrderCreatePaymentTransactionRequest instantiates a new OrderCreatePaymentTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreatePaymentTransactionRequest() *OrderCreatePaymentTransactionRequest {
	this := OrderCreatePaymentTransactionRequest{}
	var type_ OrderTransactionType = ORDERTRANSACTIONTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// NewOrderCreatePaymentTransactionRequestWithDefaults instantiates a new OrderCreatePaymentTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreatePaymentTransactionRequestWithDefaults() *OrderCreatePaymentTransactionRequest {
	this := OrderCreatePaymentTransactionRequest{}
	var type_ OrderTransactionType = ORDERTRANSACTIONTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderCreatePaymentTransactionRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentTransactionRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderCreatePaymentTransactionRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderCreatePaymentTransactionRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *OrderCreatePaymentTransactionRequest) GetPaymentId() string {
	if o == nil || isNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentTransactionRequest) GetPaymentIdOk() (*string, bool) {
	if o == nil || isNil(o.PaymentId) {
    return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *OrderCreatePaymentTransactionRequest) HasPaymentId() bool {
	if o != nil && !isNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *OrderCreatePaymentTransactionRequest) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderCreatePaymentTransactionRequest) GetType() OrderTransactionType {
	if o == nil || isNil(o.Type) {
		var ret OrderTransactionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentTransactionRequest) GetTypeOk() (*OrderTransactionType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderCreatePaymentTransactionRequest) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given OrderTransactionType and assigns it to the Type field.
func (o *OrderCreatePaymentTransactionRequest) SetType(v OrderTransactionType) {
	o.Type = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *OrderCreatePaymentTransactionRequest) GetAdditionalInfo() string {
	if o == nil || isNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentTransactionRequest) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalInfo) {
    return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *OrderCreatePaymentTransactionRequest) HasAdditionalInfo() bool {
	if o != nil && !isNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *OrderCreatePaymentTransactionRequest) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

func (o OrderCreatePaymentTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	return json.Marshal(toSerialize)
}

type NullableOrderCreatePaymentTransactionRequest struct {
	value *OrderCreatePaymentTransactionRequest
	isSet bool
}

func (v NullableOrderCreatePaymentTransactionRequest) Get() *OrderCreatePaymentTransactionRequest {
	return v.value
}

func (v *NullableOrderCreatePaymentTransactionRequest) Set(val *OrderCreatePaymentTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreatePaymentTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreatePaymentTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreatePaymentTransactionRequest(val *OrderCreatePaymentTransactionRequest) *NullableOrderCreatePaymentTransactionRequest {
	return &NullableOrderCreatePaymentTransactionRequest{value: val, isSet: true}
}

func (v NullableOrderCreatePaymentTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreatePaymentTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


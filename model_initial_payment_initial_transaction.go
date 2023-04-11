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

// InitialPaymentInitialTransaction struct for InitialPaymentInitialTransaction
type InitialPaymentInitialTransaction struct {
	Type *OrderTransactionType `json:"type,omitempty"`
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
}

// NewInitialPaymentInitialTransaction instantiates a new InitialPaymentInitialTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitialPaymentInitialTransaction() *InitialPaymentInitialTransaction {
	this := InitialPaymentInitialTransaction{}
	var type_ OrderTransactionType = ORDERTRANSACTIONTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// NewInitialPaymentInitialTransactionWithDefaults instantiates a new InitialPaymentInitialTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitialPaymentInitialTransactionWithDefaults() *InitialPaymentInitialTransaction {
	this := InitialPaymentInitialTransaction{}
	var type_ OrderTransactionType = ORDERTRANSACTIONTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InitialPaymentInitialTransaction) GetType() OrderTransactionType {
	if o == nil || isNil(o.Type) {
		var ret OrderTransactionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialPaymentInitialTransaction) GetTypeOk() (*OrderTransactionType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InitialPaymentInitialTransaction) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given OrderTransactionType and assigns it to the Type field.
func (o *InitialPaymentInitialTransaction) SetType(v OrderTransactionType) {
	o.Type = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *InitialPaymentInitialTransaction) GetAdditionalInfo() string {
	if o == nil || isNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialPaymentInitialTransaction) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalInfo) {
    return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *InitialPaymentInitialTransaction) HasAdditionalInfo() bool {
	if o != nil && !isNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *InitialPaymentInitialTransaction) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

func (o InitialPaymentInitialTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	return json.Marshal(toSerialize)
}

type NullableInitialPaymentInitialTransaction struct {
	value *InitialPaymentInitialTransaction
	isSet bool
}

func (v NullableInitialPaymentInitialTransaction) Get() *InitialPaymentInitialTransaction {
	return v.value
}

func (v *NullableInitialPaymentInitialTransaction) Set(val *InitialPaymentInitialTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableInitialPaymentInitialTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableInitialPaymentInitialTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitialPaymentInitialTransaction(val *InitialPaymentInitialTransaction) *NullableInitialPaymentInitialTransaction {
	return &NullableInitialPaymentInitialTransaction{value: val, isSet: true}
}

func (v NullableInitialPaymentInitialTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitialPaymentInitialTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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
	"bytes"
	"fmt"
)

// checks if the CreateOrderRequestInitialPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrderRequestInitialPayment{}

// CreateOrderRequestInitialPayment struct for CreateOrderRequestInitialPayment
type CreateOrderRequestInitialPayment struct {
	Code string `json:"code"`
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
	Amount OrderMoney `json:"amount"`
	CcInfo *PaymentCcInfo `json:"ccInfo,omitempty"`
	Transaction *InitialPaymentInitialTransaction `json:"transaction,omitempty"`
}

type _CreateOrderRequestInitialPayment CreateOrderRequestInitialPayment

// NewCreateOrderRequestInitialPayment instantiates a new CreateOrderRequestInitialPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrderRequestInitialPayment(code string, amount OrderMoney) *CreateOrderRequestInitialPayment {
	this := CreateOrderRequestInitialPayment{}
	this.Code = code
	this.Amount = amount
	return &this
}

// NewCreateOrderRequestInitialPaymentWithDefaults instantiates a new CreateOrderRequestInitialPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrderRequestInitialPaymentWithDefaults() *CreateOrderRequestInitialPayment {
	this := CreateOrderRequestInitialPayment{}
	return &this
}

// GetCode returns the Code field value
func (o *CreateOrderRequestInitialPayment) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CreateOrderRequestInitialPayment) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *CreateOrderRequestInitialPayment) SetCode(v string) {
	o.Code = v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *CreateOrderRequestInitialPayment) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequestInitialPayment) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *CreateOrderRequestInitialPayment) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *CreateOrderRequestInitialPayment) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetAmount returns the Amount field value
func (o *CreateOrderRequestInitialPayment) GetAmount() OrderMoney {
	if o == nil {
		var ret OrderMoney
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateOrderRequestInitialPayment) GetAmountOk() (*OrderMoney, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateOrderRequestInitialPayment) SetAmount(v OrderMoney) {
	o.Amount = v
}

// GetCcInfo returns the CcInfo field value if set, zero value otherwise.
func (o *CreateOrderRequestInitialPayment) GetCcInfo() PaymentCcInfo {
	if o == nil || IsNil(o.CcInfo) {
		var ret PaymentCcInfo
		return ret
	}
	return *o.CcInfo
}

// GetCcInfoOk returns a tuple with the CcInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequestInitialPayment) GetCcInfoOk() (*PaymentCcInfo, bool) {
	if o == nil || IsNil(o.CcInfo) {
		return nil, false
	}
	return o.CcInfo, true
}

// HasCcInfo returns a boolean if a field has been set.
func (o *CreateOrderRequestInitialPayment) HasCcInfo() bool {
	if o != nil && !IsNil(o.CcInfo) {
		return true
	}

	return false
}

// SetCcInfo gets a reference to the given PaymentCcInfo and assigns it to the CcInfo field.
func (o *CreateOrderRequestInitialPayment) SetCcInfo(v PaymentCcInfo) {
	o.CcInfo = &v
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *CreateOrderRequestInitialPayment) GetTransaction() InitialPaymentInitialTransaction {
	if o == nil || IsNil(o.Transaction) {
		var ret InitialPaymentInitialTransaction
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequestInitialPayment) GetTransactionOk() (*InitialPaymentInitialTransaction, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *CreateOrderRequestInitialPayment) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given InitialPaymentInitialTransaction and assigns it to the Transaction field.
func (o *CreateOrderRequestInitialPayment) SetTransaction(v InitialPaymentInitialTransaction) {
	o.Transaction = &v
}

func (o CreateOrderRequestInitialPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrderRequestInitialPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CcInfo) {
		toSerialize["ccInfo"] = o.CcInfo
	}
	if !IsNil(o.Transaction) {
		toSerialize["transaction"] = o.Transaction
	}
	return toSerialize, nil
}

func (o *CreateOrderRequestInitialPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"amount",
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

	varCreateOrderRequestInitialPayment := _CreateOrderRequestInitialPayment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrderRequestInitialPayment)

	if err != nil {
		return err
	}

	*o = CreateOrderRequestInitialPayment(varCreateOrderRequestInitialPayment)

	return err
}

type NullableCreateOrderRequestInitialPayment struct {
	value *CreateOrderRequestInitialPayment
	isSet bool
}

func (v NullableCreateOrderRequestInitialPayment) Get() *CreateOrderRequestInitialPayment {
	return v.value
}

func (v *NullableCreateOrderRequestInitialPayment) Set(val *CreateOrderRequestInitialPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrderRequestInitialPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrderRequestInitialPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrderRequestInitialPayment(val *CreateOrderRequestInitialPayment) *NullableCreateOrderRequestInitialPayment {
	return &NullableCreateOrderRequestInitialPayment{value: val, isSet: true}
}

func (v NullableCreateOrderRequestInitialPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrderRequestInitialPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



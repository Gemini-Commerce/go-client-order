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
	"fmt"
)

// checks if the ImportOrderRequestImportedPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportOrderRequestImportedPayment{}

// ImportOrderRequestImportedPayment struct for ImportOrderRequestImportedPayment
type ImportOrderRequestImportedPayment struct {
	Code                 string               `json:"code"`
	AdditionalInfo       *string              `json:"additionalInfo,omitempty"`
	Amounts              []OrderPaymentAmount `json:"amounts"`
	CcInfo               *PaymentCcInfo       `json:"ccInfo,omitempty"`
	IsUpfront            *bool                `json:"isUpfront,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportOrderRequestImportedPayment ImportOrderRequestImportedPayment

// NewImportOrderRequestImportedPayment instantiates a new ImportOrderRequestImportedPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportOrderRequestImportedPayment(code string, amounts []OrderPaymentAmount) *ImportOrderRequestImportedPayment {
	this := ImportOrderRequestImportedPayment{}
	this.Code = code
	this.Amounts = amounts
	return &this
}

// NewImportOrderRequestImportedPaymentWithDefaults instantiates a new ImportOrderRequestImportedPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportOrderRequestImportedPaymentWithDefaults() *ImportOrderRequestImportedPayment {
	this := ImportOrderRequestImportedPayment{}
	return &this
}

// GetCode returns the Code field value
func (o *ImportOrderRequestImportedPayment) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ImportOrderRequestImportedPayment) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ImportOrderRequestImportedPayment) SetCode(v string) {
	o.Code = v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *ImportOrderRequestImportedPayment) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportOrderRequestImportedPayment) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *ImportOrderRequestImportedPayment) IsSetAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *ImportOrderRequestImportedPayment) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetAmounts returns the Amounts field value
func (o *ImportOrderRequestImportedPayment) GetAmounts() []OrderPaymentAmount {
	if o == nil {
		var ret []OrderPaymentAmount
		return ret
	}

	return o.Amounts
}

// GetAmountsOk returns a tuple with the Amounts field value
// and a boolean to check if the value has been set.
func (o *ImportOrderRequestImportedPayment) GetAmountsOk() ([]OrderPaymentAmount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amounts, true
}

// SetAmounts sets field value
func (o *ImportOrderRequestImportedPayment) SetAmounts(v []OrderPaymentAmount) {
	o.Amounts = v
}

// GetCcInfo returns the CcInfo field value if set, zero value otherwise.
func (o *ImportOrderRequestImportedPayment) GetCcInfo() PaymentCcInfo {
	if o == nil || IsNil(o.CcInfo) {
		var ret PaymentCcInfo
		return ret
	}
	return *o.CcInfo
}

// GetCcInfoOk returns a tuple with the CcInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportOrderRequestImportedPayment) GetCcInfoOk() (*PaymentCcInfo, bool) {
	if o == nil || IsNil(o.CcInfo) {
		return nil, false
	}
	return o.CcInfo, true
}

// HasCcInfo returns a boolean if a field has been set.
func (o *ImportOrderRequestImportedPayment) IsSetCcInfo() bool {
	if o != nil && !IsNil(o.CcInfo) {
		return true
	}

	return false
}

// SetCcInfo gets a reference to the given PaymentCcInfo and assigns it to the CcInfo field.
func (o *ImportOrderRequestImportedPayment) SetCcInfo(v PaymentCcInfo) {
	o.CcInfo = &v
}

// GetIsUpfront returns the IsUpfront field value if set, zero value otherwise.
func (o *ImportOrderRequestImportedPayment) GetIsUpfront() bool {
	if o == nil || IsNil(o.IsUpfront) {
		var ret bool
		return ret
	}
	return *o.IsUpfront
}

// GetIsUpfrontOk returns a tuple with the IsUpfront field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportOrderRequestImportedPayment) GetIsUpfrontOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUpfront) {
		return nil, false
	}
	return o.IsUpfront, true
}

// HasIsUpfront returns a boolean if a field has been set.
func (o *ImportOrderRequestImportedPayment) IsSetIsUpfront() bool {
	if o != nil && !IsNil(o.IsUpfront) {
		return true
	}

	return false
}

// SetIsUpfront gets a reference to the given bool and assigns it to the IsUpfront field.
func (o *ImportOrderRequestImportedPayment) SetIsUpfront(v bool) {
	o.IsUpfront = &v
}

func (o ImportOrderRequestImportedPayment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportOrderRequestImportedPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	toSerialize["amounts"] = o.Amounts
	if !IsNil(o.CcInfo) {
		toSerialize["ccInfo"] = o.CcInfo
	}
	if !IsNil(o.IsUpfront) {
		toSerialize["isUpfront"] = o.IsUpfront
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportOrderRequestImportedPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"amounts",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varImportOrderRequestImportedPayment := _ImportOrderRequestImportedPayment{}

	err = json.Unmarshal(data, &varImportOrderRequestImportedPayment)

	if err != nil {
		return err
	}

	*o = ImportOrderRequestImportedPayment(varImportOrderRequestImportedPayment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "additionalInfo")
		delete(additionalProperties, "amounts")
		delete(additionalProperties, "ccInfo")
		delete(additionalProperties, "isUpfront")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ImportOrderRequestImportedPayment) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ImportOrderRequestImportedPayment) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableImportOrderRequestImportedPayment struct {
	value *ImportOrderRequestImportedPayment
	isSet bool
}

func (v NullableImportOrderRequestImportedPayment) Get() *ImportOrderRequestImportedPayment {
	return v.value
}

func (v *NullableImportOrderRequestImportedPayment) Set(val *ImportOrderRequestImportedPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableImportOrderRequestImportedPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableImportOrderRequestImportedPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportOrderRequestImportedPayment(val *ImportOrderRequestImportedPayment) *NullableImportOrderRequestImportedPayment {
	return &NullableImportOrderRequestImportedPayment{value: val, isSet: true}
}

func (v NullableImportOrderRequestImportedPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportOrderRequestImportedPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

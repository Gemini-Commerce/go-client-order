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

// checks if the OrderDataShipmentInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDataShipmentInfo{}

// OrderDataShipmentInfo struct for OrderDataShipmentInfo
type OrderDataShipmentInfo struct {
	Reference string `json:"reference"`
	Code string `json:"code"`
	Method *string `json:"method,omitempty"`
	Title *string `json:"title,omitempty"`
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
	Amount OrderMoney `json:"amount"`
	Fee *OrderMoney `json:"fee,omitempty"`
	VatAmount *OrderMoney `json:"vatAmount,omitempty"`
	VatPercentage *float32 `json:"vatPercentage,omitempty"`
	VatInaccurate *bool `json:"vatInaccurate,omitempty"`
	VatCalculated *bool `json:"vatCalculated,omitempty"`
	Grn *string `json:"grn,omitempty"`
	FromAddress *OrderPostalAddress `json:"fromAddress,omitempty"`
	ReturnAddress *OrderPostalAddress `json:"returnAddress,omitempty"`
}

type _OrderDataShipmentInfo OrderDataShipmentInfo

// NewOrderDataShipmentInfo instantiates a new OrderDataShipmentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataShipmentInfo(reference string, code string, amount OrderMoney) *OrderDataShipmentInfo {
	this := OrderDataShipmentInfo{}
	this.Reference = reference
	this.Code = code
	this.Amount = amount
	return &this
}

// NewOrderDataShipmentInfoWithDefaults instantiates a new OrderDataShipmentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataShipmentInfoWithDefaults() *OrderDataShipmentInfo {
	this := OrderDataShipmentInfo{}
	return &this
}

// GetReference returns the Reference field value
func (o *OrderDataShipmentInfo) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *OrderDataShipmentInfo) SetReference(v string) {
	o.Reference = v
}

// GetCode returns the Code field value
func (o *OrderDataShipmentInfo) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OrderDataShipmentInfo) SetCode(v string) {
	o.Code = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *OrderDataShipmentInfo) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// Method returns a boolean if a field has been set.
func (o *OrderDataShipmentInfo) Method() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *OrderDataShipmentInfo) SetMethod(v string) {
	o.Method = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *OrderDataShipmentInfo) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// Title returns a boolean if a field has been set.
func (o *OrderDataShipmentInfo) Title() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *OrderDataShipmentInfo) SetTitle(v string) {
	o.Title = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *OrderDataShipmentInfo) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// AdditionalInfo returns a boolean if a field has been set.
func (o *OrderDataShipmentInfo) AdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *OrderDataShipmentInfo) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetAmount returns the Amount field value
func (o *OrderDataShipmentInfo) GetAmount() OrderMoney {
	if o == nil {
		var ret OrderMoney
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetAmountOk() (*OrderMoney, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *OrderDataShipmentInfo) SetAmount(v OrderMoney) {
	o.Amount = v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *OrderDataShipmentInfo) GetFee() OrderMoney {
	if o == nil || IsNil(o.Fee) {
		var ret OrderMoney
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetFeeOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// Fee returns a boolean if a field has been set.
func (o *OrderDataShipmentInfo) Fee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given OrderMoney and assigns it to the Fee field.
func (o *OrderDataShipmentInfo) SetFee(v OrderMoney) {
	o.Fee = &v
}

// GetVatAmount returns the VatAmount field value if set, zero value otherwise.
func (o *OrderDataShipmentInfo) GetVatAmount() OrderMoney {
	if o == nil || IsNil(o.VatAmount) {
		var ret OrderMoney
		return ret
	}
	return *o.VatAmount
}

// GetVatAmountOk returns a tuple with the VatAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetVatAmountOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.VatAmount) {
		return nil, false
	}
	return o.VatAmount, true
}

// VatAmount returns a boolean if a field has been set.
func (o *OrderDataShipmentInfo) VatAmount() bool {
	if o != nil && !IsNil(o.VatAmount) {
		return true
	}

	return false
}

// SetVatAmount gets a reference to the given OrderMoney and assigns it to the VatAmount field.
func (o *OrderDataShipmentInfo) SetVatAmount(v OrderMoney) {
	o.VatAmount = &v
}

// GetVatPercentage returns the VatPercentage field value if set, zero value otherwise.
func (o *OrderDataShipmentInfo) GetVatPercentage() float32 {
	if o == nil || IsNil(o.VatPercentage) {
		var ret float32
		return ret
	}
	return *o.VatPercentage
}

// GetVatPercentageOk returns a tuple with the VatPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetVatPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.VatPercentage) {
		return nil, false
	}
	return o.VatPercentage, true
}

// VatPercentage returns a boolean if a field has been set.
func (o *OrderDataShipmentInfo) VatPercentage() bool {
	if o != nil && !IsNil(o.VatPercentage) {
		return true
	}

	return false
}

// SetVatPercentage gets a reference to the given float32 and assigns it to the VatPercentage field.
func (o *OrderDataShipmentInfo) SetVatPercentage(v float32) {
	o.VatPercentage = &v
}

// GetVatInaccurate returns the VatInaccurate field value if set, zero value otherwise.
func (o *OrderDataShipmentInfo) GetVatInaccurate() bool {
	if o == nil || IsNil(o.VatInaccurate) {
		var ret bool
		return ret
	}
	return *o.VatInaccurate
}

// GetVatInaccurateOk returns a tuple with the VatInaccurate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetVatInaccurateOk() (*bool, bool) {
	if o == nil || IsNil(o.VatInaccurate) {
		return nil, false
	}
	return o.VatInaccurate, true
}

// VatInaccurate returns a boolean if a field has been set.
func (o *OrderDataShipmentInfo) VatInaccurate() bool {
	if o != nil && !IsNil(o.VatInaccurate) {
		return true
	}

	return false
}

// SetVatInaccurate gets a reference to the given bool and assigns it to the VatInaccurate field.
func (o *OrderDataShipmentInfo) SetVatInaccurate(v bool) {
	o.VatInaccurate = &v
}

// GetVatCalculated returns the VatCalculated field value if set, zero value otherwise.
func (o *OrderDataShipmentInfo) GetVatCalculated() bool {
	if o == nil || IsNil(o.VatCalculated) {
		var ret bool
		return ret
	}
	return *o.VatCalculated
}

// GetVatCalculatedOk returns a tuple with the VatCalculated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetVatCalculatedOk() (*bool, bool) {
	if o == nil || IsNil(o.VatCalculated) {
		return nil, false
	}
	return o.VatCalculated, true
}

// VatCalculated returns a boolean if a field has been set.
func (o *OrderDataShipmentInfo) VatCalculated() bool {
	if o != nil && !IsNil(o.VatCalculated) {
		return true
	}

	return false
}

// SetVatCalculated gets a reference to the given bool and assigns it to the VatCalculated field.
func (o *OrderDataShipmentInfo) SetVatCalculated(v bool) {
	o.VatCalculated = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *OrderDataShipmentInfo) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// Grn returns a boolean if a field has been set.
func (o *OrderDataShipmentInfo) Grn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *OrderDataShipmentInfo) SetGrn(v string) {
	o.Grn = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *OrderDataShipmentInfo) GetFromAddress() OrderPostalAddress {
	if o == nil || IsNil(o.FromAddress) {
		var ret OrderPostalAddress
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetFromAddressOk() (*OrderPostalAddress, bool) {
	if o == nil || IsNil(o.FromAddress) {
		return nil, false
	}
	return o.FromAddress, true
}

// FromAddress returns a boolean if a field has been set.
func (o *OrderDataShipmentInfo) FromAddress() bool {
	if o != nil && !IsNil(o.FromAddress) {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given OrderPostalAddress and assigns it to the FromAddress field.
func (o *OrderDataShipmentInfo) SetFromAddress(v OrderPostalAddress) {
	o.FromAddress = &v
}

// GetReturnAddress returns the ReturnAddress field value if set, zero value otherwise.
func (o *OrderDataShipmentInfo) GetReturnAddress() OrderPostalAddress {
	if o == nil || IsNil(o.ReturnAddress) {
		var ret OrderPostalAddress
		return ret
	}
	return *o.ReturnAddress
}

// GetReturnAddressOk returns a tuple with the ReturnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataShipmentInfo) GetReturnAddressOk() (*OrderPostalAddress, bool) {
	if o == nil || IsNil(o.ReturnAddress) {
		return nil, false
	}
	return o.ReturnAddress, true
}

// ReturnAddress returns a boolean if a field has been set.
func (o *OrderDataShipmentInfo) ReturnAddress() bool {
	if o != nil && !IsNil(o.ReturnAddress) {
		return true
	}

	return false
}

// SetReturnAddress gets a reference to the given OrderPostalAddress and assigns it to the ReturnAddress field.
func (o *OrderDataShipmentInfo) SetReturnAddress(v OrderPostalAddress) {
	o.ReturnAddress = &v
}

func (o OrderDataShipmentInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDataShipmentInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reference"] = o.Reference
	toSerialize["code"] = o.Code
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.VatAmount) {
		toSerialize["vatAmount"] = o.VatAmount
	}
	if !IsNil(o.VatPercentage) {
		toSerialize["vatPercentage"] = o.VatPercentage
	}
	if !IsNil(o.VatInaccurate) {
		toSerialize["vatInaccurate"] = o.VatInaccurate
	}
	if !IsNil(o.VatCalculated) {
		toSerialize["vatCalculated"] = o.VatCalculated
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.FromAddress) {
		toSerialize["fromAddress"] = o.FromAddress
	}
	if !IsNil(o.ReturnAddress) {
		toSerialize["returnAddress"] = o.ReturnAddress
	}
	return toSerialize, nil
}

func (o *OrderDataShipmentInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reference",
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

	varOrderDataShipmentInfo := _OrderDataShipmentInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderDataShipmentInfo)

	if err != nil {
		return err
	}

	*o = OrderDataShipmentInfo(varOrderDataShipmentInfo)

	return err
}

type NullableOrderDataShipmentInfo struct {
	value *OrderDataShipmentInfo
	isSet bool
}

func (v NullableOrderDataShipmentInfo) Get() *OrderDataShipmentInfo {
	return v.value
}

func (v *NullableOrderDataShipmentInfo) Set(val *OrderDataShipmentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataShipmentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataShipmentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataShipmentInfo(val *OrderDataShipmentInfo) *NullableOrderDataShipmentInfo {
	return &NullableOrderDataShipmentInfo{value: val, isSet: true}
}

func (v NullableOrderDataShipmentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataShipmentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



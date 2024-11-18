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

// checks if the OrderDataPromotionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDataPromotionInfo{}

// OrderDataPromotionInfo struct for OrderDataPromotionInfo
type OrderDataPromotionInfo struct {
	PromotionGrn         *string     `json:"promotionGrn,omitempty"`
	Type                 string      `json:"type"`
	AdditionalInfo       *string     `json:"additionalInfo,omitempty"`
	Name                 string      `json:"name"`
	Description          *string     `json:"description,omitempty"`
	Amount               OrderMoney  `json:"amount"`
	CouponCode           *string     `json:"couponCode,omitempty"`
	VatAmount            *OrderMoney `json:"vatAmount,omitempty"`
	VatPercentage        *float32    `json:"vatPercentage,omitempty"`
	VatInaccurate        *bool       `json:"vatInaccurate,omitempty"`
	VatCalculated        *bool       `json:"vatCalculated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderDataPromotionInfo OrderDataPromotionInfo

// NewOrderDataPromotionInfo instantiates a new OrderDataPromotionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataPromotionInfo(type_ string, name string, amount OrderMoney) *OrderDataPromotionInfo {
	this := OrderDataPromotionInfo{}
	this.Type = type_
	this.Name = name
	this.Amount = amount
	return &this
}

// NewOrderDataPromotionInfoWithDefaults instantiates a new OrderDataPromotionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataPromotionInfoWithDefaults() *OrderDataPromotionInfo {
	this := OrderDataPromotionInfo{}
	return &this
}

// GetPromotionGrn returns the PromotionGrn field value if set, zero value otherwise.
func (o *OrderDataPromotionInfo) GetPromotionGrn() string {
	if o == nil || IsNil(o.PromotionGrn) {
		var ret string
		return ret
	}
	return *o.PromotionGrn
}

// GetPromotionGrnOk returns a tuple with the PromotionGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataPromotionInfo) GetPromotionGrnOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionGrn) {
		return nil, false
	}
	return o.PromotionGrn, true
}

// HasPromotionGrn returns a boolean if a field has been set.
func (o *OrderDataPromotionInfo) IsSetPromotionGrn() bool {
	if o != nil && !IsNil(o.PromotionGrn) {
		return true
	}

	return false
}

// SetPromotionGrn gets a reference to the given string and assigns it to the PromotionGrn field.
func (o *OrderDataPromotionInfo) SetPromotionGrn(v string) {
	o.PromotionGrn = &v
}

// GetType returns the Type field value
func (o *OrderDataPromotionInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderDataPromotionInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderDataPromotionInfo) SetType(v string) {
	o.Type = v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *OrderDataPromotionInfo) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataPromotionInfo) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *OrderDataPromotionInfo) IsSetAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *OrderDataPromotionInfo) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetName returns the Name field value
func (o *OrderDataPromotionInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrderDataPromotionInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrderDataPromotionInfo) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrderDataPromotionInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataPromotionInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrderDataPromotionInfo) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrderDataPromotionInfo) SetDescription(v string) {
	o.Description = &v
}

// GetAmount returns the Amount field value
func (o *OrderDataPromotionInfo) GetAmount() OrderMoney {
	if o == nil {
		var ret OrderMoney
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *OrderDataPromotionInfo) GetAmountOk() (*OrderMoney, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *OrderDataPromotionInfo) SetAmount(v OrderMoney) {
	o.Amount = v
}

// GetCouponCode returns the CouponCode field value if set, zero value otherwise.
func (o *OrderDataPromotionInfo) GetCouponCode() string {
	if o == nil || IsNil(o.CouponCode) {
		var ret string
		return ret
	}
	return *o.CouponCode
}

// GetCouponCodeOk returns a tuple with the CouponCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataPromotionInfo) GetCouponCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CouponCode) {
		return nil, false
	}
	return o.CouponCode, true
}

// HasCouponCode returns a boolean if a field has been set.
func (o *OrderDataPromotionInfo) IsSetCouponCode() bool {
	if o != nil && !IsNil(o.CouponCode) {
		return true
	}

	return false
}

// SetCouponCode gets a reference to the given string and assigns it to the CouponCode field.
func (o *OrderDataPromotionInfo) SetCouponCode(v string) {
	o.CouponCode = &v
}

// GetVatAmount returns the VatAmount field value if set, zero value otherwise.
func (o *OrderDataPromotionInfo) GetVatAmount() OrderMoney {
	if o == nil || IsNil(o.VatAmount) {
		var ret OrderMoney
		return ret
	}
	return *o.VatAmount
}

// GetVatAmountOk returns a tuple with the VatAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataPromotionInfo) GetVatAmountOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.VatAmount) {
		return nil, false
	}
	return o.VatAmount, true
}

// HasVatAmount returns a boolean if a field has been set.
func (o *OrderDataPromotionInfo) IsSetVatAmount() bool {
	if o != nil && !IsNil(o.VatAmount) {
		return true
	}

	return false
}

// SetVatAmount gets a reference to the given OrderMoney and assigns it to the VatAmount field.
func (o *OrderDataPromotionInfo) SetVatAmount(v OrderMoney) {
	o.VatAmount = &v
}

// GetVatPercentage returns the VatPercentage field value if set, zero value otherwise.
func (o *OrderDataPromotionInfo) GetVatPercentage() float32 {
	if o == nil || IsNil(o.VatPercentage) {
		var ret float32
		return ret
	}
	return *o.VatPercentage
}

// GetVatPercentageOk returns a tuple with the VatPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataPromotionInfo) GetVatPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.VatPercentage) {
		return nil, false
	}
	return o.VatPercentage, true
}

// HasVatPercentage returns a boolean if a field has been set.
func (o *OrderDataPromotionInfo) IsSetVatPercentage() bool {
	if o != nil && !IsNil(o.VatPercentage) {
		return true
	}

	return false
}

// SetVatPercentage gets a reference to the given float32 and assigns it to the VatPercentage field.
func (o *OrderDataPromotionInfo) SetVatPercentage(v float32) {
	o.VatPercentage = &v
}

// GetVatInaccurate returns the VatInaccurate field value if set, zero value otherwise.
func (o *OrderDataPromotionInfo) GetVatInaccurate() bool {
	if o == nil || IsNil(o.VatInaccurate) {
		var ret bool
		return ret
	}
	return *o.VatInaccurate
}

// GetVatInaccurateOk returns a tuple with the VatInaccurate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataPromotionInfo) GetVatInaccurateOk() (*bool, bool) {
	if o == nil || IsNil(o.VatInaccurate) {
		return nil, false
	}
	return o.VatInaccurate, true
}

// HasVatInaccurate returns a boolean if a field has been set.
func (o *OrderDataPromotionInfo) IsSetVatInaccurate() bool {
	if o != nil && !IsNil(o.VatInaccurate) {
		return true
	}

	return false
}

// SetVatInaccurate gets a reference to the given bool and assigns it to the VatInaccurate field.
func (o *OrderDataPromotionInfo) SetVatInaccurate(v bool) {
	o.VatInaccurate = &v
}

// GetVatCalculated returns the VatCalculated field value if set, zero value otherwise.
func (o *OrderDataPromotionInfo) GetVatCalculated() bool {
	if o == nil || IsNil(o.VatCalculated) {
		var ret bool
		return ret
	}
	return *o.VatCalculated
}

// GetVatCalculatedOk returns a tuple with the VatCalculated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataPromotionInfo) GetVatCalculatedOk() (*bool, bool) {
	if o == nil || IsNil(o.VatCalculated) {
		return nil, false
	}
	return o.VatCalculated, true
}

// HasVatCalculated returns a boolean if a field has been set.
func (o *OrderDataPromotionInfo) IsSetVatCalculated() bool {
	if o != nil && !IsNil(o.VatCalculated) {
		return true
	}

	return false
}

// SetVatCalculated gets a reference to the given bool and assigns it to the VatCalculated field.
func (o *OrderDataPromotionInfo) SetVatCalculated(v bool) {
	o.VatCalculated = &v
}

func (o OrderDataPromotionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDataPromotionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PromotionGrn) {
		toSerialize["promotionGrn"] = o.PromotionGrn
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CouponCode) {
		toSerialize["couponCode"] = o.CouponCode
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderDataPromotionInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"amount",
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

	varOrderDataPromotionInfo := _OrderDataPromotionInfo{}

	err = json.Unmarshal(data, &varOrderDataPromotionInfo)

	if err != nil {
		return err
	}

	*o = OrderDataPromotionInfo(varOrderDataPromotionInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "promotionGrn")
		delete(additionalProperties, "type")
		delete(additionalProperties, "additionalInfo")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "couponCode")
		delete(additionalProperties, "vatAmount")
		delete(additionalProperties, "vatPercentage")
		delete(additionalProperties, "vatInaccurate")
		delete(additionalProperties, "vatCalculated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderDataPromotionInfo) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderDataPromotionInfo) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderDataPromotionInfo struct {
	value *OrderDataPromotionInfo
	isSet bool
}

func (v NullableOrderDataPromotionInfo) Get() *OrderDataPromotionInfo {
	return v.value
}

func (v *NullableOrderDataPromotionInfo) Set(val *OrderDataPromotionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataPromotionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataPromotionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataPromotionInfo(val *OrderDataPromotionInfo) *NullableOrderDataPromotionInfo {
	return &NullableOrderDataPromotionInfo{value: val, isSet: true}
}

func (v NullableOrderDataPromotionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataPromotionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

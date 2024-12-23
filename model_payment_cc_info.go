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

// checks if the PaymentCcInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentCcInfo{}

// PaymentCcInfo struct for PaymentCcInfo
type PaymentCcInfo struct {
	Approval  *string `json:"approval,omitempty"`
	ExpMonth  *int32  `json:"expMonth,omitempty"`
	ExpYear   *string `json:"expYear,omitempty"`
	Last4     *string `json:"last4,omitempty"`
	NumberEnc *string `json:"numberEnc,omitempty"`
	Owner     *string `json:"owner,omitempty"`
	AvsStatus *string `json:"avsStatus,omitempty"`
	// card type MasterCard, Visa..
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentCcInfo PaymentCcInfo

// NewPaymentCcInfo instantiates a new PaymentCcInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentCcInfo() *PaymentCcInfo {
	this := PaymentCcInfo{}
	return &this
}

// NewPaymentCcInfoWithDefaults instantiates a new PaymentCcInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentCcInfoWithDefaults() *PaymentCcInfo {
	this := PaymentCcInfo{}
	return &this
}

// GetApproval returns the Approval field value if set, zero value otherwise.
func (o *PaymentCcInfo) GetApproval() string {
	if o == nil || IsNil(o.Approval) {
		var ret string
		return ret
	}
	return *o.Approval
}

// GetApprovalOk returns a tuple with the Approval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCcInfo) GetApprovalOk() (*string, bool) {
	if o == nil || IsNil(o.Approval) {
		return nil, false
	}
	return o.Approval, true
}

// HasApproval returns a boolean if a field has been set.
func (o *PaymentCcInfo) IsSetApproval() bool {
	if o != nil && !IsNil(o.Approval) {
		return true
	}

	return false
}

// SetApproval gets a reference to the given string and assigns it to the Approval field.
func (o *PaymentCcInfo) SetApproval(v string) {
	o.Approval = &v
}

// GetExpMonth returns the ExpMonth field value if set, zero value otherwise.
func (o *PaymentCcInfo) GetExpMonth() int32 {
	if o == nil || IsNil(o.ExpMonth) {
		var ret int32
		return ret
	}
	return *o.ExpMonth
}

// GetExpMonthOk returns a tuple with the ExpMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCcInfo) GetExpMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpMonth) {
		return nil, false
	}
	return o.ExpMonth, true
}

// HasExpMonth returns a boolean if a field has been set.
func (o *PaymentCcInfo) IsSetExpMonth() bool {
	if o != nil && !IsNil(o.ExpMonth) {
		return true
	}

	return false
}

// SetExpMonth gets a reference to the given int32 and assigns it to the ExpMonth field.
func (o *PaymentCcInfo) SetExpMonth(v int32) {
	o.ExpMonth = &v
}

// GetExpYear returns the ExpYear field value if set, zero value otherwise.
func (o *PaymentCcInfo) GetExpYear() string {
	if o == nil || IsNil(o.ExpYear) {
		var ret string
		return ret
	}
	return *o.ExpYear
}

// GetExpYearOk returns a tuple with the ExpYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCcInfo) GetExpYearOk() (*string, bool) {
	if o == nil || IsNil(o.ExpYear) {
		return nil, false
	}
	return o.ExpYear, true
}

// HasExpYear returns a boolean if a field has been set.
func (o *PaymentCcInfo) IsSetExpYear() bool {
	if o != nil && !IsNil(o.ExpYear) {
		return true
	}

	return false
}

// SetExpYear gets a reference to the given string and assigns it to the ExpYear field.
func (o *PaymentCcInfo) SetExpYear(v string) {
	o.ExpYear = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *PaymentCcInfo) GetLast4() string {
	if o == nil || IsNil(o.Last4) {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCcInfo) GetLast4Ok() (*string, bool) {
	if o == nil || IsNil(o.Last4) {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *PaymentCcInfo) IsSetLast4() bool {
	if o != nil && !IsNil(o.Last4) {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *PaymentCcInfo) SetLast4(v string) {
	o.Last4 = &v
}

// GetNumberEnc returns the NumberEnc field value if set, zero value otherwise.
func (o *PaymentCcInfo) GetNumberEnc() string {
	if o == nil || IsNil(o.NumberEnc) {
		var ret string
		return ret
	}
	return *o.NumberEnc
}

// GetNumberEncOk returns a tuple with the NumberEnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCcInfo) GetNumberEncOk() (*string, bool) {
	if o == nil || IsNil(o.NumberEnc) {
		return nil, false
	}
	return o.NumberEnc, true
}

// HasNumberEnc returns a boolean if a field has been set.
func (o *PaymentCcInfo) IsSetNumberEnc() bool {
	if o != nil && !IsNil(o.NumberEnc) {
		return true
	}

	return false
}

// SetNumberEnc gets a reference to the given string and assigns it to the NumberEnc field.
func (o *PaymentCcInfo) SetNumberEnc(v string) {
	o.NumberEnc = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *PaymentCcInfo) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCcInfo) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *PaymentCcInfo) IsSetOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *PaymentCcInfo) SetOwner(v string) {
	o.Owner = &v
}

// GetAvsStatus returns the AvsStatus field value if set, zero value otherwise.
func (o *PaymentCcInfo) GetAvsStatus() string {
	if o == nil || IsNil(o.AvsStatus) {
		var ret string
		return ret
	}
	return *o.AvsStatus
}

// GetAvsStatusOk returns a tuple with the AvsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCcInfo) GetAvsStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AvsStatus) {
		return nil, false
	}
	return o.AvsStatus, true
}

// HasAvsStatus returns a boolean if a field has been set.
func (o *PaymentCcInfo) IsSetAvsStatus() bool {
	if o != nil && !IsNil(o.AvsStatus) {
		return true
	}

	return false
}

// SetAvsStatus gets a reference to the given string and assigns it to the AvsStatus field.
func (o *PaymentCcInfo) SetAvsStatus(v string) {
	o.AvsStatus = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentCcInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCcInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentCcInfo) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentCcInfo) SetType(v string) {
	o.Type = &v
}

func (o PaymentCcInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentCcInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Approval) {
		toSerialize["approval"] = o.Approval
	}
	if !IsNil(o.ExpMonth) {
		toSerialize["expMonth"] = o.ExpMonth
	}
	if !IsNil(o.ExpYear) {
		toSerialize["expYear"] = o.ExpYear
	}
	if !IsNil(o.Last4) {
		toSerialize["last4"] = o.Last4
	}
	if !IsNil(o.NumberEnc) {
		toSerialize["numberEnc"] = o.NumberEnc
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.AvsStatus) {
		toSerialize["avsStatus"] = o.AvsStatus
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentCcInfo) UnmarshalJSON(data []byte) (err error) {
	varPaymentCcInfo := _PaymentCcInfo{}

	err = json.Unmarshal(data, &varPaymentCcInfo)

	if err != nil {
		return err
	}

	*o = PaymentCcInfo(varPaymentCcInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "approval")
		delete(additionalProperties, "expMonth")
		delete(additionalProperties, "expYear")
		delete(additionalProperties, "last4")
		delete(additionalProperties, "numberEnc")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "avsStatus")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PaymentCcInfo) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PaymentCcInfo) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePaymentCcInfo struct {
	value *PaymentCcInfo
	isSet bool
}

func (v NullablePaymentCcInfo) Get() *PaymentCcInfo {
	return v.value
}

func (v *NullablePaymentCcInfo) Set(val *PaymentCcInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentCcInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentCcInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentCcInfo(val *PaymentCcInfo) *NullablePaymentCcInfo {
	return &NullablePaymentCcInfo{value: val, isSet: true}
}

func (v NullablePaymentCcInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentCcInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

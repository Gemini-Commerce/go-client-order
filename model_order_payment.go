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

// checks if the OrderPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderPayment{}

// OrderPayment struct for OrderPayment
type OrderPayment struct {
	OrderId *string `json:"orderId,omitempty"`
	Id      *string `json:"id,omitempty"`
	// payment type stripe, paypal..
	Code                 *string              `json:"code,omitempty"`
	AdditionalInfo       *string              `json:"additionalInfo,omitempty"`
	Amounts              []OrderPaymentAmount `json:"amounts,omitempty"`
	CcInfo               *PaymentCcInfo       `json:"ccInfo,omitempty"`
	Transactions         []OrderTransaction   `json:"transactions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderPayment OrderPayment

// NewOrderPayment instantiates a new OrderPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderPayment() *OrderPayment {
	this := OrderPayment{}
	return &this
}

// NewOrderPaymentWithDefaults instantiates a new OrderPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderPaymentWithDefaults() *OrderPayment {
	this := OrderPayment{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderPayment) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPayment) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderPayment) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderPayment) SetOrderId(v string) {
	o.OrderId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderPayment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPayment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderPayment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderPayment) SetId(v string) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OrderPayment) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPayment) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OrderPayment) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OrderPayment) SetCode(v string) {
	o.Code = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *OrderPayment) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPayment) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *OrderPayment) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *OrderPayment) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetAmounts returns the Amounts field value if set, zero value otherwise.
func (o *OrderPayment) GetAmounts() []OrderPaymentAmount {
	if o == nil || IsNil(o.Amounts) {
		var ret []OrderPaymentAmount
		return ret
	}
	return o.Amounts
}

// GetAmountsOk returns a tuple with the Amounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPayment) GetAmountsOk() ([]OrderPaymentAmount, bool) {
	if o == nil || IsNil(o.Amounts) {
		return nil, false
	}
	return o.Amounts, true
}

// HasAmounts returns a boolean if a field has been set.
func (o *OrderPayment) HasAmounts() bool {
	if o != nil && !IsNil(o.Amounts) {
		return true
	}

	return false
}

// SetAmounts gets a reference to the given []OrderPaymentAmount and assigns it to the Amounts field.
func (o *OrderPayment) SetAmounts(v []OrderPaymentAmount) {
	o.Amounts = v
}

// GetCcInfo returns the CcInfo field value if set, zero value otherwise.
func (o *OrderPayment) GetCcInfo() PaymentCcInfo {
	if o == nil || IsNil(o.CcInfo) {
		var ret PaymentCcInfo
		return ret
	}
	return *o.CcInfo
}

// GetCcInfoOk returns a tuple with the CcInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPayment) GetCcInfoOk() (*PaymentCcInfo, bool) {
	if o == nil || IsNil(o.CcInfo) {
		return nil, false
	}
	return o.CcInfo, true
}

// HasCcInfo returns a boolean if a field has been set.
func (o *OrderPayment) HasCcInfo() bool {
	if o != nil && !IsNil(o.CcInfo) {
		return true
	}

	return false
}

// SetCcInfo gets a reference to the given PaymentCcInfo and assigns it to the CcInfo field.
func (o *OrderPayment) SetCcInfo(v PaymentCcInfo) {
	o.CcInfo = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *OrderPayment) GetTransactions() []OrderTransaction {
	if o == nil || IsNil(o.Transactions) {
		var ret []OrderTransaction
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPayment) GetTransactionsOk() ([]OrderTransaction, bool) {
	if o == nil || IsNil(o.Transactions) {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *OrderPayment) HasTransactions() bool {
	if o != nil && !IsNil(o.Transactions) {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []OrderTransaction and assigns it to the Transactions field.
func (o *OrderPayment) SetTransactions(v []OrderTransaction) {
	o.Transactions = v
}

func (o OrderPayment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	if !IsNil(o.Amounts) {
		toSerialize["amounts"] = o.Amounts
	}
	if !IsNil(o.CcInfo) {
		toSerialize["ccInfo"] = o.CcInfo
	}
	if !IsNil(o.Transactions) {
		toSerialize["transactions"] = o.Transactions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderPayment) UnmarshalJSON(data []byte) (err error) {
	varOrderPayment := _OrderPayment{}

	err = json.Unmarshal(data, &varOrderPayment)

	if err != nil {
		return err
	}

	*o = OrderPayment(varOrderPayment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "code")
		delete(additionalProperties, "additionalInfo")
		delete(additionalProperties, "amounts")
		delete(additionalProperties, "ccInfo")
		delete(additionalProperties, "transactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderPayment) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderPayment) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderPayment struct {
	value *OrderPayment
	isSet bool
}

func (v NullableOrderPayment) Get() *OrderPayment {
	return v.value
}

func (v *NullableOrderPayment) Set(val *OrderPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderPayment(val *OrderPayment) *NullableOrderPayment {
	return &NullableOrderPayment{value: val, isSet: true}
}

func (v NullableOrderPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

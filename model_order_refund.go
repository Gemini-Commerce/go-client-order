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
	"time"
)

// checks if the OrderRefund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderRefund{}

// OrderRefund struct for OrderRefund
type OrderRefund struct {
	CreatedAt            *time.Time          `json:"createdAt,omitempty"`
	PaymentId            *string             `json:"paymentId,omitempty"`
	Id                   *string             `json:"id,omitempty"`
	Items                []OrderRefundItem   `json:"items,omitempty"`
	Amounts              []OrderRefundAmount `json:"amounts,omitempty"`
	Note                 *string             `json:"note,omitempty"`
	AdditionalInfo       *string             `json:"additionalInfo,omitempty"`
	TransactionIds       []string            `json:"transactionIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderRefund OrderRefund

// NewOrderRefund instantiates a new OrderRefund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderRefund() *OrderRefund {
	this := OrderRefund{}
	return &this
}

// NewOrderRefundWithDefaults instantiates a new OrderRefund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderRefundWithDefaults() *OrderRefund {
	this := OrderRefund{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrderRefund) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefund) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrderRefund) IsSetCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrderRefund) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *OrderRefund) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefund) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *OrderRefund) IsSetPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *OrderRefund) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderRefund) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefund) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderRefund) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderRefund) SetId(v string) {
	o.Id = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *OrderRefund) GetItems() []OrderRefundItem {
	if o == nil || IsNil(o.Items) {
		var ret []OrderRefundItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefund) GetItemsOk() ([]OrderRefundItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *OrderRefund) IsSetItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrderRefundItem and assigns it to the Items field.
func (o *OrderRefund) SetItems(v []OrderRefundItem) {
	o.Items = v
}

// GetAmounts returns the Amounts field value if set, zero value otherwise.
func (o *OrderRefund) GetAmounts() []OrderRefundAmount {
	if o == nil || IsNil(o.Amounts) {
		var ret []OrderRefundAmount
		return ret
	}
	return o.Amounts
}

// GetAmountsOk returns a tuple with the Amounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefund) GetAmountsOk() ([]OrderRefundAmount, bool) {
	if o == nil || IsNil(o.Amounts) {
		return nil, false
	}
	return o.Amounts, true
}

// HasAmounts returns a boolean if a field has been set.
func (o *OrderRefund) IsSetAmounts() bool {
	if o != nil && !IsNil(o.Amounts) {
		return true
	}

	return false
}

// SetAmounts gets a reference to the given []OrderRefundAmount and assigns it to the Amounts field.
func (o *OrderRefund) SetAmounts(v []OrderRefundAmount) {
	o.Amounts = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *OrderRefund) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefund) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *OrderRefund) IsSetNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *OrderRefund) SetNote(v string) {
	o.Note = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *OrderRefund) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefund) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *OrderRefund) IsSetAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *OrderRefund) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetTransactionIds returns the TransactionIds field value if set, zero value otherwise.
func (o *OrderRefund) GetTransactionIds() []string {
	if o == nil || IsNil(o.TransactionIds) {
		var ret []string
		return ret
	}
	return o.TransactionIds
}

// GetTransactionIdsOk returns a tuple with the TransactionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefund) GetTransactionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TransactionIds) {
		return nil, false
	}
	return o.TransactionIds, true
}

// HasTransactionIds returns a boolean if a field has been set.
func (o *OrderRefund) IsSetTransactionIds() bool {
	if o != nil && !IsNil(o.TransactionIds) {
		return true
	}

	return false
}

// SetTransactionIds gets a reference to the given []string and assigns it to the TransactionIds field.
func (o *OrderRefund) SetTransactionIds(v []string) {
	o.TransactionIds = v
}

func (o OrderRefund) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderRefund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Amounts) {
		toSerialize["amounts"] = o.Amounts
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	if !IsNil(o.TransactionIds) {
		toSerialize["transactionIds"] = o.TransactionIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderRefund) UnmarshalJSON(data []byte) (err error) {
	varOrderRefund := _OrderRefund{}

	err = json.Unmarshal(data, &varOrderRefund)

	if err != nil {
		return err
	}

	*o = OrderRefund(varOrderRefund)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "paymentId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "items")
		delete(additionalProperties, "amounts")
		delete(additionalProperties, "note")
		delete(additionalProperties, "additionalInfo")
		delete(additionalProperties, "transactionIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderRefund) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderRefund) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderRefund struct {
	value *OrderRefund
	isSet bool
}

func (v NullableOrderRefund) Get() *OrderRefund {
	return v.value
}

func (v *NullableOrderRefund) Set(val *OrderRefund) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderRefund) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderRefund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderRefund(val *OrderRefund) *NullableOrderRefund {
	return &NullableOrderRefund{value: val, isSet: true}
}

func (v NullableOrderRefund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderRefund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

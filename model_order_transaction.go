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

// checks if the OrderTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderTransaction{}

// OrderTransaction struct for OrderTransaction
type OrderTransaction struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	PaymentId *string `json:"paymentId,omitempty"`
	Id *string `json:"id,omitempty"`
	Type *OrderTransactionType `json:"type,omitempty"`
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
	ChildTransactions []OrderTransaction `json:"childTransactions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderTransaction OrderTransaction

// NewOrderTransaction instantiates a new OrderTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTransaction() *OrderTransaction {
	this := OrderTransaction{}
	var type_ OrderTransactionType = ORDERTRANSACTIONTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// NewOrderTransactionWithDefaults instantiates a new OrderTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTransactionWithDefaults() *OrderTransaction {
	this := OrderTransaction{}
	var type_ OrderTransactionType = ORDERTRANSACTIONTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrderTransaction) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTransaction) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// &#39;Has&#39;CreatedAt returns a boolean if a field has been set.
func (o *OrderTransaction) &#39;Has&#39;CreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrderTransaction) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *OrderTransaction) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTransaction) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// &#39;Has&#39;PaymentId returns a boolean if a field has been set.
func (o *OrderTransaction) &#39;Has&#39;PaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *OrderTransaction) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderTransaction) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTransaction) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// &#39;Has&#39;Id returns a boolean if a field has been set.
func (o *OrderTransaction) &#39;Has&#39;Id() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderTransaction) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderTransaction) GetType() OrderTransactionType {
	if o == nil || IsNil(o.Type) {
		var ret OrderTransactionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTransaction) GetTypeOk() (*OrderTransactionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// &#39;Has&#39;Type returns a boolean if a field has been set.
func (o *OrderTransaction) &#39;Has&#39;Type() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given OrderTransactionType and assigns it to the Type field.
func (o *OrderTransaction) SetType(v OrderTransactionType) {
	o.Type = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *OrderTransaction) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTransaction) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// &#39;Has&#39;AdditionalInfo returns a boolean if a field has been set.
func (o *OrderTransaction) &#39;Has&#39;AdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *OrderTransaction) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetChildTransactions returns the ChildTransactions field value if set, zero value otherwise.
func (o *OrderTransaction) GetChildTransactions() []OrderTransaction {
	if o == nil || IsNil(o.ChildTransactions) {
		var ret []OrderTransaction
		return ret
	}
	return o.ChildTransactions
}

// GetChildTransactionsOk returns a tuple with the ChildTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTransaction) GetChildTransactionsOk() ([]OrderTransaction, bool) {
	if o == nil || IsNil(o.ChildTransactions) {
		return nil, false
	}
	return o.ChildTransactions, true
}

// &#39;Has&#39;ChildTransactions returns a boolean if a field has been set.
func (o *OrderTransaction) &#39;Has&#39;ChildTransactions() bool {
	if o != nil && !IsNil(o.ChildTransactions) {
		return true
	}

	return false
}

// SetChildTransactions gets a reference to the given []OrderTransaction and assigns it to the ChildTransactions field.
func (o *OrderTransaction) SetChildTransactions(v []OrderTransaction) {
	o.ChildTransactions = v
}

func (o OrderTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTransaction) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	if !IsNil(o.ChildTransactions) {
		toSerialize["childTransactions"] = o.ChildTransactions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderTransaction) UnmarshalJSON(data []byte) (err error) {
	varOrderTransaction := _OrderTransaction{}

	err = json.Unmarshal(data, &varOrderTransaction)

	if err != nil {
		return err
	}

	*o = OrderTransaction(varOrderTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "paymentId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "additionalInfo")
		delete(additionalProperties, "childTransactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderTransaction) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderTransaction) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderTransaction struct {
	value *OrderTransaction
	isSet bool
}

func (v NullableOrderTransaction) Get() *OrderTransaction {
	return v.value
}

func (v *NullableOrderTransaction) Set(val *OrderTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTransaction(val *OrderTransaction) *NullableOrderTransaction {
	return &NullableOrderTransaction{value: val, isSet: true}
}

func (v NullableOrderTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



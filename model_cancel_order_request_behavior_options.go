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

// checks if the CancelOrderRequestBehaviorOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelOrderRequestBehaviorOptions{}

// CancelOrderRequestBehaviorOptions struct for CancelOrderRequestBehaviorOptions
type CancelOrderRequestBehaviorOptions struct {
	Inventory            *BehaviorOptionsInventory                 `json:"inventory,omitempty"`
	Payment              *CancelOrderRequestBehaviorOptionsPayment `json:"payment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelOrderRequestBehaviorOptions CancelOrderRequestBehaviorOptions

// NewCancelOrderRequestBehaviorOptions instantiates a new CancelOrderRequestBehaviorOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelOrderRequestBehaviorOptions() *CancelOrderRequestBehaviorOptions {
	this := CancelOrderRequestBehaviorOptions{}
	return &this
}

// NewCancelOrderRequestBehaviorOptionsWithDefaults instantiates a new CancelOrderRequestBehaviorOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelOrderRequestBehaviorOptionsWithDefaults() *CancelOrderRequestBehaviorOptions {
	this := CancelOrderRequestBehaviorOptions{}
	return &this
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *CancelOrderRequestBehaviorOptions) GetInventory() BehaviorOptionsInventory {
	if o == nil || IsNil(o.Inventory) {
		var ret BehaviorOptionsInventory
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrderRequestBehaviorOptions) GetInventoryOk() (*BehaviorOptionsInventory, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *CancelOrderRequestBehaviorOptions) IsSetInventory() bool {
	if o != nil && !IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given BehaviorOptionsInventory and assigns it to the Inventory field.
func (o *CancelOrderRequestBehaviorOptions) SetInventory(v BehaviorOptionsInventory) {
	o.Inventory = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *CancelOrderRequestBehaviorOptions) GetPayment() CancelOrderRequestBehaviorOptionsPayment {
	if o == nil || IsNil(o.Payment) {
		var ret CancelOrderRequestBehaviorOptionsPayment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrderRequestBehaviorOptions) GetPaymentOk() (*CancelOrderRequestBehaviorOptionsPayment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *CancelOrderRequestBehaviorOptions) IsSetPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given CancelOrderRequestBehaviorOptionsPayment and assigns it to the Payment field.
func (o *CancelOrderRequestBehaviorOptions) SetPayment(v CancelOrderRequestBehaviorOptionsPayment) {
	o.Payment = &v
}

func (o CancelOrderRequestBehaviorOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelOrderRequestBehaviorOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Inventory) {
		toSerialize["inventory"] = o.Inventory
	}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CancelOrderRequestBehaviorOptions) UnmarshalJSON(data []byte) (err error) {
	varCancelOrderRequestBehaviorOptions := _CancelOrderRequestBehaviorOptions{}

	err = json.Unmarshal(data, &varCancelOrderRequestBehaviorOptions)

	if err != nil {
		return err
	}

	*o = CancelOrderRequestBehaviorOptions(varCancelOrderRequestBehaviorOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "inventory")
		delete(additionalProperties, "payment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CancelOrderRequestBehaviorOptions) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CancelOrderRequestBehaviorOptions) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCancelOrderRequestBehaviorOptions struct {
	value *CancelOrderRequestBehaviorOptions
	isSet bool
}

func (v NullableCancelOrderRequestBehaviorOptions) Get() *CancelOrderRequestBehaviorOptions {
	return v.value
}

func (v *NullableCancelOrderRequestBehaviorOptions) Set(val *CancelOrderRequestBehaviorOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelOrderRequestBehaviorOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelOrderRequestBehaviorOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelOrderRequestBehaviorOptions(val *CancelOrderRequestBehaviorOptions) *NullableCancelOrderRequestBehaviorOptions {
	return &NullableCancelOrderRequestBehaviorOptions{value: val, isSet: true}
}

func (v NullableCancelOrderRequestBehaviorOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelOrderRequestBehaviorOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

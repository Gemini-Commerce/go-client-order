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

// checks if the OrderCreatePaymentTransactionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreatePaymentTransactionRequest{}

// OrderCreatePaymentTransactionRequest struct for OrderCreatePaymentTransactionRequest
type OrderCreatePaymentTransactionRequest struct {
	TenantId string `json:"tenantId"`
	PaymentId string `json:"paymentId"`
	Type OrderTransactionType `json:"type"`
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderCreatePaymentTransactionRequest OrderCreatePaymentTransactionRequest

// NewOrderCreatePaymentTransactionRequest instantiates a new OrderCreatePaymentTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreatePaymentTransactionRequest(tenantId string, paymentId string, type_ OrderTransactionType) *OrderCreatePaymentTransactionRequest {
	this := OrderCreatePaymentTransactionRequest{}
	this.TenantId = tenantId
	this.PaymentId = paymentId
	this.Type = type_
	return &this
}

// NewOrderCreatePaymentTransactionRequestWithDefaults instantiates a new OrderCreatePaymentTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreatePaymentTransactionRequestWithDefaults() *OrderCreatePaymentTransactionRequest {
	this := OrderCreatePaymentTransactionRequest{}
	var type_ OrderTransactionType = ORDERTRANSACTIONTYPE_UNKNOWN
	this.Type = type_
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderCreatePaymentTransactionRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentTransactionRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderCreatePaymentTransactionRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetPaymentId returns the PaymentId field value
func (o *OrderCreatePaymentTransactionRequest) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentTransactionRequest) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *OrderCreatePaymentTransactionRequest) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetType returns the Type field value
func (o *OrderCreatePaymentTransactionRequest) GetType() OrderTransactionType {
	if o == nil {
		var ret OrderTransactionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentTransactionRequest) GetTypeOk() (*OrderTransactionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderCreatePaymentTransactionRequest) SetType(v OrderTransactionType) {
	o.Type = v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *OrderCreatePaymentTransactionRequest) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreatePaymentTransactionRequest) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// IsSetAdditionalInfo returns a boolean if a field has been set.
func (o *OrderCreatePaymentTransactionRequest) IsSetAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *OrderCreatePaymentTransactionRequest) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

func (o OrderCreatePaymentTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreatePaymentTransactionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["paymentId"] = o.PaymentId
	toSerialize["type"] = o.Type
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderCreatePaymentTransactionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"paymentId",
		"type",
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

	varOrderCreatePaymentTransactionRequest := _OrderCreatePaymentTransactionRequest{}

	err = json.Unmarshal(data, &varOrderCreatePaymentTransactionRequest)

	if err != nil {
		return err
	}

	*o = OrderCreatePaymentTransactionRequest(varOrderCreatePaymentTransactionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "paymentId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "additionalInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderCreatePaymentTransactionRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderCreatePaymentTransactionRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderCreatePaymentTransactionRequest struct {
	value *OrderCreatePaymentTransactionRequest
	isSet bool
}

func (v NullableOrderCreatePaymentTransactionRequest) Get() *OrderCreatePaymentTransactionRequest {
	return v.value
}

func (v *NullableOrderCreatePaymentTransactionRequest) Set(val *OrderCreatePaymentTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreatePaymentTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreatePaymentTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreatePaymentTransactionRequest(val *OrderCreatePaymentTransactionRequest) *NullableOrderCreatePaymentTransactionRequest {
	return &NullableOrderCreatePaymentTransactionRequest{value: val, isSet: true}
}

func (v NullableOrderCreatePaymentTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreatePaymentTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



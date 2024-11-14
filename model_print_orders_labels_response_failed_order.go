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

// checks if the PrintOrdersLabelsResponseFailedOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrintOrdersLabelsResponseFailedOrder{}

// PrintOrdersLabelsResponseFailedOrder struct for PrintOrdersLabelsResponseFailedOrder
type PrintOrdersLabelsResponseFailedOrder struct {
	OrderNumber *string `json:"orderNumber,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrintOrdersLabelsResponseFailedOrder PrintOrdersLabelsResponseFailedOrder

// NewPrintOrdersLabelsResponseFailedOrder instantiates a new PrintOrdersLabelsResponseFailedOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrintOrdersLabelsResponseFailedOrder() *PrintOrdersLabelsResponseFailedOrder {
	this := PrintOrdersLabelsResponseFailedOrder{}
	return &this
}

// NewPrintOrdersLabelsResponseFailedOrderWithDefaults instantiates a new PrintOrdersLabelsResponseFailedOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrintOrdersLabelsResponseFailedOrderWithDefaults() *PrintOrdersLabelsResponseFailedOrder {
	this := PrintOrdersLabelsResponseFailedOrder{}
	return &this
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *PrintOrdersLabelsResponseFailedOrder) GetOrderNumber() string {
	if o == nil || IsNil(o.OrderNumber) {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintOrdersLabelsResponseFailedOrder) GetOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNumber) {
		return nil, false
	}
	return o.OrderNumber, true
}

// IsSetOrderNumber returns a boolean if a field has been set.
func (o *PrintOrdersLabelsResponseFailedOrder) IsSetOrderNumber() bool {
	if o != nil && !IsNil(o.OrderNumber) {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *PrintOrdersLabelsResponseFailedOrder) SetOrderNumber(v string) {
	o.OrderNumber = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *PrintOrdersLabelsResponseFailedOrder) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintOrdersLabelsResponseFailedOrder) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// IsSetErrorMessage returns a boolean if a field has been set.
func (o *PrintOrdersLabelsResponseFailedOrder) IsSetErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *PrintOrdersLabelsResponseFailedOrder) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *PrintOrdersLabelsResponseFailedOrder) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintOrdersLabelsResponseFailedOrder) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// IsSetErrorCode returns a boolean if a field has been set.
func (o *PrintOrdersLabelsResponseFailedOrder) IsSetErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *PrintOrdersLabelsResponseFailedOrder) SetErrorCode(v string) {
	o.ErrorCode = &v
}

func (o PrintOrdersLabelsResponseFailedOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrintOrdersLabelsResponseFailedOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderNumber) {
		toSerialize["orderNumber"] = o.OrderNumber
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrintOrdersLabelsResponseFailedOrder) UnmarshalJSON(data []byte) (err error) {
	varPrintOrdersLabelsResponseFailedOrder := _PrintOrdersLabelsResponseFailedOrder{}

	err = json.Unmarshal(data, &varPrintOrdersLabelsResponseFailedOrder)

	if err != nil {
		return err
	}

	*o = PrintOrdersLabelsResponseFailedOrder(varPrintOrdersLabelsResponseFailedOrder)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderNumber")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "errorCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PrintOrdersLabelsResponseFailedOrder) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *PrintOrdersLabelsResponseFailedOrder) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullablePrintOrdersLabelsResponseFailedOrder struct {
	value *PrintOrdersLabelsResponseFailedOrder
	isSet bool
}

func (v NullablePrintOrdersLabelsResponseFailedOrder) Get() *PrintOrdersLabelsResponseFailedOrder {
	return v.value
}

func (v *NullablePrintOrdersLabelsResponseFailedOrder) Set(val *PrintOrdersLabelsResponseFailedOrder) {
	v.value = val
	v.isSet = true
}

func (v NullablePrintOrdersLabelsResponseFailedOrder) IsSet() bool {
	return v.isSet
}

func (v *NullablePrintOrdersLabelsResponseFailedOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrintOrdersLabelsResponseFailedOrder(val *PrintOrdersLabelsResponseFailedOrder) *NullablePrintOrdersLabelsResponseFailedOrder {
	return &NullablePrintOrdersLabelsResponseFailedOrder{value: val, isSet: true}
}

func (v NullablePrintOrdersLabelsResponseFailedOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrintOrdersLabelsResponseFailedOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



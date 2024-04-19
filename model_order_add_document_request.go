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
	"bytes"
	"fmt"
)

// checks if the OrderAddDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderAddDocumentRequest{}

// OrderAddDocumentRequest struct for OrderAddDocumentRequest
type OrderAddDocumentRequest struct {
	TenantId string `json:"tenantId"`
	OrderId string `json:"orderId"`
	Code string `json:"code"`
	Label *string `json:"label,omitempty"`
	AssetGrn *string `json:"assetGrn,omitempty"`
	Url *string `json:"url,omitempty"`
}

type _OrderAddDocumentRequest OrderAddDocumentRequest

// NewOrderAddDocumentRequest instantiates a new OrderAddDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAddDocumentRequest(tenantId string, orderId string, code string) *OrderAddDocumentRequest {
	this := OrderAddDocumentRequest{}
	this.TenantId = tenantId
	this.OrderId = orderId
	this.Code = code
	return &this
}

// NewOrderAddDocumentRequestWithDefaults instantiates a new OrderAddDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAddDocumentRequestWithDefaults() *OrderAddDocumentRequest {
	this := OrderAddDocumentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderAddDocumentRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderAddDocumentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderAddDocumentRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetOrderId returns the OrderId field value
func (o *OrderAddDocumentRequest) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *OrderAddDocumentRequest) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *OrderAddDocumentRequest) SetOrderId(v string) {
	o.OrderId = v
}

// GetCode returns the Code field value
func (o *OrderAddDocumentRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OrderAddDocumentRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OrderAddDocumentRequest) SetCode(v string) {
	o.Code = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *OrderAddDocumentRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddDocumentRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *OrderAddDocumentRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *OrderAddDocumentRequest) SetLabel(v string) {
	o.Label = &v
}

// GetAssetGrn returns the AssetGrn field value if set, zero value otherwise.
func (o *OrderAddDocumentRequest) GetAssetGrn() string {
	if o == nil || IsNil(o.AssetGrn) {
		var ret string
		return ret
	}
	return *o.AssetGrn
}

// GetAssetGrnOk returns a tuple with the AssetGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddDocumentRequest) GetAssetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.AssetGrn) {
		return nil, false
	}
	return o.AssetGrn, true
}

// HasAssetGrn returns a boolean if a field has been set.
func (o *OrderAddDocumentRequest) HasAssetGrn() bool {
	if o != nil && !IsNil(o.AssetGrn) {
		return true
	}

	return false
}

// SetAssetGrn gets a reference to the given string and assigns it to the AssetGrn field.
func (o *OrderAddDocumentRequest) SetAssetGrn(v string) {
	o.AssetGrn = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OrderAddDocumentRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddDocumentRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OrderAddDocumentRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OrderAddDocumentRequest) SetUrl(v string) {
	o.Url = &v
}

func (o OrderAddDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAddDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["orderId"] = o.OrderId
	toSerialize["code"] = o.Code
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.AssetGrn) {
		toSerialize["assetGrn"] = o.AssetGrn
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

func (o *OrderAddDocumentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"orderId",
		"code",
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

	varOrderAddDocumentRequest := _OrderAddDocumentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderAddDocumentRequest)

	if err != nil {
		return err
	}

	*o = OrderAddDocumentRequest(varOrderAddDocumentRequest)

	return err
}

type NullableOrderAddDocumentRequest struct {
	value *OrderAddDocumentRequest
	isSet bool
}

func (v NullableOrderAddDocumentRequest) Get() *OrderAddDocumentRequest {
	return v.value
}

func (v *NullableOrderAddDocumentRequest) Set(val *OrderAddDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAddDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAddDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAddDocumentRequest(val *OrderAddDocumentRequest) *NullableOrderAddDocumentRequest {
	return &NullableOrderAddDocumentRequest{value: val, isSet: true}
}

func (v NullableOrderAddDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAddDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// checks if the OrderDataDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDataDocument{}

// OrderDataDocument struct for OrderDataDocument
type OrderDataDocument struct {
	Code *string `json:"code,omitempty"`
	Label *string `json:"label,omitempty"`
	AssetGrn *string `json:"assetGrn,omitempty"`
	Url *string `json:"url,omitempty"`
	InsertedAt *time.Time `json:"insertedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderDataDocument OrderDataDocument

// NewOrderDataDocument instantiates a new OrderDataDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataDocument() *OrderDataDocument {
	this := OrderDataDocument{}
	return &this
}

// NewOrderDataDocumentWithDefaults instantiates a new OrderDataDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataDocumentWithDefaults() *OrderDataDocument {
	this := OrderDataDocument{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OrderDataDocument) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataDocument) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// &#39;Has&#39;Code returns a boolean if a field has been set.
func (o *OrderDataDocument) &#39;Has&#39;Code() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OrderDataDocument) SetCode(v string) {
	o.Code = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *OrderDataDocument) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataDocument) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// &#39;Has&#39;Label returns a boolean if a field has been set.
func (o *OrderDataDocument) &#39;Has&#39;Label() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *OrderDataDocument) SetLabel(v string) {
	o.Label = &v
}

// GetAssetGrn returns the AssetGrn field value if set, zero value otherwise.
func (o *OrderDataDocument) GetAssetGrn() string {
	if o == nil || IsNil(o.AssetGrn) {
		var ret string
		return ret
	}
	return *o.AssetGrn
}

// GetAssetGrnOk returns a tuple with the AssetGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataDocument) GetAssetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.AssetGrn) {
		return nil, false
	}
	return o.AssetGrn, true
}

// &#39;Has&#39;AssetGrn returns a boolean if a field has been set.
func (o *OrderDataDocument) &#39;Has&#39;AssetGrn() bool {
	if o != nil && !IsNil(o.AssetGrn) {
		return true
	}

	return false
}

// SetAssetGrn gets a reference to the given string and assigns it to the AssetGrn field.
func (o *OrderDataDocument) SetAssetGrn(v string) {
	o.AssetGrn = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OrderDataDocument) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataDocument) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// &#39;Has&#39;Url returns a boolean if a field has been set.
func (o *OrderDataDocument) &#39;Has&#39;Url() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OrderDataDocument) SetUrl(v string) {
	o.Url = &v
}

// GetInsertedAt returns the InsertedAt field value if set, zero value otherwise.
func (o *OrderDataDocument) GetInsertedAt() time.Time {
	if o == nil || IsNil(o.InsertedAt) {
		var ret time.Time
		return ret
	}
	return *o.InsertedAt
}

// GetInsertedAtOk returns a tuple with the InsertedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataDocument) GetInsertedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InsertedAt) {
		return nil, false
	}
	return o.InsertedAt, true
}

// &#39;Has&#39;InsertedAt returns a boolean if a field has been set.
func (o *OrderDataDocument) &#39;Has&#39;InsertedAt() bool {
	if o != nil && !IsNil(o.InsertedAt) {
		return true
	}

	return false
}

// SetInsertedAt gets a reference to the given time.Time and assigns it to the InsertedAt field.
func (o *OrderDataDocument) SetInsertedAt(v time.Time) {
	o.InsertedAt = &v
}

func (o OrderDataDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDataDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.AssetGrn) {
		toSerialize["assetGrn"] = o.AssetGrn
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.InsertedAt) {
		toSerialize["insertedAt"] = o.InsertedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderDataDocument) UnmarshalJSON(data []byte) (err error) {
	varOrderDataDocument := _OrderDataDocument{}

	err = json.Unmarshal(data, &varOrderDataDocument)

	if err != nil {
		return err
	}

	*o = OrderDataDocument(varOrderDataDocument)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "label")
		delete(additionalProperties, "assetGrn")
		delete(additionalProperties, "url")
		delete(additionalProperties, "insertedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderDataDocument) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderDataDocument) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderDataDocument struct {
	value *OrderDataDocument
	isSet bool
}

func (v NullableOrderDataDocument) Get() *OrderDataDocument {
	return v.value
}

func (v *NullableOrderDataDocument) Set(val *OrderDataDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataDocument(val *OrderDataDocument) *NullableOrderDataDocument {
	return &NullableOrderDataDocument{value: val, isSet: true}
}

func (v NullableOrderDataDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



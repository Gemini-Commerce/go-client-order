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

// checks if the OptionImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionImage{}

// OptionImage struct for OptionImage
type OptionImage struct {
	Grn                  *string `json:"grn,omitempty"`
	Url                  *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OptionImage OptionImage

// NewOptionImage instantiates a new OptionImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionImage() *OptionImage {
	this := OptionImage{}
	return &this
}

// NewOptionImageWithDefaults instantiates a new OptionImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionImageWithDefaults() *OptionImage {
	this := OptionImage{}
	return &this
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *OptionImage) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionImage) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *OptionImage) IsSetGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *OptionImage) SetGrn(v string) {
	o.Grn = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OptionImage) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionImage) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OptionImage) IsSetUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OptionImage) SetUrl(v string) {
	o.Url = &v
}

func (o OptionImage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OptionImage) UnmarshalJSON(data []byte) (err error) {
	varOptionImage := _OptionImage{}

	err = json.Unmarshal(data, &varOptionImage)

	if err != nil {
		return err
	}

	*o = OptionImage(varOptionImage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grn")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OptionImage) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OptionImage) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOptionImage struct {
	value *OptionImage
	isSet bool
}

func (v NullableOptionImage) Get() *OptionImage {
	return v.value
}

func (v *NullableOptionImage) Set(val *OptionImage) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionImage) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionImage(val *OptionImage) *NullableOptionImage {
	return &NullableOptionImage{value: val, isSet: true}
}

func (v NullableOptionImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

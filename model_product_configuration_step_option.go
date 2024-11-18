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

// checks if the ProductConfigurationStepOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfigurationStepOption{}

// ProductConfigurationStepOption struct for ProductConfigurationStepOption
type ProductConfigurationStepOption struct {
	Id                   *string      `json:"id,omitempty"`
	Grn                  *string      `json:"grn,omitempty"`
	Label                *string      `json:"label,omitempty"`
	PriceVariation       *OrderMoney  `json:"priceVariation,omitempty"`
	Image                *OptionImage `json:"image,omitempty"`
	HasQuantity          *bool        `json:"hasQuantity,omitempty"`
	Quantity             *int64       `json:"quantity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductConfigurationStepOption ProductConfigurationStepOption

// NewProductConfigurationStepOption instantiates a new ProductConfigurationStepOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfigurationStepOption() *ProductConfigurationStepOption {
	this := ProductConfigurationStepOption{}
	return &this
}

// NewProductConfigurationStepOptionWithDefaults instantiates a new ProductConfigurationStepOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfigurationStepOptionWithDefaults() *ProductConfigurationStepOption {
	this := ProductConfigurationStepOption{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductConfigurationStepOption) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfigurationStepOption) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductConfigurationStepOption) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductConfigurationStepOption) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ProductConfigurationStepOption) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfigurationStepOption) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ProductConfigurationStepOption) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ProductConfigurationStepOption) SetGrn(v string) {
	o.Grn = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductConfigurationStepOption) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfigurationStepOption) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductConfigurationStepOption) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ProductConfigurationStepOption) SetLabel(v string) {
	o.Label = &v
}

// GetPriceVariation returns the PriceVariation field value if set, zero value otherwise.
func (o *ProductConfigurationStepOption) GetPriceVariation() OrderMoney {
	if o == nil || IsNil(o.PriceVariation) {
		var ret OrderMoney
		return ret
	}
	return *o.PriceVariation
}

// GetPriceVariationOk returns a tuple with the PriceVariation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfigurationStepOption) GetPriceVariationOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.PriceVariation) {
		return nil, false
	}
	return o.PriceVariation, true
}

// HasPriceVariation returns a boolean if a field has been set.
func (o *ProductConfigurationStepOption) HasPriceVariation() bool {
	if o != nil && !IsNil(o.PriceVariation) {
		return true
	}

	return false
}

// SetPriceVariation gets a reference to the given OrderMoney and assigns it to the PriceVariation field.
func (o *ProductConfigurationStepOption) SetPriceVariation(v OrderMoney) {
	o.PriceVariation = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ProductConfigurationStepOption) GetImage() OptionImage {
	if o == nil || IsNil(o.Image) {
		var ret OptionImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfigurationStepOption) GetImageOk() (*OptionImage, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ProductConfigurationStepOption) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given OptionImage and assigns it to the Image field.
func (o *ProductConfigurationStepOption) SetImage(v OptionImage) {
	o.Image = &v
}

// GetHasQuantity returns the HasQuantity field value if set, zero value otherwise.
func (o *ProductConfigurationStepOption) GetHasQuantity() bool {
	if o == nil || IsNil(o.HasQuantity) {
		var ret bool
		return ret
	}
	return *o.HasQuantity
}

// GetHasQuantityOk returns a tuple with the HasQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfigurationStepOption) GetHasQuantityOk() (*bool, bool) {
	if o == nil || IsNil(o.HasQuantity) {
		return nil, false
	}
	return o.HasQuantity, true
}

// HasHasQuantity returns a boolean if a field has been set.
func (o *ProductConfigurationStepOption) HasHasQuantity() bool {
	if o != nil && !IsNil(o.HasQuantity) {
		return true
	}

	return false
}

// SetHasQuantity gets a reference to the given bool and assigns it to the HasQuantity field.
func (o *ProductConfigurationStepOption) SetHasQuantity(v bool) {
	o.HasQuantity = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ProductConfigurationStepOption) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfigurationStepOption) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ProductConfigurationStepOption) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *ProductConfigurationStepOption) SetQuantity(v int64) {
	o.Quantity = &v
}

func (o ProductConfigurationStepOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfigurationStepOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.PriceVariation) {
		toSerialize["priceVariation"] = o.PriceVariation
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.HasQuantity) {
		toSerialize["hasQuantity"] = o.HasQuantity
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductConfigurationStepOption) UnmarshalJSON(data []byte) (err error) {
	varProductConfigurationStepOption := _ProductConfigurationStepOption{}

	err = json.Unmarshal(data, &varProductConfigurationStepOption)

	if err != nil {
		return err
	}

	*o = ProductConfigurationStepOption(varProductConfigurationStepOption)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "grn")
		delete(additionalProperties, "label")
		delete(additionalProperties, "priceVariation")
		delete(additionalProperties, "image")
		delete(additionalProperties, "hasQuantity")
		delete(additionalProperties, "quantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductConfigurationStepOption) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductConfigurationStepOption) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductConfigurationStepOption struct {
	value *ProductConfigurationStepOption
	isSet bool
}

func (v NullableProductConfigurationStepOption) Get() *ProductConfigurationStepOption {
	return v.value
}

func (v *NullableProductConfigurationStepOption) Set(val *ProductConfigurationStepOption) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfigurationStepOption) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfigurationStepOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfigurationStepOption(val *ProductConfigurationStepOption) *NullableProductConfigurationStepOption {
	return &NullableProductConfigurationStepOption{value: val, isSet: true}
}

func (v NullableProductConfigurationStepOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfigurationStepOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

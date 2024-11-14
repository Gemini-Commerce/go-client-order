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

// checks if the ShipmentTracking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentTracking{}

// ShipmentTracking struct for ShipmentTracking
type ShipmentTracking struct {
	CarrierCode *string `json:"carrierCode,omitempty"`
	CarrierTitle *string `json:"carrierTitle,omitempty"`
	Url *string `json:"url,omitempty"`
	Number *string `json:"number,omitempty"`
	LabelUrl *string `json:"labelUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ShipmentTracking ShipmentTracking

// NewShipmentTracking instantiates a new ShipmentTracking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentTracking() *ShipmentTracking {
	this := ShipmentTracking{}
	return &this
}

// NewShipmentTrackingWithDefaults instantiates a new ShipmentTracking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentTrackingWithDefaults() *ShipmentTracking {
	this := ShipmentTracking{}
	return &this
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *ShipmentTracking) GetCarrierCode() string {
	if o == nil || IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentTracking) GetCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// &#39;Has&#39;CarrierCode returns a boolean if a field has been set.
func (o *ShipmentTracking) &#39;Has&#39;CarrierCode() bool {
	if o != nil && !IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *ShipmentTracking) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

// GetCarrierTitle returns the CarrierTitle field value if set, zero value otherwise.
func (o *ShipmentTracking) GetCarrierTitle() string {
	if o == nil || IsNil(o.CarrierTitle) {
		var ret string
		return ret
	}
	return *o.CarrierTitle
}

// GetCarrierTitleOk returns a tuple with the CarrierTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentTracking) GetCarrierTitleOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierTitle) {
		return nil, false
	}
	return o.CarrierTitle, true
}

// &#39;Has&#39;CarrierTitle returns a boolean if a field has been set.
func (o *ShipmentTracking) &#39;Has&#39;CarrierTitle() bool {
	if o != nil && !IsNil(o.CarrierTitle) {
		return true
	}

	return false
}

// SetCarrierTitle gets a reference to the given string and assigns it to the CarrierTitle field.
func (o *ShipmentTracking) SetCarrierTitle(v string) {
	o.CarrierTitle = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ShipmentTracking) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentTracking) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// &#39;Has&#39;Url returns a boolean if a field has been set.
func (o *ShipmentTracking) &#39;Has&#39;Url() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ShipmentTracking) SetUrl(v string) {
	o.Url = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ShipmentTracking) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentTracking) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// &#39;Has&#39;Number returns a boolean if a field has been set.
func (o *ShipmentTracking) &#39;Has&#39;Number() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *ShipmentTracking) SetNumber(v string) {
	o.Number = &v
}

// GetLabelUrl returns the LabelUrl field value if set, zero value otherwise.
func (o *ShipmentTracking) GetLabelUrl() string {
	if o == nil || IsNil(o.LabelUrl) {
		var ret string
		return ret
	}
	return *o.LabelUrl
}

// GetLabelUrlOk returns a tuple with the LabelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentTracking) GetLabelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LabelUrl) {
		return nil, false
	}
	return o.LabelUrl, true
}

// &#39;Has&#39;LabelUrl returns a boolean if a field has been set.
func (o *ShipmentTracking) &#39;Has&#39;LabelUrl() bool {
	if o != nil && !IsNil(o.LabelUrl) {
		return true
	}

	return false
}

// SetLabelUrl gets a reference to the given string and assigns it to the LabelUrl field.
func (o *ShipmentTracking) SetLabelUrl(v string) {
	o.LabelUrl = &v
}

func (o ShipmentTracking) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentTracking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierCode) {
		toSerialize["carrierCode"] = o.CarrierCode
	}
	if !IsNil(o.CarrierTitle) {
		toSerialize["carrierTitle"] = o.CarrierTitle
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.LabelUrl) {
		toSerialize["labelUrl"] = o.LabelUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ShipmentTracking) UnmarshalJSON(data []byte) (err error) {
	varShipmentTracking := _ShipmentTracking{}

	err = json.Unmarshal(data, &varShipmentTracking)

	if err != nil {
		return err
	}

	*o = ShipmentTracking(varShipmentTracking)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "carrierCode")
		delete(additionalProperties, "carrierTitle")
		delete(additionalProperties, "url")
		delete(additionalProperties, "number")
		delete(additionalProperties, "labelUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ShipmentTracking) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ShipmentTracking) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableShipmentTracking struct {
	value *ShipmentTracking
	isSet bool
}

func (v NullableShipmentTracking) Get() *ShipmentTracking {
	return v.value
}

func (v *NullableShipmentTracking) Set(val *ShipmentTracking) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentTracking) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentTracking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentTracking(val *ShipmentTracking) *NullableShipmentTracking {
	return &NullableShipmentTracking{value: val, isSet: true}
}

func (v NullableShipmentTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentTracking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



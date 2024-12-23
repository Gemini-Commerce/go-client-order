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

// checks if the OrderFulfillment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderFulfillment{}

// OrderFulfillment struct for OrderFulfillment
type OrderFulfillment struct {
	CreatedAt            *time.Time             `json:"createdAt,omitempty"`
	UpdatedAt            *time.Time             `json:"updatedAt,omitempty"`
	OrderId              *string                `json:"orderId,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Status               *string                `json:"status,omitempty"`
	Items                []OrderFulfillmentItem `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderFulfillment OrderFulfillment

// NewOrderFulfillment instantiates a new OrderFulfillment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderFulfillment() *OrderFulfillment {
	this := OrderFulfillment{}
	return &this
}

// NewOrderFulfillmentWithDefaults instantiates a new OrderFulfillment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderFulfillmentWithDefaults() *OrderFulfillment {
	this := OrderFulfillment{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrderFulfillment) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFulfillment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrderFulfillment) IsSetCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrderFulfillment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrderFulfillment) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFulfillment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrderFulfillment) IsSetUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrderFulfillment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderFulfillment) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFulfillment) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderFulfillment) IsSetOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderFulfillment) SetOrderId(v string) {
	o.OrderId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderFulfillment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFulfillment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderFulfillment) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderFulfillment) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderFulfillment) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFulfillment) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderFulfillment) IsSetStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrderFulfillment) SetStatus(v string) {
	o.Status = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *OrderFulfillment) GetItems() []OrderFulfillmentItem {
	if o == nil || IsNil(o.Items) {
		var ret []OrderFulfillmentItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFulfillment) GetItemsOk() ([]OrderFulfillmentItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *OrderFulfillment) IsSetItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrderFulfillmentItem and assigns it to the Items field.
func (o *OrderFulfillment) SetItems(v []OrderFulfillmentItem) {
	o.Items = v
}

func (o OrderFulfillment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderFulfillment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderFulfillment) UnmarshalJSON(data []byte) (err error) {
	varOrderFulfillment := _OrderFulfillment{}

	err = json.Unmarshal(data, &varOrderFulfillment)

	if err != nil {
		return err
	}

	*o = OrderFulfillment(varOrderFulfillment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderFulfillment) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderFulfillment) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderFulfillment struct {
	value *OrderFulfillment
	isSet bool
}

func (v NullableOrderFulfillment) Get() *OrderFulfillment {
	return v.value
}

func (v *NullableOrderFulfillment) Set(val *OrderFulfillment) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderFulfillment) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderFulfillment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderFulfillment(val *OrderFulfillment) *NullableOrderFulfillment {
	return &NullableOrderFulfillment{value: val, isSet: true}
}

func (v NullableOrderFulfillment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderFulfillment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

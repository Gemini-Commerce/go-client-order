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

// checks if the OrderShipment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderShipment{}

// OrderShipment struct for OrderShipment
type OrderShipment struct {
	CreatedAt            *time.Time          `json:"createdAt,omitempty"`
	UpdatedAt            *time.Time          `json:"updatedAt,omitempty"`
	OrderId              *string             `json:"orderId,omitempty"`
	Id                   *string             `json:"id,omitempty"`
	Status               *string             `json:"status,omitempty"`
	Items                []OrderShipmentItem `json:"items,omitempty"`
	Address              *OrderPostalAddress `json:"address,omitempty"`
	FromAddress          *OrderPostalAddress `json:"fromAddress,omitempty"`
	ReturnAddress        *OrderPostalAddress `json:"returnAddress,omitempty"`
	Tracking             []ShipmentTracking  `json:"tracking,omitempty"`
	ReturnTracking       []ShipmentTracking  `json:"returnTracking,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderShipment OrderShipment

// NewOrderShipment instantiates a new OrderShipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderShipment() *OrderShipment {
	this := OrderShipment{}
	return &this
}

// NewOrderShipmentWithDefaults instantiates a new OrderShipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderShipmentWithDefaults() *OrderShipment {
	this := OrderShipment{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrderShipment) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrderShipment) IsSetCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrderShipment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrderShipment) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrderShipment) IsSetUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrderShipment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderShipment) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipment) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderShipment) IsSetOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderShipment) SetOrderId(v string) {
	o.OrderId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderShipment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderShipment) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderShipment) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderShipment) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipment) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderShipment) IsSetStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrderShipment) SetStatus(v string) {
	o.Status = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *OrderShipment) GetItems() []OrderShipmentItem {
	if o == nil || IsNil(o.Items) {
		var ret []OrderShipmentItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipment) GetItemsOk() ([]OrderShipmentItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *OrderShipment) IsSetItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrderShipmentItem and assigns it to the Items field.
func (o *OrderShipment) SetItems(v []OrderShipmentItem) {
	o.Items = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrderShipment) GetAddress() OrderPostalAddress {
	if o == nil || IsNil(o.Address) {
		var ret OrderPostalAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipment) GetAddressOk() (*OrderPostalAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrderShipment) IsSetAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given OrderPostalAddress and assigns it to the Address field.
func (o *OrderShipment) SetAddress(v OrderPostalAddress) {
	o.Address = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *OrderShipment) GetFromAddress() OrderPostalAddress {
	if o == nil || IsNil(o.FromAddress) {
		var ret OrderPostalAddress
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipment) GetFromAddressOk() (*OrderPostalAddress, bool) {
	if o == nil || IsNil(o.FromAddress) {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *OrderShipment) IsSetFromAddress() bool {
	if o != nil && !IsNil(o.FromAddress) {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given OrderPostalAddress and assigns it to the FromAddress field.
func (o *OrderShipment) SetFromAddress(v OrderPostalAddress) {
	o.FromAddress = &v
}

// GetReturnAddress returns the ReturnAddress field value if set, zero value otherwise.
func (o *OrderShipment) GetReturnAddress() OrderPostalAddress {
	if o == nil || IsNil(o.ReturnAddress) {
		var ret OrderPostalAddress
		return ret
	}
	return *o.ReturnAddress
}

// GetReturnAddressOk returns a tuple with the ReturnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipment) GetReturnAddressOk() (*OrderPostalAddress, bool) {
	if o == nil || IsNil(o.ReturnAddress) {
		return nil, false
	}
	return o.ReturnAddress, true
}

// HasReturnAddress returns a boolean if a field has been set.
func (o *OrderShipment) IsSetReturnAddress() bool {
	if o != nil && !IsNil(o.ReturnAddress) {
		return true
	}

	return false
}

// SetReturnAddress gets a reference to the given OrderPostalAddress and assigns it to the ReturnAddress field.
func (o *OrderShipment) SetReturnAddress(v OrderPostalAddress) {
	o.ReturnAddress = &v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *OrderShipment) GetTracking() []ShipmentTracking {
	if o == nil || IsNil(o.Tracking) {
		var ret []ShipmentTracking
		return ret
	}
	return o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipment) GetTrackingOk() ([]ShipmentTracking, bool) {
	if o == nil || IsNil(o.Tracking) {
		return nil, false
	}
	return o.Tracking, true
}

// HasTracking returns a boolean if a field has been set.
func (o *OrderShipment) IsSetTracking() bool {
	if o != nil && !IsNil(o.Tracking) {
		return true
	}

	return false
}

// SetTracking gets a reference to the given []ShipmentTracking and assigns it to the Tracking field.
func (o *OrderShipment) SetTracking(v []ShipmentTracking) {
	o.Tracking = v
}

// GetReturnTracking returns the ReturnTracking field value if set, zero value otherwise.
func (o *OrderShipment) GetReturnTracking() []ShipmentTracking {
	if o == nil || IsNil(o.ReturnTracking) {
		var ret []ShipmentTracking
		return ret
	}
	return o.ReturnTracking
}

// GetReturnTrackingOk returns a tuple with the ReturnTracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipment) GetReturnTrackingOk() ([]ShipmentTracking, bool) {
	if o == nil || IsNil(o.ReturnTracking) {
		return nil, false
	}
	return o.ReturnTracking, true
}

// HasReturnTracking returns a boolean if a field has been set.
func (o *OrderShipment) IsSetReturnTracking() bool {
	if o != nil && !IsNil(o.ReturnTracking) {
		return true
	}

	return false
}

// SetReturnTracking gets a reference to the given []ShipmentTracking and assigns it to the ReturnTracking field.
func (o *OrderShipment) SetReturnTracking(v []ShipmentTracking) {
	o.ReturnTracking = v
}

func (o OrderShipment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderShipment) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.FromAddress) {
		toSerialize["fromAddress"] = o.FromAddress
	}
	if !IsNil(o.ReturnAddress) {
		toSerialize["returnAddress"] = o.ReturnAddress
	}
	if !IsNil(o.Tracking) {
		toSerialize["tracking"] = o.Tracking
	}
	if !IsNil(o.ReturnTracking) {
		toSerialize["returnTracking"] = o.ReturnTracking
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderShipment) UnmarshalJSON(data []byte) (err error) {
	varOrderShipment := _OrderShipment{}

	err = json.Unmarshal(data, &varOrderShipment)

	if err != nil {
		return err
	}

	*o = OrderShipment(varOrderShipment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "items")
		delete(additionalProperties, "address")
		delete(additionalProperties, "fromAddress")
		delete(additionalProperties, "returnAddress")
		delete(additionalProperties, "tracking")
		delete(additionalProperties, "returnTracking")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderShipment) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderShipment) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderShipment struct {
	value *OrderShipment
	isSet bool
}

func (v NullableOrderShipment) Get() *OrderShipment {
	return v.value
}

func (v *NullableOrderShipment) Set(val *OrderShipment) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderShipment(val *OrderShipment) *NullableOrderShipment {
	return &NullableOrderShipment{value: val, isSet: true}
}

func (v NullableOrderShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

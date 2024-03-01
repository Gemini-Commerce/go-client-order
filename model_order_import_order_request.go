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
	"bytes"
	"fmt"
)

// checks if the OrderImportOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderImportOrderRequest{}

// OrderImportOrderRequest struct for OrderImportOrderRequest
type OrderImportOrderRequest struct {
	TenantId string `json:"tenantId"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Number string `json:"number"`
	Channel *string `json:"channel,omitempty"`
	Market string `json:"market"`
	Locale string `json:"locale"`
	CustomerInfo OrderDataCustomerInfo `json:"customerInfo"`
	ShippingAddress OrderPostalAddress `json:"shippingAddress"`
	BillingAddress OrderPostalAddress `json:"billingAddress"`
	Payments []ImportOrderRequestImportedPayment `json:"payments"`
	PaymentsInfo []OrderDataPaymentInfo `json:"paymentsInfo"`
	ShipmentsInfo []OrderDataShipmentInfo `json:"shipmentsInfo"`
	Items []OrderOrderDataItem `json:"items"`
	Subtotals map[string]OrderDataSubtotal `json:"subtotals"`
	Totals map[string]OrderDataTotal `json:"totals"`
	Status string `json:"status"`
	Currency OrderCurrency `json:"currency"`
}

type _OrderImportOrderRequest OrderImportOrderRequest

// NewOrderImportOrderRequest instantiates a new OrderImportOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderImportOrderRequest(tenantId string, number string, market string, locale string, customerInfo OrderDataCustomerInfo, shippingAddress OrderPostalAddress, billingAddress OrderPostalAddress, payments []ImportOrderRequestImportedPayment, paymentsInfo []OrderDataPaymentInfo, shipmentsInfo []OrderDataShipmentInfo, items []OrderOrderDataItem, subtotals map[string]OrderDataSubtotal, totals map[string]OrderDataTotal, status string, currency OrderCurrency) *OrderImportOrderRequest {
	this := OrderImportOrderRequest{}
	this.TenantId = tenantId
	this.Number = number
	this.Market = market
	this.Locale = locale
	this.CustomerInfo = customerInfo
	this.ShippingAddress = shippingAddress
	this.BillingAddress = billingAddress
	this.Payments = payments
	this.PaymentsInfo = paymentsInfo
	this.ShipmentsInfo = shipmentsInfo
	this.Items = items
	this.Subtotals = subtotals
	this.Totals = totals
	this.Status = status
	this.Currency = currency
	return &this
}

// NewOrderImportOrderRequestWithDefaults instantiates a new OrderImportOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderImportOrderRequestWithDefaults() *OrderImportOrderRequest {
	this := OrderImportOrderRequest{}
	var currency OrderCurrency = ORDERCURRENCY_XXX
	this.Currency = currency
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderImportOrderRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderImportOrderRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrderImportOrderRequest) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// IsSetCreatedAt returns a boolean if a field has been set.
func (o *OrderImportOrderRequest) IsSetCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrderImportOrderRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetNumber returns the Number field value
func (o *OrderImportOrderRequest) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *OrderImportOrderRequest) SetNumber(v string) {
	o.Number = v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *OrderImportOrderRequest) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// IsSetChannel returns a boolean if a field has been set.
func (o *OrderImportOrderRequest) IsSetChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *OrderImportOrderRequest) SetChannel(v string) {
	o.Channel = &v
}

// GetMarket returns the Market field value
func (o *OrderImportOrderRequest) GetMarket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetMarketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *OrderImportOrderRequest) SetMarket(v string) {
	o.Market = v
}

// GetLocale returns the Locale field value
func (o *OrderImportOrderRequest) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *OrderImportOrderRequest) SetLocale(v string) {
	o.Locale = v
}

// GetCustomerInfo returns the CustomerInfo field value
func (o *OrderImportOrderRequest) GetCustomerInfo() OrderDataCustomerInfo {
	if o == nil {
		var ret OrderDataCustomerInfo
		return ret
	}

	return o.CustomerInfo
}

// GetCustomerInfoOk returns a tuple with the CustomerInfo field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetCustomerInfoOk() (*OrderDataCustomerInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerInfo, true
}

// SetCustomerInfo sets field value
func (o *OrderImportOrderRequest) SetCustomerInfo(v OrderDataCustomerInfo) {
	o.CustomerInfo = v
}

// GetShippingAddress returns the ShippingAddress field value
func (o *OrderImportOrderRequest) GetShippingAddress() OrderPostalAddress {
	if o == nil {
		var ret OrderPostalAddress
		return ret
	}

	return o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetShippingAddressOk() (*OrderPostalAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingAddress, true
}

// SetShippingAddress sets field value
func (o *OrderImportOrderRequest) SetShippingAddress(v OrderPostalAddress) {
	o.ShippingAddress = v
}

// GetBillingAddress returns the BillingAddress field value
func (o *OrderImportOrderRequest) GetBillingAddress() OrderPostalAddress {
	if o == nil {
		var ret OrderPostalAddress
		return ret
	}

	return o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetBillingAddressOk() (*OrderPostalAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingAddress, true
}

// SetBillingAddress sets field value
func (o *OrderImportOrderRequest) SetBillingAddress(v OrderPostalAddress) {
	o.BillingAddress = v
}

// GetPayments returns the Payments field value
func (o *OrderImportOrderRequest) GetPayments() []ImportOrderRequestImportedPayment {
	if o == nil {
		var ret []ImportOrderRequestImportedPayment
		return ret
	}

	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetPaymentsOk() ([]ImportOrderRequestImportedPayment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payments, true
}

// SetPayments sets field value
func (o *OrderImportOrderRequest) SetPayments(v []ImportOrderRequestImportedPayment) {
	o.Payments = v
}

// GetPaymentsInfo returns the PaymentsInfo field value
func (o *OrderImportOrderRequest) GetPaymentsInfo() []OrderDataPaymentInfo {
	if o == nil {
		var ret []OrderDataPaymentInfo
		return ret
	}

	return o.PaymentsInfo
}

// GetPaymentsInfoOk returns a tuple with the PaymentsInfo field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetPaymentsInfoOk() ([]OrderDataPaymentInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentsInfo, true
}

// SetPaymentsInfo sets field value
func (o *OrderImportOrderRequest) SetPaymentsInfo(v []OrderDataPaymentInfo) {
	o.PaymentsInfo = v
}

// GetShipmentsInfo returns the ShipmentsInfo field value
func (o *OrderImportOrderRequest) GetShipmentsInfo() []OrderDataShipmentInfo {
	if o == nil {
		var ret []OrderDataShipmentInfo
		return ret
	}

	return o.ShipmentsInfo
}

// GetShipmentsInfoOk returns a tuple with the ShipmentsInfo field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetShipmentsInfoOk() ([]OrderDataShipmentInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShipmentsInfo, true
}

// SetShipmentsInfo sets field value
func (o *OrderImportOrderRequest) SetShipmentsInfo(v []OrderDataShipmentInfo) {
	o.ShipmentsInfo = v
}

// GetItems returns the Items field value
func (o *OrderImportOrderRequest) GetItems() []OrderOrderDataItem {
	if o == nil {
		var ret []OrderOrderDataItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetItemsOk() ([]OrderOrderDataItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *OrderImportOrderRequest) SetItems(v []OrderOrderDataItem) {
	o.Items = v
}

// GetSubtotals returns the Subtotals field value
func (o *OrderImportOrderRequest) GetSubtotals() map[string]OrderDataSubtotal {
	if o == nil {
		var ret map[string]OrderDataSubtotal
		return ret
	}

	return o.Subtotals
}

// GetSubtotalsOk returns a tuple with the Subtotals field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetSubtotalsOk() (*map[string]OrderDataSubtotal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtotals, true
}

// SetSubtotals sets field value
func (o *OrderImportOrderRequest) SetSubtotals(v map[string]OrderDataSubtotal) {
	o.Subtotals = v
}

// GetTotals returns the Totals field value
func (o *OrderImportOrderRequest) GetTotals() map[string]OrderDataTotal {
	if o == nil {
		var ret map[string]OrderDataTotal
		return ret
	}

	return o.Totals
}

// GetTotalsOk returns a tuple with the Totals field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetTotalsOk() (*map[string]OrderDataTotal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Totals, true
}

// SetTotals sets field value
func (o *OrderImportOrderRequest) SetTotals(v map[string]OrderDataTotal) {
	o.Totals = v
}

// GetStatus returns the Status field value
func (o *OrderImportOrderRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrderImportOrderRequest) SetStatus(v string) {
	o.Status = v
}

// GetCurrency returns the Currency field value
func (o *OrderImportOrderRequest) GetCurrency() OrderCurrency {
	if o == nil {
		var ret OrderCurrency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *OrderImportOrderRequest) GetCurrencyOk() (*OrderCurrency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *OrderImportOrderRequest) SetCurrency(v OrderCurrency) {
	o.Currency = v
}

func (o OrderImportOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderImportOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	toSerialize["number"] = o.Number
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	toSerialize["market"] = o.Market
	toSerialize["locale"] = o.Locale
	toSerialize["customerInfo"] = o.CustomerInfo
	toSerialize["shippingAddress"] = o.ShippingAddress
	toSerialize["billingAddress"] = o.BillingAddress
	toSerialize["payments"] = o.Payments
	toSerialize["paymentsInfo"] = o.PaymentsInfo
	toSerialize["shipmentsInfo"] = o.ShipmentsInfo
	toSerialize["items"] = o.Items
	toSerialize["subtotals"] = o.Subtotals
	toSerialize["totals"] = o.Totals
	toSerialize["status"] = o.Status
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

func (o *OrderImportOrderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"number",
		"market",
		"locale",
		"customerInfo",
		"shippingAddress",
		"billingAddress",
		"payments",
		"paymentsInfo",
		"shipmentsInfo",
		"items",
		"subtotals",
		"totals",
		"status",
		"currency",
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

	varOrderImportOrderRequest := _OrderImportOrderRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderImportOrderRequest)

	if err != nil {
		return err
	}

	*o = OrderImportOrderRequest(varOrderImportOrderRequest)

	return err
}

type NullableOrderImportOrderRequest struct {
	value *OrderImportOrderRequest
	isSet bool
}

func (v NullableOrderImportOrderRequest) Get() *OrderImportOrderRequest {
	return v.value
}

func (v *NullableOrderImportOrderRequest) Set(val *OrderImportOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderImportOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderImportOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderImportOrderRequest(val *OrderImportOrderRequest) *NullableOrderImportOrderRequest {
	return &NullableOrderImportOrderRequest{value: val, isSet: true}
}

func (v NullableOrderImportOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderImportOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



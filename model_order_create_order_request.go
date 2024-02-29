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

// checks if the OrderCreateOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateOrderRequest{}

// OrderCreateOrderRequest struct for OrderCreateOrderRequest
type OrderCreateOrderRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Number *string `json:"number,omitempty"`
	Channel *string `json:"channel,omitempty"`
	Market *string `json:"market,omitempty"`
	Locale *string `json:"locale,omitempty"`
	Items []OrderOrderDataItem `json:"items,omitempty"`
	PaymentsInfo []OrderDataPaymentInfo `json:"paymentsInfo,omitempty"`
	ShipmentsInfo []OrderDataShipmentInfo `json:"shipmentsInfo,omitempty"`
	Promotions []OrderDataPromotionInfo `json:"promotions,omitempty"`
	Payments []CreateOrderRequestInitialPayment `json:"payments,omitempty"`
	Currency *OrderCurrency `json:"currency,omitempty"`
	Subtotals *map[string]OrderDataSubtotal `json:"subtotals,omitempty"`
	Totals *map[string]OrderDataTotal `json:"totals,omitempty"`
	VatIncluded *bool `json:"vatIncluded,omitempty"`
	BillingAddress *OrderPostalAddress `json:"billingAddress,omitempty"`
	ShippingAddress *OrderPostalAddress `json:"shippingAddress,omitempty"`
	CustomerInfo *OrderDataCustomerInfo `json:"customerInfo,omitempty"`
	CartGrn *string `json:"cartGrn,omitempty"`
	OnHold *bool `json:"onHold,omitempty"`
	Notes *string `json:"notes,omitempty"`
}

// NewOrderCreateOrderRequest instantiates a new OrderCreateOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateOrderRequest() *OrderCreateOrderRequest {
	this := OrderCreateOrderRequest{}
	var currency OrderCurrency = XXX
	this.Currency = &currency
	return &this
}

// NewOrderCreateOrderRequestWithDefaults instantiates a new OrderCreateOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateOrderRequestWithDefaults() *OrderCreateOrderRequest {
	this := OrderCreateOrderRequest{}
	var currency OrderCurrency = XXX
	this.Currency = &currency
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OrderCreateOrderRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *OrderCreateOrderRequest) SetNumber(v string) {
	o.Number = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *OrderCreateOrderRequest) SetChannel(v string) {
	o.Channel = &v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetMarket() string {
	if o == nil || IsNil(o.Market) {
		var ret string
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetMarketOk() (*string, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given string and assigns it to the Market field.
func (o *OrderCreateOrderRequest) SetMarket(v string) {
	o.Market = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *OrderCreateOrderRequest) SetLocale(v string) {
	o.Locale = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetItems() []OrderOrderDataItem {
	if o == nil || IsNil(o.Items) {
		var ret []OrderOrderDataItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetItemsOk() ([]OrderOrderDataItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrderOrderDataItem and assigns it to the Items field.
func (o *OrderCreateOrderRequest) SetItems(v []OrderOrderDataItem) {
	o.Items = v
}

// GetPaymentsInfo returns the PaymentsInfo field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetPaymentsInfo() []OrderDataPaymentInfo {
	if o == nil || IsNil(o.PaymentsInfo) {
		var ret []OrderDataPaymentInfo
		return ret
	}
	return o.PaymentsInfo
}

// GetPaymentsInfoOk returns a tuple with the PaymentsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetPaymentsInfoOk() ([]OrderDataPaymentInfo, bool) {
	if o == nil || IsNil(o.PaymentsInfo) {
		return nil, false
	}
	return o.PaymentsInfo, true
}

// HasPaymentsInfo returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasPaymentsInfo() bool {
	if o != nil && !IsNil(o.PaymentsInfo) {
		return true
	}

	return false
}

// SetPaymentsInfo gets a reference to the given []OrderDataPaymentInfo and assigns it to the PaymentsInfo field.
func (o *OrderCreateOrderRequest) SetPaymentsInfo(v []OrderDataPaymentInfo) {
	o.PaymentsInfo = v
}

// GetShipmentsInfo returns the ShipmentsInfo field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetShipmentsInfo() []OrderDataShipmentInfo {
	if o == nil || IsNil(o.ShipmentsInfo) {
		var ret []OrderDataShipmentInfo
		return ret
	}
	return o.ShipmentsInfo
}

// GetShipmentsInfoOk returns a tuple with the ShipmentsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetShipmentsInfoOk() ([]OrderDataShipmentInfo, bool) {
	if o == nil || IsNil(o.ShipmentsInfo) {
		return nil, false
	}
	return o.ShipmentsInfo, true
}

// HasShipmentsInfo returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasShipmentsInfo() bool {
	if o != nil && !IsNil(o.ShipmentsInfo) {
		return true
	}

	return false
}

// SetShipmentsInfo gets a reference to the given []OrderDataShipmentInfo and assigns it to the ShipmentsInfo field.
func (o *OrderCreateOrderRequest) SetShipmentsInfo(v []OrderDataShipmentInfo) {
	o.ShipmentsInfo = v
}

// GetPromotions returns the Promotions field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetPromotions() []OrderDataPromotionInfo {
	if o == nil || IsNil(o.Promotions) {
		var ret []OrderDataPromotionInfo
		return ret
	}
	return o.Promotions
}

// GetPromotionsOk returns a tuple with the Promotions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetPromotionsOk() ([]OrderDataPromotionInfo, bool) {
	if o == nil || IsNil(o.Promotions) {
		return nil, false
	}
	return o.Promotions, true
}

// HasPromotions returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasPromotions() bool {
	if o != nil && !IsNil(o.Promotions) {
		return true
	}

	return false
}

// SetPromotions gets a reference to the given []OrderDataPromotionInfo and assigns it to the Promotions field.
func (o *OrderCreateOrderRequest) SetPromotions(v []OrderDataPromotionInfo) {
	o.Promotions = v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetPayments() []CreateOrderRequestInitialPayment {
	if o == nil || IsNil(o.Payments) {
		var ret []CreateOrderRequestInitialPayment
		return ret
	}
	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetPaymentsOk() ([]CreateOrderRequestInitialPayment, bool) {
	if o == nil || IsNil(o.Payments) {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasPayments() bool {
	if o != nil && !IsNil(o.Payments) {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []CreateOrderRequestInitialPayment and assigns it to the Payments field.
func (o *OrderCreateOrderRequest) SetPayments(v []CreateOrderRequestInitialPayment) {
	o.Payments = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetCurrency() OrderCurrency {
	if o == nil || IsNil(o.Currency) {
		var ret OrderCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetCurrencyOk() (*OrderCurrency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given OrderCurrency and assigns it to the Currency field.
func (o *OrderCreateOrderRequest) SetCurrency(v OrderCurrency) {
	o.Currency = &v
}

// GetSubtotals returns the Subtotals field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetSubtotals() map[string]OrderDataSubtotal {
	if o == nil || IsNil(o.Subtotals) {
		var ret map[string]OrderDataSubtotal
		return ret
	}
	return *o.Subtotals
}

// GetSubtotalsOk returns a tuple with the Subtotals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetSubtotalsOk() (*map[string]OrderDataSubtotal, bool) {
	if o == nil || IsNil(o.Subtotals) {
		return nil, false
	}
	return o.Subtotals, true
}

// HasSubtotals returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasSubtotals() bool {
	if o != nil && !IsNil(o.Subtotals) {
		return true
	}

	return false
}

// SetSubtotals gets a reference to the given map[string]OrderDataSubtotal and assigns it to the Subtotals field.
func (o *OrderCreateOrderRequest) SetSubtotals(v map[string]OrderDataSubtotal) {
	o.Subtotals = &v
}

// GetTotals returns the Totals field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetTotals() map[string]OrderDataTotal {
	if o == nil || IsNil(o.Totals) {
		var ret map[string]OrderDataTotal
		return ret
	}
	return *o.Totals
}

// GetTotalsOk returns a tuple with the Totals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetTotalsOk() (*map[string]OrderDataTotal, bool) {
	if o == nil || IsNil(o.Totals) {
		return nil, false
	}
	return o.Totals, true
}

// HasTotals returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasTotals() bool {
	if o != nil && !IsNil(o.Totals) {
		return true
	}

	return false
}

// SetTotals gets a reference to the given map[string]OrderDataTotal and assigns it to the Totals field.
func (o *OrderCreateOrderRequest) SetTotals(v map[string]OrderDataTotal) {
	o.Totals = &v
}

// GetVatIncluded returns the VatIncluded field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetVatIncluded() bool {
	if o == nil || IsNil(o.VatIncluded) {
		var ret bool
		return ret
	}
	return *o.VatIncluded
}

// GetVatIncludedOk returns a tuple with the VatIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetVatIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.VatIncluded) {
		return nil, false
	}
	return o.VatIncluded, true
}

// HasVatIncluded returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasVatIncluded() bool {
	if o != nil && !IsNil(o.VatIncluded) {
		return true
	}

	return false
}

// SetVatIncluded gets a reference to the given bool and assigns it to the VatIncluded field.
func (o *OrderCreateOrderRequest) SetVatIncluded(v bool) {
	o.VatIncluded = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetBillingAddress() OrderPostalAddress {
	if o == nil || IsNil(o.BillingAddress) {
		var ret OrderPostalAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetBillingAddressOk() (*OrderPostalAddress, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given OrderPostalAddress and assigns it to the BillingAddress field.
func (o *OrderCreateOrderRequest) SetBillingAddress(v OrderPostalAddress) {
	o.BillingAddress = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetShippingAddress() OrderPostalAddress {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret OrderPostalAddress
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetShippingAddressOk() (*OrderPostalAddress, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given OrderPostalAddress and assigns it to the ShippingAddress field.
func (o *OrderCreateOrderRequest) SetShippingAddress(v OrderPostalAddress) {
	o.ShippingAddress = &v
}

// GetCustomerInfo returns the CustomerInfo field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetCustomerInfo() OrderDataCustomerInfo {
	if o == nil || IsNil(o.CustomerInfo) {
		var ret OrderDataCustomerInfo
		return ret
	}
	return *o.CustomerInfo
}

// GetCustomerInfoOk returns a tuple with the CustomerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetCustomerInfoOk() (*OrderDataCustomerInfo, bool) {
	if o == nil || IsNil(o.CustomerInfo) {
		return nil, false
	}
	return o.CustomerInfo, true
}

// HasCustomerInfo returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasCustomerInfo() bool {
	if o != nil && !IsNil(o.CustomerInfo) {
		return true
	}

	return false
}

// SetCustomerInfo gets a reference to the given OrderDataCustomerInfo and assigns it to the CustomerInfo field.
func (o *OrderCreateOrderRequest) SetCustomerInfo(v OrderDataCustomerInfo) {
	o.CustomerInfo = &v
}

// GetCartGrn returns the CartGrn field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetCartGrn() string {
	if o == nil || IsNil(o.CartGrn) {
		var ret string
		return ret
	}
	return *o.CartGrn
}

// GetCartGrnOk returns a tuple with the CartGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetCartGrnOk() (*string, bool) {
	if o == nil || IsNil(o.CartGrn) {
		return nil, false
	}
	return o.CartGrn, true
}

// HasCartGrn returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasCartGrn() bool {
	if o != nil && !IsNil(o.CartGrn) {
		return true
	}

	return false
}

// SetCartGrn gets a reference to the given string and assigns it to the CartGrn field.
func (o *OrderCreateOrderRequest) SetCartGrn(v string) {
	o.CartGrn = &v
}

// GetOnHold returns the OnHold field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetOnHold() bool {
	if o == nil || IsNil(o.OnHold) {
		var ret bool
		return ret
	}
	return *o.OnHold
}

// GetOnHoldOk returns a tuple with the OnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetOnHoldOk() (*bool, bool) {
	if o == nil || IsNil(o.OnHold) {
		return nil, false
	}
	return o.OnHold, true
}

// HasOnHold returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasOnHold() bool {
	if o != nil && !IsNil(o.OnHold) {
		return true
	}

	return false
}

// SetOnHold gets a reference to the given bool and assigns it to the OnHold field.
func (o *OrderCreateOrderRequest) SetOnHold(v bool) {
	o.OnHold = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderCreateOrderRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateOrderRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderCreateOrderRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OrderCreateOrderRequest) SetNotes(v string) {
	o.Notes = &v
}

func (o OrderCreateOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.PaymentsInfo) {
		toSerialize["paymentsInfo"] = o.PaymentsInfo
	}
	if !IsNil(o.ShipmentsInfo) {
		toSerialize["shipmentsInfo"] = o.ShipmentsInfo
	}
	if !IsNil(o.Promotions) {
		toSerialize["promotions"] = o.Promotions
	}
	if !IsNil(o.Payments) {
		toSerialize["payments"] = o.Payments
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Subtotals) {
		toSerialize["subtotals"] = o.Subtotals
	}
	if !IsNil(o.Totals) {
		toSerialize["totals"] = o.Totals
	}
	if !IsNil(o.VatIncluded) {
		toSerialize["vatIncluded"] = o.VatIncluded
	}
	if !IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !IsNil(o.ShippingAddress) {
		toSerialize["shippingAddress"] = o.ShippingAddress
	}
	if !IsNil(o.CustomerInfo) {
		toSerialize["customerInfo"] = o.CustomerInfo
	}
	if !IsNil(o.CartGrn) {
		toSerialize["cartGrn"] = o.CartGrn
	}
	if !IsNil(o.OnHold) {
		toSerialize["onHold"] = o.OnHold
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableOrderCreateOrderRequest struct {
	value *OrderCreateOrderRequest
	isSet bool
}

func (v NullableOrderCreateOrderRequest) Get() *OrderCreateOrderRequest {
	return v.value
}

func (v *NullableOrderCreateOrderRequest) Set(val *OrderCreateOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateOrderRequest(val *OrderCreateOrderRequest) *NullableOrderCreateOrderRequest {
	return &NullableOrderCreateOrderRequest{value: val, isSet: true}
}

func (v NullableOrderCreateOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



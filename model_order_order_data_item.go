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

// checks if the OrderOrderDataItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderOrderDataItem{}

// OrderOrderDataItem struct for OrderOrderDataItem
type OrderOrderDataItem struct {
	Id *string `json:"id,omitempty"`
	ProductGrn *string `json:"productGrn,omitempty"`
	QtyOrdered *int64 `json:"qtyOrdered,omitempty"`
	QtyCommitted *int64 `json:"qtyCommitted,omitempty"`
	UnitSalePrice *OrderMoney `json:"unitSalePrice,omitempty"`
	UnitListPrice *OrderMoney `json:"unitListPrice,omitempty"`
	UnitBasePrice *OrderMoney `json:"unitBasePrice,omitempty"`
	UnitVatAmount *OrderMoney `json:"unitVatAmount,omitempty"`
	RowSalePrice *OrderMoney `json:"rowSalePrice,omitempty"`
	RowListPrice *OrderMoney `json:"rowListPrice,omitempty"`
	RowVatAmount *OrderMoney `json:"rowVatAmount,omitempty"`
	DiscountAmount *OrderMoney `json:"discountAmount,omitempty"`
	RowBasePrice *OrderMoney `json:"rowBasePrice,omitempty"`
	VatPercentage *float32 `json:"vatPercentage,omitempty"`
	VatInaccurate *bool `json:"vatInaccurate,omitempty"`
	VatCalculated *bool `json:"vatCalculated,omitempty"`
	ProductName *string `json:"productName,omitempty"`
	ProductCode *string `json:"productCode,omitempty"`
	ProductSku *string `json:"productSku,omitempty"`
	ProductOptions *string `json:"productOptions,omitempty"`
	ProductImg *string `json:"productImg,omitempty"`
	ProductData *string `json:"productData,omitempty"`
	ShipmentInfoReference *string `json:"shipmentInfoReference,omitempty"`
	PromotionGrn []string `json:"promotionGrn,omitempty"`
	ProductIsVirtual *bool `json:"productIsVirtual,omitempty"`
}

// NewOrderOrderDataItem instantiates a new OrderOrderDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderOrderDataItem() *OrderOrderDataItem {
	this := OrderOrderDataItem{}
	return &this
}

// NewOrderOrderDataItemWithDefaults instantiates a new OrderOrderDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderOrderDataItemWithDefaults() *OrderOrderDataItem {
	this := OrderOrderDataItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderOrderDataItem) SetId(v string) {
	o.Id = &v
}

// GetProductGrn returns the ProductGrn field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetProductGrn() string {
	if o == nil || IsNil(o.ProductGrn) {
		var ret string
		return ret
	}
	return *o.ProductGrn
}

// GetProductGrnOk returns a tuple with the ProductGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetProductGrnOk() (*string, bool) {
	if o == nil || IsNil(o.ProductGrn) {
		return nil, false
	}
	return o.ProductGrn, true
}

// HasProductGrn returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasProductGrn() bool {
	if o != nil && !IsNil(o.ProductGrn) {
		return true
	}

	return false
}

// SetProductGrn gets a reference to the given string and assigns it to the ProductGrn field.
func (o *OrderOrderDataItem) SetProductGrn(v string) {
	o.ProductGrn = &v
}

// GetQtyOrdered returns the QtyOrdered field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetQtyOrdered() int64 {
	if o == nil || IsNil(o.QtyOrdered) {
		var ret int64
		return ret
	}
	return *o.QtyOrdered
}

// GetQtyOrderedOk returns a tuple with the QtyOrdered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetQtyOrderedOk() (*int64, bool) {
	if o == nil || IsNil(o.QtyOrdered) {
		return nil, false
	}
	return o.QtyOrdered, true
}

// HasQtyOrdered returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasQtyOrdered() bool {
	if o != nil && !IsNil(o.QtyOrdered) {
		return true
	}

	return false
}

// SetQtyOrdered gets a reference to the given int64 and assigns it to the QtyOrdered field.
func (o *OrderOrderDataItem) SetQtyOrdered(v int64) {
	o.QtyOrdered = &v
}

// GetQtyCommitted returns the QtyCommitted field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetQtyCommitted() int64 {
	if o == nil || IsNil(o.QtyCommitted) {
		var ret int64
		return ret
	}
	return *o.QtyCommitted
}

// GetQtyCommittedOk returns a tuple with the QtyCommitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetQtyCommittedOk() (*int64, bool) {
	if o == nil || IsNil(o.QtyCommitted) {
		return nil, false
	}
	return o.QtyCommitted, true
}

// HasQtyCommitted returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasQtyCommitted() bool {
	if o != nil && !IsNil(o.QtyCommitted) {
		return true
	}

	return false
}

// SetQtyCommitted gets a reference to the given int64 and assigns it to the QtyCommitted field.
func (o *OrderOrderDataItem) SetQtyCommitted(v int64) {
	o.QtyCommitted = &v
}

// GetUnitSalePrice returns the UnitSalePrice field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetUnitSalePrice() OrderMoney {
	if o == nil || IsNil(o.UnitSalePrice) {
		var ret OrderMoney
		return ret
	}
	return *o.UnitSalePrice
}

// GetUnitSalePriceOk returns a tuple with the UnitSalePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetUnitSalePriceOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.UnitSalePrice) {
		return nil, false
	}
	return o.UnitSalePrice, true
}

// HasUnitSalePrice returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasUnitSalePrice() bool {
	if o != nil && !IsNil(o.UnitSalePrice) {
		return true
	}

	return false
}

// SetUnitSalePrice gets a reference to the given OrderMoney and assigns it to the UnitSalePrice field.
func (o *OrderOrderDataItem) SetUnitSalePrice(v OrderMoney) {
	o.UnitSalePrice = &v
}

// GetUnitListPrice returns the UnitListPrice field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetUnitListPrice() OrderMoney {
	if o == nil || IsNil(o.UnitListPrice) {
		var ret OrderMoney
		return ret
	}
	return *o.UnitListPrice
}

// GetUnitListPriceOk returns a tuple with the UnitListPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetUnitListPriceOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.UnitListPrice) {
		return nil, false
	}
	return o.UnitListPrice, true
}

// HasUnitListPrice returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasUnitListPrice() bool {
	if o != nil && !IsNil(o.UnitListPrice) {
		return true
	}

	return false
}

// SetUnitListPrice gets a reference to the given OrderMoney and assigns it to the UnitListPrice field.
func (o *OrderOrderDataItem) SetUnitListPrice(v OrderMoney) {
	o.UnitListPrice = &v
}

// GetUnitBasePrice returns the UnitBasePrice field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetUnitBasePrice() OrderMoney {
	if o == nil || IsNil(o.UnitBasePrice) {
		var ret OrderMoney
		return ret
	}
	return *o.UnitBasePrice
}

// GetUnitBasePriceOk returns a tuple with the UnitBasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetUnitBasePriceOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.UnitBasePrice) {
		return nil, false
	}
	return o.UnitBasePrice, true
}

// HasUnitBasePrice returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasUnitBasePrice() bool {
	if o != nil && !IsNil(o.UnitBasePrice) {
		return true
	}

	return false
}

// SetUnitBasePrice gets a reference to the given OrderMoney and assigns it to the UnitBasePrice field.
func (o *OrderOrderDataItem) SetUnitBasePrice(v OrderMoney) {
	o.UnitBasePrice = &v
}

// GetUnitVatAmount returns the UnitVatAmount field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetUnitVatAmount() OrderMoney {
	if o == nil || IsNil(o.UnitVatAmount) {
		var ret OrderMoney
		return ret
	}
	return *o.UnitVatAmount
}

// GetUnitVatAmountOk returns a tuple with the UnitVatAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetUnitVatAmountOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.UnitVatAmount) {
		return nil, false
	}
	return o.UnitVatAmount, true
}

// HasUnitVatAmount returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasUnitVatAmount() bool {
	if o != nil && !IsNil(o.UnitVatAmount) {
		return true
	}

	return false
}

// SetUnitVatAmount gets a reference to the given OrderMoney and assigns it to the UnitVatAmount field.
func (o *OrderOrderDataItem) SetUnitVatAmount(v OrderMoney) {
	o.UnitVatAmount = &v
}

// GetRowSalePrice returns the RowSalePrice field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetRowSalePrice() OrderMoney {
	if o == nil || IsNil(o.RowSalePrice) {
		var ret OrderMoney
		return ret
	}
	return *o.RowSalePrice
}

// GetRowSalePriceOk returns a tuple with the RowSalePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetRowSalePriceOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.RowSalePrice) {
		return nil, false
	}
	return o.RowSalePrice, true
}

// HasRowSalePrice returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasRowSalePrice() bool {
	if o != nil && !IsNil(o.RowSalePrice) {
		return true
	}

	return false
}

// SetRowSalePrice gets a reference to the given OrderMoney and assigns it to the RowSalePrice field.
func (o *OrderOrderDataItem) SetRowSalePrice(v OrderMoney) {
	o.RowSalePrice = &v
}

// GetRowListPrice returns the RowListPrice field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetRowListPrice() OrderMoney {
	if o == nil || IsNil(o.RowListPrice) {
		var ret OrderMoney
		return ret
	}
	return *o.RowListPrice
}

// GetRowListPriceOk returns a tuple with the RowListPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetRowListPriceOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.RowListPrice) {
		return nil, false
	}
	return o.RowListPrice, true
}

// HasRowListPrice returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasRowListPrice() bool {
	if o != nil && !IsNil(o.RowListPrice) {
		return true
	}

	return false
}

// SetRowListPrice gets a reference to the given OrderMoney and assigns it to the RowListPrice field.
func (o *OrderOrderDataItem) SetRowListPrice(v OrderMoney) {
	o.RowListPrice = &v
}

// GetRowVatAmount returns the RowVatAmount field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetRowVatAmount() OrderMoney {
	if o == nil || IsNil(o.RowVatAmount) {
		var ret OrderMoney
		return ret
	}
	return *o.RowVatAmount
}

// GetRowVatAmountOk returns a tuple with the RowVatAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetRowVatAmountOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.RowVatAmount) {
		return nil, false
	}
	return o.RowVatAmount, true
}

// HasRowVatAmount returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasRowVatAmount() bool {
	if o != nil && !IsNil(o.RowVatAmount) {
		return true
	}

	return false
}

// SetRowVatAmount gets a reference to the given OrderMoney and assigns it to the RowVatAmount field.
func (o *OrderOrderDataItem) SetRowVatAmount(v OrderMoney) {
	o.RowVatAmount = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetDiscountAmount() OrderMoney {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret OrderMoney
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetDiscountAmountOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given OrderMoney and assigns it to the DiscountAmount field.
func (o *OrderOrderDataItem) SetDiscountAmount(v OrderMoney) {
	o.DiscountAmount = &v
}

// GetRowBasePrice returns the RowBasePrice field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetRowBasePrice() OrderMoney {
	if o == nil || IsNil(o.RowBasePrice) {
		var ret OrderMoney
		return ret
	}
	return *o.RowBasePrice
}

// GetRowBasePriceOk returns a tuple with the RowBasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetRowBasePriceOk() (*OrderMoney, bool) {
	if o == nil || IsNil(o.RowBasePrice) {
		return nil, false
	}
	return o.RowBasePrice, true
}

// HasRowBasePrice returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasRowBasePrice() bool {
	if o != nil && !IsNil(o.RowBasePrice) {
		return true
	}

	return false
}

// SetRowBasePrice gets a reference to the given OrderMoney and assigns it to the RowBasePrice field.
func (o *OrderOrderDataItem) SetRowBasePrice(v OrderMoney) {
	o.RowBasePrice = &v
}

// GetVatPercentage returns the VatPercentage field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetVatPercentage() float32 {
	if o == nil || IsNil(o.VatPercentage) {
		var ret float32
		return ret
	}
	return *o.VatPercentage
}

// GetVatPercentageOk returns a tuple with the VatPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetVatPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.VatPercentage) {
		return nil, false
	}
	return o.VatPercentage, true
}

// HasVatPercentage returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasVatPercentage() bool {
	if o != nil && !IsNil(o.VatPercentage) {
		return true
	}

	return false
}

// SetVatPercentage gets a reference to the given float32 and assigns it to the VatPercentage field.
func (o *OrderOrderDataItem) SetVatPercentage(v float32) {
	o.VatPercentage = &v
}

// GetVatInaccurate returns the VatInaccurate field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetVatInaccurate() bool {
	if o == nil || IsNil(o.VatInaccurate) {
		var ret bool
		return ret
	}
	return *o.VatInaccurate
}

// GetVatInaccurateOk returns a tuple with the VatInaccurate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetVatInaccurateOk() (*bool, bool) {
	if o == nil || IsNil(o.VatInaccurate) {
		return nil, false
	}
	return o.VatInaccurate, true
}

// HasVatInaccurate returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasVatInaccurate() bool {
	if o != nil && !IsNil(o.VatInaccurate) {
		return true
	}

	return false
}

// SetVatInaccurate gets a reference to the given bool and assigns it to the VatInaccurate field.
func (o *OrderOrderDataItem) SetVatInaccurate(v bool) {
	o.VatInaccurate = &v
}

// GetVatCalculated returns the VatCalculated field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetVatCalculated() bool {
	if o == nil || IsNil(o.VatCalculated) {
		var ret bool
		return ret
	}
	return *o.VatCalculated
}

// GetVatCalculatedOk returns a tuple with the VatCalculated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetVatCalculatedOk() (*bool, bool) {
	if o == nil || IsNil(o.VatCalculated) {
		return nil, false
	}
	return o.VatCalculated, true
}

// HasVatCalculated returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasVatCalculated() bool {
	if o != nil && !IsNil(o.VatCalculated) {
		return true
	}

	return false
}

// SetVatCalculated gets a reference to the given bool and assigns it to the VatCalculated field.
func (o *OrderOrderDataItem) SetVatCalculated(v bool) {
	o.VatCalculated = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *OrderOrderDataItem) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *OrderOrderDataItem) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetProductSku returns the ProductSku field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetProductSku() string {
	if o == nil || IsNil(o.ProductSku) {
		var ret string
		return ret
	}
	return *o.ProductSku
}

// GetProductSkuOk returns a tuple with the ProductSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetProductSkuOk() (*string, bool) {
	if o == nil || IsNil(o.ProductSku) {
		return nil, false
	}
	return o.ProductSku, true
}

// HasProductSku returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasProductSku() bool {
	if o != nil && !IsNil(o.ProductSku) {
		return true
	}

	return false
}

// SetProductSku gets a reference to the given string and assigns it to the ProductSku field.
func (o *OrderOrderDataItem) SetProductSku(v string) {
	o.ProductSku = &v
}

// GetProductOptions returns the ProductOptions field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetProductOptions() string {
	if o == nil || IsNil(o.ProductOptions) {
		var ret string
		return ret
	}
	return *o.ProductOptions
}

// GetProductOptionsOk returns a tuple with the ProductOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetProductOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.ProductOptions) {
		return nil, false
	}
	return o.ProductOptions, true
}

// HasProductOptions returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasProductOptions() bool {
	if o != nil && !IsNil(o.ProductOptions) {
		return true
	}

	return false
}

// SetProductOptions gets a reference to the given string and assigns it to the ProductOptions field.
func (o *OrderOrderDataItem) SetProductOptions(v string) {
	o.ProductOptions = &v
}

// GetProductImg returns the ProductImg field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetProductImg() string {
	if o == nil || IsNil(o.ProductImg) {
		var ret string
		return ret
	}
	return *o.ProductImg
}

// GetProductImgOk returns a tuple with the ProductImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetProductImgOk() (*string, bool) {
	if o == nil || IsNil(o.ProductImg) {
		return nil, false
	}
	return o.ProductImg, true
}

// HasProductImg returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasProductImg() bool {
	if o != nil && !IsNil(o.ProductImg) {
		return true
	}

	return false
}

// SetProductImg gets a reference to the given string and assigns it to the ProductImg field.
func (o *OrderOrderDataItem) SetProductImg(v string) {
	o.ProductImg = &v
}

// GetProductData returns the ProductData field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetProductData() string {
	if o == nil || IsNil(o.ProductData) {
		var ret string
		return ret
	}
	return *o.ProductData
}

// GetProductDataOk returns a tuple with the ProductData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetProductDataOk() (*string, bool) {
	if o == nil || IsNil(o.ProductData) {
		return nil, false
	}
	return o.ProductData, true
}

// HasProductData returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasProductData() bool {
	if o != nil && !IsNil(o.ProductData) {
		return true
	}

	return false
}

// SetProductData gets a reference to the given string and assigns it to the ProductData field.
func (o *OrderOrderDataItem) SetProductData(v string) {
	o.ProductData = &v
}

// GetShipmentInfoReference returns the ShipmentInfoReference field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetShipmentInfoReference() string {
	if o == nil || IsNil(o.ShipmentInfoReference) {
		var ret string
		return ret
	}
	return *o.ShipmentInfoReference
}

// GetShipmentInfoReferenceOk returns a tuple with the ShipmentInfoReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetShipmentInfoReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentInfoReference) {
		return nil, false
	}
	return o.ShipmentInfoReference, true
}

// HasShipmentInfoReference returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasShipmentInfoReference() bool {
	if o != nil && !IsNil(o.ShipmentInfoReference) {
		return true
	}

	return false
}

// SetShipmentInfoReference gets a reference to the given string and assigns it to the ShipmentInfoReference field.
func (o *OrderOrderDataItem) SetShipmentInfoReference(v string) {
	o.ShipmentInfoReference = &v
}

// GetPromotionGrn returns the PromotionGrn field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetPromotionGrn() []string {
	if o == nil || IsNil(o.PromotionGrn) {
		var ret []string
		return ret
	}
	return o.PromotionGrn
}

// GetPromotionGrnOk returns a tuple with the PromotionGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetPromotionGrnOk() ([]string, bool) {
	if o == nil || IsNil(o.PromotionGrn) {
		return nil, false
	}
	return o.PromotionGrn, true
}

// HasPromotionGrn returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasPromotionGrn() bool {
	if o != nil && !IsNil(o.PromotionGrn) {
		return true
	}

	return false
}

// SetPromotionGrn gets a reference to the given []string and assigns it to the PromotionGrn field.
func (o *OrderOrderDataItem) SetPromotionGrn(v []string) {
	o.PromotionGrn = v
}

// GetProductIsVirtual returns the ProductIsVirtual field value if set, zero value otherwise.
func (o *OrderOrderDataItem) GetProductIsVirtual() bool {
	if o == nil || IsNil(o.ProductIsVirtual) {
		var ret bool
		return ret
	}
	return *o.ProductIsVirtual
}

// GetProductIsVirtualOk returns a tuple with the ProductIsVirtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOrderDataItem) GetProductIsVirtualOk() (*bool, bool) {
	if o == nil || IsNil(o.ProductIsVirtual) {
		return nil, false
	}
	return o.ProductIsVirtual, true
}

// HasProductIsVirtual returns a boolean if a field has been set.
func (o *OrderOrderDataItem) HasProductIsVirtual() bool {
	if o != nil && !IsNil(o.ProductIsVirtual) {
		return true
	}

	return false
}

// SetProductIsVirtual gets a reference to the given bool and assigns it to the ProductIsVirtual field.
func (o *OrderOrderDataItem) SetProductIsVirtual(v bool) {
	o.ProductIsVirtual = &v
}

func (o OrderOrderDataItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderOrderDataItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProductGrn) {
		toSerialize["productGrn"] = o.ProductGrn
	}
	if !IsNil(o.QtyOrdered) {
		toSerialize["qtyOrdered"] = o.QtyOrdered
	}
	if !IsNil(o.QtyCommitted) {
		toSerialize["qtyCommitted"] = o.QtyCommitted
	}
	if !IsNil(o.UnitSalePrice) {
		toSerialize["unitSalePrice"] = o.UnitSalePrice
	}
	if !IsNil(o.UnitListPrice) {
		toSerialize["unitListPrice"] = o.UnitListPrice
	}
	if !IsNil(o.UnitBasePrice) {
		toSerialize["unitBasePrice"] = o.UnitBasePrice
	}
	if !IsNil(o.UnitVatAmount) {
		toSerialize["unitVatAmount"] = o.UnitVatAmount
	}
	if !IsNil(o.RowSalePrice) {
		toSerialize["rowSalePrice"] = o.RowSalePrice
	}
	if !IsNil(o.RowListPrice) {
		toSerialize["rowListPrice"] = o.RowListPrice
	}
	if !IsNil(o.RowVatAmount) {
		toSerialize["rowVatAmount"] = o.RowVatAmount
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discountAmount"] = o.DiscountAmount
	}
	if !IsNil(o.RowBasePrice) {
		toSerialize["rowBasePrice"] = o.RowBasePrice
	}
	if !IsNil(o.VatPercentage) {
		toSerialize["vatPercentage"] = o.VatPercentage
	}
	if !IsNil(o.VatInaccurate) {
		toSerialize["vatInaccurate"] = o.VatInaccurate
	}
	if !IsNil(o.VatCalculated) {
		toSerialize["vatCalculated"] = o.VatCalculated
	}
	if !IsNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !IsNil(o.ProductCode) {
		toSerialize["productCode"] = o.ProductCode
	}
	if !IsNil(o.ProductSku) {
		toSerialize["productSku"] = o.ProductSku
	}
	if !IsNil(o.ProductOptions) {
		toSerialize["productOptions"] = o.ProductOptions
	}
	if !IsNil(o.ProductImg) {
		toSerialize["productImg"] = o.ProductImg
	}
	if !IsNil(o.ProductData) {
		toSerialize["productData"] = o.ProductData
	}
	if !IsNil(o.ShipmentInfoReference) {
		toSerialize["shipmentInfoReference"] = o.ShipmentInfoReference
	}
	if !IsNil(o.PromotionGrn) {
		toSerialize["promotionGrn"] = o.PromotionGrn
	}
	if !IsNil(o.ProductIsVirtual) {
		toSerialize["productIsVirtual"] = o.ProductIsVirtual
	}
	return toSerialize, nil
}

type NullableOrderOrderDataItem struct {
	value *OrderOrderDataItem
	isSet bool
}

func (v NullableOrderOrderDataItem) Get() *OrderOrderDataItem {
	return v.value
}

func (v *NullableOrderOrderDataItem) Set(val *OrderOrderDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderOrderDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderOrderDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderOrderDataItem(val *OrderOrderDataItem) *NullableOrderOrderDataItem {
	return &NullableOrderOrderDataItem{value: val, isSet: true}
}

func (v NullableOrderOrderDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderOrderDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



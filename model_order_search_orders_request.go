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
	"fmt"
	"time"
)

// checks if the OrderSearchOrdersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSearchOrdersRequest{}

// OrderSearchOrdersRequest struct for OrderSearchOrdersRequest
type OrderSearchOrdersRequest struct {
	TenantId    string  `json:"tenantId"`
	SearchQuery *string `json:"searchQuery,omitempty"`
	// The maximum number of orders to return. The service may return fewer than this value. If unspecified, at most 10 orders will be returned. The maximum value is 100; values above 100 will be coerced to 100.
	PageSize *int64 `json:"pageSize,omitempty"`
	// A page token, received from a previous `ListOrders` call. Provide this to retrieve the subsequent page.   When paginating, all other parameters provided to `ListOrders` must match the call that provided the page token.
	PageToken            *string             `json:"pageToken,omitempty"`
	OrderBy              []OrderOrderBy      `json:"orderBy,omitempty"`
	StatusFilter         *OrderStatusFilter  `json:"statusFilter,omitempty"`
	FromDate             *time.Time          `json:"fromDate,omitempty"`
	ToDate               *time.Time          `json:"toDate,omitempty"`
	PaymentFilter        *OrderPaymentFilter `json:"paymentFilter,omitempty"`
	AgentGrn             *string             `json:"agentGrn,omitempty"`
	UpdatedAtFrom        *time.Time          `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo          *time.Time          `json:"updatedAtTo,omitempty"`
	OnHold               *bool               `json:"onHold,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderSearchOrdersRequest OrderSearchOrdersRequest

// NewOrderSearchOrdersRequest instantiates a new OrderSearchOrdersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchOrdersRequest(tenantId string) *OrderSearchOrdersRequest {
	this := OrderSearchOrdersRequest{}
	this.TenantId = tenantId
	return &this
}

// NewOrderSearchOrdersRequestWithDefaults instantiates a new OrderSearchOrdersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchOrdersRequestWithDefaults() *OrderSearchOrdersRequest {
	this := OrderSearchOrdersRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *OrderSearchOrdersRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *OrderSearchOrdersRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetSearchQuery returns the SearchQuery field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetSearchQuery() string {
	if o == nil || IsNil(o.SearchQuery) {
		var ret string
		return ret
	}
	return *o.SearchQuery
}

// GetSearchQueryOk returns a tuple with the SearchQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetSearchQueryOk() (*string, bool) {
	if o == nil || IsNil(o.SearchQuery) {
		return nil, false
	}
	return o.SearchQuery, true
}

// HasSearchQuery returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) IsSetSearchQuery() bool {
	if o != nil && !IsNil(o.SearchQuery) {
		return true
	}

	return false
}

// SetSearchQuery gets a reference to the given string and assigns it to the SearchQuery field.
func (o *OrderSearchOrdersRequest) SetSearchQuery(v string) {
	o.SearchQuery = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) IsSetPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *OrderSearchOrdersRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) IsSetPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *OrderSearchOrdersRequest) SetPageToken(v string) {
	o.PageToken = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetOrderBy() []OrderOrderBy {
	if o == nil || IsNil(o.OrderBy) {
		var ret []OrderOrderBy
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetOrderByOk() ([]OrderOrderBy, bool) {
	if o == nil || IsNil(o.OrderBy) {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) IsSetOrderBy() bool {
	if o != nil && !IsNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given []OrderOrderBy and assigns it to the OrderBy field.
func (o *OrderSearchOrdersRequest) SetOrderBy(v []OrderOrderBy) {
	o.OrderBy = v
}

// GetStatusFilter returns the StatusFilter field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetStatusFilter() OrderStatusFilter {
	if o == nil || IsNil(o.StatusFilter) {
		var ret OrderStatusFilter
		return ret
	}
	return *o.StatusFilter
}

// GetStatusFilterOk returns a tuple with the StatusFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetStatusFilterOk() (*OrderStatusFilter, bool) {
	if o == nil || IsNil(o.StatusFilter) {
		return nil, false
	}
	return o.StatusFilter, true
}

// HasStatusFilter returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) IsSetStatusFilter() bool {
	if o != nil && !IsNil(o.StatusFilter) {
		return true
	}

	return false
}

// SetStatusFilter gets a reference to the given OrderStatusFilter and assigns it to the StatusFilter field.
func (o *OrderSearchOrdersRequest) SetStatusFilter(v OrderStatusFilter) {
	o.StatusFilter = &v
}

// GetFromDate returns the FromDate field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetFromDate() time.Time {
	if o == nil || IsNil(o.FromDate) {
		var ret time.Time
		return ret
	}
	return *o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetFromDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FromDate) {
		return nil, false
	}
	return o.FromDate, true
}

// HasFromDate returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) IsSetFromDate() bool {
	if o != nil && !IsNil(o.FromDate) {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given time.Time and assigns it to the FromDate field.
func (o *OrderSearchOrdersRequest) SetFromDate(v time.Time) {
	o.FromDate = &v
}

// GetToDate returns the ToDate field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetToDate() time.Time {
	if o == nil || IsNil(o.ToDate) {
		var ret time.Time
		return ret
	}
	return *o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetToDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ToDate) {
		return nil, false
	}
	return o.ToDate, true
}

// HasToDate returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) IsSetToDate() bool {
	if o != nil && !IsNil(o.ToDate) {
		return true
	}

	return false
}

// SetToDate gets a reference to the given time.Time and assigns it to the ToDate field.
func (o *OrderSearchOrdersRequest) SetToDate(v time.Time) {
	o.ToDate = &v
}

// GetPaymentFilter returns the PaymentFilter field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetPaymentFilter() OrderPaymentFilter {
	if o == nil || IsNil(o.PaymentFilter) {
		var ret OrderPaymentFilter
		return ret
	}
	return *o.PaymentFilter
}

// GetPaymentFilterOk returns a tuple with the PaymentFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetPaymentFilterOk() (*OrderPaymentFilter, bool) {
	if o == nil || IsNil(o.PaymentFilter) {
		return nil, false
	}
	return o.PaymentFilter, true
}

// HasPaymentFilter returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) IsSetPaymentFilter() bool {
	if o != nil && !IsNil(o.PaymentFilter) {
		return true
	}

	return false
}

// SetPaymentFilter gets a reference to the given OrderPaymentFilter and assigns it to the PaymentFilter field.
func (o *OrderSearchOrdersRequest) SetPaymentFilter(v OrderPaymentFilter) {
	o.PaymentFilter = &v
}

// GetAgentGrn returns the AgentGrn field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetAgentGrn() string {
	if o == nil || IsNil(o.AgentGrn) {
		var ret string
		return ret
	}
	return *o.AgentGrn
}

// GetAgentGrnOk returns a tuple with the AgentGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetAgentGrnOk() (*string, bool) {
	if o == nil || IsNil(o.AgentGrn) {
		return nil, false
	}
	return o.AgentGrn, true
}

// HasAgentGrn returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) IsSetAgentGrn() bool {
	if o != nil && !IsNil(o.AgentGrn) {
		return true
	}

	return false
}

// SetAgentGrn gets a reference to the given string and assigns it to the AgentGrn field.
func (o *OrderSearchOrdersRequest) SetAgentGrn(v string) {
	o.AgentGrn = &v
}

// GetUpdatedAtFrom returns the UpdatedAtFrom field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetUpdatedAtFrom() time.Time {
	if o == nil || IsNil(o.UpdatedAtFrom) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtFrom
}

// GetUpdatedAtFromOk returns a tuple with the UpdatedAtFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetUpdatedAtFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAtFrom) {
		return nil, false
	}
	return o.UpdatedAtFrom, true
}

// HasUpdatedAtFrom returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) IsSetUpdatedAtFrom() bool {
	if o != nil && !IsNil(o.UpdatedAtFrom) {
		return true
	}

	return false
}

// SetUpdatedAtFrom gets a reference to the given time.Time and assigns it to the UpdatedAtFrom field.
func (o *OrderSearchOrdersRequest) SetUpdatedAtFrom(v time.Time) {
	o.UpdatedAtFrom = &v
}

// GetUpdatedAtTo returns the UpdatedAtTo field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetUpdatedAtTo() time.Time {
	if o == nil || IsNil(o.UpdatedAtTo) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtTo
}

// GetUpdatedAtToOk returns a tuple with the UpdatedAtTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetUpdatedAtToOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAtTo) {
		return nil, false
	}
	return o.UpdatedAtTo, true
}

// HasUpdatedAtTo returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) IsSetUpdatedAtTo() bool {
	if o != nil && !IsNil(o.UpdatedAtTo) {
		return true
	}

	return false
}

// SetUpdatedAtTo gets a reference to the given time.Time and assigns it to the UpdatedAtTo field.
func (o *OrderSearchOrdersRequest) SetUpdatedAtTo(v time.Time) {
	o.UpdatedAtTo = &v
}

// GetOnHold returns the OnHold field value if set, zero value otherwise.
func (o *OrderSearchOrdersRequest) GetOnHold() bool {
	if o == nil || IsNil(o.OnHold) {
		var ret bool
		return ret
	}
	return *o.OnHold
}

// GetOnHoldOk returns a tuple with the OnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchOrdersRequest) GetOnHoldOk() (*bool, bool) {
	if o == nil || IsNil(o.OnHold) {
		return nil, false
	}
	return o.OnHold, true
}

// HasOnHold returns a boolean if a field has been set.
func (o *OrderSearchOrdersRequest) IsSetOnHold() bool {
	if o != nil && !IsNil(o.OnHold) {
		return true
	}

	return false
}

// SetOnHold gets a reference to the given bool and assigns it to the OnHold field.
func (o *OrderSearchOrdersRequest) SetOnHold(v bool) {
	o.OnHold = &v
}

func (o OrderSearchOrdersRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSearchOrdersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	if !IsNil(o.SearchQuery) {
		toSerialize["searchQuery"] = o.SearchQuery
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	if !IsNil(o.OrderBy) {
		toSerialize["orderBy"] = o.OrderBy
	}
	if !IsNil(o.StatusFilter) {
		toSerialize["statusFilter"] = o.StatusFilter
	}
	if !IsNil(o.FromDate) {
		toSerialize["fromDate"] = o.FromDate
	}
	if !IsNil(o.ToDate) {
		toSerialize["toDate"] = o.ToDate
	}
	if !IsNil(o.PaymentFilter) {
		toSerialize["paymentFilter"] = o.PaymentFilter
	}
	if !IsNil(o.AgentGrn) {
		toSerialize["agentGrn"] = o.AgentGrn
	}
	if !IsNil(o.UpdatedAtFrom) {
		toSerialize["updatedAtFrom"] = o.UpdatedAtFrom
	}
	if !IsNil(o.UpdatedAtTo) {
		toSerialize["updatedAtTo"] = o.UpdatedAtTo
	}
	if !IsNil(o.OnHold) {
		toSerialize["onHold"] = o.OnHold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderSearchOrdersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderSearchOrdersRequest := _OrderSearchOrdersRequest{}

	err = json.Unmarshal(data, &varOrderSearchOrdersRequest)

	if err != nil {
		return err
	}

	*o = OrderSearchOrdersRequest(varOrderSearchOrdersRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "searchQuery")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "pageToken")
		delete(additionalProperties, "orderBy")
		delete(additionalProperties, "statusFilter")
		delete(additionalProperties, "fromDate")
		delete(additionalProperties, "toDate")
		delete(additionalProperties, "paymentFilter")
		delete(additionalProperties, "agentGrn")
		delete(additionalProperties, "updatedAtFrom")
		delete(additionalProperties, "updatedAtTo")
		delete(additionalProperties, "onHold")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderSearchOrdersRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *OrderSearchOrdersRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableOrderSearchOrdersRequest struct {
	value *OrderSearchOrdersRequest
	isSet bool
}

func (v NullableOrderSearchOrdersRequest) Get() *OrderSearchOrdersRequest {
	return v.value
}

func (v *NullableOrderSearchOrdersRequest) Set(val *OrderSearchOrdersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchOrdersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchOrdersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchOrdersRequest(val *OrderSearchOrdersRequest) *NullableOrderSearchOrdersRequest {
	return &NullableOrderSearchOrdersRequest{value: val, isSet: true}
}

func (v NullableOrderSearchOrdersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchOrdersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

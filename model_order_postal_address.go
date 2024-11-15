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

// checks if the OrderPostalAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderPostalAddress{}

// OrderPostalAddress Represents a postal address, e.g. for postal delivery or payments addresses. Given a postal address, a postal service can deliver items to a premise, P.O. Box or similar. It is not intended to model geographical locations (roads, towns, mountains).  In typical usage an address would be created via user input or from importing existing data, depending on the type of process.  Advice on address input / editing:  - Use an i18n-ready address widget such as    https://github.com/google/libaddressinput) - Users should not be presented with UI elements for input or editing of   fields outside countries where that field is used.  For more guidance on how to use this schema, please see: https://support.google.com/business/answer/6397478
type OrderPostalAddress struct {
	// The schema revision of the `PostalAddress`. This must be set to 0, which is the latest revision.  All new revisions **must** be backward compatible with old revisions.
	Revision *int32 `json:"revision,omitempty"`
	// Required. CLDR region code of the country/region of the address. This is never inferred and it is up to the user to ensure the value is correct. See http://cldr.unicode.org/ and http://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: \"CH\" for Switzerland.
	RegionCode *string `json:"regionCode,omitempty"`
	// Optional. BCP-47 language code of the contents of this address (if known). This is often the UI language of the input form or is expected to match one of the languages used in the address' country/region, or their transliterated equivalents. This can affect formatting in certain countries, but is not critical to the correctness of the data and will never affect any validation or other non-formatting related operations.  If this value is not known, it should be omitted (rather than specifying a possibly incorrect default).  Examples: \"zh-Hant\", \"ja\", \"ja-Latn\", \"en\".
	LanguageCode *string `json:"languageCode,omitempty"`
	// Optional. Postal code of the address. Not all countries use or require postal codes to be present, but where they are used, they may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).
	PostalCode *string `json:"postalCode,omitempty"`
	// Optional. Additional, country-specific, sorting code. This is not used in most regions. Where it is used, the value is either a string like \"CEDEX\", optionally followed by a number (e.g. \"CEDEX 7\"), or just a number alone, representing the \"sector code\" (Jamaica), \"delivery area indicator\" (Malawi) or \"post office indicator\" (e.g. Côte d'Ivoire).
	SortingCode *string `json:"sortingCode,omitempty"`
	// Optional. Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state, a province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community (e.g. \"Barcelona\" and not \"Catalonia\"). Many countries don't use an administrative area in postal addresses. E.g. in Switzerland this should be left unpopulated.
	AdministrativeArea *string `json:"administrativeArea,omitempty"`
	// Optional. Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world where localities are not well defined or do not fit into this structure well, leave locality empty and use address_lines.
	Locality *string `json:"locality,omitempty"`
	// Optional. Sublocality of the address. For example, this can be neighborhoods, boroughs, districts.
	Sublocality *string `json:"sublocality,omitempty"`
	// Unstructured address lines describing the lower levels of an address.  Because values in address_lines do not have type information and may sometimes contain multiple values in a single field (e.g. \"Austin, TX\"), it is important that the line order is clear. The order of address lines should be \"envelope order\" for the country/region of the address. In places where this can vary (e.g. Japan), address_language is used to make it explicit (e.g. \"ja\" for large-to-small ordering and \"ja-Latn\" or \"en\" for small-to-large). This way, the most specific line of an address can be selected based on the language.  The minimum permitted structural representation of an address consists of a region_code with all remaining information placed in the address_lines. It would be possible to format such an address very approximately without geocoding, but no semantic reasoning could be made about any of the address components until it was at least partially resolved.  Creating an address only containing a region_code and address_lines, and then geocoding is the recommended way to handle completely unstructured addresses (as opposed to guessing which parts of the address should be localities or administrative areas).
	AddressLines []string `json:"addressLines,omitempty"`
	// Optional. The recipient at the address. This field may, under certain circumstances, contain multiline information. For example, it might contain \"care of\" information.
	Recipients []string `json:"recipients,omitempty"`
	// Optional. The name of the organization at the address.
	Organization *string `json:"organization,omitempty"`
	// Optional.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Optional.
	AdditionalInfo map[string]interface{} `json:"additionalInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderPostalAddress OrderPostalAddress

// NewOrderPostalAddress instantiates a new OrderPostalAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderPostalAddress() *OrderPostalAddress {
	this := OrderPostalAddress{}
	return &this
}

// NewOrderPostalAddressWithDefaults instantiates a new OrderPostalAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderPostalAddressWithDefaults() *OrderPostalAddress {
	this := OrderPostalAddress{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetRevision() int32 {
	if o == nil || IsNil(o.Revision) {
		var ret int32
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetRevisionOk() (*int32, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given int32 and assigns it to the Revision field.
func (o *OrderPostalAddress) SetRevision(v int32) {
	o.Revision = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetRegionCode() string {
	if o == nil || IsNil(o.RegionCode) {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetRegionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RegionCode) {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasRegionCode() bool {
	if o != nil && !IsNil(o.RegionCode) {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *OrderPostalAddress) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetLanguageCode() string {
	if o == nil || IsNil(o.LanguageCode) {
		var ret string
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetLanguageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageCode) {
		return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasLanguageCode() bool {
	if o != nil && !IsNil(o.LanguageCode) {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given string and assigns it to the LanguageCode field.
func (o *OrderPostalAddress) SetLanguageCode(v string) {
	o.LanguageCode = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *OrderPostalAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetSortingCode returns the SortingCode field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetSortingCode() string {
	if o == nil || IsNil(o.SortingCode) {
		var ret string
		return ret
	}
	return *o.SortingCode
}

// GetSortingCodeOk returns a tuple with the SortingCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetSortingCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SortingCode) {
		return nil, false
	}
	return o.SortingCode, true
}

// HasSortingCode returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasSortingCode() bool {
	if o != nil && !IsNil(o.SortingCode) {
		return true
	}

	return false
}

// SetSortingCode gets a reference to the given string and assigns it to the SortingCode field.
func (o *OrderPostalAddress) SetSortingCode(v string) {
	o.SortingCode = &v
}

// GetAdministrativeArea returns the AdministrativeArea field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetAdministrativeArea() string {
	if o == nil || IsNil(o.AdministrativeArea) {
		var ret string
		return ret
	}
	return *o.AdministrativeArea
}

// GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetAdministrativeAreaOk() (*string, bool) {
	if o == nil || IsNil(o.AdministrativeArea) {
		return nil, false
	}
	return o.AdministrativeArea, true
}

// HasAdministrativeArea returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasAdministrativeArea() bool {
	if o != nil && !IsNil(o.AdministrativeArea) {
		return true
	}

	return false
}

// SetAdministrativeArea gets a reference to the given string and assigns it to the AdministrativeArea field.
func (o *OrderPostalAddress) SetAdministrativeArea(v string) {
	o.AdministrativeArea = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetLocality() string {
	if o == nil || IsNil(o.Locality) {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetLocalityOk() (*string, bool) {
	if o == nil || IsNil(o.Locality) {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasLocality() bool {
	if o != nil && !IsNil(o.Locality) {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *OrderPostalAddress) SetLocality(v string) {
	o.Locality = &v
}

// GetSublocality returns the Sublocality field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetSublocality() string {
	if o == nil || IsNil(o.Sublocality) {
		var ret string
		return ret
	}
	return *o.Sublocality
}

// GetSublocalityOk returns a tuple with the Sublocality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetSublocalityOk() (*string, bool) {
	if o == nil || IsNil(o.Sublocality) {
		return nil, false
	}
	return o.Sublocality, true
}

// HasSublocality returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasSublocality() bool {
	if o != nil && !IsNil(o.Sublocality) {
		return true
	}

	return false
}

// SetSublocality gets a reference to the given string and assigns it to the Sublocality field.
func (o *OrderPostalAddress) SetSublocality(v string) {
	o.Sublocality = &v
}

// GetAddressLines returns the AddressLines field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetAddressLines() []string {
	if o == nil || IsNil(o.AddressLines) {
		var ret []string
		return ret
	}
	return o.AddressLines
}

// GetAddressLinesOk returns a tuple with the AddressLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetAddressLinesOk() ([]string, bool) {
	if o == nil || IsNil(o.AddressLines) {
		return nil, false
	}
	return o.AddressLines, true
}

// HasAddressLines returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasAddressLines() bool {
	if o != nil && !IsNil(o.AddressLines) {
		return true
	}

	return false
}

// SetAddressLines gets a reference to the given []string and assigns it to the AddressLines field.
func (o *OrderPostalAddress) SetAddressLines(v []string) {
	o.AddressLines = v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetRecipients() []string {
	if o == nil || IsNil(o.Recipients) {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetRecipientsOk() ([]string, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []string and assigns it to the Recipients field.
func (o *OrderPostalAddress) SetRecipients(v []string) {
	o.Recipients = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *OrderPostalAddress) SetOrganization(v string) {
	o.Organization = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *OrderPostalAddress) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *OrderPostalAddress) GetAdditionalInfo() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostalAddress) GetAdditionalInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *OrderPostalAddress) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given map[string]interface{} and assigns it to the AdditionalInfo field.
func (o *OrderPostalAddress) SetAdditionalInfo(v map[string]interface{}) {
	o.AdditionalInfo = v
}

func (o OrderPostalAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderPostalAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.RegionCode) {
		toSerialize["regionCode"] = o.RegionCode
	}
	if !IsNil(o.LanguageCode) {
		toSerialize["languageCode"] = o.LanguageCode
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.SortingCode) {
		toSerialize["sortingCode"] = o.SortingCode
	}
	if !IsNil(o.AdministrativeArea) {
		toSerialize["administrativeArea"] = o.AdministrativeArea
	}
	if !IsNil(o.Locality) {
		toSerialize["locality"] = o.Locality
	}
	if !IsNil(o.Sublocality) {
		toSerialize["sublocality"] = o.Sublocality
	}
	if !IsNil(o.AddressLines) {
		toSerialize["addressLines"] = o.AddressLines
	}
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderPostalAddress) UnmarshalJSON(data []byte) (err error) {
	varOrderPostalAddress := _OrderPostalAddress{}

	err = json.Unmarshal(data, &varOrderPostalAddress)

	if err != nil {
		return err
	}

	*o = OrderPostalAddress(varOrderPostalAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "revision")
		delete(additionalProperties, "regionCode")
		delete(additionalProperties, "languageCode")
		delete(additionalProperties, "postalCode")
		delete(additionalProperties, "sortingCode")
		delete(additionalProperties, "administrativeArea")
		delete(additionalProperties, "locality")
		delete(additionalProperties, "sublocality")
		delete(additionalProperties, "addressLines")
		delete(additionalProperties, "recipients")
		delete(additionalProperties, "organization")
		delete(additionalProperties, "phoneNumber")
		delete(additionalProperties, "additionalInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *OrderPostalAddress) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *OrderPostalAddress) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableOrderPostalAddress struct {
	value *OrderPostalAddress
	isSet bool
}

func (v NullableOrderPostalAddress) Get() *OrderPostalAddress {
	return v.value
}

func (v *NullableOrderPostalAddress) Set(val *OrderPostalAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderPostalAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderPostalAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderPostalAddress(val *OrderPostalAddress) *NullableOrderPostalAddress {
	return &NullableOrderPostalAddress{value: val, isSet: true}
}

func (v NullableOrderPostalAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderPostalAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



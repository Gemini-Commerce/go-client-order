/*
order/order.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// OrderCurrency the model 'OrderCurrency'
type OrderCurrency string

// List of orderCurrency
const (
	ORDERCURRENCY_XXX OrderCurrency = "XXX"
	ORDERCURRENCY_ALL OrderCurrency = "ALL"
	ORDERCURRENCY_DZD OrderCurrency = "DZD"
	ORDERCURRENCY_ARS OrderCurrency = "ARS"
	ORDERCURRENCY_AUD OrderCurrency = "AUD"
	ORDERCURRENCY_BSD OrderCurrency = "BSD"
	ORDERCURRENCY_BHD OrderCurrency = "BHD"
	ORDERCURRENCY_BDT OrderCurrency = "BDT"
	ORDERCURRENCY_AMD OrderCurrency = "AMD"
	ORDERCURRENCY_BBD OrderCurrency = "BBD"
	ORDERCURRENCY_BMD OrderCurrency = "BMD"
	ORDERCURRENCY_BTN OrderCurrency = "BTN"
	ORDERCURRENCY_BOB OrderCurrency = "BOB"
	ORDERCURRENCY_BWP OrderCurrency = "BWP"
	ORDERCURRENCY_BZD OrderCurrency = "BZD"
	ORDERCURRENCY_SBD OrderCurrency = "SBD"
	ORDERCURRENCY_BND OrderCurrency = "BND"
	ORDERCURRENCY_MMK OrderCurrency = "MMK"
	ORDERCURRENCY_BIF OrderCurrency = "BIF"
	ORDERCURRENCY_KHR OrderCurrency = "KHR"
	ORDERCURRENCY_CAD OrderCurrency = "CAD"
	ORDERCURRENCY_CVE OrderCurrency = "CVE"
	ORDERCURRENCY_KYD OrderCurrency = "KYD"
	ORDERCURRENCY_LKR OrderCurrency = "LKR"
	ORDERCURRENCY_CLP OrderCurrency = "CLP"
	ORDERCURRENCY_CNY OrderCurrency = "CNY"
	ORDERCURRENCY_COP OrderCurrency = "COP"
	ORDERCURRENCY_KMF OrderCurrency = "KMF"
	ORDERCURRENCY_CRC OrderCurrency = "CRC"
	ORDERCURRENCY_HRK OrderCurrency = "HRK"
	ORDERCURRENCY_CUP OrderCurrency = "CUP"
	ORDERCURRENCY_CZK OrderCurrency = "CZK"
	ORDERCURRENCY_DKK OrderCurrency = "DKK"
	ORDERCURRENCY_DOP OrderCurrency = "DOP"
	ORDERCURRENCY_SVC OrderCurrency = "SVC"
	ORDERCURRENCY_ETB OrderCurrency = "ETB"
	ORDERCURRENCY_ERN OrderCurrency = "ERN"
	ORDERCURRENCY_FKP OrderCurrency = "FKP"
	ORDERCURRENCY_FJD OrderCurrency = "FJD"
	ORDERCURRENCY_DJF OrderCurrency = "DJF"
	ORDERCURRENCY_GMD OrderCurrency = "GMD"
	ORDERCURRENCY_GIP OrderCurrency = "GIP"
	ORDERCURRENCY_GTQ OrderCurrency = "GTQ"
	ORDERCURRENCY_GNF OrderCurrency = "GNF"
	ORDERCURRENCY_GYD OrderCurrency = "GYD"
	ORDERCURRENCY_HTG OrderCurrency = "HTG"
	ORDERCURRENCY_HNL OrderCurrency = "HNL"
	ORDERCURRENCY_HKD OrderCurrency = "HKD"
	ORDERCURRENCY_HUF OrderCurrency = "HUF"
	ORDERCURRENCY_ISK OrderCurrency = "ISK"
	ORDERCURRENCY_INR OrderCurrency = "INR"
	ORDERCURRENCY_IDR OrderCurrency = "IDR"
	ORDERCURRENCY_IRR OrderCurrency = "IRR"
	ORDERCURRENCY_IQD OrderCurrency = "IQD"
	ORDERCURRENCY_ILS OrderCurrency = "ILS"
	ORDERCURRENCY_JMD OrderCurrency = "JMD"
	ORDERCURRENCY_JPY OrderCurrency = "JPY"
	ORDERCURRENCY_KZT OrderCurrency = "KZT"
	ORDERCURRENCY_JOD OrderCurrency = "JOD"
	ORDERCURRENCY_KES OrderCurrency = "KES"
	ORDERCURRENCY_KPW OrderCurrency = "KPW"
	ORDERCURRENCY_KRW OrderCurrency = "KRW"
	ORDERCURRENCY_KWD OrderCurrency = "KWD"
	ORDERCURRENCY_KGS OrderCurrency = "KGS"
	ORDERCURRENCY_LAK OrderCurrency = "LAK"
	ORDERCURRENCY_LBP OrderCurrency = "LBP"
	ORDERCURRENCY_LSL OrderCurrency = "LSL"
	ORDERCURRENCY_LRD OrderCurrency = "LRD"
	ORDERCURRENCY_LYD OrderCurrency = "LYD"
	ORDERCURRENCY_LTL OrderCurrency = "LTL"
	ORDERCURRENCY_MOP OrderCurrency = "MOP"
	ORDERCURRENCY_MWK OrderCurrency = "MWK"
	ORDERCURRENCY_MYR OrderCurrency = "MYR"
	ORDERCURRENCY_MVR OrderCurrency = "MVR"
	ORDERCURRENCY_MRO OrderCurrency = "MRO"
	ORDERCURRENCY_MUR OrderCurrency = "MUR"
	ORDERCURRENCY_MXN OrderCurrency = "MXN"
	ORDERCURRENCY_MNT OrderCurrency = "MNT"
	ORDERCURRENCY_MDL OrderCurrency = "MDL"
	ORDERCURRENCY_MAD OrderCurrency = "MAD"
	ORDERCURRENCY_OMR OrderCurrency = "OMR"
	ORDERCURRENCY_NAD OrderCurrency = "NAD"
	ORDERCURRENCY_NPR OrderCurrency = "NPR"
	ORDERCURRENCY_ANG OrderCurrency = "ANG"
	ORDERCURRENCY_AWG OrderCurrency = "AWG"
	ORDERCURRENCY_VUV OrderCurrency = "VUV"
	ORDERCURRENCY_NZD OrderCurrency = "NZD"
	ORDERCURRENCY_NIO OrderCurrency = "NIO"
	ORDERCURRENCY_NGN OrderCurrency = "NGN"
	ORDERCURRENCY_NOK OrderCurrency = "NOK"
	ORDERCURRENCY_PKR OrderCurrency = "PKR"
	ORDERCURRENCY_PAB OrderCurrency = "PAB"
	ORDERCURRENCY_PGK OrderCurrency = "PGK"
	ORDERCURRENCY_PYG OrderCurrency = "PYG"
	ORDERCURRENCY_PEN OrderCurrency = "PEN"
	ORDERCURRENCY_PHP OrderCurrency = "PHP"
	ORDERCURRENCY_QAR OrderCurrency = "QAR"
	ORDERCURRENCY_RUB OrderCurrency = "RUB"
	ORDERCURRENCY_RWF OrderCurrency = "RWF"
	ORDERCURRENCY_SHP OrderCurrency = "SHP"
	ORDERCURRENCY_STD OrderCurrency = "STD"
	ORDERCURRENCY_SAR OrderCurrency = "SAR"
	ORDERCURRENCY_SCR OrderCurrency = "SCR"
	ORDERCURRENCY_SLL OrderCurrency = "SLL"
	ORDERCURRENCY_SGD OrderCurrency = "SGD"
	ORDERCURRENCY_VND OrderCurrency = "VND"
	ORDERCURRENCY_SOS OrderCurrency = "SOS"
	ORDERCURRENCY_ZAR OrderCurrency = "ZAR"
	ORDERCURRENCY_SSP OrderCurrency = "SSP"
	ORDERCURRENCY_SZL OrderCurrency = "SZL"
	ORDERCURRENCY_SEK OrderCurrency = "SEK"
	ORDERCURRENCY_CHF OrderCurrency = "CHF"
	ORDERCURRENCY_SYP OrderCurrency = "SYP"
	ORDERCURRENCY_THB OrderCurrency = "THB"
	ORDERCURRENCY_TOP OrderCurrency = "TOP"
	ORDERCURRENCY_TTD OrderCurrency = "TTD"
	ORDERCURRENCY_AED OrderCurrency = "AED"
	ORDERCURRENCY_TND OrderCurrency = "TND"
	ORDERCURRENCY_UGX OrderCurrency = "UGX"
	ORDERCURRENCY_MKD OrderCurrency = "MKD"
	ORDERCURRENCY_EGP OrderCurrency = "EGP"
	ORDERCURRENCY_GBP OrderCurrency = "GBP"
	ORDERCURRENCY_TZS OrderCurrency = "TZS"
	ORDERCURRENCY_USD OrderCurrency = "USD"
	ORDERCURRENCY_UYU OrderCurrency = "UYU"
	ORDERCURRENCY_UZS OrderCurrency = "UZS"
	ORDERCURRENCY_WST OrderCurrency = "WST"
	ORDERCURRENCY_YER OrderCurrency = "YER"
	ORDERCURRENCY_TWD OrderCurrency = "TWD"
	ORDERCURRENCY_CUC OrderCurrency = "CUC"
	ORDERCURRENCY_ZWL OrderCurrency = "ZWL"
	ORDERCURRENCY_TMT OrderCurrency = "TMT"
	ORDERCURRENCY_GHS OrderCurrency = "GHS"
	ORDERCURRENCY_VEF OrderCurrency = "VEF"
	ORDERCURRENCY_SDG OrderCurrency = "SDG"
	ORDERCURRENCY_UYI OrderCurrency = "UYI"
	ORDERCURRENCY_RSD OrderCurrency = "RSD"
	ORDERCURRENCY_MZN OrderCurrency = "MZN"
	ORDERCURRENCY_AZN OrderCurrency = "AZN"
	ORDERCURRENCY_RON OrderCurrency = "RON"
	ORDERCURRENCY_CHE OrderCurrency = "CHE"
	ORDERCURRENCY_CHW OrderCurrency = "CHW"
	ORDERCURRENCY_TRY OrderCurrency = "TRY"
	ORDERCURRENCY_XAF OrderCurrency = "XAF"
	ORDERCURRENCY_XCD OrderCurrency = "XCD"
	ORDERCURRENCY_XOF OrderCurrency = "XOF"
	ORDERCURRENCY_XPF OrderCurrency = "XPF"
	ORDERCURRENCY_XBA OrderCurrency = "XBA"
	ORDERCURRENCY_XBB OrderCurrency = "XBB"
	ORDERCURRENCY_XBC OrderCurrency = "XBC"
	ORDERCURRENCY_XBD OrderCurrency = "XBD"
	ORDERCURRENCY_XAU OrderCurrency = "XAU"
	ORDERCURRENCY_XDR OrderCurrency = "XDR"
	ORDERCURRENCY_XAG OrderCurrency = "XAG"
	ORDERCURRENCY_XPT OrderCurrency = "XPT"
	ORDERCURRENCY_XPD OrderCurrency = "XPD"
	ORDERCURRENCY_XUA OrderCurrency = "XUA"
	ORDERCURRENCY_ZMW OrderCurrency = "ZMW"
	ORDERCURRENCY_SRD OrderCurrency = "SRD"
	ORDERCURRENCY_MGA OrderCurrency = "MGA"
	ORDERCURRENCY_COU OrderCurrency = "COU"
	ORDERCURRENCY_AFN OrderCurrency = "AFN"
	ORDERCURRENCY_TJS OrderCurrency = "TJS"
	ORDERCURRENCY_AOA OrderCurrency = "AOA"
	ORDERCURRENCY_BYR OrderCurrency = "BYR"
	ORDERCURRENCY_BGN OrderCurrency = "BGN"
	ORDERCURRENCY_CDF OrderCurrency = "CDF"
	ORDERCURRENCY_BAM OrderCurrency = "BAM"
	ORDERCURRENCY_EUR OrderCurrency = "EUR"
	ORDERCURRENCY_MXV OrderCurrency = "MXV"
	ORDERCURRENCY_UAH OrderCurrency = "UAH"
	ORDERCURRENCY_GEL OrderCurrency = "GEL"
	ORDERCURRENCY_BOV OrderCurrency = "BOV"
	ORDERCURRENCY_PLN OrderCurrency = "PLN"
	ORDERCURRENCY_BRL OrderCurrency = "BRL"
	ORDERCURRENCY_CLF OrderCurrency = "CLF"
	ORDERCURRENCY_XSU OrderCurrency = "XSU"
	ORDERCURRENCY_USN OrderCurrency = "USN"
)

// All allowed values of OrderCurrency enum
var AllowedOrderCurrencyEnumValues = []OrderCurrency{
	"XXX",
	"ALL",
	"DZD",
	"ARS",
	"AUD",
	"BSD",
	"BHD",
	"BDT",
	"AMD",
	"BBD",
	"BMD",
	"BTN",
	"BOB",
	"BWP",
	"BZD",
	"SBD",
	"BND",
	"MMK",
	"BIF",
	"KHR",
	"CAD",
	"CVE",
	"KYD",
	"LKR",
	"CLP",
	"CNY",
	"COP",
	"KMF",
	"CRC",
	"HRK",
	"CUP",
	"CZK",
	"DKK",
	"DOP",
	"SVC",
	"ETB",
	"ERN",
	"FKP",
	"FJD",
	"DJF",
	"GMD",
	"GIP",
	"GTQ",
	"GNF",
	"GYD",
	"HTG",
	"HNL",
	"HKD",
	"HUF",
	"ISK",
	"INR",
	"IDR",
	"IRR",
	"IQD",
	"ILS",
	"JMD",
	"JPY",
	"KZT",
	"JOD",
	"KES",
	"KPW",
	"KRW",
	"KWD",
	"KGS",
	"LAK",
	"LBP",
	"LSL",
	"LRD",
	"LYD",
	"LTL",
	"MOP",
	"MWK",
	"MYR",
	"MVR",
	"MRO",
	"MUR",
	"MXN",
	"MNT",
	"MDL",
	"MAD",
	"OMR",
	"NAD",
	"NPR",
	"ANG",
	"AWG",
	"VUV",
	"NZD",
	"NIO",
	"NGN",
	"NOK",
	"PKR",
	"PAB",
	"PGK",
	"PYG",
	"PEN",
	"PHP",
	"QAR",
	"RUB",
	"RWF",
	"SHP",
	"STD",
	"SAR",
	"SCR",
	"SLL",
	"SGD",
	"VND",
	"SOS",
	"ZAR",
	"SSP",
	"SZL",
	"SEK",
	"CHF",
	"SYP",
	"THB",
	"TOP",
	"TTD",
	"AED",
	"TND",
	"UGX",
	"MKD",
	"EGP",
	"GBP",
	"TZS",
	"USD",
	"UYU",
	"UZS",
	"WST",
	"YER",
	"TWD",
	"CUC",
	"ZWL",
	"TMT",
	"GHS",
	"VEF",
	"SDG",
	"UYI",
	"RSD",
	"MZN",
	"AZN",
	"RON",
	"CHE",
	"CHW",
	"TRY",
	"XAF",
	"XCD",
	"XOF",
	"XPF",
	"XBA",
	"XBB",
	"XBC",
	"XBD",
	"XAU",
	"XDR",
	"XAG",
	"XPT",
	"XPD",
	"XUA",
	"ZMW",
	"SRD",
	"MGA",
	"COU",
	"AFN",
	"TJS",
	"AOA",
	"BYR",
	"BGN",
	"CDF",
	"BAM",
	"EUR",
	"MXV",
	"UAH",
	"GEL",
	"BOV",
	"PLN",
	"BRL",
	"CLF",
	"XSU",
	"USN",
}

func (v *OrderCurrency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderCurrency(value)
	for _, existing := range AllowedOrderCurrencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderCurrency", value)
}

// NewOrderCurrencyFromValue returns a pointer to a valid OrderCurrency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderCurrencyFromValue(v string) (*OrderCurrency, error) {
	ev := OrderCurrency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderCurrency: valid values are %v", v, AllowedOrderCurrencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderCurrency) IsValid() bool {
	for _, existing := range AllowedOrderCurrencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderCurrency value
func (v OrderCurrency) Ptr() *OrderCurrency {
	return &v
}

type NullableOrderCurrency struct {
	value *OrderCurrency
	isSet bool
}

func (v NullableOrderCurrency) Get() *OrderCurrency {
	return v.value
}

func (v *NullableOrderCurrency) Set(val *OrderCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCurrency(val *OrderCurrency) *NullableOrderCurrency {
	return &NullableOrderCurrency{value: val, isSet: true}
}

func (v NullableOrderCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


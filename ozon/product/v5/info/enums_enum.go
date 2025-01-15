// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package info

import (
	"fmt"
	"strings"
)

const (
	// PricesRequestFilterVisibilityALL is a PricesRequestFilterVisibility of type ALL.
	PricesRequestFilterVisibilityALL PricesRequestFilterVisibility = iota
	// PricesRequestFilterVisibilityVISIBLE is a PricesRequestFilterVisibility of type VISIBLE.
	PricesRequestFilterVisibilityVISIBLE
	// PricesRequestFilterVisibilityINVISIBLE is a PricesRequestFilterVisibility of type INVISIBLE.
	PricesRequestFilterVisibilityINVISIBLE
	// PricesRequestFilterVisibilityEMPTYSTOCK is a PricesRequestFilterVisibility of type EMPTY_STOCK.
	PricesRequestFilterVisibilityEMPTYSTOCK
	// PricesRequestFilterVisibilityNOTMODERATED is a PricesRequestFilterVisibility of type NOT_MODERATED.
	PricesRequestFilterVisibilityNOTMODERATED
	// PricesRequestFilterVisibilityMODERATED is a PricesRequestFilterVisibility of type MODERATED.
	PricesRequestFilterVisibilityMODERATED
	// PricesRequestFilterVisibilityDISABLED is a PricesRequestFilterVisibility of type DISABLED.
	PricesRequestFilterVisibilityDISABLED
	// PricesRequestFilterVisibilitySTATEFAILED is a PricesRequestFilterVisibility of type STATE_FAILED.
	PricesRequestFilterVisibilitySTATEFAILED
	// PricesRequestFilterVisibilityREADYTOSUPPLY is a PricesRequestFilterVisibility of type READY_TO_SUPPLY.
	PricesRequestFilterVisibilityREADYTOSUPPLY
	// PricesRequestFilterVisibilityVALIDATIONSTATEPENDING is a PricesRequestFilterVisibility of type VALIDATION_STATE_PENDING.
	PricesRequestFilterVisibilityVALIDATIONSTATEPENDING
	// PricesRequestFilterVisibilityVALIDATIONSTATEFAIL is a PricesRequestFilterVisibility of type VALIDATION_STATE_FAIL.
	PricesRequestFilterVisibilityVALIDATIONSTATEFAIL
	// PricesRequestFilterVisibilityVALIDATIONSTATESUCCESS is a PricesRequestFilterVisibility of type VALIDATION_STATE_SUCCESS.
	PricesRequestFilterVisibilityVALIDATIONSTATESUCCESS
	// PricesRequestFilterVisibilityTOSUPPLY is a PricesRequestFilterVisibility of type TO_SUPPLY.
	PricesRequestFilterVisibilityTOSUPPLY
	// PricesRequestFilterVisibilityINSALE is a PricesRequestFilterVisibility of type IN_SALE.
	PricesRequestFilterVisibilityINSALE
	// PricesRequestFilterVisibilityREMOVEDFROMSALE is a PricesRequestFilterVisibility of type REMOVED_FROM_SALE.
	PricesRequestFilterVisibilityREMOVEDFROMSALE
	// PricesRequestFilterVisibilityBANNED is a PricesRequestFilterVisibility of type BANNED.
	PricesRequestFilterVisibilityBANNED
	// PricesRequestFilterVisibilityOVERPRICED is a PricesRequestFilterVisibility of type OVERPRICED.
	PricesRequestFilterVisibilityOVERPRICED
	// PricesRequestFilterVisibilityCRITICALLYOVERPRICED is a PricesRequestFilterVisibility of type CRITICALLY_OVERPRICED.
	PricesRequestFilterVisibilityCRITICALLYOVERPRICED
	// PricesRequestFilterVisibilityEMPTYBARCODE is a PricesRequestFilterVisibility of type EMPTY_BARCODE.
	PricesRequestFilterVisibilityEMPTYBARCODE
	// PricesRequestFilterVisibilityBARCODEEXISTS is a PricesRequestFilterVisibility of type BARCODE_EXISTS.
	PricesRequestFilterVisibilityBARCODEEXISTS
	// PricesRequestFilterVisibilityQUARANTINE is a PricesRequestFilterVisibility of type QUARANTINE.
	PricesRequestFilterVisibilityQUARANTINE
	// PricesRequestFilterVisibilityARCHIVED is a PricesRequestFilterVisibility of type ARCHIVED.
	PricesRequestFilterVisibilityARCHIVED
	// PricesRequestFilterVisibilityOVERPRICEDWITHSTOCK is a PricesRequestFilterVisibility of type OVERPRICED_WITH_STOCK.
	PricesRequestFilterVisibilityOVERPRICEDWITHSTOCK
	// PricesRequestFilterVisibilityPARTIALAPPROVED is a PricesRequestFilterVisibility of type PARTIAL_APPROVED.
	PricesRequestFilterVisibilityPARTIALAPPROVED
)

var ErrInvalidPricesRequestFilterVisibility = fmt.Errorf("not a valid PricesRequestFilterVisibility, try [%s]", strings.Join(_PricesRequestFilterVisibilityNames, ", "))

const _PricesRequestFilterVisibilityName = "ALLVISIBLEINVISIBLEEMPTY_STOCKNOT_MODERATEDMODERATEDDISABLEDSTATE_FAILEDREADY_TO_SUPPLYVALIDATION_STATE_PENDINGVALIDATION_STATE_FAILVALIDATION_STATE_SUCCESSTO_SUPPLYIN_SALEREMOVED_FROM_SALEBANNEDOVERPRICEDCRITICALLY_OVERPRICEDEMPTY_BARCODEBARCODE_EXISTSQUARANTINEARCHIVEDOVERPRICED_WITH_STOCKPARTIAL_APPROVED"

var _PricesRequestFilterVisibilityNames = []string{
	_PricesRequestFilterVisibilityName[0:3],
	_PricesRequestFilterVisibilityName[3:10],
	_PricesRequestFilterVisibilityName[10:19],
	_PricesRequestFilterVisibilityName[19:30],
	_PricesRequestFilterVisibilityName[30:43],
	_PricesRequestFilterVisibilityName[43:52],
	_PricesRequestFilterVisibilityName[52:60],
	_PricesRequestFilterVisibilityName[60:72],
	_PricesRequestFilterVisibilityName[72:87],
	_PricesRequestFilterVisibilityName[87:111],
	_PricesRequestFilterVisibilityName[111:132],
	_PricesRequestFilterVisibilityName[132:156],
	_PricesRequestFilterVisibilityName[156:165],
	_PricesRequestFilterVisibilityName[165:172],
	_PricesRequestFilterVisibilityName[172:189],
	_PricesRequestFilterVisibilityName[189:195],
	_PricesRequestFilterVisibilityName[195:205],
	_PricesRequestFilterVisibilityName[205:226],
	_PricesRequestFilterVisibilityName[226:239],
	_PricesRequestFilterVisibilityName[239:253],
	_PricesRequestFilterVisibilityName[253:263],
	_PricesRequestFilterVisibilityName[263:271],
	_PricesRequestFilterVisibilityName[271:292],
	_PricesRequestFilterVisibilityName[292:308],
}

// PricesRequestFilterVisibilityNames returns a list of possible string values of PricesRequestFilterVisibility.
func PricesRequestFilterVisibilityNames() []string {
	tmp := make([]string, len(_PricesRequestFilterVisibilityNames))
	copy(tmp, _PricesRequestFilterVisibilityNames)
	return tmp
}

var _PricesRequestFilterVisibilityMap = map[PricesRequestFilterVisibility]string{
	PricesRequestFilterVisibilityALL:                    _PricesRequestFilterVisibilityName[0:3],
	PricesRequestFilterVisibilityVISIBLE:                _PricesRequestFilterVisibilityName[3:10],
	PricesRequestFilterVisibilityINVISIBLE:              _PricesRequestFilterVisibilityName[10:19],
	PricesRequestFilterVisibilityEMPTYSTOCK:             _PricesRequestFilterVisibilityName[19:30],
	PricesRequestFilterVisibilityNOTMODERATED:           _PricesRequestFilterVisibilityName[30:43],
	PricesRequestFilterVisibilityMODERATED:              _PricesRequestFilterVisibilityName[43:52],
	PricesRequestFilterVisibilityDISABLED:               _PricesRequestFilterVisibilityName[52:60],
	PricesRequestFilterVisibilitySTATEFAILED:            _PricesRequestFilterVisibilityName[60:72],
	PricesRequestFilterVisibilityREADYTOSUPPLY:          _PricesRequestFilterVisibilityName[72:87],
	PricesRequestFilterVisibilityVALIDATIONSTATEPENDING: _PricesRequestFilterVisibilityName[87:111],
	PricesRequestFilterVisibilityVALIDATIONSTATEFAIL:    _PricesRequestFilterVisibilityName[111:132],
	PricesRequestFilterVisibilityVALIDATIONSTATESUCCESS: _PricesRequestFilterVisibilityName[132:156],
	PricesRequestFilterVisibilityTOSUPPLY:               _PricesRequestFilterVisibilityName[156:165],
	PricesRequestFilterVisibilityINSALE:                 _PricesRequestFilterVisibilityName[165:172],
	PricesRequestFilterVisibilityREMOVEDFROMSALE:        _PricesRequestFilterVisibilityName[172:189],
	PricesRequestFilterVisibilityBANNED:                 _PricesRequestFilterVisibilityName[189:195],
	PricesRequestFilterVisibilityOVERPRICED:             _PricesRequestFilterVisibilityName[195:205],
	PricesRequestFilterVisibilityCRITICALLYOVERPRICED:   _PricesRequestFilterVisibilityName[205:226],
	PricesRequestFilterVisibilityEMPTYBARCODE:           _PricesRequestFilterVisibilityName[226:239],
	PricesRequestFilterVisibilityBARCODEEXISTS:          _PricesRequestFilterVisibilityName[239:253],
	PricesRequestFilterVisibilityQUARANTINE:             _PricesRequestFilterVisibilityName[253:263],
	PricesRequestFilterVisibilityARCHIVED:               _PricesRequestFilterVisibilityName[263:271],
	PricesRequestFilterVisibilityOVERPRICEDWITHSTOCK:    _PricesRequestFilterVisibilityName[271:292],
	PricesRequestFilterVisibilityPARTIALAPPROVED:        _PricesRequestFilterVisibilityName[292:308],
}

// String implements the Stringer interface.
func (x PricesRequestFilterVisibility) String() string {
	if str, ok := _PricesRequestFilterVisibilityMap[x]; ok {
		return str
	}
	return fmt.Sprintf("PricesRequestFilterVisibility(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PricesRequestFilterVisibility) IsValid() bool {
	_, ok := _PricesRequestFilterVisibilityMap[x]
	return ok
}

var _PricesRequestFilterVisibilityValue = map[string]PricesRequestFilterVisibility{
	_PricesRequestFilterVisibilityName[0:3]:     PricesRequestFilterVisibilityALL,
	_PricesRequestFilterVisibilityName[3:10]:    PricesRequestFilterVisibilityVISIBLE,
	_PricesRequestFilterVisibilityName[10:19]:   PricesRequestFilterVisibilityINVISIBLE,
	_PricesRequestFilterVisibilityName[19:30]:   PricesRequestFilterVisibilityEMPTYSTOCK,
	_PricesRequestFilterVisibilityName[30:43]:   PricesRequestFilterVisibilityNOTMODERATED,
	_PricesRequestFilterVisibilityName[43:52]:   PricesRequestFilterVisibilityMODERATED,
	_PricesRequestFilterVisibilityName[52:60]:   PricesRequestFilterVisibilityDISABLED,
	_PricesRequestFilterVisibilityName[60:72]:   PricesRequestFilterVisibilitySTATEFAILED,
	_PricesRequestFilterVisibilityName[72:87]:   PricesRequestFilterVisibilityREADYTOSUPPLY,
	_PricesRequestFilterVisibilityName[87:111]:  PricesRequestFilterVisibilityVALIDATIONSTATEPENDING,
	_PricesRequestFilterVisibilityName[111:132]: PricesRequestFilterVisibilityVALIDATIONSTATEFAIL,
	_PricesRequestFilterVisibilityName[132:156]: PricesRequestFilterVisibilityVALIDATIONSTATESUCCESS,
	_PricesRequestFilterVisibilityName[156:165]: PricesRequestFilterVisibilityTOSUPPLY,
	_PricesRequestFilterVisibilityName[165:172]: PricesRequestFilterVisibilityINSALE,
	_PricesRequestFilterVisibilityName[172:189]: PricesRequestFilterVisibilityREMOVEDFROMSALE,
	_PricesRequestFilterVisibilityName[189:195]: PricesRequestFilterVisibilityBANNED,
	_PricesRequestFilterVisibilityName[195:205]: PricesRequestFilterVisibilityOVERPRICED,
	_PricesRequestFilterVisibilityName[205:226]: PricesRequestFilterVisibilityCRITICALLYOVERPRICED,
	_PricesRequestFilterVisibilityName[226:239]: PricesRequestFilterVisibilityEMPTYBARCODE,
	_PricesRequestFilterVisibilityName[239:253]: PricesRequestFilterVisibilityBARCODEEXISTS,
	_PricesRequestFilterVisibilityName[253:263]: PricesRequestFilterVisibilityQUARANTINE,
	_PricesRequestFilterVisibilityName[263:271]: PricesRequestFilterVisibilityARCHIVED,
	_PricesRequestFilterVisibilityName[271:292]: PricesRequestFilterVisibilityOVERPRICEDWITHSTOCK,
	_PricesRequestFilterVisibilityName[292:308]: PricesRequestFilterVisibilityPARTIALAPPROVED,
}

// ParsePricesRequestFilterVisibility attempts to convert a string to a PricesRequestFilterVisibility.
func ParsePricesRequestFilterVisibility(name string) (PricesRequestFilterVisibility, error) {
	if x, ok := _PricesRequestFilterVisibilityValue[name]; ok {
		return x, nil
	}
	return PricesRequestFilterVisibility(0), fmt.Errorf("%s is %w", name, ErrInvalidPricesRequestFilterVisibility)
}

// MarshalText implements the text marshaller method.
func (x PricesRequestFilterVisibility) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *PricesRequestFilterVisibility) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParsePricesRequestFilterVisibility(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// PricesResponseItemPriceCurrencyCodeRUB is a PricesResponseItemPriceCurrencyCode of type RUB.
	PricesResponseItemPriceCurrencyCodeRUB PricesResponseItemPriceCurrencyCode = iota + 1
	// PricesResponseItemPriceCurrencyCodeBYN is a PricesResponseItemPriceCurrencyCode of type BYN.
	PricesResponseItemPriceCurrencyCodeBYN
	// PricesResponseItemPriceCurrencyCodeKZT is a PricesResponseItemPriceCurrencyCode of type KZT.
	PricesResponseItemPriceCurrencyCodeKZT
	// PricesResponseItemPriceCurrencyCodeEUR is a PricesResponseItemPriceCurrencyCode of type EUR.
	PricesResponseItemPriceCurrencyCodeEUR
	// PricesResponseItemPriceCurrencyCodeUSD is a PricesResponseItemPriceCurrencyCode of type USD.
	PricesResponseItemPriceCurrencyCodeUSD
	// PricesResponseItemPriceCurrencyCodeCNY is a PricesResponseItemPriceCurrencyCode of type CNY.
	PricesResponseItemPriceCurrencyCodeCNY
)

var ErrInvalidPricesResponseItemPriceCurrencyCode = fmt.Errorf("not a valid PricesResponseItemPriceCurrencyCode, try [%s]", strings.Join(_PricesResponseItemPriceCurrencyCodeNames, ", "))

const _PricesResponseItemPriceCurrencyCodeName = "RUBBYNKZTEURUSDCNY"

var _PricesResponseItemPriceCurrencyCodeNames = []string{
	_PricesResponseItemPriceCurrencyCodeName[0:3],
	_PricesResponseItemPriceCurrencyCodeName[3:6],
	_PricesResponseItemPriceCurrencyCodeName[6:9],
	_PricesResponseItemPriceCurrencyCodeName[9:12],
	_PricesResponseItemPriceCurrencyCodeName[12:15],
	_PricesResponseItemPriceCurrencyCodeName[15:18],
}

// PricesResponseItemPriceCurrencyCodeNames returns a list of possible string values of PricesResponseItemPriceCurrencyCode.
func PricesResponseItemPriceCurrencyCodeNames() []string {
	tmp := make([]string, len(_PricesResponseItemPriceCurrencyCodeNames))
	copy(tmp, _PricesResponseItemPriceCurrencyCodeNames)
	return tmp
}

var _PricesResponseItemPriceCurrencyCodeMap = map[PricesResponseItemPriceCurrencyCode]string{
	PricesResponseItemPriceCurrencyCodeRUB: _PricesResponseItemPriceCurrencyCodeName[0:3],
	PricesResponseItemPriceCurrencyCodeBYN: _PricesResponseItemPriceCurrencyCodeName[3:6],
	PricesResponseItemPriceCurrencyCodeKZT: _PricesResponseItemPriceCurrencyCodeName[6:9],
	PricesResponseItemPriceCurrencyCodeEUR: _PricesResponseItemPriceCurrencyCodeName[9:12],
	PricesResponseItemPriceCurrencyCodeUSD: _PricesResponseItemPriceCurrencyCodeName[12:15],
	PricesResponseItemPriceCurrencyCodeCNY: _PricesResponseItemPriceCurrencyCodeName[15:18],
}

// String implements the Stringer interface.
func (x PricesResponseItemPriceCurrencyCode) String() string {
	if str, ok := _PricesResponseItemPriceCurrencyCodeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("PricesResponseItemPriceCurrencyCode(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PricesResponseItemPriceCurrencyCode) IsValid() bool {
	_, ok := _PricesResponseItemPriceCurrencyCodeMap[x]
	return ok
}

var _PricesResponseItemPriceCurrencyCodeValue = map[string]PricesResponseItemPriceCurrencyCode{
	_PricesResponseItemPriceCurrencyCodeName[0:3]:   PricesResponseItemPriceCurrencyCodeRUB,
	_PricesResponseItemPriceCurrencyCodeName[3:6]:   PricesResponseItemPriceCurrencyCodeBYN,
	_PricesResponseItemPriceCurrencyCodeName[6:9]:   PricesResponseItemPriceCurrencyCodeKZT,
	_PricesResponseItemPriceCurrencyCodeName[9:12]:  PricesResponseItemPriceCurrencyCodeEUR,
	_PricesResponseItemPriceCurrencyCodeName[12:15]: PricesResponseItemPriceCurrencyCodeUSD,
	_PricesResponseItemPriceCurrencyCodeName[15:18]: PricesResponseItemPriceCurrencyCodeCNY,
}

// ParsePricesResponseItemPriceCurrencyCode attempts to convert a string to a PricesResponseItemPriceCurrencyCode.
func ParsePricesResponseItemPriceCurrencyCode(name string) (PricesResponseItemPriceCurrencyCode, error) {
	if x, ok := _PricesResponseItemPriceCurrencyCodeValue[name]; ok {
		return x, nil
	}
	return PricesResponseItemPriceCurrencyCode(0), fmt.Errorf("%s is %w", name, ErrInvalidPricesResponseItemPriceCurrencyCode)
}

// MarshalText implements the text marshaller method.
func (x PricesResponseItemPriceCurrencyCode) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *PricesResponseItemPriceCurrencyCode) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParsePricesResponseItemPriceCurrencyCode(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// PricesResponseItemPriceIndexesColorIndexWITHOUTINDEX is a PricesResponseItemPriceIndexesColorIndex of type WITHOUT_INDEX.
	PricesResponseItemPriceIndexesColorIndexWITHOUTINDEX PricesResponseItemPriceIndexesColorIndex = iota
	// PricesResponseItemPriceIndexesColorIndexGREEN is a PricesResponseItemPriceIndexesColorIndex of type GREEN.
	PricesResponseItemPriceIndexesColorIndexGREEN
	// PricesResponseItemPriceIndexesColorIndexYELLOW is a PricesResponseItemPriceIndexesColorIndex of type YELLOW.
	PricesResponseItemPriceIndexesColorIndexYELLOW
	// PricesResponseItemPriceIndexesColorIndexRED is a PricesResponseItemPriceIndexesColorIndex of type RED.
	PricesResponseItemPriceIndexesColorIndexRED
)

var ErrInvalidPricesResponseItemPriceIndexesColorIndex = fmt.Errorf("not a valid PricesResponseItemPriceIndexesColorIndex, try [%s]", strings.Join(_PricesResponseItemPriceIndexesColorIndexNames, ", "))

const _PricesResponseItemPriceIndexesColorIndexName = "WITHOUT_INDEXGREENYELLOWRED"

var _PricesResponseItemPriceIndexesColorIndexNames = []string{
	_PricesResponseItemPriceIndexesColorIndexName[0:13],
	_PricesResponseItemPriceIndexesColorIndexName[13:18],
	_PricesResponseItemPriceIndexesColorIndexName[18:24],
	_PricesResponseItemPriceIndexesColorIndexName[24:27],
}

// PricesResponseItemPriceIndexesColorIndexNames returns a list of possible string values of PricesResponseItemPriceIndexesColorIndex.
func PricesResponseItemPriceIndexesColorIndexNames() []string {
	tmp := make([]string, len(_PricesResponseItemPriceIndexesColorIndexNames))
	copy(tmp, _PricesResponseItemPriceIndexesColorIndexNames)
	return tmp
}

var _PricesResponseItemPriceIndexesColorIndexMap = map[PricesResponseItemPriceIndexesColorIndex]string{
	PricesResponseItemPriceIndexesColorIndexWITHOUTINDEX: _PricesResponseItemPriceIndexesColorIndexName[0:13],
	PricesResponseItemPriceIndexesColorIndexGREEN:        _PricesResponseItemPriceIndexesColorIndexName[13:18],
	PricesResponseItemPriceIndexesColorIndexYELLOW:       _PricesResponseItemPriceIndexesColorIndexName[18:24],
	PricesResponseItemPriceIndexesColorIndexRED:          _PricesResponseItemPriceIndexesColorIndexName[24:27],
}

// String implements the Stringer interface.
func (x PricesResponseItemPriceIndexesColorIndex) String() string {
	if str, ok := _PricesResponseItemPriceIndexesColorIndexMap[x]; ok {
		return str
	}
	return fmt.Sprintf("PricesResponseItemPriceIndexesColorIndex(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PricesResponseItemPriceIndexesColorIndex) IsValid() bool {
	_, ok := _PricesResponseItemPriceIndexesColorIndexMap[x]
	return ok
}

var _PricesResponseItemPriceIndexesColorIndexValue = map[string]PricesResponseItemPriceIndexesColorIndex{
	_PricesResponseItemPriceIndexesColorIndexName[0:13]:  PricesResponseItemPriceIndexesColorIndexWITHOUTINDEX,
	_PricesResponseItemPriceIndexesColorIndexName[13:18]: PricesResponseItemPriceIndexesColorIndexGREEN,
	_PricesResponseItemPriceIndexesColorIndexName[18:24]: PricesResponseItemPriceIndexesColorIndexYELLOW,
	_PricesResponseItemPriceIndexesColorIndexName[24:27]: PricesResponseItemPriceIndexesColorIndexRED,
}

// ParsePricesResponseItemPriceIndexesColorIndex attempts to convert a string to a PricesResponseItemPriceIndexesColorIndex.
func ParsePricesResponseItemPriceIndexesColorIndex(name string) (PricesResponseItemPriceIndexesColorIndex, error) {
	if x, ok := _PricesResponseItemPriceIndexesColorIndexValue[name]; ok {
		return x, nil
	}
	return PricesResponseItemPriceIndexesColorIndex(0), fmt.Errorf("%s is %w", name, ErrInvalidPricesResponseItemPriceIndexesColorIndex)
}

// MarshalText implements the text marshaller method.
func (x PricesResponseItemPriceIndexesColorIndex) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *PricesResponseItemPriceIndexesColorIndex) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParsePricesResponseItemPriceIndexesColorIndex(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

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
	// PricesRequestFilterVisibilityIMAGEABSENT is a PricesRequestFilterVisibility of type IMAGE_ABSENT.
	PricesRequestFilterVisibilityIMAGEABSENT
	// PricesRequestFilterVisibilityMODERATIONBLOCK is a PricesRequestFilterVisibility of type MODERATION_BLOCK.
	PricesRequestFilterVisibilityMODERATIONBLOCK
)

var ErrInvalidPricesRequestFilterVisibility = fmt.Errorf("not a valid PricesRequestFilterVisibility, try [%s]", strings.Join(_PricesRequestFilterVisibilityNames, ", "))

const _PricesRequestFilterVisibilityName = "ALLVISIBLEINVISIBLEEMPTY_STOCKNOT_MODERATEDMODERATEDDISABLEDSTATE_FAILEDREADY_TO_SUPPLYVALIDATION_STATE_PENDINGVALIDATION_STATE_FAILVALIDATION_STATE_SUCCESSTO_SUPPLYIN_SALEREMOVED_FROM_SALEBANNEDOVERPRICEDCRITICALLY_OVERPRICEDEMPTY_BARCODEBARCODE_EXISTSQUARANTINEARCHIVEDOVERPRICED_WITH_STOCKPARTIAL_APPROVEDIMAGE_ABSENTMODERATION_BLOCK"

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
	_PricesRequestFilterVisibilityName[308:320],
	_PricesRequestFilterVisibilityName[320:336],
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
	PricesRequestFilterVisibilityIMAGEABSENT:            _PricesRequestFilterVisibilityName[308:320],
	PricesRequestFilterVisibilityMODERATIONBLOCK:        _PricesRequestFilterVisibilityName[320:336],
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
	_PricesRequestFilterVisibilityName[308:320]: PricesRequestFilterVisibilityIMAGEABSENT,
	_PricesRequestFilterVisibilityName[320:336]: PricesRequestFilterVisibilityMODERATIONBLOCK,
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
	// PricesResponseResultItemPriceCurrencyCodeRUB is a PricesResponseResultItemPriceCurrencyCode of type RUB.
	PricesResponseResultItemPriceCurrencyCodeRUB PricesResponseResultItemPriceCurrencyCode = iota + 1
	// PricesResponseResultItemPriceCurrencyCodeBYN is a PricesResponseResultItemPriceCurrencyCode of type BYN.
	PricesResponseResultItemPriceCurrencyCodeBYN
	// PricesResponseResultItemPriceCurrencyCodeKZT is a PricesResponseResultItemPriceCurrencyCode of type KZT.
	PricesResponseResultItemPriceCurrencyCodeKZT
	// PricesResponseResultItemPriceCurrencyCodeEUR is a PricesResponseResultItemPriceCurrencyCode of type EUR.
	PricesResponseResultItemPriceCurrencyCodeEUR
	// PricesResponseResultItemPriceCurrencyCodeUSD is a PricesResponseResultItemPriceCurrencyCode of type USD.
	PricesResponseResultItemPriceCurrencyCodeUSD
	// PricesResponseResultItemPriceCurrencyCodeCNY is a PricesResponseResultItemPriceCurrencyCode of type CNY.
	PricesResponseResultItemPriceCurrencyCodeCNY
)

var ErrInvalidPricesResponseResultItemPriceCurrencyCode = fmt.Errorf("not a valid PricesResponseResultItemPriceCurrencyCode, try [%s]", strings.Join(_PricesResponseResultItemPriceCurrencyCodeNames, ", "))

const _PricesResponseResultItemPriceCurrencyCodeName = "RUBBYNKZTEURUSDCNY"

var _PricesResponseResultItemPriceCurrencyCodeNames = []string{
	_PricesResponseResultItemPriceCurrencyCodeName[0:3],
	_PricesResponseResultItemPriceCurrencyCodeName[3:6],
	_PricesResponseResultItemPriceCurrencyCodeName[6:9],
	_PricesResponseResultItemPriceCurrencyCodeName[9:12],
	_PricesResponseResultItemPriceCurrencyCodeName[12:15],
	_PricesResponseResultItemPriceCurrencyCodeName[15:18],
}

// PricesResponseResultItemPriceCurrencyCodeNames returns a list of possible string values of PricesResponseResultItemPriceCurrencyCode.
func PricesResponseResultItemPriceCurrencyCodeNames() []string {
	tmp := make([]string, len(_PricesResponseResultItemPriceCurrencyCodeNames))
	copy(tmp, _PricesResponseResultItemPriceCurrencyCodeNames)
	return tmp
}

var _PricesResponseResultItemPriceCurrencyCodeMap = map[PricesResponseResultItemPriceCurrencyCode]string{
	PricesResponseResultItemPriceCurrencyCodeRUB: _PricesResponseResultItemPriceCurrencyCodeName[0:3],
	PricesResponseResultItemPriceCurrencyCodeBYN: _PricesResponseResultItemPriceCurrencyCodeName[3:6],
	PricesResponseResultItemPriceCurrencyCodeKZT: _PricesResponseResultItemPriceCurrencyCodeName[6:9],
	PricesResponseResultItemPriceCurrencyCodeEUR: _PricesResponseResultItemPriceCurrencyCodeName[9:12],
	PricesResponseResultItemPriceCurrencyCodeUSD: _PricesResponseResultItemPriceCurrencyCodeName[12:15],
	PricesResponseResultItemPriceCurrencyCodeCNY: _PricesResponseResultItemPriceCurrencyCodeName[15:18],
}

// String implements the Stringer interface.
func (x PricesResponseResultItemPriceCurrencyCode) String() string {
	if str, ok := _PricesResponseResultItemPriceCurrencyCodeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("PricesResponseResultItemPriceCurrencyCode(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PricesResponseResultItemPriceCurrencyCode) IsValid() bool {
	_, ok := _PricesResponseResultItemPriceCurrencyCodeMap[x]
	return ok
}

var _PricesResponseResultItemPriceCurrencyCodeValue = map[string]PricesResponseResultItemPriceCurrencyCode{
	_PricesResponseResultItemPriceCurrencyCodeName[0:3]:   PricesResponseResultItemPriceCurrencyCodeRUB,
	_PricesResponseResultItemPriceCurrencyCodeName[3:6]:   PricesResponseResultItemPriceCurrencyCodeBYN,
	_PricesResponseResultItemPriceCurrencyCodeName[6:9]:   PricesResponseResultItemPriceCurrencyCodeKZT,
	_PricesResponseResultItemPriceCurrencyCodeName[9:12]:  PricesResponseResultItemPriceCurrencyCodeEUR,
	_PricesResponseResultItemPriceCurrencyCodeName[12:15]: PricesResponseResultItemPriceCurrencyCodeUSD,
	_PricesResponseResultItemPriceCurrencyCodeName[15:18]: PricesResponseResultItemPriceCurrencyCodeCNY,
}

// ParsePricesResponseResultItemPriceCurrencyCode attempts to convert a string to a PricesResponseResultItemPriceCurrencyCode.
func ParsePricesResponseResultItemPriceCurrencyCode(name string) (PricesResponseResultItemPriceCurrencyCode, error) {
	if x, ok := _PricesResponseResultItemPriceCurrencyCodeValue[name]; ok {
		return x, nil
	}
	return PricesResponseResultItemPriceCurrencyCode(0), fmt.Errorf("%s is %w", name, ErrInvalidPricesResponseResultItemPriceCurrencyCode)
}

// MarshalText implements the text marshaller method.
func (x PricesResponseResultItemPriceCurrencyCode) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *PricesResponseResultItemPriceCurrencyCode) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParsePricesResponseResultItemPriceCurrencyCode(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// PricesResponseResultItemPriceIndexesPriceIndexWITHOUTINDEX is a PricesResponseResultItemPriceIndexesPriceIndex of type WITHOUT_INDEX.
	PricesResponseResultItemPriceIndexesPriceIndexWITHOUTINDEX PricesResponseResultItemPriceIndexesPriceIndex = iota + 1
	// PricesResponseResultItemPriceIndexesPriceIndexPROFIT is a PricesResponseResultItemPriceIndexesPriceIndex of type PROFIT.
	PricesResponseResultItemPriceIndexesPriceIndexPROFIT
	// PricesResponseResultItemPriceIndexesPriceIndexAVGPROFIT is a PricesResponseResultItemPriceIndexesPriceIndex of type AVG_PROFIT.
	PricesResponseResultItemPriceIndexesPriceIndexAVGPROFIT
	// PricesResponseResultItemPriceIndexesPriceIndexNONPROFIT is a PricesResponseResultItemPriceIndexesPriceIndex of type NON_PROFIT.
	PricesResponseResultItemPriceIndexesPriceIndexNONPROFIT
)

var ErrInvalidPricesResponseResultItemPriceIndexesPriceIndex = fmt.Errorf("not a valid PricesResponseResultItemPriceIndexesPriceIndex, try [%s]", strings.Join(_PricesResponseResultItemPriceIndexesPriceIndexNames, ", "))

const _PricesResponseResultItemPriceIndexesPriceIndexName = "WITHOUT_INDEXPROFITAVG_PROFITNON_PROFIT"

var _PricesResponseResultItemPriceIndexesPriceIndexNames = []string{
	_PricesResponseResultItemPriceIndexesPriceIndexName[0:13],
	_PricesResponseResultItemPriceIndexesPriceIndexName[13:19],
	_PricesResponseResultItemPriceIndexesPriceIndexName[19:29],
	_PricesResponseResultItemPriceIndexesPriceIndexName[29:39],
}

// PricesResponseResultItemPriceIndexesPriceIndexNames returns a list of possible string values of PricesResponseResultItemPriceIndexesPriceIndex.
func PricesResponseResultItemPriceIndexesPriceIndexNames() []string {
	tmp := make([]string, len(_PricesResponseResultItemPriceIndexesPriceIndexNames))
	copy(tmp, _PricesResponseResultItemPriceIndexesPriceIndexNames)
	return tmp
}

var _PricesResponseResultItemPriceIndexesPriceIndexMap = map[PricesResponseResultItemPriceIndexesPriceIndex]string{
	PricesResponseResultItemPriceIndexesPriceIndexWITHOUTINDEX: _PricesResponseResultItemPriceIndexesPriceIndexName[0:13],
	PricesResponseResultItemPriceIndexesPriceIndexPROFIT:       _PricesResponseResultItemPriceIndexesPriceIndexName[13:19],
	PricesResponseResultItemPriceIndexesPriceIndexAVGPROFIT:    _PricesResponseResultItemPriceIndexesPriceIndexName[19:29],
	PricesResponseResultItemPriceIndexesPriceIndexNONPROFIT:    _PricesResponseResultItemPriceIndexesPriceIndexName[29:39],
}

// String implements the Stringer interface.
func (x PricesResponseResultItemPriceIndexesPriceIndex) String() string {
	if str, ok := _PricesResponseResultItemPriceIndexesPriceIndexMap[x]; ok {
		return str
	}
	return fmt.Sprintf("PricesResponseResultItemPriceIndexesPriceIndex(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PricesResponseResultItemPriceIndexesPriceIndex) IsValid() bool {
	_, ok := _PricesResponseResultItemPriceIndexesPriceIndexMap[x]
	return ok
}

var _PricesResponseResultItemPriceIndexesPriceIndexValue = map[string]PricesResponseResultItemPriceIndexesPriceIndex{
	_PricesResponseResultItemPriceIndexesPriceIndexName[0:13]:  PricesResponseResultItemPriceIndexesPriceIndexWITHOUTINDEX,
	_PricesResponseResultItemPriceIndexesPriceIndexName[13:19]: PricesResponseResultItemPriceIndexesPriceIndexPROFIT,
	_PricesResponseResultItemPriceIndexesPriceIndexName[19:29]: PricesResponseResultItemPriceIndexesPriceIndexAVGPROFIT,
	_PricesResponseResultItemPriceIndexesPriceIndexName[29:39]: PricesResponseResultItemPriceIndexesPriceIndexNONPROFIT,
}

// ParsePricesResponseResultItemPriceIndexesPriceIndex attempts to convert a string to a PricesResponseResultItemPriceIndexesPriceIndex.
func ParsePricesResponseResultItemPriceIndexesPriceIndex(name string) (PricesResponseResultItemPriceIndexesPriceIndex, error) {
	if x, ok := _PricesResponseResultItemPriceIndexesPriceIndexValue[name]; ok {
		return x, nil
	}
	return PricesResponseResultItemPriceIndexesPriceIndex(0), fmt.Errorf("%s is %w", name, ErrInvalidPricesResponseResultItemPriceIndexesPriceIndex)
}

// MarshalText implements the text marshaller method.
func (x PricesResponseResultItemPriceIndexesPriceIndex) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *PricesResponseResultItemPriceIndexesPriceIndex) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParsePricesResponseResultItemPriceIndexesPriceIndex(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// StocksRequestFilterVisibilityALL is a StocksRequestFilterVisibility of type ALL.
	StocksRequestFilterVisibilityALL StocksRequestFilterVisibility = iota
	// StocksRequestFilterVisibilityVISIBLE is a StocksRequestFilterVisibility of type VISIBLE.
	StocksRequestFilterVisibilityVISIBLE
	// StocksRequestFilterVisibilityINVISIBLE is a StocksRequestFilterVisibility of type INVISIBLE.
	StocksRequestFilterVisibilityINVISIBLE
	// StocksRequestFilterVisibilityEMPTYSTOCK is a StocksRequestFilterVisibility of type EMPTY_STOCK.
	StocksRequestFilterVisibilityEMPTYSTOCK
	// StocksRequestFilterVisibilityNOTMODERATED is a StocksRequestFilterVisibility of type NOT_MODERATED.
	StocksRequestFilterVisibilityNOTMODERATED
	// StocksRequestFilterVisibilityMODERATED is a StocksRequestFilterVisibility of type MODERATED.
	StocksRequestFilterVisibilityMODERATED
	// StocksRequestFilterVisibilityDISABLED is a StocksRequestFilterVisibility of type DISABLED.
	StocksRequestFilterVisibilityDISABLED
	// StocksRequestFilterVisibilitySTATEFAILED is a StocksRequestFilterVisibility of type STATE_FAILED.
	StocksRequestFilterVisibilitySTATEFAILED
	// StocksRequestFilterVisibilityREADYTOSUPPLY is a StocksRequestFilterVisibility of type READY_TO_SUPPLY.
	StocksRequestFilterVisibilityREADYTOSUPPLY
	// StocksRequestFilterVisibilityVALIDATIONSTATEPENDING is a StocksRequestFilterVisibility of type VALIDATION_STATE_PENDING.
	StocksRequestFilterVisibilityVALIDATIONSTATEPENDING
	// StocksRequestFilterVisibilityVALIDATIONSTATEFAIL is a StocksRequestFilterVisibility of type VALIDATION_STATE_FAIL.
	StocksRequestFilterVisibilityVALIDATIONSTATEFAIL
	// StocksRequestFilterVisibilityVALIDATIONSTATESUCCESS is a StocksRequestFilterVisibility of type VALIDATION_STATE_SUCCESS.
	StocksRequestFilterVisibilityVALIDATIONSTATESUCCESS
	// StocksRequestFilterVisibilityTOSUPPLY is a StocksRequestFilterVisibility of type TO_SUPPLY.
	StocksRequestFilterVisibilityTOSUPPLY
	// StocksRequestFilterVisibilityINSALE is a StocksRequestFilterVisibility of type IN_SALE.
	StocksRequestFilterVisibilityINSALE
	// StocksRequestFilterVisibilityREMOVEDFROMSALE is a StocksRequestFilterVisibility of type REMOVED_FROM_SALE.
	StocksRequestFilterVisibilityREMOVEDFROMSALE
	// StocksRequestFilterVisibilityBANNED is a StocksRequestFilterVisibility of type BANNED.
	StocksRequestFilterVisibilityBANNED
	// StocksRequestFilterVisibilityOVERPRICED is a StocksRequestFilterVisibility of type OVERPRICED.
	StocksRequestFilterVisibilityOVERPRICED
	// StocksRequestFilterVisibilityCRITICALLYOVERPRICED is a StocksRequestFilterVisibility of type CRITICALLY_OVERPRICED.
	StocksRequestFilterVisibilityCRITICALLYOVERPRICED
	// StocksRequestFilterVisibilityEMPTYBARCODE is a StocksRequestFilterVisibility of type EMPTY_BARCODE.
	StocksRequestFilterVisibilityEMPTYBARCODE
	// StocksRequestFilterVisibilityBARCODEEXISTS is a StocksRequestFilterVisibility of type BARCODE_EXISTS.
	StocksRequestFilterVisibilityBARCODEEXISTS
	// StocksRequestFilterVisibilityQUARANTINE is a StocksRequestFilterVisibility of type QUARANTINE.
	StocksRequestFilterVisibilityQUARANTINE
	// StocksRequestFilterVisibilityARCHIVED is a StocksRequestFilterVisibility of type ARCHIVED.
	StocksRequestFilterVisibilityARCHIVED
	// StocksRequestFilterVisibilityOVERPRICEDWITHSTOCK is a StocksRequestFilterVisibility of type OVERPRICED_WITH_STOCK.
	StocksRequestFilterVisibilityOVERPRICEDWITHSTOCK
	// StocksRequestFilterVisibilityPARTIALAPPROVED is a StocksRequestFilterVisibility of type PARTIAL_APPROVED.
	StocksRequestFilterVisibilityPARTIALAPPROVED
)

var ErrInvalidStocksRequestFilterVisibility = fmt.Errorf("not a valid StocksRequestFilterVisibility, try [%s]", strings.Join(_StocksRequestFilterVisibilityNames, ", "))

const _StocksRequestFilterVisibilityName = "ALLVISIBLEINVISIBLEEMPTY_STOCKNOT_MODERATEDMODERATEDDISABLEDSTATE_FAILEDREADY_TO_SUPPLYVALIDATION_STATE_PENDINGVALIDATION_STATE_FAILVALIDATION_STATE_SUCCESSTO_SUPPLYIN_SALEREMOVED_FROM_SALEBANNEDOVERPRICEDCRITICALLY_OVERPRICEDEMPTY_BARCODEBARCODE_EXISTSQUARANTINEARCHIVEDOVERPRICED_WITH_STOCKPARTIAL_APPROVED"

var _StocksRequestFilterVisibilityNames = []string{
	_StocksRequestFilterVisibilityName[0:3],
	_StocksRequestFilterVisibilityName[3:10],
	_StocksRequestFilterVisibilityName[10:19],
	_StocksRequestFilterVisibilityName[19:30],
	_StocksRequestFilterVisibilityName[30:43],
	_StocksRequestFilterVisibilityName[43:52],
	_StocksRequestFilterVisibilityName[52:60],
	_StocksRequestFilterVisibilityName[60:72],
	_StocksRequestFilterVisibilityName[72:87],
	_StocksRequestFilterVisibilityName[87:111],
	_StocksRequestFilterVisibilityName[111:132],
	_StocksRequestFilterVisibilityName[132:156],
	_StocksRequestFilterVisibilityName[156:165],
	_StocksRequestFilterVisibilityName[165:172],
	_StocksRequestFilterVisibilityName[172:189],
	_StocksRequestFilterVisibilityName[189:195],
	_StocksRequestFilterVisibilityName[195:205],
	_StocksRequestFilterVisibilityName[205:226],
	_StocksRequestFilterVisibilityName[226:239],
	_StocksRequestFilterVisibilityName[239:253],
	_StocksRequestFilterVisibilityName[253:263],
	_StocksRequestFilterVisibilityName[263:271],
	_StocksRequestFilterVisibilityName[271:292],
	_StocksRequestFilterVisibilityName[292:308],
}

// StocksRequestFilterVisibilityNames returns a list of possible string values of StocksRequestFilterVisibility.
func StocksRequestFilterVisibilityNames() []string {
	tmp := make([]string, len(_StocksRequestFilterVisibilityNames))
	copy(tmp, _StocksRequestFilterVisibilityNames)
	return tmp
}

var _StocksRequestFilterVisibilityMap = map[StocksRequestFilterVisibility]string{
	StocksRequestFilterVisibilityALL:                    _StocksRequestFilterVisibilityName[0:3],
	StocksRequestFilterVisibilityVISIBLE:                _StocksRequestFilterVisibilityName[3:10],
	StocksRequestFilterVisibilityINVISIBLE:              _StocksRequestFilterVisibilityName[10:19],
	StocksRequestFilterVisibilityEMPTYSTOCK:             _StocksRequestFilterVisibilityName[19:30],
	StocksRequestFilterVisibilityNOTMODERATED:           _StocksRequestFilterVisibilityName[30:43],
	StocksRequestFilterVisibilityMODERATED:              _StocksRequestFilterVisibilityName[43:52],
	StocksRequestFilterVisibilityDISABLED:               _StocksRequestFilterVisibilityName[52:60],
	StocksRequestFilterVisibilitySTATEFAILED:            _StocksRequestFilterVisibilityName[60:72],
	StocksRequestFilterVisibilityREADYTOSUPPLY:          _StocksRequestFilterVisibilityName[72:87],
	StocksRequestFilterVisibilityVALIDATIONSTATEPENDING: _StocksRequestFilterVisibilityName[87:111],
	StocksRequestFilterVisibilityVALIDATIONSTATEFAIL:    _StocksRequestFilterVisibilityName[111:132],
	StocksRequestFilterVisibilityVALIDATIONSTATESUCCESS: _StocksRequestFilterVisibilityName[132:156],
	StocksRequestFilterVisibilityTOSUPPLY:               _StocksRequestFilterVisibilityName[156:165],
	StocksRequestFilterVisibilityINSALE:                 _StocksRequestFilterVisibilityName[165:172],
	StocksRequestFilterVisibilityREMOVEDFROMSALE:        _StocksRequestFilterVisibilityName[172:189],
	StocksRequestFilterVisibilityBANNED:                 _StocksRequestFilterVisibilityName[189:195],
	StocksRequestFilterVisibilityOVERPRICED:             _StocksRequestFilterVisibilityName[195:205],
	StocksRequestFilterVisibilityCRITICALLYOVERPRICED:   _StocksRequestFilterVisibilityName[205:226],
	StocksRequestFilterVisibilityEMPTYBARCODE:           _StocksRequestFilterVisibilityName[226:239],
	StocksRequestFilterVisibilityBARCODEEXISTS:          _StocksRequestFilterVisibilityName[239:253],
	StocksRequestFilterVisibilityQUARANTINE:             _StocksRequestFilterVisibilityName[253:263],
	StocksRequestFilterVisibilityARCHIVED:               _StocksRequestFilterVisibilityName[263:271],
	StocksRequestFilterVisibilityOVERPRICEDWITHSTOCK:    _StocksRequestFilterVisibilityName[271:292],
	StocksRequestFilterVisibilityPARTIALAPPROVED:        _StocksRequestFilterVisibilityName[292:308],
}

// String implements the Stringer interface.
func (x StocksRequestFilterVisibility) String() string {
	if str, ok := _StocksRequestFilterVisibilityMap[x]; ok {
		return str
	}
	return fmt.Sprintf("StocksRequestFilterVisibility(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x StocksRequestFilterVisibility) IsValid() bool {
	_, ok := _StocksRequestFilterVisibilityMap[x]
	return ok
}

var _StocksRequestFilterVisibilityValue = map[string]StocksRequestFilterVisibility{
	_StocksRequestFilterVisibilityName[0:3]:     StocksRequestFilterVisibilityALL,
	_StocksRequestFilterVisibilityName[3:10]:    StocksRequestFilterVisibilityVISIBLE,
	_StocksRequestFilterVisibilityName[10:19]:   StocksRequestFilterVisibilityINVISIBLE,
	_StocksRequestFilterVisibilityName[19:30]:   StocksRequestFilterVisibilityEMPTYSTOCK,
	_StocksRequestFilterVisibilityName[30:43]:   StocksRequestFilterVisibilityNOTMODERATED,
	_StocksRequestFilterVisibilityName[43:52]:   StocksRequestFilterVisibilityMODERATED,
	_StocksRequestFilterVisibilityName[52:60]:   StocksRequestFilterVisibilityDISABLED,
	_StocksRequestFilterVisibilityName[60:72]:   StocksRequestFilterVisibilitySTATEFAILED,
	_StocksRequestFilterVisibilityName[72:87]:   StocksRequestFilterVisibilityREADYTOSUPPLY,
	_StocksRequestFilterVisibilityName[87:111]:  StocksRequestFilterVisibilityVALIDATIONSTATEPENDING,
	_StocksRequestFilterVisibilityName[111:132]: StocksRequestFilterVisibilityVALIDATIONSTATEFAIL,
	_StocksRequestFilterVisibilityName[132:156]: StocksRequestFilterVisibilityVALIDATIONSTATESUCCESS,
	_StocksRequestFilterVisibilityName[156:165]: StocksRequestFilterVisibilityTOSUPPLY,
	_StocksRequestFilterVisibilityName[165:172]: StocksRequestFilterVisibilityINSALE,
	_StocksRequestFilterVisibilityName[172:189]: StocksRequestFilterVisibilityREMOVEDFROMSALE,
	_StocksRequestFilterVisibilityName[189:195]: StocksRequestFilterVisibilityBANNED,
	_StocksRequestFilterVisibilityName[195:205]: StocksRequestFilterVisibilityOVERPRICED,
	_StocksRequestFilterVisibilityName[205:226]: StocksRequestFilterVisibilityCRITICALLYOVERPRICED,
	_StocksRequestFilterVisibilityName[226:239]: StocksRequestFilterVisibilityEMPTYBARCODE,
	_StocksRequestFilterVisibilityName[239:253]: StocksRequestFilterVisibilityBARCODEEXISTS,
	_StocksRequestFilterVisibilityName[253:263]: StocksRequestFilterVisibilityQUARANTINE,
	_StocksRequestFilterVisibilityName[263:271]: StocksRequestFilterVisibilityARCHIVED,
	_StocksRequestFilterVisibilityName[271:292]: StocksRequestFilterVisibilityOVERPRICEDWITHSTOCK,
	_StocksRequestFilterVisibilityName[292:308]: StocksRequestFilterVisibilityPARTIALAPPROVED,
}

// ParseStocksRequestFilterVisibility attempts to convert a string to a StocksRequestFilterVisibility.
func ParseStocksRequestFilterVisibility(name string) (StocksRequestFilterVisibility, error) {
	if x, ok := _StocksRequestFilterVisibilityValue[name]; ok {
		return x, nil
	}
	return StocksRequestFilterVisibility(0), fmt.Errorf("%s is %w", name, ErrInvalidStocksRequestFilterVisibility)
}

// MarshalText implements the text marshaller method.
func (x StocksRequestFilterVisibility) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *StocksRequestFilterVisibility) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseStocksRequestFilterVisibility(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// StocksResponseItemStockShipmentTypeSHIPMENTTYPEGENERAL is a StocksResponseItemStockShipmentType of type SHIPMENT_TYPE_GENERAL.
	StocksResponseItemStockShipmentTypeSHIPMENTTYPEGENERAL StocksResponseItemStockShipmentType = iota
	// StocksResponseItemStockShipmentTypeSHIPMENTTYPEBOX is a StocksResponseItemStockShipmentType of type SHIPMENT_TYPE_BOX.
	StocksResponseItemStockShipmentTypeSHIPMENTTYPEBOX
	// StocksResponseItemStockShipmentTypeSHIPMENTTYPEPALLET is a StocksResponseItemStockShipmentType of type SHIPMENT_TYPE_PALLET.
	StocksResponseItemStockShipmentTypeSHIPMENTTYPEPALLET
)

var ErrInvalidStocksResponseItemStockShipmentType = fmt.Errorf("not a valid StocksResponseItemStockShipmentType, try [%s]", strings.Join(_StocksResponseItemStockShipmentTypeNames, ", "))

const _StocksResponseItemStockShipmentTypeName = "SHIPMENT_TYPE_GENERALSHIPMENT_TYPE_BOXSHIPMENT_TYPE_PALLET"

var _StocksResponseItemStockShipmentTypeNames = []string{
	_StocksResponseItemStockShipmentTypeName[0:21],
	_StocksResponseItemStockShipmentTypeName[21:38],
	_StocksResponseItemStockShipmentTypeName[38:58],
}

// StocksResponseItemStockShipmentTypeNames returns a list of possible string values of StocksResponseItemStockShipmentType.
func StocksResponseItemStockShipmentTypeNames() []string {
	tmp := make([]string, len(_StocksResponseItemStockShipmentTypeNames))
	copy(tmp, _StocksResponseItemStockShipmentTypeNames)
	return tmp
}

var _StocksResponseItemStockShipmentTypeMap = map[StocksResponseItemStockShipmentType]string{
	StocksResponseItemStockShipmentTypeSHIPMENTTYPEGENERAL: _StocksResponseItemStockShipmentTypeName[0:21],
	StocksResponseItemStockShipmentTypeSHIPMENTTYPEBOX:     _StocksResponseItemStockShipmentTypeName[21:38],
	StocksResponseItemStockShipmentTypeSHIPMENTTYPEPALLET:  _StocksResponseItemStockShipmentTypeName[38:58],
}

// String implements the Stringer interface.
func (x StocksResponseItemStockShipmentType) String() string {
	if str, ok := _StocksResponseItemStockShipmentTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("StocksResponseItemStockShipmentType(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x StocksResponseItemStockShipmentType) IsValid() bool {
	_, ok := _StocksResponseItemStockShipmentTypeMap[x]
	return ok
}

var _StocksResponseItemStockShipmentTypeValue = map[string]StocksResponseItemStockShipmentType{
	_StocksResponseItemStockShipmentTypeName[0:21]:  StocksResponseItemStockShipmentTypeSHIPMENTTYPEGENERAL,
	_StocksResponseItemStockShipmentTypeName[21:38]: StocksResponseItemStockShipmentTypeSHIPMENTTYPEBOX,
	_StocksResponseItemStockShipmentTypeName[38:58]: StocksResponseItemStockShipmentTypeSHIPMENTTYPEPALLET,
}

// ParseStocksResponseItemStockShipmentType attempts to convert a string to a StocksResponseItemStockShipmentType.
func ParseStocksResponseItemStockShipmentType(name string) (StocksResponseItemStockShipmentType, error) {
	if x, ok := _StocksResponseItemStockShipmentTypeValue[name]; ok {
		return x, nil
	}
	return StocksResponseItemStockShipmentType(0), fmt.Errorf("%s is %w", name, ErrInvalidStocksResponseItemStockShipmentType)
}

// MarshalText implements the text marshaller method.
func (x StocksResponseItemStockShipmentType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *StocksResponseItemStockShipmentType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseStocksResponseItemStockShipmentType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

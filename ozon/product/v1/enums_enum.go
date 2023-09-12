// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.5
// Revision: b9e7d1ac24b2b7f6a5b451fa3d21706ffd8d79e2
// Build Date: 2023-01-30T01:49:43Z
// Built By: goreleaser

package v1

import (
	"fmt"
	"strings"
)

const (
	// ImportBySKURequestItemCurrencyCodeRUB is a ImportBySKURequestItemCurrencyCode of type RUB.
	ImportBySKURequestItemCurrencyCodeRUB ImportBySKURequestItemCurrencyCode = iota
	// ImportBySKURequestItemCurrencyCodeBYN is a ImportBySKURequestItemCurrencyCode of type BYN.
	ImportBySKURequestItemCurrencyCodeBYN
	// ImportBySKURequestItemCurrencyCodeKZT is a ImportBySKURequestItemCurrencyCode of type KZT.
	ImportBySKURequestItemCurrencyCodeKZT
	// ImportBySKURequestItemCurrencyCodeEUR is a ImportBySKURequestItemCurrencyCode of type EUR.
	ImportBySKURequestItemCurrencyCodeEUR
	// ImportBySKURequestItemCurrencyCodeUSD is a ImportBySKURequestItemCurrencyCode of type USD.
	ImportBySKURequestItemCurrencyCodeUSD
	// ImportBySKURequestItemCurrencyCodeCNY is a ImportBySKURequestItemCurrencyCode of type CNY.
	ImportBySKURequestItemCurrencyCodeCNY
)

var ErrInvalidImportBySKURequestItemCurrencyCode = fmt.Errorf("not a valid ImportBySKURequestItemCurrencyCode, try [%s]", strings.Join(_ImportBySKURequestItemCurrencyCodeNames, ", "))

const _ImportBySKURequestItemCurrencyCodeName = "RUBBYNKZTEURUSDCNY"

var _ImportBySKURequestItemCurrencyCodeNames = []string{
	_ImportBySKURequestItemCurrencyCodeName[0:3],
	_ImportBySKURequestItemCurrencyCodeName[3:6],
	_ImportBySKURequestItemCurrencyCodeName[6:9],
	_ImportBySKURequestItemCurrencyCodeName[9:12],
	_ImportBySKURequestItemCurrencyCodeName[12:15],
	_ImportBySKURequestItemCurrencyCodeName[15:18],
}

// ImportBySKURequestItemCurrencyCodeNames returns a list of possible string values of ImportBySKURequestItemCurrencyCode.
func ImportBySKURequestItemCurrencyCodeNames() []string {
	tmp := make([]string, len(_ImportBySKURequestItemCurrencyCodeNames))
	copy(tmp, _ImportBySKURequestItemCurrencyCodeNames)
	return tmp
}

var _ImportBySKURequestItemCurrencyCodeMap = map[ImportBySKURequestItemCurrencyCode]string{
	ImportBySKURequestItemCurrencyCodeRUB: _ImportBySKURequestItemCurrencyCodeName[0:3],
	ImportBySKURequestItemCurrencyCodeBYN: _ImportBySKURequestItemCurrencyCodeName[3:6],
	ImportBySKURequestItemCurrencyCodeKZT: _ImportBySKURequestItemCurrencyCodeName[6:9],
	ImportBySKURequestItemCurrencyCodeEUR: _ImportBySKURequestItemCurrencyCodeName[9:12],
	ImportBySKURequestItemCurrencyCodeUSD: _ImportBySKURequestItemCurrencyCodeName[12:15],
	ImportBySKURequestItemCurrencyCodeCNY: _ImportBySKURequestItemCurrencyCodeName[15:18],
}

// String implements the Stringer interface.
func (x ImportBySKURequestItemCurrencyCode) String() string {
	if str, ok := _ImportBySKURequestItemCurrencyCodeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ImportBySKURequestItemCurrencyCode(%d)", x)
}

var _ImportBySKURequestItemCurrencyCodeValue = map[string]ImportBySKURequestItemCurrencyCode{
	_ImportBySKURequestItemCurrencyCodeName[0:3]:   ImportBySKURequestItemCurrencyCodeRUB,
	_ImportBySKURequestItemCurrencyCodeName[3:6]:   ImportBySKURequestItemCurrencyCodeBYN,
	_ImportBySKURequestItemCurrencyCodeName[6:9]:   ImportBySKURequestItemCurrencyCodeKZT,
	_ImportBySKURequestItemCurrencyCodeName[9:12]:  ImportBySKURequestItemCurrencyCodeEUR,
	_ImportBySKURequestItemCurrencyCodeName[12:15]: ImportBySKURequestItemCurrencyCodeUSD,
	_ImportBySKURequestItemCurrencyCodeName[15:18]: ImportBySKURequestItemCurrencyCodeCNY,
}

// ParseImportBySKURequestItemCurrencyCode attempts to convert a string to a ImportBySKURequestItemCurrencyCode.
func ParseImportBySKURequestItemCurrencyCode(name string) (ImportBySKURequestItemCurrencyCode, error) {
	if x, ok := _ImportBySKURequestItemCurrencyCodeValue[name]; ok {
		return x, nil
	}
	return ImportBySKURequestItemCurrencyCode(0), fmt.Errorf("%s is %w", name, ErrInvalidImportBySKURequestItemCurrencyCode)
}

// MarshalText implements the text marshaller method.
func (x ImportBySKURequestItemCurrencyCode) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ImportBySKURequestItemCurrencyCode) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseImportBySKURequestItemCurrencyCode(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
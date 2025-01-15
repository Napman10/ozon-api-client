// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package v2

import (
	"fmt"
	"strings"
)

const (
	// LanguageDEFAULT is a Language of type DEFAULT.
	LanguageDEFAULT Language = iota
	// LanguageRU is a Language of type RU.
	LanguageRU
	// LanguageEN is a Language of type EN.
	LanguageEN
	// LanguageTR is a Language of type TR.
	LanguageTR
	// LanguageZHHANS is a Language of type ZH_HANS.
	LanguageZHHANS
)

var ErrInvalidLanguage = fmt.Errorf("not a valid Language, try [%s]", strings.Join(_LanguageNames, ", "))

const _LanguageName = "DEFAULTRUENTRZH_HANS"

var _LanguageNames = []string{
	_LanguageName[0:7],
	_LanguageName[7:9],
	_LanguageName[9:11],
	_LanguageName[11:13],
	_LanguageName[13:20],
}

// LanguageNames returns a list of possible string values of Language.
func LanguageNames() []string {
	tmp := make([]string, len(_LanguageNames))
	copy(tmp, _LanguageNames)
	return tmp
}

var _LanguageMap = map[Language]string{
	LanguageDEFAULT: _LanguageName[0:7],
	LanguageRU:      _LanguageName[7:9],
	LanguageEN:      _LanguageName[9:11],
	LanguageTR:      _LanguageName[11:13],
	LanguageZHHANS:  _LanguageName[13:20],
}

// String implements the Stringer interface.
func (x Language) String() string {
	if str, ok := _LanguageMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Language(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Language) IsValid() bool {
	_, ok := _LanguageMap[x]
	return ok
}

var _LanguageValue = map[string]Language{
	_LanguageName[0:7]:   LanguageDEFAULT,
	_LanguageName[7:9]:   LanguageRU,
	_LanguageName[9:11]:  LanguageEN,
	_LanguageName[11:13]: LanguageTR,
	_LanguageName[13:20]: LanguageZHHANS,
}

// ParseLanguage attempts to convert a string to a Language.
func ParseLanguage(name string) (Language, error) {
	if x, ok := _LanguageValue[name]; ok {
		return x, nil
	}
	return Language(0), fmt.Errorf("%s is %w", name, ErrInvalidLanguage)
}

// MarshalText implements the text marshaller method.
func (x Language) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *Language) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseLanguage(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

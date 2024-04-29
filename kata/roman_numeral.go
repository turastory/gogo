package kata

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var romanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(n uint16) string {
	var result strings.Builder

	for _, romanNumeral := range romanNumerals {
		for n >= romanNumeral.Value {
			result.WriteString(romanNumeral.Symbol)
			n -= romanNumeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(s string) (arabic uint16) {
	for _, romanNumeral := range romanNumerals {
		for strings.HasPrefix(s, romanNumeral.Symbol) {
			arabic += romanNumeral.Value
			s = strings.TrimPrefix(s, romanNumeral.Symbol)
			// s = s[len(romanNumeral.Symbol):]
		}
	}

	return arabic
}

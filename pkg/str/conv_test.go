// +build all common pkg str conv

// Package str :: conv_test.go

package str

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	translateNumberTests = []struct {
		number   uint64
		commaStr string
		expected string
	}{
		{
			0, "0", "zero",
		},
		{
			3, "3", "three",
		},
		{
			10, "10", "ten",
		},
		{
			100, "100", "one hundred",
		},
		{
			123, "123", "one hundred twenty three",
		},
		{
			200, "200", "two hundreds",
		},
		{
			1000, "1,000", "one thousand",
		},
		{
			1001, "1,001", "one thousand one",
		},
		{
			5000, "5,000", "five thousands",
		},
		{
			65535, "65,535", "sixty five thousands five hundreds thirty five", // 65,535
		},
		{
			7000000, "7,000,000", "seven millions",
		},
		{
			9000000009, "9,000,000,009", "nine billions nine",
		},
		{
			math.MaxUint32,
			"4,294,967,295",
			"four billions two hundreds ninety four millions nine hundreds sixty seven thousands two hundreds ninety five",
		},
		{
			math.MaxUint64,
			"18,446,744,073,709,551,615",
			"eighteen quintillions four hundreds forty six quadrillions seven hundreds forty four trillions seventy three billions seven hundreds nine millions five hundreds fifty one thousands six hundreds fifteen",
		},
	}
)

// TestFromNumber tests str.FromNumber function
func TestFromNumber(t *testing.T) {
	for idx, test := range translateNumberTests {
		result := FromNumber(test.number)
		msg := fmt.Sprintf("Test %2d: %d ==> %s\n",
			idx, test.number, test.expected)
		assert.Equal(t, test.expected, result, msg)
	}
}

// TestTranslate tests str.TranslateNumber and str.TranslateTo functions
func TestTranslate(t *testing.T) {
	for idx, test := range translateNumberTests {
		deform1 := TranslateNumber(test.number, nil)
		deform2 := TranslateTo("de", test.number)
		result1 := TranslateNumber(test.number, FromNumber)
		result2 := TranslateTo("en", test.number)
		msg := fmt.Sprintf("Test %2d: %d ==> [%s] %s\n",
			idx, test.number, test.commaStr, test.expected)
		assert.Equal(t, test.commaStr, deform1, msg)
		assert.Equal(t, test.commaStr, deform2, msg)
		assert.Equal(t, test.expected, result1, msg)
		assert.Equal(t, test.expected, result2, msg)
	}
}

// TestToCamel tests str.ToCamel function
func TestToCamel(t *testing.T) {
	for idx, test := range []struct {
		strInput string
		expected string
		withCaps string
	}{
		{"", "", ""},
		{"Camel case", "CamelCase", "CamelCase"},
		{"dash-board viewer", "DashBoardViewer", "DashBoardViewer"},
		{"Made in China", "MadeInChina", "MadeInChina"},
		{"Made in USA", "MadeInUsa", "MadeInUSA"},
		{"snake_test_case", "SnakeTestCase", "SnakeTestCase"},
		{"start PDF loader", "StartPdfLoader", "StartPDFLoader"},
		{"This is a test", "ThisIsATest", "ThisIsATest"},
		{"привет_мир12345678", "ПриветМир12345678", "ПриветМир12345678"},
		{"中文 没有 大小写", "中文没有大小写", "中文没有大小写"},
		{"here.we.go", "HereWeGo", "HereWeGo"},
	} {
		msg := fmt.Sprintf("Test %2d: %s ==> %s [%s]\n",
			idx, test.strInput, test.expected, test.withCaps)
		assert.Equal(t, test.expected, ToCamel(test.strInput), msg)
		assert.Equal(t, test.withCaps, ToCamel(test.strInput, true), msg)
	}
}

// TestToSnake tests str.ToSnake function
func TestToSnake(t *testing.T) {
	for idx, test := range []struct {
		strInput string
		expected string
	}{
		{"", ""},
		{"@Me", "@_me"},
		{"#Me", "#_me"},
		{"1Country", "1_country"},
		{"a", "a"},
		{"A", "a"},
		{"aToken", "a_token"},
		{"BarID", "bar_id"},
		{"FooTest", "foo_test"},
		{"ID", "id"},
		{"MadeInChina", "made_in_china"},
		{"onlysmallcase", "onlysmallcase"},
		{"PaaS", "paa_s"},
		{"pT", "p_t"},
		{"SID", "sid"},
		{"Snake_Case", "snake_case"},
		{"Snake-Case", "snake-_case"},
		{"SomeCapCase", "some_cap_case"},
		{"USA123", "usa123"},
		{"ПриветМир12345678", "привет_мир12345678"},
		{"中文没有大小写", "中文没有大小写"},
	} {
		msg := fmt.Sprintf("Test %2d: %s ==> %s\n", idx, test.strInput, test.expected)
		assert.Equal(t, test.expected, ToSnake(test.strInput), msg)
	}
}

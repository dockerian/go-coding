// +build all common pkg str conv

// Package str :: conv_test.go

package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

// +build all common pkg sig uuid md5

// Package sig :: sig_test.go

package sig

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	regexUUID = `[a-z0-9]{8}-([a-z0-9]{4}-){3}[a-z0-9]{12}`
)

// ruleTest to test Rule interface
type ruleTest struct {
	id string
	ib string
}

// Output implements Rule interface
func (rule *ruleTest) Output() string {
	return strings.Repeat(rule.id, 10)
}

// OutputIB implements IBRule interface
func (rule *ruleTest) OutputIB() string {
	return strings.Repeat(rule.ib, 10)
}

// TestCreateSources
func TestCreateSources(t *testing.T) {
	rules := make([]Rule, 0)
	for i := 0; i < 10; i++ {
		sib := 'a' - i
		rules = append(rules, &ruleTest{id: strconv.Itoa(i), ib: string(sib)})
	}
	result := CreateSources(rules, "")
	assert.Equal(t, 2, len(result))
	assert.Contains(t, result[0].Name, ddosRuleName)
	assert.Contains(t, result[1].Name, defaultIBRuleName)
}

// TestGetGUID tests sig.GetGUID
func TestGetGUID(t *testing.T) {
	prevUUID := ""
	prevName := "don't use for following test"
	for idx, test := range []string{
		"",
		"abc.com",
		"domain.com",
		"www.domain.com",
		"just a name",
		"",
		"",
	} {
		result := GetGUID(test).String()
		t.Logf("Test %2d: '%s' ==> '%s'\n", idx, test, result)
		assert.True(t, len(result) == 36)

		matched, err := regexp.MatchString(regexUUID, result)
		assert.True(t, matched)
		assert.Nil(t, err)

		if test == "" || test != prevName {
			assert.NotEqual(t, prevUUID, result)
		}
		prevUUID = result
		prevName = test
	}
}

// TestGetPattern tests sig.GetPattern
func TestGetPattern(t *testing.T) {
	for idx, test := range []struct {
		domain, pattern string
	}{
		{"", "|00|"},
		{"0bbxx.com", "|05|0bbxx|03|com|00|"},
		{"1156789012.com", "|0a|1156789012|03|com|00|"},
		{"91duofenxiang.org", "|0d|91duofenxiang|03|org|00|"},
		{"lafengzhuangyuan.cn", "|10|lafengzhuangyuan|02|cn|00|"},
		{"wbfastpay.com", "|09|wbfastpay|03|com|00|"},
		{"xcailing.net", "|08|xcailing|03|net|00|"},
	} {
		result := GetPattern(test.domain)
		t.Logf("Test %2d: '%s' ==> '%s'\n", idx, test.domain, test.pattern)
		assert.Equal(t, test.pattern, result)
	}
}

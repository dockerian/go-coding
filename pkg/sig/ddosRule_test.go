// +build all common pkg sig rule ddos

// Package sig :: ddosRule_test.go

package sig

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	regexRuleUUID = `RUL-[a-z0-9]{8}-([a-z0-9]{4}-){3}[a-z0-9]{12}`
)

// TestDDosRuleOutputDDosRule tests DDosRule method OutputDDosRule
func TestDDosRuleOutputDDosRule(t *testing.T) {
	for idx, test := range []struct {
		domain, firstSeen, lastSeen, output string
	}{
		{
			domain:    "coxhn.net",
			firstSeen: "2016-01-08 00:00:00",
			lastSeen:  "2017-10-06 00:00:00",
			output:    `drop udp any any -> any 53 (msg:"DNS attempted resource exhaustion: coxhn.net"; content:"|05|coxhn|03|net|00|"; classtype:attempted-dos; sid:3030303; rev:1;) #first seen:2016-01-08 00:00:00 - last seen:2017-10-06 00:00:00`,
		},
	} {
		t.Logf("Test %2d: %s\n", idx, test.domain)
		ddosRule := NewDDosRule(test.domain)
		assert.Equal(t, test.domain, ddosRule.Domain)
		assert.Equal(t, "attempted-dos", ddosRule.RuleType)
		assert.Equal(t, int32(3030303), ddosRule.SID)
		assert.Equal(t, 1, ddosRule.SigRev)

		firstSeen, _ := time.Parse("2006-01-02 15:04:05", test.firstSeen)
		lastSeen, _ := time.Parse("2006-01-02 15:04:05", test.lastSeen)
		expected := strings.TrimSpace(test.output)
		output := ddosRule.OutputDDosRule(firstSeen, lastSeen)
		assert.Equal(t, expected, output)
	}
}

// TestDDosRuleOutputIB tests DDosRule method OutputIB
func TestDDosRuleOutputIB(t *testing.T) {
	for idx, test := range []struct {
		domain, output string
	}{
		{
			domain: "bad-bad.com",
			output: `
#IB_CATEGORY: Potential DDoS related Domains; IB_SET: 1; IB_TYPE: SYSTEM; IB_LOG: MAJOR; IB_DISABLED: False; IB_DESC: "This rule blacklists bad-bad.com, which has been observed to be used in DDoS attacks.";
#IB_PARAM_EVENT_PER_SECOND:1;
drop udp any any -> any 53 (msg:"Potential DDoS related domain: bad-bad.com"; content:"|07|bad-bad|03|com|00|"; offset:12; classtype:attempted-dos; sid:3030303; rev:1;)
event_filter gen_id 1, sig_id 3030303, type limit, track by_src, count 1, seconds 1
`,
		},
		{
			domain: "coxhn.net",
			output: `
#IB_CATEGORY: Potential DDoS related Domains; IB_SET: 1; IB_TYPE: SYSTEM; IB_LOG: MAJOR; IB_DISABLED: False; IB_DESC: "This rule blacklists coxhn.net, which has been observed to be used in DDoS attacks.";
#IB_PARAM_EVENT_PER_SECOND:1;
drop udp any any -> any 53 (msg:"Potential DDoS related domain: coxhn.net"; content:"|05|coxhn|03|net|00|"; offset:12; classtype:attempted-dos; sid:3030303; rev:1;)
event_filter gen_id 1, sig_id 3030303, type limit, track by_src, count 1, seconds 1
`,
		},
	} {
		t.Logf("Test %2d: %s\n", idx, test.domain)
		ddosRule := NewDDosRule(test.domain)
		assert.Equal(t, test.domain, ddosRule.Domain)
		assert.Equal(t, "attempted-dos", ddosRule.RuleType)
		assert.Equal(t, int32(3030303), ddosRule.SID)
		assert.Equal(t, 1, ddosRule.SigRev)

		expected := strings.TrimSpace(test.output)
		output := ddosRule.OutputIB()
		assert.Equal(t, expected, output)
	}

}

// TestGetRuleGUID tests sig.GetRuleGUID
func TestGetRuleGUID(t *testing.T) {
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
		result := GetRuleGUID(test)
		t.Logf("Test %2d: '%s' ==> '%s'\n", idx, test, result)
		assert.True(t, len(result) == 40)

		matched, err := regexp.MatchString(regexRuleUUID, result)
		assert.True(t, matched)
		assert.Nil(t, err)

		if test == "" || test != prevName {
			assert.NotEqual(t, prevUUID, result)
		}
		prevUUID = result
		prevName = test
	}
}

// Package sig :: ddosRule.go - DDoS rule implementation
package sig

import (
	"fmt"
	"time"
)

var (
	// ddosRuleType defines DDoS rule type
	ddosRuleType = "attempted-dos"
	// ddosRuleSID defines DDoS rule SID
	ddosRuleSID = 3030303
	// ddosSigRev defines DDos rule signature rev
	ddosSigRev = 1
)

// DDosRule struct defines a rule for DDoS domain
type DDosRule struct {
	// Domain name
	Domain string
	// Domain rule GUID
	GUID string
	// DDoS rule signature pattern
	Pattern string
	// Domain rule type
	RuleType string
	// Domain rule signature revision
	SigRev int
	// Domain rule SID
	SID int32
}

// GetDDosIBRule returns DDoS rule IB output
func GetDDosIBRule(domain, sigPattern string) string {
	header, footer, part1, part2 := GetDDosIBRuleFormatters()

	ruleHeader := fmt.Sprintf(header, domain)
	ruleFooter := fmt.Sprintf(footer, ddosRuleSID)
	rulesPart1 := fmt.Sprintf(part1, domain)
	rulesPart2 := fmt.Sprintf(part2, sigPattern, ddosRuleType, ddosRuleSID, ddosSigRev)

	output := fmt.Sprintf(
		"%s\n%s %s\n%s", ruleHeader, rulesPart1, rulesPart2, ruleFooter)

	return output
}

// GetDDosIBRuleFormatters returns Infoblox (IB) rule formatters
// Here are placeholders in order
//   - header formatter: domain
//   - part-1 formatter: domain
//   - part-2 formatter: pattern, rule type, sid, sig rev
//   - footer formatter: sid
// Formatter string to generate full IB rule output
//   - "%s\n%s %s\n%s"
func GetDDosIBRuleFormatters() (header, footer, part1, part2 string) {
	// buiding header formtter
	ibCategory := fmt.Sprintf("IB_CATEGORY: %s", categoryDDoS)
	ibDesc := "IB_DESC: \"This rule blacklists %v," // domain placeholder
	ibDescDDoS := "which has been observed to be used in DDoS attacks.\""
	ibFlags := "IB_SET: 1; IB_TYPE: SYSTEM; IB_LOG: MAJOR; IB_DISABLED: False"
	ibParam := "IB_PARAM_EVENT_PER_SECOND:1"
	header = fmt.Sprintf("#%s; %s; %s %s;\n#%s;",
		ibCategory, ibFlags, ibDesc, ibDescDDoS, ibParam)

	// formatting DDoS rule in 2 parts
	proto := "udp"
	action := "drop"
	message := "Potential DDoS related domain: %v" // domain placeholder
	// ruleType := "attempted-dos"
	part1 = fmt.Sprintf("%s %s any any -> any 53 (msg:\"%s\";", action, proto, message)
	// sig pattern, rule type, sid, sig rev placeholders
	part2 = "content:\"%v\"; offset:12; classtype:%v; sid:%d; rev:%v;)"

	// building footer formtter
	ibFilter := "event_filter gen_id 1, sig_id %d, type limit," // sid placeholder
	footer = fmt.Sprintf("%s track by_src, count 1, seconds 1", ibFilter)

	return
}

// GetDDosRuleFormatter returns DDoS rule formatter, with placeholders:
// domain name, sig pattern, rule type, sid, sig rev, first seen, last seen
func GetDDosRuleFormatter() string {
	proto := "udp"
	action := "drop"
	message := "DNS attempted resource exhaustion: %v" // domain placeholder

	part1 := fmt.Sprintf("%s %s any any -> any 53 (msg:\"%s\";", action, proto, message)
	// sig pattern, rule type, sid, sig rev placeholders
	part2 := "content:\"%v\"; classtype:%v; sid:%d; rev:%v;)"
	// first-seen, last-seen placeholders
	part3 := "#first seen:%v - last seen:%v"

	return fmt.Sprintf("%s %s %s", part1, part2, part3)
}

// GetDDosRule returns DDoS rule output
func GetDDosRule(domain, sigPattern string, firstSeen, lastSeen time.Time) string {
	formatter := GetDDosRuleFormatter()
	s1 := firstSeen.Format("2006-01-02 15:04:05")
	s2 := lastSeen.Format("2006-01-02 15:04:05")

	output := fmt.Sprintf(formatter, domain,
		sigPattern, ddosRuleType, ddosRuleSID, ddosSigRev, s1, s2)

	return output
}

// GetRuleGUID returns a rule GUID with prefix "RUL-",
func GetRuleGUID(domain string) string {
	return fmt.Sprintf("RUL-%s", GetGUID(domain))
}

// NewDDosRule constucts a DomainRule by domain name
func NewDDosRule(domain string) *DDosRule {
	return &DDosRule{
		Domain:   domain,
		GUID:     GetRuleGUID(domain),
		Pattern:  GetPattern(domain),
		RuleType: ddosRuleType,
		SID:      int32(ddosRuleSID),
		SigRev:   ddosSigRev,
	}
}

// OutputDDosRule returns a ddos rule ouptput
func (ddosRule *DDosRule) OutputDDosRule(firstSeen, lastSeen time.Time) string {
	return GetDDosRule(ddosRule.Domain, ddosRule.Pattern, firstSeen, lastSeen)
}

// OutputIB implements IBRule interface to produce DDoS rule output for Infoblox NIO
func (ddosRule *DDosRule) OutputIB() string {
	return GetDDosIBRule(ddosRule.Domain, ddosRule.Pattern)
}

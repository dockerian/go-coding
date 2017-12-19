// Package sig :: sig.go - signature interface
package sig

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dockerian/go-coding/pkg/zip"
	"github.com/satori/go.uuid"
)

var (
	// category name for DDoS
	categoryDDoS = "Potential DDoS related Domains"
	// default IB (Infoblox) rule name
	defaultIBRuleName = "InfobloxRule-4.0"
	// default DDoS rule name
	ddosRuleName = "ddos_domains"
)

// IBRule interface represents Infoblox rule
type IBRule interface {
	// OutputIB Infoblox rule to formatted string
	OutputIB() string
}

// Rule interface represents rule and signature
type Rule interface {
	IBRule
	// Output rule to formatted string
	Output() string
}

// CreateSources prepares zip sources from all rules output
func CreateSources(rules []Rule, ruleName string) []*zip.Source {
	if strings.TrimSpace(ruleName) == "" {
		ruleName = ddosRuleName
	}
	var output bytes.Buffer
	var outputIB bytes.Buffer
	var size, sizeIB int
	var yyyymmdd = time.Now().UTC().Format("20060102")
	var filename = fmt.Sprintf("%s-%s.txt", ruleName, yyyymmdd)
	var ruleIBFn = fmt.Sprintf("%s-%s.txt", defaultIBRuleName, yyyymmdd)

	for _, rule := range rules {
		outputBytes := []byte(rule.Output() + "\r\n")
		output.Write(outputBytes)
		size += len(outputBytes)
		outputIBBytes := []byte(rule.OutputIB() + "\r\n")
		outputIB.Write(outputIBBytes)
		sizeIB += len(outputIBBytes)
	}

	return []*zip.Source{
		{Name: filename, Reader: &output, Size: size},
		{Name: ruleIBFn, Reader: &outputIB, Size: sizeIB},
	}
}

// GetGUID generates uuid v3 from a specific name
// See http://antoniomo.com/blog/2017/05/21/unique-ids-in-golang-part-1/
func GetGUID(name string) uuid.UUID {
	input := name
	if input == "" {
		input = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
		return uuid.NewV3(uuid.NamespaceDNS, input)
	}
	bytes := []byte(input)
	md5dg := md5.New().Sum(bytes)
	uuid3 := uuid.NewV3(uuid.FromBytesOrNil(md5dg), input)
	return uuid3
}

// GetPattern returns a signature pattern by domain name
func GetPattern(domain string) string {
	pattern := ""
	sfields := strings.Split(domain, ".")
	for _, field := range sfields {
		if field != "" {
			pattern += fmt.Sprintf("|%02x|%s", len(field), field)
		}
	}
	pattern += "|00|"

	return pattern
}

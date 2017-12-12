# sig
--
    import "github.com/dockerian/go-coding/pkg/sig"

Package sig :: ddosRule.go - DDoS rule implementation

Package sig :: sig.go - signature interface

## Usage

#### func  CreateSources

```go
func CreateSources(rules []Rule, ruleName string) []*zip.Source
```
CreateSources prepares zip sources from all rules output

#### func  GetDDosIBRule

```go
func GetDDosIBRule(domain, sigPattern string) string
```
GetDDosIBRule returns DDoS rule IB output

#### func  GetDDosIBRuleFormatters

```go
func GetDDosIBRuleFormatters() (header, footer, part1, part2 string)
```
GetDDosIBRuleFormatters returns Infoblox (IB) rule formatters Here are
placeholders in order

    - header formatter: domain
    - part-1 formatter: domain
    - part-2 formatter: pattern, rule type, sid, sig rev
    - footer formatter: sid

Formatter string to generate full IB rule output

    - "%s\n%s %s\n%s"

#### func  GetDDosRule

```go
func GetDDosRule(domain, sigPattern string, firstSeen, lastSeen time.Time) string
```
GetDDosRule returns DDoS rule output

#### func  GetDDosRuleFormatter

```go
func GetDDosRuleFormatter() string
```
GetDDosRuleFormatter returns DDoS rule formatter, with placeholders: domain
name, sig pattern, rule type, sid, sig rev, first seen, last seen

#### func  GetGUID

```go
func GetGUID(name string) uuid.UUID
```
GetGUID generates uuid v3 from a specific name See
http://antoniomo.com/blog/2017/05/21/unique-ids-in-golang-part-1/

#### func  GetPattern

```go
func GetPattern(domain string) string
```
GetPattern returns a signature pattern by domain name

#### func  GetRuleGUID

```go
func GetRuleGUID(domain string) string
```
GetRuleGUID returns a rule GUID with prefix "RUL-",

#### type DDosRule

```go
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
```

DDosRule struct defines a rule for DDoS domain

#### func  NewDDosRule

```go
func NewDDosRule(domain string) *DDosRule
```
NewDDosRule constucts a DomainRule by domain name

#### func (*DDosRule) OutputDDosRule

```go
func (ddosRule *DDosRule) OutputDDosRule(firstSeen, lastSeen time.Time) string
```
OutputDDosRule returns a ddos rule ouptput

#### func (*DDosRule) OutputIB

```go
func (ddosRule *DDosRule) OutputIB() string
```
OutputIB implements IBRule interface to produce DDoS rule output for Infoblox
NIO

#### type IBRule

```go
type IBRule interface {
	// OutputIB Infoblox rule to formatted string
	OutputIB() string
}
```

IBRule interface represents Infoblox rule

#### type Rule

```go
type Rule interface {
	IBRule
	// Output rule to formatted string
	Output() string
}
```

Rule interface represents rule and signature

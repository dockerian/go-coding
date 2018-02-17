# str
--
    import "github.com/dockerian/go-coding/pkg/str"

Package str :: conv.go - extended string formatter functions

Package str :: palindrome.go

Package str :: str.go - extended string functions

## Usage

#### func  Append

```go
func Append(slice, data []byte) []byte
```
Append concatenates byte slices

#### func  FormatNumber

```go
func FormatNumber(number uint64) string
```
FormatNumber returns a comma delimited decimal string

#### func  FromNumber

```go
func FromNumber(number uint64) string
```
FromNumber returns an English words representation for a number. ex. 1024 =>
"one thousand twenty four"

#### func  GetPalindromicSubstring

```go
func GetPalindromicSubstring(str string) string
```
GetPalindromicSubstring returns the longest palindromic substring

#### func  IsPalindrome

```go
func IsPalindrome(input interface{}) (bool, error)
```
IsPalindrome checks if interface{} is palindrome

#### func  IsPalindromeNumber

```go
func IsPalindromeNumber(input uint64) bool
```
IsPalindromeNumber checks if an input number is palindrome

#### func  IsPalindromePhase

```go
func IsPalindromePhase(input string) bool
```
IsPalindromePhase checks if an input string is palindrome phase

#### func  IsPalindromeString

```go
func IsPalindromeString(input string) bool
```
IsPalindromeString checks if an input string is palindrome

#### func  ReplaceProxyURL

```go
func ReplaceProxyURL(url, prefix, proxyURL string) string
```
ReplaceProxyURL searches prefix in url and replaces with proxyURL

#### func  StringIn

```go
func StringIn(stringInput string, stringList []string, options ...bool) bool
```
StringIn check if an input is in an array of strings; optional to ignore case

#### func  ToCamel

```go
func ToCamel(in string, keepAllCaps ...bool) string
```
ToCamel converts a string to camel case format

#### func  ToSnake

```go
func ToSnake(in string) string
```
ToSnake converts a string to snake case format with unicode support See also
https://github.com/serenize/snaker/blob/master/snaker.go

#### func  TranslateNumber

```go
func TranslateNumber(number uint64, xFunc TranslateFunc) string
```
TranslateNumber translates a number to string by specific function.

#### func  TranslateTo

```go
func TranslateTo(lang string, number uint64) string
```
TranslateTo returns a string representation of number by specific language.

#### type Palindrome

```go
type Palindrome interface {
	GetData() string
	IsPalindrome() bool
}
```

Palindrome interface

#### type PalindromeNumber

```go
type PalindromeNumber struct {
}
```

PalindromeNumber struct

#### func (*PalindromeNumber) GetData

```go
func (p *PalindromeNumber) GetData() string
```
GetData returns input data

#### func (*PalindromeNumber) IsPalindrome

```go
func (p *PalindromeNumber) IsPalindrome() bool
```
IsPalindrome checks if an input number is palindrome

#### type PalindromeString

```go
type PalindromeString struct {
}
```

PalindromeString struct

#### func (*PalindromeString) GetData

```go
func (p *PalindromeString) GetData() string
```
GetData returns input data

#### func (*PalindromeString) GetSubstring

```go
func (p *PalindromeString) GetSubstring() string
```
GetSubstring returns the longest palindromic substring

#### func (*PalindromeString) IsPalindrome

```go
func (p *PalindromeString) IsPalindrome() bool
```
IsPalindrome checks if an input string is palindrome

#### func (*PalindromeString) IsPalindromePhase

```go
func (p *PalindromeString) IsPalindromePhase() bool
```
IsPalindromePhase checks if an input string is palindrome phase

#### type TranslateFunc

```go
type TranslateFunc func(uint64) string
```

TranslateFunc defines a type of function to translate number (uint64) to string.

# utils
--
    import "github.com/dockerian/go-coding/utils"

Package utils :: aws.go - extended AWS SDK functions

Package utils :: bit.go

Package utils :: config.go Get project settings from os env or specified
config.yaml

Package utils :: conv.go - extended string formatter functions

Package utils :: env.go - extended os env functions

Package utils :: fmt.go - extended fmt functions

Package utils :: interface.go

Package utils :: io.go

Package utils :: logging.go

Package utils :: pair.go

Package utils :: string.go

Package utils :: testing.go

Package utils :: utf.go

## Usage

```go
var (

	// BinaryString is an array of all 4-bit binary representation
	BinaryString = map[string][]string{
		"b": []string{
			"0000", "0001", "0010", "0011",
			"0100", "0101", "0110", "0111",
			"1000", "1001", "1010", "1011",
			"1100", "1101", "1110", "1111",
		},
		"x": []string{
			"0", "1", "2", "3", "4", "5", "6", "7",
			"8", "9", "A", "B", "C", "D", "E", "F",
		},
	}
)
```

```go
var DebugEnv = CheckEnvBoolean("DEBUG", true)
```
DebugEnv indicates DEBUG = 1|on|true in environment variable (ignoring case)

```go
var DebugEnvStore = DebugEnv
```
DebugEnvStore keeps a backup of DebugEnv value

#### func  BitAllOne

```go
func BitAllOne() int64
```
BitAllOne returns integer with all bits are 1

#### func  BitCheck

```go
func BitCheck(x int64, n uint8) bool
```
BitCheck checks on nth bit of x

#### func  BitClear

```go
func BitClear(x int64, n uint8) int64
```
BitClear sets 0 on nth bit of x

#### func  BitCountOne

```go
func BitCountOne(x int64) int
```
BitCountOne returns number of 1 in x (aka Hamming weight)

#### func  BitCountOneUint64

```go
func BitCountOneUint64(x uint64) int
```
BitCountOneUint64 returns number of 1 in x (aka Hamming weight)

#### func  BitIntersection

```go
func BitIntersection(a, b int64) int64
```
BitIntersection applies bitwise AND (&) operator on a and b (interaction)

#### func  BitInverse

```go
func BitInverse(x int64) int64
```
BitInverse returns inverted x

#### func  BitIsPowerOf2

```go
func BitIsPowerOf2(number int64) bool
```
BitIsPowerOf2 checks if the number is power of 2

#### func  BitIsPowerOf4

```go
func BitIsPowerOf4(number int64) bool
```
BitIsPowerOf4 checks if the number is power of 4

#### func  BitNegativeInt

```go
func BitNegativeInt(x int64) int64
```
BitNegativeInt returns negative number of x (0 - x)

#### func  BitSet

```go
func BitSet(x int64, n uint8) int64
```
BitSet sets 1 on nth bit of x

#### func  BitString

```go
func BitString(x uint64, key, delimiter string) string
```
BitString converts uint64 x to zero-padding bits representation

#### func  BitSubstraction

```go
func BitSubstraction(a, b int64) int64
```
BitSubstraction applies A & ~B

#### func  BitSumInt

```go
func BitSumInt(x, y int) int
```
BitSumInt calculates sum of two integers without using arithmetic operators

#### func  BitSumInt64

```go
func BitSumInt64(x, y int64) int64
```
BitSumInt64 calculates sum of two integers without using arithmetic operators

#### func  BitUnion

```go
func BitUnion(a, b int64) int64
```
BitUnion applies bitwise OR (|) operator on a and b (union)

#### func  CheckEnvBoolean

```go
func CheckEnvBoolean(name string, ignoreCase bool) bool
```
CheckEnvBoolean checks if an environment variable is set to non-false/non-zero

#### func  Debug

```go
func Debug(format string, v ...interface{})
```
Debug prints logging message if DEBUG is set

#### func  DebugOff

```go
func DebugOff()
```
DebugOff turns off debug

#### func  DebugOn

```go
func DebugOn()
```
DebugOn turns on debug

#### func  DebugReset

```go
func DebugReset()
```
DebugReset reset to debug mode by environment setting

#### func  Debugln

```go
func Debugln(a ...interface{})
```
Debugln prints logging message (with new line) if DEBUG is set

#### func  DecryptKeyTextByKMS

```go
func DecryptKeyTextByKMS(key, text string) string
```
DecryptKeyTextByKMS checks possible encrypted KMS key/value and retruns
decrypted text

#### func  Flatten

```go
func Flatten(prefix string, value interface{}, kvmap map[string]string)
```
Flatten builds a flattened key/value string pairs map

#### func  FlattenConfig

```go
func FlattenConfig(file string) map[string]string
```
FlattenConfig loads a config file (.yaml) to flattened key/value map

#### func  FmtComma

```go
func FmtComma(number string) string
```
FmtComma formats number with thousands comma

#### func  GetEnvron

```go
func GetEnvron() map[string]string
```
GetEnvron get a map of environment variables

#### func  GetGraphemeCount

```go
func GetGraphemeCount(str string) int
```
GetGraphemeCount function

#### func  GetGraphemeCountInString

```go
func GetGraphemeCountInString(str string) int
```
GetGraphemeCountInString function

#### func  GetSliceAtIndex

```go
func GetSliceAtIndex(input string, index int) string
```
GetSliceAtIndex returns indexed one-byte slice, or empty string

#### func  GetTernary

```go
func GetTernary(condition bool, a interface{}, b interface{}) interface{}
```
GetTernary returns a if condition is true; otherwise returns b

#### func  GetTestName

```go
func GetTestName(t *testing.T) string
```
GetTestName returns the name of the test function from testing.T.

#### func  GetTestNameByCaller

```go
func GetTestNameByCaller() string
```
GetTestNameByCaller returns the name of the test function from the call stack.

#### func  HasEnv

```go
func HasEnv(name string, ignoreCase bool) bool
```
HasEnv checks if an environment variable is set

#### func  IsDigitOrLetter

```go
func IsDigitOrLetter(char rune) bool
```
IsDigitOrLetter checks if a unicode char is digit or letter

#### func  Readlines

```go
func Readlines(r *bufio.Reader) ([]string, error)
```
Readlines returns lines (without the ending \n) from buffered reader,
additionally returns error from buffered reader if there is any.

#### func  ReplaceProxyURL

```go
func ReplaceProxyURL(url, prefix, proxyURL string) string
```
ReplaceProxyURL searches prefix in url and replaces with proxyURL

#### func  ShiftSlice

```go
func ShiftSlice(input string, shift int) string
```
ShiftSlice returns slice by shift index

#### func  StringIn

```go
func StringIn(stringInput string, stringList []string, options ...bool) bool
```
StringIn check if an input is in an array of strings; optional to ignore case

#### func  ToBinaryString

```go
func ToBinaryString(x uint64, delimiter string) string
```
ToBinaryString converts uint64 x to zero-padding binary representation

#### func  ToCamel

```go
func ToCamel(in string, keepAllCaps ...bool) string
```
ToCamel converts a string to camel case format

#### func  ToHexString

```go
func ToHexString(x uint64, delimiter string) string
```
ToHexString converts uint64 x to zero-padding hexadecimal representation

#### func  ToJSON

```go
func ToJSON(t interface{}) string
```
ToJSON function returns pretty-printed JSON for a struct.

#### func  ToSnake

```go
func ToSnake(in string) string
```
ToSnake converts a string to snake case format with unicode support See also
https://github.com/serenize/snaker/blob/master/snaker.go

#### type Config

```go
type Config struct {
}
```

Config represents a flattened settings per config file

#### func  GetConfig

```go
func GetConfig(file string) *Config
```
GetConfig gets a singleton instance of Config

#### func (Config) Get

```go
func (c Config) Get(key string, defaultValues ...string) string
```
Get gets string value of the key in os.Environ() or Config.settings or using
defaultValues[0] if provided; otherwise return ""

#### func (Config) GetBool

```go
func (c Config) GetBool(key string, defaultValues ...bool) bool
```
GetBool gets boolean value of the key, or defaultValues[0], or false

#### func (Config) GetInt32

```go
func (c Config) GetInt32(key string, defaultValues ...int32) int32
```
GetInt32 gets int32 value of the key, or defaultValues[0], or 0

#### func (Config) GetInt64

```go
func (c Config) GetInt64(key string, defaultValues ...int64) int64
```
GetInt64 gets int64 value of the key, or defaultValues[0], or 0

#### func (Config) GetUint32

```go
func (c Config) GetUint32(key string, defaultValues ...uint32) uint32
```
GetUint32 gets uint32 value of the key, or defaultValues[0], or 0

#### func (Config) GetUint64

```go
func (c Config) GetUint64(key string, defaultValues ...uint64) uint64
```
GetUint64 gets uint64 value of the key, or defaultValues[0], or 0

#### type ConfigParserFunc

```go
type ConfigParserFunc func([]byte, interface{}) error
```

ConfigParserFunc is a generic parser function

#### type ConfigReaderFunc

```go
type ConfigReaderFunc func(string) ([]byte, error)
```

ConfigReaderFunc is a generic reader function

#### type DecryptFunc

```go
type DecryptFunc func(string, string) string
```

DecryptFunc is a generic decrypt function

#### type Env

```go
type Env map[string]interface{}
```

Env struct stores application-wide configuration

#### func (Env) Delete

```go
func (env Env) Delete(key string)
```
Delete removes a key and the mapping value

#### func (Env) Get

```go
func (env Env) Get(key string) string
```
Get returns string for the mapping value by the key

#### func (Env) GetInt

```go
func (env Env) GetInt(key string) int
```
GetInt returns int for the mapping value by the key

#### func (Env) GetValue

```go
func (env Env) GetValue(key string) interface{}
```
GetValue returns the mapping value by the key

#### func (Env) Set

```go
func (env Env) Set(key string, value interface{})
```
Set overwrite the mapping value by the key

#### type KMSDecryptInterface

```go
type KMSDecryptInterface interface {
	Decrypt(input *kms.DecryptInput) (*kms.DecryptOutput, error)
}
```

KMSDecryptInterface interface

#### type Pair

```go
type Pair struct {
	Item1, Item2 interface{}
}
```

Pair struct respresents a pair of anything

#### func (*Pair) AreEqual

```go
func (p *Pair) AreEqual(other *Pair) bool
```
AreEqual method receiver compares 'this' Pair to 'other' Pair

#### func (*Pair) String

```go
func (p *Pair) String() string
```
String method receiver for Pair struct

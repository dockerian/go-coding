# cfg
--
    import "github.com/dockerian/go-coding/pkg/cfg"

Package cfg :: aws.go - extended AWS SDK functions

Package cfg :: config.go Get project settings from os env or specified
config.yaml

Package cfg :: context.go

Package cfg :: env.go

## Usage

#### func  Decrypt

```go
func Decrypt(text string) string
```
Decrypt tries to decrypt a text

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

#### type Context

```go
type Context struct {
	context.Context
	Cookie  *http.Cookie
	Session *sessions.Session
	Env     *Env
}
```

Context struct wraps Env, http.Cookie, gorilla Session, and Context

#### func (*Context) Value

```go
func (ctx *Context) Value(key interface{}) interface{}
```
Value implements context.Context

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

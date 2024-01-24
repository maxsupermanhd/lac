# lac

Live application config

LAC is a configuration library for Golang with focus on gorutine safety.

Entierly for use in my own projects, no guarantees on anything.

## General methods and usage

All operations are guarded with a mutex, it is safe to get/set/load/safe from within multiple gorutines.

### Loading/saving

```go
func FromBytesJSON(b []byte) (*Conf, error)
func FromFileJSON(path string) (*Conf, error)

func (c *Conf) SetFromBytesJSON(b []byte) error
func (c *Conf) SetFromFileJSON(path string) error

func (c *Conf) ToBytesIndentJSON() ([]byte, error)
func (c *Conf) ToFileIndentJSON(path string, perm os.FileMode) error
```

### Getting/setting data

Raw get operations. They return **copy** of the data in the config tree.

```go
func (c *Conf) Get(k ...string) (any, bool)
func (c *Conf) GetString(k ...string) (string, bool)
func (c *Conf) GetFloat64(k ...string) (float64, bool)
func (c *Conf) GetInt64(k ...string) (int64, bool)
func (c *Conf) GetInt(k ...string) (int, bool) // just trimmed int64, beware of overflows
func (c *Conf) GetBool(k ...string) (bool, bool)
func (c *Conf) GetMapStringAny(k ...string) (map[string]any, bool)
```

Use the prefix "D" to specify a default value to return if value in config tree does not exist or of a wrong type.

```go
func (c *Conf) GetDString(d string, k ...string) string
func (c *Conf) GetDFloat64(d float64, k ...string) float64
func (c *Conf) GetDInt64(d int64, k ...string) int64
func (c *Conf) GetDInt(d int, k ...string) int
func (c *Conf) GetDBool(d bool, k ...string) bool
```

Combine with prefix "S" to also set the default if it is of wrong type or does not exist. (To initialize tree for later editing/writing)

```go
func (c *Conf) GetDSString(d string, k ...string) string
func (c *Conf) GetDSFloat64(d float64, k ...string) float64
func (c *Conf) GetDSInt64(d int64, k ...string) int64
func (c *Conf) GetDSInt(d int, k ...string) int
func (c *Conf) GetDSBool(d bool, k ...string) bool
```

If you just want to list keys of the path, there is a helper for that

```go
func (c *Conf) GetKeys(k ...string) ([]string, bool)
```

Raw set operation is also available to update the configuration tree.

```go
func (c *Conf) Set(v any, k ...string)
```

## Example

```json
{
    "hello": "world",
    "testing": {
        "stuff": 42
    }
}
```

```go
conf, err := lac.FromFileJSON("config.json")
str, ok := conf.GetString("hello") // str = "world", ok = true
str, ok = conf.GetString("hello2") // str = "", ok = false
str = conf.GetDString("a", "doesnotexist") // str = "a"
str = conf.GetDString("a", "hello") // str = "world"
i = conf.GetDInt(0, "testing", "stuff") // i = 42
str = conf.GetDSString("a", "doesnotexist") // str = "a", config updated
err = conf.ToFileIndentJSON("config.json", 0644)
```

```json
{
    "doesnotexist": "a",
    "hello": "world",
    "testing": {
        "stuff": 42
    }
}
```


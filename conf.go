package lac

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/mitchellh/mapstructure"
)

type Conf struct {
	tree map[string]any
	// listeners []ConfListener
	lock sync.Mutex
}

func (c *Conf) ToBytesJSON() ([]byte, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return json.Marshal(c.tree)
}

func (c *Conf) ToBytesIndentJSON() ([]byte, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return json.MarshalIndent(c.tree, "", "\t")
}

func (c *Conf) ToFileJSON(path string, perm os.FileMode) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	b, err := json.Marshal(c.tree)
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, perm)
}

func (c *Conf) ToFileIndentJSON(path string, perm os.FileMode) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	b, err := json.MarshalIndent(c.tree, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, perm)
}

func (c *Conf) SetFromBytesJSON(b []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return json.Unmarshal(b, &c.tree)
}

func (c *Conf) SetFromFileJSON(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return c.SetFromBytesJSON(b)
}

func (c *Conf) SetTree(t map[string]any) {
	c.lock.Lock()
	c.tree = t
	c.lock.Unlock()
}

func (c *Conf) CopyTree(t map[string]any) {
	c.lock.Lock()
	c.tree = copyMap(t)
	c.lock.Unlock()
}

func (c *Conf) Set(v any, k ...string) {
	c.lock.Lock()
	setTree(c.tree, v, k)
	// for _, l := range c.listeners {
	// 	if areStringSlicesEqual(l.path, k) {
	// 		l.f(k, v)
	// 	}
	// }
	c.lock.Unlock()
}

func (c *Conf) GetToStruct(t any, k ...string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return ErrNoKey
	}
	return mapstructure.Decode(v, t)
}

func (c *Conf) Get(k ...string) (any, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return nil, false
	}
	switch vv := v.(type) {
	case string, float64:
		return vv, true
	case map[string]any:
		return copyMap(vv), true
	case []any:
		ret := make([]any, len(vv))
		copy(ret, vv)
		return ret, true
	case nil:
		return nil, true
	default:
		panic(fmt.Sprintf("Unknown type: %#v", vv))
	}
}

func (c *Conf) GetMapStringAny(k ...string) (map[string]any, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return nil, false
	}
	r, ok := v.(map[string]any)
	return r, ok
}

func (c *Conf) GetString(k ...string) (string, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return "", false
	}
	r, ok := v.(string)
	return r, ok
}

func (c *Conf) GetDString(d string, k ...string) string {
	r, ok := c.GetString(k...)
	if ok {
		return r
	}
	return d
}

func (c *Conf) GetDSString(d string, k ...string) string {
	r, ok := c.GetString(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *Conf) GetFloat64(k ...string) (float64, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return 0.0, false
	}
	r, ok := v.(float64)
	return r, ok
}

func (c *Conf) GetDFloat64(d float64, k ...string) float64 {
	r, ok := c.GetFloat64(k...)
	if ok {
		return r
	}
	return d
}

func (c *Conf) GetDSFloat64(d float64, k ...string) float64 {
	r, ok := c.GetFloat64(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *Conf) GetInt64(k ...string) (int64, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return 0, false
	}
	r, ok := v.(int64)
	return r, ok
}

func (c *Conf) GetDInt64(d int64, k ...string) int64 {
	r, ok := c.GetInt64(k...)
	if ok {
		return r
	}
	return d
}

func (c *Conf) GetDSInt64(d int64, k ...string) int64 {
	r, ok := c.GetInt64(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *Conf) GetInt(k ...string) (int, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return 0, false
	}
	r, ok := v.(int)
	return r, ok
}

func (c *Conf) GetDInt(d int, k ...string) int {
	r, ok := c.GetInt(k...)
	if ok {
		return r
	}
	return d
}

func (c *Conf) GetDSInt(d int, k ...string) int {
	r, ok := c.GetInt(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *Conf) GetBool(k ...string) (bool, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return false, false
	}
	r, ok := v.(bool)
	return r, ok
}

func (c *Conf) GetDBool(d bool, k ...string) bool {
	r, ok := c.GetBool(k...)
	if ok {
		return r
	}
	return d
}

func (c *Conf) GetDSBool(d bool, k ...string) bool {
	r, ok := c.GetBool(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

type ConfWalkFunc func(k []string, v any)

func (c *Conf) Walk(f ConfWalkFunc) {
	c.lock.Lock()
	walkTree(c.tree, []string{}, f)
	c.lock.Unlock()
}

func (c *Conf) SubTree(path ...string) *ConfSubtree {
	return NewSubTree(c, path...)
}

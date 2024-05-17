package lac

import (
	"sync"

	"github.com/mitchellh/mapstructure"
)

type MapConf struct {
	tree map[string]any
	lock *sync.Mutex
	path []string
}

func (c *MapConf) CopyTree(t map[string]any) {
	c.lock.Lock()
	c.tree = copyMap(t)
	c.lock.Unlock()
}

func (c *MapConf) Set(v any, k ...string) {
	c.lock.Lock()
	setTree(c.tree, v, k)
	c.lock.Unlock()
}

func (c *MapConf) Get(k ...string) (any, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return nil, false
	}
	return copyAny(v), true
}

func (c *MapConf) GetToStruct(t any, k ...string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return ErrNoKey
	}
	return mapstructure.Decode(v, t)
}

func (c *MapConf) GetMapStringAny(k ...string) (map[string]any, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return nil, false
	}
	r, ok := v.(map[string]any)
	if !ok {
		return nil, false
	}
	return copyMap(r), true
}

func (c *MapConf) GetString(k ...string) (string, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return "", false
	}
	r, ok := v.(string)
	return r, ok
}

func (c *MapConf) GetDString(d string, k ...string) string {
	r, ok := c.GetString(k...)
	if ok {
		return r
	}
	return d
}

func (c *MapConf) GetDSString(d string, k ...string) string {
	r, ok := c.GetString(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *MapConf) GetFloat64(k ...string) (float64, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return 0.0, false
	}
	r, ok := v.(float64)
	return r, ok
}

func (c *MapConf) GetDFloat64(d float64, k ...string) float64 {
	r, ok := c.GetFloat64(k...)
	if ok {
		return r
	}
	return d
}

func (c *MapConf) GetDSFloat64(d float64, k ...string) float64 {
	r, ok := c.GetFloat64(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *MapConf) GetInt64(k ...string) (int64, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return 0, false
	}
	r, ok := v.(int64)
	return r, ok
}

func (c *MapConf) GetDInt64(d int64, k ...string) int64 {
	r, ok := c.GetInt64(k...)
	if ok {
		return r
	}
	return d
}

func (c *MapConf) GetDSInt64(d int64, k ...string) int64 {
	r, ok := c.GetInt64(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *MapConf) GetInt(k ...string) (int, bool) {
	i, ok := c.GetInt64(k...)
	return int(int32(i)), ok
}

func (c *MapConf) GetDInt(d int, k ...string) int {
	r, ok := c.GetInt(k...)
	if ok {
		return r
	}
	return d
}

func (c *MapConf) GetDSInt(d int, k ...string) int {
	r, ok := c.GetInt(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *MapConf) GetBool(k ...string) (bool, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, k)
	if !ok {
		return false, false
	}
	r, ok := v.(bool)
	return r, ok
}

func (c *MapConf) GetDBool(d bool, k ...string) bool {
	r, ok := c.GetBool(k...)
	if ok {
		return r
	}
	return d
}

func (c *MapConf) GetDSBool(d bool, k ...string) bool {
	r, ok := c.GetBool(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *MapConf) GetKeys(k ...string) ([]string, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	vAny, ok := lookupTree(c.tree, k)
	if !ok {
		return nil, false
	}
	v, ok := vAny.(map[string]any)
	if !ok {
		return nil, false
	}
	ret := []string{}
	for k := range v {
		ret = append(ret, k)
	}
	return ret, true
}

type ConfWalkFunc func(k []string, v any)

func (c *MapConf) Walk(f ConfWalkFunc) {
	c.lock.Lock()
	walkTree(c.tree, []string{}, f)
	c.lock.Unlock()
}

func (c *MapConf) SubTree(path ...string) Conf {
	return &MapConf{
		tree: c.tree,
		lock: c.lock,
		path: append(c.path, path...),
	}
}

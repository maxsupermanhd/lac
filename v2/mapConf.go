package lac

import (
	"encoding/json"
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
	setTree(c.tree, copyAny(v), append(c.path, k...))
	c.lock.Unlock()
}

func (c *MapConf) Get(k ...string) (any, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, append(c.path, k...))
	if !ok {
		return nil, false
	}
	return copyAny(v), true
}

func (c *MapConf) GetToStruct(t any, k ...string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, append(c.path, k...))
	if !ok {
		return ErrNoKey
	}
	return mapstructure.Decode(v, t)
}

func (c *MapConf) GetMapStringAny(k ...string) (map[string]any, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, append(c.path, k...))
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
	v, ok := lookupTree(c.tree, append(c.path, k...))
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
	v, ok := lookupTree(c.tree, append(c.path, k...))
	if !ok {
		return 0.0, false
	}
	switch r := v.(type) {
	case float64:
		return r, true
	case float32:
		return float64(r), true
	case json.Number:
		rr, err := r.Float64()
		if err == nil {
			return rr, true
		}
	}
	return 0.0, false
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
	v, ok := lookupTree(c.tree, append(c.path, k...))
	if !ok {
		return 0, false
	}
	switch r := v.(type) {
	case json.Number:
		rr, err := r.Int64()
		if err == nil {
			return rr, true
		}
	case int:
		return int64(r), true
	case int8:
		return int64(r), true
	case int16:
		return int64(r), true
	case int32:
		return int64(r), true
	case int64:
		return int64(r), true
	case uint:
		return int64(r), true
	case uint8:
		return int64(r), true
	case uint16:
		return int64(r), true
	case uint32:
		return int64(r), true
	case uint64:
		return int64(r), true
	}
	return 0, false
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
	v, ok := lookupTree(c.tree, append(c.path, k...))
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

func (c *MapConf) GetSliceAny(k ...string) ([]any, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, append(c.path, k...))
	if !ok {
		return nil, false
	}
	r, ok := v.([]any)
	if !ok {
		return nil, false
	}
	return copySliceAny(r), true
}

func (c *MapConf) GetDSliceAny(d []any, k ...string) []any {
	r, ok := c.GetSliceAny(k...)
	if ok {
		return r
	}
	return d
}

func (c *MapConf) GetDSSliceAny(d []any, k ...string) []any {
	r, ok := c.GetSliceAny(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *MapConf) GetSliceString(k ...string) ([]string, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, append(c.path, k...))
	if !ok {
		return nil, false
	}
	r, ok := v.([]any)
	if !ok {
		return nil, false
	}
	return copySlice[string](r)
}

func (c *MapConf) GetDSliceString(d []string, k ...string) []string {
	r, ok := c.GetSliceString(k...)
	if ok {
		return r
	}
	return d
}

func (c *MapConf) GetDSSliceString(d []string, k ...string) []string {
	r, ok := c.GetSliceString(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *MapConf) GetSliceInt(k ...string) ([]int, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, append(c.path, k...))
	if !ok {
		return nil, false
	}
	r, ok := v.([]any)
	if !ok {
		return nil, false
	}
	ret := make([]int, len(r))
	for i, rr := range r {
		switch vv := rr.(type) {
		case json.Number:
			vvv, err := vv.Int64()
			if err != nil {
				return nil, false
			}
			ret[i] = int(vvv)
		case int:
			ret[i] = int(vv)
		case int8:
			ret[i] = int(vv)
		case int16:
			ret[i] = int(vv)
		case int32:
			ret[i] = int(vv)
		case int64:
			ret[i] = int(vv)
		case uint:
			ret[i] = int(vv)
		case uint8:
			ret[i] = int(vv)
		case uint16:
			ret[i] = int(vv)
		case uint32:
			ret[i] = int(vv)
		case uint64:
			ret[i] = int(vv)
		default:
			return nil, false
		}
	}
	return ret, true
}

func (c *MapConf) GetDSliceInt(d []int, k ...string) []int {
	r, ok := c.GetSliceInt(k...)
	if ok {
		return r
	}
	return d
}

func (c *MapConf) GetDSSliceInt(d []int, k ...string) []int {
	r, ok := c.GetSliceInt(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *MapConf) GetSliceFloat(k ...string) ([]float64, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, ok := lookupTree(c.tree, append(c.path, k...))
	if !ok {
		return nil, false
	}
	r, ok := v.([]any)
	if !ok {
		return nil, false
	}
	ret := make([]float64, len(r))
	for i, rr := range r {
		switch vv := rr.(type) {
		case json.Number:
			vvv, err := vv.Float64()
			if err != nil {
				return nil, false
			}
			ret[i] = float64(vvv)
		case float32:
			ret[i] = float64(vv)
		case float64:
			ret[i] = vv
		default:
			return nil, false
		}
	}
	return ret, true
}

func (c *MapConf) GetDSliceFloat(d []float64, k ...string) []float64 {
	r, ok := c.GetSliceFloat(k...)
	if ok {
		return r
	}
	return d
}

func (c *MapConf) GetDSSliceFloat(d []float64, k ...string) []float64 {
	r, ok := c.GetSliceFloat(k...)
	if ok {
		return r
	}
	c.Set(d, k...)
	return d
}

func (c *MapConf) GetKeys(k ...string) ([]string, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	vAny, ok := lookupTree(c.tree, append(c.path, k...))
	if !ok {
		return nil, false
	}
	v, ok := vAny.(map[string]any)
	if !ok {
		return nil, false
	}
	ret := []string{}
	for kk := range v {
		ret = append(ret, kk)
	}
	return ret, true
}

type ConfWalkFunc func(k []string, v any)

func (c *MapConf) Walk(f ConfWalkFunc) {
	c.lock.Lock()
	walkTree(c.tree, []string{}, f)
	c.lock.Unlock()
}

func (c *MapConf) LinkSubTree(path ...string) Conf {
	return &MapConf{
		tree: c.tree,
		lock: c.lock,
		path: append(c.path, path...),
	}
}

func (c *MapConf) DupSubTree(path ...string) Conf {
	m, ok := c.GetMapStringAny(path...)
	if !ok {
		return nil
	}
	return &MapConf{
		tree: m,
		lock: &sync.Mutex{},
		path: []string{},
	}
}

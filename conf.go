package lac

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/mitchellh/mapstructure"
)

type Conf struct {
	tree      map[string]any
	listeners []ConfListener
	lock      sync.Mutex
}

func (c *Conf) SetBytesJSON(b []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return json.Unmarshal(b, &c.tree)
}

func (c *Conf) SetFileJSON(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return c.SetBytesJSON(b)
}

func (c *Conf) SetTree(t map[string]any) {
	c.lock.Lock()
	c.tree = t
	c.lock.Unlock()
}

func (c *Conf) Set(v any, k ...string) {
	c.lock.Lock()
	setTree(c.tree, v, k)
	for _, l := range c.listeners {
		if areStringSlicesEqual(l.path, k) {
			l.f(k, v)
		}
	}
	c.lock.Unlock()
}

func (c *Conf) GetToStruct(t *any, k ...string) error {
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
	return lookupTree(c.tree, k)
}

type ConfWalkFunc func(k []string, v any)

func (c *Conf) Walk(f ConfWalkFunc) {
	c.lock.Lock()
	walkTree(c.tree, []string{}, f)
	c.lock.Unlock()
}

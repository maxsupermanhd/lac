package lac

import (
	"bytes"
	"encoding/json"
	"os"
)

func FromBytesJSON(b []byte) (Conf, error) {
	c := NewConf()
	err := c.SetFromBytesJSON(b)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func FromFileJSON(path string) (Conf, error) {
	c := NewConf()
	err := c.SetFromFileJSON(path)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *MapConf) ToBytesJSON() ([]byte, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return json.Marshal(c.tree)
}

func (c *MapConf) ToBytesIndentJSON() ([]byte, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return json.MarshalIndent(c.tree, "", "\t")
}

func (c *MapConf) ToFileJSON(path string, perm os.FileMode) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	b, err := json.Marshal(c.tree)
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, perm)
}

func (c *MapConf) ToFileIndentJSON(path string, perm os.FileMode) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	b, err := json.MarshalIndent(c.tree, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, perm)
}

func (c *MapConf) SetFromBytesJSON(b []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	p, err := parseJSON(b)
	if err == nil {
		c.tree = p
	}
	return err
}

func (c *MapConf) SetFromFileJSON(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return c.SetFromBytesJSON(b)
}

func parseJSON(b []byte) (map[string]any, error) {
	var parsed map[string]interface{}
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err := d.Decode(&parsed)
	if err != nil {
		return parsed, err
	}
	walkTree(parsed, []string{}, func(k []string, v any) {
		n, ok := v.(json.Number)
		if !ok {
			return
		}
		if i, err := n.Int64(); err == nil {
			setTree(parsed, i, k)
			return
		}
		if f, err := n.Float64(); err == nil {
			setTree(parsed, f, k)
			return
		}
	})
	return parsed, nil
}

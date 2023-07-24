package lac

import (
	"encoding/json"
	"os"
)

func FromBytesJSON(b []byte) (*Conf, error) {
	c := NewConf()
	err := c.SetFromBytesJSON(b)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func FromFileJSON(path string) (*Conf, error) {
	c := NewConf()
	err := c.SetFromFileJSON(path)
	if err != nil {
		return nil, err
	}
	return c, nil
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

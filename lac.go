package lac

import (
	"errors"
)

var (
	ErrNoKey = errors.New("key not found")
)

// type ConfListenerFunc func(k []string, newval any)

// type ConfListener struct {
// 	f    ConfListenerFunc
// 	path []string
// }

func NewConf() *Conf {
	return &Conf{
		tree: map[string]any{},
		// listeners: []ConfListener{},
	}
}

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

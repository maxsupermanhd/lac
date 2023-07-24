package lac

import (
	"errors"
)

var (
	ErrNoKey = errors.New("key not found")
)

func NewConf() *Conf {
	return &Conf{
		tree: map[string]any{},
	}
}

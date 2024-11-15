package lac

import (
	"errors"
	"os"
	"sync"
)

var (
	ErrNoKey = errors.New("key not found")
)

type Conf interface {
	CopyTree(t map[string]any)

	Set(v any, k ...string)
	Get(k ...string) (any, bool)
	GetToStruct(t any, k ...string) error
	GetMapStringAny(k ...string) (map[string]any, bool)

	GetString(k ...string) (string, bool)
	GetDString(d string, k ...string) string
	GetDSString(d string, k ...string) string

	GetFloat64(k ...string) (float64, bool)
	GetDFloat64(d float64, k ...string) float64
	GetDSFloat64(d float64, k ...string) float64

	GetInt64(k ...string) (int64, bool)
	GetDInt64(d int64, k ...string) int64
	GetDSInt64(d int64, k ...string) int64

	GetInt(k ...string) (int, bool)
	GetDInt(d int, k ...string) int
	GetDSInt(d int, k ...string) int

	GetBool(k ...string) (bool, bool)
	GetDBool(d bool, k ...string) bool
	GetDSBool(d bool, k ...string) bool

	GetSliceAny(k ...string) ([]any, bool)
	GetDSliceAny(d []any, k ...string) []any
	GetDSSliceAny(d []any, k ...string) []any

	GetSliceString(k ...string) ([]string, bool)
	GetDSliceString(d []string, k ...string) []string
	GetDSSliceString(d []string, k ...string) []string

	GetSliceInt(k ...string) ([]int, bool)
	GetDSliceInt(d []int, k ...string) []int
	GetDSSliceInt(d []int, k ...string) []int

	GetSliceFloat(k ...string) ([]float64, bool)
	GetDSliceFloat(d []float64, k ...string) []float64
	GetDSSliceFloat(d []float64, k ...string) []float64

	GetKeys(k ...string) ([]string, bool)
	Walk(f ConfWalkFunc)

	LinkSubTree(path ...string) Conf
	DupSubTree(path ...string) Conf

	ToBytesJSON() ([]byte, error)
	MarshalJSON() ([]byte, error)
	ToBytesIndentJSON() ([]byte, error)
	ToFileJSON(path string, perm os.FileMode) error
	ToFileIndentJSON(path string, perm os.FileMode) error
	SetFromBytesJSON(b []byte) error
	SetFromFileJSON(path string) error
}

func NewConf() Conf {
	return &MapConf{
		tree: map[string]any{},
		lock: &sync.Mutex{},
		path: []string{},
	}
}

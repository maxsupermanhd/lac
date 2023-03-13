package lac

type ConfSubtree struct {
	root *Conf
	path []string
}

func NewSubTree(c *Conf, path ...string) *ConfSubtree {
	return &ConfSubtree{
		root: c,
		path: path,
	}
}

func (c *ConfSubtree) Set(v any, k ...string) {
	c.root.Set(v, append(c.path, k...)...)
}

func (c *ConfSubtree) Get(k ...string) (any, bool) {
	return c.root.Get(append(c.path, k...)...)
}

func (c *ConfSubtree) GetToStruct(t any, k ...string) error {
	return c.root.GetToStruct(t, append(c.path, k...)...)
}

func (c *ConfSubtree) GetString(k ...string) (string, bool) {
	return c.root.GetString(append(c.path, k...)...)
}

func (c *ConfSubtree) GetDString(d string, k ...string) string {
	return c.root.GetDString(d, append(c.path, k...)...)
}

func (c *ConfSubtree) GetDSString(d string, k ...string) string {
	return c.root.GetDSString(d, append(c.path, k...)...)
}

func (c *ConfSubtree) GetFloat64(k ...string) (float64, bool) {
	return c.root.GetFloat64(append(c.path, k...)...)
}

func (c *ConfSubtree) GetDFloat64(d float64, k ...string) float64 {
	return c.root.GetDFloat64(d, append(c.path, k...)...)
}

func (c *ConfSubtree) GetDSFloat64(d float64, k ...string) float64 {
	return c.root.GetDSFloat64(d, append(c.path, k...)...)
}

func (c *ConfSubtree) GetInt64(k ...string) (int64, bool) {
	return c.root.GetInt64(append(c.path, k...)...)
}

func (c *ConfSubtree) GetDInt64(d int64, k ...string) int64 {
	return c.root.GetDInt64(d, append(c.path, k...)...)
}

func (c *ConfSubtree) GetDSInt64(d int64, k ...string) int64 {
	return c.root.GetDSInt64(d, append(c.path, k...)...)
}

func (c *ConfSubtree) GetInt(k ...string) (int, bool) {
	return c.root.GetInt(append(c.path, k...)...)
}

func (c *ConfSubtree) GetDInt(d int, k ...string) int {
	return c.root.GetDInt(d, append(c.path, k...)...)
}

func (c *ConfSubtree) GetDSInt(d int, k ...string) int {
	return c.root.GetDSInt(d, append(c.path, k...)...)
}

func (c *ConfSubtree) GetBool(k ...string) (bool, bool) {
	return c.root.GetBool(append(c.path, k...)...)
}

func (c *ConfSubtree) GetDBool(d bool, k ...string) bool {
	return c.root.GetDBool(d, append(c.path, k...)...)
}

func (c *ConfSubtree) GetDSBool(d bool, k ...string) bool {
	return c.root.GetDSBool(d, append(c.path, k...)...)
}

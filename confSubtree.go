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

func (c *ConfSubtree) GetToStruct(t *any, k ...string) error {
	return c.root.GetToStruct(t, append(c.path, k...)...)
}

func (c *ConfSubtree) GetString(k ...string) (string, bool) {
	return c.root.GetString(append(c.path, k...)...)
}

func (c *ConfSubtree) GetDString(d string, k ...string) string {
	r, ok := c.GetString(k...)
	if ok {
		return r
	}
	return d
}

func (c *ConfSubtree) GetFloat64(k ...string) (float64, bool) {
	return c.root.GetFloat64(append(c.path, k...)...)
}

func (c *ConfSubtree) GetDFloat64(d float64, k ...string) float64 {
	r, ok := c.GetFloat64(k...)
	if ok {
		return r
	}
	return d
}

func (c *ConfSubtree) GetInt64(k ...string) (int64, bool) {
	return c.root.GetInt64(append(c.path, k...)...)
}

func (c *ConfSubtree) GetDInt64(d int64, k ...string) int64 {
	r, ok := c.GetInt64(k...)
	if ok {
		return r
	}
	return d
}

func (c *ConfSubtree) GetInt(k ...string) (int, bool) {
	return c.root.GetInt(append(c.path, k...)...)
}

func (c *ConfSubtree) GetDInt(d int, k ...string) int {
	r, ok := c.GetInt(k...)
	if ok {
		return r
	}
	return d
}

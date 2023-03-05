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

package lac

import (
	"fmt"
	"sort"
)

func lookupTree(tree map[string]any, k []string) (any, bool) {
	if len(k) == 0 {
		return nil, false
	}
	v, ok := tree[k[0]]
	if ok {
		if len(k) > 1 {
			m, ok := v.(map[string]any)
			if ok {
				return lookupTree(m, k[1:])
			}
		} else {
			return v, true
		}
	}
	return nil, false
}

func setTree(tree map[string]any, v any, k []string) {
	if len(k) == 1 {
		tree[k[0]] = v
		return
	}
	n, ok := tree[k[0]]
	if !ok {
		newmap := map[string]any{}
		tree[k[0]] = newmap
		setTree(newmap, v, k[1:])
		return
	}
	t, ok := n.(map[string]any)
	if !ok {
		newmap := map[string]any{}
		tree[k[0]] = newmap
		setTree(newmap, v, k[1:])
		return
	}
	setTree(t, v, k[1:])
}

func walkTree(tree map[string]any, passed []string, fn ConfWalkFunc) {
	keys := []string{}
	for k := range tree {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := tree[k]
		p := append(passed, k)
		m, ok := v.(map[string]any)
		if ok {
			walkTree(m, p, fn)
			continue
		}
		fn(p, v)
	}
}

// func areStringSlicesEqual(a, b []string) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	for i := range a {
// 		if a[i] != b[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func copyMap(m map[string]any) map[string]any {
	ret := map[string]any{}
	for k, v := range m {
		switch vv := v.(type) {
		case string, float64:
			ret[k] = vv
		case map[string]any:
			ret[k] = copyMap(vv)
		case []any:
			a := make([]any, len(vv))
			copy(a, vv)
		case nil:
			ret[k] = nil
		default:
			panic(fmt.Sprintf("Unknown type: %#v", vv))
		}
	}
	return ret
}

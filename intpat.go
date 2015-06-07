// Package intpat is an integer patricia tree
/*
Based on
    https://github.com/liuxinyu95/AlgoXY/blob/algoxy/datastruct/tree/trie/src/intpatricia.py

Licensed under the GPL, like the original.
*/
package intpat

type Key uint32

type Tree struct {
	key         Key
	value       interface{}
	prefix      Key
	hasPrefix   bool
	mask        Key
	left, right *Tree
}

func (t *Tree) setChildren(left, right *Tree) {
	t.left = left
	t.right = right
}

func (t *Tree) replaceChild(x, y *Tree) {
	if t.left == x {
		t.left = y
	} else {
		t.right = y
	}
}

func (t *Tree) isLeaf() bool {
	return t.left == nil && t.right == nil
}

func (t *Tree) getPrefix() Key {
	if t.hasPrefix {
		return t.prefix
	}
	return t.key
}

func maskbit(x, mask Key) Key {
	return x &^ mask
}

func match(key Key, t *Tree) bool {
	return !t.isLeaf() && (maskbit(key, t.mask) == t.prefix)
}

func zero(x, mask Key) bool {
	return x&((mask>>1)+1) == 0
}

func lcp(p1, p2 Key) (mask, prefix Key) {
	diff := (p1 ^ p2)
	mask = Key(1)
	for diff != 0 {
		diff >>= 1
		mask <<= 1
	}
	mask--
	return maskbit(p1, mask), mask
}

func branch(t1, t2 *Tree) *Tree {
	t := &Tree{}
	t.prefix, t.mask = lcp(t1.getPrefix(), t2.getPrefix())
	t.hasPrefix = true
	if zero(t1.getPrefix(), t.mask) {
		t.setChildren(t1, t2)
	} else {
		t.setChildren(t2, t1)
	}
	return t
}

func (t *Tree) Insert(key Key, value interface{}) *Tree {
	if t == nil {
		return &Tree{key: key, value: value}
	}

	node := t
	var parent *Tree

	for {
		if match(key, node) {
			parent = node
			if zero(key, node.mask) {
				node = node.left
			} else {
				node = node.right
			}
		} else {
			if node.isLeaf() && key == node.key {
				node.value = value
			} else {
				new_node := branch(node, &Tree{key: key, value: value})
				if parent == nil {
					t = new_node
				} else {
					parent.replaceChild(node, new_node)
				}
			}
			break
		}
	}
	return t
}

func (t *Tree) Lookup(key Key) (value interface{}, ok bool) {
	if t == nil {
		return nil, false
	}
	for !t.isLeaf() && match(key, t) {
		if zero(key, t.mask) {
			t = t.left
		} else {
			t = t.right
		}
	}
	if t.isLeaf() && t.key == key {
		return t.value, true
	}

	return nil, false
}

func (t *Tree) Prefix(key Key) (p Key) {
	if t == nil {
		return 0
	}
	for match(key, t) {
		if zero(key, t.mask) {
			t = t.left
		} else {
			t = t.right
		}
	}
	return t.getPrefix()
}

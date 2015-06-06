package intpat

import (
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {

	var inserts = []struct {
		k Key
		v interface{}
	}{
		{6, nil},
		{7, nil},
		{1, "x"},
		{4, "y"},
		{5, "z"},
	}

	var tree *Tree

	if v, ok := tree.Lookup(0); ok {
		t.Errorf("<nil>.Lookup(0)=%v,%v, want nil,false", v, ok)
	}

	for _, tt := range inserts {
		tree = tree.Insert(tt.k, tt.v)
	}

	for _, tt := range inserts {
		if v, ok := tree.Lookup(tt.k); !ok || v != tt.v {
			t.Errorf("Lookup(%v)=%v,%v, want %v", tt.k, v, ok, tt.v)
		}
	}

	if v, ok := tree.Lookup(0); ok {
		t.Errorf("Lookup(0)=%v,%v, want nil, false", v, ok)
	}

	val := "xx"
	tree = tree.Insert(1, val)
	if v, ok := tree.Lookup(1); !ok || v != val {
		t.Errorf("Lookup(%v)=%v,%v, want %v,true", 1, v, ok, val)
	}
}

func TestLots(t *testing.T) {

	var tree *Tree

	m := make(map[Key]int)

	for i := 0; i < 1e5; i++ {
		k := Key(rand.Int63())
		m[k] = i
		tree = tree.Insert(k, i)
	}

	for mk, mv := range m {
		if v, ok := tree.Lookup(mk); !ok || v != mv {
			t.Errorf("Lookup(%v)=%v,%v, want %v", mk, v, ok, mv)
		}
	}
}

func BenchmarkMap(b *testing.B) {

	m := make(map[Key]int)

	rand.Seed(0)

	k := Key(rand.Int63())
	m[k] = 42

	for i := 0; i < 1e5; i++ {
		k := Key(rand.Int63())
		m[k] = i
	}

	b.ResetTimer()

	var total int

	for i := 0; i < b.N; i++ {
		_, ok := m[k]
		if ok {
			total++
		}
	}
}

func BenchmarkTree(b *testing.B) {

	var tree *Tree

	rand.Seed(0)

	k := Key(rand.Int63())

	for i := 0; i < 1e5; i++ {
		k := Key(rand.Int63())
		tree = tree.Insert(k, i)
	}

	b.ResetTimer()

	var total int

	for i := 0; i < b.N; i++ {
		_, ok := tree.Lookup(k)
		if ok {
			total++
		}
	}
}

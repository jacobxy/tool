package tool

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}]bool
}

func (set *HashSet) Contain(e interface{}) bool {
	return set.m[e]
}

func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}

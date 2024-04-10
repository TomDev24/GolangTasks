package main

import (
	"errors"
	"fmt"
)

type Set struct {
	dict map[string]bool
}

func NewSet() *Set {
	return &Set{ map[string]bool{} }
}

func (set *Set) Init (elements []string) {
	for _, el := range elements {
		set.Add(el)
	}
}

func (set *Set) Exist (element string) bool {
	_, exist := set.dict[element]
	return exist
}

func (set *Set) Add (element string) error {
	if _, exist := set.dict[element]; exist {
		return errors.New("Element already present in Set")
	}
	set.dict[element] = true
	return nil
}

func (set *Set) Delete (element string) {
	delete(set.dict, element)
}

func (set *Set) Elements() []string {
	res := []string{}
	for k := range set.dict {
		res = append(res, k)
	}
	return res
}

func (set *Set) Copy(otherSet *Set) {
	elements := otherSet.Elements()
	set.Init(elements)
}

func (set *Set) Len() int {
	return len(set.dict)
}

func main() {
	s := NewSet()
	s.Init([]string{"cat", "cat", "dog", "cat", "tree"})

	fmt.Println(s.Elements())
}

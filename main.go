package main

import (
	"fmt"
	"strconv"
)

const (
	Prime              = (1 << 31) - 1
	a                  = 5
	DefaultHashsetSize = 16
)

type Chain []string

type Hashset struct {
	size     uint32
	elements []Chain
	len      uint32
}

func New() Hashset {
	return Hashset{
		size:     DefaultHashsetSize,
		elements: make([]Chain, DefaultHashsetSize),
		len:      uint32(0),
	}
}

func Copy(h Hashset) Hashset {
	return Hashset{
		size:     h.size,
		elements: h.elements,
		len:      h.len,
	}
}

func (c Chain) contains(target string) bool {
	for _, el := range c {
		if el == target {
			return true
		}
	}
	return false
}

func (c Chain) remove(target string) []string {
	for i, el := range c {
		if el == target {
			return append(c[:i], c[i+1:]...)
		}
	}
	return c
}

func (h *Hashset) Insert(element string) {
	if h.Has(element) {
		return
	}

	index := h.hash(element)
	h.elements[index] = append(h.elements[index], element)

	h.len++

	if h.len > h.size {
		h.size = h.size * 2
		newElements := make([]Chain, h.size)
		for _, el := range h.flatten() {
			index := h.hash(el)
			newElements[index] = append(newElements[index], el)
		}
		h.elements = newElements
	}
}

func (h *Hashset) Remove(element string) {
	if !h.Has(element) {
		return
	}

	index := h.hash(element)
	h.elements[index] = h.elements[index].remove(element)

	h.len -= 1

	if h.len < h.size/4 && h.size/4 >= DefaultHashsetSize {
		h.size = h.size / 4
		newElements := make([]Chain, h.size)
		for _, el := range h.flatten() {
			index := h.hash(el)
			newElements[index] = append(newElements[index], el)
		}
		h.elements = newElements
	}
}

func (h Hashset) Has(element string) bool {
	index := h.hash(element)
	return h.elements[index].contains(element)
}

func (x Hashset) Intersect(y Hashset) Hashset {
	var intersection = New()

	for _, el := range y.flatten() {
		if x.Has(el) {
			intersection.Insert(el)
		}
	}

	return intersection
}

func (x Hashset) Union(y Hashset) Hashset {
	union := Copy(x)

	for _, el := range y.flatten() {
		union.Insert(el)
	}

	return union
}

func (x Hashset) Superset(y Hashset) bool {
	intersection := x.Intersect(y)
	if(intersection.len != y.len) {return false}
	for _, el := range(intersection.flatten()){
		if(!y.Has(el)){return false}
	}
	return true
}

func (x Hashset) Subset(y Hashset) bool {
	return y.Superset(x)
}

func (h Hashset) Println() {
	fmt.Println(h.flatten())
}

func (h Hashset) flatten() []string {
	var flatArray []string
	for _, el := range h.elements {
		flatArray = append(flatArray, el...)
	}
	return flatArray
}

func (h Hashset) hash(element string) uint32 {
	hash := uint32(0)

	for _, char := range element {
		hash = (hash*uint32(a) + uint32(char)) % Prime
	}
	return hash % h.size
}

func main() {
	hashset := New()

	for i := 0; i < 63; i++ {
		el := strconv.Itoa(i)
		hashset.Insert(el)
	}
	for i := 0; i < 32+17; i++ {
		el := strconv.Itoa(i)
		hashset.Remove(el)
	}

	hashset.Println()
	fmt.Println("hashset len:", hashset.len)
	fmt.Println("hashset size:", hashset.size)
}

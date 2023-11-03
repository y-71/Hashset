package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Prime              = (1 << 31) - 1
	a                  = 5
	InitialHashsetSize = 16
)

type Chain []string

type Hashset struct {
	size     uint32
	elements []Chain
	len      uint32
}

func New() Hashset {
	return Hashset{
		size:     InitialHashsetSize,
		elements: make([]Chain, InitialHashsetSize),
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

func (h Hashset) hash(element string) uint32 {
	hash := uint32(0)

	for _, char := range element {
		hash = (hash*uint32(a) + uint32(char)) % Prime
	}
	return hash % h.size
}

func (h Hashset) Add(element string) {
	if h.Has(element) {
		return
	}
	h.len += 1
	if h.len < h.size {
		index := h.hash(element)
		if len(h.elements[index]) == 0 {
			h.elements[index] = []string{element}
		} else {
			h.elements[index] = append(h.elements[index], element)
		}
	} else {
		h.size = h.size * 2
		newElements := make([]Chain, h.size)
		for _, el := range(h.flatten()){
			index := h.hash(el)
			if len(newElements[index]) == 0 {
				newElements[index] = []string{el}
			} else {
				newElements[index] = append(newElements[index], el)
			}
		}
	}
}

func (h Hashset)flatten() []string{
	output := fmt.Sprint(h.elements)
	output = strings.ReplaceAll(output, "[", "")
	output = strings.ReplaceAll(output, "]", "")
	return strings.Split(output, " ")
}

func (h Hashset) Println() {
	fmt.Println(h.flatten())
}

func (h Hashset) Has(element string) bool {
	index := h.hash(element)
	return h.elements[index].contains(element)
}

func main() {
	hashset := New()

	for i := 0; i < 32; i++ {
		el := strconv.Itoa(i)
		fmt.Println(hashset.hash(el))
		hashset.Add(el)
	}
	hashset.Println()
}

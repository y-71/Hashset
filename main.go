package main

import (
	"fmt"
)

const (
	Prime            = (1 << 31) - 1
	a                = 5
	InitialHashsetSize = 16
)

type Hashset struct {
	size     uint32
	elements []string
}

func CreateHashset() Hashset {
	return Hashset{
		size:     InitialHashsetSize,
		elements: make([]string, InitialHashsetSize),
	}
}

func (h Hashset) hash(element string) uint32 {
	hash := uint32(0)

	for _, char := range element {
		hash = (hash*uint32(a) + uint32(char)) % Prime
	}
	return hash % h.size
}

func (h Hashset) add(element string){
	index := h.hash(element)
	h.elements[index] = element
}

func (h Hashset) has(element string)bool{
	index := h.hash(element)
	return h.elements[index] == element
}

func main() {
	hashset := CreateHashset()
	myHash := hashset.hash("test")
	fmt.Println(myHash)
}

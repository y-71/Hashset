package main

import "fmt"

const Prime = (1 << 31) - 1
const a = 5
const InitialHashsetSize = 16

type Hashset struct {
	size     uint32
	elements []string
}

func create_hashset() Hashset {
	return Hashset{
		size: InitialHashsetSize,
		elements: make([]string, InitialHashsetSize),
	}
}

func (h Hashset) hash(element string) uint32 {
	hash := uint32(0)

	for _, char := range element {
		hash = (uint32(hash)*uint32(a) + uint32(char)) % Prime
	}
	return hash % h.size
}

func main() {
	hashset := create_hashset()
	my_hash := hashset.hash("test")
	fmt.Println(my_hash)
}

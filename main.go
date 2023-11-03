
package main

import "fmt"

const (
	Prime            = (1 << 31) - 1
	a                = 5
	InitialHashsetSize = 16
)

type Hashset struct {
	size     uint32
	elements []string
}

func createHashset() Hashset {
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

func main() {
	hashset := createHashset()
	myHash := hashset.hash("test")
	fmt.Println(myHash)
}

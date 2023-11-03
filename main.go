package main

import "fmt"

func hash(element string, a, p int, m uint32)uint32{
	hash := uint32(0)

	for _, char := range(element){ 
		hash = (uint32(hash)*uint32(a) + uint32(char)) % uint32(p)
	}
	return hash % m;
}



const Prime = (2<<61)-1 
const InitialHashsetSize = 16

func main(){
	m := uint32(InitialHashsetSize)
	my_hash := hash("test", 5, Prime, m)
	fmt.Println(my_hash)
}

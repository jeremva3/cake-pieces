package main

import (
	"cake-pieces/encrypt"
	"cake-pieces/graph"
	"fmt"
)

func main() {

	key2 := "dcj"
	message2 := "I love prOgrAmming!"
	fmt.Println(encrypt.EncryptComunication(&key2, &message2))
	// fmt.Println(encrypt.EncryptComunication(nil, &message2))
	// fmt.Println(encrypt.EncryptComunication(&key2, nil))

	input := []int{3, 4, -7, 5, -6, 2, 5, -1, 8}
	result := graph.RemoveZeroSumConsecutive(input)
	fmt.Println(result)
}

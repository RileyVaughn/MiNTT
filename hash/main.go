package main

import (
	"fmt"
	"math/rand"
	"strconv"

	m16 "github.com/RileyVaughn/MiNTT/hash/MiNNT16"
)

func main() {

	seed, _ := strconv.Atoi("MiNNT")
	rand.Seed(int64(seed))

	var m16_input [3264]byte
	for i := 0; i < 3264; i++ {
		m16_input[i] = byte(rand.Intn(256))
	}
	//fmt.Println(m16_input)

	out := m16.MinNNT16(m16_input)
	fmt.Println(out)
}

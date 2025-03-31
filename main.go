package main

import (
	"fmt"
	"math/rand"
	"strconv"

	m8 "github.com/RileyVaughn/MiNTT/hash/MiNTT8"
)

func main() {
	RunM8()
}

func RunM8() {

	seed, _ := strconv.Atoi("MiNTT")
	rand.Seed(int64(seed))
	var input [1728]byte
	for i := 0; i < 1728; i++ {
		input[i] = byte(rand.Intn(256))
	}

	m8.SetupM8()

	out := m8.MinNTT8(input)
	fmt.Println(out)

}

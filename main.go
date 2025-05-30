package main

import (
	"fmt"
	"math/rand"
	"strconv"

	m128 "github.com/RileyVaughn/MiNTT/hash/int64/simd/MiNTT128"
	m64 "github.com/RileyVaughn/MiNTT/hash/int64/simd/MiNTT64"
)

func main() {
	// RunM128()
	// RunM64()

}

func RunM64() {
	seed, _ := strconv.Atoi("MiNTT")
	rand.Seed(int64(seed))
	var input [1728]byte
	for i := 0; i < 1728; i++ {
		input[i] = byte(rand.Intn(256))
	}

	m64.SetupM64()

	out := m64.MinNTT64(input)
	fmt.Println(out)
}

func RunM128() {

	seed, _ := strconv.Atoi("MiNTT")
	rand.Seed(int64(seed))
	var input [1728]byte
	for i := 0; i < 1728; i++ {
		input[i] = byte(rand.Intn(256))
	}

	m128.SetupM128()
	out := m128.MinNTT128(input)
	fmt.Println(out)

}

package main

import (
	"fmt"
	"math/rand"
	"strconv"

	m128_16 "github.com/RileyVaughn/MiNTT/hash/int16/simd/MiNTT128"
	m128_64 "github.com/RileyVaughn/MiNTT/hash/int64/simd/MiNTT128"
	m64 "github.com/RileyVaughn/MiNTT/hash/int64/simd/MiNTT64"
)

func main() {
	RunM128()
	// RunM64()

	// var vec1 [8]int16 = [8]int16{1, 2, 3, 4, 5, 6, 7, 8}
	// var vec2 [8]int16 = [8]int16{9, 10, 11, 12, 13, 14, 15, 16}
	// product := util.SIMD_Mult(&vec1, &vec2)
	// fmt.Println(product)

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

	m128_16.SetupM128()
	m128_64.SetupM128()
	m128_16.MinNTT128(input)
	m128_64.MinNTT128(input)
	//fmt.Println(out)

}

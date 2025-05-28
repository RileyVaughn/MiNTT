package main

import (
	"fmt"
	"math/rand"
	"strconv"

	m128 "github.com/RileyVaughn/MiNTT/hash/int64/MiNTT128"
	m64 "github.com/RileyVaughn/MiNTT/hash/int64/MiNTT64"
	m8 "github.com/RileyVaughn/MiNTT/hash/int64/MiNTT8"
	"github.com/RileyVaughn/MiNTT/hash/int64/util"
)

func main() {
	// RunM128()
	// RunM64()
	//RunM8()

	var vec [8]int64 = [8]int64{1, 2, 3, 4, 5, 6, 7, 8}
	util.SIMD_Shift(&vec, 2)
	fmt.Println(vec)
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

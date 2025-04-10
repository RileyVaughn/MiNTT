package main

import (
	"fmt"
	"math/rand"
	"strconv"

	m64 "github.com/RileyVaughn/MiNTT/hash/MiNTT64"
	m8 "github.com/RileyVaughn/MiNTT/hash/MiNTT8"
)

func main() {
	//RunM64()
	m8.SetupM8()
	if m8.NTT8_TABLE != m8.NTT8_TABLE_B {
		fmt.Println(m8.NTT8_TABLE[64])
		fmt.Println(m8.NTT8_TABLE_B[64])
	} else {
		fmt.Println("All good")
	}

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

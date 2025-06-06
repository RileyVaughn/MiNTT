package main

import (
	"fmt"
	"math/rand"
	"strconv"

	m128_16 "github.com/RileyVaughn/MiNTT/hash/int16/simd/MiNTT128"
	"github.com/RileyVaughn/MiNTT/hash/int16/util"
	m128_64 "github.com/RileyVaughn/MiNTT/hash/int64/simd/MiNTT128"
	m64 "github.com/RileyVaughn/MiNTT/hash/int64/simd/MiNTT64"
)

func main() {
	//RunM128()
	// RunM64()

	N := 100000
	for i := 0; i < N; i++ {
		var init [8]int16
		var init2 [8]int16
		var init3 [8]int16
		for j := 0; j < 8; j++ {
			init[j] = int16(rand.Intn(65536) - 32768)
			init2[j] = init[j]
			init3[j] = util.Mod(init[j], 257)
			util.Mod_257(&init[j])

		}
		util.SIMD_Mod_257(&init2)

		if init != init2 || init2 != init3 {
			fmt.Println(init, init2, init3)
		}
		fmt.Println(init, init2, init3)
	}

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

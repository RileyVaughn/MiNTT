package main

import (
	// m128_norm_int64 "github.com/RileyVaughn/MiNTT/hash/int64/normal/MiNTT128"
	"fmt"
	"math/rand"

	m64_norm_int64 "github.com/RileyVaughn/MiNTT/hash/int64/normal/MiNTT64"
)

// m128_simd_int64 "github.com/RileyVaughn/MiNTT/hash/int64/simd/MiNTT128"
// m64_simd_int64 "github.com/RileyVaughn/MiNTT/hash/int64/simd/MiNTT64"

// m128_simd_int16 "github.com/RileyVaughn/MiNTT/hash/int16/simd/MiNTT128"
// m64_simd_int16 "github.com/RileyVaughn/MiNTT/hash/int16/simd/MiNTT64"
// m8_simd_int16 "github.com/RileyVaughn/MiNTT/hash/int16/simd/MiNTT8"

// m128_norm_int16 "github.com/RileyVaughn/MiNTT/hash/int16/normal/MiNTT128"
// m64_norm_int16 "github.com/RileyVaughn/MiNTT/hash/int16/normal/MiNTT64"
// m8_norm_int16 "github.com/RileyVaughn/MiNTT/hash/int16/normal/MiNTT8"

func main() {

	Setup()

	var input [1728]byte

	for i := 0; i < 1728; i++ {
		input[i] = byte(rand.Intn(256))
	}

	out := m64_norm_int64.MinNTT64(input)
	fmt.Println(out)
}

// // Sets up all hash variants
func Setup() {

	// 	m8_norm_int16.SetupM8()
	// 	m64_norm_int16.SetupM64()
	// 	m128_norm_int16.SetupM128()

	// 	m8_simd_int16.SetupM8()
	// 	m64_simd_int16.SetupM64()
	// 	m128_simd_int16.SetupM128()

	// 	m128_norm_int64.SetupM128()
	m64_norm_int64.SetupM64()

	// 	m128_simd_int64.SetupM128()
	// 	m64_simd_int64.SetupM64()

}

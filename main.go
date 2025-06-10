package main

import (
	"fmt"
	"math/rand"

	m128_norm_int64 "github.com/RileyVaughn/MiNTT/hash/int64/normal/MiNTT128"
	m64_norm_int64 "github.com/RileyVaughn/MiNTT/hash/int64/normal/MiNTT64"

	m128_simd_int64 "github.com/RileyVaughn/MiNTT/hash/int64/simd/MiNTT128"
	m64_simd_int64 "github.com/RileyVaughn/MiNTT/hash/int64/simd/MiNTT64"

	m128_simd_int16 "github.com/RileyVaughn/MiNTT/hash/int16/simd/MiNTT128"
	m64_simd_int16 "github.com/RileyVaughn/MiNTT/hash/int16/simd/MiNTT64"
	m8_simd_int16 "github.com/RileyVaughn/MiNTT/hash/int16/simd/MiNTT8"

	m128_norm_int16 "github.com/RileyVaughn/MiNTT/hash/int16/normal/MiNTT128"
	m64_norm_int16 "github.com/RileyVaughn/MiNTT/hash/int16/normal/MiNTT64"
	m8_norm_int16 "github.com/RileyVaughn/MiNTT/hash/int16/normal/MiNTT8"
	"github.com/RileyVaughn/MiNTT/hash/int64/util"
)

func main() {

	N := 1000
	for i := 0; i < N; i++ {
		var init [8]int64
		var init2 [8]int64
		var init3 [8]int64

		for j := 0; j < 8; j++ {
			init[j] = rand.Int63()
			init2[j] = init[j]
			init3[j] = init[j]
			init2[j] = util.Mod(init2[j], 257)

			util.Mod_257(&init[j])

		}
		util.SIMD_Mod_257(&init3)

		if init != init2 || init != init3 {
			fmt.Println(init, init2, init3)
		}

	}

}

// Sets up all hash variants
func Setup() {

	m8_norm_int16.SetupM8()
	m64_norm_int16.SetupM64()
	m128_norm_int16.SetupM128()

	m8_simd_int16.SetupM8()
	m64_simd_int16.SetupM64()
	m128_simd_int16.SetupM128()

	m128_norm_int64.SetupM128()
	m64_norm_int64.SetupM64()

	m128_simd_int64.SetupM128()
	m64_simd_int64.SetupM64()

}

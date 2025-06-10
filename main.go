package main

import (
	"fmt"
	"math/rand"

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

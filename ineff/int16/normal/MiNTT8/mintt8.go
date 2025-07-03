package MiNTT8

import (
	util "github.com/RileyVaughn/MiNTT/hash/int16/util"
)

func MiNTT8(input [ndiv8 * m]byte) [864]byte {

	return ChangeBase(ntt_sum(input))

}

func ntt_sum(input [ndiv8 * m]byte) [d][8]int16 {

	var solution [d][8]int16
	for i := int16(0); i < m; i++ {
		x := NTT8_TABLE[input[i]]
		for j := int16(0); j < d; j++ {
			util.Fake_SIMD_Add_Mult(&solution[j], &x, &A[i][j][0])
		}

	}

	return solution
}

// Assumes MOD has not yet occured
// Changes base from  257 to 256, moves the extra bits to the end (<first N*8 256 bits>N + <N end bits> ndiv8*d)
func ChangeBase(val [d][8]int16) [OUT_SIZE]byte {

	var output [OUT_SIZE]byte

	for i := int16(0); i < d; i++ {
		util.Fake_SIMD_Mod(&val[i])
		for k := int16(0); k < 8; k++ {
			output[i*n+k] = byte(val[i][k])
			val[i][k] = val[i][k] >> 8
			output[N+i] = output[N+i] | byte(val[i][k]>>k)
		}

	}

	return output
}

package MiNTT8

import (
	util "github.com/RileyVaughn/MiNTT/hash/int16/util"
)

func MinNTT8(input [ndiv8 * m]byte) [864]byte {

	return ChangeBase(ntt_sum(input))

}

func ntt_sum(input [ndiv8 * m]byte) [N]int16 {

	var solution [d][8]int16
	for i := int16(0); i < m; i++ {
		x := NTT8_TABLE[input[i]]
		for j := int16(0); j < d; j++ {
			//solution[n*j+k] = solution[n*j+k] + x[k]*A[i][n*j+k]
			util.SIMD_Add_Mult(&solution[j], &x, &A[i][j][0])
		}

	}

	// This needs to happen but is slow
	var out [N]int16
	for i := int16(0); i < d; i++ {
		for j := int16(0); j < 8; j++ {
			out[i] = util.Mod(solution[i][j], q)
		}

	}

	return out
}

// Assume vaules haveaady been ModQ'd
func ChangeBase(val [N]int16) [864]byte {

	var output [864]byte

	for i := int16(0); i < N; i++ {
		output[i] = byte(val[i])
		val[i] = val[i] >> 8
	}

	for i := int16(0); i < Ndiv8; i++ {
		for k := int16(0); k < 8; k++ {
			output[N+i] = output[N+i] | byte(val[8*i+k]>>k)
		}
	}

	return output
}

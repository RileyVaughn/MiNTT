package MiNTT64

import (
	util "github.com/RileyVaughn/MiNTT/hash/int64/util"
)

func MiNTT64(input [ndiv8 * m]byte) [OUT_SIZE]byte {

	return ChangeBase(ntt_sum(input))

}

func ncc(input [ndiv8]byte) [ndiv8][8]int64 {

	var intermed [ndiv8][8]int64

	for i := int64(0); i < ndiv8; i++ {
		intermed[i] = util.Fake_SIMD_Mult(NTT8_TABLE[input[i]], MULT_TABLE[i])

	}

	util.SIMD_AddSub(&intermed[0], &intermed[1])
	util.SIMD_AddSub(&intermed[2], &intermed[3])
	util.SIMD_AddSub(&intermed[4], &intermed[5])
	util.SIMD_AddSub(&intermed[6], &intermed[7])

	util.SIMD_Shift(&intermed[3], 4)
	util.SIMD_Shift(&intermed[7], 4)

	util.SIMD_AddSub(&intermed[0], &intermed[2])
	util.SIMD_AddSub(&intermed[1], &intermed[3])
	util.SIMD_AddSub(&intermed[4], &intermed[6])
	util.SIMD_AddSub(&intermed[5], &intermed[7])

	util.SIMD_Shift(&intermed[5], 2)
	util.SIMD_Shift(&intermed[6], 4)
	util.SIMD_Shift(&intermed[7], 6)

	util.SIMD_AddSub(&intermed[0], &intermed[4])
	util.SIMD_AddSub(&intermed[1], &intermed[5])
	util.SIMD_AddSub(&intermed[2], &intermed[6])
	util.SIMD_AddSub(&intermed[3], &intermed[7])

	return intermed
}

func ntt_sum(input [ndiv8 * m]byte) [d][ndiv8][8]int64 {

	var solution [d][ndiv8][8]int64
	for i := int64(0); i < m; i++ {
		x := ncc(sepInput(input, i))
		for j := int64(0); j < d; j++ {
			for k := int64(0); k < ndiv8; k++ {
				util.SIMD_Add_Mult(&solution[j][k], &x[k], &A[i][j][k])
			}

		}
	}

	return solution
}

// Assume values have already been ModQ'd
func ChangeBase(val [d][ndiv8][8]int64) [OUT_SIZE]byte {

	var output [OUT_SIZE]byte

	for i := int64(0); i < d; i++ {
		for j := int64(0); j < ndiv8; j++ {
			util.SIMD_Mod_257(&val[i][j])
			for k := int64(0); k < 8; k++ {
				output[i*n+j*8+k] = byte(val[i][j][k])
				val[i][j][k] = val[i][j][k] >> 8
				output[N+i*ndiv8+j] = output[N+i*ndiv8+j] | byte(val[i][j][k]>>k)
			}
		}
	}

	return output
}

// Seperates out an ndiv8 length byte array from input
func sepInput(input [ndiv8 * m]byte, i int64) [ndiv8]byte {

	var sec [ndiv8]byte

	for j := int64(0); j < ndiv8; j++ {
		sec[j] = input[ndiv8*i+j]
	}

	return sec
}

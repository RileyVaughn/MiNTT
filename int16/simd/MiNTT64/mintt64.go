package MiNTT64

import (
	util "github.com/RileyVaughn/MiNTT/hash/int16/util"
)

func MinNTT64(input [ndiv8 * m]byte) [864]byte {

	return ChangeBase(ntt_sum(input))

}

func ncc(input [ndiv8]byte) [ndiv8][8]int16 {

	var intermed [ndiv8][8]int16

	for i := int16(0); i < ndiv8; i++ {
		intermed[i] = util.SIMD_Mult(&NTT8_TABLE[input[i]], &MULT_TABLE[i])
		util.Fake_SIMD_Mod(&intermed[i], q)
	}

	util.SIMD_AddSub(&intermed[0], &intermed[1])
	util.SIMD_AddSub(&intermed[2], &intermed[3])
	util.SIMD_AddSub(&intermed[4], &intermed[5])
	util.SIMD_AddSub(&intermed[6], &intermed[7])

	util.SIMD_Shift(&intermed[3], 4)
	util.SIMD_Shift(&intermed[7], 4)

	util.Fake_SIMD_Mod(&intermed[3], q)
	util.Fake_SIMD_Mod(&intermed[7], q)

	util.SIMD_AddSub(&intermed[0], &intermed[2])
	util.SIMD_AddSub(&intermed[1], &intermed[3])
	util.SIMD_AddSub(&intermed[4], &intermed[6])
	util.SIMD_AddSub(&intermed[5], &intermed[7])

	util.SIMD_Shift(&intermed[5], 2)
	util.SIMD_Shift(&intermed[6], 4)
	util.SIMD_Shift(&intermed[7], 6)

	util.Fake_SIMD_Mod(&intermed[5], q)
	util.Fake_SIMD_Mod(&intermed[6], q)
	util.Fake_SIMD_Mod(&intermed[7], q)

	util.SIMD_AddSub(&intermed[0], &intermed[4])
	util.SIMD_AddSub(&intermed[1], &intermed[5])
	util.SIMD_AddSub(&intermed[2], &intermed[6])
	util.SIMD_AddSub(&intermed[3], &intermed[7])

	// var out [n]int16

	// for i := int16(0); i < ndiv8; i++ {
	// 	for j := int16(0); j < 8; j++ {
	// 		out[8*i+j] = intermed[i][j]
	// 	}
	// }

	return intermed
}

func ntt_sum(input [ndiv8 * m]byte) [N]int16 {

	var solution [d][ndiv8][8]int16
	for i := int16(0); i < m; i++ {
		x := ncc(sepInput(input, i))
		for j := int16(0); j < d; j++ {
			for k := int16(0); k < n; k++ {
				//solution[n*j+k] = solution[n*j+k] + x[k]*A[i][n*j+k]
				util.SIMD_Add_Mult(&solution[j][k], &x[k], &A[i][j][k])
			}

		}
	}
	// This needs to happen but is slow
	var out [N]int16
	for i := int16(0); i < d; i++ {
		for j := int16(0); j < ndiv8; j++ {
			for k := int16(0); k < 8; k++ {
				out[i] = util.Mod(solution[i][j][k], q)
			}
		}

	}

	return out
}

// Assume values have already been ModQ'd
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

// Seperates out an ndiv8 length byte array from input
func sepInput(input [ndiv8 * m]byte, i int16) [ndiv8]byte {

	var sec [ndiv8]byte

	for j := int16(0); j < ndiv8; j++ {
		sec[j] = input[ndiv8*i+j]
	}

	return sec
}

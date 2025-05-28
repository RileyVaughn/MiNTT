package MiNTT128

import (
	util "github.com/RileyVaughn/MiNTT/hash/int64/util"
)

func MinNTT128(input [ndiv8 * m]byte) [864]byte {

	return ChangeBase(ntt_sum(input))

}

func ncc(input [ndiv8]byte) [n]int64 {

	var intermed [ndiv8][8]int64
	for k := int64(0); k < ndiv8; k++ {
		intermed[k] = util.Fake_SIMD_Mult(NTT8_TABLE[input[k]], MULT_TABLE[k])

	}

	util.Fake_SIMD_AddSub(&intermed[0], &intermed[1])
	util.Fake_SIMD_AddSub(&intermed[2], &intermed[3])
	util.Fake_SIMD_AddSub(&intermed[4], &intermed[5])
	util.Fake_SIMD_AddSub(&intermed[6], &intermed[7])
	util.Fake_SIMD_AddSub(&intermed[8], &intermed[9])
	util.Fake_SIMD_AddSub(&intermed[10], &intermed[11])
	util.Fake_SIMD_AddSub(&intermed[12], &intermed[13])
	util.Fake_SIMD_AddSub(&intermed[14], &intermed[15])

	util.Fake_SIMD_Shift(&intermed[3], 4)
	util.Fake_SIMD_Shift(&intermed[7], 4)
	util.Fake_SIMD_Shift(&intermed[11], 4)
	util.Fake_SIMD_Shift(&intermed[15], 4)

	util.Fake_SIMD_AddSub(&intermed[0], &intermed[2])
	util.Fake_SIMD_AddSub(&intermed[1], &intermed[3])
	util.Fake_SIMD_AddSub(&intermed[4], &intermed[6])
	util.Fake_SIMD_AddSub(&intermed[5], &intermed[7])
	util.Fake_SIMD_AddSub(&intermed[8], &intermed[10])
	util.Fake_SIMD_AddSub(&intermed[9], &intermed[11])
	util.Fake_SIMD_AddSub(&intermed[12], &intermed[14])
	util.Fake_SIMD_AddSub(&intermed[13], &intermed[15])

	util.Fake_SIMD_Shift(&intermed[5], 2)
	util.Fake_SIMD_Shift(&intermed[6], 4)
	util.Fake_SIMD_Shift(&intermed[7], 6)
	util.Fake_SIMD_Shift(&intermed[13], 2)
	util.Fake_SIMD_Shift(&intermed[14], 4)
	util.Fake_SIMD_Shift(&intermed[15], 6)

	util.Fake_SIMD_AddSub(&intermed[0], &intermed[4])
	util.Fake_SIMD_AddSub(&intermed[1], &intermed[5])
	util.Fake_SIMD_AddSub(&intermed[2], &intermed[6])
	util.Fake_SIMD_AddSub(&intermed[3], &intermed[7])
	util.Fake_SIMD_AddSub(&intermed[8], &intermed[12])
	util.Fake_SIMD_AddSub(&intermed[9], &intermed[13])
	util.Fake_SIMD_AddSub(&intermed[10], &intermed[14])
	util.Fake_SIMD_AddSub(&intermed[11], &intermed[15])

	util.Fake_SIMD_Shift(&intermed[9], 1)
	util.Fake_SIMD_Shift(&intermed[10], 2)
	util.Fake_SIMD_Shift(&intermed[11], 3)
	util.Fake_SIMD_Shift(&intermed[12], 4)
	util.Fake_SIMD_Shift(&intermed[13], 5)
	util.Fake_SIMD_Shift(&intermed[14], 6)
	util.Fake_SIMD_Shift(&intermed[15], 7)

	util.Fake_SIMD_AddSub(&intermed[0], &intermed[8])
	util.Fake_SIMD_AddSub(&intermed[1], &intermed[9])
	util.Fake_SIMD_AddSub(&intermed[2], &intermed[10])
	util.Fake_SIMD_AddSub(&intermed[3], &intermed[11])
	util.Fake_SIMD_AddSub(&intermed[4], &intermed[12])
	util.Fake_SIMD_AddSub(&intermed[5], &intermed[13])
	util.Fake_SIMD_AddSub(&intermed[6], &intermed[14])
	util.Fake_SIMD_AddSub(&intermed[7], &intermed[15])

	var out [n]int64

	for i := int64(0); i < ndiv8; i++ {
		for j := int64(0); j < 8; j++ {
			out[8*i+j] = intermed[i][j] //util.Mod(intermed[i][j], q)
		}
	}

	return out
}

func ntt_sum(input [ndiv8 * m]byte) [N]int64 {

	var solution [N]int64
	for i := int64(0); i < m; i++ {
		x := ncc(sepInput(input, i))
		for j := int64(0); j < d; j++ {
			for k := int64(0); k < n; k++ {
				solution[n*j+k] = solution[n*j+k] + x[k]*A[i][n*j+k]
			}

		}
	}
	for i := int64(0); i < N; i++ {
		solution[i] = util.Mod(solution[i], q)
	}

	return solution
}

// Assume values have already been ModQ'd
func ChangeBase(val [N]int64) [864]byte {

	var output [864]byte

	for i := int64(0); i < N; i++ {
		output[i] = byte(val[i])
		val[i] = val[i] >> 8
	}

	for i := int64(0); i < Ndiv8; i++ {
		for k := int64(0); k < 8; k++ {
			output[N+i] = output[N+i] | byte(val[8*i+k]>>k)
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

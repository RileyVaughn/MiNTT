package MiNTT64

import (
	util "github.com/RileyVaughn/MiNTT/hash/util"
)

func MinNTT64(input [ndiv8 * m]byte) [864]byte {

	return ChangeBase(ntt_sum(input))

}

func ncc(input [ndiv8]byte) [n]int {

	var intermed [ndiv8][8]int

	for i := 0; i < ndiv8; i++ {
		intermed[i] = util.SIMD_Mult(NTT8_TABLE[input[i]], MULT_TABLE[i])

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

	var out [n]int

	for i := 0; i < ndiv8; i++ {
		for j := 0; j < 8; j++ {
			out[8*i+j] = util.Mod(intermed[i][j], q)
		}
	}

	return out
}

func ntt_sum(input [ndiv8 * m]byte) [N]int {

	var solution [N]int
	for i := 0; i < m; i++ {
		x := ncc(sepInput(input, i))
		for j := 0; j < d; j++ {
			for k := 0; k < n; k++ {
				solution[n*j+k] = solution[n*j+k] + x[k]*A[i][n*j+k]
			}

		}
	}
	for i := 0; i < N; i++ {
		solution[i] = util.Mod(solution[i], q)
	}

	return solution
}

//Assume values have already been ModQ'd
func ChangeBase(val [N]int) [864]byte {

	var output [864]byte

	for i := 0; i < N; i++ {
		output[i] = byte(val[i])
		val[i] = val[i] >> 8
	}

	for i := 0; i < Ndiv8; i++ {
		for k := 0; k < 8; k++ {
			output[N+i] = output[N+i] | byte(val[8*i+k]>>k)
		}
	}

	return output
}

// Seperates out an ndiv8 length byte array from input
func sepInput(input [ndiv8 * m]byte, i int) [ndiv8]byte {

	var sec [ndiv8]byte

	for j := 0; j < ndiv8; j++ {
		sec[j] = input[ndiv8*i+j]
	}

	return sec
}

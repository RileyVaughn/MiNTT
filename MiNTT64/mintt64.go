package MiNTT64

import (
	util "github.com/RileyVaughn/MiNTT/hash/util"
)

//C GO
const d int = 12
const n int = 64
const N int = n * d
const m int = 216
const q int = 257
const ndiv8 int = n / 8
const Ndiv8 int = N / 8

var A [m][d * n]int
var NTT8_TABLE [256][8]int
var MULT_TABLE [8][8]int

func MinNTT64(input [ndiv8 * m]byte) [864]byte {

	return ChangeBase(ntt_sum(input))

}

func ncc(input [ndiv8]byte) [n]int {

	var intermed [ndiv8][8]int

	for i := 0; i < ndiv8; i++ {
		intermed[i] = SIMD_Mult(NTT8_TABLE[input[i]], MULT_TABLE[i])

	}

	// SIMD_Shift(&intermed[1], 8)
	// SIMD_Shift(&intermed[3], 8)
	// SIMD_Shift(&intermed[5], 8)
	// SIMD_Shift(&intermed[7], 8)

	SIMD_AddSub(&intermed[0], &intermed[1])
	SIMD_AddSub(&intermed[2], &intermed[3])
	SIMD_AddSub(&intermed[4], &intermed[5])
	SIMD_AddSub(&intermed[6], &intermed[7])

	// SIMD_Shift(&intermed[2], 2)
	// SIMD_Shift(&intermed[3], 6)
	// SIMD_Shift(&intermed[6], 2)
	// SIMD_Shift(&intermed[7], 6)

	SIMD_Shift(&intermed[3], 4)
	SIMD_Shift(&intermed[7], 4)

	SIMD_AddSub(&intermed[0], &intermed[2])
	SIMD_AddSub(&intermed[1], &intermed[3])
	SIMD_AddSub(&intermed[4], &intermed[6])
	SIMD_AddSub(&intermed[5], &intermed[7])

	// SIMD_Shift(&intermed[4], 1)
	// SIMD_Shift(&intermed[5], 3)
	// SIMD_Shift(&intermed[6], 5)
	// SIMD_Shift(&intermed[7], 7)

	SIMD_Shift(&intermed[5], 2)
	SIMD_Shift(&intermed[6], 4)
	SIMD_Shift(&intermed[7], 6)

	SIMD_AddSub(&intermed[0], &intermed[4])
	SIMD_AddSub(&intermed[1], &intermed[5])
	SIMD_AddSub(&intermed[2], &intermed[6])
	SIMD_AddSub(&intermed[3], &intermed[7])

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

//Fake for now
func SIMD_AddSub(vec1 *[8]int, vec2 *[8]int) {
	for i := 0; i < 8; i++ {
		addSub(&vec1[i], &vec2[i])
	}
}

//Fake for now
func SIMD_Mult(vec1 [8]int, vec2 [8]int) [8]int {

	var product [8]int
	for i := 0; i < 8; i++ {
		product[i] = vec1[i] * vec2[i]
	}
	return product
}

//Fake for now
func SIMD_Shift(vec *[8]int, shift int) {

	for i := 0; i < 8; i++ {
		vec[i] = vec[i] << shift
	}

}

//Takes two int ptrs a and b, replaces *a with *a+*b *b with *a-*b
func addSub(a *int, b *int) {
	temp := *b
	*b = *a - *b
	*a = *a + temp
}

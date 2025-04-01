package MiNTT128

import (
	util "github.com/RileyVaughn/MiNTT/hash/util"
)

//C GO
const d int = 6
const n int = 128
const N int = n * d
const m int = 108
const q int = 257
const ndiv8 int = n / 8
const Ndiv8 int = N / 8

var A [m][d * n]int
var bit2ByteTable [256][8]int
var NTT8_TABLE [256][8]int

func MinNTT8(input [ndiv8 * m]byte) [864]byte {

	return ChangeBase(ntt_sum(input))

}

func ncc(input [ndiv8]byte) [n]int {

	var intermed [ndiv8][8]int

	for i := 0; i < ndiv8; i++ {
		intermed[i] = NTT8_TABLE[input[i]]
	}

}

func ntt_sum(input [ndiv8 * m]byte) [N]int {

	var solution [N]int
	for i := 0; i < m; i++ {
		x := NTT8_TABLE[input[i]]
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

//Takes two int ptrs a and b, replaces *a with *a+*b *b with *a-*b
func AddSub(a *int, b *int) {
	temp := *b
	*b = *a - *b
	*a = *a + temp
}

//Assume vaules haveaady been ModQ'd
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

//Fake for now
func SIMD_AddSub() {

}

//Fake for now
func SIMD_Shift() {

}

package MiNTT8

import (
	util "github.com/RileyVaughn/MiNTT/hash/util"
)

const d int = 96

const n int = 8
const ndiv8 int = n / 8
const N int = n * d
const Ndiv8 = N / 8
const q int = 257

var A [m][d * n]int
var bit2ByteTable [256][8]int

// m must be greater than d*log_2(q)
const m int = 1728

func MinNNT8(input [n * m / 8]byte) [864]byte {

	return ChangeBase(ntt_sum(input))

}

// Y_i = Sum256(x_k*omega^k(2i+1))
//Assume input has previously been bit reversed
func ntt_part(input byte) [n]int {

	var intermed [n]int

	x := bit2ByteTable[input]
	for j := 0; j < 8; j++ {
		intermed[j] = x[j]
	}

	intermed[1] = intermed[1] << 4
	intermed[3] = intermed[3] << 4
	intermed[5] = intermed[5] << 4
	intermed[7] = intermed[7] << 4

	AddSub(&intermed[0], &intermed[1])
	AddSub(&intermed[2], &intermed[3])
	AddSub(&intermed[4], &intermed[5])
	AddSub(&intermed[6], &intermed[7])

	intermed[2] = intermed[2] << 2
	intermed[3] = intermed[3] << 6
	intermed[6] = intermed[6] << 2
	intermed[7] = intermed[7] << 6

	AddSub(&intermed[0], &intermed[2])
	AddSub(&intermed[1], &intermed[3])
	AddSub(&intermed[4], &intermed[6])
	AddSub(&intermed[5], &intermed[7])

	intermed[4] = intermed[4] << 1
	intermed[5] = intermed[5] << 3
	intermed[6] = intermed[6] << 5
	intermed[7] = intermed[7] << 7

	AddSub(&intermed[0], &intermed[4])
	AddSub(&intermed[1], &intermed[5])
	AddSub(&intermed[2], &intermed[6])
	AddSub(&intermed[3], &intermed[7])

	return intermed
}

func ntt_sum(input [ndiv8 * m]byte) [N]int {

	var solution [N]int

	for i := 0; i < m; i++ {
		x := ntt_part(input[i])
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

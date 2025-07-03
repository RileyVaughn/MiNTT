package MiNTT16

import (
	util "github.com/RileyVaughn/MiNTT/hash/util"
)

const d int = 48

const n int = 16
const N int = n * d
const q int = 65537

// m must be greater than d*log_2(q)
const m int = 1632

// //const A[m][d][n] uint64

// //const fft_table [8][256]uint64

//Inputs n*m bits as n*m/8 bytes. m is based on d
//Outputs log_2(q^N)=N*17 bits = 1632 bytes
func MinNNT16(input [n * m / 8]byte) [1632]byte {

	return ChangeBase(ntt_sum(input))

}

// Y_i = Sum256(x_k*omega^k(2i+1))
func ntt_part(input [n / 8]byte) [n]int {

	var out [n]int
	var intermed [n]int

	bit2ByteTable := util.BitsFromByteTable()

	//Assume input has previously been bit reversed

	for i := 0; i < 2; i++ {
		x := bit2ByteTable[input[i]]

		for j := 0; j < 8; j++ {
			intermed[8*i+j] = x[j]
		}
	}
	intermed[1] = intermed[1] << 8
	intermed[3] = intermed[3] << 8
	intermed[5] = intermed[5] << 8
	intermed[7] = intermed[7] << 8
	intermed[9] = intermed[9] << 8
	intermed[11] = intermed[11] << 8
	intermed[13] = intermed[13] << 8
	intermed[15] = intermed[15] << 8

	AddSub(&intermed[0], &intermed[1])
	AddSub(&intermed[2], &intermed[3])
	AddSub(&intermed[4], &intermed[5])
	AddSub(&intermed[6], &intermed[7])
	AddSub(&intermed[8], &intermed[9])
	AddSub(&intermed[10], &intermed[11])
	AddSub(&intermed[12], &intermed[13])
	AddSub(&intermed[14], &intermed[15])

	intermed[2] = intermed[2] << 4
	intermed[3] = intermed[3] << 12
	intermed[6] = intermed[6] << 4
	intermed[7] = intermed[7] << 12
	intermed[10] = intermed[10] << 4
	intermed[11] = intermed[11] << 12
	intermed[14] = intermed[14] << 4
	intermed[15] = intermed[15] << 12

	AddSub(&intermed[0], &intermed[2])
	AddSub(&intermed[1], &intermed[3])
	AddSub(&intermed[4], &intermed[6])
	AddSub(&intermed[5], &intermed[7])
	AddSub(&intermed[8], &intermed[10])
	AddSub(&intermed[9], &intermed[11])
	AddSub(&intermed[12], &intermed[14])
	AddSub(&intermed[13], &intermed[15])

	intermed[4] = intermed[4] << 2
	intermed[5] = intermed[5] << 6
	intermed[6] = intermed[6] << 10
	intermed[7] = intermed[7] << 14
	intermed[12] = intermed[12] << 2
	intermed[13] = intermed[13] << 6
	intermed[14] = intermed[14] << 10
	intermed[15] = intermed[15] << 14

	AddSub(&intermed[0], &intermed[4])
	AddSub(&intermed[1], &intermed[5])
	AddSub(&intermed[2], &intermed[6])
	AddSub(&intermed[3], &intermed[7])
	AddSub(&intermed[8], &intermed[12])
	AddSub(&intermed[9], &intermed[13])
	AddSub(&intermed[10], &intermed[14])
	AddSub(&intermed[11], &intermed[15])

	intermed[8] = intermed[8] << 1
	intermed[9] = intermed[9] << 3
	intermed[10] = intermed[10] << 5
	intermed[11] = intermed[11] << 7
	intermed[12] = intermed[12] << 9
	intermed[13] = intermed[13] << 11
	intermed[14] = intermed[14] << 13
	intermed[15] = intermed[15] << 15

	out[0] = util.Mod(intermed[0]+intermed[8], q)
	out[8] = util.Mod(intermed[0]-intermed[8], q)
	out[1] = util.Mod(intermed[1]+intermed[9], q)
	out[9] = util.Mod(intermed[1]-intermed[9], q)
	out[2] = util.Mod(intermed[2]+intermed[10], q)
	out[10] = util.Mod(intermed[2]-intermed[10], q)
	out[3] = util.Mod(intermed[3]+intermed[11], q)
	out[11] = util.Mod(intermed[3]-intermed[11], q)
	out[4] = util.Mod(intermed[4]+intermed[12], q)
	out[12] = util.Mod(intermed[4]-intermed[12], q)
	out[5] = util.Mod(intermed[5]+intermed[13], q)
	out[13] = util.Mod(intermed[5]-intermed[13], q)
	out[6] = util.Mod(intermed[6]+intermed[14], q)
	out[14] = util.Mod(intermed[6]-intermed[14], q)
	out[7] = util.Mod(intermed[7]+intermed[15], q)
	out[15] = util.Mod(intermed[7]-intermed[15], q)

	return out
}

func ntt_sum(input [n * m / 8]byte) [N]int {

	A := ReadKey("./MiNNT16/key.csv")
	var solution [N]int

	for i := 0; i < m; i++ {
		t_array := [2]byte{input[i*(n/8)], input[i*(n/8)+1]}
		x := ntt_part(t_array)
		for j := 0; j < d; j++ {
			for k := 0; k < n; k++ {
				solution[n*j+k] = solution[n*j+k] + x[k]*A[i][n*j+k]
			}

		}
	}

	for i := 0; i < n*d; i++ {
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
func ChangeBase(val [N]int) [1632]byte {

	//n = 16
	//d = 48
	//ceil[log_2(q=65537)] = 17
	//=1632
	var output [1632]byte

	for i := 0; i < N; i++ {
		output[2*i] = byte(val[i])
		output[2*i+1] = byte(val[i] >> 8)
		val[i] = val[i] >> 16
	}

	for i := 0; i < N/8; i++ {
		for k := 0; k < 8; k++ {
			output[N+i] = output[i] | byte(val[8*i+k]>>k)
		}
	}

	return output
}

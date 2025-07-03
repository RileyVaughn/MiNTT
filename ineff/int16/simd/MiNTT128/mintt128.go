package MiNTT128

import (
	util "github.com/RileyVaughn/MiNTT/hash/int16/util"
)

func MiNTT128(input [ndiv8 * m]byte) [OUT_SIZE]byte {

	return ChangeBase(ntt_sum(input))

}

func ncc(input [ndiv8]byte) [ndiv8][8]int16 {

	var intermed [ndiv8][8]int16
	for k := int16(0); k < ndiv8; k++ {
		intermed[k] = util.SIMD_Mult(&NTT8_TABLE[input[k]], &MULT_TABLE[k])
		util.SIMD_Q_Reduce(&intermed[k])

	}

	util.SIMD_AddSub(&intermed[0], &intermed[1])
	util.SIMD_AddSub(&intermed[2], &intermed[3])
	util.SIMD_AddSub(&intermed[4], &intermed[5])
	util.SIMD_AddSub(&intermed[6], &intermed[7])
	util.SIMD_AddSub(&intermed[8], &intermed[9])
	util.SIMD_AddSub(&intermed[10], &intermed[11])
	util.SIMD_AddSub(&intermed[12], &intermed[13])
	util.SIMD_AddSub(&intermed[14], &intermed[15])

	util.SIMD_Shift(&intermed[3], 4)
	util.SIMD_Shift(&intermed[7], 4)
	util.SIMD_Shift(&intermed[11], 4)
	util.SIMD_Shift(&intermed[15], 4)

	util.SIMD_Q_Reduce(&intermed[3])
	util.SIMD_Q_Reduce(&intermed[7])
	util.SIMD_Q_Reduce(&intermed[11])
	util.SIMD_Q_Reduce(&intermed[15])

	util.SIMD_AddSub(&intermed[0], &intermed[2])
	util.SIMD_AddSub(&intermed[1], &intermed[3])
	util.SIMD_AddSub(&intermed[4], &intermed[6])
	util.SIMD_AddSub(&intermed[5], &intermed[7])
	util.SIMD_AddSub(&intermed[8], &intermed[10])
	util.SIMD_AddSub(&intermed[9], &intermed[11])
	util.SIMD_AddSub(&intermed[12], &intermed[14])
	util.SIMD_AddSub(&intermed[13], &intermed[15])

	util.SIMD_Shift(&intermed[5], 2)
	util.SIMD_Shift(&intermed[6], 4)
	util.SIMD_Q_Reduce(&intermed[7])
	util.SIMD_Shift(&intermed[7], 6)
	util.SIMD_Shift(&intermed[13], 2)
	util.SIMD_Shift(&intermed[14], 4)
	util.SIMD_Q_Reduce(&intermed[15])
	util.SIMD_Shift(&intermed[15], 6)

	util.SIMD_Q_Reduce(&intermed[5])
	util.SIMD_Q_Reduce(&intermed[6])
	util.SIMD_Q_Reduce(&intermed[7])
	util.SIMD_Q_Reduce(&intermed[13])
	util.SIMD_Q_Reduce(&intermed[14])
	util.SIMD_Q_Reduce(&intermed[15])

	util.SIMD_AddSub(&intermed[0], &intermed[4])
	util.SIMD_AddSub(&intermed[1], &intermed[5])
	util.SIMD_AddSub(&intermed[2], &intermed[6])
	util.SIMD_AddSub(&intermed[3], &intermed[7])
	util.SIMD_AddSub(&intermed[8], &intermed[12])
	util.SIMD_AddSub(&intermed[9], &intermed[13])
	util.SIMD_AddSub(&intermed[10], &intermed[14])
	util.SIMD_AddSub(&intermed[11], &intermed[15])

	util.SIMD_Shift(&intermed[9], 1)
	util.SIMD_Shift(&intermed[10], 2)
	util.SIMD_Shift(&intermed[11], 3)
	util.SIMD_Shift(&intermed[12], 4)
	util.SIMD_Shift(&intermed[13], 5)
	util.SIMD_Q_Reduce(&intermed[14])
	util.SIMD_Shift(&intermed[14], 6)
	util.SIMD_Q_Reduce(&intermed[15])
	util.SIMD_Shift(&intermed[15], 6)
	util.SIMD_Q_Reduce(&intermed[15])
	util.SIMD_Shift(&intermed[15], 1)

	util.SIMD_Q_Reduce(&intermed[9])
	util.SIMD_Q_Reduce(&intermed[10])
	util.SIMD_Q_Reduce(&intermed[11])
	util.SIMD_Q_Reduce(&intermed[12])
	util.SIMD_Q_Reduce(&intermed[13])
	util.SIMD_Q_Reduce(&intermed[14])
	util.SIMD_Q_Reduce(&intermed[15])

	util.SIMD_AddSub(&intermed[0], &intermed[8])
	util.SIMD_AddSub(&intermed[1], &intermed[9])
	util.SIMD_AddSub(&intermed[2], &intermed[10])
	util.SIMD_AddSub(&intermed[3], &intermed[11])
	util.SIMD_AddSub(&intermed[4], &intermed[12])
	util.SIMD_AddSub(&intermed[5], &intermed[13])
	util.SIMD_AddSub(&intermed[6], &intermed[14])
	util.SIMD_AddSub(&intermed[7], &intermed[15])

	return intermed
}

func ntt_sum(input [ndiv8 * m]byte) [d][ndiv8][8]int16 {

	var solution [d][ndiv8][8]int16
	for i := int16(0); i < m; i++ {
		x := ncc(sepInput(input, i))
		for j := int16(0); j < d; j++ {
			for k := int16(0); k < ndiv8; k++ {
				util.SIMD_Add_Mult(&solution[j][k], &x[k], &A[i][j][k])
			}

		}
	}

	return solution
}

// Assumes MOD has not yet occured
// Changes base from  257 to 256, moves the extra bits to the end (<first N*8 256 bits>N + <N end bits> ndiv8*d)
func ChangeBase(val [d][ndiv8][8]int16) [OUT_SIZE]byte {

	var output [OUT_SIZE]byte

	for i := int16(0); i < d; i++ {
		for j := int16(0); j < ndiv8; j++ {
			util.SIMD_Mod_257(&val[i][j])
			for k := int16(0); k < 8; k++ {
				output[i*n+j*8+k] = byte(val[i][j][k])
				val[i][j][k] = val[i][j][k] >> 8
				output[N+i*ndiv8+j] = output[N+i*ndiv8+j] | byte(val[i][j][k]>>k)
			}
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

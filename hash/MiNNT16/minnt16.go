package MiNNT16

const d int = 1

const n int = 16
const N int = n * d
const q int = 65537
const m int = 32

// //const A[m][d][n] uint64

// //const fft_table [8][256]uint64

// //Inputs n*m bits as n*m/8 bytes. m is based on d
// //Outputs log_2(q^N)= N*17 bits
// func MinNNT(input [n * m / 8]byte) [N * 17]byte {

// 	//Split input into m,n length arrays
// 	//FFT each Array with A_i

// 	return ChangeBase(SumArrays())
// }

// Y_i = Sum256(x_k*omega^k(2i+1))
func fft(input [n / 8]byte) [n]int {

	var out [n]int
	var intermed [n]int

	bit2ByteTable := bitsFromByteTable()

	//n=16 varibles set
	//input can be applied arbitraily (assume input has previously been bit reversed)
	//can use xor instead of plus

	for i := 0; i < 2; i++ {
		x := bit2ByteTable[input[i]]
		for j := 0; j < 4; j = j + 2 {
			intermed[8*i+2*j] = x[j]
			intermed[8*i+2*j+1] = x[j] << 8
			AddSub(&intermed[8*i+2*j], &intermed[8*i+2*j+1])
		}
	}

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
	intermed[5] = intermed[4] << 10
	intermed[6] = intermed[4] << 6
	intermed[7] = intermed[4] << 14
	intermed[12] = intermed[4] << 2
	intermed[13] = intermed[4] << 10
	intermed[14] = intermed[4] << 6
	intermed[15] = intermed[4] << 14

	AddSub(&intermed[0], &intermed[4])
	AddSub(&intermed[1], &intermed[5])
	AddSub(&intermed[2], &intermed[6])
	AddSub(&intermed[3], &intermed[7])
	AddSub(&intermed[8], &intermed[12])
	AddSub(&intermed[9], &intermed[13])
	AddSub(&intermed[10], &intermed[14])
	AddSub(&intermed[11], &intermed[15])

	intermed[8] = intermed[8] << 1
	intermed[9] = intermed[9] << 9
	intermed[10] = intermed[10] << 5
	intermed[11] = intermed[11] << 13
	intermed[12] = intermed[12] << 3
	intermed[13] = intermed[13] << 11
	intermed[14] = intermed[14] << 7
	intermed[15] = intermed[15] << 15

	out[0] = intermed[0] + intermed[8]
	out[1] = intermed[0] - intermed[8]
	out[2] = intermed[1] + intermed[9]
	out[3] = intermed[1] - intermed[9]
	out[4] = intermed[2] + intermed[10]
	out[5] = intermed[2] - intermed[10]
	out[6] = intermed[3] + intermed[11]
	out[7] = intermed[3] - intermed[11]
	out[8] = intermed[4] + intermed[12]
	out[9] = intermed[4] - intermed[12]
	out[10] = intermed[5] + intermed[13]
	out[11] = intermed[5] - intermed[13]
	out[12] = intermed[6] + intermed[14]
	out[13] = intermed[6] - intermed[14]
	out[14] = intermed[7] + intermed[15]
	out[15] = intermed[7] - intermed[15]

	return out
}

//Takes two int ptrs a and b, replaces *a with *a+*b *b with *a-*b
func AddSub(a *int, b *int) {
	temp := *b
	*b = *a - *b
	*a = *a + temp
}

// func SumArrays(val [m][d][n]uint64) [d][n]uint64 {

// 	var sum [d][n]uint64
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < d; j++ {
// 			for k := 0; k < n; k++ {
// 				sum[j][k] = sum[j][k] + val[i][j][k]
// 			}
// 		}
// 	}

// 	return sum
// }

// func ChangeBase(val [d][n]uint64) [N * 17]byte {

// 	// mod 65537
// 	// seperate last bit to end

func bitsFromByteTable() [256][8]int {

	var table [256][8]int

	for i := int(0); i < 256; i++ {
		table[i][0] = i % 2
		table[i][1] = (i >> 1) % 2
		table[i][2] = (i >> 2) % 2
		table[i][3] = (i >> 3) % 2
		table[i][4] = (i >> 4) % 2
		table[i][5] = (i >> 5) % 2
		table[i][6] = (i >> 6) % 2
		table[i][7] = (i >> 7) % 2
	}

	return table
}

package main

// import (
// 	keygen "github.com/RileyVaughn/MiNTT/hash/key"
// )

// const d int = 1
// const n int = 128
// const N int = n * d
// const q int = 65537
// const m int = 32

// // //const A[m][d][n] uint64

// // //const fft_table [8][256]uint64

// // //Inputs n*m bits as n*m/8 bytes. m is based on d
// // //Outputs log_2(q^N)= N*17 bits
// // func MinNNT(input [n * m / 8]byte) [N * 17]byte {

// // 	//Split input into m,n length arrays
// // 	//FFT each Array with A_i

// // 	return ChangeBase(SumArrays())
// // }

// // Y_i = Sum256(x_k*omega^k(2i+1))
// func fft(input [n / 8]byte) [n]uint64 {

// 	var x [n]uint64
// 	//var out [n]uint64
// 	fft_table := keygen.TableGen()
// 	omegas := keygen.OmegaGen()
// 	var partial_sums [n / 8][2][8]uint64

// 	for i0 := 0; i0 < n/8; i0++ {
// 		for i1 := 0; i1 < 2; i1++ {
// 			for k0 := 0; k0 < n/8; k0++ {
// 				partial_sums[][][] := fft_table[input[k0]][i1][i0] * omegas[k0*(2*i0+1)]
// 			}

// 		}

// 	}

// 	return same
// }

// // func SumArrays(val [m][d][n]uint64) [d][n]uint64 {

// // 	var sum [d][n]uint64
// // 	for i := 0; i < m; i++ {
// // 		for j := 0; j < d; j++ {
// // 			for k := 0; k < n; k++ {
// // 				sum[j][k] = sum[j][k] + val[i][j][k]
// // 			}
// // 		}
// // 	}

// // 	return sum
// // }

// // func ChangeBase(val [d][n]uint64) [N * 17]byte {

// // 	// mod 65537
// // 	// seperate last bit to end

// // }

// // 4080 64th root
// // 4938 128th root
// // 59963 256th root
// // 44120 512th root
// // 10423 1024th root
// // 57968 2048th root
// // 43265 4096th root
// // 56153 8192th root
// // 7348 16384th root
// // 927 32768th root
// // 40264 65536th root

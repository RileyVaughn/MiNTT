package hash

const d int = 1
const n int = 256
const N int = n * d
const q int = 65537
const m int = 32

//const A[m][d][n] uint64

//const fft_table [8][256]uint64

//Inputs N*m bits as n*m/8 bytes
//Outputs log_2(q^N)= N*17 bits
func MinNNT(input [N * m / 8]byte) [N * 17]byte {

	//Split input into m,d,n arrays
	//FFT each Array with A_i

	return ChangeBase(SumArrays())
}

func fft(polys [d][n / 8]byte) [d][n]uint64 {

	var product [d][n]uint64

	for i := 0; i < d; i++ {
		for j := 0; j < n/8; j++ {
			//Assume fft table is just 1 and 0s
			var val uint64 = fft_table[j]

		}
	}
	return product
}

func SumArrays(val [m][d][n]uint64) [d][n]uint64 {

	var sum [d][n]uint64
	for i := 0; i < m; i++ {
		for j := 0; j < d; j++ {
			for k := 0; k < n; k++ {
				sum[j][k] = sum[j][k] + val[i][j][k]
			}
		}
	}

	return sum
}

func ChangeBase(val [d][n]uint64) [N * 17]byte {

	// mod 65537
	// seperate last bit to end

}

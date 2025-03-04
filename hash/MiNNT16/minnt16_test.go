package MiNNT16

import (
	"testing"
)

func TestFFT(t *testing.T) {

	r

}

func matVecMult(mat [n][n]int, input [n / 8]byte) [n]int {

	var product [n]int
	var vec [n]int
	btb_table := bitsFromByteTable()

	for i := 0; i < n/8; i++ {
		temp := btb_table[input[i]]
		for j := 0; j < 8; j++ {
			vec[8*i+j] = temp[j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			product[i] = product[i] + mat[i][j]*vec[j]
		}
	}

	return product
}

func genNCCMat(omega int) [n][n]int {

	var ncc_mat [n][n]int

	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			if (k*(2*i+1))%(2*n) <= n {
				ncc_mat[i][k] = util.intPow(omega, (k*(2*i+1))%n)

			} else {
				ncc_mat[i][k] = -1 * util.intPow(omega, ((k*(2*i+1))%n))

			}
		}

	}

	return ncc_mat
}

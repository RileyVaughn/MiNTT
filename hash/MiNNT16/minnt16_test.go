package MiNNT16

import (
	"testing"

	util "github.com/RileyVaughn/MiNTT/hash/util"
)

func TestNTT_part(t *testing.T) {

	input := [2]byte{25, 34}
	// input := [2]byte{5, 2}

	want := NCCVecMult(2, input)
	result := ntt_part(input)

	if result != want {
		t.Fatalf("(Test fft) Bad FFT: %v != %v", result, want)
	}

}

func NCCVecMult(omega int, input [n / 8]byte) [n]int {

	var product [n]int
	var vec [n]int
	mat := genNCCMat(omega)
	btb_table := util.BitsFromByteTable()

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

	for i := 0; i < n; i++ {
		product[i] = util.Mod(product[i], q)
	}

	return product
}

//Returns NCC matrix for bitreversed output
//quick and dirty, assume n = 16
func genNCCMat(omega int) [n][n]int {

	var ncc_mat [n][n]int

	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			if (k*(2*i+1))%(2*n) <= n {
				ncc_mat[i][k] = util.IntPow(omega, (k*(2*i+1))%n)

			} else {
				ncc_mat[i][k] = -1 * util.IntPow(omega, ((k*(2*i+1))%n))

			}
		}

	}

	var br_ncc_mat [n][n]int
	var br_arr [n]int = [n]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15} //dirty
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			br_ncc_mat[j][br_arr[i]] = ncc_mat[j][i]
		}

	}

	return br_ncc_mat
}

package MiNNT16

import (
	"testing"

	util "github.com/RileyVaughn/MiNTT/hash/util"
)

func TestFFT(t *testing.T) {

	input := [2]byte{1, 2}

	want := NCCVecMult(2, input)
	result := fft(input)

	if result != want {
		t.Fatalf("(Test fft) Bad FFT: %v != %v", result, want)
	}

}

// func TestIsEqual(t *testing.T) {

// 	polyTest := ReadPolys("polynomials.csv")

// 	var wantTest []bool = []bool{true, false}

// 	result := polyTest[0].IsEqual(polyTest[0])
// 	if result != wantTest[0] {
// 		t.Fatalf("(TestIsEqual) Bad Equiv: %v != %v", result, wantTest[0])
// 	}

// 	result = polyTest[0].IsEqual(polyTest[1])
// 	if result != wantTest[1] {
// 		t.Fatalf("(TestIsEqual) Bad Equiv: %v != %v", result, wantTest[1])
// 	}

// }

func NCCVecMult(omega int, input [n / 8]byte) [n]int {

	var product [n]int
	var vec [n]int
	mat := genNCCMat(omega)
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
				ncc_mat[i][k] = util.IntPow(omega, (k*(2*i+1))%n)

			} else {
				ncc_mat[i][k] = -1 * util.IntPow(omega, ((k*(2*i+1))%n))

			}
		}

	}

	return ncc_mat
}

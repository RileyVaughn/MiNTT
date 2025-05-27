package MiNTT16

import (
	"math/rand"
	"strconv"
	"testing"

	util "github.com/RileyVaughn/MiNTT/hash/util"
)

const TEST_SIZE = 1000

func TestNTT_part(t *testing.T) {

	seed, _ := strconv.Atoi("MiNNT")
	rand.Seed(int64(seed))

	for i := 0; i < TEST_SIZE; i++ {

		input := [2]byte{byte(rand.Intn(256)), byte(rand.Intn(256))}

		want := NCCVecMult(2, input)
		result := ntt_part(input)

		if result != want {
			t.Fatalf("(Test ntt_part) Bad FFT: %v != %v", result, want)
		}
	}
}

func TestNTT_sum(t *testing.T) {
	seed, _ := strconv.Atoi("MiNNT")
	rand.Seed(int64(seed))

	for i := 0; i < TEST_SIZE; i++ {
		var input [n * m / 8]byte
		for j := 0; j < n*m/8; j++ {
			input[j] = byte(rand.Intn(256))
		}
		result := ntt_sum(input)
		want := VecMultSum(input)

		if result != want {
			t.Fatalf("(Test ntt_sum) Bad SUM: %v != %v", result, want)
		}
	}
}

//Test Change Base
// func TestChangeBase(t *testing.T) {

// }

//Test Full function

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

//Assumes n=16
func VecMultSum(input [n * m / 8]byte) [N]int {

	A := ReadKey("key.csv")
	var output [N]int

	for i := 0; i < m; i++ {
		part_input := [2]byte{input[2*i], input[2*i+1]}
		part_want := NCCVecMult(2, part_input)
		for j := 0; j < d; j++ {
			for k := 0; k < n; k++ {
				output[j*n+k] = output[j*n+k] + part_want[k]*A[i][j*n+k]
			}

		}
	}
	for i := 0; i < N; i++ {
		output[i] = util.Mod(output[i], q)
	}
	return output
}

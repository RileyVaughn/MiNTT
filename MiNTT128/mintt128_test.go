package MiNTT128

import (
	"math/rand"
	"strconv"
	"testing"

	util "github.com/RileyVaughn/MiNTT/hash/util"
)

const TEST_SIZE = 1000

var bit2ByteTable [256][8]int

func TestNCC(t *testing.T) {

	seed, _ := strconv.Atoi("MiNNT")
	rand.Seed(int64(seed))

	SetupM128()
	bit2ByteTable = util.BitsFromByteTable()

	for i := 0; i < TEST_SIZE; i++ {

		var input [ndiv8]byte
		for j := 0; j < ndiv8; j++ {
			input[j] = byte(rand.Intn(256))
		}

		want := NCCVecMult(27, input)
		result := ncc(input)
		for i := 0; i < n; i++ {
			result[i] = util.Mod(result[i], q)
		}

		if result != want {
			t.Fatalf("(Test ntt_part) Bad FFT:\n %v\n%v", result, want)
		}

	}
}

func NCCVecMult(omega int, input [ndiv8]byte) [n]int {

	var product [n]int
	var vec [n]int
	mat := genNCCMat(omega)

	for i := 0; i < ndiv8; i++ {
		t_vec := bit2ByteTable[input[i]]

		for j := 0; j < 8; j++ {
			vec[8*i+j] = t_vec[j]
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
func genNCCMat(omega int) [n][n]int {

	var ncc_mat [n][n]int

	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			if (k*(2*i+1))%(2*n) <= n {
				ncc_mat[i][k] = util.IntPow(omega, (k*(2*i+1))%n, q)

			} else {
				ncc_mat[i][k] = -1 * util.IntPow(omega, ((k*(2*i+1))%n), q)

			}
		}

	}

	var br_ncc_mat [n][n]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			br_ncc_mat[j][util.Bit_Rev(i, n)] = ncc_mat[j][i]
		}
	}

	return br_ncc_mat
}

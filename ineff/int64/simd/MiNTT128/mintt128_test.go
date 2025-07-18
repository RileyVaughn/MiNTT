package MiNTT128

import (
	"math/rand"
	"strconv"
	"testing"

	util "github.com/RileyVaughn/MiNTT/hash/int64/util"
)

const TEST_SIZE = 1000

var bit2ByteTable [256][8]int64

func TestNCC(t *testing.T) {

	seed, _ := strconv.Atoi("MiNNT")
	rand.Seed(int64(seed))

	SetupM128()
	bit2ByteTable = util.BitsFromByteTable()

	for i := int64(0); i < TEST_SIZE; i++ {

		var input [ndiv8]byte
		for j := int64(0); j < ndiv8; j++ {
			input[j] = byte(rand.Intn(256))
		}

		want := NCCVecMult(27, input)
		out := ncc(input)

		var result [n]int64
		for i := 0; i < int(ndiv8); i++ {
			for j := 0; j < 8; j++ {
				result[8*i+j] = out[i][j]
			}
		}

		for i := int64(0); i < n; i++ {
			util.Mod_257(&result[i])
		}

		if result != want {
			t.Fatalf("(Test ntt_part) Bad FFT:\n %v\n%v", result, want)
		}

	}
}

func NCCVecMult(omega int64, input [ndiv8]byte) [n]int64 {

	var product [n]int64
	var vec [n]int64
	mat := genNCCMat(omega)

	for i := int64(0); i < ndiv8; i++ {
		t_vec := bit2ByteTable[input[i]]

		for j := int64(0); j < 8; j++ {
			vec[8*i+j] = t_vec[j]
		}
	}

	for i := int64(0); i < n; i++ {
		for j := int64(0); j < n; j++ {
			product[i] = product[i] + mat[i][j]*vec[j]
		}
	}

	for i := int64(0); i < n; i++ {
		product[i] = util.Mod(product[i], q)
	}

	return product
}

// Returns NCC matrix for bitreversed output
func genNCCMat(omega int64) [n][n]int64 {

	var ncc_mat [n][n]int64

	for i := int64(0); i < n; i++ {
		for k := int64(0); k < n; k++ {
			if int64(k*(2*i+1))%(2*n) <= n {
				ncc_mat[i][k] = util.IntPow(omega, (k*(2*i+1))%n, q)

			} else {
				ncc_mat[i][k] = -1 * util.IntPow(omega, ((k*(2*i+1))%n), q)

			}
		}

	}

	var br_ncc_mat [n][n]int64
	for i := int64(0); i < n; i++ {
		for j := int64(0); j < n; j++ {
			br_ncc_mat[j][util.Bit_Rev(i, n)] = ncc_mat[j][i]
		}
	}

	return br_ncc_mat
}

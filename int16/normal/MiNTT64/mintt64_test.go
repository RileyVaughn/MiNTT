package MiNTT64

import (
	"math/rand"
	"strconv"
	"testing"

	util "github.com/RileyVaughn/MiNTT/hash/int16/util"
)

const TEST_SIZE = 1000

var bit2ByteTable [256][8]int16

func TestNCC(t *testing.T) {

	seed, _ := strconv.Atoi("MiNNT")
	rand.Seed(int64(seed))

	SetupM64()
	bit2ByteTable = util.BitsFromByteTable()

	for i := int16(0); i < TEST_SIZE; i++ {

		var input [ndiv8]byte
		for j := int16(0); j < ndiv8; j++ {
			input[j] = byte(rand.Intn(256))
		}

		want := NCCVecMult(42, input)
		out := ncc(input)
		var result [n]int16
		for i := 0; i < int(ndiv8); i++ {
			for j := 0; j < 8; j++ {
				result[8*i+j] = out[i][j]
			}
		}

		for i := int16(0); i < n; i++ {
			util.Mod_257(&result[i])
		}

		if result != want {
			t.Fatalf("(Test ntt_part) Bad FFT: %v != %v", result, want)
		}

	}
}

func NCCVecMult(omega int16, input [ndiv8]byte) [n]int16 {

	var product [n]int16
	var vec [n]int16
	mat := genNCCMat(omega)

	for i := int16(0); i < ndiv8; i++ {
		t_vec := bit2ByteTable[input[i]]

		for j := int16(0); j < 8; j++ {
			vec[8*i+j] = t_vec[j]
		}
	}

	for i := int16(0); i < n; i++ {
		for j := int16(0); j < n; j++ {
			product[i] = product[i] + mat[i][j]*vec[j]
		}
	}

	for i := int16(0); i < n; i++ {
		util.Mod_257(&product[i])
	}

	return product
}

// Returns NCC matrix for bitreversed output
func genNCCMat(omega int16) [n][n]int16 {

	var ncc_mat [n][n]int16

	for i := int16(0); i < n; i++ {
		for k := int16(0); k < n; k++ {
			if (k*(2*i+1))%(2*n) <= n {
				ncc_mat[i][k] = util.IntPow(omega, (k*(2*i+1))%n, q)

			} else {
				ncc_mat[i][k] = -1 * util.IntPow(omega, ((k*(2*i+1))%n), q)

			}
		}

	}

	var br_ncc_mat [n][n]int16
	for i := int16(0); i < n; i++ {
		for j := int16(0); j < n; j++ {
			br_ncc_mat[j][util.Bit_Rev(i, n)] = ncc_mat[j][i]
		}
	}

	return br_ncc_mat
}

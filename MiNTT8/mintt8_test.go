package MiNTT8

import (
	util "github.com/RileyVaughn/MiNTT/hash/util"
)

const TEST_SIZE = 1000

// func TestNTT_part(t *testing.T) {

// 	seed, _ := strconv.Atoi("MiNNT")
// 	rand.Seed(int64(seed))

// 	for i := 0; i < TEST_SIZE; i++ {

// 		input := byte(rand.Intn(256))

// 		want := NCCVecMult(2, input)
// 		result := ntt_part(input)
// 		for i := 0; i < n; i++ {
// 			result[i] = util.Mod(result[i], q)
// 		}

// 		if result != want {
// 			t.Fatalf("(Test ntt_part) Bad FFT: %v != %v", result, want)
// 		}
// 	}
// }

func NCCVecMult(omega int, input byte) [n]int {

	var product [n]int
	var vec [n]int
	mat := genNCCMat(omega)

	vec = bit2ByteTable[input]

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
	var br_arr [n]int = [n]int{0, 4, 2, 6, 1, 5, 3, 7} //dirty
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			br_ncc_mat[j][br_arr[i]] = ncc_mat[j][i]
		}

	}

	return br_ncc_mat
}

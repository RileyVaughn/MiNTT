package MiNTT128

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	util "github.com/RileyVaughn/MiNTT/hash/int16/util"
)

const KEY_PATH string = "./int16/normal/MiNTT128/key.csv"

// Assumes key is eactly M x (n*d)
func ReadKey(filepath string) [m][d][ndiv8][8]int16 {

	var key [m][d * n]int16

	fi, _ := os.Open(filepath)
	r := csv.NewReader(fi)
	keystring, _ := r.ReadAll()

	for i := range keystring {
		for j := range keystring[i] {
			num, _ := strconv.Atoi(keystring[i][j])
			key[i][j] = int16(num)
		}
	}

	// converts format
	var simd_key [m][d][ndiv8][8]int16
	for i := int16(0); i < m; i++ {
		for j := int16(0); j < d; j++ {
			for k := int16(0); k < ndiv8; k++ {
				for l := int16(0); l < 8; l++ {
					simd_key[i][j][k][l] = key[i][(n*j)+(8*k)+l]
				}
			}
		}
	}

	return simd_key
}

func NTT8Table(omega int16) [256][8]int16 {

	var table [256][8]int16
	var ncc8Mat = gen8NCCMat(omega)

	bit2ByteTable := util.BitsFromByteTable()
	for j := 0; j < 256; j++ {
		var product [8]int16
		vec := bit2ByteTable[j]

		for i := 0; i < 8; i++ {
			for k := 0; k < 8; k++ {
				product[i] = util.Center(product[i]+ncc8Mat[i][k]*vec[k], q)

			}
		}
		table[j] = product
	}

	return table
}

func MultTable(omega int16) [16][8]int16 {

	var table [16][8]int16

	for i := int16(0); i < 8; i++ {
		for k := int16(0); k < 16; k++ {
			table[k][i] = util.Center(util.IntPow(omega, k*(2*i+1), q), q)
		}
	}

	//Bit reverse the k index
	var br_table [16][8]int16
	for k := int16(0); k < 16; k++ {
		for i := int16(0); i < 8; i++ {
			br_table[util.Bit_Rev(k, 16)][i] = table[k][i]
		}
	}

	return br_table
}

func gen8NCCMat(omega int16) [8][8]int16 {

	var ncc_mat [8][8]int16
	for i := int16(0); i < 8; i++ {
		for k := int16(0); k < 8; k++ {
			pow := k * (2*i + 1)
			if (pow)%(2*8) < 8 {
				ncc_mat[i][k] = util.IntPow(omega, pow%8, q)
			} else {
				ncc_mat[i][k] = -1 * util.IntPow(omega, (pow%8), q)
			}
		}

	}
	var br_ncc_mat [8][8]int16
	for i := int16(0); i < 8; i++ {
		for j := int16(0); j < 8; j++ {
			br_ncc_mat[j][util.Bit_Rev(i, 8)] = ncc_mat[j][i]
		}

	}
	return br_ncc_mat
}

func SetupM128() {

	if _, err := os.Stat(KEY_PATH); errors.Is(err, os.ErrNotExist) {
		util.GenWriteKey(m, n, d, q, KEY_PATH)
	}

	A = ReadKey(KEY_PATH)

	NTT8_TABLE = NTT8Table(2)

	MULT_TABLE = MultTable(27)

	fmt.Println("Setup Finished")
}

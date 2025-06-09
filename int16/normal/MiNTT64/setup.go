package MiNTT64

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	util "github.com/RileyVaughn/MiNTT/hash/int16/util"
)

const KEY_PATH string = ".int16/normal/MiNTT64/key.csv"

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
	for i := 0; i < 256; i++ {
		var product [8]int16
		vec := bit2ByteTable[i]
		for j := 0; j < 8; j++ {
			for k := 0; k < 8; k++ {
				product[j] = util.Center(product[j]+ncc8Mat[j][k]*vec[k], q)
			}
		}
		table[i] = product
	}

	return table
}

func MultTable(omega int16) [8][8]int16 {

	var table [8][8]int16

	for i0 := int16(0); i0 < 8; i0++ {
		for k0 := int16(0); k0 < 8; k0++ {
			table[k0][i0] = util.Center(util.IntPow(omega, util.Mod(util.Bit_Rev(k0, 8)*(2*i0+1), 2*n), q), q)
		}
	}

	return table
}

func gen8NCCMat(omega int16) [8][8]int16 {

	var ncc_mat [8][8]int16
	for i := int16(0); i < 8; i++ {
		for k := int16(0); k < 8; k++ {
			if (k*(2*i+1))%(2*8) <= 8 {
				ncc_mat[i][k] = util.IntPow(omega, (k*(2*i+1))%8, q)
			} else {
				ncc_mat[i][k] = -1 * util.IntPow(omega, ((k*(2*i+1))%8), q)
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

func SetupM64() {

	if _, err := os.Stat(KEY_PATH); errors.Is(err, os.ErrNotExist) {
		util.GenWriteKey(m, n, d, q, KEY_PATH)
	}

	A = ReadKey(KEY_PATH)

	NTT8_TABLE = NTT8Table(2)
	MULT_TABLE = MultTable(42)

	fmt.Println("Setup Finished")
}

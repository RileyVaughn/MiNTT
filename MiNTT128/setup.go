package MiNTT128

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	util "github.com/RileyVaughn/MiNTT/hash/util"
)

const KEY_PATH string = "./MiNTT128/key.csv"

//Assumes key is eactly M x (n*d)
func ReadKey(filepath string) [m][d * n]int {

	var key [m][d * n]int

	fi, _ := os.Open(filepath)
	r := csv.NewReader(fi)
	keystring, _ := r.ReadAll()

	for i := range keystring {
		for j := range keystring[i] {
			num, _ := strconv.Atoi(keystring[i][j])
			key[i][j] = int(num)
		}
	}
	return key
}

func NTT8Table(omega int) [256][8]int {

	var table [256][8]int
	var ncc8Mat = gen8NCCMat(omega)

	bit2ByteTable := util.BitsFromByteTable()
	for j := 0; j < 256; j++ {
		var product [8]int
		vec := bit2ByteTable[j]

		for i := 0; i < 8; i++ {
			for k := 0; k < 8; k++ {
				product[i] = util.Mod(product[i]+ncc8Mat[i][k]*vec[k], q)

			}
		}
		table[j] = product
	}

	return table
}

func MultTable(omega int) [16][8]int {

	var table [16][8]int

	for i := 0; i < 8; i++ {
		for k := 0; k < 16; k++ {
			table[k][i] = util.IntPow(omega, k*(2*i+1), q)
		}
	}

	//Bit reverse the k index
	var br_table [16][8]int
	for k := 0; k < 16; k++ {
		for i := 0; i < 8; i++ {
			br_table[util.Bit_Rev(k, 16)][i] = table[k][i]
		}
	}

	return br_table
}

func gen8NCCMat(omega int) [8][8]int {

	var ncc_mat [8][8]int
	for i := 0; i < 8; i++ {
		for k := 0; k < 8; k++ {
			pow := k * (2*i + 1)
			if (pow)%(2*8) < 8 {
				ncc_mat[i][k] = util.IntPow(omega, pow%8, q)
			} else {
				ncc_mat[i][k] = -1 * util.IntPow(omega, (pow%8), q)
			}
		}

	}
	var br_ncc_mat [8][8]int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
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

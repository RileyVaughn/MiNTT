package MiNTT64

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	util "github.com/RileyVaughn/MiNTT/hash/util"
)

const KEY_PATH string = "./MiNTT64/key.csv"

//Seed rng with MiNNT
//returns 2d string slice length M x (n*d)
func tempKeyGen() [][]string {
	seed, _ := strconv.Atoi("MiNNT")
	rand.Seed(int64(seed))

	key := make([][]string, m)
	for i := range key {
		keyrow := make([]string, n*d)
		for j := range keyrow {
			keyrow[j] = strconv.Itoa(rand.Intn(int(q)))
		}
		key[i] = keyrow
	}

	return key
}

func GenWriteKey(filepath string) {

	key := tempKeyGen()

	fo, _ := os.Create(filepath)
	w := csv.NewWriter(fo)
	w.WriteAll([][]string(key))

}

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
	for i := 0; i < 256; i++ {
		var product [8]int
		vec := bit2ByteTable[i]
		for j := 0; j < 8; j++ {
			for k := 0; k < 8; k++ {
				product[j] = util.Mod(product[j]+ncc8Mat[j][k]*vec[k], q)
			}
		}
		table[i] = product
	}

	return table
}

func MultTable(omega int) [8][8]int {

	var table [8][8]int

	for i0 := 0; i0 < 8; i0++ {
		for k0 := 0; k0 < 8; k0++ {
			table[k0][i0] = util.IntPow(omega, util.Mod(util.Bit_Rev(k0, 8)*(2*i0+1), 2*n), q)
		}
	}

	return table
}

func gen8NCCMat(omega int) [8][8]int {

	var ncc_mat [8][8]int
	for i := 0; i < 8; i++ {
		for k := 0; k < 8; k++ {
			if (k*(2*i+1))%(2*8) <= 8 {
				ncc_mat[i][k] = util.IntPow(omega, (k*(2*i+1))%n)
			} else {
				ncc_mat[i][k] = -1 * util.IntPow(omega, ((k*(2*i+1))%n))
			}
		}

	}
	var br_ncc_mat [8][8]int
	var br_arr [8]int = [8]int{0, 4, 2, 6, 1, 5, 3, 7}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			br_ncc_mat[j][br_arr[i]] = ncc_mat[j][i]
		}

	}
	return br_ncc_mat
}

func SetupM64() {

	if _, err := os.Stat(KEY_PATH); errors.Is(err, os.ErrNotExist) {
		GenWriteKey(KEY_PATH)
	}

	A = ReadKey(KEY_PATH)

	NTT8_TABLE = NTT8Table(42)
	MULT_TABLE = MultTable(42)

	fmt.Println("Setup Finished")
}

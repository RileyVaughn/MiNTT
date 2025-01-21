package util

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
)

func WriteIntCSV(filepath string, nums [][]int) {

	var records [][]string

	for i := range nums {
		record := []string{}
		for j := range nums[i] {
			record = append(record, strconv.Itoa(nums[i][j]))
		}
		records = append(records, record)
	}
	fo, _ := os.Create(filepath)
	w := csv.NewWriter(fo)
	w.WriteAll(records)

}

func ReadIntCSV(filepath string) [][]int {

	var nums [][]int

	fi, _ := os.Open(filepath)
	r := csv.NewReader(fi)
	records, _ := r.ReadAll()

	for i := range records {
		numline := []int{}
		for j := range records[i] {
			num, _ := strconv.Atoi(records[i][j])
			numline = append(numline, num)
		}
		nums = append(nums, numline)
	}

	return nums
}

//Generated a 2d slice of random ints from 0-1000
func GenRandInt2d(rowlen int, collen int, max int) [][]int {

	var records [][]int

	for i := 0; i < rowlen; i++ {
		record := []int{}
		for j := 0; j < collen; j++ {
			record = append(record, rand.Intn(max))
		}
		records = append(records, record)
	}

	return records
}

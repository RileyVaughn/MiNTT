package polynom

import (
	"testing"

	c "github.com/RileyVaughn/MiNTT/ineff/constant"
	"github.com/RileyVaughn/MiNTT/ineff/util"
)

func TestIsEqual(t *testing.T) {

	polyTest := ReadPolys("polynomials.csv")

	var wantTest []bool = []bool{true, false}

	result := polyTest[0].IsEqual(polyTest[0])
	if result != wantTest[0] {
		t.Fatalf("(TestIsEqual) Bad Equiv: %v != %v", result, wantTest[0])
	}

	result = polyTest[0].IsEqual(polyTest[1])
	if result != wantTest[1] {
		t.Fatalf("(TestIsEqual) Bad Equiv: %v != %v", result, wantTest[1])
	}

}

func TestMult(t *testing.T) {

	polyTest := ReadPolys("polynomials.csv")
	wantTest := ReadPolys("poly_mult_ans.csv")

	for i := 0; i < 25; i++ {
		result := polyTest[i].Mult(polyTest[i+25])
		if !(result.IsEqual(wantTest[i])) {
			t.Fatalf("(TestMult) Bad Math: %v != %v", result, wantTest[i])
		}
	}

	for i := 0; i < 25; i++ {
		pt := polyTest[i+25]
		for j := 0; j < c.N; j++ {
			pt[j] = pt[j] % 2
		}
		result := polyTest[i].Mult(pt)
		if !(result.IsEqual(wantTest[i+25])) {
			t.Fatalf("(TestMult) Bad Binary Math, Test (%v): %v != %v", i, result, wantTest[i+25])
		}
	}
}

func TestAdd(t *testing.T) {

	polyTest := ReadPolys("polynomials.csv")
	wantTest := ReadPolys("poly_add_ans.csv")

	for i := 0; i < 25; i++ {
		result := polyTest[i].Add(polyTest[i+25])
		if !(result.IsEqual(wantTest[i])) {
			t.Fatalf("(TestMult) Bad Math: %v != %v", result, wantTest[i])
		}
	}

}

func ReadPolys(filename string) []Polynom {
	coefs := util.ReadIntCSV(filename)
	var polyTest []Polynom
	for i := range coefs {
		var coefs_array [c.N]int
		for j := 0; j < c.N; j++ {
			coefs_array[j] = coefs[i][j]
		}

		polyTest = append(polyTest, Polynom(coefs_array))
	}
	return polyTest
}

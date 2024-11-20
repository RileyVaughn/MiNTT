package main

import "fmt"

// Each index i represnts the powers of x, i.e. [1,2,3,4] -> 1 + 2x + 3x^2 + 4x^3
type polynomial []int

//Creates a new polynomial of specified length n
func NewPolynomial(n int) polynomial {
	var p polynomial = make(polynomial, n)

	return p
}

// Add is the simple addition of two polynomials
func (p1 polynomial) Add(p2 polynomial) polynomial {

	p1, p2 = forceSameLength(p1, p2)
	for i := 0; i < len(p1); i++ {
		p1[i] = p1[i] + p2[i]
	}
	return p1
}

// Scales multipliers a scalar by the polynomial
func (p polynomial) Scale(s int) polynomial {
	for i := 0; i < len(p); i++ {
		p[i] = p[i] * s
	}
	return p
}

// Mult is the naive multiplication of polynomials without NTT
func (p1 polynomial) PolyMult(p2 polynomial) polynomial {

	p1, p2 = forceSameLength(p1, p2)
	polyLen := len(p1)
	p3 := NewPolynomial(2 * polyLen)

	for i := 0; i < polyLen; i++ {
		for j := 0; j < polyLen; j++ {
			p3[i+j] = p3[i+j] + p1[i]*p2[j]
		}

	}

	return simplify(p3)
}

//forceSameLength adds traling 0's to the shorter polynomial so both lengths match
func forceSameLength(p1 polynomial, p2 polynomial) (polynomial, polynomial) {

	for len(p1) != len(p2) {
		if len(p1) < len(p2) {
			p1 = append(p1, 0)
		} else {
			p2 = append(p2, 0)
		}
	}

	return p1, p2
}

//simplify removes traliing 0's
func simplify(p polynomial) polynomial {

	for p[len(p)-1] == 0 {
		p = p[:len(p)-1]
	}

	return p
}

func main() {

	p1 := polynomial{1, 2, 3}
	p2 := polynomial{4, 5, 0, 6}

	p1 = p1.PolyMult(p2)
	fmt.Println(p1)

}

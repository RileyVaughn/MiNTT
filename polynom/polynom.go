package polynom

import (
	num "github.com/RileyVaughn/MiNTT/numbers"
)

// Polynom is a polynomial of length N.
// Each index i represnts the power of x, i.e. [1,2,3,4] -> 1 + 2x + 3x^2 + 4x^3
// All methods of Polynom assume both polynoms are the same length.
type Polynom []num.Number

// Add is the simple addition of two polynomials.
func (p1 Polynom) Add(p2 Polynom) Polynom {

	for i := 0; i < len(p1); i++ {
		p1[i] = p1[i].Add(p2[i])
	}
	return p1
}

// Scales multipliers a scalar by the Polynom
func (p Polynom) Scale(s num.Number) Polynom {
	for i := 0; i < len(p); i++ {
		p[i] = p[i].Mult(s)
	}
	return p
}

// PolyMult is the naive multiplication of polynomials without NTT
func (p1 Polynom) PolyMult(p2 Polynom) Polynom {

	polyLen := len(p1)
	p3 := make(Polynom, 2*polyLen)

	for i := 0; i < polyLen; i++ {
		for j := 0; j < polyLen; j++ {
			p3[i+j] = p3[i+j].Add(p1[i].Mult(p2[j]))
		}

	}

	return p3
}

// PolyMultModXnplus1 is naive polynomial multiplication mod X^n+1
// n is predefined as the shared length of the polynomials p1 and p2
func (p1 Polynom) PolyMultModXnplus1(p2 Polynom) Polynom {

	p3 := p1.PolyMult(p2)
	n := len(p1)
	p4 := make(Polynom, n)

	for i := 0; i < n; i++ {
		p4[i] = p3[i].Sub(p3[i+n])
	}

	return p4
}

// Checks if the smae length and contain the same elements
func (p1 Polynom) IsEqual(p2 Polynom) bool {

	isequal := true

	if len(p1) == len(p2) {
		for i := 0; i < len(p1); i++ {
			if !p1[i].IsEqual(p2[i]) {
				isequal = false
			}
		}
	} else {
		isequal = false
	}

	return isequal
}

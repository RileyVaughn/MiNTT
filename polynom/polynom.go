package polynom

import num "github.com/RileyVaughn/MiNTT/numbers"

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

// Mult is the naive multiplication of polynomials without NTT
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

func (p1 Polynom) Len() int {
	return len([]num.Number(p1))
}

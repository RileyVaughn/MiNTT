package poly

// Each index i represnts the power of x, i.e. [1,2,3,4] -> 1 + 2x + 3x^2 + 4x^3
type Polynomial []int

//Creates a new Polynomial of specified length n
func NewPolynomial(n int) Polynomial {
	var p Polynomial = make(Polynomial, n)

	return p
}

// Add is the simple addition of two Polynomials
func (p1 Polynomial) Add(p2 Polynomial) Polynomial {

	p1, p2 = forceSameLength(p1, p2)
	for i := 0; i < len(p1); i++ {
		p1[i] = p1[i] + p2[i]
	}
	return p1
}

// Scales multipliers a scalar by the Polynomial
func (p Polynomial) Scale(s int) Polynomial {
	for i := 0; i < len(p); i++ {
		p[i] = p[i] * s
	}
	return p
}

// Mult is the naive multiplication of Polynomials without NTT
func (p1 Polynomial) PolyMult(p2 Polynomial) Polynomial {

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

//forceSameLength adds trailing 0's to the shorter Polynomial so both lengths match
func forceSameLength(p1 Polynomial, p2 Polynomial) (Polynomial, Polynomial) {

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
func simplify(p Polynomial) Polynomial {

	for p[len(p)-1] == 0 {
		p = p[:len(p)-1]
	}

	return p
}

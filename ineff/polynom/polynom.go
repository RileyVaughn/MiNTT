package polynom

import (
	"fmt"

	c "github.com/RileyVaughn/MiNTT/ineff/constant"
)

// Polynom is a polynomial element of the set Z_q/(X^n+1)
// Each index i represents the power of x, i.e. [1,2,3,4] -> 1 + 2x + 3x^2 + 4x^3
// All methods of Polynom assume both polynoms share the values q and n (in the efficient version, these values will likely be constants).
type Polynom []int

// Add is the simple addition of two polynomials.
func (p1 Polynom) Add(p2 Polynom) Polynom {

	for i := 0; i < c.N; i++ {
		p1[i] = Mod((p1[i] + p2[i]), c.Q)
	}
	return p1
}

// PolyMult is the naive multiplication of polynomials in Z_q/(X^n+1) without NTT
func (p1 Polynom) Mult(p2 Polynom) Polynom {

	p3 := make(Polynom, c.N)
	fmt.Println(p3)

	for i := 0; i < c.N; i++ {
		for j := 0; j < c.N; j++ {
			k := (i + j) % c.N
			if k == (i + j) {
				p3[k] = Mod(p3[k]+(p1[i]*p2[j]), c.Q)
			} else {
				p3[k] = Mod(p3[k]-(p1[i]*p2[j]), c.Q)
			}
		}
	}

	return p3
}

// Checks if the smae length and contain the same elements
func (p1 Polynom) IsEqual(p2 Polynom) bool {

	isequal := true

	for i := 0; i < c.N; i++ {
		if !(p1[i] == p2[i]) {
			isequal = false
		}
	}

	return isequal
}

//Golang mod of a negative returns a negative, this returns the positive
func Mod(a int, b int) int {
	return ((a % b) + b) % b
}

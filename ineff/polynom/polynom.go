package polynom

// Polynom is a polynomial element of the set Z_q/(X^n+1)
// Each index i represents the power of x, i.e. [1,2,3,4] -> 1 + 2x + 3x^2 + 4x^3
// All methods of Polynom assume both polynoms share the values q and n (in the efficient version, these values will likely be constants).
type Polynom []int

const POLY_Q int = 7681
const POLY_N = 256

// Add is the simple addition of two polynomials.
func (p1 Polynom) Add(p2 Polynom) Polynom {

	for i := 0; i < POLY_N; i++ {
		p1[i] = ((p1[i] + p2[i]) + POLY_Q) % POLY_Q
	}
	return p1
}

// PolyMult is the naive multiplication of polynomials in Z_q/(X^n+1) without NTT
func (p1 Polynom) Mult(p2 Polynom) Polynom {

	p3 := make(Polynom, POLY_N)

	for i := 0; i < POLY_N; i++ {
		for j := 0; j < POLY_N; j++ {
			k := (i + j) % POLY_N
			if k == (i + j) {
				p3[k] = p3[k] + (p1[i] * p2[j])
			} else {
				p3[k] = p3[k] - (p1[i] * p2[j])
			}
		}
	}

	return p3
}

// // Checks if the smae length and contain the same elements
// func (p1 Polynom) IsEqual(p2 Polynom) bool {

// 	isequal := true

// 	if len(p1) == len(p2) {
// 		for i := 0; i < len(p1); i++ {
// 			if !p1[i].IsEqual(p2[i]) {
// 				isequal = false
// 			}
// 		}
// 	} else {
// 		isequal = false
// 	}

// 	return isequal
// }

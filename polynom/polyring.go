package polynom

import "fmt"

// A PolyRing is a polynomial ring element with 'Number' coefficients.
// While slcies are used in the implimintation, the polyring cannot change degree.
type PolyRing struct {
	polynom       Polynom
	ideal_polynom Polynom //Guaranteed to be same size as polynom
	max_degree    int
}

// "Guarantees" that polynoms are the same degree as max degree
func MakePolyRing(polynom Polynom, ideal_polynom Polynom, max_degree int) PolyRing {

	if polynom.Len() != ideal_polynom.Len() || polynom.Len() != max_degree {
		fmt.Println("MakePolyRing ERROR: Sizes are not the same")
	}

	pr := new(PolyRing)
	pr.polynom = polynom
	pr.ideal_polynom = ideal_polynom
	pr.max_degree = max_degree

	return *pr
}

func (pr PolyRing) GetValue() Polynom {
	return pr.polynom
}

func (pr PolyRing) GetIdeal() Polynom {
	return pr.ideal_polynom
}

func (pr PolyRing) GetMaxDegree() int {
	return pr.max_degree
}

// Add is the simple addition of two PolyRings.
// An error is thrown if
func (p1 PolyRing) Add(p2 PolyRing) PolyRing {

	// for i := 0; i < p1.max_degree; i++ {
	// 	p1.polynom[i] = p1.polynom[i].Add(p2.polynom[i])
	// }
	// return p1
}

// // Scales multipliers a scalar by the Polynom
// func (p Polynom) Scale(s int) Polynom {
// 	for i := 0; i < len(p); i++ {
// 		p[i] = p[i] * s
// 	}
// 	return p
// }

// // Mult is the naive multiplication of Polynoms without NTT
// func (p1 Polynom) PolyMult(p2 Polynom) Polynom {

// 	p1, p2 = forceSameLength(p1, p2)
// 	polyLen := len(p1)
// 	p3 := NewPolynom(2 * polyLen)

// 	for i := 0; i < polyLen; i++ {
// 		for j := 0; j < polyLen; j++ {
// 			p3[i+j] = p3[i+j] + p1[i]*p2[j]
// 		}

// 	}

// 	return simplify(p3)
// }

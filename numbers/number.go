package numbers

//Number is an interface designed so polynomials can use a a variety of coefficients
type Number interface {
	Add(Number) Number
	Sub(Number) Number
	Mult(Number) Number
	// LessThan(Number) bool
	// IsEqual(Number) bool
}

package numbers

//Number is an interface designed so polynomials can use a a variety of coefficients
type number interface {
	Add() number
	Sub() number
	Mult() number
}

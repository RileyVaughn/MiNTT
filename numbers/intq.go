package numbers

// Intq is an integer modulo q.
// All calculations return results n in [0,q-1]
// The values n and q are implememnted as 64bit integers.
type Intq struct {
	Z int64
	Q int64
}

// Add is the simple addition of two integers modulo q
// Use q of n1, regardless of n2.q
func (n1 Intq) Add(n2 Intq) Intq {

	return Intq{Z: (n1.Z + n2.Z) % n1.Q, Q: n1.Q}
}

// Sub is the simple subtraction of two integers modulo q
// Use q of n1, regardless of n2.q
func (n1 Intq) Sub(n2 Intq) Intq {

	return Intq{Z: (n1.Z - n2.Z) % n1.Q, Q: n1.Q}
}

// Mult is the simple multiplication of two integers modulo q
func (n1 Intq) Mult(n2 Intq) Intq {

	return Intq{Z: (n1.Z * n2.Z) % n1.Q, Q: n1.Q}
}

// func (n1 Intq) LessThan(n2 Intq) bool {

// 	return n1.Z < n2.Z
// }

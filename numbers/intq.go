package numbers

// Intq is an integer modulo q.
// All calculations return results n in [0,q-1]
// The values n and q are implememnted as 64bit integers.
type Intq struct {
	z int64
	q int64
}

// Add is the simple addition of two integers modulo q
// Use q of n1, regardless of n2.q
func (n1 Intq) Add(n2 Intq) Intq {

	return Intq{z: (n1.z + n2.z) % n1.q, q: n1.q}
}

// Sub is the simple subtraction of two integers modulo q
// Use q of n1, regardless of n2.q
func (n1 Intq) Sub(n2 Intq) Intq {

	return Intq{z: (n1.z - n2.z) % n1.q, q: n1.q}
}

// Mult is the simple multiplication of two integers modulo q
func (n1 Intq) Mult(n2 Intq) Intq {

	return Intq{z: (n1.z * n2.z) % n1.q, q: n1.q}
}

package numbers

import (
	"testing"
)

//Assume that Q will never be zero
//Just testing that all will

func TestAdd(t *testing.T) {

	var z1Test []int64 = []int64{6, 0, 9, 9, 2}
	var z2Test []int64 = []int64{4, 0, 9, 8, 6}
	var QTest []int64 = []int64{9, 9, 9, 9, 9}
	var wantTest []int64 = []int64{1, 0, 0, 8, 8}

	var z1 Intq
	var z2 Intq
	var want int64

	for i := 0; i < len(z1Test); i++ {

		z1 = Intq{Z: z1Test[i], Q: QTest[i]}
		z2 = Intq{Z: z2Test[i], Q: QTest[i]}
		want = wantTest[i]
	}

	result := z1.Add(z2).Z

	if result != want {
		t.Fatalf("(TestAdd) Bad Math: %v != %v", result, want)
	}

}

func TestSub(t *testing.T) {

	var z1Test []int64 = []int64{6, 0, 9, 9, 2}
	var z2Test []int64 = []int64{4, 0, 9, 8, 6}
	var QTest []int64 = []int64{9, 9, 9, 9, 9}
	var wantTest []int64 = []int64{1, 0, 0, 8, 5}

	var z1 Intq
	var z2 Intq
	var want int64

	for i := 0; i < len(z1Test); i++ {

		z1 = Intq{Z: z1Test[i], Q: QTest[i]}
		z2 = Intq{Z: z2Test[i], Q: QTest[i]}
		want = wantTest[i]
	}

	result := z1.Sub(z2).Z

	if result != want {
		t.Fatalf("(TestSub) Bad Math: %v != %v", result, want)
	}

}

func TestMult(t *testing.T) {

	var z1Test []int64 = []int64{6, 0, 9, 9, 2}
	var z2Test []int64 = []int64{4, 0, 9, 8, 6}
	var QTest []int64 = []int64{9, 9, 9, 9, 9}
	var wantTest []int64 = []int64{6, 0, 0, 0, 3}

	var z1 Intq
	var z2 Intq
	var want int64

	for i := 0; i < len(z1Test); i++ {

		z1 = Intq{Z: z1Test[i], Q: QTest[i]}
		z2 = Intq{Z: z2Test[i], Q: QTest[i]}
		want = wantTest[i]
	}

	result := z1.Mult(z2).Z

	if result != want {
		t.Fatalf("(TestMult) Bad Math: %v != %v", result, want)
	}

}

func TestIsEqual(t *testing.T) {

	var z1Test []int64 = []int64{6, 0, 9, 9, 5, 17}
	var z2Test []int64 = []int64{4, 0, 9, 18, 5, 7}
	var QTest []int64 = []int64{9, 9, 9, 9, 9, 10}
	var wantTest []bool = []bool{false, true, true, true, false, true}

	var z1 Intq
	var z2 Intq
	var want bool

	for i := 0; i < len(z1Test); i++ {

		z1 = Intq{Z: z1Test[i], Q: QTest[i]}
		z2 = Intq{Z: z2Test[i], Q: QTest[i]}
		want = wantTest[i]
	}

	result := z1.IsEqual(z2)

	if result != want {
		t.Fatalf("(TestIsEqual) Bad Math: %v != %v", result, want)
	}

}

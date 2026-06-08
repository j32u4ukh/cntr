package cntr

import "testing"

func TestBinaryDataRoundTrip(t *testing.T) {
	bd := NewBinaryData()
	if err := bd.AddInt8(-1); err != nil {
		t.Fatal(err)
	}
	if err := bd.AddUInt16(1000); err != nil {
		t.Fatal(err)
	}
	if err := bd.AddBoolean(true); err != nil {
		t.Fatal(err)
	}
	if err := bd.AddMapInt64Int64Array(map[int64][]int64{1: {2, 3}}); err != nil {
		t.Fatal(err)
	}
	if err := bd.AddNumber(int(42)); err != nil {
		t.Fatal(err)
	}
	if err := bd.AddNumber(float64(3.14)); err != nil {
		t.Fatal(err)
	}

	i8, err := bd.PopInt8()
	if err != nil || i8 != -1 {
		t.Fatalf("PopInt8: got %d, err %v", i8, err)
	}
	u16, err := bd.PopUInt16()
	if err != nil || u16 != 1000 {
		t.Fatalf("PopUInt16: got %d, err %v", u16, err)
	}
	ok, err := bd.PopBoolean()
	if err != nil || !ok {
		t.Fatalf("PopBoolean: got %v, err %v", ok, err)
	}
	m, err := bd.PopMapInt64Int64Array()
	if err != nil || len(m) != 1 || len(m[1]) != 2 {
		t.Fatalf("PopMapInt64Int64Array: got %v, err %v", m, err)
	}
	v, err := bd.PopScalar(int(0))
	if err != nil || v.(int) != 42 {
		t.Fatalf("PopScalar int: got %v, err %v", v, err)
	}
	f, err := bd.PopScalar(float64(0))
	if err != nil || f.(float64) != 3.14 {
		t.Fatalf("PopScalar float64: got %v, err %v", f, err)
	}
}

package test

import (
	"testing"

	"github.com/j32u4ukh/cntr"
	"github.com/pkg/errors"
)

func BM1() (*cntr.BikeyMap[string, int32, float32], error) {
	var err error
	bm := cntr.NewBikeyMap[string, int32, float32]()
	err = bm.Add("a", 1, 1.0)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to add bikey-value.")
	}
	err = bm.Add("b", 2, 1.414)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to add bikey-value.")
	}
	err = bm.Add("c", 3, 1.71)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to add bikey-value.")
	}
	return bm, nil
}

func TestGet(t *testing.T) {
	var value, answer float32
	var ok bool
	bm, err := BM1()
	if err != nil {
		t.Errorf("BM1 | Error: %+v\n", err)
	}
	// (a, 1)
	answer = 1.0
	value, ok = bm.GetByKey1("a")
	if (ok == false) || (value != answer) {
		t.Errorf("TestGet | Error: %+v\n", err)
	}
	value, ok = bm.GetByKey2(1)
	if (ok == false) || (value != answer) {
		t.Errorf("TestGet | Error: %+v\n", err)
	}
	// (b, 2)
	answer = 1.414
	value, ok = bm.GetByKey1("b")
	if (ok == false) || (value != answer) {
		t.Errorf("TestGet | Error: %+v\n", err)
	}
	value, ok = bm.GetByKey2(2)
	if (ok == false) || (value != answer) {
		t.Errorf("TestGet | Error: %+v\n", err)
	}
	// (c, 3)
	answer = 1.71
	value, ok = bm.GetByKey1("c")
	if (ok == false) || (value != answer) {
		t.Errorf("TestGet | Error: %+v\n", err)
	}
	value, ok = bm.GetByKey2(3)
	if (ok == false) || (value != answer) {
		t.Errorf("TestGet | Error: %+v\n", err)
	}
}

func TestContain(t *testing.T) {
	var ok bool
	bm, err := BM1()
	if err != nil {
		t.Errorf("BM1 | Error: %+v\n", err)
	}
	ok = bm.ContainKey1("a")
	if ok != true {
		t.Errorf("TestContain | Error: %+v\n", err)
	}
	ok = bm.ContainKey2(1)
	if ok != true {
		t.Errorf("TestContain | Error: %+v\n", err)
	}
	ok = bm.ContainKey1("A")
	if ok != false {
		t.Errorf("TestContain | Error: %+v\n", err)
	}
	ok = bm.ContainKey2(100)
	if ok != false {
		t.Errorf("TestContain | Error: %+v\n", err)
	}
}

func TestUpdate(t *testing.T) {
	var value, answer float32
	var ok bool
	bm, err := BM1()
	if err != nil {
		t.Errorf("BM1 | Error: %+v\n", err)
	}
	// (A, 10)
	answer = 10.0
	bm.UpdateByKey1("a", cntr.NewBivalue[string, int32, float32]("A", 10, answer))
	value, ok = bm.GetByKey1("A")
	if (ok == false) || (value != answer) {
		t.Errorf("TestUpdate | Error: %+v\n", err)
	}
	value, ok = bm.GetByKey2(10)
	if (ok == false) || (value != answer) {
		t.Errorf("TestUpdate | Error: %+v\n", err)
	}
	// (B, 20)
	answer = 14.14
	bm.UpdateByKey1("b", cntr.NewBivalue[string, int32, float32]("B", 20, answer))
	value, ok = bm.GetByKey1("B")
	if (ok == false) || (value != answer) {
		t.Errorf("TestUpdate | Error: %+v\n", err)
	}
	value, ok = bm.GetByKey2(20)
	if (ok == false) || (value != answer) {
		t.Errorf("TestUpdate | Error: %+v\n", err)
	}
	// (c, 3)
	answer = 17.1
	bm.UpdateByKey1("c", cntr.NewBivalue[string, int32, float32]("C", 30, answer))
	value, ok = bm.GetByKey1("C")
	if (ok == false) || (value != answer) {
		t.Errorf("TestUpdate | Error: %+v\n", err)
	}
	value, ok = bm.GetByKey2(30)
	if (ok == false) || (value != answer) {
		t.Errorf("TestUpdate | Error: %+v\n", err)
	}
}

func TestDelete(t *testing.T) {
	var ok bool
	bm, err := BM1()
	if err != nil {
		t.Errorf("BM1 | Error: %+v\n", err)
	}
	bm.DelByKey1("a")
	_, ok = bm.GetByKey1("a")
	if ok != false {
		t.Error("TestDelete | DelByKey1 failed\n")
	}
	bm.DelByKey2(2)
	_, ok = bm.GetByKey2(2)
	if ok != false {
		t.Error("TestDelete | DelByKey2 failed\n")
	}
}

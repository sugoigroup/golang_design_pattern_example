package factory

import "testing"

func TestFactory(t *testing.T) {
	bmw := makeBMW()
	if bmw.getName() != "BMW" {
		t.Error("BMW 제조공장 에러")
	}

	audi := makeAUDI()
	if audi.whichLine.getName() != "AUDI" {
		t.Error("AUDI 제조공장 에러")
	}
}

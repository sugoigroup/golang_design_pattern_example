package bridge

import "testing"

func TestFactory(t *testing.T) {
	mac := &mac{}

	epson := &epson{}
	mac.setPrinter(epson)

	if mac.print() != "epson" {
		t.Error("epson 프린터 오류")
	}
	samsung := &samsung{}
	mac.setPrinter(samsung)

	if mac.print() != "samsung" {
		t.Error("samsung 프린터 오류")
	}

}

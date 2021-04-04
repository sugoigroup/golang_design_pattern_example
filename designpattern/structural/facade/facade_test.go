package facade

import "testing"

func TestFactory(t *testing.T) {
	facade := facade{}
	if facade.idontknow() != "idontknow" {
		t.Error("idontknow 아님")
	}
}

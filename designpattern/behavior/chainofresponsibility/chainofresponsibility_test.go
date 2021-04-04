package chainofresponsibility

import "testing"

func TestFactory(t *testing.T) {
	bank3 := &bank3{
		remain: 1000,
		nextBank:  nil,
	}
	bank2 := &bank2{
		remain: 5000,
		nextBank:  bank3,
	}

	bank1 := &bank1{
		remain: 1000,
		nextBank:  bank2,
	}
	wantMoney := 3000
	remoteController := &chainOfBank{
		bankinfo: bank1,
	}
	if remoteController.bankinfo.canBill(wantMoney) {
		t.Error("에헤헤..은행1 미쳤냐?")
	}
	remoteController.next()

	if !remoteController.bankinfo.canBill(wantMoney) {
		t.Error("에헤헤..은행2 미쳤냐?")
	}
}


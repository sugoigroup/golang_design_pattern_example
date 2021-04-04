package visitor

import (
"testing"
)


func TestFactory(t *testing.T) {

	visitor1 := visitor{visitorName: "책임운영자"}

	room1 := Room1{
		visior: &visitor1,
	}

	if !room1.visit() {
		t.Error("책임운영자 1번방 확인실패")
	}

	visitor2 := visitor{visitorName: "감시자"}
	room1 = Room1{
		visior: &visitor2,
	}

	if !room1.visit() {
		t.Error("감시자 1번방 확인실패")
	}

}



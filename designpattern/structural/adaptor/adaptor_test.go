package adaptor

import (
	"testing"
)

func TestFactory(t *testing.T) {

	computer := &computer{}
	msg := computer.connectCable(&usb1{})

	if msg != "usb1 연결" {
		t.Error("usb1 연결오류")
	}

	msg = computer.connectCable(&thunderbird{})
	if msg != "썬더보드 연결" {
		t.Error("썬더보드 연결오류")
	}

}

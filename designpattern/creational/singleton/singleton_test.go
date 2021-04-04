package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		//Test of acceptance criteria 1 failed
		t.Error("싱글톤 생성에러")
	}
}

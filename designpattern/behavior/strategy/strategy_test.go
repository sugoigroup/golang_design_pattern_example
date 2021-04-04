package strategy

import "testing"


func TestFactory(t *testing.T) {
	busChager := busChager{
		basePrice: 1500,
	}
	child := &ruleChild{}

	if busChager.getPrice(child) != 500 {
		t.Error("아동 할인 오류")
	}
	adult := &ruleAdult{}

	if busChager.getPrice(adult) != 1500 {
		t.Error("어른은 얄짤 없음")
	}
}


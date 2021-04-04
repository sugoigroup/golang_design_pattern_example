package builder

import (
	"testing"
)

func TestFactory(t *testing.T) {
	coffeeMachine := coffeMachine("espresso")
	if coffeeMachine.coffe != 2 {
		t.Error("에소프레소는 커피2")
	}

	coffeeMachine = coffeMachine("mix")
	if coffeeMachine.cream != 1  || coffeeMachine.sugar != 1 || coffeeMachine.coffe != 1{
		t.Error("믹스 커피의 조합은 1:1:1")
	}

}

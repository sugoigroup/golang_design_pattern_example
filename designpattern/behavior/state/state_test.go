package state

import (
	"testing"
)


func TestFactory(t *testing.T) {
	mylife := &lifeIs{ state: &born{}}

	if mylife.state.getStateTest() != "태어남" {
		t.Error("아직 안태어남")
	}

	mylife.nextStep()
	mylife.nextStep()

	if mylife.state.getStateTest() != "어른" {
		t.Error("아직 어른이?")
	}
}



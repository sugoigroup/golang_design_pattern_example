package decorator

import (
	"testing"
)

func TestFactory(t *testing.T) {
	baseBugger := &hambugger{}

	meetOnlyBugger := &meetTopping{
		iHambugger: baseBugger,
	}

	if meetOnlyBugger.getPrice() != 1300 {
		t.Error("미트버거 계산오류")
	}

	eggOnlyBugger := &eggTopping{
		iHambugger: baseBugger,
	}

	if eggOnlyBugger.getPrice() != 1200 {
		t.Error("에그버거 계산오류")
	}


	meetEggBugger := &eggTopping{
		iHambugger: meetOnlyBugger,
	}

	if meetEggBugger.getPrice() != 1500 {
		t.Error("미트에그버거 계산오류")
	}


}

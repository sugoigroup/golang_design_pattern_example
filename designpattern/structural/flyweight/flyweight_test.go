package flyweight

import "testing"

func TestFactory(t *testing.T) {
	enermyFactory := &enermyFactory{
		enermies: map[string]iEnermy{},
	}

	enermy1, _ := enermyFactory.getEnermy("goblin")

	if enermy1.getName() != "goblin" {
		t.Error("goblin이 아님")
	}

	enermy2 , _  := enermyFactory.getEnermy("devil")
	if enermy2.getName() != "devil" {
		t.Error("devil 아님")
	}

	enermy3, _ := enermyFactory.getEnermy("goblin")
	if enermy3.getCached() != true {
		t.Error("goblin이 캐쉬되지 않음")
	}

}


package abstractfactory

import (
	"testing"
)

func TestFactory(t *testing.T) {
	produce, _ := produceProduct("car")
	myCar := produce.makeProduct()
	if myCar.getPower() != "200MP" {
		t.Error("자동차  마력은 200MP")
	}

	produce2, _ := produceProduct("airplane")
	myAirplain := produce2.makeProduct()
	if myAirplain.getPower() != "3200MP" {
		t.Error("비행기 마력은 3200MP")
	}

}

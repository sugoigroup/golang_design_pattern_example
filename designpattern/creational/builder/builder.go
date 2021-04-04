package builder

type iCoffeResources interface {
	setCoffee(coffee int) iCoffeResources
	setSugar(suger int) iCoffeResources
	setCream(cream int) iCoffeResources
	getResources() *coffeResource
}

type coffeResource struct {
	coffe int
	sugar int
	cream int
}

func (coffeMix *coffeResource) getResources()  *coffeResource {
	return coffeMix;
}

func (coffeMixer *coffeResource) setCoffee(coffee int) iCoffeResources {
	coffeMixer.coffe = coffee
	return coffeMixer
}
func (coffeMixer *coffeResource) setSugar(sugar int) iCoffeResources {
	coffeMixer.sugar = sugar
	return coffeMixer
}
func (coffeMixer *coffeResource) setCream(cream int) iCoffeResources {
	coffeMixer.cream = cream
	return coffeMixer
}

// 기본 재료 세팅
func newCoffeBulder() iCoffeResources {
	return &coffeResource{
		coffe: 0,
		sugar: 0,
		cream: 0,
	}
}

// 재료로 커피 생산
func  produceCoffee(r iCoffeResources) coffeResource {
	return coffeResource{
		coffe: r.getResources().coffe,
		sugar: r.getResources().sugar,
		cream: r.getResources().cream,
	}
}


func coffeMachine(coffeType string) coffeResource {

	baseCoffee := newCoffeBulder()
	if coffeType == "espresso" {
		baseCoffee.setCoffee(2)
	}
	if coffeType == "mix" {
		baseCoffee.setCoffee(1).setCream(1).setSugar(1)
	}
	return produceCoffee(baseCoffee)
}

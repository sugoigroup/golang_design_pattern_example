package flyweight

import "fmt"

type enermy struct {
	name string
	cached bool
}
type iEnermy interface {
	getName() string
	attack()
	setCached(cache bool)
	getCached() bool
}

func (e *enermy) getName() string{
	return e.name
}

func (e *enermy) attack() {
	fmt.Print("attack")
}

func (e *enermy) setCached(cache bool) {
	e.cached = cache
}

func (e *enermy) getCached()  bool{
	return e.cached
}

type enermyFactory struct {
	enermies map[string]iEnermy
}

func (ef *enermyFactory) getEnermy(ename string) (iEnermy, error)  {
	if ef.enermies[ename] != nil {
		ef.enermies[ename].setCached(true)
		return ef.enermies[ename], nil
	}
	ef.enermies[ename] = &enermy{
		name:   ename,
		cached: false,
	}
	return ef.enermies[ename], nil
}



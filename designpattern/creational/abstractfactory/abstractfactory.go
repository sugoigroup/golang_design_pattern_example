package abstractfactory

import "fmt"

//추상팩토리 인터페이스에서는 각 공장에서 OO를 만들자 라는 공통점을 가지도록한다.
type iAbstractFactory interface {
	makeProduct() iVechicle
}
//탈것 형식에 맞게 생산한다.
func produceProduct(vechicleType string) (iAbstractFactory, error) {
	if vechicleType == "car" {
		return &carFactory{}, nil
	}

	if vechicleType == "airplane" {
		return &airplaneFactory{}, nil
	}

	return nil, fmt.Errorf("그런거 못 만들어요")
}

// 자동차 공장
type carFactory struct {
}
// 탈것을 만들되 자동차공장 구조체를 통해서 각부품을 생산한다.
func (cf *carFactory) makeProduct() iVechicle {
	return &vechicle{
		factoring: factoring{
			engineFactory: engineFactory{ power: "200MP" },
			wheelFactory:  wheelFactory{count: "4"},
			wingFactory:   wingFactory{wing: "0"},
		},
	}
}


type airplaneFactory struct {
}

func (cf *airplaneFactory) makeProduct() iVechicle {
	return &vechicle{
		factoring: factoring{
			engineFactory: engineFactory{ power: "3200MP" },
			wheelFactory:  wheelFactory{count: "3"},
			wingFactory:   wingFactory{wing: "2"},
		},
	}
}

// 탈것이 공통으로 가질 속성을 정의한다.
type iVechicle interface {
	setPower(power string)
	getPower() string
	setWheel(wheel string)
	getWheel() string
	setWing(wing string)
	getWing() string
}

// 생산과정에서 엔진, 바퀴, 날개 각 공장을 통해 생산한다.
type factoring struct {
	engineFactory
	wheelFactory
	wingFactory
}

// 탈것은 생산을 통해 만들어 진다.
type vechicle struct {
	factoring
}

// 탈것의 공통속성에 대한 구현을 한다.
func (v vechicle) getPower() string {
	return v.power
}
func (v vechicle) getWheel() string {
	return v.count
}
func (v vechicle) getWing() string {
	return v.wing
}
func (v vechicle) setPower(power string) {
	v.power = power
}
func (v vechicle) setWheel(wheel string) {
	v.count = wheel
}
func (v vechicle) setWing(wing string) {
	v.wing = wing
}

//각 공장을 정의한다.
type engineFactory struct {
	power string
}
type wheelFactory struct {
	count string
}
type wingFactory struct {
	wing string
}
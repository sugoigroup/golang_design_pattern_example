package factory



type factoryLine struct {
	name string
}

func (factoryLine *factoryLine) setName(name string) {
	factoryLine.name = name
}

func (factoryLine *factoryLine) getName() string {
	return factoryLine.name
}

// 요소의 이름과 타입을 같이 지정하면 bmw.getName() 처럼 함축해서 기술할수 있다
type BMW struct {
	factoryLine
}

func makeBMW() *BMW {
	return &BMW{
		factoryLine: factoryLine{
			name: "BMW",
		},
	}
}

// struct 의 요소에 이름을 지정하면 audi.whichLine.getName() 처럼 호출해야한다.
type AUDI struct {
	whichLine factoryLine
}


func makeAUDI() *AUDI {
	return &AUDI{
		whichLine: factoryLine{
			name: "AUDI",
		},
	}
}


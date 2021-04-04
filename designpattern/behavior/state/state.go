package state

type states interface {
	getStateTest() string
	nextSteate() states
}

type born struct {}
func (b *born) getStateTest() string {
	return "태어남"
}
func (b *born) nextSteate() states {
	return &infant{}
}

type infant struct {}
func (b *infant) getStateTest() string {
	return "아기때"
}
func (b *infant) nextSteate() states {
	return &adult{}
}

type adult struct {}
func (b *adult) getStateTest() string {
	return "어른"
}
func (b *adult) nextSteate() states {
	return &die{}
}

type die struct {}
func (b *die) getStateTest() string {
	return "죽어서 환생준비"
}
func (b *die) nextSteate() states {
	return &born{}
}

type lifeIs struct {
	state states
}

func (life *lifeIs) nextStep()  {
	life.state = life.state.nextSteate()
}
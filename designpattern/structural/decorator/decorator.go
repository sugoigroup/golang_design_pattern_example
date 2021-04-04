package decorator

type iHambugger interface {
	getPrice() int
}
type hambugger struct {
	price int
}
func (h *hambugger) getPrice() int {
	return 1000
}

type meetTopping struct {
	iHambugger
}

func (m *meetTopping) getPrice() int {
	return m.iHambugger.getPrice() + 300
}

type eggTopping struct {
	iHambugger
}

func (e *eggTopping) getPrice() int {
	return  e.iHambugger.getPrice() + 200
}
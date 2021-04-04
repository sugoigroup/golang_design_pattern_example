package facade

type facade struct {
	toResult iVeryComplexLogic
}

type iVeryComplexLogic interface {
	idontknow() string
}

func (f *facade) idontknow() string {
	return "idontknow"
}

package bridge

type mac struct {
	printer
}

func (m *mac) print() string  {
	return m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

type printer interface {
	printFile() string
}

type epson struct {}

func (e *epson) printFile()  string {
	return "epson"
}

type samsung struct {}

func (e *samsung) printFile()  string {
	return "samsung"
}


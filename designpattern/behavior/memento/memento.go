package memento

//https://golangbyexample.com/memento-design-pattern-go/ 참고함
//설명은 여기서 : https://yunhos.blogspot.com/2021/03/kotlindesignpatternebookkth.html#strategy-pattern
// 기념비에 대한 설명을 보존
type originator struct {
	state string
}

func (e *originator) createMemento() *memento {
	return &memento{state: e.state}
}

func (e *originator) restoreMemento(m *memento) {
	e.state = m.getSavedState()
}

func (e *originator) setState(state string) {
	e.state = state
}

func (e *originator) getState() string {
	return e.state
}

//순수객체
type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}

// 기념비목록을 추가,관리
type caretaker struct {
	mementoArray []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *caretaker) getMemento(index int) *memento {
	return c.mementoArray[index]
}
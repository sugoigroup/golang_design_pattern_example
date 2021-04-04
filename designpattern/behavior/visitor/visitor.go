package visitor

import "fmt"

type camera interface {
	display()
	cleanCheck()
	visit() bool
}

type Room1 struct {
	visior iVisitor
}

func (r *Room1) display() {
	fmt.Print("1번방 카메라보기")
}

func (r *Room1) cleanCheck() {
	fmt.Print("1번방 청소확인")
}

func (r *Room1) visit() bool{
	r.visior.visitForRoom1(r)
	return true
}
type Room2 struct {
	visior iVisitor
}

func (r *Room2) display() {
	fmt.Println("1번방 카메라보기")
}

func (r *Room2) cleanCheck() {
	fmt.Println("1번방 청소확인")
}

func (r *Room2) visit() bool{
	r.visior.visitForRoom2(r)
	return true;
}

type iVisitor interface {
	visitForRoom1(room1 *Room1)
	visitForRoom2(room2 *Room2)
}

type visitor struct {
	visitorName string
}

func (v *visitor) visitForRoom1(room1 *Room1) {
	fmt.Println(v.visitorName)
	room1.display()
	room1.cleanCheck()
}

func (v *visitor) visitForRoom2(room2 *Room2) {
	fmt.Println(v.visitorName)
	room2.display()
	room2.cleanCheck()
}


package template

import "fmt"

type dailyRoutine interface {
	commute()
	toWork
	eatLunch()
	goHome()
	showTeamsDailyroutine()
}
type toWork interface {
	toWork()
}

type commonRoutine struct {
	toWork
}

func (c *commonRoutine) commute() {
	fmt.Println("commute")
}

func (c *commonRoutine) eatLunch() {
	fmt.Println("eatLunch")
}

func (c *commonRoutine) goHome() {
	fmt.Println("gohome")
}

func (c *commonRoutine) showTeamsDailyroutine() {
	c.commute()
	c.toWork.toWork()
	c.eatLunch()
	c.goHome()
}

type hrTeam struct {

}

func (h *hrTeam) toWork() {
	fmt.Println("work is seek!seek!seek!")
}

type devTeam struct {

}

func (h *devTeam) toWork() {
	fmt.Println("work is coding!coding!coding!")
}




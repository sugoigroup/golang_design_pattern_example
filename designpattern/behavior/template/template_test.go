package template

import "testing"


func TestFactory(t *testing.T) {
	hrTeamWork := hrTeam{}

	baseRoutine := commonRoutine{
		toWork: &hrTeamWork,
	}
	baseRoutine.showTeamsDailyroutine()

	baseRoutine = commonRoutine{
		toWork: &devTeam{},
	}
	baseRoutine.showTeamsDailyroutine()

}


package main

import (
	itschool "strategy/ITSchool"
	"strategy/core"
	"strategy/strategies"
	"time"
)

func main() {
	codingSet := []core.Coding{
		strategies.BackEnd{},
		strategies.FrontEnd{},
	}
	CheckTrainee(codingSet)
}
func CheckTrainee(codingSet []core.Coding) {
	for {
		tr := itschool.GenerateTrainee()
		for _, coding := range codingSet {
			if coding.TestTrainee(&tr) {
				tr.Learn(coding)
			}
		}
		time.Sleep(time.Second * 1)
	}
}

type Coding interface {
	Name() string
	TestTrainee(Trainee) Coding
}

type Trainee interface {
	Age() int
	Predisposition() string
}

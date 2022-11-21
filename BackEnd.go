package strategies

import "strategy/core"

type BackEnd struct {
}

func (b BackEnd) Name() string {
	return `coding lover`
}

func (b BackEnd) TestTrainee(trainee core.Trainee) bool {
	return trainee.Predisposition() == `brain`
}

package strategies

import "strategy/core"

type FrontEnd struct {
}

func (f FrontEnd) Name() string {
	return `user experience artist`
}

func (f FrontEnd) TestTrainee(trainee core.Trainee) bool {
	return trainee.Predisposition() == `creativity`
}

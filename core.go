package core

type Coding interface {
	Name() string
	TestTrainee(Trainee) bool
}

type Trainee interface {
	Age() int
	Predisposition() string
	Learn(Coding)
}

package itschool

import (
	"fmt"
	"math/rand"
	"strategy/core"
	"time"
)

type Trainee struct {
	age            int
	predisposition string
	coding         core.Coding
}

func (t Trainee) Age() int {
	return t.age
}

func (t *Trainee) Predisposition() string {
	return t.predisposition
}

func (t *Trainee) Learn(c core.Coding) {
	t.coding = c
	fmt.Printf("Programmer with age of %d has become a %s\n", t.Age(), c.Name())
}
func GenerateTrainee() Trainee {
	rand.Seed(time.Now().Unix())
	a := rand.Intn(60)
	p := PredispositionGenerator()

	return Trainee{age: a, predisposition: p}
}

func PredispositionGenerator() string {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(21)

	switch r {
	case 0:
		return `brain`
	case 1:
		return `creativity`
	default:
		fallthrough
	case 3:
		return `nothing`
	}
}

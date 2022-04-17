package domains

import (
	"log"
	"math/rand"
	"time"
)

type Architecture struct {
	name string
}

func NewArchitecture(name string) *Architecture {
	return &Architecture{
		name: name,
	}
}

func (a Architecture) Calculate(project *Project) {
	log.Println("Project Calculating")
	totalCost := 250 * project.difficult
	for {
		if totalCost <= 0.0 {
			break
		}
		cost, level := a.getTaskDetails(&totalCost, project.difficult)
		NewTask(cost, level, "Story", project)
		project.wg.Add(1)
	}
}

func (a Architecture) getTaskDetails(totalCost *float64, difficult float64) (float64, int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	cost := float64(rnd.Int() % 20)
	if cost > *totalCost {
		cost = *totalCost
	}
	*totalCost -= cost
	level := 1
	for i := int(difficult); i > 1; i-- {
		if int(*totalCost)%(i*2) == 0 {
			level = i
			break
		}
	}
	return cost, level
}

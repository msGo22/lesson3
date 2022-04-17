package domains

import (
	"github.com/google/uuid"
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
	totalCost := 100 * project.difficult
	for {
		if totalCost <= 0.0 {
			break
		}
		taskID := uuid.New().String()
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		cost := float64(rnd.Int() % 20)
		level := 1
		if (rnd.Int() % 100) > 70 {
			level = 2
		}
		if cost > totalCost {
			cost = totalCost
		}
		totalCost -= cost
		project.tasks = append(project.tasks, &Task{
			Id:            taskID,
			assigned:      nil,
			cost:          cost,
			requiredLevel: level,
			Status:        false,
		})
		project.wg.Add(1)
	}
}

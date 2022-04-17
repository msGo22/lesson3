package domains

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type Task struct {
	Id            string
	taskType      string
	assigned      *Developer
	cost          float64
	requiredLevel int
	Status        bool
	failed        time.Time
	project       *Project
}

func (t *Task) Done(developer *Developer) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if rnd.Int()%100 < 35 {
		if subTaskCost := t.cost * 0.25; subTaskCost > 0.01 {
			t.project.wg.Add(1)
			level := developer.level - 1
			if level == 0 {
				level = 1
			}
			newTask := NewTask(subTaskCost, level, "Bug", t.project)
			developer.project.workList <- newTask
		}
	}
	t.assigned = developer
	time.Sleep(time.Duration(t.cost*100) * time.Millisecond)
	t.Status = true
	t.project.wg.Done()
}

func (t *Task) check(developer *Developer) bool {

	return true
}

func NewTask(cost float64, level int, taskType string, project *Project) *Task {
	task := &Task{
		Id:            uuid.New().String(),
		taskType:      taskType,
		assigned:      nil,
		cost:          cost,
		requiredLevel: level,
		Status:        false,
		failed:        time.Time{},
		project:       project,
	}
	task.project.tasks = append(task.project.tasks, task)
	return task
}

package domains

import (
	"fmt"
	"sync"
)

type Task struct {
	Id            string
	assigned      *Developer
	cost          float64
	requiredLevel int
	Status        bool
}

type Project struct {
	title     string
	difficult float64
	tasks     []*Task
	wg        *sync.WaitGroup
}

func NewProject(title string, difficult float64) *Project {
	return &Project{
		title:     title,
		difficult: difficult,
		tasks:     []*Task{},
		wg:        new(sync.WaitGroup),
	}
}

func (p Project) Statuses() {
	isCompleted := true
	fmt.Printf("Proje Raporu \n")
	fmt.Printf("=============================\n")
	for _, task := range p.tasks {
		if !task.Status {
			isCompleted = false
		}
		fmt.Printf("%s (lv %d) görevini : %s tamamladı\n", task.Id, task.requiredLevel, task.assigned.name)
	}
	if isCompleted {
		fmt.Printf("Proje Bitmiştir\n")
	}
}

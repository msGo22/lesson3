package domains

import (
	"fmt"
	"sync"
)

type Project struct {
	title            string
	difficult        float64
	tasks            []*Task
	wg               *sync.WaitGroup
	workList         chan *Task
	urgentList       chan *Task
	teamLeadMessages chan *Task
}

func NewProject(title string, difficult float64) *Project {
	return &Project{
		title:            title,
		difficult:        difficult,
		tasks:            []*Task{},
		wg:               new(sync.WaitGroup),
		workList:         make(chan *Task, 100),
		urgentList:       make(chan *Task, 100),
		teamLeadMessages: make(chan *Task, 100),
	}
}

func (p Project) Statuses() {
	isCompleted := true
	fmt.Printf("Proje Raporu \n")
	fmt.Printf("=============================\n")
	for _, task := range p.tasks {
		statusMessage := "tamamlandı"
		if !task.Status {
			isCompleted = false
			statusMessage = "bitirilemedi."
		}
		fmt.Printf("%s - %s  (lv %d) görevi %s\n", task.taskType, task.Id, task.requiredLevel, statusMessage)
	}
	if isCompleted {
		fmt.Printf("Proje Bitmiştir\n")
	}
}

package domains

import (
	"context"
	"fmt"
	"time"
)

type Developer struct {
	name       string
	level      int
	project    *Project
	urgentList chan *Task
}

func NewDeveloper(name string, level int, project *Project) *Developer {
	var urgentList chan *Task
	if level >= int(project.difficult) {
		urgentList = project.urgentList
	}
	return &Developer{
		name:       name,
		level:      level,
		project:    project,
		urgentList: urgentList,
	}
}

func (d *Developer) Run(ctx context.Context) {
	for {
		select {
		case task := <-d.urgentList:
			fmt.Printf("%s \t bir %s \t acil olarak aldı\n", d.name, task.taskType)
			task.Done(d)
		case task := <-d.project.workList:
			fmt.Printf("%s \t bir %s \t aldı\n", d.name, task.taskType)
			if !task.failed.IsZero() && task.failed.Before(time.Now()) {
				continue
			}
			if task.failed.IsZero() && d.level < task.requiredLevel {
				task.failed = time.Now().Add(time.Duration(task.cost*100) * 5 * time.Millisecond)
				d.project.urgentList <- task
				fmt.Printf("%s \t bir %s  \t iade edildi\n", d.name, task.taskType)
				continue
			}
			task.Done(d)
		case <-ctx.Done():
			fmt.Printf("%s \t işten çıktı \n", d.name)
			return
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func (d *Developer) SetLevel(level int) {
	d.level = level
}

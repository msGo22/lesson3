package domains

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type TeamLead struct {
	developers   []*Developer
	architecture *Architecture
	project      *Project
}

func NewTeamLead(project *Project) *TeamLead {
	return &TeamLead{
		developers:   nil,
		architecture: nil,
		project:      project,
	}
}

func (l *TeamLead) Start() error {
	if l.project == nil {
		return errors.New("project is not exist")
	}
	l.architecture.Calculate(l.project)
	return l.runAllDevelopers()
}

func (l *TeamLead) sendAllRequestToBoard(project *Project) {
	for _, task := range project.tasks {
		l.project.workList <- task
		time.Sleep(250 * time.Millisecond)
	}
}

func (l *TeamLead) runAllDevelopers() error {
	ctx, cancel := context.WithCancel(context.Background())
	go l.sendAllRequestToBoard(l.project)
	go l.newAssignments(ctx)
	for _, v := range l.developers {
		go v.Run(ctx)
	}
	l.project.wg.Wait()
	// tüm developerlar paydos verir
	cancel()
	return nil
}

func (l *TeamLead) newAssignments(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case task := <-l.project.teamLeadMessages:
			fmt.Println("Takım liderine mesaj geldi")
			if task.failed.IsZero() {
				l.project.workList <- task
				continue
			}
			l.project.urgentList <- task
		}
	}
}

func (l *TeamLead) AddDeveloper(name string, level int) {
	developer := NewDeveloper(name, level, l.project)
	if l.developers == nil {
		l.developers = []*Developer{developer}
		return
	}
	l.developers = append(l.developers, developer)
}

func (l *TeamLead) AddArchitecture(name string) {
	l.architecture = NewArchitecture(name)
}

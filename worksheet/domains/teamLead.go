package domains

import (
	"context"
	"log"
)

type TeamLead struct {
	developers   []*Developer
	architecture *Architecture
	levels       map[int]Level
}

func NewTeamLead(architecture *Architecture, developers []*Developer, levels ...Level) *TeamLead {
	tmpLevels := map[int]Level{}
	for _, level := range levels {
		if _, ok := tmpLevels[level.lv]; !ok {
			tmpLevels[level.lv] = level
		}
	}
	return &TeamLead{
		developers:   developers,
		architecture: architecture,
		levels:       tmpLevels,
	}
}

func (l *TeamLead) Start(project *Project) error {
	l.architecture.Calculate(project)
	return l.runAllDevelopers(project)
}

func (l *TeamLead) sendAllRequestToBoard(project *Project) {
	for _, task := range project.tasks {
		if v, ok := l.levels[task.requiredLevel]; ok {
			v.lvChan <- task
			continue
		}
		log.Println("Unsupported work")
	}
}

func (l *TeamLead) runAllDevelopers(project *Project) error {
	ctx, cancel := context.WithCancel(context.Background())
	for _, v := range l.developers {
		go v.Run(ctx, project.wg)
	}
	go l.sendAllRequestToBoard(project)
	project.wg.Wait()
	// tüm developerlar paydos verir
	cancel()
	return nil
}

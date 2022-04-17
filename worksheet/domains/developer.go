package domains

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Developer struct {
	name  string
	level Level
}

func NewDeveloper(name string, level Level) *Developer {
	return &Developer{
		name:  name,
		level: level,
	}
}

func (d *Developer) Run(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case task := <-d.level.lvChan:
			fmt.Printf("%s bir görevi aldı\n", d.name)
			task.assigned = d
			time.Sleep(time.Duration(task.cost*100) * time.Millisecond)
			task.Status = true
			fmt.Printf("%s bir görevi bitirdi\n", d.name)
			wg.Done()
		case <-ctx.Done():
			fmt.Printf("%s gorevlerini tamamladı", d.name)
			return
		default:
			fmt.Printf("%s görev bekliyor (Aktif Görev: %d) \n", d.name, len(d.level.lvChan))
			time.Sleep(time.Second * 2)
		}
	}
}

func (d *Developer) SetLevel(level Level) {
	d.level = level
}

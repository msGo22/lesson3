package domains

type Level struct {
	lv     int
	lvChan chan *Task
}

func NewLevel(lv int) Level {
	return Level{
		lv:     lv,
		lvChan: make(chan *Task, 3),
	}
}

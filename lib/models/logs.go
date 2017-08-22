package models

import (
	"time"

	"github.com/vence722/gcoll/list"
)

// Log entity
type Log struct {
	Content string
	Time    time.Time
}

// The store of logs
type Logs struct {
	key  string
	logs *list.ArrayList
}

func NewLogs(key string) *Logs {
	return &Logs{key: key, logs: list.NewArrayList()}
}

func (this *Logs) Put(logLine string) {
	this.logs.Add(&Log{Content: logLine, Time: time.Now()})
}

func (this *Logs) GetOne(from int) *Log {
	var logEntity *Log = nil
	if from >= 0 && from < this.logs.Size() {
		logLine := this.logs.Get(from)
		if logLine != nil {
			logEntity = logLine.(*Log)
		}
	}
	return logEntity
}

func (this *Logs) Get(from int, num int) []*Log {
	logs := []*Log{}
	// Get All the existing logs
	if num < 0 {
		for {
			log := this.GetOne(from)
			from++
			if log != nil {
				logs = append(logs, log)
			} else {
				break
			}
		}
	} else {
		for i := 0; i < num; i++ {
			log := this.GetOne(i)
			if log != nil {
				logs = append(logs, log)
			}
		}
	}
	return logs
}

func (this *Logs) Clear() {
	this.logs.Clear()
}

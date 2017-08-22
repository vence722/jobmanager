package models

import (
	"time"

	"github.com/vence722/jobmanager/lib/errors"
)

// Job status codes
const (
	// normal cases
	JOB_STATUS_CREATED  = 0
	JOB_STATUS_RUNNING  = 1
	JOB_STATUS_FINISHED = 2
	// error cases
	JOB_STATUS_ERROR_START        = -1
	JOB_STATUS_ERROR_RUNNING      = -2
	JOB_STATUS_ERROR_POST_PROCESS = -3
)

// Job struct
type Job struct {
	Id         string
	Params     string
	Logs       *Logs
	StartTime  time.Time
	FinishTime time.Time
	Status     int
	LastError  *errors.Error
}

// Job Constructor
func NewJob(jobId string, params string) *Job {
	return &Job{
		Id:        jobId,
		Params:    params,
		Logs:      NewLogs(jobId),
		StartTime: time.Now(),
		Status:    JOB_STATUS_RUNNING,
	}
}

// Job list

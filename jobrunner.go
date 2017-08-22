package jobmanager

import (
	"github.com/vence722/jobmanager/lib/errors"
	"github.com/vence722/jobmanager/lib/models"
)

type JobRunner interface {
	Run(manager *JobManager, job *models.Job) (chan *models.Log, *errors.Error)
	ProcessRunningJob(manager *JobManager, job *models.Job, ch chan *models.Log) *errors.Error
	PostProcessJob(manager *JobManager, job *models.Job) *errors.Error
}

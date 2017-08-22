package jobmanager

import (
	"fmt"
	"testing"
	"time"

	"github.com/vence722/convert"
	"github.com/vence722/jobmanager/lib/errors"
	"github.com/vence722/jobmanager/lib/models"
)

type TestJobRunner struct{}

func (this *TestJobRunner) Run(manager *JobManager, job *models.Job) (chan *models.Log, *errors.Error) {
	ch := make(chan *models.Log)
	go func() {
		for i := 0; i < 10000000; i++ {
			ch <- &models.Log{Content: convert.Int2Str(i), Time: time.Now()}
		}
	}()
	return ch, nil
}

func (this *TestJobRunner) ProcessRunningJob(manager *JobManager, job *models.Job, ch chan *models.Log) *errors.Error {
	for {
		action := 0
		select {
		case v := <-ch:
			fmt.Println(v)
			if v.Content == "1000000" {
				action = 1
			}
		}
		if action == 1 {
			break
		}
	}
	return nil
}

func (this *TestJobRunner) PostProcessJob(manager *JobManager, job *models.Job) *errors.Error {
	fmt.Println("postprocessing job")
	return nil
}

func testJobManager(t *testing.T) {
	jobmanager := NewJobManager(&TestJobRunner{})
	job, _ := jobmanager.CreateJob("test1", "no params")
	jobmanager.StartJob(job)
	fmt.Println(jobmanager.ListJobs())
	fmt.Println(jobmanager.GetJob("test1"))
}

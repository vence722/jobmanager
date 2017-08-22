package jobmanager

import (
	"time"

	"github.com/vence722/jobmanager/lib/errors"
	"github.com/vence722/jobmanager/lib/models"
)

// Job manager struct
type JobManager struct {
	jobMap    map[string]*models.Job
	jobRunner JobRunner
}

// Job manager constructor, accepts a jobRunner instance
func NewJobManager(jobRunner JobRunner) *JobManager {
	return &JobManager{
		jobMap:    make(map[string]*models.Job),
		jobRunner: jobRunner,
	}
}

// Create a new job
func (this *JobManager) CreateJob(jobId string, params string) (*models.Job, *errors.Error) {
	var job *models.Job
	if job = this.jobMap[jobId]; job != nil {
		return nil, errors.NewError(errors.ERR_JOB_EXISTS, "Job already exists", nil)
	}
	this.jobMap[jobId] = models.NewJob(jobId, params)
	return this.jobMap[jobId], nil
}

// Get a single job object by job id
func (this *JobManager) GetJob(jobId string) (*models.Job, *errors.Error) {
	var job *models.Job
	if job = this.jobMap[jobId]; job == nil {
		return nil, errors.NewError(errors.ERR_JOB_NOT_EXISTS, "Job not exists", nil)
	}
	return job, nil
}

// List all job ids and job objects
func (this *JobManager) ListJobs() ([]string, []*models.Job) {
	var jobIds []string
	var jobList []*models.Job
	for jobId, job := range this.jobMap {
		jobIds = append(jobIds, jobId)
		jobList = append(jobList, job)
	}
	return jobIds, jobList
}

// Start a job, will update job status base on the return value of each stage of jobRunner
func (this *JobManager) StartJob(job *models.Job) *errors.Error {
	ch, err := this.jobRunner.Run(this, job)
	if err != nil {
		// Fail to start job, mark job status
		job.LastError = err
		job.Status = models.JOB_STATUS_ERROR_START
		job.FinishTime = time.Now()
		return err
	}
	// Job started, mark status
	job.Status = models.JOB_STATUS_RUNNING
	// Set start time
	job.StartTime = time.Now()

	// Asynchronous process the running job
	go func() {
		err = this.jobRunner.ProcessRunningJob(this, job, ch)
		if err != nil {
			// Fail during running, mark job status
			job.LastError = err
			job.Status = models.JOB_STATUS_ERROR_RUNNING
			job.FinishTime = time.Now()
			return
		}
		err = this.jobRunner.PostProcessJob(this, job)
		if err != nil {
			// Fail during running, mark job status
			job.LastError = err
			job.Status = models.JOB_STATUS_ERROR_POST_PROCESS
			job.FinishTime = time.Now()
			return
		}
		// All finished, mark job status
		job.Status = models.JOB_STATUS_FINISHED
		job.FinishTime = time.Now()
	}()

	return nil
}

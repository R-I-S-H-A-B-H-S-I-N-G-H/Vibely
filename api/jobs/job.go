package jobs

import "github.com/robfig/cron/v3"

type Job struct{}

var songJob *SongJob

func (j *Job) InitiateJob() {
	songJob.InitiateSongJobs()
}

func (j *Job) CreateJob(spec string, cmd func()) (cron.EntryID, error) {
	c := cron.New(cron.WithSeconds())
	defer c.Start()

	res, err := c.AddFunc(spec, cmd)
	return res, err
}

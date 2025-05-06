package jobs

import (
	"fmt"

	"github.com/R-I-S-H-A-B-H-S-I-N-G-H/Vibely/api/service"
)

type SongJob struct{}

var job *Job
var songService *service.SongService

func (s *SongJob) InitiateSongJobs() {
	_, err := job.CreateJob("*/60 * * * * *", s.generateHls)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *SongJob) generateHls() {
	fmt.Println("Process HLS job running")
	songService.GenerateHLSJob()
}

package lld

import (
	"fmt"
	"sort"
)

type Job struct {
	JobTitle string
	Priority Priority
}
type Priority int

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)

var jobs []*Job

func AddJob(job *Job) {

	if job.JobTitle != "" {
		jobs = append(jobs, job)

	}

}

func GetJob() *Job {

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Priority > jobs[j].Priority
	})

	top := jobs[0]
	jobs = jobs[1:]
	return top

}

func jobDemo() {
	job1 := &Job{

		JobTitle: "Low Priority",
		Priority: PriorityLow,
	}

	job2 := &Job{
		JobTitle: "Medium Priority",
		Priority: PriorityMedium,
	}

	AddJob(job1)
	AddJob(job2)
	fmt.Println(GetJob())

}

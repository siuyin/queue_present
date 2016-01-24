package q

// re-implement from algorithm

import (
	"time"
)

// 01 OMIT
const N = 10

var maxTime = time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
var clk Nower
var queues [N]FIFO
var lastVirtFinish [N]time.Time

// 02 OMIT
func receive(j *Job) {
	queueNum := chooseQueue(*j)
	queues[queueNum].Enqueue(*j) // HL01
	updateTime(j, queueNum)
}

// 03 OMIT
func chooseQueue(j Job) int {
	return 1 // assign to a queue. For now we always assign to queue 1
}

// 04 OMIT
func updateTime(j *Job, qn int) {
	// virtStart is the virtual start of service
	virtStart := max(now(), lastVirtFinish[qn])
	j.VirtFinish = virtStart.Add(j.Cost()) // HL01
	lastVirtFinish[qn] = j.VirtFinish
}

// 05 OMIT
func send() Job {
	queueNum := selectQueue() // HL01
	if queueNum == -1 {
		return Job{Id: "idle_job"}
	}
	j, err := queues[queueNum].Dequeue()
	if err != nil {
		return Job{Id: "idle_job"}
	}
	return j.(Job)
}

// 06 OMIT
func selectQueue() int {
	it := 0
	queueNum := -1
	minVirtFinish := maxTime
	for it < N {
		queue := queues[it]
		if queue.Len() > 0 {
			h, _ := queue.Head()
			j := h.(Job)
			if j.VirtFinish.Before(minVirtFinish) {
				minVirtFinish = j.VirtFinish
				queueNum = it
			}
		}
		it = it + 1
	}
	return queueNum
}

// 07 OMIT
type Job struct {
	VirtFinish time.Time
	Id         string
}

func (j Job) Cost() time.Duration {
	return 10 * time.Second
}

type Nower interface {
	Now() time.Time
}

type Clock struct{}

func (c Clock) Now() time.Time {
	return time.Now()
}

func now() time.Time {
	return clk.Now()
}

func max(a, b time.Time) time.Time {
	if b.After(a) {
		return b
	} else {
		return a
	}
}

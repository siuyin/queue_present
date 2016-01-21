package q

// re-implement from algorithm

import (
	"time"
)

const N = 10

var maxTime = time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)

type Job struct {
	VirtFinish time.Time
}

func (j Job) Cost() time.Duration {
	return 10 * time.Second
}

type FQueue struct {
	FIFO
}

var queues []FQueue
var lastVirtFinish []time.Time

func chooseQueue(j Job) int {
	return 1 // assign to a queue. For now we always assign to queue 1
}

func now() time.Time {
	return time.Now()
}

func max(a, b time.Time) time.Time {
	if b.After(a) {
		return b
	} else {
		return a
	}
}

func updateTime(j *Job, qn int) {
	// virtStart is the virtual start of service
	virtStart := max(now(), lastVirtFinish[qn])
	j.VirtFinish = virtStart.Add(j.Cost())
	lastVirtFinish[qn] = j.VirtFinish
}

func receive(j Job) {
	queueNum := chooseQueue(j)
	queues[queueNum].Enqueue(j)
	updateTime(&j, queueNum)
}

func selectQueue() int {
	it := 1
	minVirtFinish := maxTime
	for it <= N {
		queue := queues[it]
		if queue.Len() > 0 {
			h, _ := queue.Head()
			if h.VirtFinish < minVirtFinish {
				minVirtFinish = queue.Head.VirtFinish
				queueNum = it
			}
		}
		it = it + 1
	}
	return queueNum
}

func send() Job {
	queueNum = selectQueue()
	j = queues[queueNum].Dequeue()
	return j
}

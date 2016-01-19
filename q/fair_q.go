package q

import (
	"time"
)

type FQJobber interface {
	Cost() time.Duration
	EstFinishTime()
	Now() time.Duration
}
type FQJob struct {
	dat        interface{}
	finishTime time.Time
}

func (j *FQJob) Cost() time.Duration {
	return 10 * time.Second
}

func (j *FQJob) EstFinishTime() time.Time {
	j.finishTime = time.Now().Add(j.Cost())
	return j.finishTime
}

func (j FQJob) Now() time.Time {
	return time.Now()
}

type FairQ struct {
	finishTime time.Time
	j          []FQJobber
}

func NewFairQ() *FairQ {
	f := new(FairQ)
	f.j = make([]FQJobber, 0)
	return f
}

func (e FairQ) Len() int {
	return len(e.j)
}

func max(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	} else {
		return b
	}
}
func (e *FairQ) Enqueue(jp FQJobber) error {
	e.j = append(e.j, jp)
	e.finishTime = max(jp.Now(), jp.EstFinishTime())
	return nil
}

func (e FairQ) Head() (*FQJobber, error) {
	if e.Len() == 0 {
		return nil, &QError{"Queue is empty"}
	}
	return e.j[0], nil
}

func (e *FairQ) Dequeue() (*FQJobber, error) {
	if e.Len() == 0 {
		return nil, &QError{"Queue is empty"}
	}
	v := e.j[0]
	e.j[0] = nil // avoid memory leak
	e.j = e.j[1:]
	return v, nil
}

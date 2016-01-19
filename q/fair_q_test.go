package q

import (
	"testing"
	"time"
)

type FQJobFake struct {
	dat        interface{}
	finishTime time.Time
}

func (j *FQJobFake) Cost() time.Duration {
	return 10 * time.Second
}

func (j *FQJobFake) EstFinishTime() {
	j.finishTime = j.Now().Add(j.Cost())
}

func (j FQJobFake) Now() time.Time {
	return time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
}

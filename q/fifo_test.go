package q

import (
	"testing"

	"github.com/siuyin/queues_present/q"
)

func TestLen(t *testing.T) {
	qu := q.NewFIFO()
	if qu == nil {
		t.Error("q is null")
	}
	if qu.Len() != 0 {
		t.Error("did not start with an empty queue")
	}
}

func TestEnqueue(t *testing.T) {
	qu := q.NewFIFO()
	qu.Enqueue(1)
	if qu.Len() != 1 {
		t.Error("enqueue failed")
	}
}
func TestDequeueEmpty(t *testing.T) {
	qu := q.NewFIFO()
	_, e := qu.Dequeue()
	if e == nil {
		t.Error("should have errored")
	}
}
func TestDequeue(t *testing.T) {
	qu := q.NewFIFO()
	qu.Enqueue(2)
	o, _ := qu.Dequeue()
	if o.(int) != 2 {
		t.Error("dequeue failed")
	}
}

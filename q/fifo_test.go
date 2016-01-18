package q

import (
	"testing"
)

// 01 OMIT
func TestLen(t *testing.T) {
	qu := NewFIFO()
	if qu == nil { // HL01
		t.Error("q is null")
	}
	if qu.Len() != 0 { // HL01
		t.Error("did not start with an empty queue")
	}
}

// 02 OMIT
func TestEnqueue(t *testing.T) {
	qu := NewFIFO()
	qu.Enqueue(1)
	if qu.Len() != 1 {
		t.Error("enqueue failed")
	}
}
// 03 OMIT
func TestDequeueEmpty(t *testing.T) {
	qu := NewFIFO()
	_, e := qu.Dequeue()
	if e == nil {
		t.Error("should have errored")
	}
}
// 04 OMIT
func TestDequeue(t *testing.T) {
	qu := NewFIFO()
	qu.Enqueue(2) // HL02
	o, _ := qu.Dequeue()
	if o.(int) != 2 { // HL02
		t.Error("dequeue failed")
	}
}
// 05 OMIT
func TestDequeue2(t *testing.T) {
	qu := NewFIFO()
	qu.Enqueue(1) // HL03
	qu.Enqueue("boy") 
	o, _ := qu.Dequeue()
	if o.(int) != 1 { // HL03
		t.Error("dequeue failed")
	}
	
	o, _ = qu.Dequeue()
	if o.(string) != "boy" { 
		t.Error("dequeue failed")
	}
}
// 06 OMIT
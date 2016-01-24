package q

import (
	"testing"
	"time"
)

type ClockFake struct{}

func (c ClockFake) Now() time.Time {
	return time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
}
func TestNow(t *testing.T) {
	clk = ClockFake{}
	exp := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	if !now().Equal(exp) {
		t.Error("incorrect time for Fake Now()")
	}
}
func TestNowReal(t *testing.T) {
	clk = Clock{}
	exp := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	if now().Equal(exp) {
		t.Error("incorrect time for RealNow()")
	}
}
func TestTimeMax(t *testing.T) {
	t1 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC)
	if !max(t1, t2).Equal(t2) {
		t.Error("max should return t2")
	}
	if !max(t2, t1).Equal(t2) {
		t.Error("max should return t2")
	}
}

func TestChooseQueue(t *testing.T) {
	if chooseQueue(Job{}) != 1 {
		t.Error("should return 1 in initial versions")
	}
}

// ut1 OMIT
func TestUpdateTime(t *testing.T) {
	clk = ClockFake{} // HL
	j := Job{}
	qn := 1
	lastVirtFinish[qn] = time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC)
	updateTime(&j, qn)
	if !j.VirtFinish.Equal(time.Date(2000, 1, 1, 0, 0, 10, 0, time.UTC)) {
		t.Error("should be 10 seconds later")
		t.Error(j.VirtFinish)
	}
	if !lastVirtFinish[qn].Equal(j.VirtFinish) {
		t.Error("times should be equal")
	}
}

// ut2 OMIT

func TestReceive(t *testing.T) {
	j := Job{Id: "j1"}
	receive(&j)
	qn := 1
	if queues[qn].Len() != 1 {
		t.Error("Incorrect queueing")
	}

	jout, _ := queues[qn].Dequeue()
	if j.Id != jout.(Job).Id {
		t.Error("incorrect job queued")
	}
}

func TestSelectQueue(t *testing.T) {
	j1 := Job{Id: "j1"}
	j2 := Job{Id: "j2"}
	receive(&j1)
	defer send()
	receive(&j2)
	defer send()
	qn := selectQueue()
	if qn != 1 {
		t.Error("incorrect queue number returned")
	}
}

// 01 OMIT
func TestSend(t *testing.T) {
	j1 := Job{Id: "j1"}
	j2 := Job{Id: "j2"}
	receive(&j1) // HL
	receive(&j2) // HL
	js := send()
	if js.Id != j1.Id { // HL
		t.Error("should be j1 that is sent first")
	}

	if send().Id != j2.Id { // HL
		t.Error("should be j2 that is sent next")
	}
}

// 02 OMIT
func TestSendUnderflow(t *testing.T) {
	j1 := Job{Id: "j1"}
	receive(&j1)            // HL
	if send().Id != j1.Id { // HL
		t.Error("should be j1 that is sent first")
	}

	js := send()
	if js.Id != "idle_job" { // HL
		t.Errorf("should be idle %v\n", js)

	}

}

// 03 OMIT

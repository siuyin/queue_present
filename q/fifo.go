package q

type FIFO struct {
	s []interface{}
}

func NewFIFO() *FIFO {
	f := new(FIFO)
	f.s = make([]interface{}, 0)
	return f
}

func (e *FIFO) Enqueue(i interface{}) error {
	e.s = append(e.s, i)
	return nil
}

func (e *FIFO) Dequeue() (interface{}, error) {
	if e.Len() == 0 {
		return nil, &QError{1, "Queue is empty"}
	}
	v := e.s[0]
	e.s[0] = nil
	e.s = e.s[1:]
	return v, nil
}

func (e FIFO) Len() int {
	return len(e.s)
}

type QError struct {
	I int
	M string
}

func (e QError) Error() string {
	return e.M
}

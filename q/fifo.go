package q

// 01 OMIT
type FIFO struct {
	s []interface{}
}

func NewFIFO() *FIFO {
	f := new(FIFO)
	f.s = make([]interface{}, 0)
	return f
}

// 02 OMIT
func (e *FIFO) Enqueue(i interface{}) error { // HL01
	e.s = append(e.s, i)
	return nil
}

// 05 OMIT
func (e *FIFO) Dequeue() (interface{}, error) { // HL01
	if e.Len() == 0 {
		return nil, &QError{"Queue is empty"} // HL02
	}
	v := e.s[0]
	e.s[0] = nil // prevents memory leak
	e.s = e.s[1:]
	return v, nil
}

// 06 OMIT
func (e FIFO) Len() int { // HL01
	return len(e.s)
}

// 03 OMIT
type QError struct { // HL02
	Msg string
}

func (e QError) Error() string { // HL02
	return e.Msg
}

// 07 OMIT

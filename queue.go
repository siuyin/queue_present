package q

type Queuer interface {
	Enqueue() error
	Dequeueue() (interface{}, error)
	Len() int
}

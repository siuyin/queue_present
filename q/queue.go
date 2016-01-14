package q

type Queuer interface {
	Enqueue(interface{}) error
	Dequeueue() (interface{}, error)
	Len() int
}

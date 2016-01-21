package q

type Queuer interface {
	Enqueue(interface{}) error
	Dequeueue() (interface{}, error)
	Head() (interface{}, error)
	Len() int
}

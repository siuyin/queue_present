package q

type Queue interface {
	Enqueue(interface{}) error
	Dequeueue() (interface{}, error)
	Head() (interface{}, error)
	Len() int
}

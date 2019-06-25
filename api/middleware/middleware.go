package middleware

import (
	"fmt"

	"github.com/kataras/iris"
)

type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

// Que Reservar memoria
var Que []*Queue

// Repository should implement common methods
type Repository interface {
	Read() []*Queue
}

func (q *Queue) Read() []*Queue {
	return MockQueue()
}

// MockQueue For Testing
func MockQueue() []*Queue {
	return []*Queue{
		{
			Domain:   "alpha",
			Weight:   5,
			Priority: 5,
		},
		{
			Domain:   "omega",
			Weight:   1,
			Priority: 5,
		},
		{
			Domain:   "beta",
			Weight:   5,
			Priority: 1,
		},
	}
}

// InitQueue initializes the Queue
func InitQueue() {
	Que = append(Que, &Queue{})
}

// ProxyMiddleware kdk
func ProxyMiddleware(c iris.Context) {
	c.GetHeader("domain")
	var repo Repository
	repo = &Queue{}
	fmt.Println(repo.Read())
}

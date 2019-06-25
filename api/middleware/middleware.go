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
		{
			Domain:   "alpha",
			Weight:   10,
			Priority: 6,
		},
		{
			Domain:   "beta",
			Weight:   1,
			Priority: 4,
		},
	}
}

// InitQueue initializes the Queue
func InitQueue() {
	Que = append(Que, &Queue{})
}

func onLowSegment(priority int, weight int) bool {
	x, y := priority, weight
	return x < 5 && y < 5
}

func onMediumSegment(priority int, weight int) bool {
	x, y := priority, weight
	return x >= 5 && y <= 5 || y >= 5 && x <= 5
}

func onHighSegment(priority int, weight int) bool {
	x, y := priority, weight
	return x >= 6 && y >= 6
}

// ProxyMiddleware kdk
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	var repo Repository
	repo = &Queue{}

	mediumPriority := []*Queue{}
	lowPriority := []*Queue{}

	fmt.Println("From header ", domain)
	for _, row := range repo.Read() {
		fmt.Println("From source ", row.Domain)
		if onHighSegment(row.Priority, row.Weight) {
		} else if onMediumSegment(row.Priority, row.Weight) {
			mediumPriority = append(mediumPriority, row)
		} else if onLowSegment(row.Priority, row.Weight) {
			lowPriority = append(lowPriority, row)
		}
	}

	for _, s := range mediumPriority {
		fmt.Print("MEdum ", s)
	}

	c.Next()
}

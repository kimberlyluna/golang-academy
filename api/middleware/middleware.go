package middleware

import (
	"sort"

	"github.com/kataras/iris"
	"github.com/kimberly.luna/proxy-app/api/utils"
)

// Node used for sorting
type Node struct {
	Domain   string
	Priority int
}

// Queue declaration
var Queue []string

// nodeQueue a temporary node array used for sorting by priority
var nodeQueue = []*Node{}

// ProxyMiddleware should queue our incoming requests
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")

	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "Domain error"})
		return
	}

	// Add a new node into the the queue
	nodeQueue = append(nodeQueue, &Node{domain, utils.GetDomainPriority(domain)})

	// Sorts nodeQueue decreasingly from highest priority to lowest priority
	sort.Slice(nodeQueue, func(i, j int) bool {
		return nodeQueue[i].Priority > nodeQueue[j].Priority
	})

	// 'Map' the domain names after sorting
	Queue = Queue[:0]
	for _, node := range nodeQueue {
		Queue = append(Queue, node.Domain)
	}

	// Doubt: then if a request gets solved it should be removed into both the nodeQueue and the Queue to avoid possible conflicts

	c.Next()
}

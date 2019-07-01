package middleware

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	"github.com/kataras/iris"
)

// Queue
type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

type Node struct {
	Domain   string
	Priority int
}

// Que declaration
var Que []string

// Repository should implement common methods
type Repository interface {
	Read() []*Queue
}

func (q *Queue) Read() []*Queue {
	path, _ := filepath.Abs("")
	file, err := os.Open(path + "/api/middleware/domain.txt")

	var dataStructure = []*Queue{}
	tempNode := Queue{}

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			//reading a whitespace
			continue
		} else {
			currentLine := scanner.Text()

			isWeight, _ := regexp.MatchString("weight", currentLine)
			isPriority, _ := regexp.MatchString("priority", currentLine)

			numRegex := regexp.MustCompile("[0-9]")

			if isWeight {
				tempNode.Weight, _ = strconv.Atoi(numRegex.FindString(currentLine))
			} else if isPriority {
				tempNode.Priority, _ = strconv.Atoi(numRegex.FindString(currentLine))
				newNode := Queue{tempNode.Domain, tempNode.Weight, tempNode.Priority}
				dataStructure = append(dataStructure, &newNode)
			} else {
				tempNode.Domain = currentLine
			}
		}
	}
	return dataStructure
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

// ProxyMiddleware should queue our incoming requests
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")

	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	var repository Repository
	repository = &Queue{}

	var nodeQueue = []*Node{}
	var priority int

	fmt.Println("@#@#$#$", domain)

	for _, row := range repository.Read() {
		if onLowSegment(row.Priority, row.Weight) {
			priority = 1
		} else if onMediumSegment(row.Priority, row.Weight) {
			priority = 2
		} else if onHighSegment(row.Priority, row.Weight) {
			priority = 3
		}
		nodeQueue = append(nodeQueue, &Node{row.Domain, priority})
	}

	sort.Slice(nodeQueue, func(i, j int) bool {
		return nodeQueue[i].Priority > nodeQueue[j].Priority
	})

	for _, node := range nodeQueue {
		Que = append(Que, node.Domain)
	}

	c.Next()
}

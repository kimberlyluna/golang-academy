package middleware

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/kataras/iris"
)

// Queue
type Queue struct {
	Domain   string
	Weight   int
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
			fmt.Println("OUT", scanner.Text())
			//reading a whitespace
			continue
		}
		currentLine := scanner.Text()
		isWeight, _ := regexp.MatchString("weight", currentLine)
		isPriority, _ := regexp.MatchString("priority", currentLine)

		reg := regexp.MustCompile("[0-9]")

		if isWeight {
			tempNode.Weight, _ = strconv.Atoi(reg.FindString(currentLine))
		} else if isPriority {

			// add to dataStructure
			tempNode.Priority, _ = strconv.Atoi(reg.FindString(currentLine))
			newNode := Queue{tempNode.Domain, tempNode.Weight, tempNode.Priority}
			dataStructure = append(dataStructure, &newNode)

		} else {
			// is domain name
			tempNode.Domain = currentLine

		}
		fmt.Println(" ")
	}
	return dataStructure
}

// MockQueue should mock an Array of Queues
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

// ProxyMiddleware should queue our incoming requests
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	var repo Repository
	repo = &Queue{}
	fmt.Println("FROM HEADER", domain)

	// read from texfile
	for _, row := range repo.Read() {
		fmt.Println("hey ", row.Domain, row.Priority, row.Weight)
	}

	// El queue final
	Que = append(Que, domain)

	c.Next()
}

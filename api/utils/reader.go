package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// Queue a queue node
type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

// Repository should implement common methods
type Repository interface {
	Read() []*Queue
}

// Read reads from a textfile
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
			// reading a text line

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

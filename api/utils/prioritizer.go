package utils

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

// CalculatePriority receives an x and y value and maps the coordinate it into a segment
func CalculatePriority(priority int, weight int) int {
	var calculatedPriority int
	if onLowSegment(priority, weight) {
		calculatedPriority = 1
	} else if onMediumSegment(priority, weight) {
		calculatedPriority = 2
	} else if onHighSegment(priority, weight) {
		calculatedPriority = 3
	}
	return calculatedPriority
}

// GetDomainPriority receives a domain name and returns a priority number
// 1 - lowest, 3 - highest
func GetDomainPriority(domain string) int {
	var repository Repository
	repository = &Queue{}

	for _, row := range repository.Read() {
		if domain == row.Domain {
			return CalculatePriority(row.Priority, row.Weight)
		}
	}
	return 0
}

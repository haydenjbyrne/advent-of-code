package ParseUtils

import (
	"log"
	"strconv"
	"strings"
)

func ParseNumbers(s string) []int {
	var numbers []int
	rawNumbers := strings.Fields(s)
	for _, rawNumber := range rawNumbers {
		number, err := strconv.Atoi(rawNumber)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

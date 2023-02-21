package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func LargeNumberToReadingString(numberString string) string {
	numberString = strings.ReplaceAll(numberString, ",", "")
	numberString = strings.ReplaceAll(numberString, " ", "")
	numberString = strings.ReplaceAll(numberString, "$", "")
	numberString = strings.ReplaceAll(numberString, "%", "")

	number, err := strconv.ParseFloat(numberString, 64)
	if err != nil {
		return numberString
	}

	if number >= 1000000000 {
		return fmt.Sprintf("%.2f billion", number/1000000000)
	} else if number >= 1000000 {
		return fmt.Sprintf("%.2f million", number/1000000)
	} else if number >= 1000 {
		return fmt.Sprintf("%.2f thousand", number/1000)
	} else {
		return numberString
	}
}

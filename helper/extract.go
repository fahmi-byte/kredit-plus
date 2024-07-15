package helper

import (
	"fmt"
	"regexp"
	"strconv"
)

func ExtractNumber(s string) (int, error) {
	re := regexp.MustCompile(`CTR-(\d+)/\d+`)
	matches := re.FindStringSubmatch(s)
	var number int

	if len(matches) > 1 {
		// Konversi string angka menjadi integer
		result, err := strconv.Atoi(matches[1])
		PanicIfError(err)
		number = result
	} else {
		fmt.Println("No match found")
	}
	return number, nil
}

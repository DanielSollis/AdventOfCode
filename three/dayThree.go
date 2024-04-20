package three

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func PartOne() (sum int, err error) {
	text, _ := os.ReadFile("three/input.txt")
	matrix := strings.Split(string(text), "\n")

	symbolFound := false
	numStartI, numStartj, numLen := 0, 0, 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			space := rune(matrix[i][j])

			if unicode.IsDigit(space) {
				if numLen == 0 {
					numStartI = i
					numStartj = j
				}
				numLen++
				for isquare := i - 1; isquare <= i+1; isquare++ {
					for jsquare := j - 1; jsquare <= j+1; jsquare++ {
						if isquare >= 0 && isquare < len(matrix) &&
							jsquare >= 0 && jsquare < len(matrix[i]) {

							checkSpace := rune(matrix[isquare][jsquare])
							if !unicode.IsDigit(checkSpace) && checkSpace != '.' {
								symbolFound = true
							}
						}
					}
				}
			} else {
				if numLen > 0 && symbolFound {
					var number int
					numstring := matrix[numStartI][numStartj : numStartj+numLen]
					if number, err = strconv.Atoi(numstring); err != nil {
						return 0, err
					}
					sum += number
				}
				numLen = 0
				symbolFound = false
			}
		}
	}
	return sum, nil
}

func PartTwo() (sum int, err error) {
	return 0, nil
}

package three

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func PartOne() (result int, err error) {
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
					result += number
				}
				numLen = 0
				symbolFound = false
			}
		}
	}
	return result, nil
}

type location struct {
	i int
	j int
}

func PartTwo() (result int, err error) {
	text, _ := os.ReadFile("three/input.txt")
	matrix := strings.Split(string(text), "\n")

	LocationsToNums := map[location]int{}
	gearLocations := []location{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			space := rune(matrix[i][j])
			if space == '*' {
				gearLocations = append(gearLocations, location{i, j})
			} else if unicode.IsDigit(space) {
				foundDigitBytes := []byte{matrix[i][j]}
				foundLocations := []location{{i, j}}
				for k := j + 1; k < len(matrix[i]); k++ {
					if unicode.IsDigit(rune(matrix[i][k])) {
						foundDigitBytes = append(foundDigitBytes, matrix[i][k])
						foundLocations = append(foundLocations, location{i, k})
					} else {
						j = k
						break
					}
				}
				var foundDigit int
				if foundDigit, err = strconv.Atoi(string(foundDigitBytes)); err != nil {
					return 0, err
				}
				for _, location := range foundLocations {
					LocationsToNums[location] = foundDigit
				}
			}
		}
	}

	for _, gear := range gearLocations {
		numbers := map[int]struct{}{}
		for i := gear.i - 1; i < gear.i+2; i++ {
			for j := gear.j - 1; j < gear.j+2; j++ {
				if val, ok := LocationsToNums[location{i, j}]; ok {
					if _, ok := numbers[val]; !ok {
						numbers[val] = struct{}{}
					}
				}
			}
		}
		if len(numbers) == 2 {
			gearRatio := 1
			for num := range numbers {
				gearRatio *= num
			}
			result += gearRatio
		}
	}
	return result, nil
}

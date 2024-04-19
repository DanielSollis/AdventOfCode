package two

import (
	"os"
	"strconv"
	"strings"
)

func PartOne() (sum int, err error) {
	var text []byte
	if text, err = os.ReadFile("two/input.txt"); err != nil {
		return 0, err
	}
	lines := strings.Split(string(text), "\n")

	for inx, line := range lines {
		tokens := strings.FieldsFunc(line, split)[2:]
		sum += inx + 1
		for i := 0; i < len(tokens)-1; i += 2 {
			var numberOfCubes int
			if numberOfCubes, err = strconv.Atoi(tokens[i]); err != nil {
				return 0, err
			}
			color := tokens[i+1]

			if color == "red" {
				if numberOfCubes > 12 {
					sum -= inx
					break
				}
			} else if color == "green" {
				if numberOfCubes > 13 {
					sum -= inx
					break
				}
			} else if color == "blue" {
				if numberOfCubes > 14 {
					sum -= inx
					break
				}
			}
		}
	}
	return sum, nil
}

func split(r rune) bool {
	return r == ':' || r == ';' || r == ',' || r == ' '
}

func PartTwo() {

}

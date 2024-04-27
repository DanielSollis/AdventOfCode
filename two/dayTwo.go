package two

import (
	"os"
	"strconv"
	"strings"
)

func PartOne() (result int, err error) {
	var text []byte
	if text, err = os.ReadFile("two/input.txt"); err != nil {
		return 0, err
	}
	lines := strings.Split(string(text), "\n")

	for inx, line := range lines {
		gameId := inx + 1
		split := func(r rune) bool {
			return r == ':' || r == ';' || r == ',' || r == ' '
		}
		tokens := strings.FieldsFunc(line, split)[2:]

		result += gameId
		for i := 0; i < len(tokens)-1; i += 2 {
			var numberOfCubes int
			if numberOfCubes, err = strconv.Atoi(tokens[i]); err != nil {
				return 0, err
			}
			color := tokens[i+1]

			if color == "red" {
				if numberOfCubes > 12 {
					result -= gameId
					break
				}
			} else if color == "green" {
				if numberOfCubes > 13 {
					result -= gameId
					break
				}
			} else if color == "blue" {
				if numberOfCubes > 14 {
					result -= gameId
					break
				}
			}
		}
	}
	return result, nil
}

func PartTwo() (result int, err error) {
	var text []byte
	if text, err = os.ReadFile("two/input.txt"); err != nil {
		return 0, err
	}
	lines := strings.Split(string(text), "\n")

	for _, line := range lines {
		split := func(r rune) bool {
			return r == ':' || r == ';' || r == ',' || r == ' '
		}
		tokens := strings.FieldsFunc(line, split)[2:]

		maxRed, maxBlue, maxGreen := 0, 0, 0
		for i := 0; i < len(tokens)-1; i += 2 {
			var numberOfCubes int
			if numberOfCubes, err = strconv.Atoi(tokens[i]); err != nil {
				return 0, err
			}
			color := tokens[i+1]

			if color == "red" {
				if maxRed < numberOfCubes {
					maxRed = numberOfCubes
				}
			} else if color == "green" {
				if maxBlue < numberOfCubes {
					maxBlue = numberOfCubes
				}
			} else if color == "blue" {
				if maxGreen < numberOfCubes {
					maxGreen = numberOfCubes
				}
			}
		}
		result += maxRed * maxGreen * maxBlue
	}
	return result, nil
}

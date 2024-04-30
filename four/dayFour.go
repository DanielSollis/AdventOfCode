package four

import (
	"os"
	"strings"
)

func PartOne() (result int, err error) {
	var text []byte
	if text, err = os.ReadFile("four/input.txt"); err != nil {
		return 0, err
	}

	split := func(r rune) bool {
		return r == '\n' || r == '|'
	}
	lines := strings.FieldsFunc(string(text), split)

	for inx := 0; inx < len(lines); inx += 2 {
		haveSplit := func(r rune) bool {
			return r == ':' || r == ' '
		}
		need := strings.FieldsFunc(lines[inx], haveSplit)[2:]
		have := strings.Split(lines[inx+1], " ")

		haveSet := map[string]struct{}{}
		for _, num := range have {
			haveSet[num] = struct{}{}
		}

		sum := 0
		for _, num := range need {
			if _, ok := haveSet[num]; ok {
				if sum == 0 {
					sum = 1
				} else {
					sum *= 2
				}
			}
		}
		result += sum
	}

	return result, nil
}

func PartTwo() (result int, err error) {
	var text []byte
	if text, err = os.ReadFile("four/input.txt"); err != nil {
		return 0, err
	}
	lines := strings.Split(string(text), "\n")

	cardCounts := map[int]int{}
	for inx := range lines {
		cardCounts[inx] = 1
	}

	for inx, line := range lines {
		splitfunc := func(r rune) bool {
			return r == ':' || r == ' ' || r == '|'
		}
		splitLine := strings.FieldsFunc(line, splitfunc)[2:]

		need := splitLine[:10]
		have := splitLine[10:]

		haveSet := map[string]struct{}{}
		for _, val := range have {
			haveSet[val] = struct{}{}
		}

		count := 0
		for _, val := range need {
			if _, ok := haveSet[val]; ok {
				count += 1
			}
		}

		for i := inx + 1; i <= inx+count; i++ {
			cardCounts[i] += cardCounts[inx]
		}

	}

	for _, val := range cardCounts {
		result += val
	}

	return result, nil
}

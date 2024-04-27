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

func PartTwo() {

}

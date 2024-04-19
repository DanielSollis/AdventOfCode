package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var err error
	var result int
	if result, err = dayOnePartOne("input.txt"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	if result, err = dayOnePartTwo("input.txt"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func dayOnePartOne(input string) (sum int, err error) {
	var file *os.File
	if file, err = os.Open(input); err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		digits := []byte{}
		var firstDigitInx int
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				digits = append(digits, line[i])
				break
			}
		}

		for i := len(line) - 1; i >= firstDigitInx-1; i-- {
			if unicode.IsDigit(rune(line[i])) {
				digits = append(digits, line[i])
				break
			}
		}

		var digit int
		if digit, err = strconv.Atoi(string(digits)); err != nil {
			return 0, err
		}
		sum += digit
	}
	return sum, nil
}

var digitMap = map[int]map[string]int{
	3: {
		"one": 1,
		"two": 2,
		"six": 6,
	},
	4: {
		"four": 4,
		"five": 5,
		"nine": 9,
	},
	5: {
		"three": 3,
		"seven": 7,
		"eight": 8,
	},
}

func dayOnePartTwo(input string) (sum int, err error) {
	var text []byte
	if text, err = os.ReadFile(input); err != nil {
		return 0, err
	}

	lines := strings.Split(string(text), "\n")
	for _, line := range lines {
		digits := []byte{}
		firstDigitInx := 0
		for i := 0; i < len(line); i++ {
			var found bool
			if digits, found = checkForNum(line, digits, i); found {
				firstDigitInx = i
				break
			}
		}

		for i := len(line) - 1; i > firstDigitInx-1; i-- {
			var found bool
			if digits, found = checkForNum(line, digits, i); found {
				break
			}
		}

		var digit int
		if digit, err = strconv.Atoi(string(digits)); err != nil {
			panic(err)
		}
		sum += digit
	}

	return sum, nil
}

func checkForNum(line string, digits []byte, inx int) (_ []byte, found bool) {
	if unicode.IsDigit(rune(line[inx])) {
		digits = append(digits, byte(line[inx]))
		return digits, true
	} else {
		lengthLeft := len(line) - inx
		switch {
		case lengthLeft >= 5:
			toFind := line[inx : inx+5]
			if val, ok := digitMap[5][toFind]; ok {
				digits = append(digits, byte(val)+'0')
				return digits, true
			}
			fallthrough

		case lengthLeft >= 4:
			toFind := line[inx : inx+4]
			if val, ok := digitMap[4][toFind]; ok {
				digits = append(digits, byte(val)+'0')
				return digits, true
			}
			fallthrough

		case lengthLeft >= 3:
			toFind := line[inx : inx+3]
			if val, ok := digitMap[3][toFind]; ok {
				digits = append(digits, byte(val)+'0')
				return digits, true
			}
		}
	}
	return digits, false
}

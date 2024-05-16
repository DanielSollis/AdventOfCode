package six

import (
	"os"
	"strconv"
	"strings"
)

func PartOne() (result int, err error) {
	var inputText []byte
	if inputText, err = os.ReadFile("six/input.txt"); err != nil {
		return 0, err
	}

	inputlines := strings.Split(string(inputText), "\n")

	temp := strings.Split(inputlines[0], " ")
	times := []int{}
	for i := 0; i < len(temp); i++ {
		if val, err := strconv.Atoi(temp[i]); err == nil {
			times = append(times, val)
		}
	}

	temp = strings.Split(inputlines[1], " ")
	distances := []int{}
	for i := 0; i < len(temp); i++ {
		if val, err := strconv.Atoi(temp[i]); err == nil {
			distances = append(distances, val)
		}
	}

	result = 1
	for i := 0; i < len(times); i++ {
		count := 0
		for j := 0; j < times[i]; j++ {
			if j*(times[i]-j) > distances[i] {
				count++
			}
		}
		result *= count
	}

	return result, nil
}

func PartTwo() (result int, err error) {
	var inputText []byte
	if inputText, err = os.ReadFile("six/input.txt"); err != nil {
		return 0, err
	}

	lines := strings.Split(string(inputText), "\n")

	var time int
	splitTimes := strings.Split(lines[0], ":")[1]
	timeString := strings.ReplaceAll(splitTimes, " ", "")
	if time, err = strconv.Atoi(timeString); err != nil {
		return 0, err
	}

	var distance int
	splitDistance := strings.Split(lines[1], ":")[1]
	distanceString := strings.ReplaceAll(splitDistance, " ", "")
	if distance, err = strconv.Atoi(distanceString); err != nil {
		return 0, err
	}

	start := 0
	for i := 0; i < time; i++ {
		if i*(time-i) > distance {
			start = i
			break
		}
	}

	stop := 0
	for i := time; i > 0; i-- {
		if i*(time-i) > distance {
			stop = i
			break
		}
	}
	return stop - start + 1, nil
}

package main

import (
	freqCount "Training_GoLanguage/IntegerCount/freqCount"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputVal string
	if len(os.Args) > 1 {
		inputVal = os.Args[1]
	} else {
		input, err := getInputFromUser()
		if err != nil {
			fmt.Println("Received Error:", err)
		}
		inputVal = input
	}

	seriesOfDigits, err := freqCount.ValidateIntegerInput(inputVal)
	if err != nil {
		fmt.Println("Received Error:", err)
	} else {
		generateOutput(seriesOfDigits)
	}
}

func getInputFromUser() (string, error) {
	fmt.Println("Enter the Input String of Numbers")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("Error Reading the Input from the User")
	}
	inputVal := strings.TrimSpace(input)
	return inputVal, nil
}

func generateOutput(seriesOfDigits int) {
	fmt.Println("Output - ")
	for i := 0; i <= 9; i++ {
		copySeries := seriesOfDigits
		integerSlice := make([]int, 0)
		for copySeries > 0 {
			digit := copySeries % 10
			copySeries /= 10
			if digit == i {
				integerSlice = append(integerSlice, digit)
			}
		}
		fmt.Println("     The digit " + strconv.Itoa(i) + " occurs " + strconv.Itoa(len(integerSlice)) + " time(s)")
	}
}

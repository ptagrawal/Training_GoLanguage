package validateInput

import (
	"fmt"
	"strconv"
)

func ValidateIntegerInput(inputVal string) (int, error) {
	digits, err := strconv.Atoi(inputVal)
	if err != nil {
		return 0, fmt.Errorf("Invalid input. Program Exited")
	}
	return digits, nil
}

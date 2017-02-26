package validateInput

import (
	"testing"
)

func TestValidateIntegerInputSuccess(t *testing.T) {
	input := "1234567890"

	actual, _ := ValidateIntegerInput(input)
	expected := 1234567890

	if actual != expected {
		t.Error("Expected: %d, Got: %v", expected, actual)
	}
}

func TestValidateIntegerInputFailure(t *testing.T) {
	input := "asgasde"

	_, err := ValidateIntegerInput(input)

	if err == nil {
		t.Error("Expected an error for Invalid Input")
	}
}

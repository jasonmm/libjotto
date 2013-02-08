package libjotto

import "testing"

func TestGuessResult(t *testing.T) {
	result := GuessResult("sum", "sun")

	if result != 2 {
		t.Errorf("GuessResult returned", result, ". Should have been 2.")
	}
}

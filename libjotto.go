package libjotto


import (
	"container/list"
	"strings"
)

// Compares the guess to the given secret word and returns the number of 
// letters the guess got correct and the number of letters the guess put
// in the correct place.
func GuessResult(guess string, secretWord string) int {
	var cnt int = 0

	for i := 0; i < len(guess); i++ {
		if strings.Contains(secretWord, guess[i:i+1]) {
			cnt++
		}
	}
	return cnt
}

// Determines what the next guess should be based on past guesses and their
// results.
func NextGuess(pastGuesses *list.List) string {
	return ""
}



libjotto
========

A support library for Jotto games.

Functions
---------

**func GuessResult(guess string, secretWord string) int**

> Compares the guess to the given secret word and returns the number of
> letters the guess got correct and the number of letters the guess put in
> the correct place.

**func NextGuess(pastGuesses \*list.List) string**
    
> Determines what the next guess should be based on past guesses and their
> results.


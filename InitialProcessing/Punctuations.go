package Module

import (
	"strings"
)

func PunctuationProcessing(Words []string) []rune {
	Punctuations := []rune{',', ';', ':', '!', '?', '.'}

	String := strings.Join(Words, "\n")
	Letters := []rune(String)
	for i := 0; i < len(Letters); i++ {
		for j := 0; j < len(Punctuations); j++ {
			// if char is a punctuation mark
			if Letters[i] == Punctuations[j] {
				// delete spaces before
				if i > 0 && Letters[i-1] == ' ' {
					Letters[i-1] = '\x00'
				}
				// add space after
				if i < len(Letters)-1 && Letters[i+1] != ' ' && Letters[i+1] != '\n' {
					Letters = append(Letters[:i+1], append([]rune{' '}, Letters[i+1:]...)...)
				}
			}
		}
	}

	return Letters
}

package Module

func SingleQuotesProcessing(Letters []rune) string {
	var inQuote bool
	quoteStart := -1 // -1 can't be an index
	for j := 0; j < len(Letters); j++ {
		// if single quote between two none space chars, it's an apostroph ignore it
		if Letters[j] == '\'' {
			if j-1 >= 0 && j+1 < len(Letters)-1 && j != len(Letters)-1 && (Letters[j+1] != ' ' && Letters[j-1] != ' ') && Letters[j-1] != '\n' && Letters[j+1] != '\n' {
				continue
			}

			if !inQuote {
				// if inquote is false and you just found a single quote switch it to true and save the value of j
				quoteStart = j
				inQuote = true
				// fmt.Println(" step 1 : inquote false became true, j conserved")

			} else {
				// opening quote
				// add space before opening quote if there's none, don't if it comes right after a newline
				if quoteStart > 0 && Letters[quoteStart-1] != '\n' && Letters[quoteStart-1] != ' ' {
					Letters = append(Letters[:quoteStart], append([]rune{' '}, Letters[quoteStart:]...)...)
					j++
					quoteStart = quoteStart + 1
					// fmt.Println("step 2 : append space before opening quote", quoteStart, inQuote, string(Letters[quoteStart+1]))

				}
				// delete space after opening quote
				if quoteStart+1 < len(Letters) && Letters[quoteStart+1] == ' ' {
					Letters = append(Letters[:quoteStart+1], Letters[quoteStart+2:]...)
					j--
					// fmt.Println("step : 3 delete space after opening quote", quoteStart, inQuote, "letters quotestart = ", string(Letters[quoteStart]), "letters quotestart +1 =", string(Letters[quoteStart+1]))

				}

				// closing quote
				// delete space before closing quote
				if j > 0 && Letters[j-1] == ' ' {
					Letters = append(Letters[:j-1], Letters[j:]...)
					j--
					// fmt.Println("step : 4 delete space before closing quote", j, inQuote, "letters quotestart = ", string(Letters[quoteStart]), "letters quotestart +1 =", string(Letters[quoteStart+1]))

				}
				// add space after closing quote
				if j+2 < len(Letters) && Letters[j+1] != ' ' && Letters[j+1] != '\n' && (j+2 >= len(Letters) || (Letters[j+2] != ',' && Letters[j+2] != ';' && Letters[j+2] != '.' && Letters[j+2] != ':' && Letters[j+2] != '?' && Letters[j+2] != '!')) {
					Letters = append(Letters[:j+1], append([]rune{' '}, Letters[j+1:]...)...)
					// fmt.Println("step 5 : adding space after single", quoteStart, inQuote, "letters quotestart = ", string(Letters[quoteStart]), "letters quotestart +1 =", string(Letters[quoteStart+1]))
					j++
				}

				inQuote = false
				quoteStart = -1
			}
		}
		// reset in newlines
		if j < len(Letters)-1 && Letters[j] == '\n' {
			inQuote = false
			quoteStart = -1
		}
		// remove null char
		if len(Letters) > 1 && j < len(Letters) && Letters[j] == '\x00' {
			Letters = append(Letters[:j], Letters[j+1:]...)
			j--
		}
	}

	return string(Letters)
}

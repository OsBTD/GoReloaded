package Module

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// this function checks if a strings only contains punctuation marks, if so it returns true
func ContainsOnly(char string) bool {
	for i := 0; i < len(char); i++ {
		if !strings.ContainsAny(string(char[i]), ",;:!?.'\"") {
			return false
		}
	}
	return true
}

func IsFlag(Word string) bool {
	if Word == "(up)" || Word == "(cap)" || Word == "(low)" || Word == "(hex)" || Word == "(bin)" || strings.HasPrefix(Word, "(up,") || strings.HasPrefix(Word, "(cap,") || strings.HasPrefix(Word, "(low,") || strings.HasPrefix(Word, "(hex,") || strings.HasPrefix(Word, "(bin,") {
		return true
	} else {
		return false
	}
}

func FlagProcessing() []string {
	counter := 1
	FileName := os.Args[1]

	content, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatal("Error : unable to read file", err)
	}
	Input := string(content)
	Lines := strings.Split(Input, "\n")
	var Words []string
	var Processed []string
	for _, Line := range Lines {
		Words = strings.Fields(Line)
		for i := 0; i < len(Words)-1; i++ {
			if len(Words) > 1 && (Words[i] == "a" || Words[i] == "A") {
				if strings.ContainsAny(string(Words[i+1][0]), "aeoiuhAEOIUH") {
					Words[i] += "n"
				}
			}
			if len(Words) > 1 && len(Words[i]) > 1 && strings.ContainsAny(string(Words[i][len(Words[i])-1]), "Aa") && strings.ContainsAny(string(Words[i][:len(Words[i])-1]), ",;:?!'(\"") {
				if strings.ContainsAny(string(Words[i+1][0]), "aeoiuhAEOIUH") {
					Words[i] += "n"
				}
			}
		}

		for i := 0; i < len(Words); i++ {
			// this checks if a string is a flag to avoid adding spaces inside of it
			// if a string only contains punctuation marks increment counter
			if ContainsOnly(Words[i]) {
				counter++
			} else if !IsFlag(Words[i]) {
				counter = 1
			}
			// reset counter after newlines
			if i == 0 {
				counter = 1
			}

			// handle flags while ignoring punctuation marks
			switch {
			case Words[i] == "(up,":
				if i+1 < len(Words) && strings.HasSuffix(Words[i+1], ")") {
					num := strings.TrimSuffix(Words[i+1], ")")
					x, err := strconv.Atoi(num)
					if err != nil || x < 0 {
						fmt.Println("Error : ", Words[i+1], " invalid flag number")
						continue
					}
					start := i - x - (counter - 1)
					if start < 0 {
						start = 0
					}
					for j := start; j < i; j++ {
						Words[j] = strings.ToUpper(Words[j])
					}
				}
				Words = append(Words[:i], Words[i+2:]...)
				i--
			case Words[i] == "(up)":
				if i-counter >= 0 {
					Words[i-counter] = strings.ToUpper(Words[i-counter])
				}
				Words = append(Words[:i], Words[i+1:]...)
				i--
			case Words[i] == "(cap,":
				if i+1 < len(Words) && strings.HasSuffix(Words[i+1], ")") {
					num := strings.TrimSuffix(Words[i+1], ")")
					x, err := strconv.Atoi(num)
					if err != nil || x < 0 {
						fmt.Println("Error : ", Words[i+1], " invalid flag number")
						continue
					}
					start := i - x - (counter - 1)
					if start < 0 {
						start = 0
					}
					for j := start; j < i; j++ {
						if len(Words[j]) > 0 {
							firstRune, body := utf8.DecodeRuneInString(Words[j])
							Words[j] = strings.ToUpper(string(firstRune)) + strings.ToLower(Words[j][body:])
						}
					}
					Words = append(Words[:i], Words[i+2:]...)
					i--

				}
			case Words[i] == "(cap)":
				if i-counter >= 0 {
					if len(Words[i-counter]) > 0 {
						firstRune, body := utf8.DecodeRuneInString(Words[i-counter])
						Words[i-counter] = strings.ToUpper(string(firstRune)) + strings.ToLower(Words[i-counter][body:])
					}
				}
				Words = append(Words[:i], Words[i+1:]...)
				i--
			case Words[i] == "(low,":
				if i+1 < len(Words) && strings.HasSuffix(Words[i+1], ")") {
					num := strings.TrimSuffix(Words[i+1], ")")
					x, err := strconv.Atoi(num)
					if err != nil || x < 0 {
						fmt.Println("Error : ", Words[i+1], " invalid flag number")
						continue
					}
					start := i - x - (counter - 1)
					if start < 0 {
						start = 0
					}
					for j := start; j < i; j++ {
						Words[j] = strings.ToLower(Words[j])
					}
					Words = append(Words[:i], Words[i+2:]...)
					i--

				}
			case Words[i] == "(low)":
				if i-counter >= 0 {
					Words[i-counter] = strings.ToLower(Words[i-counter])
				}
				Words = append(Words[:i], Words[i+1:]...)
				i--
			case Words[i] == "(hex,":
				if i+1 < len(Words) && strings.HasSuffix(Words[i+1], ")") {
					num := strings.TrimSuffix(Words[i+1], ")")
					x, err := strconv.Atoi(num)
					if err != nil || x < 0 {
						fmt.Println("Error : ", Words[i+1], " invalid flag number")
						continue
					}
					start := i - x - (counter - 1)
					if start < 0 {
						start = 0
					}
					for j := start; j < i; j++ {
						y, err := strconv.ParseInt(Words[j], 16, 64)
						if err != nil {
							fmt.Println("Error : ", Words[j], " this hexadecimal value is incorrect, it can't be converted")
							continue
						} else {
							Words[j] = strconv.FormatInt(y, 10)
						}

					}
					Words = append(Words[:i], Words[i+2:]...)
					i--

				}
			case Words[i] == "(hex)":
				if i-counter >= 0 {
					y, err := strconv.ParseInt(Words[i-counter], 16, 64)
					if err != nil {
						fmt.Println("Error : ", Words[i-counter], " this hexadecimal value is incorrect, it can't be converted")
						continue
					} else {
						Words[i-counter] = strconv.FormatInt(y, 10)
					}
				}

				Words = append(Words[:i], Words[i+1:]...)
				i--
			case Words[i] == "(bin,":
				if i+1 < len(Words) && strings.HasSuffix(Words[i+1], ")") {
					num := strings.TrimSuffix(Words[i+1], ")")
					x, err := strconv.Atoi(num)
					if err != nil || x < 0 {
						fmt.Println("Error : ", Words[i+1], " invalid flag number")
						continue
					}
					start := i - x - (counter - 1)
					if start < 0 {
						start = 0
					}
					for j := start; j < i; j++ {

						y, err := strconv.ParseInt(Words[j], 2, 64)
						if err != nil {
							fmt.Println("Error : ", Words[j], " this binary value is incorrect, it can't be converted")
							continue
						} else {
							Words[j] = strconv.FormatInt(y, 10)
						}

					}
					Words = append(Words[:i], Words[i+2:]...)
					i--

				}
			case Words[i] == "(bin)":

				if i-counter >= 0 {
					y, err := strconv.ParseInt(Words[i-counter], 2, 64)
					if err != nil {
						fmt.Println("Error : ", Words[i-counter], " this binary value is incorrect, it can't be converted")
						continue
					} else {
						Words[i-counter] = strconv.FormatInt(y, 10)
					}
				}

				Words = append(Words[:i], Words[i+1:]...)
				i--
			}
		}
		// handeling a and an if a is upped
		for i := 0; i < len(Words)-1; i++ {
			if len(Words) > 1 && Words[i] == "A" {
				if strings.ContainsAny(string(Words[i+1][0]), "aeoiuhAEOIUH") {
					Words[i] += "N"
				}
			}
			if len(Words) > 1 && len(Words[i]) > 1 && strings.ContainsAny(string(Words[i][len(Words[i])-1]), "A") && strings.ContainsAny(string(Words[i][:len(Words[i])-1]), ",;:?!'(\"") {
				if strings.ContainsAny(string(Words[i+1][0]), "aeoiuhAEOIUH") {
					Words[i] += "N"
				}
			}
		}

		// handle a and an

		Processed = append(Processed, strings.Join(Words, " "))

	}
	return Processed
}

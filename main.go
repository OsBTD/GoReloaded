package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	Module "go-reloaded/InitialProcessing"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Error : Not enough arguments provided")
		return
	} else if len(args) > 2 {
		fmt.Println("Error : too many arguments")
	}
	FileName := os.Args[1]
	ResultName := os.Args[2]
	// input and output files must have a .txt extension
	if filepath.Ext(FileName) != ".txt" || filepath.Ext(ResultName) != ".txt" {
		log.Fatal("Error : input and output files must have a .txt extension")
	}
	// input and output must have different names
	if FileName == ResultName {
		fmt.Println("Error : input and output files must have different names to avoid overwriting")
		return
	}
	// goes from input file to flag processing which inclues a and an handeling then to punctuation handeling to single quotes handeling
	ProcessedFlags := Module.FlagProcessing()
	ProcessedPunctuations := Module.PunctuationProcessing(ProcessedFlags)
	ProcessedQuotes := Module.SingleQuotesProcessing(ProcessedPunctuations)
	// creating result file and write the result
	resfile, err := os.Create(ResultName)
	if err != nil {
		log.Fatal("Error creating file : ", err)
	}
	defer resfile.Close()

	_, err = io.WriteString(resfile, string(ProcessedQuotes))
	if err != nil {
		log.Fatal("Error writing to file : ", err)
	}
	fmt.Println("File processed successfully :) ")
}

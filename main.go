package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	DefaultSourceLanguage    = "en"
	DefaultTranslateLanguage = "fa"
)

type Inputs struct {
	SourceLanguage    string
	TranslateLanguage string
	Words             []string
}

func main() {
	inputs := Inputs{
		DefaultSourceLanguage,
		DefaultTranslateLanguage,
		[]string{},
	}

	fmt.Print("Enter source language (default is '" + DefaultSourceLanguage + "'): ")
	fmt.Scanln(&inputs.SourceLanguage)
	fmt.Print("Enter translate language (default is '" + DefaultTranslateLanguage + "'): ")
	fmt.Scanln(&inputs.TranslateLanguage)

	fmt.Println("Enter a comma-separated string containing the words to translate: ")
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	inputs.Words = strings.Split(line, ",")
}

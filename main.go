package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"translator/translate"
)

const (
	DefaultSourceLanguage    = "en"
	DefaultTranslateLanguage = "fa"
)

func main() {
	translator := translate.Translator{
		Sl:          DefaultSourceLanguage,
		Tl:          DefaultTranslateLanguage,
		SourceWords: []string{},
	}

	// get user inputs
	fmt.Print("Enter source language (default is '" + DefaultSourceLanguage + "'): ")
	fmt.Scanln(&translator.Sl)
	fmt.Print("Enter translate language (default is '" + DefaultTranslateLanguage + "'): ")
	fmt.Scanln(&translator.Tl)
	fmt.Println("Enter a comma-separated string containing the sourceWords to translate: ")
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	translator.SourceWords = strings.Split(line, ",")

	// translate
	results, err := translator.Translate()
	if err != nil {
		log.Fatal("Translation failed, ", err)
	}

	fmt.Println(results)
}

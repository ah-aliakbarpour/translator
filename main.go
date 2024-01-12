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
	var (
		sourceLanguage    string
		translateLanguage string
		sourceWords       []string
	)

	// get user inputs
	fmt.Print("Enter source language (default is '" + DefaultSourceLanguage + "'): ")
	fmt.Scanln(&sourceLanguage)
	fmt.Print("Enter translate language (default is '" + DefaultTranslateLanguage + "'): ")
	fmt.Scanln(&translateLanguage)
	fmt.Println("Enter a comma-separated string containing the sourceWords to translate: ")
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	sourceWords = strings.Split(line, ",")

	// translate
	translatedWords, err := translate.Translate(sourceLanguage, translateLanguage, sourceWords)
	if err != nil {
		log.Fatal("Translation failed, ", err)
	}
	fmt.Println(translatedWords)
}

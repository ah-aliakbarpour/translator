package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"translator/dictionary"
	"translator/export"
)

const (
	DefaultSourceLanguage    = "en"
	DefaultTranslateLanguage = "fa"
)

func main() {
	translator := dictionary.GoogleTranslate{
		SourceLanguage:    DefaultSourceLanguage,
		TranslateLanguage: DefaultTranslateLanguage,
	}

	// get user inputs
	fmt.Print("Enter source language (default is '" + DefaultSourceLanguage + "'): ")
	fmt.Scanln(&translator.SourceLanguage)
	fmt.Print("Enter translate language (default is '" + DefaultTranslateLanguage + "'): ")
	fmt.Scanln(&translator.TranslateLanguage)
	fmt.Println("Enter sourceWords in a comma-separated string: ")
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	sourceWords := strings.Split(line, ",")

	// translate
	results, err := translator.Translate(sourceWords)
	if err != nil {
		log.Fatal("Translation failed, ", err)
	}

	// export excel
	exporter := export.NewExcel("translate", 'Z')
	err = exporter.Export(results)
	if err != nil {
		log.Fatal("Export failed, ", err)
	}

	fmt.Printf("\nDone!\n")
}

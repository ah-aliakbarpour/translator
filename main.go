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
	translator := dictionary.Glosbe{
		Sl:          DefaultSourceLanguage,
		Tl:          DefaultTranslateLanguage,
		SourceWords: []string{},
	}

	// get user inputs
	fmt.Print("Enter source language (default is '" + DefaultSourceLanguage + "'): ")
	fmt.Scanln(&translator.Sl)
	fmt.Print("Enter translate language (default is '" + DefaultTranslateLanguage + "'): ")
	fmt.Scanln(&translator.Tl)
	fmt.Println("Enter sourceWords in a comma-separated string: ")
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	translator.SourceWords = strings.Split(line, ",")

	// translate
	results, err := translator.Translate()
	if err != nil {
		log.Fatal("Translation failed, ", err)
	}

	// export excel
	exporter := export.ExcelExporter{
		Data: results,
	}
	err = exporter.Export()
	if err != nil {
		log.Fatal("Export failed, ", err)
	}
}

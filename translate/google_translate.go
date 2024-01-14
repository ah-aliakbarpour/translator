package translate

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"strings"
	"time"
)

type GoogleTranslator struct {
	Sl, Tl      string
	SourceWords []string
}

const Domain = "https://translate.google.com"
const TranslatedElementClass = "ryNqvb"

// Translate scrapes translated data from Google Translate
func (translator *GoogleTranslator) Translate() ([]Result, error) {

	service, err := startSeleniumService()
	if err != nil {
		return nil, err
	}

	driver, err := setupSelenium()
	if err != nil {
		return nil, err
	}

	words, err := translator.scrapeTranslatedWords(driver)
	if err != nil {
		return nil, err
	}

	err = stopSeleniumService(service)
	if err != nil {
		return nil, err
	}

	return words, nil
}

// scrapeTranslatedWords iterates through SourceWords and stores the translation in result array
func (translator *GoogleTranslator) scrapeTranslatedWords(driver selenium.WebDriver) ([]Result, error) {
	var results []Result

	total := len(translator.SourceWords)
	for i := 0; i < total; i++ {
		sourceWord := strings.ToLower(strings.TrimSpace(translator.SourceWords[i]))
		if len(sourceWord) == 0 {
			continue
		}

		for {
			// visit the target page
			url := Domain + "/?sl=" + translator.Sl + "&tl=" + translator.Tl + "&text=" + sourceWord
			err := driver.Get(url)
			if err != nil {
				return nil, fmt.Errorf("can't visit the target page: %w", err)
			}

			// wait 10 second for the translated element except reload the page
			err = driver.WaitWithTimeout(func(driver selenium.WebDriver) (bool, error) {
				translatedElement, _ := driver.FindElement(selenium.ByClassName, TranslatedElementClass)

				if translatedElement != nil {
					return true, nil
				}
				return false, nil
			}, 10*time.Second)
			if err == nil {
				break
			}
		}

		// find the translated elements
		translatedElements, err := driver.FindElements(selenium.ByClassName, TranslatedElementClass)
		if err != nil {
			log.Println("error finding element for'"+sourceWord+"':", err.Error())
		}

		result := Result{Source: sourceWord}

		for _, element := range translatedElements {
			// extract the text of element
			translatedText, err := element.Text()
			if err != nil {
				log.Println("error extracting text for'"+sourceWord+"':", err.Error())
				continue
			}

			// add the scraped data to the results
			result.Translations = append(result.Translations, translatedText)
		}
		results = append(results, result)
		fmt.Printf("\n[%.0f%%] %v: %v\n\n", float64(i+1)/float64(total)*100, result.Source, result.Translations)
	}

	return results, nil
}

func startSeleniumService() (*selenium.Service, error) {
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		return nil, fmt.Errorf("can't initialize a Chrome browser: %w", err)
	}

	return service, nil
}

func stopSeleniumService(service *selenium.Service) error {
	err := service.Stop()
	if err != nil {
		return fmt.Errorf("can't stop selenium service: %w", err)
	}
	return nil
}

func setupSelenium() (selenium.WebDriver, error) {
	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless", // comment out this line for testing
	}})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return driver, fmt.Errorf("can't create a new remote client: %w", err)
	}

	// maximize the current window to avoid responsive rendering
	err = driver.MaximizeWindow("")
	if err != nil {
		return driver, fmt.Errorf("can't maximize the window: %w", err)
	}

	return driver, nil
}

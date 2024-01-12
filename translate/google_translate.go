package translate

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"strings"
)

const Domain = "https://translate.google.com"
const TranslatedElementClass = "ryNqvb"

// googleTranslate scrapes translated data
func (translator *Translator) googleTranslate() (Result, error) {

	service, err := startSeleniumService()
	if err != nil {
		return nil, err
	}

	driver, err := setupSelenium()
	if err != nil {
		return nil, err
	}

	words, err := translator.getTranslatedWords(driver)
	if err != nil {
		return nil, err
	}

	err = stopSeleniumService(service)
	if err != nil {
		return nil, err
	}

	return words, nil
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

// getTranslatedWords iterates through SourceWords and stores the translation in result array
func (translator *Translator) getTranslatedWords(driver selenium.WebDriver) (Result, error) {
	results := make(Result)

	for _, sourceWord := range translator.SourceWords {
		sourceWord = strings.TrimSpace(sourceWord)

		// visit the target page
		url := Domain + "/?sl=" + translator.Sl + "&tl=" + translator.Tl + "&text=" + sourceWord
		err := driver.Get(url)
		if err != nil {
			return nil, fmt.Errorf("can't visit the target page: %w", err)
		}

		// wait for the translated element to load
		err = driver.Wait(func(driver selenium.WebDriver) (bool, error) {
			translatedElement, _ := driver.FindElement(selenium.ByCSSSelector, "."+TranslatedElementClass)

			if translatedElement != nil {
				return true, nil
			}
			return false, nil
		})
		if err != nil {
			return nil, fmt.Errorf("can't wait for translated element: %w", err)
		}

		results[sourceWord] = []string{}

		// find the translated elements
		translatedElements, err := driver.FindElements(selenium.ByCSSSelector, "."+TranslatedElementClass)
		if err != nil {
			log.Println("error finding element for'"+sourceWord+"':", err.Error())
			continue
		}

		for _, element := range translatedElements {
			// extract the text of element
			translatedText, err := element.Text()
			if err != nil {
				log.Println("error extracting text for'"+sourceWord+"':", err.Error())
				continue
			}

			// add the scraped data to the results
			results[sourceWord] = append(results[sourceWord], translatedText)
		}
	}

	return results, nil
}
package main

import (
	"fmt"
	"github.com/fedesog/webdriver"
	"log"
	"time"
)

func main() {
	chromeDriver := webdriver.NewChromeDriver("chromedriver")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}

	desired := webdriver.Capabilities{"Platform": "Linux"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		log.Println(err)
	}

	err = session.Url("http://google.pl")
	if err != nil {
		log.Println(err)
	}

	element, _ := session.FindElement(webdriver.CSS_Selector, "[name='q']")
	fmt.Printf("%v", element)

	element.SendKeys("Wedriver without selenium :)")
	element, err = session.FindElement(webdriver.CSS_Selector, "[name='btnG']")
	if err != nil {
		log.Println(err)
	}
	element.Click()

	time.Sleep(10 * time.Second)
	session.Delete()
	chromeDriver.Stop()

}

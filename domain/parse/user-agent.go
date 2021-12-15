package parse

import (
	"github.com/PuerkitoBio/goquery"
)

var browserList = []string{
	"Internet Explorer",
	"Opera",
	"Chrome",
	"Safari",
	"FireFox",
}

// UserAgentCom return a string list
func UserAgentCom(doc *goquery.Document) ([]string, error) {
	var newBrowserList = make([]string, 1)

	doc.Find("div#liste ul li").Each(func(i int, selection *goquery.Selection) {
		userAgent := selection.Find("a").Text()
		newBrowserList = append(newBrowserList, userAgent)
		//fmt.Println(userAgent)
		newBrowserList = append(newBrowserList, userAgent)
	})

	return newBrowserList, nil
}
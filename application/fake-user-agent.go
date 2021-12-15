package application

import (
	"errors"
	"fmt"
	"log"

	"github.com/lixvyang/fakeuseragent/domain/global"
	"github.com/lixvyang/fakeuseragent/infra/download"

	"github.com/lixvyang/fakeuseragent/domain/parse"

	"math/rand"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// FakeUserAgent ...
type FakeUserAgent struct {
	UserAgentStringOk	bool
	Cache	bool
}

var (
	ErrUserAgent = errors.New("user agent err")
)

func NewFakeUserAgent(UserAgentStringOk bool, CacheOK bool) *FakeUserAgent {
	return &FakeUserAgent{
		UserAgentStringOk: UserAgentStringOk,
		Cache:             CacheOK,
	}
}

func (f *FakeUserAgent) common(browserType string) string {
	if !f.valid() {
		log.Println("UserAgentStringOk or CacheOk should be true")
		return "None"
	}

	r := rand.NewSource(time.Now().Unix())
	randomChoice := rand.New(r)
	if f.Cache {
		index := randomChoice.Intn(len(global.LOCALUSERAGENT[browserType]))
		return global.LOCALUSERAGENT[browserType][index]
	}

	var url string
	if f.UserAgentStringOk {
		url = fmt.Sprintf(global.BROWSER_BASE_PAGE, browserType)
	}
	var (
		doc *goquery.Document
		err error
	)

	doc, err = download.ResponseDownload(url)
	if err != nil {
		fmt.Println(ErrUserAgent)
		panic(ErrUserAgent)
	}

	var (
		userAgentList []string
	)

	if f.UserAgentStringOk {
		userAgentList, err = parse.UserAgentCom(doc)
		if err != nil {
			fmt.Println(ErrUserAgent)
			panic(ErrUserAgent)
		}
		return userAgentList[randomChoice.Intn(len(userAgentList))]
	}
	return ""
}

func (f *FakeUserAgent) valid() bool {
	// UserAgentStringOk
	// Cache
	// must one true
	return (f.UserAgentStringOk || f.Cache)
}

// IE UserAgent
func (f *FakeUserAgent) IE() string {
	return f.common("Internet+Explorer")
}

// InternetExplorer Usernet
func (f *FakeUserAgent) InternetExplorer() string {
	return f.IE()
}

# fakeuseragent
Simple useragent faker with real world database

### 0. Install
```go
go get github.com/lixvyang/fakeuseragent/application
```

### 1.Feature
Fetch User-Agent from the following URL
- [https://www.w3schools.com/browsers/default.asp](https://www.w3schools.com/browsers/default.asp)
- [https://fake-useragent.herokuapp.com/browsers/0.1.8](https://fake-useragent.herokuapp.com/browsers/0.1.8)


### 3. Usage
```go
package main

import (
  "fmt"
  "github.com/lixvyang/fakeuseragent/application"
)

func main() {
  fakeUserAgent := application.NewFakeUserAgent(true, false)
	fmt.Println(fakeUserAgent.Random())
	fmt.Println(fakeUserAgent.Safari())
	fmt.Println(fakeUserAgent.Chrome())
	fmt.Println(fakeUserAgent.IE())
	fmt.Println(fakeUserAgent.Opera())
}
```

Output:
```
Mozilla/5.0 (iPad; U; CPU OS 3_2_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B500 Safari/53
Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_3; nb-no) AppleWebKit/525.18 (KHTML, 
like Gecko) Version/3.1.1 Safari/525.20
Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/534.13 (KHTML, like Gecko) Chrome/9.0.599.0 Safari/534.13
Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0; YComp 5.0.2.6; Hotbar 3.0)
Mozilla/4.0 (compatible; MSIE 5.0; Windows 98) Opera 6.04  [pl]
```

#### 3.1 Method
- Random get UserAgent
- IE/Msie/InternetExplorer return IE Browser UserAgent
- FF/FireFox return FireFox Browser UserAgent
- Google/Chrome return Chrome Browser UserAgent
- Opera return Opera Browser UserAgent

#### 3.2 Param
- UserAgentStringOk: http://useragentstring.com/
- CacheOk : Random 50 UserAgent

Reference：[Python fake-useragent](https://github.com/hellysmile/fake-useragent)

### License
MIT ©[Lixvyang](LICENSE)
package onu

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"golang.org/x/net/html/charset"
)

type ONU struct {
	Url_ip string
	cookie []*http.Cookie
	client *http.Client
}

// Constructor of the ONU object
func Constructor() (*ONU, error) {

	jar, err := cookiejar.New(nil)
	if err != nil {
		//TODO: Implentar LOG
		return nil, err
	}
	client := &http.Client{
		Timeout: time.Duration(10) * time.Second,
		Jar:     jar,
	}
	onu := ONU{
		client: client,
	}
	return &onu, nil
}

// Funciton Test
func (onu ONU) Test() {
	fmt.Println("Test")
}

func (onu *ONU) Login() {

	response, err := onu.client.PostForm(URL, Payload)
	if err != nil {
		//TODO: Implentar LOG
		panic(err)
	}
	defer response.Body.Close()
	onu.cookie = response.Cookies()
}

func (onu *ONU) PrintCookies() {
	for _, cookie := range onu.cookie {
		fmt.Printf("Cookie: %s=%s\n", cookie.Name, cookie.Value)
	}
}

func (onu ONU) AddCookies(req *http.Request) {

	for _, cookie := range onu.cookie {
		req.AddCookie(&http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}
}

func (onu ONU) AddHeader(req *http.Request) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", onu.cookie[0].Name+"="+onu.cookie[0].Value)
}

func (onu ONU) Page() {

	request, err := http.NewRequest("GET",
		URL_LOGIN,
		nil)
	if err != nil {
		//TODO: Implentar LOG
		panic(err)
	}
	//onu.AddCookies(request)
	onu.AddHeader(request)
	fmt.Println(request.Header["Cookie"])

	_, err = onu.client.Do(request)
	if err != nil {
		//TODO: Implentar LOG
		panic(err)
	}

	// response, err := Client.Do(request)
	// if err != nil {
	// 	panic(err)
	// }
	// defer response.Body.Close()

	// utf8Body, err := charset.NewReader(response.Body, response.Header.Get("Content-Type"))
	// if err != nil {
	// 	panic(err)
	// }
	// bytes, err := io.ReadAll(utf8Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(bytes))

}

// func (onu *ONU) Login() {

// 	response, err := onu.client.PostForm(URL, Payload)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()
// 	onu.cookie = response.Cookies()
// 	fmt.Println(onu.cookie)
// }

func (onu ONU) Cookies(cookie []*http.Cookie) {
	for _, cookie := range cookie {
		fmt.Printf("Cookie: %s=----=%s\n", cookie.Name, cookie.Value)
	}
}

func (onu ONU) SetHeader(r *http.Request) {

}

func (onu ONU) setCookies(r *http.Request) {
	for _, cookie := range onu.cookie {
		value := cookie.Name + "=" + cookie.Value
		fmt.Println(value)
		r.Header.Add("Cookie", value)
	}
}

func (onu ONU) GetPage() {

	request, err := http.NewRequest("GET",
		URL_LOGIN,
		nil)
	if err != nil {
		panic(err)
	}

	onu.setCookies(request)

	response, err := Client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	utf8Body, err := charset.NewReader(response.Body, response.Header.Get("Content-Type"))
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(utf8Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}

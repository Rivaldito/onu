package onu

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html/charset"
)

type ONU struct {
	URL_IP string
	cookie []*http.Cookie
}

func (onu *ONU) Login() {

	response, err := Client.PostForm(URL, Payload)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	onu.cookie = response.Cookies()
}

func (onu ONU) Cookies() {
	for _, cookie := range onu.cookie {
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

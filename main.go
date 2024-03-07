package main

import onu "github.com/rivaldito/onu/ONU"

func main() {

	onu := onu.ONU{}
	onu.Login()
	onu.GetPage()
	// URL := "http://192.168.101.1/cgi-bin/login.cgi"

	// client := &http.Client{}

	// payload := url.Values{
	// 	"onSubmit":       {"1"},
	// 	"randomkey_id":   {" "},
	// 	"login_language": {"English"},
	// 	"encryPassword":  {"f5b18d4744a0f3618765cdac9bed8a0c"},
	// 	"username":       {"InterAdmin"}}

	// request, err := http.NewRequest("POST",
	// 	URL,
	// 	strings.NewReader(payload.Encode()))
	// if err != nil {
	// 	panic(err)
	// }

	// request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// resp, err := client.Do(request)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// cookies := resp.Cookies()
	// for _, cookie := range cookies {
	// 	fmt.Printf("Cookie: %s=----=%s\n", cookie.Name, cookie.Value)
	// }

	// payload := url.Values{
	// 	"onSubmit":       {"1"},
	// 	"randomkey_id":   {" "},
	// 	"login_language": {"English"},
	// 	"encryPassword":  {"f5b18d4744a0f3618765cdac9bed8a0c"},
	// 	"username":       {"InterAdmin"}}

	// resp, err := client.PostForm(URL, payload)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// fmt.Println(resp.Cookies()[0].String())

}

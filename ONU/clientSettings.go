package onu

import "net/url"

//Paylod for login
var Payload = url.Values{
	"onSubmit":       {"1"},
	"randomkey_id":   {" "},
	"login_language": {"English"},
	//"encryPassword":  {"f5b18d4744a0f3618765cdac9bed8a0c"},
	"encryPassword": {"1887d79411a4c97c42e1494359f8eca8"},
	"username":      {"InterAdmin"}}

//

var URL = "http://190.142.136.6/cgi-bin/login.cgi"
var URL_LOGIN = "http://190.142.136.6/cgi-bin/main.cgi"

package onu

import "net/url"

//Paylod for login
var Payload = url.Values{
	"onSubmit":       {"1"},
	"randomkey_id":   {" "},
	"login_language": {"English"},
	"encryPassword":  {"f5b18d4744a0f3618765cdac9bed8a0c"},
	"username":       {"InterAdmin"}}

//

var URL = "http://192.168.101.1/cgi-bin/login.cgi"
var URL_LOGIN = "http://192.168.101.1/cgi-bin/main.cgi"

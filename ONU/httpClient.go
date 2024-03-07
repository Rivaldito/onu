package onu

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

var (
	jar, _ = cookiejar.New(nil)

	Client = &http.Client{
		Timeout: time.Duration(10) * time.Second,
		Jar:     jar,
	}
)

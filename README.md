# GoCaptchaClients
A go library for multiple captcha clients. 

**Supports Capmonster, Anti-Captcha and Capsolver**

Usage: 
```go
package main

import (
	"fmt"

	captcha "github.com/Aran404/GoCaptchaClients"
)

func main() {
	capmonster, err := captcha.NewClient("API_KEY", "capmonster")

	if err != nil {
		panic(err)
	}

	err = capmonster.CreateTask(captcha.HCaptchaTask{
		Type:    "HCaptchaTaskProxyless",
		Url:     "https://discord.com/register",
		SiteKey: "f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34",
	})

	if err != nil {
		panic(err)
	}

	response := capmonster.GetTaskResult(60)

	fmt.Println(response.HcaptchaToken)
}

```

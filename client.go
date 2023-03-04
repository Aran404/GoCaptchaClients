package captchaclients

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Mutable Constants
var (
	ValidServices = map[string][]string{
		"capmonster":   {"capmonster", "monster", "capmonster.cloud", "api.capmonster.cloud"},
		"anti-captcha": {"anti", "anti-captcha", "anticaptcha", "anti-captcha.com", "api.anti-captcha.com"},
		"capsolver":    {"capsolver", "capsolver.com", "api.capsolver.com"},
	}
	ApiUrls = map[string]string{
		"capmonster":   "api.capmonster.cloud",
		"anti-captcha": "api.anti-captcha.com",
		"capsolver":    "api.capsolver.com",
	}
)

func CheckValidService(service string) string {
	for services, aliases := range ValidServices {
		for _, name := range aliases {
			if name == service {
				return services
			}
		}
	}

	return ""
}

func NewClient(apiKey, service string) (*Instance, error) {
	validate := CheckValidService(strings.ToLower(service))

	if validate == "" {
		return nil, fmt.Errorf("%v is not a captcha service we support. Please choose another one.", service)
	}

	return &Instance{
		Client:  &http.Client{Timeout: time.Second * 60},
		ApiKey:  apiKey,
		Service: ApiUrls[validate],
	}, nil
}

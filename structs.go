package captchaclients

import (
	"net/http"
	"time"
)

type Instance struct {
	ApiKey  string
	Service string
	TaskId  string
	Client  *http.Client
}

type HCaptchaTask struct {
	Type              string `json:"type"`
	UserAgent         string `json:"userAgent,omitempty"`
	Url               string `json:"websiteURL"`
	SiteKey           string `json:"websiteKey"`
	Invisble          bool   `json:"isInvisible,omitempty"`
	Proxy             string `json:"proxy,omitempty"`
	EnterprisePayload *struct {
		RqData string `json:"rqdata,omitempty"`
	} `json:"enterprisePayload,omitempty"`
}

type CHCaptchaTask struct {
	Type          string `json:"type"`
	UserAgent     string `json:"userAgent,omitempty"`
	Url           string `json:"websiteURL"`
	SiteKey       string `json:"websiteKey"`
	Invisble      bool   `json:"isInvisible,omitempty"`
	RqData        string `json:"rqdata,omitempty"`
	ProxyType     string `json:"proxyType,omitempty"`
	ProxyAddress  string `json:"proxyAddress,omitempty"`
	ProxyPort     int    `json:"proxyPort,omitempty"`
	ProxyLogin    string `json:"proxyLogin,omitempty"`
	ProxyPassword string `json:"proxyPassword,omitempty"`
}

type CreateTaskPayload struct {
	ClientKey string      `json:"clientKey"`
	Task      interface{} `json:"task"`
}

type GetResultPayload struct {
	ClientKey string      `json:"clientKey"`
	TaskId    interface{} `json:"taskId"`
}

type GetTaskResultResponse struct {
	HcaptchaToken string
	SolveTime     time.Duration
	ErrorId       string
	Error         error
}

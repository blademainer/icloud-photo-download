package icloud

import (
	"fmt"
	"github.com/blademainer/icloud-photo-download/pkg/icloud/config"
	"github.com/go-resty/resty/v2"
	"log"
)

const (
	HttpContentTypeJSON = "application/json"
)

const (
	ApiURL   = "https://icloud.com.cn"
	LoginURL = "https://idmsa.apple.com/appleauth/auth/signin?isRememberMeEnabled=true"
)

type LoginRequest struct {
	AccountName string   `json:"accountName"`
	RememberMe  bool     `json:"rememberMe"`
	Password    string   `json:"password"`
	TrustTokens []string `json:"trustTokens"`
}

type LoginResponse struct {
	AccountName string   `json:"accountName"`
	RememberMe  bool     `json:"rememberMe"`
	Password    string   `json:"password"`
	TrustTokens []string `json:"trustTokens"`
}

type ICloud struct {
	config *config.Config
	client *resty.Client
}

// NewICloud create icloud
func NewICloud(config *config.Config) (*ICloud, error) {
	i := &ICloud{
		config: config,
		client: resty.New().SetDebug(config.Debug),
	}
	return i, nil
}

func (c *ICloud) Login() {
	get, err := c.client.R().
		Get("https://setup.icloud.com.cn/setup/ws/1/validate?clientBuildNumber=2301Hotfix20&clientMasteringNumber=2301Hotfix20&clientId=e98cc553-326f-4875-a405-611d87aad2c7")
	if err != nil {
		log.Printf("err: %v", err.Error())
		return
	}
	fmt.Printf("header: %s", get.Header())
	r2, err := c.client.R().SetBody(
		&LoginRequest{
			AccountName: c.config.User.UserName,
			RememberMe:  true,
		},
	).
		Post("https://idmsa.apple.com/appleauth/auth/federate?isRememberMeEnabled=true")
	if err != nil {
		log.Printf("err: %v", err.Error())
		return
	}
	// if r3.StatusCode() != http.StatusOK {
	// 	log.Println(r3.Header())
	// 	return
	// 	// panic(string(r3.Body()))
	// }
	fmt.Printf("header: %s", r2.Header())

	r3, err := c.client.R().SetBody(
		&LoginRequest{
			AccountName: c.config.User.UserName,
			Password:    c.config.User.Password,
			RememberMe:  true,
		},
	).Post("https://idmsa.apple.com/appleauth/auth/signin?isRememberMeEnabled=true")
	if err != nil {
		log.Printf("err: %v", err.Error())
		return
	}
	log.Printf("login code: %v", r3.Status())
	fmt.Printf("header: %s", r3.Header())
	sessionToken := r3.Header().Get("X-Apple-Session-Token")

	r4, err := c.client.R().SetBody(
		map[string]interface{}{
			"accountCountryCode": "CHN",
			"dsWebAuthToken":     sessionToken,
			"extended_login":     false,
		},
	).Post("https://setup.icloud.com.cn/setup/ws/1/accountLogin?clientBuildNumber=2301Hotfix20&clientMasteringNumber=2301Hotfix20&clientId=e98cc553-326f-4875-a405-611d87aad2c7")
	if err != nil {
		log.Printf("err: %v", err.Error())
		return
	}
	log.Printf("login code: %v", r4.Status())
	fmt.Printf("header: %s", r4.Header())

	// r4, err := c.client.R().Get("https://idmsa.apple.com/appleauth/auth")
	// if err != nil {
	// 	log.Printf("err: %v", err.Error())
	// 	return
	// }
	// if r4.StatusCode() != http.StatusOK {
	// 	panic(string(r4.Body()))
	// }
	// fmt.Printf("header: %s", r4.Header())
	// "https://idmsa.apple.com/appleauth/auth"

	// req := &LoginRequest{
	// 	AccountName: c.config.User.UserName,
	// 	RememberMe:  true,
	// 	Password:    c.config.User.Password,
	// 	TrustTokens: nil,
	// }
	// resp, err := c.client.R().
	// 	EnableTrace().
	// 	SetHeader("Origin", "https://idmsa.apple.com").
	// 	SetHeader("Referer", "https://idmsa.apple.com").
	// 	SetHeader(
	// 		"X-Apple-I-FD-Client-Info",
	// 		"{\"U\":\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36\",\"L\":\"zh-CN\",\"Z\":\"GMT+08:00\",\"V\":\"1.1\",\"F\":\"Fta44j1e3NlY5BNlY5BSmHACVZXnN9.6JjKHJ_9_Cyhk6Hb9LarUqUdHz16rgNNlejV2pNk0ug49RJdmcK9rTJfwrKyNjjNklY5BNleBBNlYCa1nkBMfs.CQV\"}",
	// 	).
	// 	SetBody(req).
	// 	Post(LoginURL)
	// if err != nil {
	// 	log.Printf("err: %v", err.Error())
	// 	return
	// }
	// log.Printf("header: %v", resp.Header())
	// // log.Printf("body: %s", resp.Body())
	// token := resp.Header().Get("X-Apple-Session-Token")
	// log.Printf("token: %v", token)

}

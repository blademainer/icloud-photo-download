package icloud

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/blademainer/icloud-photo-download/pkg/icloud/config"
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/go-uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

var (
	CommonHeaders = map[string]string{
		"User-Agent":              "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
		"X-Apple-Widget-Key":      OAuthClientID,
		"X-Apple-OAuth-Client-Id": OAuthClientID,
	}
	OAuthClientID = "d39ba9916b7251055b22c7f910e2ea796ee65e98b2ddecea8f5dde8d9d1a815d"
)

const (
	HttpContentTypeJSON = "application/json"
)

type ICloud struct {
	config   *config.Config
	client   *resty.Client
	clientID string
}

// NewICloud create icloud
func NewICloud(config *config.Config) (*ICloud, error) {
	logrus.SetLevel(logrus.DebugLevel)

	jar, err := NewCookieJar()
	if err != nil {
		return nil, err
	}
	clientID, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}
	i := &ICloud{
		config: config,
		client: resty.New().
			SetDebug(config.Debug).
			SetLogger(logrus.StandardLogger()).
			SetCookieJar(jar),
		clientID: clientID,
	}
	return i, nil
}

func (c *ICloud) Login() error {
	validate, err := c.client.R().
		SetHeader("Origin", "https://icloud.com.cn").
		SetHeader("Referer", "https://icloud.com.cn").
		Get(
			fmt.Sprintf(
				"https://setup.icloud.com.cn/setup/ws/1/validate?clientBuildNumber=2301Hotfix20&clientMasteringNumber=2301Hotfix20&clientId=%v",
				c.clientID,
			),
		)
	if err != nil {
		logrus.Debugf("err: %v", err.Error())
		return err
	}
	logrus.Debugf("validate status: %v", validate.Status())
	logrus.Debugf("validate header: %s", validate.RawResponse.Header)

	r2, err := c.client.R().
		SetHeader("Origin", "https://idmsa.apple.com").
		SetHeader("Referer", "https://idmsa.apple.com").
		SetHeaders(CommonHeaders).
		SetBody(
			&LoginRequest{
				AccountName: c.config.User.UserName,
				RememberMe:  true,
			},
		).
		Post("https://idmsa.apple.com/appleauth/auth/federate?isRememberMeEnabled=true")
	if err != nil {
		logrus.Errorf("err: %v", err.Error())
		return err
	}
	logrus.Debugf("federate status: %v", r2.Status())
	logrus.Debugf("federate header: %s", r2.RawResponse.Header)
	// if signin.StatusCode() != http.StatusOK {
	// 	log.Println(signin.Header())
	// 	return
	// 	// panic(string(signin.Body()))
	// }
	// logrus.Debugf("header: %s", r2.Header())

	signin, err := c.client.R().
		SetHeaders(CommonHeaders).
		SetHeader("Origin", "https://idmsa.apple.com").
		SetHeader("Referer", "https://idmsa.apple.com").
		SetBody(
			&LoginRequest{
				AccountName: c.config.User.UserName,
				Password:    c.config.User.Password,
				RememberMe:  true,
			},
		).
		SetResult(make(map[string]string)).
		Post("https://idmsa.apple.com/appleauth/auth/signin?isRememberMeEnabled=true")
	if err != nil {
		logrus.Errorf("err: %v", err.Error())
		return err
	}
	sessionToken := signin.RawResponse.Header.Get("X-Apple-Session-Token")
	sessionID := signin.RawResponse.Header.Get("X-Apple-Id-Session-Id")
	scnt := signin.RawResponse.Header.Get("Scnt")
	logrus.Debugf("sessionToken: %s", sessionToken)
	logrus.Debugf("sessionToken: %s", sessionToken)
	logrus.Debugf("scnt: %s", scnt)
	logrus.Debugf("signin status: %v", signin.Status())
	logrus.Debugf("sign header: %s", signin.RawResponse.Header)
	logrus.Debugf("sign resp: %v", signin.RawResponse)
	logrus.Debugf("sign resp: %s", signin.Result())

	if sessionToken == "" {
		return fmt.Errorf("can't get session token")
	}

	// authReq, err := c.client.R().
	// 	SetHeader("Origin", "https://idmsa.apple.com").
	// 	SetHeader("Referer", "https://idmsa.apple.com").
	// 	SetHeader("X-Apple-Id-Session-Id", sessionID).
	// 	SetHeaders(CommonHeaders).
	// 	Get("https://idmsa.apple.com/appleauth/auth")
	// if err != nil {
	// 	logrus.Errorf("err: %v", err.Error())
	// 	return err
	// }
	// logrus.Debugf("auth status: %v", authReq.Status())
	// logrus.Debugf("auth header: %s", authReq.RawResponse.Header)
	// if authReq.StatusCode() != http.StatusOK {
	// 	return fmt.Errorf("auth failed")
	// }

	userInfo := &UserInfo{}
	login, err := c.client.R().
		SetHeaders(CommonHeaders).
		// SetHeader("X-Apple-Id-Session-Id", sessionID).
		SetHeader("Origin", "https://www.icloud.com.cn").
		SetBody(
			map[string]interface{}{
				"accountCountryCode": "CHN",
				"dsWebAuthToken":     sessionToken,
				"extended_login":     false,
			},
		).
		SetResult(userInfo).
		Post(
			fmt.Sprintf(
				"https://setup.icloud.com.cn/setup/ws/1/accountLogin?clientBuildNumber=2301Hotfix20&clientMasteringNumber=2301Hotfix20&clientId=%v",
				c.clientID,
			),
		)
	if err != nil {
		logrus.Errorf("err: %v", err.Error())
		return err
	}
	userStr, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}
	logrus.Debugf("login status: %v", login.Status())
	logrus.Debugf("user info: %v", string(userStr))
	logrus.Debugf("user info: %v", login.Result())

	if !userInfo.DsInfo.HsaEnabled && userInfo.DsInfo.HsaVersion <= 1 {
		return nil
	}

	fmt.Println("input security code: ")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	if err != nil {
		return err
	}

	securityCodeStr := strings.TrimSpace(string(input))
	logrus.Debugf("security code: %v", securityCodeStr)
	securityCode, err := c.client.R().
		SetHeaders(CommonHeaders).
		SetHeader("Origin", "https://idmsa.apple.com").
		SetHeader("Referer", "https://idmsa.apple.com").
		SetHeader("X-Apple-ID-Session-Id", sessionID).
		SetHeader("X-Apple-OAuth-Redirect-URI", "https://www.icloud.com.cn").
		SetHeader("X-Apple-OAuth-Client-Type", "firstPartyAuth").
		SetHeader("X-Apple-OAuth-Require-Grant-Code", "true").
		SetHeader("X-Apple-OAuth-Response-Mode", "web_message").
		SetHeader("X-Apple-OAuth-Response-Type", "code").
		SetHeader("X-Requested-With", "XMLHttpRequest").
		SetHeader("X-Apple-Domain-Id", "6").
		SetHeader("X-Apple-Frame-Id", "auth-7hlrtbcm-4qxg-yrgs-yfkw-34xkak3k").
		SetHeader("scnt", scnt).
		SetBody(
			map[string]interface{}{
				"securityCode": map[string]interface{}{
					"code": securityCodeStr,
				},
			},
		).
		Post("https://idmsa.apple.com/appleauth/auth/verify/trusteddevice/securitycode")
	logrus.Debugf("security code status: %v", securityCode.Status())
	if securityCode.StatusCode() >= http.StatusBadRequest || securityCode.StatusCode() < http.StatusOK {
		logrus.Errorf(
			"failed to verify security code: %v resp: %s", securityCode.Status(),
			securityCode.Body(),
		)
		return fmt.Errorf("failed to verify security code")
	}
	logrus.Infof(
		"login status: %v result: %v", securityCode.Status(),
		securityCode.Body(),
	)

	// fmt.Printf("header: %s", login.Header())

	// login, err := c.client.R().Get("https://idmsa.apple.com/appleauth/auth")
	// if err != nil {
	// 	logrus.Debugf("err: %v", err.Error())
	// 	return
	// }
	// if login.StatusCode() != http.StatusOK {
	// 	panic(string(login.Body()))
	// }
	// fmt.Printf("header: %s", login.Header())
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
	// 	logrus.Debugf("err: %v", err.Error())
	// 	return
	// }
	// logrus.Debugf("header: %v", resp.Header())
	// // logrus.Debugf("body: %s", resp.Body())
	// token := resp.Header().Get("X-Apple-Session-Token")
	// logrus.Debugf("token: %v", token)
	return nil
}

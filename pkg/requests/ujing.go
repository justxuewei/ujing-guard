package requests

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/xavier-niu/ujing-guard/pkg/conf"
	"net/http"
)

const (
	apiPrefix  = "https://phoenix.ujing.online/api/v1"
	captchaApi = apiPrefix + "/captcha"
	// header keys
	mobileBrandKey    = "x-mobile-brand"
	acceptEncodingKey = "Accept-Encoding"
	acceptLanguageKey = "Accept-Language"
	acceptKey         = "Accept"
	appCodeKey        = "x-app-code"
	mobileIdKey       = "x-mobile-id"
	userAgentKey      = "User-Agent"
	appVersionKey     = "x-app-version"
	mobileModelKey    = "x-mobile-model"
)

// errors
var ServiceError = errors.New("ujing service request error")

var basicHeader = map[string]string{
	mobileBrandKey:    conf.DeviceConfig.MobileBrand,
	acceptEncodingKey: "gzip, deflate, br",
	acceptLanguageKey: "zh-Hans-CN;q=1, en-CN;q=0.9",
	acceptKey:         "*/*",
	appCodeKey:        conf.UJingAppConfig.AppCode,
	mobileIdKey:       conf.DeviceConfig.MobileId,
	userAgentKey:      conf.UJingAppConfig.UserAgent,
	appVersionKey:     conf.UJingAppConfig.AppVersion,
	mobileModelKey:    conf.DeviceConfig.MobileModel,
}

type uJIngJsonResponse struct {
	code    int
	message string
	data    interface{}
}

func NewUJingSession() *UJingSession {
	return &UJingSession{client: &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}}
}

type UJingSession struct {
	token  string
	client *http.Client
}

func (s *UJingSession) Captcha() error {
	//phone := conf.UserConfig.Phone

	resp, _ := http.Get("https://www.aliyun.com")
	fmt.Println(resp)

	//req, err := http.NewRequest("GET", "https://www.nxw.name", nil)
	//if err != nil {
	//	return err
	//}
	//// headers
	//for key, value := range basicHeader {
	//	req.Header.Set(key, value)
	//}
	//// url queries
	//q := req.URL.Query()
	//q.Add("mobile", phone)
	//q.Add("sessionId", "AFS_SWITCH_OFF")
	//q.Add("sig", "AFS_SWITCH_OFF")
	//q.Add("token", "AFS_SWITCH_OFF")
	//q.Add("type", "1")
	//req.URL.RawQuery = q.Encode()

	//resp, err := s.client.Do(req)
	//if err != nil {
	//	return err
	//}
	//if resp.StatusCode != 200 {
	//	util.Log().Error("The captcha request is not OK: %s.", resp.Status)
	//	return ServiceError
	//}
	//
	//defer resp.Body.Close()
	//
	//respBody := &uJIngJsonResponse{}
	//err = json.NewDecoder(resp.Body).Decode(respBody)
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Println(respBody)
	//if respBody.code != 0 {
	//	util.Log().Error("The captcha request returns an error: %s.", respBody.message)
	//	return ServiceError
	//}
	//
	return nil
}

package conf

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestInitNoPanic(t *testing.T) {
	asserts := assert.New(t)
	testCase := `
phone: 13888888888
mobileBrand: apple
appCode: ZI
mobileId: A58AEFED-83D2-38DB-DCB2-BC5B892FD9D4
userAgent: U jing/2.1.12 (iPhone; iOS 14.2.1; Scale/3.00)
appVersion: 2.1.12
mobileModel: iPhone13,2
location: 19.23,14.21
deviceNames:
 - device1
 - device2
refreshInterval: 1
`
	err := ioutil.WriteFile("testConf.yaml", []byte(testCase), 0644)
	defer func() {
		_ = os.Remove("testConf.yaml")
		_ = os.Remove(".ujing_guard_cache_test_case")
	}()
	if err != nil {
		panic(err)
	}
	asserts.NotPanics(func() {
		Init("testConf.yaml", ".ujing_guard_cache_test_case")
	})
}

func TestValidParseConf(t *testing.T) {
	asserts := assert.New(t)
	testCase := `
phone: 13888888888
mobileBrand: apple
appCode: ZI
mobileId: A58AEFED-83D2-38DB-DCB2-BC5B892FD9D4
userAgent: U jing/2.1.12 (iPhone; iOS 14.2.1; Scale/3.00)
appVersion: 2.1.12
x-mobile-model: iPhone13,2
location: 19.23,14.21
deviceNames:
 - device1
 - device2
refreshInterval: 1
`
	err := ioutil.WriteFile("testConf.yaml", []byte(testCase), 0644)
	defer func() {
		_ = os.Remove("testConf.yaml")
		_ = os.Remove(".ujing_guard_cache_test_case")
	}()
	if err != nil {
		panic(err)
	}
	Init("testConf.yaml", ".ujing_guard_cache_test_case")
	err = parseConf("testConf.yaml")
	asserts.NoError(err)

	t.Log(UserConfig)
	t.Log(StoreConfig)
	t.Log(AppConfig)
	asserts.Equal("13888888888", UserConfig.Phone)
}

func TestParseConfMissingPhone(t *testing.T) {
	asserts := assert.New(t)
	testCase := `
mobileBrand: apple
appCode: ZI
mobileId: A58AEFED-83D2-38DB-DCB2-BC5B892FD9D4
userAgent: U jing/2.1.12 (iPhone; iOS 14.2.1; Scale/3.00)
appVersion: 2.1.12
x-mobile-model: iPhone13,2
refreshInterval: 1
`
	err := ioutil.WriteFile("testConf.yaml", []byte(testCase), 0644)
	defer func() {
		_ = os.Remove("testConf.yaml")
		_ = os.Remove(".ujing_guard_cache_test_case")
	}()
	if err != nil {
		panic(err)
	}
	//Init("testConf.yaml", ".ujing_guard_cache_test_case")
	err = parseConf("testConf.yaml")
	asserts.Error(err)
	t.Logf("err: %s\n", err)
}

func TestParseConfWithInvalidPhoneFormat(t *testing.T) {
	asserts := assert.New(t)
	testCase := `
phone: 1388888888
mobileBrand: apple
appCode: ZI
mobileId: A58AEFED-83D2-38DB-DCB2-BC5B892FD9D4
userAgent: U jing/2.1.12 (iPhone; iOS 14.2.1; Scale/3.00)
appVersion: 2.1.12
x-mobile-model: iPhone13,2
refreshInterval: 1
`
	err := ioutil.WriteFile("testConf.yaml", []byte(testCase), 0644)
	defer func() {
		_ = os.Remove("testConf.yaml")
		_ = os.Remove(".ujing_guard_cache_test_case")
	}()
	if err != nil {
		panic(err)
	}
	err = parseConf("testConf.yaml")
	asserts.Error(err)
	t.Logf("err: %s\n", err)
}

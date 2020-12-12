package conf

import (
	"github.com/go-playground/validator/v10"
	"github.com/xavier-niu/ujing-guard/pkg/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var CachePath string

//region Public Conf Struct
var UserConfig = &user{}

var DeviceConfig = &device{
	MobileBrand: "apple",
	MobileId:    "A58AEFED-83D2-38DB-DCB2-BC5B892FD9D4",
	MobileModel: "iPhone13,2",
}

var AppConfig = &app{
	AppVersion: "iPhone13,2",
	AppCode:    "ZI",
	UserAgent:  "U jing/2.1.12 (iPhone; iOS 14.2.1; Scale/3.00)",
}

var StoreConfig = &store{}
//endregion

//region Conf Entity Struct Definition
type user struct {
	Phone string `yaml:"phone" validate:"required,number,len=11"`
}

type device struct {
	MobileBrand string `yaml:"mobileBrand"`
	MobileId    string `yaml:"mobileId"`
	MobileModel string `yaml:"mobileModel"`
}

type app struct {
	AppVersion string `yaml:"appVersion"`
	AppCode    string `yaml:"appCode"`
	UserAgent  string `yaml:"userAgent"`
}

type store struct {
	Location    string   `yaml:"location" validate:"required"`
	DeviceNames []string `yaml:"deviceNames" validate:"required"`
}
//endregion

func Init(confPath string, cachePath string) {
	if confPath == "" || !util.Exists(confPath) {
		util.Log().Panic("Config file is required, but nothing is found at \"%s\".", confPath)
	}
	if cachePath == "" {
		util.Log().Panic("Cache file should not be an empty string.")
	}
	if !util.Exists(cachePath) {
		// create a cache file if the file is not existed
		f, err := util.CreatNestedFile(cachePath)
		if err != nil {
			util.Log().Panic("Creating cache file is failed: %s", err)
		}
		f.Close()
	}

	CachePath = cachePath

	err := parseConf(confPath)
	if err != nil {
		util.Log().Panic("Parsing conf is failed: %s.", err)
	}
}

func parseConf(confPath string) error {
	f, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}

	sections := []interface{}{UserConfig, DeviceConfig, AppConfig, StoreConfig}
	validate := validator.New()

	for _, section := range sections {
		err = yaml.Unmarshal(f, section)
		if err != nil {
			return err
		}
		err = validate.Struct(section)
		if err != nil {
			return err
		}
	}

	return nil
}

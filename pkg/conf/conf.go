package conf

import "github.com/xavier-niu/ujing-guard/pkg/util"

var (
	UserConfig user
	CachePath string
)

type user struct {
	Phone       string `yaml:"phone" validate:"required"`
	MobileBrand string `yaml:"mobileBrand" validate:"required"`
	AppCode     string `yaml:"appCode" validate:"required"`
	MobileId    string `yaml:"mobileId" validate:"required"`
	UserAgent   string `yaml:"userAgent" validate:"required"`
	AppVersion  string `yaml:"appVersion" validate:"required"`
	MobileModel string `yaml:"mobileModel" validate:"required"`
}

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

}

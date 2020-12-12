package main

import (
	"flag"
	"fmt"
	"github.com/xavier-niu/ujing-guard/pkg/conf"
	"github.com/xavier-niu/ujing-guard/pkg/util"
)

var (
	confPath string
	cachePath string
)

func init() {
	flag.StringVar(&confPath, "conf", util.RelativePath("conf.yaml"), "the path of conf")
	flag.StringVar(&cachePath, "cache", util.RelativePath(".ujing_guard_cache"), "the path of cache")
	flag.Parse()
	conf.Init(confPath, cachePath)
}

func main() {
	fmt.Println("11111")
}

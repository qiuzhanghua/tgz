package main

import (
	"github.com/qiuzhanghua/common/log"
	"os"
)

func init() {
	level := os.Getenv("LOGGING_LEVEL")
	log.SetGlobalLevel(level)
}

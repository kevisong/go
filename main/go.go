package main

import (
	"github.com/KEVISONG/go/common/log"
)

func init() {
	log.Init(&log.Config{Level: "info", Formatter: "text", ReportCaller: true})
}

func main() {
}

package main

import (
	"github.com/KEVISONG/go/common/log"
)

func init() {
	log.Init(&log.Config{Level: "info", Formatter: "json", ReportCaller: true})
}

func main() {

}

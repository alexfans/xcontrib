package xcontrib

import "log"

func Debug(v ...interface{}) {
	log.Println(v...)
}

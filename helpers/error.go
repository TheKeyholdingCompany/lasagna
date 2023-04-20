package helpers

import (
	"log"
	"runtime/debug"
)

func CheckError(e error) {
	if e != nil {
		debug.PrintStack()
		log.Fatalln(e.Error())
	}
}

func CheckWarn(e error) {
	if e != nil {
		log.Println(e.Error())
	}
}

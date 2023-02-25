package utils

import (
	"encoding/json"
	"log"
	"runtime"
)

func PrintObject(object interface{}) {
	_, file, line, _ := runtime.Caller(1)
	indent, err := json.MarshalIndent(object, "", " ")
	if err == nil {
		log.Printf("%s:%d %s\n", file, line, indent)
	}
}

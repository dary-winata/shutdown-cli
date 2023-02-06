package src

import (
	"os"
)

func Run() {
	paramString := os.Args[1:]
	value := ParamToTime(paramString)
	Logic(value)
}

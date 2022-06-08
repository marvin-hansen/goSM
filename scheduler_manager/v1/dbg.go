package v1

import "log"

func dbgPrint(msg string) {
	if dbg {
		log.Println(main + msg)
	}
}

func nilCheck(value interface{}, msg string) {
	if value == nil {
		println(msg)
		panic(value)
	}
}

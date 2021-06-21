package tools

import "log"

func Unwrap(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

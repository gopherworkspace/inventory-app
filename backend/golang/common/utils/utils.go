package utils

import (
	"log"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

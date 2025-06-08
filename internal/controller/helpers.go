package controller

import (
	"log"
	"strconv"
)

func ParseId64(s string) (int64, error) {
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Println("strconv.ParseInt error: ", err)
	}
	return id, err
}

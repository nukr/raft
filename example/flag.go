package main

import (
	"errors"
	"log"
	"net"
	"strconv"
	"strings"
)

type peers []string

func (j *peers) String() string {
	var ss string
	for i, s := range *j {
		if i == len(*j)-1 {
			ss += s
		} else {
			ss += s + ","
		}
	}
	return ss
}

func (j *peers) Set(value string) error {
	// Check value whether valid IP:PORT
	valueSlice := strings.Split(value, ":")
	if len(valueSlice) != 2 {
		return errors.New("not valid string format must be IP:PORT")
	}
	ip := net.ParseIP(valueSlice[0])
	if ip == nil {
		return errors.New("not valid IP")
	}
	_, err := strconv.ParseUint(valueSlice[1], 10, 16)
	if err != nil {
		log.Fatalln(err)
	}
	*j = append(*j, value)
	return nil
}

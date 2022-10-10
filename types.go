package main

import "errors"

var errSome = errors.New("some error")

type Response struct {
	Status   string      `json:"status"`
	Error    string      `json:"error,omitempty"`
	Endpoint string      `json:"endpoint"`
	Data     interface{} `json:"data"`
}

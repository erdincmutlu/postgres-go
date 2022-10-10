package main

import "errors"

var errSome = errors.New("some error")

type Response struct {
	Id   int         `json:"id"`
	Data interface{} `json:"data"`
}

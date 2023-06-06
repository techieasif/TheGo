package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurfCSRF(t *testing.T) {
	var myH myHandler
	h := NoSurfCSRF(&myH)

	switch v := h.(type) {
	case http.Handler:
	//do nothing
	default:
		fmt.Println(fmt.Sprintf("h is not Handler type, but instead it is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)
	switch v := h.(type) {
	case http.Handler:
	//do nothing
	default:
		fmt.Println(fmt.Sprintf("h is not Handler type, but instead it is %T", v))
	}
}

package demoapi

import (
	"fmt"
	"net/http"
)

func Demo2API(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Demo 2 API")
}

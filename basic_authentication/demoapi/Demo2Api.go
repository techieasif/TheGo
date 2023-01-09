package demoapi

import (
	"fmt"
	"net/http"
)

func Route2(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Demo 2 API")
}

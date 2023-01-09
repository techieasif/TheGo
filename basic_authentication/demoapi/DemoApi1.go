package demoapi

import (
	"fmt"
	"net/http"
)

func Route1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Demo 1 API")

}

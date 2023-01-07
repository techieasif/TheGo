package demoapi

import (
	"fmt"
	"net/http"
)

func Demo1API(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Demo 1 API")

}

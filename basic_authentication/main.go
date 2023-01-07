package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/techieasif/TheGo/basic_authentication/basicauthmiddleware"
	"github.com/techieasif/TheGo/basic_authentication/demoapi"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/api/demo/demo1", basicauthmiddleware.BasicAuthMiddleware(demoapi.Demo1API)).Methods("GET")

	router.HandleFunc("/api/demo/demo2", demoapi.Demo2API).Methods("GET")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}

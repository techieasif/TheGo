package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals-quarters", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors-suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	//{"reservation-summary", "/reservation-summary", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "POST", []postData{
		{"first_name", "Karan"},
		{"last_name", "Arjun"},
		{"email", "karan@arjun.com"},
		{"phone", "222-222-2222"},
	}, http.StatusOK},
	{"search-availability", "/search-availability", "POST", []postData{
		{
			"start",
			"2022-09-02",
		},
		{
			"end",
			"2022-09-02",
		},
	}, http.StatusOK},
	{"search-availability-json", "/search-availability-json", "POST", []postData{
		{
			"start",
			"2022-09-02",
		},
		{
			"end",
			"2022-09-02",
		},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				log.Fatal(err.Error())
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("Method %s, has %d, expected was %d", e.method, resp.StatusCode, e.expectedStatusCode)
			}
			resp.Body.Close()

		} else {

			values := url.Values{}

			for _, v := range e.params {
				values.Add(v.key, v.value)
			}

			resp, err := ts.Client().PostForm(ts.URL+e.url, values)

			if err != nil {
				t.Log(err)
				log.Fatal(err.Error())
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("Method %s, has %d, expected was %d", e.method, resp.StatusCode, e.expectedStatusCode)
			}
			resp.Body.Close()
		}
	}

}

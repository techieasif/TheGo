package render

import (
	"github.com/techieasif/TheGo/bookings/models"
	"net/http"
	"testing"
)

func TestAddDefaultTemplateData(t *testing.T) {
	var templateData models.TemplateData

	resp, e := getSession()
	if e != nil {
		t.Error(e)
	}

	session.Put(resp.Context(), "flash", "123")

	result := AddDefaultTemplateData(&templateData, resp)

	if result.Flash != "123" {
		t.Error("RESULT FAILED..")
	}
}

func TestParseTemplateV2(t *testing.T) {
	pathToTemplates = "./../../templates"

	tc, err := CreateTemplateCacheV2()

	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc

	r, err := getSession()

	if err != nil {
		t.Error(err)
	}

	ww := myWriter{}

	err = ParseTemplateV2(&ww, "home.page.tmpl", &models.TemplateData{}, r)
	if err != nil {
		t.Error(err)
	}
	err = ParseTemplateV2(&ww, "non0.page.tmpl", &models.TemplateData{}, r)
	if err == nil {
		t.Error(err)
	}

}

func getSession() (*http.Request, error) {
	resp, err := http.NewRequest("GET", "/dummy-url", nil)

	if err != nil {
		return nil, err
	}
	ctx := resp.Context()
	ctx, _ = session.Load(ctx, resp.Header.Get("X-Session"))
	resp = resp.WithContext(ctx)

	return resp, nil
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCacheV2(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCacheV2()

	if err != nil {
		t.Error(err)
	}
}

type myWriter struct {
}

func (m *myWriter) Header() http.Header {
	var header http.Header
	return header
}

func (m *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

func (m *myWriter) WriteHeader(status int) {

}

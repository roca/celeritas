package render

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRender_Page(t *testing.T) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	testRender.Renderer = "go"
	testRender.RootPath = "./testdata"

	err = testRender.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page",err)
	}

	if strings.Contains(w.Body.String(), "Hello World") == false {
		t.Error("Body does not match expected output")
	}

	testRender.Renderer = "jet"
	err = testRender.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page",err)
	}
}

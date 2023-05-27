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
		t.Error("Error rendering page", err)
	}

	if strings.Contains(w.Body.String(), "Hello World") == false {
		t.Error("Body does not match expected output")
	}

	err = testRender.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Rendered template that does not exist")
	}

	testRender.Renderer = "jet"
	err = testRender.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

	err = testRender.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Rendered Jet template that does not exist")
	}

	testRender.Renderer = "unknown"
	err = testRender.Page(w, r, "home", nil, nil)
	if err == nil {
		t.Error("No error return while rendering with unknown renderer")
	}
}

func TestRender_GoPage(t *testing.T) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	testRender.Renderer = "go"
	testRender.RootPath = "./testdata"

	err = testRender.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}
}

func TestRender_JetPage(t *testing.T) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	testRender.Renderer = "Jet"
	testRender.RootPath = "./testdata"

	err = testRender.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}
}

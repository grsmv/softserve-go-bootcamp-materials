package main

import (
	"net/http/httptest"
	"testing"
)

var ResponseReceivier = httptest.NewRecorder()
var testPage = Page{Title: "TestTitle", Body: []byte("Test Content")}
var testView = "View Title: TestTitle, Body: Test Content\n"
var testEdit = "Edit Title: TestTitle, Body: Test Content\n"

func TestViewTemplateRendering(t *testing.T) {
	defer func() { ResponseReceivier.Body.Reset() }()
	staticRender := DirectTemplate{
		path: "test_suites",
	}
	staticRender.renderTemplate(ResponseReceivier, "view", &testPage)
	expect(t, ResponseReceivier.Body.String(), testView, "Direct rendering view according testing template")
}
func TestViewCompiledTemplateRendering(t *testing.T) {
	defer func() { ResponseReceivier.Body.Reset() }()

	render := NewCTemplates("test_suites")
	render.renderTemplate(ResponseReceivier, "view", &testPage)
	expect(t, ResponseReceivier.Body.String(), testView, "Compiled rendering edit according testing template")

}
func TestEditTemplateRendering(t *testing.T) {
	defer func() { ResponseReceivier.Body.Reset() }()
	staticRender := DirectTemplate{
		path: "test_suites",
	}
	staticRender.renderTemplate(ResponseReceivier, "edit", &testPage)
	expect(t, ResponseReceivier.Body.String(), testEdit, "Direct rendering editaccording testing template")
}

func TestEditCompiledTemplateRendering(t *testing.T) {
	defer func() { ResponseReceivier.Body.Reset() }()
	render := NewCTemplates("test_suites")
	render.renderTemplate(ResponseReceivier, "edit", &testPage)
	expect(t, ResponseReceivier.Body.String(), testEdit, "Compiled rendering edit according testing template")
}

func BenchmarkEditTemplateRendering(b *testing.B) {
	defer func() { ResponseReceivier.Body.Reset() }()
	staticRender := DirectTemplate{
		path: "test_suites",
	}
	for n := 0; n < b.N; n++ {
		staticRender.renderTemplate(ResponseReceivier, "edit", &testPage)
	}
}

func BenchmarkEditCompiledTemplateRendering(b *testing.B) {
	defer func() { ResponseReceivier.Body.Reset() }()
	render := NewCTemplates("test_suites")
	for n := 0; n < b.N; n++ {
		render.renderTemplate(ResponseReceivier, "edit", &testPage)
	}
}

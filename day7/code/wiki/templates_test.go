package main

import (
	"net/http/httptest"
	"testing"
)

var ResponseReceivier = httptest.NewRecorder()
var testPage = Page{Title: "TestTitle", Body: []byte("Test Content")}
var testView = "View Title: TestTitle, Body: Test Content"
var testEdit = "View Title: TestTitle, Body: Test Content"

func TestViewTemplateRendering(t *testing.T) {
	renderTemplate(ResponseReceivier, "view", &testPage)
	expect(t, ResponseReceivier.Body.String(), testView, "Rendering view according testing template")
}

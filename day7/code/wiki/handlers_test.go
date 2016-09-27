package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server *httptest.Server
)

type testCase struct {
	uri         string
	status      int
	contentType string
	content     string
	description string
}

func init() {
	templates = DirectTemplate{
		path: "test_suites",
	}
	server = httptest.NewServer(nil)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
}

func TestRedirectOnUnexistingPage(t *testing.T) {
	tc := testCase{
		uri:         "/view/newpage",
		description: "view unexisting page",
	}
	req, err := http.NewRequest("GET", server.URL+tc.uri, nil)
	if err != nil {
		t.Fatalf("Can't create request %s", err)
	}
	hc := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("Redirect")
		},
	}
	rsp, err := hc.Do(req)
	if err != nil {
		t.Errorf("Err: %s\n\n", err)
	}
	expect(t, rsp.StatusCode, tc.status, tc.description+" status")
}

func TestRedirectToEditPage(t *testing.T) {
	tc := testCase{
		uri:         "/view/newpage",
		description: "view unexisting page",
	}
	req, err := http.NewRequest("GET", server.URL+tc.uri, nil)
	if err != nil {
		t.Fatalf("Can't create request %s", err)
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Err: %s\n\n", err)
	}

	//*/
	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("Can't read response's body %s ", err)
	}
	t.Logf("Response: %+v\n\n", string(buf))
	if err != nil {
		t.Fatalf("Can't make request %s", err)
	}
	//*/
	expect(t, rsp.StatusCode, tc.status, tc.description+" status")
}

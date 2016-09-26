package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testIdGenHandlerFunc(idgen idGenerator, t *testing.T) {
	server := httptest.NewServer(newIdGenHandler(idgen))

	defer server.Close()

	for i := 1; i <= 10; i++ {
		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
		}
		expected := fmt.Sprint(i)
		actual, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if expected != string(actual) {
			t.Errorf("Expected the message %q, recieved %q\n", expected, string(actual))
		}
	}
}

func TestIdGenHandlerAtomic(t *testing.T) {
	testIdGenHandlerFunc(newIdGeneratorAtomic(), t)
}

func TestIdGenHandlerMutex(t *testing.T) {
	testIdGenHandlerFunc(newIdGeneratorMutex(), t)
}

func TestIdGenHandlerChan(t *testing.T) {
	testIdGenHandlerFunc(newIdGeneratorChan(), t)
}

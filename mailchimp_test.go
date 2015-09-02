package mailchimp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func equals(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestBadAPIKey(t *testing.T) {
	_, err := NewClient("asz", nil)
	if err == nil {
		t.Fail()
	}
}

func TestURL(t *testing.T) {
	client, _ := NewClient("a-lit11", nil)
	equals(t, "https://lit11.api.mailchimp.com/3.0", client.BaseURL.String())
}

func TestSubscribeError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(500)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client, _ := NewClient("a-lit11", &http.Client{Transport: transport})
	client.BaseURL, _ = url.Parse("http://localhost/")
	_, err := client.Subscribe("me@matthewbrown.io", "abc_test")
	if err == nil {
		t.Fatal(err)
	}
	equals(t, err.Error(), "Error 0  ()")
}

func TestSubscribe(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{"email":"bob@example.com","status":"sent","reject_reason":"hard-bounce","_id":"1"}`)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client, _ := NewClient("a-lit11", &http.Client{Transport: transport})
	client.BaseURL, _ = url.Parse("http://localhost/")
	_, err := client.Subscribe("me@matthewbrown.io", "abc_test")
	if err != nil {
		t.Fatal(err)
	}
}

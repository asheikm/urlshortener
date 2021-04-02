package api_test

import "net/http/httptest"
import "testing"
import "net/http"
import "os"
import "urlshortener/api"

type TestGetHandler struct{}
type TestPostHandler struct{}

func (h *TestGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.GetShortenedURL(w, r)
}

func (h *TestPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        api.CreateShortenedURL(w, r)
}

func TestGetDomain(t *testing.T) {
	// Make a test request
	resp, err := http.Get(os.Getenv("SHORT_DOMAIN") + os.Getenv("LISTEN_PORT") + "/")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
}

func TestPostReq(t *testing.T) {
        h := &TestGetHandler{}
        server := httptest.NewServer(h)
        defer server.Close()

        // Make a test request
        resp, err := http.Get(os.Getenv("SHORT_DOMAIN"))
        if err != nil {
                t.Fatal(err)
        }
        if resp.StatusCode != 200 {
                t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
        }
}




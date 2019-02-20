package multiplerequests_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"

	. "github.com/jszroberto/go-algorithms/concurrency/multiplerequests"
)

func TestMaxInMultipleRequests(t *testing.T) {
	tests := []struct {
		Name    string
		Version func([]string) (int, error)
	}{
		{"Simple", MaxInMultipleRequestsSimple},
		{"WithErrors", MaxInMultipleRequestsWithErrors},
		{"WithCancellation", MaxInMultipleRequestsWithCancellation},
		{"WithErrorGroup", MaxInMultipleRequestsWithErrorGroup},
	}

	t.Log("When given different versions of the algorithm")
	{
		for _, tt := range tests {
			tf := func(t *testing.T) {
				var counter int
				var mtx sync.Mutex
				handler := func(w http.ResponseWriter, r *http.Request) {
					mtx.Lock()
					defer mtx.Unlock()
					counter++
					counter %= 6
					w.WriteHeader(200)
					w.Write([]byte(strconv.Itoa(counter)))
				}

				server := httptest.NewServer(http.HandlerFunc(handler))
				numRequests := 100
				urls := make([]string, numRequests)
				for i := 0; i < numRequests; i++ {
					urls[i] = server.URL
				}
				expected := 5
				got, err := tt.Version(urls)
				if err != nil {
					t.Fatalf("got unexpected error %s", err)
				}

				if got != expected {
					t.Errorf("test failed expected %d, got %d\n", got, expected)
				}
			}
			t.Run(tt.Name, tf)
		}
	}
}
